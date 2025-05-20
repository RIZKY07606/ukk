package getbyid

type GetKategoriKaryaByIDResponse struct {
	KategoriID string `json:"kategori_id"`
	Nama       string `json:"nama"`
}

type GetKategoriKaryaByIDResponseWrapper struct {
	Code    int                          `json:"code"`
	Message string                       `json:"message"`
	Data    GetKategoriKaryaByIDResponse `json:"data"`
}
