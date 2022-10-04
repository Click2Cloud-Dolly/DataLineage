package configs

type Cred struct {
	ClientId       string `json:"clientId"`
	ClientSecret   string `json:"clientSecret"`
	SubscriptionId string `json:"subscriptionId"`
	TenantId       string `json:"tenantId"`
	AssertionToken string `json:"assertionToken"`
}

type Auth struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	ExtExpiresIn string `json:"ext_expires_in"`
	ExpiresOn    string `json:"expires_on"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	AccessToken  string `json:"access_token"`
}
type PowerBIAuth struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
}

// const PowerBI_Endpoint = "https://api.powerbi.com/v1.0/myorg/groups"
const PowerBIResource = "https://analysis.windows.net/powerbi/api"
const PowerBIResourceDefault = "https://analysis.windows.net/powerbi/api/.default"
const AzureSynapseResource = "https://dev.azuresynapse.net/"
const AzureStorageResource = "https://storage.azure.com/"
const AzureToken = "https://login.microsoftonline.com/"
const AzureCore = "https://management.core.windows.net/"
const AzureDataFactory = "https://management.azure.com/"
