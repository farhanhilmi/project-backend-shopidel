package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
)

func GetRajaOngkirCost(reqData dtousecase.CheckDeliveryFeeRequest) ([]dtousecase.DeliveryFeeResponse, error) {
	formData := url.Values{
		"origin":      {reqData.Origin},
		"destination": {reqData.Destination},
		"weight":      {reqData.Weight},
		"courier":     {reqData.Courier},
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/cost", config.GetEnv("RAJA_ONGKOR_ENDPOINT")), bytes.NewBufferString(formData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("key", config.GetEnv("RAJA_ONGKIR_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	var response dtousecase.RajaOngkirFee

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, nil
	}

	return response.RajaOngkir.Results, nil
}
