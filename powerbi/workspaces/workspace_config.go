package workspaces

import "time"

type WorkspaceDetails struct {
	Workspaces []struct {
		ID                    string `json:"id"`
		Name                  string `json:"name"`
		Type                  string `json:"type"`
		State                 string `json:"state"`
		IsOnDedicatedCapacity bool   `json:"isOnDedicatedCapacity"`
		Reports               []struct {
			ReportType       string `json:"reportType"`
			ID               string `json:"id"`
			Name             string `json:"name"`
			DatasetID        string `json:"datasetId"`
			CreatedDateTime  string `json:"createdDateTime"`
			ModifiedDateTime string `json:"modifiedDateTime"`
			ModifiedBy       string `json:"modifiedBy"`
			CreatedBy        string `json:"createdBy"`
			ModifiedByID     string `json:"modifiedById"`
			CreatedByID      string `json:"createdById"`
			Users            []struct {
				ReportUserAccessRight string `json:"reportUserAccessRight"`
				EmailAddress          string `json:"emailAddress"`
				DisplayName           string `json:"displayName"`
				Identifier            string `json:"identifier"`
				GraphID               string `json:"graphId"`
				PrincipalType         string `json:"principalType"`
				UserType              string `json:"userType"`
			} `json:"users"`
		} `json:"reports"`
		Dashboards []struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
			IsReadOnly  bool   `json:"isReadOnly"`
			Tiles       []struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				SubTitle string `json:"subTitle"`
			} `json:"tiles"`
			Users []struct {
				DashboardUserAccessRight string `json:"dashboardUserAccessRight"`
				EmailAddress             string `json:"emailAddress"`
				DisplayName              string `json:"displayName"`
				Identifier               string `json:"identifier"`
				GraphID                  string `json:"graphId"`
				PrincipalType            string `json:"principalType"`
				UserType                 string `json:"userType"`
			} `json:"users"`
		} `json:"dashboards"`
		Datasets []struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Tables []struct {
				Name    string `json:"name"`
				Columns []struct {
					Name       string `json:"name"`
					DataType   string `json:"dataType"`
					IsHidden   bool   `json:"isHidden"`
					ColumnType string `json:"columnType"`
				} `json:"columns"`
				Measures []interface{} `json:"measures"`
				IsHidden bool          `json:"isHidden"`
				Source   []struct {
					Expression string `json:"expression"`
				} `json:"source"`
			} `json:"tables"`
			ConfiguredBy                     string `json:"configuredBy"`
			ConfiguredByID                   string `json:"configuredById"`
			IsEffectiveIdentityRequired      bool   `json:"isEffectiveIdentityRequired"`
			IsEffectiveIdentityRolesRequired bool   `json:"isEffectiveIdentityRolesRequired"`
			TargetStorageMode                string `json:"targetStorageMode"`
			CreatedDate                      string `json:"createdDate"`
			ContentProviderType              string `json:"contentProviderType"`
			DatasourceUsages                 []struct {
				DatasourceInstanceID string `json:"datasourceInstanceId"`
			} `json:"datasourceUsages,omitempty"`
			Users []struct {
				DatasetUserAccessRight string `json:"datasetUserAccessRight"`
				EmailAddress           string `json:"emailAddress"`
				DisplayName            string `json:"displayName"`
				Identifier             string `json:"identifier"`
				GraphID                string `json:"graphId"`
				PrincipalType          string `json:"principalType"`
				UserType               string `json:"userType"`
			} `json:"users"`
		} `json:"datasets"`
		Dataflows []interface{} `json:"dataflows"`
		Datamarts []interface{} `json:"datamarts"`
		Users     []struct {
			GroupUserAccessRight string `json:"groupUserAccessRight"`
			EmailAddress         string `json:"emailAddress"`
			DisplayName          string `json:"displayName"`
			Identifier           string `json:"identifier"`
			GraphID              string `json:"graphId"`
			PrincipalType        string `json:"principalType"`
			UserType             string `json:"userType"`
		} `json:"users"`
	} `json:"workspaces"`
	DatasourceInstances []struct {
		DatasourceType    string `json:"datasourceType"`
		ConnectionDetails struct {
			Server   string `json:"server"`
			Database string `json:"database"`
		} `json:"connectionDetails"`
		DatasourceID string `json:"datasourceId"`
		GatewayID    string `json:"gatewayId"`
	} `json:"datasourceInstances"`
}

type WorkspaceInfo struct {
	ID              string    `json:"id"`
	CreatedDateTime time.Time `json:"createdDateTime"`
	Status          string    `json:"status"`
}

type ScanStatus struct {
	ID              string `json:"id"`
	CreatedDateTime string `json:"createdDateTime"`
	Status          string `json:"status"`
}

type ScanResult struct {
	Workspaces []struct {
		ID                    string `json:"id"`
		Name                  string `json:"name"`
		Type                  string `json:"type"`
		State                 string `json:"state"`
		IsOnDedicatedCapacity bool   `json:"isOnDedicatedCapacity"`
		Reports               []struct {
			ReportType       string `json:"reportType"`
			ID               string `json:"id"`
			Name             string `json:"name"`
			DatasetID        string `json:"datasetId"`
			CreatedDateTime  string `json:"createdDateTime"`
			ModifiedDateTime string `json:"modifiedDateTime"`
			ModifiedBy       string `json:"modifiedBy"`
			CreatedBy        string `json:"createdBy"`
			ModifiedByID     string `json:"modifiedById"`
			CreatedByID      string `json:"createdById"`
			Users            []struct {
				ReportUserAccessRight string `json:"reportUserAccessRight"`
				EmailAddress          string `json:"emailAddress"`
				DisplayName           string `json:"displayName"`
				Identifier            string `json:"identifier"`
				GraphID               string `json:"graphId"`
				PrincipalType         string `json:"principalType"`
				UserType              string `json:"userType"`
			} `json:"users"`
		} `json:"reports"`
		Dashboards []struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
			IsReadOnly  bool   `json:"isReadOnly"`
			Tiles       []struct {
				ID       string `json:"id"`
				Title    string `json:"title"`
				SubTitle string `json:"subTitle"`
			} `json:"tiles"`
			Users []struct {
				DashboardUserAccessRight string `json:"dashboardUserAccessRight"`
				EmailAddress             string `json:"emailAddress"`
				DisplayName              string `json:"displayName"`
				Identifier               string `json:"identifier"`
				GraphID                  string `json:"graphId"`
				PrincipalType            string `json:"principalType"`
				UserType                 string `json:"userType"`
			} `json:"users"`
		} `json:"dashboards"`
		Datasets []struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Tables []struct {
				Name    string `json:"name"`
				Columns []struct {
					Name       string `json:"name"`
					DataType   string `json:"dataType"`
					IsHidden   bool   `json:"isHidden"`
					ColumnType string `json:"columnType"`
				} `json:"columns"`
				Measures []interface{} `json:"measures"`
				IsHidden bool          `json:"isHidden"`
				Source   []struct {
					Expression string `json:"expression"`
				} `json:"source"`
			} `json:"tables"`
			ConfiguredBy                     string `json:"configuredBy"`
			ConfiguredByID                   string `json:"configuredById"`
			IsEffectiveIdentityRequired      bool   `json:"isEffectiveIdentityRequired"`
			IsEffectiveIdentityRolesRequired bool   `json:"isEffectiveIdentityRolesRequired"`
			TargetStorageMode                string `json:"targetStorageMode"`
			CreatedDate                      string `json:"createdDate"`
			ContentProviderType              string `json:"contentProviderType"`
			DatasourceUsages                 []struct {
				DatasourceInstanceID string `json:"datasourceInstanceId"`
			} `json:"datasourceUsages,omitempty"`
			Users []struct {
				DatasetUserAccessRight string `json:"datasetUserAccessRight"`
				EmailAddress           string `json:"emailAddress"`
				DisplayName            string `json:"displayName"`
				Identifier             string `json:"identifier"`
				GraphID                string `json:"graphId"`
				PrincipalType          string `json:"principalType"`
				UserType               string `json:"userType"`
			} `json:"users"`
		} `json:"datasets"`
		Dataflows []interface{} `json:"dataflows"`
		Datamarts []interface{} `json:"datamarts"`
		Users     []struct {
			GroupUserAccessRight string `json:"groupUserAccessRight"`
			EmailAddress         string `json:"emailAddress"`
			DisplayName          string `json:"displayName"`
			Identifier           string `json:"identifier"`
			GraphID              string `json:"graphId"`
			PrincipalType        string `json:"principalType"`
			UserType             string `json:"userType"`
		} `json:"users"`
	} `json:"workspaces"`
	DatasourceInstances []struct {
		DatasourceType    string `json:"datasourceType"`
		ConnectionDetails struct {
			Server   string `json:"server"`
			Database string `json:"database"`
		} `json:"connectionDetails"`
		DatasourceID string `json:"datasourceId"`
		GatewayID    string `json:"gatewayId"`
	} `json:"datasourceInstances"`
}
