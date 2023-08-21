# # ApiExecuteSchedulesSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A name for the execute schedule |
**description** | **string** | A description for the execute schedule | [optional]
**schedule_type** | **string** | Type of schedule | [optional]
**schedule_timezone** | **string** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional] [default to 'UTC']
**cron** | **string** | Cron Expression. The default is daily at midnight | [optional] [default to '0 0 * * *']
**enabled** | **bool** | Is enabled | [optional] [default to true]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
