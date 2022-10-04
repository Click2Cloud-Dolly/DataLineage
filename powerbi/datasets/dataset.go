package datasets

import (
	"azureDataLineage/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetDatasets(auth configs.Auth, groupId string) (*PowerBiDataset, error) {

	var datasets PowerBiDataset
	url := strings.Join([]string{"https://api.powerbi.com/v1.0/myorg/groups", groupId, "datasets"}, "/")
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
	err = json.Unmarshal(body, &datasets)
	if err != nil {
		return nil, err
	}

	fmt.Printf("datasets body:%+v \n", string(body))
	return &datasets, nil
}
func GetDataset(auth configs.Auth, groupId string) (*PowerBiDataset, error) {

	var datasets PowerBiDataset
	url := strings.Join([]string{"https://api.powerbi.com/v1.0/myorg/groups", groupId, "datasets"}, "/")
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
	err = json.Unmarshal(body, &datasets)
	if err != nil {
		return nil, err
	}

	fmt.Printf("datasets body:%+v \n", string(body))
	return &datasets, nil
}

func GetDataSources(auth configs.Auth, groupId, datasetId string) (*PowerBiDatasource, error) {

	var datasource PowerBiDatasource
	url := strings.Join([]string{"https://api.powerbi.com/v1.0/myorg/groups", groupId, "datasets", datasetId, "datasources"}, "/")
	fmt.Println("check url:", url)
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
	//fmt.Println("datasource body:", string(body))
	err = json.Unmarshal(body, &datasource)
	if err != nil {
		return nil, err
	}
	return &datasource, nil
}
