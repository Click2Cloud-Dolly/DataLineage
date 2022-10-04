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
	GetDataset(body []byte, auth configs.Auth, workspaceEndpoint string, mappings *[]Mappings) error
}

func (d *Dataset) GetDatasets(auth configs.Auth, workspaceEndpoint, datasetName string, mappings *[]Mappings) error {

	var dataset *GetDatasetDetails

	url := strings.Join([]string{workspaceEndpoint, "/datasets/", datasetName, "?api-version=2020-12-01"}, "")
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
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &dataset)
	if err != nil {
		return err
	}
	fmt.Println("dataset type:", datasetName, dataset.Properties.Type, string(body))

	if dataset.Properties.Type == Parquet { // for outputs
		fmt.Println("Parquet")
		err = d.Parquet.GetDataset(body, auth, workspaceEndpoint, mappings)
		if err != nil {
			return err
		}

		fmt.Println("check unmarshalled parquet:", d.Parquet)

	} else if dataset.Properties.Type == SqlPoolTable { // for outputs
		fmt.Println("SqlPoolTable")
		err = d.SQLPool.GetDataset(body, auth, workspaceEndpoint, mappings)
		if err != nil {
			return err
		}
		fmt.Println("check unmarshalled sqldw:", mappings)

	} else if dataset.Properties.Type == AzureSQLTable { // for inputs
		fmt.Println("AzureSQLTable")
		err = d.SQLTable.GetDataset(body, auth, workspaceEndpoint, mappings)
		if err != nil {
			return err
		}
		fmt.Println("check unmarshalled mappings:", mappings)

	} else if dataset.Properties.Type == AzureSQLDWTable { // for outputs
		fmt.Println("AzureSQLDWTable")
		err = d.SQLPool.GetDataset(body, auth, workspaceEndpoint, mappings)
		if err != nil {
			return err
		}
		fmt.Println("check unmarshalled sqldw:", mappings)

	}

	/*else if dataset.Properties.Type == AzureSQLDWTable { // for outputs
		fmt.Println("AzureSQLDWTable")
		//finaldataset.Graphs = append(finaldataset.Graphs, *graph1)
		err = d.DWSQLPool.GetDataset(body, &finaldataset.Graphs, auth, workspaceEndpoint, mappings)
		if err != nil {
			return err
		}
		fmt.Println("check unmarshalled sqldw:", finaldataset.SQLPool)

	}*/
	return nil
}
