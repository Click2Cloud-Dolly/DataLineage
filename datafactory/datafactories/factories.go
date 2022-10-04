package datafactories

import (
	"azureDataLineage/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ListDataFactory(auth configs.Auth, cred configs.Cred) (*DataFactoryList, error) {

	var datafactorylist *DataFactoryList
	url := strings.Join([]string{"https://management.azure.com/subscriptions/", cred.SubscriptionId, "/providers/Microsoft.DataFactory/factories?api-version=2018-06-01"}, "")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return datafactorylist, err
	}
	token := strings.Join([]string{"Bearer", auth.AccessToken}, " ")
	req.Header.Add("Authorization", token)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return datafactorylist, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return datafactorylist, err
	}
	err = json.Unmarshal(body, &datafactorylist)
	if err != nil {
		return nil, err
	}
	return datafactorylist, nil
}

func GetDataFactory(auth configs.Auth, cred configs.Cred, resourceGroups string, datafactoryName string) (*DataFactory, error) {
	var datafactory *DataFactory
	url := strings.Join([]string{"https://management.azure.com/subscriptions/", cred.SubscriptionId, "/resourceGroups/", resourceGroups, "/providers/Microsoft.DataFactory/factories/", datafactoryName, "?api-version=2018-06-01"}, "")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return datafactory, err
	}
	token := strings.Join([]string{"Bearer", auth.AccessToken}, " ")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return datafactory, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return datafactory, err
	}
	err = json.Unmarshal(body, &datafactory)
	if err != nil {
		return nil, err
	}
	return datafactory, nil
}
