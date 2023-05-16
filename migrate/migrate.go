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
	infrastructure.DB.AutoMigrate(&entity.Post{})
}
