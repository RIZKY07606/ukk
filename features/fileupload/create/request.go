package create

type CreateFileUploadRequest struct {
	Nama    string `json:"nama"`
	Tipe    string `json:"tipe"`
	URL     string `json:"url"`
	KaryaID string `json:"karya_id"`
}
