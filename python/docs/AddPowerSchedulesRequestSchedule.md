# AddPowerSchedulesRequestSchedule


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the power schedule | 
**description** | **str** | A description for the power schedule | [optional] 
**schedule_type** | **str** | Type of schedule &#x60;power&#x60; on or &#x60;power off&#x60; | [optional] 
**schedule_timezone** | **str** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional]  if omitted the server will use the default value of "UTC"
**enabled** | **bool** | Is the power schedule enabled | [optional]  if omitted the server will use the default value of True
**monday_on_time** | **str** | Monday Start time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "00:00"
**monday_off_time** | **str** | Monday Off time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "24:00"
**tuesday_on_time** | **str** | Tuesday Start time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "00:00"
**tuesday_off_time** | **str** | Tuesday Off time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "24:00"
**wednesday_on_time** | **str** | Wednesday Start time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "00:00"
**wednesday_off_time** | **str** | Wednesday Off time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "24:00"
**thursday_on_time** | **str** | Thursday Start time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "00:00"
**thursday_off_time** | **str** | Thursday Off time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "24:00"
**friday_on_time** | **str** | Friday Start time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "00:00"
**friday_off_time** | **str** | Friday Off time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "24:00"
**saturday_on_time** | **str** | Saturday Start time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "00:00"
**saturday_off_time** | **str** | Saturday Off time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "24:00"
**sunday_on_time** | **str** | Sunday Start time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "00:00"
**sunday_off_time** | **str** | Sunday Off time of the day in 24-hour format | [optional]  if omitted the server will use the default value of "24:00"
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


