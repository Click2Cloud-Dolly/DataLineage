package datasets

import (
	"azureDataLineage/configs"
	"azureDataLineage/synapse/linked_services"

	//"azureDataLineage/synapse/sql_pools"
	"encoding/json"
	"fmt"
	"strings"
)

func (pt *SqlPoolTableDetails) GetDataset(body []byte, auth configs.Auth, workspaceEndpoint string, mappings *[]Mappings) error {
	var newMap Mappings
	err := json.Unmarshal(body, &pt)
	if err != nil {
		return err
	}

	fmt.Printf("\nSQLPoolProp: %v", pt)

	//fmt.Println("check ls:", inlinked)
	//cred, err := config.GetEnv()
	//if err != nil {
	//	return err
	//}
	////Get Token
	//auth, err := cred.GetToken()
	//if err != nil {
	//	return err
	//}
	//fmt.Println("\nGetTokenpool:", auth.AccessToken)
	//wlist, err := workspaces.ListWorkspaces(auth, cred)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("\nworkspacelist: ", wlist)
	//
	//var sqlep string
	//for _, ep := range wlist.Value {
	//
	//	sqlep = ep.Properties.ConnectivityEndpoints.Sql
	//}
	//fmt.Println("sqlep: ", sqlep)
	//pooldbname := strings.Join([]string{pt.Properties.SqlPool.ReferenceName}, "/")
	cpMap := *mappings
	newMap.Name = pt.Properties.TypeProperties.Table
	newMap.Type = pt.Properties.Type
	newMap.Source = strings.Join([]string{workspaceEndpoint, pt.Name}, ";")
	//newMap.Destination = nil
	inlinked, err := linked_services.GetLinkedServices(auth, workspaceEndpoint, pt.Properties.LinkedServiceName.ReferenceName)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	fmt.Println("checking outPut:007:", inlinked)
	fmt.Println("check123007:", GetConnstring(inlinked.Properties.TypeProperties.ConnectionString))
	newMap.Location = strings.Join([]string{sqlEnd(workspaceEndpoint), pt.Properties.SqlPool.ReferenceName}, ",")
	//newMap.Location = GetConnstring(inlinked.Properties.TypeProperties.ConnectionString)
	cpMap[len(cpMap)-1].TDestination = append(cpMap[len(cpMap)-1].TDestination, strings.Join([]string{newMap.Location, pt.Properties.TypeProperties.Table}, ";"))

	*mappings = append(*mappings, newMap)
	fmt.Println("checkingmaps007:", mappings)
	//mappings.TSource = append(mappings.TSource, newMap)
	return nil

}
func sqlEnd(workspaceEnd string) string {

	endPoint := strings.Split(strings.Split(workspaceEnd, "//")[1], ".")
	endPoint[1] = "sql"
	newEndpoint := strings.Join(endPoint, ".")
	fmt.Println("checking newEndpoint", newEndpoint)
	return newEndpoint
}
