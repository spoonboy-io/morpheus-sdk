# BlueprintHelmCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the blueprint | 
**Image** | **String** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **String** | Blueprint Type | 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Helm** | [**BlueprintHelmCreateHelm**](BlueprintHelmCreateHelm.md) |  | 

## Examples

- Prepare the resource
```powershell
$BlueprintHelmCreate = Initialize-PSOpenAPIToolsBlueprintHelmCreate  -Name null `
 -Image null `
 -Type null `
 -Labels null `
 -Helm null
```

- Convert the resource to JSON
```powershell
$BlueprintHelmCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

