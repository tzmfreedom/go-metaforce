package metaforce

import (
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
)

const (
	DefaultApiVersion = "44.0"
	DefaultLoginUrl   = "login.salesforce.com"
)

type Client struct {
	ApiVersion string
	ServerUrl  string
	LoginUrl   string
	portType    *MetadataPortType
	loginResult *LoginResult
	debug       bool
}

func NewClient() *Client {
	portType := NewMetadataPortType("", true, nil)
	return &Client{
		portType: portType,
		LoginUrl: DefaultLoginUrl,
		ApiVersion: DefaultApiVersion,
	}
}

func (client *Client) SetDebug(debug bool) {
	client.portType.SetDebug(debug)
}

func (c *Client) SetApiVersion(v string) {
	c.ApiVersion = v
	c.setLoginUrl()
}

func (c *Client) SetAccessToken(sid string) {
	sessionHeader := &SessionHeader{
		SessionId: sid,
	}
	c.portType.SetHeader(sessionHeader)
}

func (c *Client) GetSessionID() string {
	if c.loginResult == nil {
		return ""
	}
	return c.loginResult.SessionId
}

func (c *Client) GetServerURL() string {
	return c.portType.client.GetServerUrl()
}

func (c *Client) SetLoginUrl(url string) {
	c.LoginUrl = url
	c.setLoginUrl()
}

func (c *Client) setLoginUrl() {
	url := fmt.Sprintf("https://%s/services/Soap/u/%s", c.LoginUrl, c.ApiVersion)
	c.portType.SetServerUrl(url)
}

func (c *Client) SetLogger(logger io.Writer) {
	c.portType.SetLogger(logger)
}

func (c *Client) SetGzip(gz bool) {
	c.portType.SetGzip(gz)
}

//func (c *Client) Logout() error {
//	_, err := c.portType.Logout(&soapforce.Logout{})
//	if err != nil {
//		return err
//	}
//	c.ServerUrl = ""
//	c.setLoginUrl()
//	c.portType.ClearHeader()
//	return nil
//}

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
	f, err := strconv.ParseFloat(client.ApiVersion, 32)
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
	f, err := strconv.ParseFloat(client.ApiVersion, 32)
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

func (client *Client) DeployRecentValidation(validationId string) (*DeployRecentValidationResponse, error) {
	return client.portType.DeployRecentValidation(&DeployRecentValidation{
		ValidationId: ID(validationId),
	})
}

func (client *Client) ReadMetadataInto(typeName string, fullNames []string, response interface{}) error {
	request := ReadMetadata{
		FullNames: fullNames,
		Type:      typeName,
	}
	return client.portType.ReadMetadataInto(&request, response)
}

