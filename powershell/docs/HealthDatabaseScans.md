# HealthDatabaseScans
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerReadFirst** | **String** |  | [optional] 
**HandlerReadKey** | **String** |  | [optional] 
**HandlerReadLast** | **String** |  | [optional] 
**HandlerReadNext** | **String** |  | [optional] 
**HandlerReadPrev** | **String** |  | [optional] 
**HandlerReadRnd** | **String** |  | [optional] 
**HandlerReadRndNext** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthDatabaseScans = Initialize-PSOpenAPIToolsHealthDatabaseScans  -HandlerReadFirst null `
 -HandlerReadKey null `
 -HandlerReadLast null `
 -HandlerReadNext null `
 -HandlerReadPrev null `
 -HandlerReadRnd null `
 -HandlerReadRndNext null
```

- Convert the resource to JSON
```powershell
$HealthDatabaseScans | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

