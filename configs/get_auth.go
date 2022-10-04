package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetPowerBiAuth(clientID, clientSecret, tenantID string) (Auth, error) {
	var auth Auth

	url := strings.Join([]string{"https://login.microsoftonline.com", tenantID, "oauth2/token"}, "/")
	method := "GET"

	data := strings.Join([]string{"grant_type=client_credentials&client_id=", clientID, "&client_secret=", clientSecret, "&resource=", PowerBIResource}, "")
	payload := strings.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return auth, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return auth, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return auth, err
	}
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return auth, err
	}
	return auth, nil
}

func GetPowerBiAuthDefault(clientID, clientSecret, tenantID, assertionToken string) (PowerBIAuth, error) {
	var auth PowerBIAuth
	url := strings.Join([]string{"https://login.microsoftonline.com", tenantID, "oauth2/v2.0/token"}, "/")
	//url := "https://login.microsoftonline.com/fa0fd8d2-d414-482c-a2a3-d86969b5c5fc/oauth2/v2.0/token"
	method := "POST"

	data := strings.Join([]string{"grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer&client_id=", clientID, "&client_secret=", clientSecret, "&requested_token_use=on_behalf_of&assertion=", assertionToken, "&scope=", PowerBIResourceDefault}, "")
	payload := strings.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return auth, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return auth, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return auth, err
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &auth)
	if err != nil {
		fmt.Println(err)
		return auth, err
	}
	return auth, nil
}

func (c *Cred) GetToken() (Auth, error) {
	var auth Auth
	url := strings.Join([]string{AzureToken, c.TenantId, "/oauth2/token"}, "")
	method := "GET"
	data := strings.Join([]string{"grant_type=client_credentials&client_id=", c.ClientId, "&client_secret=", c.ClientSecret, "&resource=", AzureCore}, "")
	payload := strings.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return auth, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return auth, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return auth, err
	}
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return auth, err
	}
	return auth, nil
}
func (c *Cred) GetToken1() (Auth, error) {
	var auth Auth
	url := strings.Join([]string{AzureToken, c.TenantId, "/oauth2/token"}, "")
	method := "GET"
	data := strings.Join([]string{"grant_type=client_credentials&client_id=", c.ClientId, "&client_secret=", c.ClientSecret, "&resource=", AzureSynapseResource}, "")
	payload := strings.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return auth, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return auth, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return auth, err
	}
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return auth, err
	}
	return auth, nil
}

func (c *Cred) GetTokenDataFactory() (Auth, error) {
	var auth Auth
	url := strings.Join([]string{AzureToken, c.TenantId, "/oauth2/token"}, "")
	method := "GET"
	data := strings.Join([]string{"grant_type=client_credentials&client_id=", c.ClientId, "&client_secret=", c.ClientSecret, "&resource=", AzureDataFactory}, "")
	payload := strings.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return auth, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return auth, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return auth, err
	}
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return auth, err
	}
	return auth, nil
}
