package powerbi

import (
	"azureDataLineage/configs"
	"azureDataLineage/powerbi/datasets"
	"azureDataLineage/powerbi/groups"
	"azureDataLineage/powerbi/reports"
	"azureDataLineage/powerbi/workspaces"
	"encoding/json"

	//"azureDataLineage/powerbi/workspaces"
	"azureDataLineage/synapse"
	synDatasets "azureDataLineage/synapse/datasets"
	"fmt"
	"strings"
)

type powerbiPayload struct {
	WorkspaceIds []string `json:"workspaces"`
}

func DataLineagePowerBi(datas *synapse.Data, assertionToken string) (*synapse.Data, error) {
	var pbiPayload powerbiPayload

	cred, err := configs.GetEnv()
	if err != nil {
		return nil, err
	}
	auth, err := configs.GetPowerBiAuth(cred.ClientId, cred.ClientSecret, cred.TenantId)
	if err != nil {
		return nil, err

	}

	groupsVal, err := groups.GetGroups(auth)
	if err != nil {
		return nil, err

	}

	fmt.Println("len of groupsVAl", groupsVal)

	for _, grp := range groupsVal.Value {
		datasetList, err := datasets.GetDatasets(auth, grp.ID)
		if err != nil {
			return nil, err

		}
		for _, dataset := range datasetList.Value {
			dataSource, err := datasets.GetDataSources(auth, grp.ID, dataset.ID)
			if err != nil {
				return nil, err
			}
			fmt.Println("checkds:", dataSource.Value)
			if checkMatch(datas.Maps, dataSource) {
				pbiPayload.WorkspaceIds = append(pbiPayload.WorkspaceIds, grp.ID)
			}
			fmt.Println("payload checkL:", pbiPayload)
		}
	}
	if len(pbiPayload.WorkspaceIds) == 0 {
		return datas, nil
	}
	finByte, err := json.MarshalIndent(pbiPayload, "", " ")
	fmt.Println(string(finByte))

	workspaceAuth, err := configs.GetPowerBiAuthDefault(cred.ClientId, cred.ClientSecret, cred.TenantId, assertionToken)
	if err != nil {
		return nil, err
	}
	fmt.Println("Workspaceauth details:", workspaceAuth)

	wksinfo, err := workspaces.GetWorkspaceInfo(workspaceAuth.AccessToken)
	if err != nil {
		return nil, err
	}
	fmt.Println("The value in WorkspaceInfo:", wksinfo)
	var inpMap synDatasets.Mappings
	for {
		scanstat, err := workspaces.GetStatus(&wksinfo, workspaceAuth.AccessToken)
		if err != nil {
			return nil, err
		}
		fmt.Println("The value in ScanStatus:", scanstat)
		if scanstat.Status == "Succeeded" {
			break
		}
	}

	workspacesInfo, err := workspaces.GetTableDetails(&wksinfo, workspaceAuth.AccessToken)
	if err != nil {
		return nil, err
	}

	for _, dataSourceInstance := range workspacesInfo.DatasourceInstances {
		for i, data := range datas.Maps {
			sr := strings.Join([]string{data.Location, data.Name}, ";")
			dataSourceId := dataSourceInstance.DatasourceID
			dst := strings.Join([]string{strings.Split(dataSourceInstance.ConnectionDetails.Server, ",")[0], dataSourceInstance.ConnectionDetails.Database}, ",")
			//dst:=strings.Join(dataSourceInstance.)
			for _, workspaceInfo := range workspacesInfo.Workspaces {
				for _, workDataset := range workspaceInfo.Datasets {
					for _, dataUsage := range workDataset.DatasourceUsages {
						if dataSourceId == dataUsage.DatasourceInstanceID {
							for _, datasetTable := range workDataset.Tables {
								if sr == strings.Join([]string{dst, datasetTable.Name}, ";") {
									for _, report := range workspaceInfo.Reports {
										if report.DatasetID == workDataset.ID {

											inpMap.Name = report.Name
											inpMap.Type = report.ReportType

											rep, err := reports.GetReport(auth, workspaceInfo.ID, report.ID)
											if err != nil {
												return nil, err
											}
											inpMap.Location = rep.WebURL
											inpMap.Source = strings.Join([]string{data.Location, data.Name}, ";")
											inpMap.TDestination = nil
											fmt.Println("checking DataMap:", datas.Maps[i].Name, datas.Maps[i].Type)
											datas.Maps[i].TDestination = append(datas.Maps[i].TDestination, strings.Join([]string{inpMap.Location, inpMap.Name}, ";"))
											datas.Maps = append(datas.Maps, inpMap)

											fmt.Println("reportName:", report.Name, workDataset.Name, datasetTable.Name)
										}
									}
								}
							}

						}
					}

				}
			}
		}
	}

	//
	//if datas == nil {
	//	return nil, fmt.Errorf("Empty Data")
	//}
	//groups, err = groups.GetGroups(auth)
	//if err != nil {
	//	return nil, err
	//
	//}

	//fmt.Println(groupsVal.Value[0].ID)

	//for _, grp := range groups.Value {
	//	datasetList, err := datasets.GetDatasets(auth, grp.ID)
	//	if err != nil {
	//		return nil, err
	//
	//	}
	//
	//}

	//var inpMap synDatasets.Mappings
	//reportList, err := reports.GetReports(auth, groupsVal.Value[0].ID)
	//if err != nil {
	//	return nil, err
	//
	//}
	//fmt.Println("GetReportValue:", reportList.Value)

	//for _, report := range reportList.Value {
	//	if report.Name == reportname {
	//		fmt.Println("GetReportValue:", report)
	//		dataSources, err := datasets.GetDatasources(auth, groupsVal.Value[0].ID, report.DatasetID)
	//		if err != nil {
	//			return nil, err
	//
	//		}
	//
	//		inpMap.Name = reportname
	//		inpMap.Type = report.ReportType
	//		inpMap.Location = report.WebURL
	//		for i, dataMap := range datas.Maps {
	//
	//			dataM, newMap, ok := getMatch(&dataMap, dataSources, &inpMap)
	//
	//			if ok {
	//				datas.Maps[i] = *dataM
	//				datas.Maps = append(datas.Maps, *newMap)
	//			}
	//			fmt.Println("Datas:", datas.Maps, "Supposed:", dataMap)
	//			//fmt.Println()
	//		}
	//		fmt.Println("check datasMaps", datas.Maps)
	//		//for _,dataSource:=range dataSources.Value{
	//		//
	//		//}
	//
	//	}
	//}

	//for _, dataset := range datasetList.Value {
	//	if dataset.Name != "DatalineagePBI" {
	//		continue
	//	}
	//	reportList, err := reports.GetReports(auth, groups.Value[0].ID)
	//	if err != nil {
	//		return err
	//
	//	}
	//
	//	for _, report := range reportList.Value {
	//		if dataset.ID == report.DatasetID {
	//
	//			inpMap.Name = report.Name
	//			inpMap.Type = report.ReportType
	//			inpMap.Location = report.WebURL
	//			inpMap.Source = ""
	//			inpMap.Destination = nil
	//
	//			datasources, err := datasets.GetDatasources(auth, groups.Value[0].ID, dataset.ID)
	//			if err != nil {
	//				return err
	//
	//			}
	//			//dashboardList, err := dashboards.GetDashboard(auth, groups.Value[0].ID)
	//			//if err != nil {
	//			//	return err
	//			//
	//			//}
	//
	//			fmt.Printf("Check datasource: %s:%+v\n\n", report.Name, datasources.Value)
	//
	//			//for _, val := range datasources.Value {
	//			//	fmt.Println("check inpmapbf:", inpMap, datas.Maps)
	//			//	fmt.Println("check getInsight:")
	//			//	if insights, ok := getInsight(datas.Maps, val.ConnectionDetails, inpMap); ok {
	//			//		rdata[report.Name] = insights
	//			//	}
	//			//
	//			//}
	//
	//			//fmt.Printf("Check Dashobards: %+v\n", dashboardList)
	//		}
	//	}
	//	//fmt.Printf("FinalMaps: %+v", datas.Maps)
	//	//byt, err := json.MarshalIndent(datas.Maps, "", " ")
	//	//fmt.Println("\n\n finalstring:\n", string(byt))
	//
	//}
	//fmt.Println("Checking rdata")
	//for k, v := range rdata {
	//	fmt.Println(len(v))
	//	fmt.Printf("%s: %+v \n", k, v)
	//}
	return datas, nil
}

func checkMatch(dataMaps []synDatasets.Mappings, dS *datasets.PowerBiDatasource) bool {
	var found bool
	for _, dataSource := range dS.Value {
		connStr := strings.Join([]string{strings.Split(dataSource.ConnectionDetails.Server, ",")[0], dataSource.ConnectionDetails.Database}, ",")
		for _, dataMap := range dataMaps {
			fmt.Println("check1200", dataMap)
			fmt.Println("checking daaMapLo: ", dataMap.Location, connStr, dataMap.Location == connStr)
			if dataMap.Location == connStr {
				fmt.Println("found")
				found = true
				return found
			}
		}

	}
	return found
}

func getMatch(dataMap *synDatasets.Mappings, dS *datasets.PowerBiDatasource, inpMap *synDatasets.Mappings) (*synDatasets.Mappings, *synDatasets.Mappings, bool) {
	var found bool
	for _, dataSource := range dS.Value {
		connStr := strings.Join([]string{dataSource.ConnectionDetails.Server, dataSource.ConnectionDetails.Database}, ",")
		fmt.Println("check1200", dataMap)
		fmt.Println("checking daaMapLo: ", dataMap.Location, connStr, dataMap.Location == connStr)
		if dataMap.Location == connStr {
			found = true
			fmt.Println("found")
			dataMap.TDestination = append(dataMap.TDestination, strings.Join([]string{inpMap.Location, inpMap.Name}, ";"))
			fmt.Println("check datasMapsint", dataMap, "\n\n")
		}
	}
	return dataMap, inpMap, found
}
