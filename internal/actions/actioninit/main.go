package actioninit

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testflowkit/internal/config"
	"testflowkit/internal/utils/fileutils"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/logger"
)

var validatePath = fileutils.ValidatePath

const featuresDir = "features"
const cursorRulesDir = ".cursor/rules"

//go:embed boilerplate/config.boilerplate.yml
var configTemplate string

//go:embed boilerplate/sample.boilerplate.feature
var sampleFeatureTemplate string

//go:embed boilerplate/testflowkit.agent.boilerplate.yml
var agentConfigTemplate string

//go:embed boilerplate/testflowkit-agent.mdc
var cursorRuleTemplate string

//go:embed boilerplate/copilot-instructions.md
var copilotInstructionsTemplate string

func Execute(_ *config.Config, _ error) {
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

	if err := createAgentConfigFile(state); err != nil {
		cleanup(err)
		return
	}

	if err := createCursorRule(state); err != nil {
		cleanup(err)
		return
	}

	if err := createCopilotInstructions(state); err != nil {
		cleanup(err)
		return
	}

	displayGuidance()
}

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

	const dirPerm = 0755
	err := os.MkdirAll(featuresDir, dirPerm)
	if err != nil {
		logger.Error("Failed to create features directory: "+err.Error(), nil, nil)
		return err
	}

	state.addCreatedDir(featuresDir)
	logger.Info("Created features directory")
	return nil
}

func createSampleFeature(state *InitializationState) error {
	sampleFeaturePath := filepath.Join(featuresDir, "sample.feature")

	if err := validatePath(sampleFeaturePath, featuresDir); err != nil {
		logger.Error("Invalid sample feature path: "+err.Error(), nil, nil)
		return err
	}

	if err := validateFilePermissions(featuresDir); err != nil {
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

	err := os.WriteFile(sampleFeaturePath, []byte(sampleFeatureTemplate), fileutils.FilePermission)
	if err != nil {
		logger.Error("Failed to create sample feature file: "+err.Error(), nil, nil)
		return err
	}

	state.addCreatedFile(sampleFeaturePath)
	logger.Info("Created sample feature file: " + sampleFeaturePath)
	return nil
}

func configFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func createConfigFile(state *InitializationState) error {
	configPath := config.DefaultConfigFile

	if err := validatePath(configPath); err != nil {
		logger.Error("Invalid config path: "+err.Error(), nil, nil)
		return err
	}

	if err := validateFilePermissions("."); err != nil {
		logger.Error("Permission validation failed: "+err.Error(), nil, nil)
		return err
	}

	if configFileExists(configPath) || configFileExists(config.LegacyConfigFile) {
		const instruction = "Please remove testflowkit.yml (or legacy config.yml) or run init in a different directory."
		logger.Error("Config file already exists. "+instruction, nil, nil)
		return fmt.Errorf("%w: %s", apperrors.ErrConfigAlreadyExists, instruction)
	}

	err := os.WriteFile(configPath, []byte(configTemplate), fileutils.FilePermission)
	if err != nil {
		logger.Error("Failed to create config file: "+err.Error(), nil, nil)
		return err
	}

	state.addCreatedFile(configPath)
	logger.Info("Created config file: " + configPath)
	return nil
}

func createAgentConfigFile(state *InitializationState) error {
	agentConfigPath := "testflowkit.agent.yml"

	if err := validatePath(agentConfigPath); err != nil {
		logger.Error("Invalid agent config path: "+err.Error(), nil, nil)
		return err
	}

	if configFileExists(agentConfigPath) {
		logger.Info("Agent config already exists, skipping creation")
		return nil
	}

	if err := os.WriteFile(agentConfigPath, []byte(agentConfigTemplate), fileutils.FilePermission); err != nil {
		logger.Error("Failed to create agent config file: "+err.Error(), nil, nil)
		return err
	}

	state.addCreatedFile(agentConfigPath)
	logger.Info("Created agent config file: " + agentConfigPath)
	return nil
}

func createCursorRule(state *InitializationState) error {
	const dirPerm = 0755

	if err := os.MkdirAll(cursorRulesDir, dirPerm); err != nil {
		logger.Warn(fmt.Sprintf("Could not create %s directory, skipping Cursor rule: %v", cursorRulesDir, err), nil)
		return nil
	}

	rulePath := filepath.Join(cursorRulesDir, "testflowkit-agent.mdc")

	if configFileExists(rulePath) {
		logger.Info("Cursor rule already exists, skipping creation")
		return nil
	}

	if err := os.WriteFile(rulePath, []byte(cursorRuleTemplate), fileutils.FilePermission); err != nil {
		logger.Warn("Failed to create Cursor rule file (non-fatal): "+err.Error(), nil)
		return nil
	}

	state.addCreatedFile(rulePath)
	logger.Info("Created Cursor rule: " + rulePath)
	return nil
}

func createCopilotInstructions(state *InitializationState) error {
	copilotPath := "copilot-instructions.md"

	if configFileExists(copilotPath) {
		logger.Info("copilot-instructions.md already exists, skipping creation")
		return nil
	}

	if err := os.WriteFile(copilotPath, []byte(copilotInstructionsTemplate), fileutils.FilePermission); err != nil {
		logger.Warn("Failed to create copilot-instructions.md (non-fatal): "+err.Error(), nil)
		return nil
	}

	state.addCreatedFile(copilotPath)
	logger.Info("Created VS Code Copilot instructions: " + copilotPath)
	return nil
}

func displayGuidance() {
	logger.Success("🎉 TestFlowKit project initialized successfully!")
	logger.Info("")

	logger.Info("📁 Generated files:")
	logger.Info("   ├── testflowkit.yml          (TestFlowKit runtime configuration)")
	logger.Info("   ├── testflowkit.agent.yml    (IDE agent configuration)")
	logger.Info("   ├── features/sample.feature  (Sample test file)")
	logger.Info("   ├── .cursor/rules/testflowkit-agent.mdc  (Cursor rules)")
	logger.Info("   └── copilot-instructions.md  (VS Code Copilot instructions)")
	logger.Info("")

	logger.Info("🚀 Next steps:")
	logger.Info("   1. Run your first test:")
	logger.Info("      ./tkit run")
	logger.Info("")
	logger.Info("   2. Enable the IDE agent (Cursor):")
	logger.Info("      Add .cursor/mcp.json with the testflowkit-mcp server.")
	logger.Info("      See: https://testflowkit.github.io/testflowkit/docs/getting-started/ide-agent")
	logger.Info("")
	logger.Info("   3. Explore the sample test:")
	logger.Info("      cat features/sample.feature")
	logger.Info("")

	logger.Info("📚 Learn more:")
	logger.Info("   • TestFlowKit Documentation: https://testflowkit.github.io")
	logger.Info("   • Available Test Sentences: https://testflowkit.github.io/testflowkit/sentences")
	logger.Info("   • Configuration Guide: https://testflowkit.github.io/testflowkit/configuration")
	logger.Info("   • IDE Agent Guide: https://testflowkit.github.io/testflowkit/docs/getting-started/ide-agent")
	logger.Info("")

	logger.Info("✨ Welcome to TestFlowKit! Your sample test is ready to run against our documentation site.")
	logger.Info("   Happy testing! 🧪")
}
