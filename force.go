package metaforce

import (
	"encoding/base64"
	"strconv"
)

type ForceClient struct {
	portType    *MetadataPortType
	loginResult *LoginResult
	apiVersion  string
}

func NewForceClient(endpoint string, apiversion string) *ForceClient {
	if endpoint == "" {
		endpoint = "login.salesforce.com"
	}
	portType := NewMetadataPortType("https://"+endpoint+"/services/Soap/u/"+apiversion, true, nil)
	return &ForceClient{
		portType: portType,
		apiVersion: apiversion,
	}
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
	request := CheckDeployStatus{AsyncProcessId: resultId, IncludeDetails: includeDetails}
	return client.portType.CheckDeployStatus(&request)
}

func (client *ForceClient) CancelDeploy(processId *ID) (*CancelDeployResponse, error) {
	request := CancelDeploy{AsyncProcessId: processId}
	return client.portType.CancelDeploy(&request)
}

func (client *ForceClient) DescribeMetadata() (*DescribeMetadataResponse, error) {
	f, err := strconv.ParseFloat(client.apiVersion, 32)
	if err != nil {
		f = 37.0
	}

	request := DescribeMetadata{AsOfVersion: f}
	return client.portType.DescribeMetadata(&request)
}

func (client *ForceClient) DescribeValueType(desc_type string) (*DescribeValueTypeResponse, error) {
	request := DescribeValueType{
		Type: desc_type,
	}
	return client.portType.DescribeValueType(&request)
}

func (client *ForceClient) ListMetadata(listMetadataQuery []*ListMetadataQuery) (*ListMetadataResponse, error) {
	f, err := strconv.ParseFloat(client.apiVersion, 32)
	if err != nil {
		f = 37.0
	}

	request := ListMetadata{
		Queries: listMetadataQuery,
		AsOfVersion: f,
	}
	return client.portType.ListMetadata(&request)
}


//func (client *ForceClient) Retrieve() () {
//	retrieve_request := Retrieve{
//		RetrieveRequest: &RetrieveRequest{
//			ApiVersion: "37.0",
//			PackageNames,
//			SinglePackage,
//			SpecificFiles,
//			Unpackaged *Package,
//		},
//	}
//	return client.portType.Retrieve(retrieve_request)
//}
