# MorpheusApi.InstancesConfigVMWare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**noAgent** | **Boolean** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**resourcePoolId** | **String** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional] 
**hostId** | **String** | Specific host to deploy to if so desired. | [optional] 
**smbiosAssetTag** | **String** | Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used. | [optional] 
**nestedVirtualization** | **String** | Enable Nested Virtualization | [optional] [default to &#39;off&#39;]
**vmwareFolderId** | **String** | VMWare Folder External ID (as a String) or ID (as an Integer or String) | [optional] 



## Enum: NestedVirtualizationEnum


* `on` (value: `"on"`)

* `off` (value: `"off"`)




