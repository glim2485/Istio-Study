package main

import (
	"net/http"
	"encoding/json"
	"io"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"fmt"
)

type serverResponse struct {
	SecretCode	string `json:"secretcode"`
}

func main() {
	uniqueID := uuid.New().String()
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "serviceA",
			"unique_id": uniqueID,
		})
	})

	router.GET("/secretcode", func(c *gin.Context) {
		secretcode := getSecretCode()
		c.JSON(http.StatusOK, gin.H{
			"message": "serviceA",
			"secretcode": secretcode,
		})
	})

	router.Run(":6700")
}

func getSecretCode() string {
	url := "http://servicec.servicec.svc.cluster.local:6702/getsecretcode"
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var response serverResponse
	json.Unmarshal(body, &response)
	fmt.Println(response)
	return response.SecretCode
}