# InstanceScheduleCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**schedule_type** | **str** |  | [optional]  if omitted the server will use the default value of "dayOfWeek"
**schedule_timezone** | **str** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional]  if omitted the server will use the default value of "UTC"
**start_day_of_week** | **int** | Start day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**start_time** | **str** | Start time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**end_day_of_week** | **int** | End day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**end_time** | **str** | End time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**start_date** | **datetime** | Start Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional] 
**end_date** | **datetime** | End Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional] 
**threshold** | [**InstanceScheduleCreateThreshold**](InstanceScheduleCreateThreshold.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


