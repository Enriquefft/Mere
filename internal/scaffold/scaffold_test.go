package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDirectory(t *testing.T) {
	// Create a temp directory for testing
	tmpDir := t.TempDir()
	testDir := filepath.Join(tmpDir, "test", "nested", "dir")

	err := CreateDirectory(testDir)
	if err != nil {
		t.Fatalf("CreateDirectory() error = %v", err)
	}

	// Verify directory was created
	info, err := os.Stat(testDir)
	if err != nil {
		t.Fatalf("Directory not created: %v", err)
	}

	if !info.IsDir() {
		t.Error("CreateDirectory() did not create a directory")
	}

	// Check permissions (should be 0755)
	expectedPerms := os.FileMode(0755)
	if info.Mode().Perm() != expectedPerms {
		t.Errorf("Directory permissions = %v, expected %v", info.Mode().Perm(), expectedPerms)
	}
}

func TestWriteFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	testContent := "test content\n"

	err := WriteFile(testFile, testContent)
	if err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	// Verify file was created
	info, err := os.Stat(testFile)
	if err != nil {
		t.Fatalf("File not created: %v", err)
	}

	if info.IsDir() {
		t.Error("WriteFile() created a directory instead of a file")
	}

	// Verify content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(content) != testContent {
		t.Errorf("File content = %q, expected %q", string(content), testContent)
	}

	// Check permissions (should be 0644)
	expectedPerms := os.FileMode(0644)
	if info.Mode().Perm() != expectedPerms {
		t.Errorf("File permissions = %v, expected %v", info.Mode().Perm(), expectedPerms)
	}
}

func TestInitProject_New(t *testing.T) {
	// Create a temp directory for testing
	tmpDir := t.TempDir()

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Initialize project
	err = InitProject(false)
	if err != nil {
		t.Fatalf("InitProject() error = %v", err)
	}

	// Verify .claude directory was created
	claudeDir := ".claude"
	info, err := os.Stat(claudeDir)
	if err != nil {
		t.Fatalf(".claude directory not created: %v", err)
	}

	if !info.IsDir() {
		t.Error("InitProject() did not create .claude directory")
	}

	// Verify CLAUDE.md was created at project root
	if _, err := os.Stat("CLAUDE.md"); err != nil {
		t.Fatalf("CLAUDE.md not created at project root: %v", err)
	}

	// Verify commands directory was created
	commandsDir := filepath.Join(claudeDir, "commands")
	if _, err := os.Stat(commandsDir); err != nil {
		t.Fatalf("commands directory not created: %v", err)
	}

	// Verify command templates were created
	expectedFiles := []string{
		filepath.Join(commandsDir, "module.md"),
		filepath.Join(commandsDir, "status.md"),
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); err != nil {
			t.Errorf("Expected file not created: %s", file)
		}
	}
}

func TestInitProject_ExistingWithoutForce(t *testing.T) {
	tmpDir := t.TempDir()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Initialize project first time
	err = InitProject(false)
	if err != nil {
		t.Fatalf("First InitProject() error = %v", err)
	}

	// Try to initialize again without force
	err = InitProject(false)
	if err == nil {
		t.Error("InitProject() should return error when .claude already exists without force flag")
	}

	expectedErrMsg := ".claude/ directory already exists"
	if err == nil || !containsString(err.Error(), expectedErrMsg) {
		t.Errorf("Expected error containing %q, got: %v", expectedErrMsg, err)
	}
}

func TestInitProject_ExistingWithForce(t *testing.T) {
	tmpDir := t.TempDir()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Initialize project first time
	err = InitProject(false)
	if err != nil {
		t.Fatalf("First InitProject() error = %v", err)
	}

	// Add custom content to CLAUDE.md at root to verify it's overwritten
	err = os.WriteFile("CLAUDE.md", []byte("custom content"), 0644)
	if err != nil {
		t.Fatalf("Failed to write custom file: %v", err)
	}

	// Initialize again with force
	err = InitProject(true)
	if err != nil {
		t.Fatalf("InitProject(force=true) error = %v", err)
	}

	// Verify custom content was replaced
	content, err := os.ReadFile("CLAUDE.md")
	if err != nil {
		t.Fatalf("Failed to read CLAUDE.md: %v", err)
	}

	if string(content) == "custom content" {
		t.Error("InitProject(force=true) did not overwrite existing files")
	}

	// Verify content is from template (not the custom content)
	expectedContent := "## What This Is"
	if !containsString(string(content), expectedContent) {
		t.Errorf("CLAUDE.md does not contain expected template content: %s", expectedContent)
	}
}

func TestCreateDirectory_ErrorHandling(t *testing.T) {
	// Test with an invalid path that cannot be created
	invalidPath := "/root/nonexistent/test"

	err := CreateDirectory(invalidPath)
	if err == nil {
		t.Error("CreateDirectory() should return error for invalid path")
	}
}

func TestWriteFile_ErrorHandling(t *testing.T) {
	// Test with an invalid path
	invalidPath := "/root/nonexistent/test.txt"

	err := WriteFile(invalidPath, "content")
	if err == nil {
		t.Error("WriteFile() should return error for invalid path")
	}
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || containsString(s[1:], substr)))
}
