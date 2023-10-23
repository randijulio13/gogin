package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchData(c *gin.Context) {
	url := "http://localhost:3002/test"
	jsonBytes, err := json.Marshal(map[string]interface{}{
		"name": "Randi",
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request."})
		return
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error making post request."})
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not read response body."})
		return
	}

	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	c.JSON(http.StatusOK, responseBody)
}
