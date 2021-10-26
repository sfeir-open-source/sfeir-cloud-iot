package db

import (
	"context"
	"log"

	"github.com/sfeir-cloud-iot/bicycle-api/models"
)

func GetAllData() (map[string]models.BicycleData, error) {
	ctx := context.Background()
	var data map[string]models.BicycleData

	ref := Db.NewRef("bicycle_data")

	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
		return nil, err
	}

	return data, nil
}
