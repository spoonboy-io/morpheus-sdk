# Cypher
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int32** |  | [optional] 
**ItemKey** | **String** |  | [optional] 
**LeaseTimeout** | **Int64** |  | [optional] 
**ExpireDate** | **System.DateTime** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**LastAccessed** | **System.DateTime** |  | [optional] 
**CreatedBy** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Cypher = Initialize-PSOpenAPIToolsCypher  -Id null `
 -ItemKey null `
 -LeaseTimeout null `
 -ExpireDate null `
 -DateCreated null `
 -LastUpdated null `
 -LastAccessed null `
 -CreatedBy null
```

- Convert the resource to JSON
```powershell
$Cypher | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

