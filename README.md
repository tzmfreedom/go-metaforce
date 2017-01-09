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
client := metaforce.NewForceClient("", "v37.0")
err := client.Login("username", "password")
```

* Login to sandbox
```golang
client := metaforce.NewForceClient("test.salesforce.com", "v37.0")
err := client.Login("username", "password")
```

* Retrieve Metadata
```golang

```

* Deploy
```golang
response, err := client.Deploy(buf.Bytes())
```

* Check Deploy Status

```golang
response, err := client.CheckDeployStatus(resultId)
```

* Cancel Deploy

```golang
response, err := client.
```

* Check Retrieve Status

* Create Metadata

* Delete Metadata

* Deploy Recent Validation

* Describe Metadata

* Describe ValueType

* List Metadata

* Read Metadata

* Rename Metadata

* Retrieve

* Update Metadata

* Upsert Metadata
