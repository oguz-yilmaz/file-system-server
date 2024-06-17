package fsmod

import (
	"path/filepath"
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
	 */
	Content string `json:"content"`
	/**
	 * The size of the file in bytes
	 */
	Size int `json:"size"`
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
	MetaData map[string]any `json:"metaData"`
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

func NewFileFromCreateFileParams(fileParams *CreateFileParams) *File {
	file := NewFile()

	ext := filepath.Ext(fileParams.Name)

	file.Name = fileParams.Name
	file.Dir = fileParams.Dir
	file.Content = fileParams.Content
	file.FileType = fileParams.FileType
	file.Permissions = fileParams.Permissions
	file.Extension = ext[1:] // Remove the leading dot from the extension

	return file
}

func NewFile() *File {
	return &File{
		Version:     1,
		MIMEType:    MIMETextPlain,
		Permissions: 438, // 0666
	}
}
