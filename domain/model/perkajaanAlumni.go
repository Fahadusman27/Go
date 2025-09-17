package model

import "time"

type PerkajaanAlumni struct {
	ID				int		`json:"id"`
	IDAlumni		 string	`json:"id_alumni"`
	StatusKerja		 string	`json:"status_kerja"`
	JenisIndustri	string	`json:"jenis_industri"`
	Jabatan			string	`json:"jabatan"`
	Gaji			int		`json:"gaji"`
	LamaBekerja		int		`json:"lama_bekerja"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}