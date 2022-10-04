package workspaces

type WorkspaceList struct {
	Value []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		Location   string `json:"location"`
		Properties struct {
			DefaultDataLakeStorage struct {
				AccountURL string `json:"accountUrl"`
				Filesystem string `json:"filesystem"`
			} `json:"defaultDataLakeStorage"`
			ConnectivityEndpoints struct {
				Dev string `json:"dev"`
				Sql string `json:"sql"`
			} `json:"connectivityEndpoints"`
			SqlAdministratorLogin    string `json:"sqlAdministratorLogin"`
			ManagedResourceGroupName string `json:"managedResourceGroupName"`
			ProvisioningState        string `json:"provisioningState"`
			WorkspaceUID             string `json:"workspaceUID"`
		} `json:"properties"`
		Tags struct {
			Key string `json:"key"`
		} `json:"tags"`
	} `json:"value"`
}

type Workspace struct {
	Properties struct {
		Settings struct {
			AzureADOnlyAuthentication struct {
				Value bool   `json:"value"`
				State string `json:"state"`
			} `json:"azureADOnlyAuthentication"`
		} `json:"settings"`
		DefaultDataLakeStorage struct {
			ResourceID                   string `json:"resourceId"`
			CreateManagedPrivateEndpoint bool   `json:"createManagedPrivateEndpoint"`
			AccountURL                   string `json:"accountUrl"`
			Filesystem                   string `json:"filesystem"`
		} `json:"defaultDataLakeStorage"`
		Encryption struct {
			DoubleEncryptionEnabled bool `json:"doubleEncryptionEnabled"`
		} `json:"encryption"`
		ProvisioningState     string `json:"provisioningState"`
		ConnectivityEndpoints struct {
			Web         string `json:"web"`
			Dev         string `json:"dev"`
			SqlOnDemand string `json:"sqlOnDemand"`
			Sql         string `json:"sql"`
		} `json:"connectivityEndpoints"`
		ManagedResourceGroupName   string        `json:"managedResourceGroupName"`
		SqlAdministratorLogin      string        `json:"sqlAdministratorLogin"`
		PrivateEndpointConnections []interface{} `json:"privateEndpointConnections"`
		WorkspaceUID               string        `json:"workspaceUID"`
		ExtraProperties            struct {
			WorkspaceType  string `json:"WorkspaceType"`
			IsScopeEnabled bool   `json:"IsScopeEnabled"`
		} `json:"extraProperties"`
		PublicNetworkAccess         string `json:"publicNetworkAccess"`
		CspWorkspaceAdminProperties struct {
			InitialWorkspaceAdminObjectID string `json:"initialWorkspaceAdminObjectId"`
		} `json:"cspWorkspaceAdminProperties"`
		AzureADOnlyAuthentication   bool `json:"azureADOnlyAuthentication"`
		TrustedServiceBypassEnabled bool `json:"trustedServiceBypassEnabled"`
	} `json:"properties"`
	Type     string `json:"type"`
	ID       string `json:"id"`
	Location string `json:"location"`
	Name     string `json:"name"`
	Identity struct {
		Type        string `json:"type"`
		TenantID    string `json:"tenantId"`
		PrincipalID string `json:"principalId"`
	} `json:"identity"`
	Tags struct {
	} `json:"tags"`
}
