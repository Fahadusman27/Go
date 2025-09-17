package model

import "time"

type Alumni	struct	{
				NIM			string			`json:"nim"`
				Nama		string			`json:"nama"`
				Angkatan	*int				`json:"angkatan"`
				TahunLulus	*int				`json:"tahun_lulus"`
				IDFakultas	*int				`json:"id_fakultas"`
				IDProdi		*int				`json:"id_prodi"`
				Sumber		*string			`json:"sumber"`
				CreatedAt	time.Time		`json:"created_at"`
				UpdatedAt	time.Time		`json:"updated_at"`
}

type Jumlah struct {
				Angkatan       int `json:"angkatan"`
				JumlahAngkatan int `json:"jumlah_angkatan"`
}