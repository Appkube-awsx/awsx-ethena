package athenacmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-athena/authenticator"
	"github.com/Appkube-awsx/awsx-athena/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		// print(authFlag)
		// authFlag := true
		if authFlag {
			workGroupName, _ := cmd.Flags().GetString("workGroupName")
			if workGroupName != "" {
				getWorkGroupDetails(region, crossAccountRoleArn, acKey, secKey, workGroupName, externalId)
			} else {
				log.Fatalln("workGroupName not provided. Program exit")
			}
		}
	},
}

func getWorkGroupDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, workGroupName string, externalId string) *athena.GetWorkGroupOutput {
	log.Println("Getting aws cluster data")
	
	listWorkClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &athena.GetWorkGroupInput{
		WorkGroup: aws.String(workGroupName),
	}

	workDetailsResponse, err := listWorkClient.GetWorkGroup(input)

	log.Println(workDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return workDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("workGroupName", "t", "", "workGroupName")

	if err := GetConfigDataCmd.MarkFlagRequired("workGroupName"); err != nil {
		fmt.Println(err)
	}
}
