	package service

	import (
		"database/sql"
		"tugas/domain/model"
		."tugas/domain/repository"

		"github.com/gofiber/fiber/v2"
	)

	func CheckPerkajaanAlumniService(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "ID wajib diisi",
				"success": false,
			})
		}

		perkajaan, err := CheckPerkajaanAlumniByID(id)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message":  "Data perkajaan alumni tidak ditemukan",
					"success":  true,
					"exists":   false,
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal cek perkajaan alumni karena " + err.Error(),
				"success": false,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":    "Berhasil mendapatkan data perkajaan alumni",
			"success":    true,
			"exists":     true,
			"perkajaan":  perkajaan,
		})
	}

	func CreatePerkajaanAlumniService(c *fiber.Ctx) error {
		var perkajaan model.PerkajaanAlumni
		if err := c.BodyParser(&perkajaan); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
				"success": false,
			})
		}

		if perkajaan.IDAlumni == "" || perkajaan.StatusKerja == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "IDAlumni dan StatusKerja wajib diisi",
				"success": false,
			})
		}

		if err := CreatePerkajaanAlumni(&perkajaan); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal membuat perkajaan alumni karena " + err.Error(),
				"success": false,
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message":   "Berhasil membuat data perkajaan alumni",
			"success":   true,
			"perkajaan": perkajaan,
		})
	}

	func UpdatePerkajaanAlumniService(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "ID wajib diisi",
				"success": false,
			})
		}

		var perkajaan model.PerkajaanAlumni
		if err := c.BodyParser(&perkajaan); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
				"success": false,
			})
		}

		if perkajaan.ID == 10 || perkajaan.StatusKerja == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "ID dan StatusKerja wajib diisi",
				"success": false,
			})
		}

		if err := UpdatePerkajaanAlumni(id, &perkajaan); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal update perkajaan alumni karena " + err.Error(),
				"success": false,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":   "Berhasil update data perkajaan alumni",
			"success":   true,
			"perkajaan": perkajaan,
		})
	}

	func DeletePerkajaanAlumniService(c *fiber.Ctx) error {
		idAlumni := c.Params("id_alumni")
		if idAlumni == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "IDAlumni wajib diisi",
				"success": false,
			})
		}

		if err := DeletePerkajaanAlumni(idAlumni); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal menghapus perkajaan alumni karena " + err.Error(),
				"success": false,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Berhasil menghapus data perkajaan alumni",
			"success": true,
		})
	}

	func GetAllPerkajaanAlumniService(c *fiber.Ctx) error {
		perkajaanList, err := GetAllPerkajaanAlumni()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal mendapatkan daftar perkajaan alumni karena " + err.Error(),
				"success": false,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":   "Berhasil mendapatkan daftar perkajaan alumni",
			"success":   true,
			"perkajaan": perkajaanList,
		})
	}