package embed

import (
	"io/fs"
	"strings"
	"testing"
)

func TestGetClaudeMD(t *testing.T) {
	content, err := GetClaudeMD()
	if err != nil {
		t.Fatalf("GetClaudeMD() error = %v", err)
	}

	if len(content) == 0 {
		t.Error("GetClaudeMD() returned empty content")
	}

	// Verify it contains expected content
	expected := "# Project Context"
	if !strings.Contains(content, expected) {
		t.Errorf("GetClaudeMD() does not contain expected content: %s", expected)
	}

	// Check approximate length (should be ~50-80 lines as per PRD with new Language-Agnostic Philosophy section)
	lines := strings.Split(content, "\n")
	if len(lines) < 50 || len(lines) > 90 {
		t.Errorf("GetClaudeMD() line count = %d, expected ~50-80 lines", len(lines))
	}
}

func TestGetModuleMD(t *testing.T) {
	content, err := GetModuleMD()
	if err != nil {
		t.Fatalf("GetModuleMD() error = %v", err)
	}

	if len(content) == 0 {
		t.Error("GetModuleMD() returned empty content")
	}

	expected := "/module"
	if !strings.Contains(content, expected) {
		t.Errorf("GetModuleMD() does not contain expected content: %s", expected)
	}
}

func TestGetStatusMD(t *testing.T) {
	content, err := GetStatusMD()
	if err != nil {
		t.Fatalf("GetStatusMD() error = %v", err)
	}

	if len(content) == 0 {
		t.Error("GetStatusMD() returned empty content")
	}

	expected := "/status"
	if !strings.Contains(content, expected) {
		t.Errorf("GetStatusMD() does not contain expected content: %s", expected)
	}
}

func TestGetInterfaceMD(t *testing.T) {
	content, err := GetInterfaceMD()
	if err != nil {
		t.Fatalf("GetInterfaceMD() error = %v", err)
	}

	if len(content) == 0 {
		t.Error("GetInterfaceMD() returned empty content")
	}

	expected := "Module:"
	if !strings.Contains(content, expected) {
		t.Errorf("GetInterfaceMD() does not contain expected content: %s", expected)
	}
}

func TestListTemplates(t *testing.T) {
	files, err := ListTemplates()
	if err != nil {
		t.Fatalf("ListTemplates() error = %v", err)
	}

	if len(files) == 0 {
		t.Error("ListTemplates() returned no files")
	}

	// Check for expected files
	expectedFiles := map[string]bool{
		"templates/claude/CLAUDE.md":      false,
		"templates/commands/module.md":      false,
		"templates/commands/status.md":      false,
		"templates/INTERFACE.md.template": false,
	}

	for _, file := range files {
		if _, exists := expectedFiles[file]; exists {
			expectedFiles[file] = true
		}
	}

	for file, found := range expectedFiles {
		if !found {
			t.Errorf("ListTemplates() missing expected file: %s", file)
		}
	}
}

func TestGetTemplatesFS(t *testing.T) {
	fileSystem := GetTemplatesFS()
	if fileSystem == nil {
		t.Error("GetTemplatesFS() returned nil")
	}

	// Try to read a file through the filesystem
	_, err := fs.ReadFile(fileSystem, "templates/claude/CLAUDE.md")
	if err != nil {
		t.Errorf("GetTemplatesFS() filesystem error: %v", err)
	}
}

func TestAllTemplatesEmbedded(t *testing.T) {
	// Verify that all template functions return valid content
	tests := []struct {
		name    string
		getFunc func() (string, error)
	}{
		{"CLAUDE.md", GetClaudeMD},
		{"module.md", GetModuleMD},
		{"status.md", GetStatusMD},
		{"INTERFACE.md.template", GetInterfaceMD},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := tt.getFunc()
			if err != nil {
				t.Errorf("%s: error = %v", tt.name, err)
				return
			}

			if len(content) == 0 {
				t.Errorf("%s: returned empty content", tt.name)
			}

			// Verify it's valid markdown (has at least one # header)
			if !strings.Contains(content, "#") {
				t.Errorf("%s: content doesn't appear to be valid markdown (no # headers)", tt.name)
			}
		})
	}
}

func TestEmbeddedFilesystemStructure(t *testing.T) {
	fileSystem := GetTemplatesFS()

	// Walk the filesystem and verify structure
	expectedDirs := []string{
		"templates",
		"templates/claude",
		"templates/commands",
	}

	foundDirs := make(map[string]bool)
	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			foundDirs[path] = true
		}
		return nil
	})

	if err != nil {
		t.Fatalf("Failed to walk embedded filesystem: %v", err)
	}

	for _, dir := range expectedDirs {
		if !foundDirs[dir] {
			t.Errorf("Expected directory not found: %s", dir)
		}
	}
}

func TestTemplatesAreLanguageAgnostic(t *testing.T) {
	templates := map[string]func() (string, error){
		"CLAUDE.md": GetClaudeMD,
		"module.md": GetModuleMD,
		"status.md": GetStatusMD,
	}

	// Check for file extensions - these should never appear in templates
	fileExtensionPatterns := []string{
		"\\.ts\"",      // TypeScript extensions
		"\\.go\"",      // Go extensions
		"\\.py\"",      // Python extensions
		"\\.java\"",    // Java extensions
		"\\.rs\"",      // Rust extensions
	}

	// Check for actual test commands at start of lines (commands, not examples)
	testCommandPatterns := []string{
		"npm test",     // npm test at line start
		"go test",      // go test at line start
		"pytest",       // pytest at line start
		"mvn test",     // maven test at line start
		"cargo test",    // cargo test at line start
	}

	for name, getFunc := range templates {
		t.Run(name, func(t *testing.T) {
			content, err := getFunc()
			if err != nil {
				t.Fatalf("%s: error = %v", name, err)
			}

			// Check for file extensions
			for _, pattern := range fileExtensionPatterns {
				if strings.Contains(content, pattern) {
					t.Errorf("%s contains language-specific file extension pattern %q", name, pattern)
				}
			}

			// Check for test commands at line start (not in examples)
			lines := strings.Split(content, "\n")
			for _, line := range lines {
				trimmed := strings.TrimSpace(line)
				for _, pattern := range testCommandPatterns {
					// Only flag if pattern appears at line start (actual command, not example)
					if strings.HasPrefix(trimmed, pattern) && !strings.Contains(line, "Example patterns") {
						t.Errorf("%s contains language-specific test command %q", name, pattern)
					}
				}
			}
		})
	}
}

func TestTemplatesContainAgnosticGuidance(t *testing.T) {
	templates := map[string]func() (string, error){
		"module.md": GetModuleMD,
		"CLAUDE.md": GetClaudeMD,
	}
	// Note: status.md is excluded as it's a reporting command, not a scaffolding command

	agnosticPatterns := []string{
		"language-agnostic",
		"AI Inference Guidelines",
		"detect the project's language",
		"module entry point file",
		"boundary test file",
	}

	for name, getFunc := range templates {
		t.Run(name, func(t *testing.T) {
			content, err := getFunc()
			if err != nil {
				t.Fatalf("%s: error = %v", name, err)
			}

			// module.md should have AI inference guidelines
			if name == "module.md" {
				if !strings.Contains(content, "AI Inference Guidelines") {
					t.Errorf("%s should contain 'AI Inference Guidelines' section", name)
				}
			}

			// All should contain at least one agnostic pattern
			found := false
			for _, pattern := range agnosticPatterns {
				if strings.Contains(content, pattern) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("%s should contain at least one language-agnostic guidance pattern", name)
			}
		})
	}
}
