package repository

import (
	"tugas/domain/config"
	"tugas/domain/model"
)

func CheckAlumniByNim(nim string) (*model.Alumni, error) {
	alumni := new(model.Alumni)
	query := `SELECT nim, full_name, angkatan, id_fakultas, id_prodi, tahun_lulus, sumber
        FROM alumni WHERE nim = $1 LIMIT 1`
	err := config.DB.QueryRow(query, nim).Scan(&alumni.NIM, &alumni.Nama, &alumni.Angkatan, &alumni.IDFakultas, &alumni.IDProdi,
		&alumni.TahunLulus, &alumni.Sumber)
	if err != nil {
		return nil, err
	}
	return alumni, nil
}

func GetJumlahAlumni(angkatan int) (*model.Jumlah, error) {
    jumlah := new(model.Jumlah)
    query := `SELECT angkatan, COUNT(*) AS jumlah_angkatan 
              FROM alumni 
              WHERE angkatan = $1 
              GROUP BY angkatan`
    err := config.DB.QueryRow(query, angkatan).Scan(&jumlah.Angkatan, &jumlah.JumlahAngkatan)
    if err != nil {
        return nil, err
    }
    return jumlah, nil
}

func CreateAlumni(alumni *model.Alumni) error {
	query := `INSERT INTO alumni (nim, full_name, angkatan, id_fakultas, id_prodi, tahun_lulus, sumber)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := config.DB.Exec(query, alumni.NIM, alumni.Nama, alumni.Angkatan, alumni.IDFakultas, alumni.IDProdi,
		alumni.TahunLulus, alumni.Sumber)
	return err
}

func UpdateAlumni(nim string, alumni *model.Alumni) error {
	query := `UPDATE alumni SET full_name=$1, angkatan=$2, tahun_lulus=$3, id_fakultas=$4, id_prodi=$5, sumber=$6
        WHERE nim=$7`
	_, err := config.DB.Exec(query, alumni.Nama, alumni.Angkatan, alumni.TahunLulus, alumni.IDFakultas, alumni.IDProdi,
		alumni.Sumber, nim)
	return err
}

func DeleteAlumni(nim string) error {
	query := `DELETE FROM alumni WHERE nim=$1`
	_, err := config.DB.Exec(query, nim)
	return err
}

func GetAllAlumni() ([]model.Alumni, error) {
	query := `SELECT nim, full_name, angkatan, tahun_lulus, id_fakultas, id_prodi, sumber FROM alumni`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alumniList []model.Alumni
	for rows.Next() {
		alumni := model.Alumni{}
		err := rows.Scan(&alumni.NIM, &alumni.Nama, &alumni.Angkatan, &alumni.TahunLulus, &alumni.IDFakultas, &alumni.IDProdi,
			&alumni.Sumber)
		if err != nil {
			return nil, err
		}
		alumniList = append(alumniList, alumni)
	}
	return alumniList, nil
}
