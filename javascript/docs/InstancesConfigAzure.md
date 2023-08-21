# MorpheusApi.InstancesConfigAzure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resourcePoolId** | **String** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional] 
**availabilityOptions** | **String** | Availability Options | [optional] 
**availabilitySet** | **String** | Availability Set | [optional] 
**availabilityZone** | **Number** | Availability Zone | [optional] 
**azurefloatingIp** | **String** | Assign Public IP | [optional] 
**bootDiagnostics** | **String** | Boot Diagnostics | [optional] 
**osGuestDiagnostics** | **String** | OS Guest Diagnostics | [optional] 
**diagnosticsStorageAccount** | **String** | Diagnostics Storage Account | [optional] 
**createUser** | **Boolean** | Create User | [optional] [default to true]



## Enum: AvailabilityOptionsEnum


* `zone` (value: `"zone"`)

* `set` (value: `"set"`)





## Enum: AzurefloatingIpEnum


* `on` (value: `"on"`)

* `off` (value: `"off"`)





## Enum: BootDiagnosticsEnum


* `enable` (value: `"enable"`)

* `enable_custom_storage` (value: `"enable_custom_storage"`)





## Enum: OsGuestDiagnosticsEnum


* `on` (value: `"on"`)

* `off` (value: `"off"`)




