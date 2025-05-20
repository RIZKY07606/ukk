package getbyid

type GetFileUploadByIDResponseWrapper struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    FileUploadDetailResponse `json:"data"`
}

type FileUploadDetailResponse struct {
	FileID    string `json:"file_id"`
	Nama      string `json:"nama"`
	Tipe      string `json:"tipe"`
	URL       string `json:"url"`
	KaryaID   string `json:"karya_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
