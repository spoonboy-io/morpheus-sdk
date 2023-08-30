# ClusterLayoutCreateMastersInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | **Int64** | Number of nodes | [optional] [default to 1]
**ContainerType** | [**AddStorageVolumesRequestStorageVolumeStorageServer**](AddStorageVolumesRequestStorageVolumeStorageServer.md) |  | 

## Examples

- Prepare the resource
```powershell
$ClusterLayoutCreateMastersInner = Initialize-PSOpenAPIToolsClusterLayoutCreateMastersInner  -NodeCount null `
 -ContainerType null
```

- Convert the resource to JSON
```powershell
$ClusterLayoutCreateMastersInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

