package fsmod

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	jsoniter "github.com/json-iterator/go"
	Conf "github.com/oguz-yilmaz/file-system-server/pkg/config"
)

func TestOverridesTheContent(t *testing.T) {
	tempDir := t.TempDir()

	file1, err := createFile(map[string]any{
		"name":      "example.txt",
		"content":   "this is an example file.",
		"overwrite": true,
	}, tempDir, t)
	if err != nil {
		t.Error(err)
	}

	ensureFileIsCreated(*file1, t)

	file2, err := createFile(map[string]any{
		"name":      "example.txt",
		"content":   "Updated",
		"overwrite": true,
	}, tempDir, t)
	if err != nil {
		t.Error(err)
	}

	_, content := ensureFileIsCreated(*file2, t)
	if content != "Updated" {
		t.Error("Content is not updated")
	}

	_, err = createFile(map[string]any{
		"name":      "example.txt",
		"content":   "Updated",
		"overwrite": false,
	}, tempDir, t)

	if err == nil {
		t.Error("Error should be thrown when overwrite is false, when the file already exists")
	}
}

func TestCreatesFileInCorrectDirectory(t *testing.T) {
	tempDir := t.TempDir()

	file, err := createFile(map[string]any{
		"name":      "example.txt",
		"content":   "this is an example file.",
		"dir":       "./test",
		"root":      tempDir,
		"file-type": "txt",
		"overwrite": true,
	}, tempDir, t)

	if err != nil {
		t.Error(err)
	}

	ensureFileIsCreated(*file, t)
}

func TestCreatesFile(t *testing.T) {
	tempDir := t.TempDir()

	file, err := createFile(map[string]any{
		"name":      "example.txt",
		"content":   "this is an example file.",
		"file-type": "txt",
		"overwrite": true,
	}, tempDir, t)

	if err != nil {
		t.Error(err)
	}

	ensureFileIsCreated(*file, t)
}

func ensureFileIsCreated(file File, t *testing.T) (fs.FileInfo, string) {
	fullPath := filepath.Join(file.Dir, file.Name)
	info, err := os.Stat(fullPath) // ensure the file is created
	if err != nil {
		t.Error(err)
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Error(err)

		return nil, ""
	}

	if string(file.Content) != string(content) {
		t.Error("Content is not correct")
	}

	if info.Name() != file.Name {
		t.Error("File name is not correct")
	}

	if info.IsDir() != false {
		t.Error("File is not a file, it is a directory")
	}

	return info, string(content)
}

func createFile(params map[string]any, dir string, t *testing.T) (*File, error) {
	params["dir"] = dir
	req := createFileParamsJsonString(params)

	conf := Conf.NewDefaultConfig()
	cfp := NewCreateFileParams(map[string]any{
		"dir": dir,
	}, conf)

	if err := jsoniter.Unmarshal([]byte(req), &cfp); err != nil {
		t.Error("Error decoding CreateFileRequest")
	}

	file, err := CreateFile(cfp)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func createFileParamsJsonString(params map[string]any) string {
	jsonMap := make(map[string]any)

	if val, exists := params["name"]; exists {
		jsonMap["name"] = val.(string)
	}
	if val, exists := params["content"]; exists {
		if content, ok := val.([]byte); ok {
			jsonMap["content"] = content
		} else if contentStr, ok := val.(string); ok {
			jsonMap["content"] = []byte(contentStr) // Convert string to []byte if necessary
		}
	}
	if val, exists := params["file-type"]; exists {
		jsonMap["file-type"] = val.(string)
	}
	if val, exists := params["permissions"]; exists {
		jsonMap["permissions"] = val.(int)
	}
	if val, exists := params["overwrite"]; exists {
		jsonMap["overwrite"] = val.(bool)
	}

	jsonData, err := jsoniter.MarshalIndent(jsonMap, "", "    ")
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return ""
	}

	return string(jsonData)
}
