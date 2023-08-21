# MorpheusApi.InstanceSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Number** |  | [optional] 
**scheduleType** | **String** |  | [optional] [default to &#39;dayOfWeek&#39;]
**scheduleTimezone** | **String** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] [default to &#39;UTC&#39;]
**startDayOfWeek** | **Number** | Start day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**startTime** | **String** | Start time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**endDayOfWeek** | **Number** | End day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**endTime** | **String** | End time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**startDate** | **Date** | Start Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional] 
**endDate** | **Date** | End Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional] 
**startDisplay** | **String** | Start day and time or start date formatted for display | [optional] 
**endDisplay** | **String** | End day and time or end date formatted for display | [optional] 
**threshold** | [**InstanceScheduleThreshold**](InstanceScheduleThreshold.md) |  | [optional] 
**dateCreated** | **Date** |  | [optional] 
**lastUpdated** | **Date** |  | [optional] 



## Enum: ScheduleTypeEnum


* `dayOfWeek` (value: `"dayOfWeek"`)

* `exact` (value: `"exact"`)




