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

```

```

* Delete Metadata

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

* Rename Metadata

* Retrieve

* Update Metadata

* Upsert Metadata
