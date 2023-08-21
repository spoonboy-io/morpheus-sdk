

# ApiMonitoringGroupsCheckGroup

Payload for creating a new monitoring check group
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the check group | 
**description** | **String** | Optional description field |  [optional]
**minHappy** | **Integer** | This specifies the minimum number of checks within the group that must be happy to keep the group from becoming unhealthy. |  [optional]
**inUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations |  [optional]
**severity** | [**SeverityEnum**](#SeverityEnum) | Determines the maximum severity level this group can incur on an incident when failing |  [optional]
**active** | **Boolean** | Used to determine if check group is active |  [optional]
**checks** | **List&lt;Integer&gt;** |  |  [optional]



## Enum: SeverityEnum

Name | Value
---- | -----
INFO | &quot;info&quot;
WARNING | &quot;warning&quot;
CRITICAL | &quot;critical&quot;



