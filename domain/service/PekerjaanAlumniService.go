package service

import (
	"database/sql"
	"tugas/domain/model"
	. "tugas/domain/repository"

	"github.com/gofiber/fiber/v2"
)

func CheckpekerjaanAlumniService(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID tidak ditemukan",
			"success": false,
		})
	}

	pekerjaan, err := CheckpekerjaanAlumniByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Data pekerjaan alumni tidak ditemukan",
				"success": true,
				"exists":  false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal cek pekerjaan alumni karena " + err.Error(),
			"success": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Berhasil mendapatkan data pekerjaan alumni",
		"success":   true,
		"exists":    true,
		"pekerjaan": pekerjaan,
	})
}

func CreatepekerjaanAlumniService(c *fiber.Ctx) error {
	var pekerjaan model.PekerjaanAlumni
	if err := c.BodyParser(&pekerjaan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"success": false,
		})
	}

	if pekerjaan.NimAlumni == "" || pekerjaan.StatusKerja == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "IDAlumni dan StatusKerja wajib diisi",
			"success": false,
		})
	}

	if err := CreatepekerjaanAlumni(&pekerjaan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat pekerjaan alumni karena " + err.Error(),
			"success": false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "Berhasil membuat data pekerjaan alumni",
		"success":   true,
		"pekerjaan": pekerjaan,
	})
}

func UpdatepekerjaanAlumniService(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID wajib diisi",
			"success": false,
		})
	}

	var pekerjaan model.PekerjaanAlumni
	if err := c.BodyParser(&pekerjaan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"success": false,
		})
	}

	if pekerjaan.ID == 10 || pekerjaan.StatusKerja == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID dan StatusKerja wajib diisi",
			"success": false,
		})
	}

	if err := UpdatepekerjaanAlumni(id, &pekerjaan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal update pekerjaan alumni karena " + err.Error(),
			"success": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Berhasil update data pekerjaan alumni",
		"success":   true,
		"pekerjaan": pekerjaan,
	})
}

func DeletepekerjaanAlumniService(c *fiber.Ctx) error {
	idAlumni := c.Params("nim_alumni")
	if idAlumni == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "IDAlumni wajib diisi",
			"success": false,
		})
	}

	if err := DeletepekerjaanAlumni(idAlumni); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus pekerjaan alumni karena " + err.Error(),
			"success": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menghapus data pekerjaan alumni",
		"success": true,
	})
}

func GetAllpekerjaanAlumniService(c *fiber.Ctx) error {
	pekerjaanList, err := GetAllpekerjaanAlumni()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendapatkan daftar pekerjaan alumni karena " + err.Error(),
			"success": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Berhasil mendapatkan daftar pekerjaan alumni",
		"success":   true,
		"pekerjaan": pekerjaanList,
	})
}
