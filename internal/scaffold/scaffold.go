package scaffold

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Enriquefft/Mere/internal/embed"
)

// InitProject initializes the .claude/ directory structure with embedded templates
func InitProject(force bool) error {
	claudeDir := ".claude"
	commandsDir := filepath.Join(claudeDir, "commands")

	// Check if .claude/ already exists
	if _, err := os.Stat(claudeDir); !os.IsNotExist(err) {
		if !force {
			return fmt.Errorf(".claude/ directory already exists. Use --force to overwrite")
		}
		fmt.Println("Removing existing .claude/ directory...")
		if err := os.RemoveAll(claudeDir); err != nil {
			return fmt.Errorf("failed to remove existing .claude/: %w", err)
		}
	}

	fmt.Println("Initializing Mere project structure...")

	// Create .claude/ directory
	if err := CreateDirectory(claudeDir); err != nil {
		return err
	}

	// Create commands/ subdirectory
	if err := CreateDirectory(commandsDir); err != nil {
		return err
	}

	// Write CLAUDE.md
	claudeMD, err := embed.GetClaudeMD()
	if err != nil {
		return fmt.Errorf("failed to read CLAUDE.md template: %w", err)
	}
	if err := WriteFile(filepath.Join(claudeDir, "CLAUDE.md"), claudeMD); err != nil {
		return err
	}
	fmt.Println("  Created .claude/CLAUDE.md")

	// Write module.md
	moduleMD, err := embed.GetModuleMD()
	if err != nil {
		return fmt.Errorf("failed to read module.md template: %w", err)
	}
	if err := WriteFile(filepath.Join(commandsDir, "module.md"), moduleMD); err != nil {
		return err
	}
	fmt.Println("  Created .claude/commands/module.md")

	// Write status.md
	statusMD, err := embed.GetStatusMD()
	if err != nil {
		return fmt.Errorf("failed to read status.md template: %w", err)
	}
	if err := WriteFile(filepath.Join(commandsDir, "status.md"), statusMD); err != nil {
		return err
	}
	fmt.Println("  Created .claude/commands/status.md")

	fmt.Println("\n✓ Mere initialized successfully!")
	fmt.Println("\nNext steps:")
	fmt.Println("  1. Update .claude/CLAUDE.md with your project context")
	fmt.Println("  2. Use /module to create your first deep module")
	fmt.Println("  3. Use /status to check project state")

	return nil
}

// ValidateProject validates that the current directory is a valid project
func ValidateProject() error {
	// Placeholder implementation
	return nil
}

// CreateDirectory creates a directory with proper permissions
func CreateDirectory(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}

// WriteFile writes content to a file with proper permissions
func WriteFile(path string, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}
	return nil
}
