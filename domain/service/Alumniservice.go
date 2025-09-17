package service

import (
	"database/sql"
	"strconv"
	"tugas/domain/model"
	"tugas/domain/repository"

	"github.com/gofiber/fiber/v2"
)

func CheckAlumniService(c *fiber.Ctx) error {
    nim := c.FormValue("nim")
    if nim == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "NIM wajib diisi",
            "success": false,
        })
    }
    alumni, err := repository.CheckAlumniByNim(nim)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.Status(fiber.StatusOK).JSON(fiber.Map{
                "message":  "Mahasiswa bukan alumni",
                "success":  true,
                "isAlumni": false,
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal cek alumni karena " + err.Error(),
            "success": false,
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message":  "Berhasil mendapatkan data alumni",
        "success":  true,
        "isAlumni": true,
        "alumni":   alumni,
    })
}

func GetJumlahAlumniService(c *fiber.Ctx) error {
    jumlahAngkatanStr := c.Query("angkatan")
    if jumlahAngkatanStr == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "angkatan wajib diisi",
            "success": false,
        })
    }

    jumlahAngkatan, err := strconv.Atoi(jumlahAngkatanStr)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "angkatan harus berupa angka",
            "success": false,
        })
    }

    alumni, err := repository.GetJumlahAlumni(jumlahAngkatan)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.Status(fiber.StatusOK).JSON(fiber.Map{
                "message": "Data alumni tidak ditemukan",
                "success": true,
                "exists":  false,
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal cek alumni karena " + err.Error(),
            "success": false,
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Berhasil mendapatkan data alumni",
        "success": true,
        "exists":  true,
        "alumni":  alumni,
    })
}


func CreateAlumniService(c *fiber.Ctx) error {
    var alumni model.Alumni
    if err := c.BodyParser(&alumni); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
            "success": false,
        })
    }

    if alumni.NIM == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "NIM wajib diisi",
            "success": false,
        })
    }

    if err := repository.CreateAlumni(&alumni); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal membuat alumni karena " + err.Error(),
            "success": false,
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Berhasil membuat data alumni",
        "success": true,
        "alumni":  alumni,
    })
}

func UpdateAlumniService(c *fiber.Ctx) error {
    nim := c.Params("nim")
    if nim == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "NIM wajib diisi",
            "success": false,
        })
    }

    var alumni model.Alumni
    if err := c.BodyParser(&alumni); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
            "success": false,
        })
    }

    if err := repository.UpdateAlumni(nim, &alumni); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal update alumni karena " + err.Error(),
            "success": false,
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Berhasil update data alumni",
        "success": true,
        "alumni":  alumni,
    })
}

func DeleteAlumniService(c *fiber.Ctx) error {
    nim := c.Params("nim")
    if nim == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "NIM wajib diisi",
            "success": false,
        })
    }

    if err := repository.DeleteAlumni(nim); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal menghapus alumni karena " + err.Error(),
            "success": false,
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Berhasil menghapus data alumni",
        "success": true,
    })
}

func GetAllAlumniService(c *fiber.Ctx) error {
    alumniList, err := repository.GetAllAlumni()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Gagal mendapatkan daftar alumni karena " + err.Error(),
            "success": false,
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Berhasil mendapatkan daftar alumni",
        "success": true,
        "alumni":  alumniList,
    })
}