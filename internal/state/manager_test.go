package state

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func tempLockFile(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	return filepath.Join(dir, LockFile)
}

func TestNewManager_Empty(t *testing.T) {
	mgr := NewManager(tempLockFile(t))
	if mgr == nil {
		t.Fatal("NewManager returned nil")
	}
}

func TestCreateEmpty_CreatesFile(t *testing.T) {
	path := tempLockFile(t)
	if err := CreateEmpty(path); err != nil {
		t.Fatalf("CreateEmpty failed: %v", err)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("lock file was not created")
	}
}

func TestCreateEmpty_AlreadyExists(t *testing.T) {
	path := tempLockFile(t)
	if err := CreateEmpty(path); err != nil {
		t.Fatalf("first CreateEmpty failed: %v", err)
	}
	// Second call should be silent (return nil, not overwrite).
	if err := CreateEmpty(path); err != nil {
		t.Errorf("second CreateEmpty should be a no-op, got error: %v", err)
	}
}

func TestLoadSave_RoundTrip(t *testing.T) {
	path := tempLockFile(t)
	mgr := NewManager(path)

	entry := &TokenEntry{
		AccessToken: "tok123",
		TokenType:   "Bearer",
		ObtainedAt:  time.Now().UTC().Truncate(time.Second),
		ExpiresAt:   time.Now().UTC().Add(1 * time.Hour).Truncate(time.Second),
		SchemeHash:  "abc123",
	}
	mgr.Put("abc123", entry)

	if err := mgr.Save(); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	mgr2 := NewManager(path)
	if err := mgr2.Load(); err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	got := mgr2.Get("abc123")
	if got == nil {
		t.Fatal("Get returned nil after load")
	}
	if got.AccessToken != "tok123" {
		t.Errorf("expected AccessToken=tok123, got %q", got.AccessToken)
	}
}

func TestGet_ReturnsNilForMissingKey(t *testing.T) {
	mgr := NewManager(tempLockFile(t))
	if mgr.Get("nonexistent") != nil {
		t.Error("expected nil for missing key")
	}
}

func TestGet_ReturnsNilForExpired(t *testing.T) {
	mgr := NewManager(tempLockFile(t))
	entry := &TokenEntry{
		AccessToken: "old",
		ObtainedAt:  time.Now().UTC().Add(-2 * time.Hour),
		ExpiresAt:   time.Now().UTC().Add(-1 * time.Hour), // expired
		SchemeHash:  "h1",
	}
	mgr.Put("h1", entry)
	if mgr.Get("h1") != nil {
		t.Error("expected nil for expired token")
	}
}

func TestInvalidate_RemovesEntry(t *testing.T) {
	mgr := NewManager(tempLockFile(t))
	entry := &TokenEntry{
		AccessToken: "tok",
		ObtainedAt:  time.Now().UTC(),
		ExpiresAt:   time.Now().UTC().Add(1 * time.Hour),
		SchemeHash:  "h2",
	}
	mgr.Put("h2", entry)
	if mgr.Get("h2") == nil {
		t.Fatal("token should exist before invalidation")
	}
	mgr.Invalidate("h2")
	if mgr.Get("h2") != nil {
		t.Error("token should be nil after invalidation")
	}
}

func TestLoad_MissingFile_OK(t *testing.T) {
	mgr := NewManager(tempLockFile(t))
	// File does not exist — should silently succeed (empty cache).
	if err := mgr.Load(); err != nil {
		t.Errorf("Load on missing file should succeed, got %v", err)
	}
}

func TestLoad_CorruptFile_OK(t *testing.T) {
	path := tempLockFile(t)
	if err := os.WriteFile(path, []byte("not valid json{{{{"), 0600); err != nil {
		t.Fatal(err)
	}
	mgr := NewManager(path)
	// Corrupt file — should silently discard and return no error.
	if err := mgr.Load(); err != nil {
		t.Errorf("Load on corrupt file should succeed (silent discard), got %v", err)
	}
	if mgr.Get("any") != nil {
		t.Error("expected empty cache after corrupt load")
	}
}

func TestIsExpired_ZeroTime(t *testing.T) {
	e := &TokenEntry{ExpiresAt: time.Time{}}
	if e.IsExpired() {
		t.Error("zero ExpiresAt should be treated as non-expiring")
	}
}
