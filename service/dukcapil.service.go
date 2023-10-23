package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/randijulio13/gogin/config"
)

type DetailRequestBody struct {
	NIK      string `json:"nik"`
	USER_ID  string `json:"userId"`
	PASSWORD string `json:"password"`
	IP_USER  string `json:"ipUser"`
}

func Detail(nik string) {
	userId := config.Config("DUKCAPIL_USER_ID", "")
	password := config.Config("DUKCAPIL_PASSWORD", "")
	ipUser := config.Config("DUKCAPIL_IP_USER", "")

	detailEndpoint := fmt.Sprintf("dukcapil/get_json/%s/CALL_NIK_GET_KK", config.Config("DUKCAPIL_CONSUMER_ID", ""))
	url := fmt.Sprintf("%s/%s", config.Config("DUKCAPIL_BASE_URL", ""), detailEndpoint)

	detailRequest := DetailRequestBody{
		nik, userId, password, ipUser,
	}
	jsonBytes, err := json.Marshal(detailRequest)

	if err != nil {
		log.Fatal(fmt.Println("failed parse json data"))
		return
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))

	log.Printf("%s %s", res, err.Error())

}
