package dashboards

type PowerBiDashboards struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		ID            string        `json:"id"`
		DisplayName   string        `json:"displayName"`
		IsReadOnly    bool          `json:"isReadOnly"`
		WebURL        string        `json:"webUrl"`
		EmbedURL      string        `json:"embedUrl"`
		Users         []interface{} `json:"users"`
		Subscriptions []interface{} `json:"subscriptions"`
	} `json:"value"`
}
