package workspaces

import (
	"azureDataLineage/configs"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func ListWorkspaces(auth configs.Auth, cred configs.Cred) (*WorkspaceList, error) {
	//x := "{subscriptionId}/providers/Microsoft.Synapse/workspaces?api-version=2021-06-01"
	var workspaceList *WorkspaceList
	url := strings.Join([]string{"https://management.azure.com/subscriptions/", cred.SubscriptionId, "/providers/Microsoft.Synapse/workspaces?api-version=2021-06-01"}, "")
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {

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
	err = json.Unmarshal(body, &workspaceList)
	if err != nil {
		return nil, err
	}
	return workspaceList, nil
}

func GetWorkspace(auth configs.Auth, cred configs.Cred, resourceGroup, workspaceName string) (*Workspace, error) {

	var workspaceDet *Workspace
	url := strings.Join([]string{"https://management.azure.com/subscriptions/", cred.SubscriptionId, "/resourceGroups/", resourceGroup, "/providers/Microsoft.Synapse/workspaces/", workspaceName, "?api-version=2021-06-01"}, "")
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {

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
	err = json.Unmarshal(body, &workspaceDet)
	if err != nil {
		return nil, err
	}
	return workspaceDet, nil
}
