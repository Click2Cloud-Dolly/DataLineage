package configs

import (
	"fmt"
	"os"
)

func GetEnv() (Cred, error) {
	var c Cred
	c.ClientId = os.Getenv("AZURE_CLIENT_ID")
	c.ClientSecret = os.Getenv("AZURE_CLIENT_SECRET")
	c.TenantId = os.Getenv("AZURE_TENANT_ID")
	c.SubscriptionId = os.Getenv("AZURE_SUBSCRIPTION_ID")

	if c.ClientId == "" || c.ClientSecret == "" || c.TenantId == "" || c.SubscriptionId == "" {
		return c, fmt.Errorf("Please provide proper AZURE_CLIENT_ID, AZURE_CLIENT_SECRET, AZURE_SUBSCRIPTION_ID and AZURE_TENANT_ID")
	}

	return c, nil
}
