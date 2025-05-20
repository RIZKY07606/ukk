package getall

type FileUploadResponse struct {
	FileID    string `json:"file_id"`
	Nama      string `json:"nama"`
	Tipe      string `json:"tipe"`
	URL       string `json:"url"`
	KaryaID   string `json:"karya_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetAllFileUploadResponseWrapper struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    []FileUploadResponse `json:"data"`
}
