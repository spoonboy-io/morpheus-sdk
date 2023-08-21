# # ApiMonitoringGroupsCheckGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Unique name scoped to your account for the check group |
**description** | **string** | Optional description field | [optional]
**min_happy** | **int** | This specifies the minimum number of checks within the group that must be happy to keep the group from becoming unhealthy. | [optional] [default to 1]
**in_uptime** | **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**severity** | **string** | Determines the maximum severity level this group can incur on an incident when failing | [optional] [default to 'critical']
**active** | **bool** | Used to determine if check group is active | [optional] [default to true]
**checks** | **int[]** |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
