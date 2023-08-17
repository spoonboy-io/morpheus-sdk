# MorpheusApi.AddCheckAppsRequestMonitorApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the check app | 
**description** | **String** | Optional description field | [optional] 
**inUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**severity** | **String** | Severity level of incidents that are created when this check fails | [optional] [default to &#39;critical&#39;]
**active** | **Boolean** | Used to determine if check app is active | [optional] [default to true]
**checks** | **[Number]** |  | [optional] 
**checkGroups** | **[Number]** |  | [optional] 



## Enum: SeverityEnum


* `info` (value: `"info"`)

* `warning` (value: `"warning"`)

* `critical` (value: `"critical"`)




