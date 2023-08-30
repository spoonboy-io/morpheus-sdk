# PolicyCloudCreatePolicyType


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **str** | The policy type | [optional] 
**config** | [**PolicyGroupCreatePolicyTypeConfig**](PolicyGroupCreatePolicyTypeConfig.md) |  | [optional] 
**enabled** | **bool** | Set to false to disable | [optional]  if omitted the server will use the default value of True
**ref_type** | **str** | Scope object type | [optional]  if omitted the server will use the default value of "ComputeZone"
**ref_id** | **int** | Scope object ID (&#x60;cloud&#x60;) | [optional] 
**accounts** | **[int]** | Array of tenants to scope the policy to | [optional] 
**each_user** | **bool** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


