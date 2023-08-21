# # BlueprintCFTCreateCloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config_type** | **string** | Configuration Type |
**json** | **string** | CloudFormation Template in JSON | [optional]
**yaml** | **string** | CloudFormation Template in YAML | [optional]
**git** | [**\OpenAPI\Client\Model\BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  | [optional]
**iam** | **bool** | CloudFormation Attribute CAPABILITY_IAM | [optional] [default to false]
**capability_named_iam** | **bool** | CloudFormation Attribute CAPABILITY_NAMED_IAM | [optional] [default to false]
**capability_auto_expand** | **bool** | CloudFormation Attribute CAPABILITY_AUTO_EXPAND | [optional] [default to false]
**install_agent** | **bool** | Install Morpheus Agent | [optional] [default to false]
**cloud_init_enabled** | **bool** | Cloud Init Enabled | [optional] [default to false]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
