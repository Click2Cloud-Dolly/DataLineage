package datafactories

import "time"

type DataFactoryList struct {
	Value []struct {
		Name       string `json:"name"`
		ID         string `json:"id"`
		Type       string `json:"type"`
		Properties struct {
			ProvisioningState   string    `json:"provisioningState"`
			CreateTime          time.Time `json:"createTime"`
			Version             string    `json:"version"`
			PublicNetworkAccess string    `json:"publicNetworkAccess"`
		} `json:"properties"`
		ETag     string `json:"eTag"`
		Location string `json:"location"`
		Identity struct {
			Type        string `json:"type"`
			PrincipalID string `json:"principalId"`
			TenantID    string `json:"tenantId"`
		} `json:"identity"`
		Tags struct {
		} `json:"tags"`
	} `json:"value"`
}

type DataFactory struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	Type       string `json:"type"`
	Properties struct {
		ProvisioningState string    `json:"provisioningState"`
		CreateTime        time.Time `json:"createTime"`
		Version           string    `json:"version"`
		FactoryStatistics struct {
			TotalResourceCount             int `json:"totalResourceCount"`
			MaxAllowedResourceCount        int `json:"maxAllowedResourceCount"`
			FactorySizeInGbUnits           int `json:"factorySizeInGbUnits"`
			MaxAllowedFactorySizeInGbUnits int `json:"maxAllowedFactorySizeInGbUnits"`
		} `json:"factoryStatistics"`
		PublicNetworkAccess string `json:"publicNetworkAccess"`
	} `json:"properties"`
	ETag     string `json:"eTag"`
	Location string `json:"location"`
	Identity struct {
		Type        string `json:"type"`
		PrincipalID string `json:"principalId"`
		TenantID    string `json:"tenantId"`
	} `json:"identity"`
	Tags struct {
	} `json:"tags"`
}
