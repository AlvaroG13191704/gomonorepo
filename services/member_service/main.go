package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/velticare/gomonorepo/shared/models"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// create a doctor from the models package from shared
		doctor := models.Doctor{
			Uid: "123",
			Credentials: models.Credentials{
				Email:    "ga1318garcia@gmail.com",
				Password: "password",
			},
			Profile: models.Profile{
				Uid:        "123",
				Name:       "Gabe",
				Email:      "ga1318garcia@gmail.com",
				Phone:      "123-456-7890",
				Speciality: []string{"General"},
				DoctorID:   "1",
				Role:       "Doctor",
			},
			Status: "Active",
			Clinics: []models.Clinics{
				{
					Id:     "1",
					Status: "Active",
				},
			},
		}

		return c.JSON(doctor)
	})

	app.Listen(":8080")
}
