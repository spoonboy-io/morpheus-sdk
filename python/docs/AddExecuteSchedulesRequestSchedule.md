# AddExecuteSchedulesRequestSchedule


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the execute schedule | 
**description** | **str** | A description for the execute schedule | [optional] 
**schedule_type** | **str** | Type of schedule | [optional]  if omitted the server will use the default value of "execute"
**schedule_timezone** | **str** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional]  if omitted the server will use the default value of "UTC"
**cron** | **str** | Cron Expression. The default is daily at midnight | [optional]  if omitted the server will use the default value of "0 0 * * *"
**enabled** | **bool** | Is enabled | [optional]  if omitted the server will use the default value of True
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


