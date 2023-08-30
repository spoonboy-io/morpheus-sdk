# ClusterApiConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceUrl** | **String** |  | [optional] 
**ServiceHost** | **String** |  | [optional] 
**ServicePath** | **String** |  | [optional] 
**ServiceHostname** | **String** |  | [optional] 
**ServicePort** | **Int64** |  | [optional] 
**ServiceUsername** | **String** |  | [optional] 
**ServicePassword** | **String** |  | [optional] 
**ServicePasswordHash** | **String** |  | [optional] 
**ServiceToken** | **String** | API Token | [optional] 
**ServiceAccess** | **String** | Kube Config | [optional] 
**ServiceCert** | **String** |  | [optional] 
**ServiceVersion** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterApiConfig = Initialize-PSOpenAPIToolsClusterApiConfig  -ServiceUrl null `
 -ServiceHost null `
 -ServicePath null `
 -ServiceHostname null `
 -ServicePort null `
 -ServiceUsername null `
 -ServicePassword null `
 -ServicePasswordHash null `
 -ServiceToken null `
 -ServiceAccess null `
 -ServiceCert null `
 -ServiceVersion null
```

- Convert the resource to JSON
```powershell
$ClusterApiConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

