package methods

type CreateFileRequest struct {
	Request

	Params CreateFileParams `json:"params"`
}

type CreateFileParams struct {
	/**
	 * The name of the file to create
	 */
	Name string `json:"name"`
	/**
	 * The directory where the file should be created
	 */
	Dir string `json:"dir"`
	/**
	 * The content of the file
	 */
	Text string `json:"text"`
}

type CreateFileResponse struct {
	Response

	Result TextFile `json:"result"`
}
