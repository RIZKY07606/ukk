package create

type Request struct {
	Nama    string `json:"nama"`
	Tipe    string `json:"tipe,omitempty"`
	URL     string `json:"url,omitempty"`
	KaryaID string `json:"karya_id"`
}
