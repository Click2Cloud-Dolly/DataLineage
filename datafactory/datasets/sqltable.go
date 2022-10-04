package datasets

import (
	"azureDataLineage/configs"
	"azureDataLineage/datafactory/linkedservice"
	"encoding/json"
	"fmt"
	"strings"
)

func (s *DfAzureSQLTableDetails) GetDataset(body []byte, accessToken configs.Auth, cred configs.Cred, resourceGroups string, datafactoryName string, mappings *[]DFMappings) error {
	err := json.Unmarshal(body, &s)
	if err != nil {
		return err
	}
	var newMap DFMappings
	inlinked, err := linkedservice.GetDFLinkedService(accessToken, cred, resourceGroups, datafactoryName, s.Properties.LinkedServiceName.ReferenceName)
	if err != nil {
		return err
	}
	cpMap := *mappings
	newMap.Name = s.Properties.TypeProperties.Table
	newMap.Type = s.Properties.Type
	newMap.Location = GetConnstring(inlinked.Properties.TypeProperties.ConnectionString)
	newMap.Source = ""
	newMap.TDestination = append(newMap.TDestination, strings.Join([]string{datafactoryName, cpMap[len(cpMap)-1].Name}, ";"))
	cpMap[len(cpMap)-1].Source = strings.Join([]string{newMap.Location, newMap.Name}, ";")
	fmt.Println("SQLName:", newMap.Name, "SQLType:", newMap.Type, "SQLlocation:", newMap.Location)
	*mappings = append(*mappings, newMap)
	fmt.Println("mappings sqltable:", mappings)
	return nil
}

func GetConnstring(connString string) string {
	var source string
	databasedet := strings.Split(connString, ";")
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
