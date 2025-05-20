package create

type CreateKategoriKaryaResponse struct {
	KategoriID string `json:"kategori_id"`
	Nama       string `json:"nama"`
}

type CreateKategoriKaryaResponseWrapper struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    CreateKategoriKaryaResponse `json:"data"`
}
