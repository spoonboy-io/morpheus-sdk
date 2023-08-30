# AddStorageVolumesRequestStorageVolume
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name scoped to your account for the storage volume | 
**Type** | **String** | Storage Type Code or ID | 
**Config** | [**SystemCollectionsHashtable**](.md) | Configuration object with parameters that vary by &#x60;type&#x60;. | [optional] 
**StorageServer** | [**AddStorageVolumesRequestStorageVolumeStorageServer**](AddStorageVolumesRequestStorageVolumeStorageServer.md) |  | 
**StorageGroup** | [**AddStorageVolumesRequestStorageVolumeStorageServer**](AddStorageVolumesRequestStorageVolumeStorageServer.md) |  | 

## Examples

- Prepare the resource
```powershell
$AddStorageVolumesRequestStorageVolume = Initialize-PSOpenAPIToolsAddStorageVolumesRequestStorageVolume  -Name null `
 -Type null `
 -Config null `
 -StorageServer null `
 -StorageGroup null
```

- Convert the resource to JSON
```powershell
$AddStorageVolumesRequestStorageVolume | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

