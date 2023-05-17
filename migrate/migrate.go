package main

import (
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/infrastructure"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/initializer"
)

func init() {
	initializer.LoadEnvVariables()
	infrastructure.ConnectDB()

}

func main() {
	err := infrastructure.DB.AutoMigrate(&entity.Post{})
	if err != nil {
		panic("failed to migrate")
	}
}
