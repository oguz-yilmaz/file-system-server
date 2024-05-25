package resources

import (
	"path/filepath"

	"github.com/oguz-yilmaz/file-system-server/pkg/fsmod/request"
)

const (
	TEXT_FILE   = "text"
	BINARY_FILE = "binary"
	IMAGE_FILE  = "image"
	PDF_FILE    = "pdf"
	HTML_FILE   = "html"
)

type File struct {
	/**
	 * Name of the file
	 */
	Name string `json:"name"`
	/**
	 * Absolute path of the parent directory of the file
	 */
	Dir string `json:"dir"`
	/**
	* The content of the file
	* TODO: Currently, we only support text files
	 */
	Content string `json:"content"`
	/**
	 * The size of the file in bytes
	 */
	Size *int `json:"size"`
	/**
	 * The type of the file, e.g. text, image, etc.
	 */
	FileType string `json:"fileType"`
	/**
	 * The version number of this document (it will increase after each
	 * change, including undo/redo).
	 */
	Version int `json:"version"`
	/**
	 * The extension of the TextFile
	 */
	Extension string `json:"extension"`
	/**
	 * A map for storing any additional data
	 */
	MetaData map[string]interface{} `json:"metaData"`
	/**
	 * The MIME type of the file
	 */
	MIMEType MIMEType `json:"mimeType"`
	/**
	 * The permissions of the file, e.g. 0777
	 */
	Permissions int `json:"permissions"`
}

type TextFile struct {
	File
	Content string `json:"text"`
}

type MIMEType string

const (
	MIMETextPlain      MIMEType = "text/plain"
	MIMETextHTML       MIMEType = "text/html"
	MIMEImageJPEG      MIMEType = "image/jpeg"
	MIMEImagePNG       MIMEType = "image/png"
	MIMEApplicationPDF MIMEType = "application/pdf"
)

func (mt MIMEType) String() string {
	return string(mt)
}

func NewFile(fileParams *request.CreateFileParams) *File {
	defaultFile := &File{
		Version:     1,
		MIMEType:    MIMETextPlain,
		Permissions: 438, // 0666
	}

	if fileParams != nil {
		ext := filepath.Ext(fileParams.Name)

		defaultFile.Name = fileParams.Name
		defaultFile.Dir = fileParams.Dir
		defaultFile.Content = fileParams.Content
		defaultFile.FileType = fileParams.FileType
		defaultFile.Permissions = fileParams.Permissions
		defaultFile.Extension = ext[1:] // Remove the leading dot from the extension
	}

	return defaultFile
}
