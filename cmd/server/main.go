package main


import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"miniauth/internal/config"
	"miniauth/internal/db"

)

func main(){
	//Load .env variables
	if err := godotenv.Load(); err != nil{
		log.Println("No .env file found")
	}


	cfg,err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	r := gin.Default()


	r.GET("/health", func( c *gin.Context){
		c.JSON(http.StatusOK, gin.H {
			"status": "OK",
		})
	})

	r.Run(":"+ cfg.AppPort)
}