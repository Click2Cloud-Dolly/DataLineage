package datasets

import (
	"azureDataLineage/configs"
	"azureDataLineage/synapse/linked_services"
	"encoding/json"
	"fmt"
	"strings"
)

func (pq *ParquetDetails) GetDataset(body []byte, accessToken configs.Auth, workspaceEndpoint string, mappings *[]Mappings) error {
	var newMap Mappings
	err := json.Unmarshal(body, &pq)
	if err != nil {
		return err
	}
	fmt.Printf("Check pq %+v\n", pq)
	parquetLinked, err := linked_services.GetLinkedServices(accessToken, workspaceEndpoint, pq.Properties.LinkedServiceName.ReferenceName)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	cpMap := *mappings
	newMap.Type = pq.Properties.Type
	newMap.Name = pq.Properties.TypeProperties.Location.FileName
	newMap.Location = strings.Join([]string{parquetLinked.Properties.TypeProperties.Url, pq.Properties.TypeProperties.Location.FileSystem}, "/")
	newMap.Source = strings.Join([]string{workspaceEndpoint, pq.Name}, ";")
	//newMap.Destination = nil

	cpMap[len(cpMap)-1].TDestination = append(cpMap[len(cpMap)-1].TDestination, strings.Join([]string{newMap.Location, newMap.Name}, ";"))
	*mappings = append(*mappings, newMap)
	//mappings.Destination = append(mappings.Destination, newMap)
	//mappings.TSource = append(mappings.TSource, newMap)
	fmt.Printf("Check pqls %+v\n", parquetLinked)

	return nil
}
