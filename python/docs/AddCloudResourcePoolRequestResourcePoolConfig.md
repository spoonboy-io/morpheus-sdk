# AddCloudResourcePoolRequestResourcePoolConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cidr_block** | **str** | Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block) | [optional] 
**tenancy** | **str** | default or dedicated | [optional]  if omitted the server will use the default value of "default"
**managers** | **[str]** | Array of manager usernames | [optional] 
**developers** | **[str]** | Array of developer usernames | [optional] 
**auditors** | **[str]** | Array of auditor usernames | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


