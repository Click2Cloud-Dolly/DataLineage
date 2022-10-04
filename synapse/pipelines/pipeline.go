package pipelines

import (
	"azureDataLineage/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ListPipelines(auth configs.Auth, workspaceEndpoint string) (*PipelinesList, error) {
	//url := "https://synapsepurviewworkspace.dev.azuresynapse.net/pipelines?api-version=2020-12-01"
	var pipelines *PipelinesList
	url := strings.Join([]string{workspaceEndpoint, "/pipelines?api-version=2020-12-01"}, "")
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
	err = json.Unmarshal(body, &pipelines)
	if err != nil {
		return nil, err
	}
	return pipelines, nil
}
func GetPipeline(auth configs.Auth, workspaceEndpoint, pipelineName string) (*GetPipelineDet, error) {
	var pipeline *GetPipelineDet
	url := strings.Join([]string{workspaceEndpoint, "/pipelines/", pipelineName, "?api-version=2020-12-01"}, "")
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
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("pipelineBody:", string(body))
	err = json.Unmarshal(body, &pipeline)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return pipeline, nil
}
