package getbyid

type GetKaryaUKKByIDResponse struct {
	KaryaID    string `json:"karya_id"`
	Judul      string `json:"judul"`
	Deskripsi  string `json:"deskripsi,omitempty"`
	Thumbnail  string `json:"thumbnail,omitempty"`
	Link       string `json:"link,omitempty"`
	SiswaID    string `json:"siswa_id"`
	KategoriID string `json:"kategori_id"`
	AdminID    string `json:"admin_id"`
}

type GetKaryaUKKByIDResponseWrapper struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    GetKaryaUKKByIDResponse `json:"data"`
}
