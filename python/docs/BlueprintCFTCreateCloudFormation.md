# BlueprintCFTCreateCloudFormation


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config_type** | **str** | Configuration Type | 
**json** | **str** | CloudFormation Template in JSON | [optional] 
**yaml** | **str** | CloudFormation Template in YAML | [optional] 
**git** | [**BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  | [optional] 
**iam** | **bool** | CloudFormation Attribute CAPABILITY_IAM | [optional]  if omitted the server will use the default value of False
**capability_named_iam** | **bool** | CloudFormation Attribute CAPABILITY_NAMED_IAM | [optional]  if omitted the server will use the default value of False
**capability_auto_expand** | **bool** | CloudFormation Attribute CAPABILITY_AUTO_EXPAND | [optional]  if omitted the server will use the default value of False
**install_agent** | **bool** | Install Morpheus Agent | [optional]  if omitted the server will use the default value of False
**cloud_init_enabled** | **bool** | Cloud Init Enabled | [optional]  if omitted the server will use the default value of False
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


