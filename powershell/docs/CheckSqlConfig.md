# CheckSqlConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbHost** | **String** | Hostname or IP address of the database | 
**DbPort** | **String** | Database Port (defaults to default port of DB type selected) | 
**DbUser** | **String** | Database username | 
**DbPassword** | **String** | Database password, (all check data is encrypted inside the database) | 
**DbPasswordHash** | **String** | Database password hash | [optional] 
**DbName** | **String** | Database name you would like to connect to | 
**DbQuery** | **String** | Query to test | 
**CheckOperator** | **String** | Can be set to &#x60;lt&#x60; (less than), &#x60;gt&#x60; (greater than), &#x60;equal&#x60; (Equal to) for comparison | [optional] 
**CheckResult** | **Decimal** |  | [optional] 
**CheckUser** | **String** |  | [optional] 
**TextCheckOn** | **String** |  | [optional] 
**CheckPassword** | **String** |  | [optional] 
**WebTextMatch** | **String** |  | [optional] 
**CheckPasswordHash** | **String** |  | [optional] 
**TunnelOn** | **String** | Set to on to turn on tunneling | [optional] [default to "off"]
**SshHost** | **String** | Hostname or IP address of the proxy host | [optional] 
**SshPort** | **Int64** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**SshUser** | **String** | SSH user on the proxy host to login as | [optional] 
**SshPassword** | **String** | Password for user, if not using key based authentication | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckSqlConfig = Initialize-PSOpenAPIToolsCheckSqlConfig  -DbHost db.example.org `
 -DbPort 3306 `
 -DbUser basicUser `
 -DbPassword basicPassword `
 -DbPasswordHash 61236d68c76405c9c7b40838a9c5d120c76e4b222222222221943c0f340f10 `
 -DbName testDb `
 -DbQuery select 1 `
 -CheckOperator null `
 -CheckResult 3 `
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
$CheckSqlConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

