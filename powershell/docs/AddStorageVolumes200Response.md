# AddStorageVolumes200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageVolume** | [**StorageVolume**](StorageVolume.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddStorageVolumes200Response = Initialize-PSOpenAPIToolsAddStorageVolumes200Response  -StorageVolume null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddStorageVolumes200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

