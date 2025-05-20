package update

type UpdateFileUploadResponse struct {
	FileID    string `json:"file_id"`
	Nama      string `json:"nama"`
	Tipe      string `json:"tipe"`
	URL       string `json:"url"`
	KaryaID   string `json:"karya_id"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateFileUploadResponseWrapper struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    UpdateFileUploadResponse `json:"data"`
}
