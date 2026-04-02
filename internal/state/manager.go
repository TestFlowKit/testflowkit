// Package state manages the lock file – a JSON key/value store
// that persists acquired auth tokens across runs.
//
// Entries are keyed by the SHA-256 hash of the fully-resolved SecurityScheme
// (see internal/security/hash.go), so a change in any environment variable
// that feeds a scheme automatically invalidates its cached token.
//
// Thread safety
//   - In-process: protected by sync.RWMutex.
//   - Cross-process: a sidecar <lockfile>.pid file is acquired with O_EXCL
//     before writes and released afterwards, preventing concurrent runners
//     from corrupting the file.
package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	lockFileSchemaVersion = 1
	// safetyBuffer is subtracted from the reported expiry so tokens are never
	// used in the last few seconds of their lifetime.
	safetyBuffer = 30 * time.Second
	LockFile     = "testflowkit.lock"
)

// TokenEntry holds the persisted state for one acquired token.
type TokenEntry struct {
	// AccessToken is the raw bearer / access token value.
	AccessToken string `json:"access_token"` //nolint:gosec // Lock state intentionally persists token material.
	// TokenType is e.g. "Bearer" (for display / injection).
	TokenType string `json:"token_type,omitempty"`
	// ObtainedAt is when the token was fetched (UTC).
	ObtainedAt time.Time `json:"obtained_at"`
	// ExpiresAt is the absolute expiry time reported by the IDP (UTC).
	// Zero means no expiry information is available.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// SchemeHash is the key under which this entry is stored (redundant but
	// useful for debugging lock-file contents by hand).
	SchemeHash string `json:"scheme_hash"`
}

// IsExpired reports whether the token should be considered expired,
// applying the 30-second safety buffer to avoid race conditions.
func (e *TokenEntry) IsExpired() bool {
	if e.ExpiresAt.IsZero() {
		return false
	}
	return time.Now().UTC().After(e.ExpiresAt.Add(-safetyBuffer))
}

// lockFileContent is the on-disk representation of lock file.
type lockFileContent struct {
	Version   int                    `json:"version"`
	UpdatedAt time.Time              `json:"updated_at"`
	Entries   map[string]*TokenEntry `json:"entries"`
}

// Manager is the runtime handle to lock file.
// Create one per process via NewManager, then call Load before the test suite
// starts and Save after it ends.
type Manager struct {
	mu      sync.RWMutex
	path    string
	pidPath string
	entries map[string]*TokenEntry
}

// NewManager creates a Manager that will use lockFilePath for persistence.
// The file does not need to exist yet; it will be created on the first Save.
func NewManager(lockFilePath string) *Manager {
	return &Manager{
		path:    lockFilePath,
		pidPath: lockFilePath + ".pid",
		entries: make(map[string]*TokenEntry),
	}
}

// Load reads lock file from disk into memory.
// If the file does not exist the manager starts with an empty state (not an error).
// If the file is corrupted it is silently discarded and an empty state is used.
func (m *Manager) Load() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, err := os.ReadFile(m.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil // fresh start
		}
		return fmt.Errorf("state: read lock file: %w", err)
	}

	var content lockFileContent
	decodeErr := json.Unmarshal(data, &content)
	if decodeErr != nil {
		// Corrupted file – discard and start fresh.
		m.entries = make(map[string]*TokenEntry)
		return nil
	}

	if content.Entries != nil {
		m.entries = content.Entries
	}
	return nil
}

// Save persists the current in-memory state to lock file.
// It acquires a cross-process file lock before writing and releases it after.
func (m *Manager) Save() error {
	m.mu.RLock()
	entries := make(map[string]*TokenEntry, len(m.entries))
	for k, v := range m.entries {
		entries[k] = v
	}
	m.mu.RUnlock()

	if err := m.acquireFileLock(); err != nil {
		return err
	}
	defer m.releaseFileLock()

	content := lockFileContent{
		Version:   lockFileSchemaVersion,
		UpdatedAt: time.Now().UTC(),
		Entries:   entries,
	}

	data, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return fmt.Errorf("state: marshal lock file: %w", err)
	}

	const perm = 0600
	writeErr := os.WriteFile(m.path, data, perm)
	if writeErr != nil {
		return fmt.Errorf("state: write lock file: %w", writeErr)
	}
	return nil
}

// Get returns the TokenEntry for the given scheme hash, or nil if absent or expired.
func (m *Manager) Get(schemeHash string) *TokenEntry {
	m.mu.RLock()
	defer m.mu.RUnlock()
	entry, ok := m.entries[schemeHash]
	if !ok || entry.IsExpired() {
		return nil
	}
	return entry
}

// Put stores (or replaces) a TokenEntry under the given scheme hash.
func (m *Manager) Put(schemeHash string, entry *TokenEntry) {
	m.mu.Lock()
	defer m.mu.Unlock()
	entry.SchemeHash = schemeHash
	m.entries[schemeHash] = entry
}

// Invalidate removes the entry for schemeHash from the in-memory cache,
// forcing the next Get to return nil and trigger a fresh authentication.
// Used by the retry_on_401 flow.
func (m *Manager) Invalidate(schemeHash string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.entries, schemeHash)
}

// CreateEmpty writes a new, empty lock file if none exists yet.
// Called by the init command so users see the file from the first run.
func CreateEmpty(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil // already exists
	}
	content := lockFileContent{
		Version:   lockFileSchemaVersion,
		UpdatedAt: time.Now().UTC(),
		Entries:   make(map[string]*TokenEntry),
	}
	data, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return fmt.Errorf("state: marshal empty lock file: %w", err)
	}
	const perm = 0600
	return os.WriteFile(path, data, perm)
}

// acquireFileLock creates a sidecar <lockfile>.pid file using O_EXCL (atomic
// create-if-not-exists).  It spins with a short back-off for up to 5 seconds
// before giving up to avoid indefinite blocking.
func (m *Manager) acquireFileLock() error {
	const (
		maxWait  = 5 * time.Second
		interval = 50 * time.Millisecond
	)
	deadline := time.Now().Add(maxWait)
	for {
		f, err := os.OpenFile(m.pidPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
		if err == nil {
			_ = f.Close()
			return nil
		}
		if !errors.Is(err, os.ErrExist) {
			return fmt.Errorf("state: acquire file lock: %w", err)
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("state: timed out waiting for lock file %s", m.pidPath)
		}
		time.Sleep(interval)
	}
}

func (m *Manager) releaseFileLock() {
	_ = os.Remove(m.pidPath)
}
