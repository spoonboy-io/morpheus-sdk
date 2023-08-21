

# InstanceScheduleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**scheduleType** | [**ScheduleTypeEnum**](#ScheduleTypeEnum) |  |  [optional]
**scheduleTimezone** | **String** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. Only used and required for scheduleType &#x60;dayOfWeek&#x60; |  [optional]
**startDayOfWeek** | **Integer** | Start day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; |  [optional]
**startTime** | **String** | Start time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; |  [optional]
**endDayOfWeek** | **Integer** | End day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; |  [optional]
**endTime** | **String** | End time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; |  [optional]
**startDate** | **OffsetDateTime** | Start Date. Only used and required for scheduleType &#x60;exact&#x60; |  [optional]
**endDate** | **OffsetDateTime** | End Date. Only used and required for scheduleType &#x60;exact&#x60; |  [optional]
**threshold** | [**InstanceScheduleCreateThreshold**](InstanceScheduleCreateThreshold.md) |  |  [optional]



## Enum: ScheduleTypeEnum

Name | Value
---- | -----
DAYOFWEEK | &quot;dayOfWeek&quot;
EXACT | &quot;exact&quot;



