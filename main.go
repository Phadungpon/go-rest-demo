package main

import (
	"fmt"
	_ "fmt"
	"go-rest-api/config"
	"go-rest-api/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("can not load goenv")
	}
    config.InitDb()  
	defer config.CloseDb()

	r := gin.Default()
	r.Static("/uploads", "./uploads")

	uploadDirs := [...]string{"aricles", "users"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755)
	}

	routes.Serve(r)
	r.Run(":"+os.Getenv("PORT"))
}
