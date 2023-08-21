

# InstancesConfigAzure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resourcePoolId** | **String** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. |  [optional]
**availabilityOptions** | [**AvailabilityOptionsEnum**](#AvailabilityOptionsEnum) | Availability Options |  [optional]
**availabilitySet** | **String** | Availability Set |  [optional]
**availabilityZone** | **Long** | Availability Zone |  [optional]
**azurefloatingIp** | [**AzurefloatingIpEnum**](#AzurefloatingIpEnum) | Assign Public IP |  [optional]
**bootDiagnostics** | [**BootDiagnosticsEnum**](#BootDiagnosticsEnum) | Boot Diagnostics |  [optional]
**osGuestDiagnostics** | [**OsGuestDiagnosticsEnum**](#OsGuestDiagnosticsEnum) | OS Guest Diagnostics |  [optional]
**diagnosticsStorageAccount** | **String** | Diagnostics Storage Account |  [optional]
**createUser** | **Boolean** | Create User |  [optional]



## Enum: AvailabilityOptionsEnum

Name | Value
---- | -----
ZONE | &quot;zone&quot;
SET | &quot;set&quot;



## Enum: AzurefloatingIpEnum

Name | Value
---- | -----
ON | &quot;on&quot;
OFF | &quot;off&quot;



## Enum: BootDiagnosticsEnum

Name | Value
---- | -----
ENABLE | &quot;enable&quot;
ENABLE_CUSTOM_STORAGE | &quot;enable_custom_storage&quot;



## Enum: OsGuestDiagnosticsEnum

Name | Value
---- | -----
ON | &quot;on&quot;
OFF | &quot;off&quot;



