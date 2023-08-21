# ApiStorageVolumesStorageVolume
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name scoped to your account for the storage volume | 
**Type** | **String** | Storage Type Code or ID | 
**Config** | [**SystemCollectionsHashtable**](.md) | Configuration object with parameters that vary by &#x60;type&#x60;. | [optional] 
**StorageServer** | [**ApiStorageVolumesStorageVolumeStorageServer**](ApiStorageVolumesStorageVolumeStorageServer.md) |  | 
**StorageGroup** | [**ApiStorageVolumesStorageVolumeStorageServer**](ApiStorageVolumesStorageVolumeStorageServer.md) |  | 

## Examples

- Prepare the resource
```powershell
$ApiStorageVolumesStorageVolume = Initialize-PSOpenAPIToolsApiStorageVolumesStorageVolume  -Name null `
 -Type null `
 -Config null `
 -StorageServer null `
 -StorageGroup null
```

- Convert the resource to JSON
```powershell
$ApiStorageVolumesStorageVolume | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

