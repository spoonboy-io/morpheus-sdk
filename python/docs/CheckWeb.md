# CheckWeb

Web check type allows you to perform a standard web request and validate the response came back successfully.  Additionally, you can check for matching text within the result. 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Unique name scoped to your account for the check | [optional] 
**description** | **str, none_type** | Optional description field | [optional] 
**check_type** | [**CheckWebCheckType**](CheckWebCheckType.md) |  | [optional] 
**check_interval** | **int** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) | [optional]  if omitted the server will use the default value of 300
**in_uptime** | **bool** | Used to determine if check should affect account wide availability calculations | [optional]  if omitted the server will use the default value of True
**active** | **bool** | Used to determine if check should be scheduled to execute | [optional]  if omitted the server will use the default value of True
**severity** | **str** | Severity level threshold for sending notifications. | [optional]  if omitted the server will use the default value of "critical"
**config** | [**CheckWebConfig**](CheckWebConfig.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


