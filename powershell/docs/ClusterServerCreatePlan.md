# ClusterServerCreatePlan
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | The id for the memory and storage option pre-configured within Morpheus. | [optional] 
**Code** | **String** | The code for the memory and storage option pre-configured within Morpheus. | [optional] 
**Options** | [**SystemCollectionsHashtable**](.md) | Map of custom options depending on selected service plan . An example would be &#x60;maxMemory&#x60;, or &#x60;maxCores&#x60;. | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterServerCreatePlan = Initialize-PSOpenAPIToolsClusterServerCreatePlan  -Id null `
 -Code null `
 -Options null
```

- Convert the resource to JSON
```powershell
$ClusterServerCreatePlan | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

