package main

import (
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/go-metaforce"
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
	//deploy()
	deployRecentValidation("0Af2K00000LuQNASA3")
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
	res, err := client.DeleteMetadata("CustomObject", []string{ "GO222__c" })
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func retrieve() {
	res, err := client.Retrieve(&metaforce.RetrieveRequest{
		ApiVersion: 37.0,
		PackageNames: []string{"CustomObject"},
		SinglePackage: true,
		SpecificFiles: []string{},
		Unpackaged: nil,
	})
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func describeMetadata() {
	res, err := client.DescribeMetadata()
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func describeValueType() {
	res, err := client.DescribeValueType("{http://soap.sforce.com/2006/04/metadata}ApexClass")
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func renameMetadata() {
	res, err := client.RenameMetadata(&metaforce.RenameMetadata{
		Type: "CustomObject",
		OldFullName: "GO1__c",
		NewFullName: "GO2__c",
	})
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func updateMetadata() {
	res, err := client.UpdateMetadata([]metaforce.MetadataInterface{
		&metaforce.CustomObject{
			FullName: "GO2__c",
			Type: "CustomObject",
			DeploymentStatus: metaforce.DeploymentStatusDeployed,
			Description: "これはGOから作ってるよ 2",
			Label: "GOミラクルオブジェクト",
			NameField: &metaforce.CustomField{
				Label: "GO名",
				Length: 80,
				Type: metaforce.FieldTypeText,
			},
			PluralLabel: "GOミラクルオブジェクツ",
			SharingModel: metaforce.SharingModelReadWrite,
		},
	})
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func upsertMetadata() {
	res, err := client.UpsertMetadata([]metaforce.MetadataInterface{
		&metaforce.CustomObject{
			FullName: "GO1__c",
			Type: "CustomObject",
			DeploymentStatus: metaforce.DeploymentStatusDeployed,
			Description: "これはGOから作ってるよ 3",
			Label: "GOミラクルオブジェクト",
			NameField: &metaforce.CustomField{
				Label: "GO名",
				Length: 80,
				Type: metaforce.FieldTypeText,
			},
			PluralLabel: "GOミラクルオブジェクツ",
			SharingModel: metaforce.SharingModelReadWrite,
		},
	})
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func deployRecentValidation(validationId string) {
	res, err := client.DeployRecentValidation(validationId)
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}

func deploy() {
	buf := []byte("bytes")
	res, err := client.Deploy(buf, &metaforce.DeployOptions{
		CheckOnly: true,
	})
	if err != nil {
		panic(err)
	}
	pp.Println(res)
}
