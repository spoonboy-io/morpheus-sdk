# MonitoringSettings


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_manage_checks** | **bool** | Auto Create Checks | [optional] 
**availability_time_frame** | **int, none_type** | Availability Time Frame. The number of days availability should be calculated for. Changes will not take effect until your checks have passed their check interval. | [optional] 
**availability_precision** | **int, none_type** | Availability Precision. The number of decimal places availability should be displayed in. Can be anywhere between 0 and 5. | [optional] 
**default_check_interval** | **int, none_type** | Default Check Interval. The number of minutes to use as the default interval to use when creating new checks. | [optional] 
**service_now** | [**MonitoringSettingsServiceNow**](MonitoringSettingsServiceNow.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


