package groups

import (
	"azureDataLineage/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetGroups(auth configs.Auth) (*Groups, error) {
	var groups Groups
	url := "https://api.powerbi.com/v1.0/myorg/groups"

	method := "GET"
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	fmt.Println("auth", auth.AccessToken)
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
	fmt.Println("Body:", string(body))
	err = json.Unmarshal(body, &groups)
	if err != nil {

		return nil, err
	}
	return &groups, nil
}
