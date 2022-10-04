package workspaces

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetWorkspaceInfo(assertionToken string) (WorkspaceInfo, error) {

	fmt.Println("assertionToken:", assertionToken)
	var workInfo WorkspaceInfo
	url := strings.Join([]string{"https://api.powerbi.com/v1.0/myorg/admin/workspaces/getInfo?datasetExpressions=true&datasetSchema=true&datasourceDetails=true&getArtifactUsers=true&lineage=true"}, "/")
	//url := "https://api.powerbi.com/v1.0/myorg/admin/workspaces/getInfo?datasetExpressions=true&datasetSchema=true&datasourceDetails=true&getArtifactUsers=true&lineage=true"
	method := "POST"

	payload := strings.NewReader(`{` + "" + `"workspaces": [` + "" + `"9224e2c1-c244-425e-b9d9-355a88498df4"` + "" + `]` + "" + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return workInfo, err
	}

	token := strings.Join([]string{"Bearer ", assertionToken}, "")
	fmt.Println("Token of WKSINFO:", token)
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return workInfo, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return workInfo, err
	}
	//fmt.Println("The value in WorkspaceInfo:", string(body))

	err = json.Unmarshal(body, &workInfo)
	if err != nil {
		fmt.Println(err)
		return workInfo, err
	}
	return workInfo, nil
}

func GetStatus(workspacescanId *WorkspaceInfo, assertionToken string) (ScanStatus, error) {

	var scanstatus ScanStatus

	//fmt.Println("Check: ", workspaceId.ID)
	url := strings.Join([]string{"https://api.powerbi.com/v1.0/myorg/admin/workspaces/scanStatus/", workspacescanId.ID}, "/")
	fmt.Println("URL of status:", url)
	//url := "https://api.powerbi.com/v1.0/myorg/admin/workspaces/scanStatus/278bc376-55bb-4d9e-a81b-feab1d9cdb0b"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return scanstatus, err
	}
	token := strings.Join([]string{"Bearer ", assertionToken}, "")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return scanstatus, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return scanstatus, err
	}
	//fmt.Println("The value in ScanStatus:",string(body))

	err = json.Unmarshal(body, &scanstatus)
	if err != nil {
		fmt.Println(err)
		return scanstatus, err
	}

	return scanstatus, err
}

func GetTableDetails(workspacescanId *WorkspaceInfo, assertionToken string) (ScanResult, error) {

	var scanresult ScanResult

	url := strings.Join([]string{"https://api.powerbi.com/v1.0/myorg/admin/workspaces/scanResult/", workspacescanId.ID}, "/")
	//url := "https://api.powerbi.com/v1.0/myorg/admin/workspaces/scanResult/278bc376-55bb-4d9e-a81b-feab1d9cdb0b"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return scanresult, err
	}

	token := strings.Join([]string{"Bearer ", assertionToken}, "")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return scanresult, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return scanresult, err
	}

	err = json.Unmarshal(body, &scanresult)
	if err != nil {
		fmt.Println(err)
		return scanresult, err
	}
	return scanresult, err
}
