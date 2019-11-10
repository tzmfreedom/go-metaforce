package main

import (
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/metaforce"
	"os"
)

var client *metaforce.Client

func main() {
	client = metaforce.NewClient("", "47.0")
	client.SetDebug(true)
	err := client.Login(os.Getenv("SALESFORCE_USERNAME"), os.Getenv("SALESFORCE_PASSWORD"))
	if err != nil {
		panic(err)
	}
	createMetadata()
	//deleteMetadata()
}

func listMetadata() {
	query := []*metaforce.ListMetadataQuery{
		{
			Type: "ApexClass",
		},
	}
	res, err := client.ListMetadata(query)
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func readMetadata() {
	res, err := client.ReadMetadata("CustomObject", []string{ "GO__c" })
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func createMetadata() {
	request := []metaforce.MetadataInterface{
		&metaforce.CustomObject{
			FullName: "GO222__c",
			Type: "CustomObject",
			DeploymentStatus: metaforce.DeploymentStatusDeployed,
			Description: "これはGOから作ってるよ",
			Label: "GOミラクルオブジェクト",
			NameField: &metaforce.CustomField{
				Label: "GO名",
				Length: 80,
				Type: metaforce.FieldTypeText,
			},
			PluralLabel: "GOミラクルオブジェクツ",
			SharingModel: metaforce.SharingModelReadWrite,
		},
	}
	res, err := client.CreateMetadata(request)
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func deleteMetadata() {
	res, err := client.DeleteMetadata("CustomObject", []string{ "GO__c" })
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

