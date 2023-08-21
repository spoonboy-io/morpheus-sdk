# MorpheusApi.CheckPush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the check | [optional] 
**description** | **String** | Optional description field | [optional] 
**checkType** | [**CheckPushCheckType**](CheckPushCheckType.md) |  | [optional] 
**checkInterval** | **Number** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) | [optional] [default to 300]
**inUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**active** | **Boolean** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**severity** | **String** | Severity level threshold for sending notifications. | [optional] [default to &#39;critical&#39;]



## Enum: SeverityEnum


* `info` (value: `"info"`)

* `warning` (value: `"warning"`)

* `critical` (value: `"critical"`)




