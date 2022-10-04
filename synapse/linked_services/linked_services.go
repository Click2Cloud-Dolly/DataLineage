package linked_services

import (
	"azureDataLineage/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetLinkedServices(auth configs.Auth, workspaceEndpoint, linkedServiceName string) (*LinkedService, error) {
	var lService *LinkedService
	url := strings.Join([]string{workspaceEndpoint, "/linkedservices/", linkedServiceName, "?api-version=2020-12-01"}, "")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	token := strings.Join([]string{"Bearer", auth.AccessToken}, " ")
	req.Header.Add("Authorization", token)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("ls:", string(body))

	err = json.Unmarshal(body, &lService)
	if err != nil {
		return nil, err
	}

	return lService, nil
}
