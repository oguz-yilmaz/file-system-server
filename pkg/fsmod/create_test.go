package fsmod

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	jsoniter "github.com/json-iterator/go"
	Conf "github.com/oguz-yilmaz/file-system-server/pkg/config"
)

func createFileParamsJsonString(params map[string]string) string {
	req := fmt.Sprintf(`
            {
                "name": "%s",
                "content": "%s",
                "file-type": "%s",
                "permissions": %s,
                "overwrite": %s
            }
        `, params["name"], params["content"], params["file-type"], params["permissions"], params["overwrite"])

	return req
}

func TestCreateFile(t *testing.T) {
	// creates a new temp dir, clear everything inside when the test has completed
	tempDir := t.TempDir()

	req := createFileParamsJsonString(map[string]string{
		"name":        "example.txt",
		"content":     "this is an example file.",
		"file-type":   "txt",
		"permissions": "438",
		"overwrite":   "true",
		"dir":         tempDir,
	})

	conf := Conf.NewDefaultConfig()
	cfp := NewCreateFileParams(map[string]any{
		"dir": tempDir,
	}, conf)

	if err := jsoniter.Unmarshal([]byte(req), &cfp); err != nil {
		t.Error(err)
	}

	file, err := CreateFile(cfp)
	if err != nil {
		t.Error(err)
	}

	fullPath := filepath.Join(tempDir, file.Name)

	info, err := os.Stat(fullPath)
	if err != nil {
		t.Error(err)
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Error(err)

		return
	}

	if file.Content != string(content) {
		t.Error("Content is not correct")
	}

	if info.Name() != file.Name {
		t.Error("File name is not correct")
	}

	if info.IsDir() != false {
		t.Error("File is not a file, it is a directory")
	}
}
