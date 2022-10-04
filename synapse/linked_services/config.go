package linked_services

type LinkedService struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		Annotations    []interface{} `json:"annotations"`
		Type           string        `json:"type"`
		TypeProperties struct {
			Url                 string `json:"url"`
			ConnectionString    string `json:"connectionString"`
			EncryptedCredential string `json:"encryptedCredential"`
		} `json:"typeProperties"`
		ConnectVia struct {
			ReferenceName string `json:"referenceName"`
			Type          string `json:"type"`
		} `json:"connectVia"`
	} `json:"properties"`
	Etag string `json:"etag"`
}
