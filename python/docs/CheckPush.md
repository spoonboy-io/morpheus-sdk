# CheckPush

A push Check is not polled regularly by the standard monitoring system. Instead it is expected that an external API push updates as to the status of the check timed closely with the configured check interval setting. This is used to throttle the push from performing too many status updates. To push an update using the api key one must send a json payload like so: `curl -XPOST https://<morpheus url>/api/monitoring/push?apiKey=<apiKey> -H 'Content-Type: application/json' -d '{\"success\":true, \"message\": \"any comment goes here\"}'` The API Key will be returned on successful creation or can be found by getting this check. 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Unique name scoped to your account for the check | [optional] 
**description** | **str, none_type** | Optional description field | [optional] 
**check_type** | [**CheckPushCheckType**](CheckPushCheckType.md) |  | [optional] 
**check_interval** | **int** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) | [optional]  if omitted the server will use the default value of 300
**in_uptime** | **bool** | Used to determine if check should affect account wide availability calculations | [optional]  if omitted the server will use the default value of True
**active** | **bool** | Used to determine if check should be scheduled to execute | [optional]  if omitted the server will use the default value of True
**severity** | **str** | Severity level threshold for sending notifications. | [optional]  if omitted the server will use the default value of "critical"
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


