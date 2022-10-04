package groups

type Groups struct {
	OdataContext string `json:"@odata.context"`
	OdataCount   int    `json:"@odata.count"`
	Value        []struct {
		ID                    string `json:"id"`
		IsReadOnly            bool   `json:"isReadOnly"`
		IsOnDedicatedCapacity bool   `json:"isOnDedicatedCapacity"`
		Type                  string `json:"type"`
		Name                  string `json:"name"`
	} `json:"value"`
}
