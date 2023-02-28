package model

import (
	"github.com/xavierxcn/apiserver/internal/serve/model/image"
	"github.com/xavierxcn/apiserver/internal/serve/pkg/storage"
)

// AutoMigrate auto migrate
func AutoMigrate() {
	var models = []interface{}{
		&image.TImage{},
	}

	err := storage.Client().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}
