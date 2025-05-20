package update

type UpdateKategoriKaryaResponse struct {
	KategoriID string `json:"kategori_id"`
	Nama       string `json:"nama"`
}

type UpdateKategoriKaryaResponseWrapper struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    UpdateKategoriKaryaResponse `json:"data"`
}
