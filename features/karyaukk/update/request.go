package update

type Request struct {
	Judul      string `json:"judul"`
	Deskripsi  string `json:"deskripsi,omitempty"`
	Link       string `json:"link,omitempty"`
	SiswaID    string `json:"siswa_id"`
	KategoriID string `json:"kategori_id"`
}
