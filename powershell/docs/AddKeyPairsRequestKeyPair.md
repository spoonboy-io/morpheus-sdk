# AddKeyPairsRequestKeyPair
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** |  | 
**PublicKey** | **String** |  | 
**PrivateKey** | **String** |  | [optional] 
**Passphrase** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddKeyPairsRequestKeyPair = Initialize-PSOpenAPIToolsAddKeyPairsRequestKeyPair  -Name null `
 -PublicKey null `
 -PrivateKey null `
 -Passphrase null
```

- Convert the resource to JSON
```powershell
$AddKeyPairsRequestKeyPair | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

