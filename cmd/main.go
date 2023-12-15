package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func gateway(c *gin.Context) {

	relay := c.Query("relay")

	os.Chdir("../fluence")
	cmd := exec.Command("../cli/node_modules/.bin/fluence", "run", "-f", `helloWorld("wonder")`, "--relay="+relay)

	output, _ := cmd.Output()
	response := map[string]interface{}{
		"msg": string(output),
	}
	
	c.JSON(http.StatusOK, response)
	
}

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Adjust this to your needs
		AllowMethods:     []string{"POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/api", gateway)

	log.Fatal(router.Run(":8000"))
	
}
