package resources

type TextFile struct {
	/**
	 * Name of the file
	 */
	Name string `json:"name"`

	/**
	 * Absolute path of the parent directory of the file
	 */
	Dir string `json:"dir"`

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
	 * The content of the opened text document.
	 */
	Text string `json:"text"`

	/**
	 * A map for storing any additional data
	 */
	MetaData map[string]interface{} `json:"metaData"`

	/**
	 * The MIME type of the file
	 */
	MIMEType *MIMEType `json:"mimeType, omitempty"`

	/**
	 * The size of the file in bytes
	 */
	Size *int `json:"size"`
}

type MIMEType string

const (
	MIMETextPlain      MIMEType = "text/plain"
	MIMETextHTML       MIMEType = "text/html"
	MIMEImageJPEG      MIMEType = "image/jpeg"
	MIMEImagePNG       MIMEType = "image/png"
	MIMEApplicationPDF MIMEType = "application/pdf"

	// Add more MIME types as needed
)

func (mt MIMEType) String() string {
	return string(mt)
}
