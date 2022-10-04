package linkedservice

import (
	"azureDataLineage/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetDFLinkedService(auth configs.Auth, cred configs.Cred, resourceGroups string, datafactoryName string, linkedServiceName string) (*LinkedService, error) {
	var linkedservice *LinkedService

	//url := "GET https://management.azure.com/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}?api-version=2018-06-01"
	url := strings.Join([]string{"https://management.azure.com/subscriptions/", cred.SubscriptionId, "/resourceGroups/", resourceGroups, "/providers/Microsoft.DataFactory/factories/", datafactoryName, "/linkedservices/", linkedServiceName, "?api-version=2018-06-01"}, "")

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return linkedservice, err
	}

	token := strings.Join([]string{"Bearer", auth.AccessToken}, " ")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return linkedservice, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return linkedservice, err
	}
	//fmt.Println("Get_linked_service:", string(body))

	err = json.Unmarshal(body, &linkedservice)
	if err != nil {
		return linkedservice, err
	}

	return linkedservice, nil
}
