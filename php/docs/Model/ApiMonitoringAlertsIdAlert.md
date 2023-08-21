# # ApiMonitoringAlertsIdAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Unique name scoped to your account for the alert | [optional]
**min_duration** | **int** | Duration in minutes of the delay before sending notification(s) | [optional] [default to 0]
**min_severity** | **string** | Severity level threshold for sending notifications. | [optional] [default to 'critical']
**active** | **bool** | Set to false to disable notifications | [optional] [default to true]
**all_checks** | **bool** | Trigger for all checks | [optional] [default to false]
**all_groups** | **bool** | Trigger for all check groups | [optional] [default to false]
**all_apps** | **bool** | Trigger for all monitor apps | [optional] [default to false]
**checks** | **int[]** |  | [optional]
**groups** | **int[]** |  | [optional]
**apps** | **int[]** |  | [optional]
**contacts** | [**\OpenAPI\Client\Model\ApiMonitoringAlertsAlertContacts[]**](ApiMonitoringAlertsAlertContacts.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
