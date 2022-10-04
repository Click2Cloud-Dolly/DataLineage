package reports

import (
	"azureDataLineage/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetReports(auth configs.Auth, groupId string) (*PowerBiReports, error) {

	var reports PowerBiReports
	url := strings.Join([]string{"https://api.powerbi.com/v1.0/myorg/groups", groupId, "reports"}, "/")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", auth.AccessToken}, " "))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &reports)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return &reports, nil
}

func GetReport(auth configs.Auth, groupId, reportID string) (*PowerBIReport, error) {

	var reports PowerBIReport
	url := strings.Join([]string{"https://api.powerbi.com/v1.0/myorg/groups", groupId, "reports", reportID}, "/")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", auth.AccessToken}, " "))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &reports)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return &reports, nil
}
