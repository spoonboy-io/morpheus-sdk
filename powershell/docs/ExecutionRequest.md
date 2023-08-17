# ExecutionRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**UniqueId** | **String** |  | [optional] 
**ContainerId** | **String** |  | [optional] 
**ServerId** | **String** |  | [optional] 
**InstanceId** | **Int64** |  | [optional] 
**ResourceId** | **String** |  | [optional] 
**AppId** | **String** |  | [optional] 
**StdOut** | **String** |  | [optional] 
**StdErr** | **String** |  | [optional] 
**ExitCode** | **Int64** |  | [optional] 
**Status** | **String** |  | [optional] 
**ExpiresAt** | **System.DateTime** |  | [optional] 
**CreatedById** | **Int64** |  | [optional] 
**StatusMessage** | **String** |  | [optional] 
**ErrorMessage** | **String** |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**RawData** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ExecutionRequest = Initialize-PSOpenAPIToolsExecutionRequest  -Id null `
 -UniqueId null `
 -ContainerId null `
 -ServerId null `
 -InstanceId null `
 -ResourceId null `
 -AppId null `
 -StdOut null `
 -StdErr null `
 -ExitCode null `
 -Status null `
 -ExpiresAt null `
 -CreatedById null `
 -StatusMessage null `
 -ErrorMessage null `
 -Config null `
 -RawData null
```

- Convert the resource to JSON
```powershell
$ExecutionRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

