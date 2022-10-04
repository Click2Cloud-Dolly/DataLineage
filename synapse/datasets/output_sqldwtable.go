package datasets

//func (dw *AzureSQLDWTableDetails) GetDataset(body []byte, datasetGraph *[]Graph1, accessToken config.Auth, workspaceEndpoint string) error {
//	err := json.Unmarshal(body, &dw)
//	if err != nil {
//		return err
//	}
//
//	dwtlinked, err := linked_services.GetLinkedServices(accessToken, workspaceEndpoint, dw.Properties.LinkedServiceName.ReferenceName)
//	if err != nil {
//		fmt.Println("error:", err)
//		return err
//	}
//	fmt.Printf("Check dwtlinked %+v\n", dwtlinked)
//	testGraph := *datasetGraph
//	if len(*datasetGraph) > 1 {
//		testGraph[len(testGraph)-1].Destination = strings.Join([]string{dwtlinked.Properties.TypeProperties.Url, dw.Properties.TypeProperties.Table}, "/")
//	} else {
//		testGraph[0].Destination = strings.Join([]string{dwtlinked.Properties.TypeProperties.Url, dw.Properties.TypeProperties.Table}, "/")
//	}
//	*datasetGraph = append(*datasetGraph, Graph1{
//		Name:        dw.Properties.TypeProperties.Table,
//		Type:        dw.Properties.Type,
//		Source:      dw.Name,
//		Destination: "",
//		Location:    strings.Join([]string{dwtlinked.Properties.TypeProperties.Url, dw.Properties.TypeProperties.Table}, "/"),
//	})
//	return nil
//
//}
