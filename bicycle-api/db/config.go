package db

import (
	"context"
	"github.com/sfeir-cloud-iot/bicycle-api/models"
)

func GetCurrentConfig() (*models.ConfigData, error) {
	ctx := context.Background()
	var configData models.ConfigData

	ref := Db.NewRef("bicycle_config")

	if err := ref.Get(ctx, &configData); err != nil {
		return nil, err
	}

	return &configData, nil
}
