# MorpheusApi.AddAlertsRequestAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the alert | 
**minDuration** | **Number** | Duration in minutes of the delay before sending notification(s) | [optional] [default to 0]
**minSeverity** | **String** | Severity level threshold for sending notifications. | [optional] [default to &#39;critical&#39;]
**active** | **Boolean** | Set to false to disable notifications | [optional] [default to true]
**allChecks** | **Boolean** | Trigger for all checks | [optional] [default to false]
**allGroups** | **Boolean** | Trigger for all check groups | [optional] [default to false]
**allApps** | **Boolean** | Trigger for all monitor apps | [optional] [default to false]
**checks** | **[Number]** |  | [optional] 
**groups** | **[Number]** |  | [optional] 
**apps** | **[Number]** |  | [optional] 
**contacts** | [**[AddAlertsRequestAlertContactsInner]**](AddAlertsRequestAlertContactsInner.md) |  | [optional] 



## Enum: MinSeverityEnum


* `info` (value: `"info"`)

* `warning` (value: `"warning"`)

* `critical` (value: `"critical"`)




