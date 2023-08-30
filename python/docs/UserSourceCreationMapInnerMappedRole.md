# UserSourceCreationMapInnerMappedRole

This object defines a mapping from the Morpheus role to the identity source role. The Morpheus role is identified by passing `mappedRole.id` or `mappedRole.authority` and the identity source role must be identified by passing either `sourceRoleFqn` or `sourceRoleName` to match on. **NOTE** This replaces the current mappings and must contain the entire set of role mappings you would like to set. Anything not defined will be removed. 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Role ID of the Morpheus role in the identity source tenant | [optional] 
**authority** | **str** | Role authority of the Morpheus role in the identity source tenant | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


