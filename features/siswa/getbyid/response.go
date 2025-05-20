package getbyid

type GetSiswaResponse struct {
	SiswaID string `json:"siswa_id"`
	NIS     string `json:"nis"`
	Nama    string `json:"nama"`
	Kelas   string `json:"kelas"`
	Jurusan string `json:"jurusan"`
}

type GetSiswaResponseWrapper struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    GetSiswaResponse `json:"data"`
}
