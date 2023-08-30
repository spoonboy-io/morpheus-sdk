# CheckElasticConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsHost** | **String** | Hostname or IP address of the Elasticsearch server | 
**EsPort** | **String** | Port to connect to the HTTP service | 
**CheckUser** | **String** |  | [optional] 
**TextCheckOn** | **String** |  | [optional] 
**CheckPassword** | **String** |  | [optional] 
**WebTextMatch** | **String** |  | [optional] 
**CheckPasswordHash** | **String** |  | [optional] 
**TunnelOn** | **String** |  | [optional] 
**SshHost** | **String** |  | [optional] 
**SshPort** | **Int64** |  | [optional] 
**SshUser** | **String** |  | [optional] 
**SshPassword** | **String** | Password for user, if not using key based authentication | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckElasticConfig = Initialize-PSOpenAPIToolsCheckElasticConfig  -EsHost test.example.org `
 -EsPort 9200 `
 -CheckUser null `
 -TextCheckOn null `
 -CheckPassword null `
 -WebTextMatch null `
 -CheckPasswordHash null `
 -TunnelOn null `
 -SshHost null `
 -SshPort null `
 -SshUser null `
 -SshPassword null
```

- Convert the resource to JSON
```powershell
$CheckElasticConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

