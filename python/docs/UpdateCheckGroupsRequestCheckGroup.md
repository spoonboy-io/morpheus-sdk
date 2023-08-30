# UpdateCheckGroupsRequestCheckGroup

Payload for creating a new monitoring check group

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Unique name scoped to your account for the check group | [optional] 
**description** | **str** | Optional description field | [optional] 
**min_happy** | **int** | This specifies the minimum number of checks within the group that must be happy to keep the group from becoming unhealthy. | [optional]  if omitted the server will use the default value of 1
**in_uptime** | **bool** | Used to determine if check should affect account wide availability calculations | [optional]  if omitted the server will use the default value of True
**severity** | **str** | Determines the maximum severity level this group can incur on an incident when failing | [optional]  if omitted the server will use the default value of "critical"
**active** | **bool** | Used to determine if check group is active | [optional]  if omitted the server will use the default value of True
**checks** | **[int]** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


