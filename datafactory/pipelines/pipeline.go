package pipelines

import (
	"azureDataLineage/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ListPipelines(auth configs.Auth, cred configs.Cred, resourceGroups string, datafactoryName string) (*DFPipelineList, error) {
	var pipelinelist *DFPipelineList

	url := strings.Join([]string{"https://management.azure.com/subscriptions/", cred.SubscriptionId, "/resourceGroups/", resourceGroups, "/providers/Microsoft.DataFactory/factories/", datafactoryName, "/pipelines/", "?api-version=2018-06-01"}, "")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return pipelinelist, err
	}
	token := strings.Join([]string{"Bearer", auth.AccessToken}, " ")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return pipelinelist, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return pipelinelist, err
	}
	err = json.Unmarshal(body, &pipelinelist)
	if err != nil {
		return nil, err
	}

	return pipelinelist, nil
}

func GetPipeline(auth configs.Auth, cred configs.Cred, resourceGroups string, datafactoryName string, pipelineName string) (*DFGetPipelineDet, error) {
	var pipelinedetails *DFGetPipelineDet
	url := strings.Join([]string{"https://management.azure.com/subscriptions/", cred.SubscriptionId, "/resourceGroups/", resourceGroups, "/providers/Microsoft.DataFactory/factories/", datafactoryName, "/pipelines/", pipelineName, "?api-version=2018-06-01"}, "")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return pipelinedetails, err
	}
	token := strings.Join([]string{"Bearer", auth.AccessToken}, " ")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return pipelinedetails, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return pipelinedetails, err
	}
	err = json.Unmarshal(body, &pipelinedetails)
	if err != nil {
		return nil, err
	}
	return pipelinedetails, nil
}
