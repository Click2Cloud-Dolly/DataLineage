package datasets

import (
	"azureDataLineage/configs"
	"azureDataLineage/synapse/linked_services"
	"encoding/json"
	"fmt"
	"strings"
)

func (s *AzureSQLTableDetails) GetDataset(body []byte, auth configs.Auth, workspaceEndpoint string, mappings *[]Mappings) error {
	fmt.Println("enter ")
	err := json.Unmarshal(body, &s)
	if err != nil {
		return err
	}
	var newMap Mappings
	inlinked, err := linked_services.GetLinkedServices(auth, workspaceEndpoint, s.Properties.LinkedServiceName.ReferenceName)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	fmt.Println("check ls:", inlinked)
	cpMap := *mappings
	newMap.Name = s.Properties.TypeProperties.Table
	newMap.Type = s.Properties.Type
	newMap.Location = GetConnstring(inlinked.Properties.TypeProperties.ConnectionString)
	newMap.Source = ""
	newMap.TDestination = append(newMap.TDestination, strings.Join([]string{workspaceEndpoint, cpMap[len(cpMap)-1].Name}, ";"))
	cpMap[len(cpMap)-1].Source = strings.Join([]string{GetConnstring(inlinked.Properties.TypeProperties.ConnectionString), s.Properties.TypeProperties.Table}, ";")
	*mappings = append(*mappings, newMap)
	fmt.Println("mappings sqltable", mappings)

	return nil
}
func GetConnstring(connString string) string {
	var source string
	databasedet := strings.Split(connString, ";")
	//fmt.Println("connstring", databasedet)
	for _, datadet := range databasedet {
		if strings.HasPrefix(strings.ToLower(datadet), "data source") || strings.HasPrefix(strings.ToLower(datadet), "initial catalog") {
			str := strings.Split(datadet, "=")[1]
			if len(source) != 0 {
				source = strings.Join([]string{source, str}, ",")
			} else {
				source = str
			}

		}
	}
	return source
}
