package datasets

import "time"

type PowerBiDataset struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		ID                               string        `json:"id"`
		Name                             string        `json:"name"`
		WebURL                           string        `json:"webUrl"`
		AddRowsAPIEnabled                bool          `json:"addRowsAPIEnabled"`
		ConfiguredBy                     string        `json:"configuredBy"`
		IsRefreshable                    bool          `json:"isRefreshable"`
		IsEffectiveIdentityRequired      bool          `json:"isEffectiveIdentityRequired"`
		IsEffectiveIdentityRolesRequired bool          `json:"isEffectiveIdentityRolesRequired"`
		IsOnPremGatewayRequired          bool          `json:"isOnPremGatewayRequired"`
		TargetStorageMode                string        `json:"targetStorageMode"`
		CreatedDate                      time.Time     `json:"createdDate"`
		CreateReportEmbedURL             string        `json:"createReportEmbedURL"`
		QnaEmbedURL                      string        `json:"qnaEmbedURL"`
		UpstreamDatasets                 []interface{} `json:"upstreamDatasets"`
		Users                            []interface{} `json:"users"`
	} `json:"value"`
}

type PowerBiDatasource struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		DatasourceType    string `json:"datasourceType"`
		ConnectionDetails struct {
			Server   string `json:"server"`
			Database string `json:"database"`
		} `json:"connectionDetails"`
		DatasourceID string `json:"datasourceId"`
		GatewayID    string `json:"gatewayId"`
	} `json:"value"`
}

type ConnectionDetails struct {
	Server   string `json:"server"`
	Database string `json:"database"`
}
