package sql_pools

import "time"

type SqlPool struct {
	Properties struct {
		Status             string    `json:"status"`
		MaxSizeBytes       int64     `json:"maxSizeBytes"`
		Collation          string    `json:"collation"`
		CreationDate       time.Time `json:"creationDate"`
		StorageAccountType string    `json:"storageAccountType"`
		ProvisioningState  string    `json:"provisioningState"`
	} `json:"properties"`
	Sku struct {
		Name     string `json:"name"`
		Capacity int    `json:"capacity"`
	} `json:"sku"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Location string `json:"location"`
	Tags     struct {
	} `json:"tags"`
}
