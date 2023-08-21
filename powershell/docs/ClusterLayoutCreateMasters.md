# ClusterLayoutCreateMasters
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | **Int64** | Number of nodes | [optional] [default to 1]
**ContainerType** | [**ApiStorageVolumesStorageVolumeStorageServer**](ApiStorageVolumesStorageVolumeStorageServer.md) |  | 

## Examples

- Prepare the resource
```powershell
$ClusterLayoutCreateMasters = Initialize-PSOpenAPIToolsClusterLayoutCreateMasters  -NodeCount null `
 -ContainerType null
```

- Convert the resource to JSON
```powershell
$ClusterLayoutCreateMasters | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

