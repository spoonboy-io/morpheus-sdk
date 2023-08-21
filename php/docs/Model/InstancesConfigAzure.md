# # InstancesConfigAzure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resource_pool_id** | **string** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional]
**availability_options** | **string** | Availability Options | [optional]
**availability_set** | **string** | Availability Set | [optional]
**availability_zone** | **int** | Availability Zone | [optional]
**azurefloating_ip** | **string** | Assign Public IP | [optional]
**boot_diagnostics** | **string** | Boot Diagnostics | [optional]
**os_guest_diagnostics** | **string** | OS Guest Diagnostics | [optional]
**diagnostics_storage_account** | **string** | Diagnostics Storage Account | [optional]
**create_user** | **bool** | Create User | [optional] [default to true]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
