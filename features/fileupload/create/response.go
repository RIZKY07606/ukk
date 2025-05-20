package create

type CreateFileUploadResponse struct {
	FileID    string `json:"file_id"`
	Nama      string `json:"nama"`
	Tipe      string `json:"tipe"`
	URL       string `json:"url"`
	KaryaID   string `json:"karya_id"`
	CreatedAt string `json:"created_at"`
}

type CreateFileUploadResponseWrapper struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    CreateFileUploadResponse `json:"data"`
}
