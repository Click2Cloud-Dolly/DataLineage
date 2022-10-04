package dashboards

import (
	"azureDataLineage/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetDashboard(auth configs.Auth, groupId string) (*PowerBiDashboards, error) {

	var dashboards PowerBiDashboards
	url := strings.Join([]string{"https://api.powerbi.com/v1.0/myorg/groups", groupId, "dashboards"}, "/")
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
	fmt.Println("dashboard body:", string(body))
	err = json.Unmarshal(body, &dashboards)
	if err != nil {
		return nil, err
	}

	return &dashboards, nil
}
