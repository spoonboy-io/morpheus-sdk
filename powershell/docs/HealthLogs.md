# HealthLogs
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeCode** | **String** |  | [optional] 
**Ts** | **System.DateTime** |  | [optional] 
**Level** | **String** |  | [optional] 
**SourceType** | **String** |  | [optional] 
**Message** | **String** |  | [optional] 
**Hostname** | **String** |  | [optional] 
**Title** | **String** |  | [optional] 
**LogSignature** | **String** |  | [optional] 
**ObjectId** | **String** |  | [optional] 
**Seq** | **Int64** |  | [optional] 
**Id** | **String** |  | [optional] 
**SignatureVerified** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthLogs = Initialize-PSOpenAPIToolsHealthLogs  -TypeCode null `
 -Ts null `
 -Level null `
 -SourceType null `
 -Message null `
 -Hostname null `
 -Title null `
 -LogSignature null `
 -ObjectId null `
 -Seq null `
 -Id null `
 -SignatureVerified null
```

- Convert the resource to JSON
```powershell
$HealthLogs | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

