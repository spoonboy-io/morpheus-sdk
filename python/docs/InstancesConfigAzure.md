# InstancesConfigAzure


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resource_pool_id** | **str** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional] 
**availability_options** | **str** | Availability Options | [optional] 
**availability_set** | **str** | Availability Set | [optional] 
**availability_zone** | **int** | Availability Zone | [optional] 
**azurefloating_ip** | **str** | Assign Public IP | [optional] 
**boot_diagnostics** | **str** | Boot Diagnostics | [optional] 
**os_guest_diagnostics** | **str** | OS Guest Diagnostics | [optional] 
**diagnostics_storage_account** | **str** | Diagnostics Storage Account | [optional] 
**create_user** | **bool** | Create User | [optional]  if omitted the server will use the default value of True
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


