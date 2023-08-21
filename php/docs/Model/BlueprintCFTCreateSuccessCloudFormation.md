# # BlueprintCFTCreateSuccessCloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config_type** | **string** | Configuration Type |
**json** | **string** | CloudFormation Template in JSON | [optional]
**yaml** | **string** | CloudFormation Template in YAML | [optional]
**git** | [**\OpenAPI\Client\Model\BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  | [optional]
**iam** | [**OneOfBooleanString**](OneOfBooleanString.md) | CloudFormation Attribute CAPABILITY_IAM | [optional]
**capability_named_iam** | [**OneOfBooleanString**](OneOfBooleanString.md) | CloudFormation Attribute CAPABILITY_NAMED_IAM | [optional]
**capability_auto_expand** | [**OneOfBooleanString**](OneOfBooleanString.md) | CloudFormation Attribute CAPABILITY_AUTO_EXPAND | [optional]
**install_agent** | [**OneOfBooleanString**](OneOfBooleanString.md) | Install Morpheus Agent | [optional]
**cloud_init_enabled** | [**OneOfBooleanString**](OneOfBooleanString.md) | Cloud Init Enabled | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
