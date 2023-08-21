# MorpheusApi.InstanceScheduleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**scheduleType** | **String** |  | [optional] [default to &#39;dayOfWeek&#39;]
**scheduleTimezone** | **String** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] [default to &#39;UTC&#39;]
**startDayOfWeek** | **Number** | Start day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**startTime** | **String** | Start time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**endDayOfWeek** | **Number** | End day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**endTime** | **String** | End time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**startDate** | **Date** | Start Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional] 
**endDate** | **Date** | End Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional] 
**threshold** | [**InstanceScheduleUpdateThreshold**](InstanceScheduleUpdateThreshold.md) |  | [optional] 



## Enum: ScheduleTypeEnum


* `dayOfWeek` (value: `"dayOfWeek"`)

* `exact` (value: `"exact"`)




