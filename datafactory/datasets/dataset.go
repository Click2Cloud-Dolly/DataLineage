package datasets

import (
	"azureDataLineage/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type DatasetInterface interface {
	GetDataset(body []byte, auth configs.Auth, workspaceEndpoint string, resourceGroups string, datafactoryName string, mappings *[]DFMappings) error
}

func (d *Dataset) GetDatasets(auth configs.Auth, cred configs.Cred, resourceGroups string, datafactoryName string, datasetName string, mapping *[]DFMappings) error {
	var listDatasets *DFListDatasets
	url := strings.Join([]string{"https://management.azure.com/subscriptions/", cred.SubscriptionId, "/resourceGroups/", resourceGroups, "/providers/Microsoft.DataFactory/factories/", datafactoryName, "/datasets/", datasetName, "?api-version=2018-06-01"}, "")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	token := strings.Join([]string{"Bearer", auth.AccessToken}, " ")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(body, &listDatasets)
	if err != nil {
		return err
	}
	fmt.Println("dataset type:", datasetName, listDatasets.Properties.Type, string(body))

	if listDatasets.Properties.Type == Parquet { // for outputs
		fmt.Println("Parquet")
		err = d.DfParquet.GetDataset(body, auth, cred, resourceGroups, datafactoryName, mapping)
		if err != nil {
			return err
		}

	} else if listDatasets.Properties.Type == AzureSQLTable { // for inputs
		fmt.Println("AzureSQLTable")
		err = d.DfSQLTable.GetDataset(body, auth, cred, resourceGroups, datafactoryName, mapping)
		if err != nil {
			return err
		}
	}
	return nil
}
