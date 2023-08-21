# # InstancesConfigVMWare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**no_agent** | **bool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**resource_pool_id** | **string** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional]
**host_id** | **string** | Specific host to deploy to if so desired. | [optional]
**smbios_asset_tag** | **string** | Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used. | [optional]
**nested_virtualization** | **string** | Enable Nested Virtualization | [optional] [default to 'off']
**vmware_folder_id** | **string** | VMWare Folder External ID (as a String) or ID (as an Integer or String) | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
