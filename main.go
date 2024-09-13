package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)

	if err == nil {
		println("Nenhum erro ")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {

			switch v.Tag() {
			case "required":
				println(v.StructField() + " is required")
			case "min":
				println(v.StructField() + " is less than the minimum: " + v.Param())
			case "max":
				println(v.StructField() + " is greater than the maximum: " + v.Param())
			case "email":
				println(v.StructField() + " is not a valid email")
			}
		}
	}
}
