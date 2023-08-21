# CatalogCartStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **Decimal** |  | [optional] 
**Currency** | **String** |  | [optional] 
**Unit** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogCartStats = Initialize-PSOpenAPIToolsCatalogCartStats  -Price null `
 -Currency null `
 -Unit null
```

- Convert the resource to JSON
```powershell
$CatalogCartStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

