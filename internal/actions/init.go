package actions

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

const (
	featuresDir = "features"
	configFile  = "config.yml"
)

//go:embed boilerplate/config.boilerplate.yml
var configTemplate string

//go:embed boilerplate/sample.boilerplate.feature
var sampleFeatureTemplate string

type InitializationState struct {
	createdFiles []string
	createdDirs  []string
}

func (state *InitializationState) addCreatedFile(path string) {
	state.createdFiles = append(state.createdFiles, path)
}

func (state *InitializationState) addCreatedDir(path string) {
	state.createdDirs = append(state.createdDirs, path)
}

func (state *InitializationState) cleanup() {
	for _, cDir := range state.createdDirs {
		if err := os.RemoveAll(cDir); err != nil {
			if os.IsNotExist(err) {
				continue
			}
			logger.Warn(fmt.Sprintf("Failed to cleanup directory %s: %v", cDir, err), nil)
		}
	}

	for _, file := range state.createdFiles {
		if err := os.Remove(file); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			logger.Warn(fmt.Sprintf("Failed to cleanup file %s: %v", file, err), nil)
		}
	}
}

func validatePath(path string, allowedPrefixes ...string) error {
	cleanPath := filepath.Clean(path)

	if strings.Contains(cleanPath, "..") {
		return fmt.Errorf("invalid path contains directory traversal: %s", path)
	}

	if len(allowedPrefixes) == 0 {
		return nil
	}

	for _, prefix := range allowedPrefixes {
		if cleanPath == prefix || strings.HasPrefix(cleanPath, prefix+string(filepath.Separator)) {
			return nil
		}
	}

	return fmt.Errorf("path %s is not within allowed locations: %v", path, allowedPrefixes)
}

func validateFilePermissions(dir string) error {
	testFile := filepath.Join(dir, ".testflowkit_permission_test")

	file, err := os.Create(testFile)
	if err != nil {
		return fmt.Errorf("insufficient permissions to create files in %s: %w", dir, err)
	}

	file.Close()
	os.Remove(testFile)

	return nil
}

func createDirectoryStructure(state *InitializationState) error {
	if err := validatePath(featuresDir); err != nil {
		logger.Error("Invalid directory path: "+err.Error(), nil, nil)
		return err
	}

	if err := validateFilePermissions("."); err != nil {
		logger.Error("Permission validation failed: "+err.Error(), nil, nil)
		return err
	}

	if _, err := os.Stat(featuresDir); err == nil {
		logger.Warn("Features directory already exists, continuing...", nil)
		return nil
	} else if !os.IsNotExist(err) {
		logger.Error("Failed to check features directory: "+err.Error(), nil, nil)
		return err
	}

	err := os.MkdirAll(featuresDir, 0755)
	if err != nil {
		logger.Error("Failed to create features directory: "+err.Error(), nil, nil)
		return err
	}

	state.addCreatedDir(featuresDir)
	logger.Info("Created features directory")
	return nil
}

func createSampleFeature(state *InitializationState) error {
	sampleFeaturePath := filepath.Join("features", "sample.feature")

	if err := validatePath(sampleFeaturePath, "features"); err != nil {
		logger.Error("Invalid sample feature path: "+err.Error(), nil, nil)
		return err
	}

	if err := validateFilePermissions("features"); err != nil {
		logger.Error("Permission validation failed for features directory: "+err.Error(), nil, nil)
		return err
	}

	if _, err := os.Stat(sampleFeaturePath); err == nil {
		logger.Info("Sample feature file already exists, skipping creation")
		return nil
	} else if !os.IsNotExist(err) {
		logger.Error("Failed to check sample feature file: "+err.Error(), nil, nil)
		return err
	}

	err := os.WriteFile(sampleFeaturePath, []byte(sampleFeatureTemplate), 0600)
	if err != nil {
		logger.Error("Failed to create sample feature file: "+err.Error(), nil, nil)
		return err
	}

	state.addCreatedFile(sampleFeaturePath)
	logger.Info("Created sample feature file: " + sampleFeaturePath)
	return nil
}

func createConfigFile(state *InitializationState) error {
	configPath := configFile

	if err := validatePath(configPath); err != nil {
		logger.Error("Invalid config path: "+err.Error(), nil, nil)
		return err
	}

	if err := validateFilePermissions("."); err != nil {
		logger.Error("Permission validation failed: "+err.Error(), nil, nil)
		return err
	}

	if _, err := os.Stat(configPath); err == nil {
		logger.Error("Config file already exists. Please remove config.yml or run init in a different directory.", nil, nil)
		return errors.New("config.yml already exists")
	} else if !os.IsNotExist(err) {
		logger.Error("Failed to check config file: "+err.Error(), nil, nil)
		return err
	}

	err := os.WriteFile(configPath, []byte(configTemplate), 0600)
	if err != nil {
		logger.Error("Failed to create config file: "+err.Error(), nil, nil)
		return err
	}

	state.addCreatedFile(configPath)
	logger.Info("Created config file: " + configPath)
	return nil
}

func displayGuidance() {
	logger.Success("🎉 TestFlowKit project initialized successfully!")
	logger.Info("")

	logger.Info("📁 Generated files:")
	logger.Info("   ├── config.yml (TestFlowKit configuration)")
	logger.Info("   └── features/sample.feature (Sample test file)")
	logger.Info("")

	logger.Info("🚀 Next steps:")
	logger.Info("   1. Run your first test:")
	logger.Info("      ./tkit run")
	logger.Info("")
	logger.Info("   2. Explore the sample test:")
	logger.Info("      cat features/sample.feature")
	logger.Info("")

	logger.Info("📚 Learn more:")
	logger.Info("   • TestFlowKit Documentation: https://testflowkit.dreamsfollowers.me")
	logger.Info("   • Available Test Sentences: https://testflowkit.dreamsfollowers.me/sentences")
	logger.Info("   • Configuration Guide: https://testflowkit.dreamsfollowers.me/configuration")
	logger.Info("")

	logger.Info("✨ Welcome to TestFlowKit! Your sample test is ready to run against our documentation site.")
	logger.Info("   Happy testing! 🧪")
}

func initMode(_ *config.Config) {
	logger.Info("Initializing TestFlowKit project...")

	state := &InitializationState{
		createdFiles: make([]string, 0),
		createdDirs:  make([]string, 0),
	}

	cleanup := func(err error) {
		logger.Error("Initialization failed: "+err.Error(), nil, nil)
		logger.Info("Cleaning up partially created files...")
		state.cleanup()
		logger.Fatal("Initialization aborted due to errors", err)
	}

	if err := createConfigFile(state); err != nil {
		cleanup(err)
		return
	}

	if err := createDirectoryStructure(state); err != nil {
		cleanup(err)
		return
	}

	if err := createSampleFeature(state); err != nil {
		cleanup(err)
		return
	}

	displayGuidance()
}
