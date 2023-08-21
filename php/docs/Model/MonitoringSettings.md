# # MonitoringSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_manage_checks** | **bool** | Auto Create Checks | [optional]
**availability_time_frame** | **int** | Availability Time Frame. The number of days availability should be calculated for. Changes will not take effect until your checks have passed their check interval. | [optional]
**availability_precision** | **int** | Availability Precision. The number of decimal places availability should be displayed in. Can be anywhere between 0 and 5. | [optional]
**default_check_interval** | **int** | Default Check Interval. The number of minutes to use as the default interval to use when creating new checks. | [optional]
**service_now** | [**\OpenAPI\Client\Model\MonitoringSettingsServiceNow**](MonitoringSettingsServiceNow.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
