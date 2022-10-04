package controllers

import (
	"azureDataLineage/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

type LoginOutput struct {
	Message string `json:"message"`
	URL     string `json:"url"`
}

func Login(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=604800, immutable")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	var newAuth configs.Cred
	var output LoginOutput

	if err := c.BindJSON(&newAuth); err != nil {
		return
	}

	os.Setenv("AZURE_CLIENT_ID", newAuth.ClientId)
	os.Setenv("AZURE_CLIENT_SECRET", newAuth.ClientSecret)
	os.Setenv("AZURE_TENANT_ID", newAuth.TenantId)
	os.Setenv("AZURE_SUBSCRIPTION_ID", newAuth.SubscriptionId)
	fmt.Println("check bounded newAuth", newAuth)

	// Send SSO URL
	output.Message = "URL Found"
	output.URL = strings.Join([]string{"https://login.microsoftonline.com/", newAuth.TenantId, "/oauth2/v2.0/authorize?client_id=", newAuth.ClientId, "&response_type=id_token&redirect_uri=http%3A%2F%2Flocalhost%3A4300%2Fdashboard&scope=openid profile email&response_mode=fragment&state=12345&nonce=678910&prompt=login"}, "")

	//c.IndentedJSON(http.StatusCreated, output)

}
