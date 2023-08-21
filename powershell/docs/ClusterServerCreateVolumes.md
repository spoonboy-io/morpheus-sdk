# ClusterServerCreateVolumes
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | The id for the LV configuration being created | [optional] [default to -1]
**RootVolume** | **Boolean** | If set to false then a non-root LV will be created | [optional] [default to $true]
**Name** | **String** | Name/type of the LV being created | [default to "root"]
**Size** | **Int64** | Size of the LV to be created in GBs  Default is from the service plan  | [optional] 
**SizeId** | **String** | Can be used to select pre-existing LV choices from Morpheus | [optional] 
**StorageType** | **Int64** | Identifier for LV type | [optional] 
**DatastoreId** | **String** | The ID of the specific datastore. Auto selection can be specified as auto or &#x60;autoCluster&#x60; (for clusters). | 

## Examples

- Prepare the resource
```powershell
$ClusterServerCreateVolumes = Initialize-PSOpenAPIToolsClusterServerCreateVolumes  -Id null `
 -RootVolume null `
 -Name null `
 -Size null `
 -SizeId null `
 -StorageType null `
 -DatastoreId null
```

- Convert the resource to JSON
```powershell
$ClusterServerCreateVolumes | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

