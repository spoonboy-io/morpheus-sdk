# MorpheusApi.ApiMonitoringGroupsIdCheckGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the check group | [optional] 
**description** | **String** | Optional description field | [optional] 
**minHappy** | **Number** | This specifies the minimum number of checks within the group that must be happy to keep the group from becoming unhealthy. | [optional] [default to 1]
**inUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**severity** | **String** | Determines the maximum severity level this group can incur on an incident when failing | [optional] [default to &#39;critical&#39;]
**active** | **Boolean** | Used to determine if check group is active | [optional] [default to true]
**checks** | **[Number]** |  | [optional] 



## Enum: SeverityEnum


* `info` (value: `"info"`)

* `warning` (value: `"warning"`)

* `critical` (value: `"critical"`)




