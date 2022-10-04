package datafactory

import (
	"azureDataLineage/configs"
	"azureDataLineage/datafactory/datafactories"
	"azureDataLineage/datafactory/datasets"
	"azureDataLineage/datafactory/pipelines"
	"fmt"
	"strings"
)

type Data struct {
	Dataset datasets.Dataset
	Maps    []datasets.DFMappings
	dataops datasets.DatasetInterface
}

func DataFactory() (*Data, error) {

	var d = new(Data)
	cred, err := configs.GetEnv()
	if err != nil {
		return nil, err
	}
	auth, err := cred.GetTokenDataFactory()
	if err != nil {
		fmt.Println(err)
	}

	var datafactorylist *datafactories.DataFactoryList
	for i := 0; i < 3; i++ {
		datafactorylist, err = datafactories.ListDataFactory(auth, cred)
		if err != nil {
			fmt.Println(err)
		}
		if datafactorylist != nil {
			break
		}
	}

	var ids string
	for _, datafact := range datafactorylist.Value {
		ids = datafact.ID
		fmt.Println("\nID:", ids)
		resourceGroups := strings.Split(ids, "/")[4]
		datafactoryName := strings.Split(ids, "/")[8]

		//GetDataFactory
		getdatafactory, err := datafactories.GetDataFactory(auth, cred, resourceGroups, datafactoryName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("\nDataFactory:", getdatafactory)

		//ListPipelines
		listpipelines, err := pipelines.ListPipelines(auth, cred, resourceGroups, datafactoryName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("\nPipelineList:", listpipelines)

		for _, pipelinedet := range listpipelines.Value {
			pipeline, err := pipelines.GetPipeline(auth, cred, resourceGroups, datafactoryName, pipelinedet.Name)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("\nPipelines:", pipeline)

			var dfInMap, dfOutMap []datasets.DFMappings
			var dfFillMap, dfNullMap datasets.DFMappings
			for _, activity := range pipeline.Properties.Activities {
				if activity.TypeProperties.Activities != nil {
					for _, nestedActivity := range activity.TypeProperties.Activities {
						fmt.Println("\nDataset_Activities:", activity)

						for i, input := range activity.Inputs {
							dfFillMap.Name = input.ReferenceName
							dfFillMap.Type = input.Type
							dfFillMap.Location = datafactoryName
							dfFillMap.TDestination = append(dfFillMap.TDestination, strings.Join([]string{datafactoryName, nestedActivity.Outputs[i].ReferenceName}, ";"))
							dfInMap = append(dfInMap, dfFillMap)
							err = d.Dataset.GetDatasets(auth, cred, resourceGroups, datafactoryName, input.ReferenceName, &dfInMap)
							if err != nil {
								fmt.Println(err)
							}
							d.Maps = append(d.Maps, dfInMap...)
							dfFillMap = dfNullMap
							fmt.Printf("\ninpmaps: %v", dfInMap)
						}

						for j, output := range activity.Outputs {
							dfFillMap.Name = output.ReferenceName
							dfFillMap.Type = output.Type
							dfFillMap.Location = datafactoryName
							dfFillMap.Source = strings.Join([]string{datafactoryName, nestedActivity.Inputs[j].ReferenceName}, ";")
							dfOutMap = append(dfOutMap, dfFillMap)
							dfFillMap = dfNullMap
							err = d.Dataset.GetDatasets(auth, cred, resourceGroups, datafactoryName, output.ReferenceName, &dfOutMap)
							if err != nil {
								fmt.Println(err)
							}
						}
					}
				}

				for _, activity := range pipeline.Properties.Activities {
					for i, input := range activity.Inputs {
						dfFillMap.Name = input.ReferenceName
						dfFillMap.Type = input.Type
						dfFillMap.Location = datafactoryName
						dfFillMap.TDestination = append(dfFillMap.TDestination, strings.Join([]string{datafactoryName, activity.Outputs[i].ReferenceName}, ";"))
						dfInMap = append(dfInMap, dfFillMap)
						err = d.Dataset.GetDatasets(auth, cred, resourceGroups, datafactoryName, input.ReferenceName, &dfInMap)
						if err != nil {
							fmt.Println(err)
						}
						d.Maps = append(d.Maps, dfInMap...)
						dfFillMap = dfNullMap
						fmt.Printf("\ninpmaps: %v", dfInMap)
					}
					for j, output := range activity.Outputs {
						dfFillMap.Name = output.ReferenceName
						dfFillMap.Type = output.Type
						dfFillMap.Location = datafactoryName
						dfFillMap.Source = strings.Join([]string{datafactoryName, activity.Inputs[j].ReferenceName}, ";")
						dfOutMap = append(dfOutMap, dfFillMap)

						dfFillMap = dfNullMap
						err = d.Dataset.GetDatasets(auth, cred, resourceGroups, datafactoryName, output.ReferenceName, &dfOutMap)
						if err != nil {
							fmt.Println(err)
						}
						d.Maps = append(d.Maps, dfOutMap...)

					}
				}

			}
		}
	}
	return d, nil
}
