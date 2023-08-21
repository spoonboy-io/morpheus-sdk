# ContainerPort
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Index** | **Int64** |  | [optional] 
**External** | **Int64** |  | [optional] 
**Internal** | **Int64** |  | [optional] 
**DisplayName** | **String** |  | [optional] 
**PrimaryPort** | **Boolean** |  | [optional] 
**Export** | **Boolean** |  | [optional] 
**Visible** | **Boolean** |  | [optional] 
**ExportName** | **String** |  | [optional] 
**LoadBalanceProtocol** | **String** |  | [optional] 
**LoadBalance** | **Boolean** |  | [optional] 
**Protocol** | **String** |  | [optional] 
**Link** | **Boolean** |  | [optional] 
**ExternalIp** | **String** |  | [optional] 
**InternalIp** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ContainerPort = Initialize-PSOpenAPIToolsContainerPort  -Id null `
 -Index null `
 -External null `
 -Internal null `
 -DisplayName null `
 -PrimaryPort null `
 -Export null `
 -Visible null `
 -ExportName null `
 -LoadBalanceProtocol null `
 -LoadBalance null `
 -Protocol null `
 -Link null `
 -ExternalIp null `
 -InternalIp null
```

- Convert the resource to JSON
```powershell
$ContainerPort | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

