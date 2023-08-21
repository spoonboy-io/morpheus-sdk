

# ApiExecuteSchedulesSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the execute schedule | 
**description** | **String** | A description for the execute schedule |  [optional]
**scheduleType** | [**ScheduleTypeEnum**](#ScheduleTypeEnum) | Type of schedule |  [optional]
**scheduleTimezone** | **String** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. |  [optional]
**cron** | **String** | Cron Expression. The default is daily at midnight |  [optional]
**enabled** | **Boolean** | Is enabled |  [optional]



## Enum: ScheduleTypeEnum

Name | Value
---- | -----
EXECUTE | &quot;execute&quot;



