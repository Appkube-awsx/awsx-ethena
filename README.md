    - [What is awsx-athena](#awsx-athena)

- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [what it does ](#what-it-does)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx-athena

This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture-phase2.svg)

This plugin subcommand will implement the Apis' related to Athena services , primarily the following API's:

- getConfigData

This cli collect data from metric/logs/traces of the Athena services and produce the data in a form that Appkube Platform expects.

This CLI , interacts with other Appkube services like Appkube vault , Appkube cloud CMDB so that it can talk with cloud services as
well as filter and sort the information in terms of product/services, so that Appkube platform gets the data that it expects from the cli.

# How to write plugin subcommand

Please refer to the instruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build/test/debug/publish and integrate into the main commmand.

# How to build / Test

            go run main.go
                - Program will print Calling aws-cloudelements on console

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-athena) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
           awsx-athena getConfigData --zone=us-east-1 --accessKey=xxxxxxxxxx --secretKey=xxxxxxxxxx --crossAccountRoleArn=xxxxxxxxxx  --externalId=xxxxxxxxxx

# what it does

This subcommand implement the following functionalities -
getConfigData - It will get the resource count summary for a given AWS account id and region.

# command input

1. --valutURL = specifies the URL of the AWS Key Management Service (KMS) customer master key (CMK) that you want to use to encrypt a lambda.
2. --acountId = specifies the AWS account ID that the lambda belongs to.
3. --zone = specifies the AWS region where the athena is located.
4. --accessKey = specifies the AWS access key to use for authentication.
5. --secretKey = specifies the AWS secret key to use for authentication.
6. --crossAccountRoleArn = specifies the Amazon Resource Name (ARN) of the role that allows access to a athena in another account.
7. --external Id = The AWS External id.
8. --workGroupName= Insert your workGroupName from athena service in aws account.

# command output

WorkGroups: [
{
CreationTime: 2022-04-26 07:33:07.362 +0000 UTC,
Description: "",
EngineVersion: {
EffectiveEngineVersion: "Athena engine version 2",
SelectedEngineVersion: "AUTO"
},
Name: "Abdul-WG",
State: "ENABLED"
},
{
CreationTime: 2022-04-26 07:14:07.144 +0000 UTC,
Description: "",
EngineVersion: {
EffectiveEngineVersion: "Athena engine version 2",
SelectedEngineVersion: "AUTO"
},
Name: "Trainee-WG",
State: "ENABLED"
}
]

# How to run

From main awsx command , it is called as follows:

```bash
awsx-athena  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>
```

If you build it locally , you can simply run it as standalone command as:

```bash
go run main.go  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

# awsx-athena

athena extension

# AWSX Commands for AWSX-Athena Cli's :

1. CMD used to get list of athena instance's :

```bash
./awsx-athena --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

2. CMD used to get Config data (metadata) of AWS Athena instances :

```bash
./awsx-athena --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<> getConfigData --workGroupName=<>
```
