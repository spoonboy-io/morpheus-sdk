# BlueprintMorpheusCreateSuccessConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **String** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Name** | **String** | A name for the blueprint | [optional] 
**Type** | **String** | Blueprint Type | [optional] 
**Tiers** | [**SystemCollectionsHashtable**](.md) | Tier definitions - Create in UI to view a baseline for object | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintMorpheusCreateSuccessConfig = Initialize-PSOpenAPIToolsBlueprintMorpheusCreateSuccessConfig  -Image null `
 -Name null `
 -Type null `
 -Tiers null
```

- Convert the resource to JSON
```powershell
$BlueprintMorpheusCreateSuccessConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

