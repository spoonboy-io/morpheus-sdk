# CheckConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebMethod** | **String** | HTTP method to use for testing | 
**WebUrl** | **String** | Web URL you wish to use to run a check on | 
**IgnoreSSL** | **Boolean** | Ignore SSL Errors | [optional] [default to $false]
**CheckUser** | **String** |  | [optional] 
**CheckPassword** | **String** |  | [optional] 
**TextCheckOn** | **String** |  | [optional] 
**WebTextMatch** | **String** |  | [optional] 
**TunnelOn** | **String** | Set to on to turn on tunneling | [optional] [default to "off"]
**SshHost** | **String** | Hostname or IP address of the proxy host | [optional] 
**SshPort** | **Int64** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**SshUser** | **String** | SSH user on the proxy host to login as | [optional] 
**SshPassword** | **String** | Password for user, if not using key based authentication | [optional] 
**DbHost** | **String** | Hostname or IP address of the database | 
**DbPort** | **String** | Database Port (defaults to default port of DB type selected) | 
**DbUser** | **String** | Database username | 
**DbPassword** | **String** | Database password, (all check data is encrypted inside the database) | 
**DbPasswordHash** | **String** | Database password hash | [optional] 
**DbName** | **String** | Database name you would like to connect to | 
**DbQuery** | **String** | Query to test | 
**CheckOperator** | **String** | Can be set to &#x60;lt&#x60; (less than), &#x60;gt&#x60; (greater than), &#x60;equal&#x60; (Equal to) for comparison | [optional] 
**CheckResult** | **Decimal** |  | [optional] 
**CheckPasswordHash** | **String** |  | [optional] 
**EsHost** | **String** | Hostname or IP address of the Elasticsearch server | 
**EsPort** | **String** | Port to connect to the HTTP service | 
**VarHost** | **String** | Hostname or IP address of the socket server | 
**Port** | **String** | TCP port | 
**Send** | **String** | Connection string you might want to send to the service | 
**ResponseMatch** | **String** | Response from the service to match against | 
**ContainerName** | **String** |  | 
**ExternalId** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckConfig = Initialize-PSOpenAPIToolsCheckConfig  -WebMethod null `
 -WebUrl https://google.com `
 -IgnoreSSL true `
 -CheckUser null `
 -CheckPassword null `
 -TextCheckOn null `
 -WebTextMatch null `
 -TunnelOn null `
 -SshHost null `
 -SshPort null `
 -SshUser null `
 -SshPassword null `
 -DbHost db.example.org `
 -DbPort 3306 `
 -DbUser basicUser `
 -DbPassword basicPassword `
 -DbPasswordHash 61236d68c76405c9c7b40838a9c5d120c76e4b222222222221943c0f340f10 `
 -DbName testDb `
 -DbQuery select 1 `
 -CheckOperator null `
 -CheckResult 3 `
 -CheckPasswordHash null `
 -EsHost test.example.org `
 -EsPort 9200 `
 -VarHost test.example.org `
 -Port 3306 `
 -Send blah `
 -ResponseMatch OK `
 -ContainerName null `
 -ExternalId null
```

- Convert the resource to JSON
```powershell
$CheckConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

