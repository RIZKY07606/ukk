package getall

type GetAllSiswaResponse struct {
	SiswaID string `json:"siswa_id"`
	NIS     string `json:"nis"`
	Nama    string `json:"nama"`
	Kelas   string `json:"kelas"`
	Jurusan string `json:"jurusan"`
}

type GetAllSiswaResponseWrapper struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    []GetAllSiswaResponse `json:"data"`
}
