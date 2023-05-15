package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-athena/authenticator"
	"github.com/Appkube-awsx/awsx-athena/client"
	"github.com/Appkube-awsx/awsx-athena/cmd/athenacmd"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/spf13/cobra"
)

var AwsxAthenaMetadataCmd = &cobra.Command{
	Use:   "getListWorkMetaDataDetails",
	Short: "getListWorkMetaDataDetails command gets resource counts",
	Long:  `getListWorkMetaDataDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command athena started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			getListWorkGroups(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}

// json.Unmarshal
func getListWorkGroups(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*athena.ListWorkGroupsOutput, error) {
	log.Println("getting workGroup metadata list summary")

	listClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	listRequest := &athena.ListWorkGroupsInput{}

	listResponse, err := listClient.ListWorkGroups(listRequest)

	if err != nil {
		log.Fatalln("Error:in getting  workGroup list", err)
	}
	log.Println(listResponse)
	return listResponse, err
}

func Execute() {
	err := AwsxAthenaMetadataCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxAthenaMetadataCmd.AddCommand(athenacmd.GetConfigDataCmd)

	AwsxAthenaMetadataCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxAthenaMetadataCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxAthenaMetadataCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxAthenaMetadataCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxAthenaMetadataCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxAthenaMetadataCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxAthenaMetadataCmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
