package metaforce

import (
	"encoding/base64"
)

type ForceClient struct {
	portType    *MetadataPortType
	loginResult *LoginResult
}

func NewForceClient(endpoint string, apiversion string) *ForceClient {
	if endpoint == "" {
		endpoint = "login.salesforce.com"
	}
	portType := NewMetadataPortType("https://"+endpoint+"/services/Soap/u/"+apiversion, true, nil)
	return &ForceClient{portType: portType}
}

func (client *ForceClient) Login(username string, password string) error {
	loginRequest := LoginRequest{Username: username, Password: password}
	loginResponse, err := client.portType.Login(&loginRequest)
	if err != nil {
		return err
	}
	client.loginResult = &loginResponse.LoginResult
	sessionHeader := SessionHeader{
		SessionId: client.loginResult.SessionId,
	}
	client.portType.SetHeader(&sessionHeader)
	client.portType.SetServerUrl(client.loginResult.MetadataServerUrl)
	return nil
}

func (client *ForceClient) Deploy(buf []byte, options *DeployOptions) (*DeployResponse, error) {
	request := Deploy{
		ZipFile:       base64.StdEncoding.EncodeToString(buf),
		DeployOptions: options,
	}
	return client.portType.Deploy(&request)
}

func (client *ForceClient) CheckDeployStatus(resultId *ID, includeDetails bool) (*CheckDeployStatusResponse, error) {
	check_request := CheckDeployStatus{AsyncProcessId: resultId, IncludeDetails: includeDetails}
	return client.portType.CheckDeployStatus(&check_request)
}

func (client *ForceClient) CancelDeploy(resultId *ID) (*CancelDeployResponse, error) {
	cancel_deploy_request := CancelDeploy{AsyncProcessId: resultId}
	return client.portType.CancelDeploy(cancel_deploy_request)
}
