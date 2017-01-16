# Go Metaforce

force.com metadata api client for golang

## Install

```bash
$ go get github.com/tzmfreedom/metaforce
```

## Usage

```golang
import "github.com/tzmfreedom/metaforce"
```

* Login to production/developer
```
client := metaforce.NewForceClient("", "37.0")
err := client.Login("username", "password")
```

* Login to sandbox
```golang
client := metaforce.NewForceClient("test.salesforce.com", "37.0")
err := client.Login("username", "password")
```

* Retrieve Metadata
```golang

```

* Deploy
```golang
res, err := client.Deploy(buf.Bytes())
```

* Check Deploy Status

```golang
res, err := client.CheckDeployStatus(resultId)
```

* Cancel Deploy

```golang
var rId metaforce.ID = "0Af*********"
res, err := client.CancelDeploy(&rId)
```

* Check Retrieve Status

```

```

* Create Metadata

```golang
status := metaforce.DeploymentStatus("Deployed")
sharing := metaforce.SharingModel("ReadWrite")
meta_type := metaforce.FieldType("Text")
request := []metaforce.MetadataInterface{
  &metaforce.CustomObject{
    FullName: "Go__c",
    Type: "CustomObject",
    DeploymentStatus: &status,
    Description: "from golang",
    Label: "Go",
    NameField: &metaforce.CustomField{
      Label: "Go name",
      Length: 80,
      Type: &meta_type,
    },
    PluralLabel: "Go objects",
    SharingModel: &sharing,
  },
}
res, err := client.CreateMetadata(request)
```

* Delete Metadata
```golang
res, err := client.DeleteMetadata("CustomObject", []string{ "GO__c" })
```

* Deploy Recent Validation

* Describe Metadata
```golang
res, err := client.DescribeMetadata()
```

* Describe ValueType
```golang
res, err := client.DescribeValueType("{http://soap.sforce.com/2006/04/metadata}ApexClass")
```

* List Metadata
```golang
query := []*metaforce.ListMetadataQuery{
  &metaforce.ListMetadataQuery{
    Type: "ApexClass",
  },
}
res, err := client.ListMetadata(query)
```

* Read Metadata
```golang
res, err := client.ReadMetadata("CustomObject", []string{ "GO__c" })
```

* Rename Metadata

* Update Metadata

* Upsert Metadata
