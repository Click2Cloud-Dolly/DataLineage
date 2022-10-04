package datasets

import (
	"azureDataLineage/configs"
	"azureDataLineage/datafactory/linkedservice"
	"encoding/json"
	"fmt"
	"strings"
)

func (dfParquet *DfParquetDetails) GetDataset(body []byte, accessToken configs.Auth, cred configs.Cred, resourceGroups string, datafactoryName string, mappings *[]DFMappings) error {
	var newMap DFMappings
	err := json.Unmarshal(body, &dfParquet)
	if err != nil {
		return err
	}
	parquetLinked, err := linkedservice.GetDFLinkedService(accessToken, cred, resourceGroups, datafactoryName, dfParquet.Properties.LinkedServiceName.ReferenceName)
	if err != nil {
		return err
	}
	cpMap := *mappings
	newMap.Type = dfParquet.Properties.Type
	newMap.Name = dfParquet.Properties.TypeProperties.Location.FileName
	newMap.Location = strings.Join([]string{parquetLinked.Properties.TypeProperties.Url, dfParquet.Properties.TypeProperties.Location.FileSystem}, "/")
	newMap.Source = strings.Join([]string{datafactoryName, dfParquet.Name}, ";")
	fmt.Println("type:", newMap.Type, "Filename:", newMap.Name, "Location:", newMap.Location)
	cpMap[len(cpMap)-1].TDestination = append(cpMap[len(cpMap)-1].TDestination, strings.Join([]string{newMap.Location, newMap.Name}, ";"))
	*mappings = append(*mappings, newMap)
	fmt.Println("mappings Parquet:", mappings)

	return nil
}
