# BootScript
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Account** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**FileName** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Content** | **String** |  | [optional] 
**CreatedBy** | [**ArchiveBucketCreatedBy**](ArchiveBucketCreatedBy.md) |  | [optional] 
**Visibility** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BootScript = Initialize-PSOpenAPIToolsBootScript  -Id null `
 -Account null `
 -FileName null `
 -Description null `
 -Content null `
 -CreatedBy null `
 -Visibility null
```

- Convert the resource to JSON
```powershell
$BootScript | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

