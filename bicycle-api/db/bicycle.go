package db

import (
	"context"
	"log"

	"github.com/sfeir-cloud-iot/bicycle-api/models"
)

func GetAllBicycleData() (map[string]models.BicycleData, error) {
	ctx := context.Background()
	var data map[string]models.BicycleData

	ref := Db.NewRef("bicycle_data")

	if err := ref.Get(ctx, &data); err != nil {
		log.Println("Error reading from database:", err)
		return nil, err
	}

	return data, nil
}

func GetLastBicycleData() (*models.BicycleData, error) {
	ctx := context.Background()
	var data map[string]models.BicycleData

	ref := Db.NewRef("bicycle_data")

	if err := ref.OrderByChild("time").LimitToLast(1).Get(ctx, &data); err != nil {
		log.Println("Error reading from database:", err)
		return nil, err
	}

	var bicycleLastData models.BicycleData
	for _, currentData := range data {
		bicycleLastData = currentData
	}

	return &bicycleLastData, nil
}

func GetBicycleDataBetweenDate(start string, end string) (map[string]models.BicycleData, error) {
	ctx := context.Background()
	var data map[string]models.BicycleData

	ref := Db.NewRef("bicycle_data")

	if err := ref.OrderByKey().StartAt(start).EndAt(end).Get(ctx, &data); err != nil {
		log.Println("Error reading from database:", err)
		return nil, err
	}

	return data, nil
}
