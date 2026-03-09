package embed

import (
	"embed"
	"io/fs"
)

//go:embed templates
var templatesFS embed.FS

// GetTemplatesFS returns the embedded filesystem
func GetTemplatesFS() fs.FS {
	return templatesFS
}

// GetTemplate returns the content of a specific template file
func GetTemplate(path string) (string, error) {
	content, err := templatesFS.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// GetClaudeMD returns the CLAUDE.md template content
func GetClaudeMD() (string, error) {
	return GetTemplate("templates/claude/CLAUDE.md")
}

// GetModuleMD returns the module.md template content
func GetModuleMD() (string, error) {
	return GetTemplate("templates/commands/module.md")
}

// GetStatusMD returns the status.md template content
func GetStatusMD() (string, error) {
	return GetTemplate("templates/commands/status.md")
}

// GetInterfaceMD returns the INTERFACE.md template content
func GetInterfaceMD() (string, error) {
	return GetTemplate("templates/INTERFACE.md.template")
}

// GetArchitectureMD returns the ARCHITECTURE.md template content
func GetArchitectureMD() (string, error) {
	return GetTemplate("templates/claude/ARCHITECTURE.md")
}

// GetManifestMD returns the MANIFEST.md template content
func GetManifestMD() (string, error) {
	return GetTemplate("templates/claude/MANIFEST.md")
}

// ListTemplates returns a list of all embedded template files
func ListTemplates() ([]string, error) {
	var files []string
	err := fs.WalkDir(templatesFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
