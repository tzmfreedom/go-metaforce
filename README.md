# Go Metaforce

force.com metadata api client for go

## Install

```bash
$ go get github.com/tzmfreedom/metaforce
```

## Usage

```go
import "github.com/tzmfreedom/metaforce"
```

* Login to production/developer
```
client := metaforce.NewClient("", "37.0")
err := client.Login("username", "password")
```

* Login to sandbox
```go
client := metaforce.NewClient("test.salesforce.com", "37.0")
err := client.Login("username", "password")
```

* Retrieve Metadata
```go
res, err := client.Retrieve(&metaforce.RetrieveRequest{
 	ApiVersion: 37.0,
   	PackageNames: []string{"CustomObject"},
   	SinglePackage: true,
   	SpecificFiles: []string{},
   	Unpackaged: nil,
})
```

* Deploy
```go
res, err := client.Deploy(buf.Bytes())
```

* Check Deploy Status

```go
res, err := client.CheckDeployStatus(resultId)
```

* Cancel Deploy

```go
var rId metaforce.ID = "0Af*********"
res, err := client.CancelDeploy(&rId)
```

* Check Retrieve Status

```

```

* Create Metadata

```go
request := []metaforce.MetadataInterface{
  &metaforce.CustomObject{
    FullName: "Go__c",
    Type: "CustomObject",
    DeploymentStatus: metaforce.DeploymentStatusDeployed,
    Description: "from go",
    Label: "Go",
    NameField: &metaforce.CustomField{
      Label: "Go name",
      Length: 80,
      Type: metaforce.FieldTypeText,
    },
    PluralLabel: "Go objects",
    SharingModel: metaforce.SharingModelReadWrite,
  },
}
res, err := client.CreateMetadata(request)
```

* Delete Metadata
```go
res, err := client.DeleteMetadata("CustomObject", []string{ "GO__c" })
```

* Deploy Recent Validation

* Describe Metadata
```go
res, err := client.DescribeMetadata()
```

* Describe ValueType
```go
res, err := client.DescribeValueType("{http://soap.sforce.com/2006/04/metadata}ApexClass")
```

* List Metadata
```go
query := []*metaforce.ListMetadataQuery{
  &metaforce.ListMetadataQuery{
    Type: "ApexClass",
  },
}
res, err := client.ListMetadata(query)
```

* Read Metadata
```go
res, err := client.ReadMetadata("CustomObject", []string{ "GO__c" })
```

* Rename Metadata

```go
res, err := client.RenameMetadata(&metaforce.RenameMetadata{
	Type: "CustomObject",
    OldFullName: "OLD__c",
    NewFullName: "NEW__c",
})
```

* Update Metadata

* Upsert Metadata


