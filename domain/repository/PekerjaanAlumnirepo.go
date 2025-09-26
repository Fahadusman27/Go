package repository

import (
	"tugas/domain/config"
	"tugas/domain/model"
)

func CheckpekerjaanAlumniByID(id string) (*model.PekerjaanAlumni, error) {
	pekerjaan := new(model.PekerjaanAlumni)
	query := `
		SELECT id, nim_alumni, status_kerja, jenis_industri, pekerjaan,
		    jabatan, gaji, lama_bekerja
		FROM pekerjaan_alumni WHERE id = $1 LIMIT 1`
	err := config.DB.QueryRow(query, id).Scan(
		&pekerjaan.ID,
		&pekerjaan.NimAlumni,
		&pekerjaan.StatusKerja,
		&pekerjaan.JenisIndustri,
		&pekerjaan.Pekerjaan,
		&pekerjaan.Jabatan,
		&pekerjaan.Gaji,
		&pekerjaan.LamaBekerja,
	)
	if err != nil {
		return nil, err
	}
	return pekerjaan, nil
}

func CreatepekerjaanAlumni(pekerjaan *model.PekerjaanAlumni) error {
	query := `
		INSERT INTO pekerjaan_alumni (
			nim_alumni, status_kerja, jenis_industri, pekerjaan,
			jabatan, gaji, lama_bekerja
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	err := config.DB.QueryRow(query,
		pekerjaan.NimAlumni,
		pekerjaan.StatusKerja,
		pekerjaan.JenisIndustri,
		pekerjaan.Pekerjaan,
		pekerjaan.Jabatan,
		pekerjaan.Gaji,
		pekerjaan.LamaBekerja,
	).Scan(&pekerjaan.ID)
	return err
}

func UpdatepekerjaanAlumni(idAlumni string, pekerjaan *model.PekerjaanAlumni) error {
	query := `
		UPDATE pekerjaan_alumni
		SET status_kerja = $1, jenis_industri = $2, pekerjaan=$3, jabatan = $4,
		    gaji = $5, lama_bekerja = $6, pekerjaan = $7
		WHERE nim_alumni = $8`
	_, err := config.DB.Exec(query,
		pekerjaan.StatusKerja,
		pekerjaan.JenisIndustri,
		pekerjaan.Pekerjaan,
		pekerjaan.Jabatan,
		pekerjaan.Gaji,
		pekerjaan.LamaBekerja,
		pekerjaan.Pekerjaan,
		idAlumni,
	)
	return err
}

func DeletepekerjaanAlumni(id string) error {
	query := `DELETE FROM pekerjaan_alumni WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	return err
}

func GetAllpekerjaanAlumni() ([]model.PekerjaanAlumni, error) {
	query := `SELECT id, nim_alumni, status_kerja, jenis_industri, pekerjaan, jabatan, gaji, lama_bekerja
		FROM pekerjaan_alumni`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pekerjaanList []model.PekerjaanAlumni
	for rows.Next() {
		var pekerjaan model.PekerjaanAlumni
		err := rows.Scan(
			&pekerjaan.ID,
			&pekerjaan.NimAlumni,
			&pekerjaan.StatusKerja,
			&pekerjaan.JenisIndustri,
			&pekerjaan.Pekerjaan,
			&pekerjaan.Jabatan,
			&pekerjaan.Gaji,
			&pekerjaan.LamaBekerja,
		)
		if err != nil {
			return nil, err
		}
		pekerjaanList = append(pekerjaanList, pekerjaan)
	}
	return pekerjaanList, nil
}
