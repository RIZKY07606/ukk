package create

type Request struct {
	NIS     string `json:"nis"`
	Nama    string `json:"nama"`
	Kelas   string `json:"kelas"`
	Jurusan string `json:"jurusan"`
	UserID  string `json:"user_id"`
}
