package main

import (
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/metaforce"
	"os"
)

func main() {
	client := metaforce.NewForceClient("", "v37.0")
	err := client.Login(os.Getenv("SF_USERNAME"), os.Getenv("SF_PASSWORD"))
	if err != nil {
		panic(err)
	}

	pp.Print(client)
	client.CancelDeploy("")

}
