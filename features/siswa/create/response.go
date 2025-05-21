package create

type CreateSiswaResponse struct {
	SiswaID  string `json:"siswa_id"`
	NIS      string `json:"nis"`
	Nama     string `json:"nama"`
	Kelas    string `json:"kelas"`
	Jurusan  string `json:"jurusan"`
	Angkatan string `json:"angkatan"`
}

type CreateSiswaResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    CreateSiswaResponse `json:"data"`
}
