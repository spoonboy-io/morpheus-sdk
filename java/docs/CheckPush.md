

# CheckPush

A push Check is not polled regularly by the standard monitoring system. Instead it is expected that an external API push updates as to the status of the check timed closely with the configured check interval setting. This is used to throttle the push from performing too many status updates. To push an update using the api key one must send a json payload like so: `curl -XPOST https://<morpheus url>/api/monitoring/push?apiKey=<apiKey> -H 'Content-Type: application/json' -d '{\"success\":true, \"message\": \"any comment goes here\"}'` The API Key will be returned on successful creation or can be found by getting this check. 
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the check |  [optional]
**description** | **String** | Optional description field |  [optional]
**checkType** | [**CheckPushCheckType**](CheckPushCheckType.md) |  |  [optional]
**checkInterval** | **Integer** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) |  [optional]
**inUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations |  [optional]
**active** | **Boolean** | Used to determine if check should be scheduled to execute |  [optional]
**severity** | [**SeverityEnum**](#SeverityEnum) | Severity level threshold for sending notifications. |  [optional]



## Enum: SeverityEnum

Name | Value
---- | -----
INFO | &quot;info&quot;
WARNING | &quot;warning&quot;
CRITICAL | &quot;critical&quot;



