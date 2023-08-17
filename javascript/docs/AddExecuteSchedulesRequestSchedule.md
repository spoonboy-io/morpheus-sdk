# MorpheusApi.AddExecuteSchedulesRequestSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the execute schedule | 
**description** | **String** | A description for the execute schedule | [optional] 
**scheduleType** | **String** | Type of schedule | [optional] 
**scheduleTimezone** | **String** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional] [default to &#39;UTC&#39;]
**cron** | **String** | Cron Expression. The default is daily at midnight | [optional] [default to &#39;0 0 * * *&#39;]
**enabled** | **Boolean** | Is enabled | [optional] [default to true]



## Enum: ScheduleTypeEnum


* `execute` (value: `"execute"`)




