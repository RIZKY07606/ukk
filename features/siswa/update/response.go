package update

type UpdateSiswaResponse struct {
	SiswaID string `json:"siswa_id"`
	NIS     string `json:"nis"`
	Nama    string `json:"nama"`
	Kelas   string `json:"kelas"`
	Jurusan string `json:"jurusan"`
}

type UpdateSiswaResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    UpdateSiswaResponse `json:"data"`
}
