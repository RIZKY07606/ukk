package create

type CreateKaryaUKKRequest struct {
	Judul      string `json:"judul"`
	Deskripsi  string `json:"deskripsi,omitempty"`
	Thumbnail  string `json:"thumbnail,omitempty"`
	Link       string `json:"link,omitempty"`
	SiswaID    string `json:"siswa_id"`
	KategoriID string `json:"kategori_id"`
	AdminID    string `json:"admin_id"`
}
