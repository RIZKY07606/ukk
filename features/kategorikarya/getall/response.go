package getall

type GetAllKategoriKaryaResponse struct {
	KategoriID string `json:"kategori_id"`
	Nama       string `json:"nama"`
}

type GetAllKategoriKaryaResponseWrapper struct {
	Code    int                           `json:"code"`
	Message string                        `json:"message"`
	Data    []GetAllKategoriKaryaResponse `json:"data"`
}
