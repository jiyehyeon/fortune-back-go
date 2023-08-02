package repo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Response struct {
	Header struct {
		ResultCode 		string 	`xml:"resultCode"`
		ResultMsg  		string 	`xml:"resultMsg"`
	} `xml:"header"`
	Body struct {
		Items struct {
			Item struct{
				LunIljin 	string 	`xml:"lunIljin"`
			} `xml:"item"`
		} `xml:"items"`
		NumOfRows  		int 		`xml:"numOfRows"`
		PageNo     		int 		`xml:"pageNo"`
		TotalCount 		int 		`xml:"totalCount"`
	} `xml:"body"`
}

type ApiConfig struct {
	ApiURL     string
	ServiceKey string
}

func getApiConfig() *ApiConfig{
	return &ApiConfig{
		ApiURL: os.Getenv("LUNAR_CALANDER_API_URL"),
		ServiceKey: os.Getenv("OPEN_API_KEY"),
	} 
}

func GetGanjiWithSolar(year, month, day string) (string, error){
	endpoint := "getLunCalInfo"

	apiConfig := getApiConfig()
	apiURL := apiConfig.ApiURL
	serviceKey := apiConfig.ServiceKey

	u, err := url.Parse(fmt.Sprintf("%s/%s", apiURL, endpoint))
	if err != nil {
		return "", err
	}

	params := url.Values{
		"serviceKey": {serviceKey},
		"solYear":    {year},
		"solMonth":   {month},
		"solDay":     {day},
	}

	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(fmt.Sprintf("Failed to get ganji - %s", resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response Response
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	ganji := response.Body.Items.Item.LunIljin

	return ganji, nil
}

func GetGanjiWithLunar(year, month, day string) (string, error){
	endpoint := "getSolCalInfo"

	apiConfig := getApiConfig()
	apiURL := apiConfig.ApiURL
	serviceKey := apiConfig.ServiceKey

	u, err := url.Parse(fmt.Sprintf("%s/%s", apiURL, endpoint))
	if err != nil {
		return "", err
	}

	params := url.Values{
		"serviceKey": {serviceKey},
		"lunYear":    {year},
		"lunMonth":   {month},
		"lunDay":     {day},
	}

	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(fmt.Sprintf("Failed to get ganji - %s", resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response Response
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	ganji := response.Body.Items.Item.LunIljin

	return ganji, nil
}
