package metaforce

import (
	"encoding/base64"
	"strconv"
)

type Client struct {
	portType    *MetadataPortType
	loginResult *LoginResult
	debug       bool
	apiVersion  string
}

func NewClient(endpoint string, apiversion string) *Client {
	if endpoint == "" {
		endpoint = "login.salesforce.com"
	}
	portType := NewMetadataPortType("https://"+endpoint+"/services/Soap/u/"+apiversion, true, nil)
	return &Client{
		portType: portType,
		apiVersion: apiversion,
	}
}

func (client *Client) SetDebug(debug bool) {
	client.portType.SetDebug(debug)
}

func (client *Client) Login(username string, password string) error {
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

func (client *Client) Deploy(buf []byte, options *DeployOptions) (*DeployResponse, error) {
	request := Deploy{
		ZipFile:       base64.StdEncoding.EncodeToString(buf),
		DeployOptions: options,
	}
	return client.portType.Deploy(&request)
}

func (client *Client) CheckDeployStatus(resultId string, includeDetails bool) (*CheckDeployStatusResponse, error) {
	request := CheckDeployStatus{AsyncProcessId: ID(resultId), IncludeDetails: includeDetails}
	return client.portType.CheckDeployStatus(&request)
}

func (client *Client) CancelDeploy(processId string) (*CancelDeployResponse, error) {
	request := CancelDeploy{AsyncProcessId: ID(processId)}
	return client.portType.CancelDeploy(&request)
}

func (client *Client) DescribeMetadata() (*DescribeMetadataResponse, error) {
	f, err := strconv.ParseFloat(client.apiVersion, 32)
	if err != nil {
		f = 37.0
	}

	request := DescribeMetadata{AsOfVersion: f}
	return client.portType.DescribeMetadata(&request)
}

func (client *Client) DescribeValueType(desc_type string) (*DescribeValueTypeResponse, error) {
	request := DescribeValueType{
		Type: desc_type,
	}
	return client.portType.DescribeValueType(&request)
}

func (client *Client) ListMetadata(listMetadataQuery []*ListMetadataQuery) (*ListMetadataResponse, error) {
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

func (client *Client) CreateMetadata(metadata []MetadataInterface) (*CreateMetadataResponse, error) {
	request := CreateMetadata{
		Metadata: metadata,
	}
	return client.portType.CreateMetadata(&request)
}

func (client *Client) DeleteMetadata(typeName string, fullNames []string) (*DeleteMetadataResponse, error) {
	request := DeleteMetadata{
		FullNames: fullNames,
		Type: typeName,
	}
	return client.portType.DeleteMetadata(&request)
}

func (client *Client) ReadMetadata(typeName string, fullNames []string) (*ReadMetadataResponse, error) {
	request := ReadMetadata{
		FullNames: fullNames,
		Type: typeName,
	}
	return client.portType.ReadMetadata(&request)
}

func (client *Client) Retrieve(retrieveRequest *RetrieveRequest) (*RetrieveResponse, error) {
	r := &Retrieve{
		RetrieveRequest: retrieveRequest,
	}
	return client.portType.Retrieve(r)
}

func (client *Client) RenameMetadata(r *RenameMetadata) (*RenameMetadataResponse, error) {
	return client.portType.RenameMetadata(r)
}

func (client *Client) UpdateMetadata(metadata []MetadataInterface) (*UpdateMetadataResponse, error) {
	return client.portType.UpdateMetadata(&UpdateMetadata{Metadata: metadata})
}

func (client *Client) UpsertMetadata(metadata []MetadataInterface) (*UpsertMetadataResponse, error) {
	return client.portType.UpsertMetadata(&UpsertMetadata{Metadata: metadata})
}