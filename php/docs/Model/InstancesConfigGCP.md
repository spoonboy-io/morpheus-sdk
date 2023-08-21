# # InstancesConfigGCP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**no_agent** | **bool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**google_zone_id** | **int** | External ID of the Google zone to use for instance. | [optional]
**external_ip_type** | **int** | External IP Type.  &#x60;-1&#x60; for ephemeral IP. | [optional]
**network_tags** | **string** | Network Tags | [optional]
**service_account** | **string** | Service Account | [optional]
**access_scope** | **string** | Access Scope | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
