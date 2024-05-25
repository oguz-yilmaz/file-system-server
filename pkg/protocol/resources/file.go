package resources

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
	Content []byte `json:"content"`
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
	Version *int `json:"version"`
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
	MIMEType *MIMEType `json:"mimeType, omitempty"`
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
