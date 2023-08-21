# # ApiMonitoringAppsMonitorApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Unique name scoped to your account for the check app |
**description** | **string** | Optional description field | [optional]
**in_uptime** | **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**severity** | **string** | Severity level of incidents that are created when this check fails | [optional] [default to 'critical']
**active** | **bool** | Used to determine if check app is active | [optional] [default to true]
**checks** | **int[]** |  | [optional]
**check_groups** | **int[]** |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
