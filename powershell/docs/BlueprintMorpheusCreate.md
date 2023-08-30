# BlueprintMorpheusCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the blueprint | 
**Type** | **String** | Blueprint Type | 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Tiers** | [**SystemCollectionsHashtable**](.md) | Tier definitions - Create in UI to view a baseline for object | 

## Examples

- Prepare the resource
```powershell
$BlueprintMorpheusCreate = Initialize-PSOpenAPIToolsBlueprintMorpheusCreate  -Name null `
 -Type null `
 -Labels null `
 -Tiers null
```

- Convert the resource to JSON
```powershell
$BlueprintMorpheusCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

