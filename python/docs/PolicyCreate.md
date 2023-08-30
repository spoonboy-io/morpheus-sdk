# PolicyCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the policy | 
**policy_type** | [**PolicyCreatePolicyType**](PolicyCreatePolicyType.md) |  | 
**config** | [**PolicyCreateConfig**](PolicyCreateConfig.md) |  | 
**description** | **str** | A description for the policy | [optional] 
**enabled** | **bool** | Set to false to disable | [optional]  if omitted the server will use the default value of True
**ref_type** | **str, none_type** | Scope object type.  If none specified, will default to Global (null) | [optional] 
**ref_id** | **int** | Scope object ID (&#x60;group&#x60;,&#x60;cloud&#x60;,&#x60;user&#x60;, etc) | [optional] 
**accounts** | **[int]** | Array of tenants to scope the policy to | [optional] 
**each_user** | **bool** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


