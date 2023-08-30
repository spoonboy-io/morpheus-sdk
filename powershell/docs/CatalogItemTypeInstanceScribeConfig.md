# CatalogItemTypeInstanceScribeConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePoolId** | **String** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional] 
**AvailabilityOptions** | **String** | Availability Options | [optional] 
**AvailabilitySet** | **String** | Availability Set | [optional] 
**AvailabilityZone** | **Int64** | Availability Zone | [optional] 
**AzurefloatingIp** | **String** | Assign Public IP | [optional] 
**BootDiagnostics** | **String** | Boot Diagnostics | [optional] 
**OsGuestDiagnostics** | **String** | OS Guest Diagnostics | [optional] 
**DiagnosticsStorageAccount** | **String** | Diagnostics Storage Account | [optional] 
**CreateUser** | **Boolean** | Create User | [optional] [default to $true]
**NoAgent** | **Boolean** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to $false]
**HostId** | **String** | Specific host to deploy to if so desired. | [optional] 
**SmbiosAssetTag** | **String** | Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used. | [optional] 
**NestedVirtualization** | **String** | Enable Nested Virtualization | [optional] [default to "off"]
**VmwareFolderId** | **String** | VMWare Folder External ID (as a String) or ID (as an Integer or String) | [optional] 
**GoogleZoneId** | **Int64** | External ID of the Google zone to use for instance. | [optional] 
**ExternalIpType** | **Int64** | External IP Type.  &#x60;-1&#x60; for ephemeral IP. | [optional] 
**NetworkTags** | **String** | Network Tags | [optional] 
**ServiceAccount** | **String** | Service Account | [optional] 
**AccessScope** | **String** | Access Scope | [optional] 
**IsEC2** | **String** | Amazon Cloud Type | [optional] [default to "false"]
**AvailabilityId** | **String** | Amazon Zone | [optional] 
**SecurityId** | **String** | Security Group | [optional] 
**PublicIpType** | **String** | Public IP | [optional] 
**InstanceProfile** | **String** | IAM Profile | [optional] 
**KmsKeyId** | **String** | KMS Key ID | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogItemTypeInstanceScribeConfig = Initialize-PSOpenAPIToolsCatalogItemTypeInstanceScribeConfig  -ResourcePoolId null `
 -AvailabilityOptions null `
 -AvailabilitySet null `
 -AvailabilityZone null `
 -AzurefloatingIp null `
 -BootDiagnostics null `
 -OsGuestDiagnostics null `
 -DiagnosticsStorageAccount null `
 -CreateUser null `
 -NoAgent null `
 -HostId null `
 -SmbiosAssetTag null `
 -NestedVirtualization null `
 -VmwareFolderId null `
 -GoogleZoneId null `
 -ExternalIpType -1 `
 -NetworkTags null `
 -ServiceAccount null `
 -AccessScope null `
 -IsEC2 null `
 -AvailabilityId null `
 -SecurityId null `
 -PublicIpType elasticIp `
 -InstanceProfile null `
 -KmsKeyId null
```

- Convert the resource to JSON
```powershell
$CatalogItemTypeInstanceScribeConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

