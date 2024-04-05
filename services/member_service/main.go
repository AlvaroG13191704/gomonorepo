package main

import (
	"fmt"
	"log"

	"github.com/AlvaroG13191704/gomonorepo/shared/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello, World!")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(createUser())
	})

	app.Listen(":3000")
}

func createUser() models.Member {
	// create a doctor from the models package from shared
	doctor := models.Member{
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

	log.Println(doctor)

	return doctor
}
