package update

type UpdateKaryaUKKRequest struct {
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Link      string `json:"link,omitempty"`
}
