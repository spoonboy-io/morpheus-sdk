# InstancesConfigVMWare


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**no_agent** | **bool, none_type** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional]  if omitted the server will use the default value of False
**resource_pool_id** | **str** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional] 
**host_id** | **str** | Specific host to deploy to if so desired. | [optional] 
**smbios_asset_tag** | **str** | Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used. | [optional] 
**nested_virtualization** | **str** | Enable Nested Virtualization | [optional]  if omitted the server will use the default value of "off"
**vmware_folder_id** | **str** | VMWare Folder External ID (as a String) or ID (as an Integer or String) | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


