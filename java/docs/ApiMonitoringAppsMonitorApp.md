

# ApiMonitoringAppsMonitorApp

Payload for creating a new monitoring check app
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the check app | 
**description** | **String** | Optional description field |  [optional]
**inUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations |  [optional]
**severity** | [**SeverityEnum**](#SeverityEnum) | Severity level of incidents that are created when this check fails |  [optional]
**active** | **Boolean** | Used to determine if check app is active |  [optional]
**checks** | **List&lt;Integer&gt;** |  |  [optional]
**checkGroups** | **List&lt;Integer&gt;** |  |  [optional]



## Enum: SeverityEnum

Name | Value
---- | -----
INFO | &quot;info&quot;
WARNING | &quot;warning&quot;
CRITICAL | &quot;critical&quot;



