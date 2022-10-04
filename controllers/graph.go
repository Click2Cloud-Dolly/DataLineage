package controllers

import (
	"azureDataLineage/configs"
	"azureDataLineage/datafactory"
	"azureDataLineage/powerbi"
	"azureDataLineage/synapse"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type TokenAssert struct {
	AssertionToken string `json:"assertionToken"`
}

type GraphOutput struct {
	Message       string `json:"message"`
	NodeList      string `json:"nodeList"`
	NodeDirection string `json:"nodeDirection"`
}

func Graph(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=604800, immutable")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	var authToken TokenAssert
	var newAuth configs.Cred
	var output GraphOutput
	if err := c.BindJSON(&authToken); err != nil {
		return
	}

	fmt.Println("check Assertion0071", authToken.AssertionToken)
	newAuth.ClientId = os.Getenv("AZURE_CLIENT_ID")
	newAuth.ClientSecret = os.Getenv("AZURE_CLIENT_SECRET")
	newAuth.TenantId = os.Getenv("AZURE_TENANT_ID")
	newAuth.SubscriptionId = os.Getenv("AZURE_SUBSCRIPTION_ID")
	newAuth.AssertionToken = authToken.AssertionToken
	auth, err := newAuth.GetToken()
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, err)

	}
	fmt.Println(auth)

	data, err := synapse.DataLineageSynapse()
	if err != nil {
		fmt.Println(err)
		return
	}
	fol1, err := json.MarshalIndent(data.Maps, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("fol::11:", string(fol1))

	dataFactory, err := datafactory.DataFactory()
	if err != nil {
		fmt.Println(err)
		return
	}
	datafact, err := json.MarshalIndent(dataFactory.Maps, "", " ")
	fmt.Println("DataFactory_Pipelines:", string(datafact))

	dataList, err := powerbi.DataLineagePowerBi(data, newAuth.AssertionToken)
	if err != nil {
		fmt.Println(err)
	}
	nodesJson, nodeDirs, err := NodesLineage(dataList.Maps)
	if err != nil {
		fmt.Println(err)
	}

	output.Message = "Client Connected"
	output.NodeList = string(nodesJson)
	output.NodeDirection = string(nodeDirs)
	c.IndentedJSON(http.StatusCreated, output)
}
