# # InstanceScheduleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**schedule_type** | **string** |  | [optional] [default to 'dayOfWeek']
**schedule_timezone** | **string** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] [default to 'UTC']
**start_day_of_week** | **int** | Start day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional]
**start_time** | **string** | Start time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional]
**end_day_of_week** | **int** | End day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional]
**end_time** | **string** | End time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional]
**start_date** | [**\DateTime**](\DateTime.md) | Start Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional]
**end_date** | [**\DateTime**](\DateTime.md) | End Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional]
**threshold** | [**\OpenAPI\Client\Model\InstanceScheduleUpdateThreshold**](InstanceScheduleUpdateThreshold.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
