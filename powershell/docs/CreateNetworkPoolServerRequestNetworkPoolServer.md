# CreateNetworkPoolServerRequestNetworkPoolServer
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** | Type Code (SolarWinds) | 
**Name** | **String** | Name | 
**Enabled** | **Boolean** | Can be used to enable / disable the network pool server. | [optional] [default to $true]
**ServiceUrl** | **String** | URL | 
**ServiceUsername** | **String** | Username | [optional] 
**ServicePassword** | **String** | Password | [optional] 
**ServiceThrottleRate** | **Int64** | Throttle Rate | [optional] [default to 0]
**IgnoreSsl** | **Boolean** | Disable SSL SNI Verification | [optional] [default to $true]
**NetworkFilter** | **String** | Network Filter | [optional] 
**Config** | [**NetworkPoolServerCreateBluecatConfig**](NetworkPoolServerCreateBluecatConfig.md) |  | 
**Credential** | [**NetworkPoolServerCreateBluecatCredential**](NetworkPoolServerCreateBluecatCredential.md) |  | [optional] 
**ZoneFilter** | **String** | Zone Filter | [optional] 
**TenantMatch** | **String** | Tenant Match | [optional] 
**ServiceMode** | **String** | IP Mode | [optional] [default to "static"]

## Examples

- Prepare the resource
```powershell
$CreateNetworkPoolServerRequestNetworkPoolServer = Initialize-PSOpenAPIToolsCreateNetworkPoolServerRequestNetworkPoolServer  -Type null `
 -Name null `
 -Enabled null `
 -ServiceUrl https://solarwinds-server/wapi/v2.9 `
 -ServiceUsername null `
 -ServicePassword null `
 -ServiceThrottleRate null `
 -IgnoreSsl null `
 -NetworkFilter null `
 -Config null `
 -Credential null `
 -ZoneFilter null `
 -TenantMatch null `
 -ServiceMode static
```

- Convert the resource to JSON
```powershell
$CreateNetworkPoolServerRequestNetworkPoolServer | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

