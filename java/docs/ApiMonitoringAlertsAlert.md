

# ApiMonitoringAlertsAlert

Payload for creating a new monitoring alert
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the alert | 
**minDuration** | **Integer** | Duration in minutes of the delay before sending notification(s) |  [optional]
**minSeverity** | [**MinSeverityEnum**](#MinSeverityEnum) | Severity level threshold for sending notifications. |  [optional]
**active** | **Boolean** | Set to false to disable notifications |  [optional]
**allChecks** | **Boolean** | Trigger for all checks |  [optional]
**allGroups** | **Boolean** | Trigger for all check groups |  [optional]
**allApps** | **Boolean** | Trigger for all monitor apps |  [optional]
**checks** | **List&lt;Integer&gt;** |  |  [optional]
**groups** | **List&lt;Integer&gt;** |  |  [optional]
**apps** | **List&lt;Integer&gt;** |  |  [optional]
**contacts** | [**List&lt;ApiMonitoringAlertsAlertContacts&gt;**](ApiMonitoringAlertsAlertContacts.md) |  |  [optional]



## Enum: MinSeverityEnum

Name | Value
---- | -----
INFO | &quot;info&quot;
WARNING | &quot;warning&quot;
CRITICAL | &quot;critical&quot;



