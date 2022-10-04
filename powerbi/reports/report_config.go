package reports

type PowerBiReports struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		ID                 string        `json:"id"`
		ReportType         string        `json:"reportType"`
		Name               string        `json:"name"`
		WebURL             string        `json:"webUrl"`
		EmbedURL           string        `json:"embedUrl"`
		IsFromPbix         bool          `json:"isFromPbix"`
		IsOwnedByMe        bool          `json:"isOwnedByMe"`
		DatasetID          string        `json:"datasetId"`
		DatasetWorkspaceID string        `json:"datasetWorkspaceId"`
		Users              []interface{} `json:"users"`
		Subscriptions      []interface{} `json:"subscriptions"`
	} `json:"value"`
}
type PowerBIReport struct {
	OdataContext       string        `json:"@odata.context"`
	ID                 string        `json:"id"`
	ReportType         string        `json:"reportType"`
	Name               string        `json:"name"`
	WebURL             string        `json:"webUrl"`
	EmbedURL           string        `json:"embedUrl"`
	IsFromPbix         bool          `json:"isFromPbix"`
	IsOwnedByMe        bool          `json:"isOwnedByMe"`
	DatasetID          string        `json:"datasetId"`
	DatasetWorkspaceID string        `json:"datasetWorkspaceId"`
	Users              []interface{} `json:"users"`
	Subscriptions      []interface{} `json:"subscriptions"`
}
