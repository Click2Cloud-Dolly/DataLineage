package synapse

import (
	"azureDataLineage/configs"
	"azureDataLineage/synapse/datasets"
	"azureDataLineage/synapse/pipelines"
	"azureDataLineage/synapse/workspaces"
	"fmt"
	"strings"
)

type Data struct {
	Dataset datasets.Dataset
	Maps    []datasets.Mappings
	dataops datasets.DatasetInterface
}

func DataLineageSynapse() (*Data, error) {
	var d = new(Data)
	cred1, err := configs.GetEnv()
	if err != nil {
		return nil, err
	}
	auth, err := cred1.GetToken()
	if err != nil {
		return nil, err
	}
	fmt.Println("\nGetToken:", auth.AccessToken)
	var wlist *workspaces.WorkspaceList

	//Listing Synapse Workspaces
	for i := 0; i < 3; i++ {
		wlist, err = workspaces.ListWorkspaces(auth, cred1)
		if err != nil {
			return nil, err
		}
		if wlist != nil {
			break
		}
	}
	var resourceGroup, workspacename string
	fmt.Println("\nwlist:", wlist)
	for _, val := range wlist.Value {
		//if val.Name == "datalineageworkspace" {
		resourceGroup = strings.Split(val.ID, "/")[4]
		workspacename = val.Name
		//}

		//GetWorkspace
		wdet, err := workspaces.GetWorkspace(auth, cred1, resourceGroup, workspacename)
		if err != nil {
			return nil, err
		}

		fmt.Println("\nwdet:", wdet)
		auth1, err := cred1.GetToken1()
		if err != nil {
			return nil, err
		}
		fmt.Println("auth1", auth1.AccessToken)
		fmt.Println("\nwdet:", wdet.Properties.ConnectivityEndpoints.Dev)

		//ListPipelines
		pipelinesList, err := pipelines.ListPipelines(auth1, wdet.Properties.ConnectivityEndpoints.Dev)
		if err != nil {
			return nil, err
		}
		fmt.Println("\nList pipeline:", pipelinesList, "workspace:", wdet.Properties.ConnectivityEndpoints.Dev)

		//GetPipeline
		for _, pipelinesdet := range pipelinesList.Value {
			fmt.Println("Pipeline:", pipelinesdet.Name)
			//if pipelinesdet.Name != "AllCopyPipeline_2m5" {
			//	continue
			//}

			pipeline, err := pipelines.GetPipeline(auth1, wdet.Properties.ConnectivityEndpoints.Dev, pipelinesdet.Name)
			if err != nil {
				return nil, err
			}
			fmt.Println("Yo:", pipeline.Properties.Activities)
			var inpMap, outMap []datasets.Mappings
			var /*inpMap*/ fillMap, nullMap datasets.Mappings
			for _, activity := range pipeline.Properties.Activities {
				if activity.TypeProperties.Activities != nil {
					fmt.Println("Ifpart")
					for _, nestedActivity := range activity.TypeProperties.Activities {
						for i, input := range nestedActivity.Inputs {
							fmt.Println("check activityInputs", input)

							fillMap.Name = input.ReferenceName
							fillMap.Type = input.Type
							fillMap.Location = wdet.Properties.ConnectivityEndpoints.Dev
							fillMap.TDestination = append(fillMap.TDestination, strings.Join([]string{wdet.Properties.ConnectivityEndpoints.Dev, nestedActivity.Outputs[i].ReferenceName}, ";"))
							inpMap = append(inpMap, fillMap)
							err = d.Dataset.GetDatasets(auth1, wdet.Properties.ConnectivityEndpoints.Dev, input.ReferenceName, &inpMap)
							if err != nil {
								return nil, err
							}
							d.Maps = append(d.Maps, inpMap...)
							fillMap = nullMap
							fmt.Printf("\ninpmaps: %v", inpMap)

						}
						for j, output := range nestedActivity.Outputs {
							fillMap.Name = output.ReferenceName
							fillMap.Type = output.Type
							fillMap.Location = wdet.Properties.ConnectivityEndpoints.Dev
							fillMap.Source = strings.Join([]string{wdet.Properties.ConnectivityEndpoints.Dev, nestedActivity.Inputs[j].ReferenceName}, ";")
							outMap = append(outMap, fillMap)
							fillMap = nullMap
							err = d.Dataset.GetDatasets(auth1, wdet.Properties.ConnectivityEndpoints.Dev, output.ReferenceName, &outMap)
							if err != nil {
								return nil, err
							}
						}
					}
				}

				for _, activity := range pipeline.Properties.Activities {
					for i, input := range activity.Inputs {
						fillMap.Name = input.ReferenceName
						fillMap.Type = input.Type
						fillMap.Location = wdet.Properties.ConnectivityEndpoints.Dev
						fillMap.TDestination = append(fillMap.TDestination, strings.Join([]string{wdet.Properties.ConnectivityEndpoints.Dev, activity.Outputs[i].ReferenceName}, ";"))
						inpMap = append(inpMap, fillMap)
						err = d.Dataset.GetDatasets(auth1, wdet.Properties.ConnectivityEndpoints.Dev, input.ReferenceName, &inpMap)
						if err != nil {
							return nil, err
						}
						d.Maps = append(d.Maps, inpMap...)
						fillMap = nullMap

					}
					for j, output := range activity.Outputs {
						fillMap.Name = output.ReferenceName
						fillMap.Type = output.Type
						fillMap.Location = wdet.Properties.ConnectivityEndpoints.Dev
						fillMap.Source = strings.Join([]string{wdet.Properties.ConnectivityEndpoints.Dev, activity.Inputs[j].ReferenceName}, ";")
						outMap = append(outMap, fillMap)

						fillMap = nullMap
						err = d.Dataset.GetDatasets(auth1, wdet.Properties.ConnectivityEndpoints.Dev, output.ReferenceName, &outMap)
						if err != nil {
							return nil, err
						}
						d.Maps = append(d.Maps, outMap...)
					}
				}

			}

		}
	}
	return d, nil
}
