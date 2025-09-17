package repository

import (
	"tugas/domain/config"
	"tugas/domain/model"
)

// CheckPerkajaanAlumniByIDAlumni checks if a PerkajaanAlumni exists by IDAlumni
func CheckPerkajaanAlumniByIDAlumni(idAlumni string) (*model.PerkajaanAlumni, error) {
	perkajaan := new(model.PerkajaanAlumni)
	query := `
		SELECT id, id_alumni, status_kerja, jenis_industri,
		    jabatan, gaji, lama_bekerja
		FROM perkajaan_alumni WHERE id_alumni = $1 LIMIT 1`
	err := config.DB.QueryRow(query, idAlumni).Scan(
		&perkajaan.ID,
		&perkajaan.IDAlumni,
		&perkajaan.StatusKerja,
		&perkajaan.JenisIndustri,
		&perkajaan.Jabatan,
		&perkajaan.Gaji,
		&perkajaan.LamaBekerja,
	)
	if err != nil {
		return nil, err
	}
	return perkajaan, nil
}

func CreatePerkajaanAlumni(perkajaan *model.PerkajaanAlumni) error {
	query := `
		INSERT INTO perkajaan_alumni (
			id_alumni, status_kerja, jenis_industri,
			jabatan, gaji, lama_bekerja
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	err := config.DB.QueryRow(query,
		perkajaan.IDAlumni,
		perkajaan.StatusKerja,
		perkajaan.JenisIndustri,
		perkajaan.Jabatan,
		perkajaan.Gaji,
		perkajaan.LamaBekerja,
	).Scan(&perkajaan.ID)
	return err
}

func UpdatePerkajaanAlumni(idAlumni string, perkajaan *model.PerkajaanAlumni) error {
	query := `
		UPDATE perkajaan_alumni
		SET status_kerja = $1, jenis_industri = $2, jabatan = $3,
		    gaji = $4, lama_bekerja = $5
		WHERE id_alumni = $6`
	_, err := config.DB.Exec(query,
		perkajaan.StatusKerja,
		perkajaan.JenisIndustri,
		perkajaan.Jabatan,
		perkajaan.Gaji,
		perkajaan.LamaBekerja,
		idAlumni,
	)
	return err
}

func DeletePerkajaanAlumni(idAlumni string) error {
	query := `DELETE FROM perkajaan_alumni WHERE id_alumni = $1`
	_, err := config.DB.Exec(query, idAlumni)
	return err
}

func GetAllPerkajaanAlumni() ([]model.PerkajaanAlumni, error) {
	query := `t
		SELECT id, id_alumni, status_kerja, jenis_industri, jabatan, gaji, lama_bekerja
		FROM perkajaan_alumni`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var perkajaanList []model.PerkajaanAlumni
	for rows.Next() {
		var perkajaan model.PerkajaanAlumni
		err := rows.Scan(
			&perkajaan.ID,
			&perkajaan.IDAlumni,
			&perkajaan.StatusKerja,
			&perkajaan.JenisIndustri,
			&perkajaan.Jabatan,
			&perkajaan.Gaji,
			&perkajaan.LamaBekerja,
		)
		if err != nil {
			return nil, err
		}
		perkajaanList = append(perkajaanList, perkajaan)
	}
	return perkajaanList, nil
}