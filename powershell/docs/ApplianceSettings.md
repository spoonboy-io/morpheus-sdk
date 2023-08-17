# ApplianceSettings
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | **String** |  | [optional] 
**InternalApplianceUrl** | **String** |  | [optional] 
**CorsAllowed** | **String** |  | [optional] 
**RegistrationEnabled** | **Boolean** |  | [optional] 
**DefaultRoleId** | **String** |  | [optional] 
**DefaultUserRoleId** | **String** |  | [optional] 
**DockerPrivilegedMode** | **Boolean** |  | [optional] 
**ExpirePwdDays** | **String** |  | [optional] 
**DisableAfterAttempts** | **String** |  | [optional] 
**DisableAfterDaysInactive** | **String** |  | [optional] 
**WarnUserDaysBefore** | **String** |  | [optional] 
**SmtpMailFrom** | **String** |  | [optional] 
**SmtpServer** | **String** |  | [optional] 
**SmtpPort** | **String** |  | [optional] 
**SmtpSSL** | **Boolean** |  | [optional] 
**SmtpTLS** | **Boolean** |  | [optional] 
**SmtpUser** | **String** |  | [optional] 
**SmtpPassword** | **String** |  | [optional] 
**SmtpPasswordHash** | **String** |  | [optional] 
**ProxyHost** | **String** |  | [optional] 
**ProxyPort** | **String** |  | [optional] 
**ProxyUser** | **String** |  | [optional] 
**ProxyPassword** | **String** |  | [optional] 
**ProxyPasswordHash** | **String** |  | [optional] 
**ProxyDomain** | **String** |  | [optional] 
**ProxyWorkstation** | **String** |  | [optional] 
**CurrencyProvider** | **String** |  | [optional] 
**CurrencyKey** | **String** |  | [optional] 
**EnabledZoneTypes** | [**ApplianceSettingsEnabledZoneTypesInner[]**](ApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**StatsRetainmentPeriod** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApplianceSettings = Initialize-PSOpenAPIToolsApplianceSettings  -ApplianceUrl null `
 -InternalApplianceUrl null `
 -CorsAllowed null `
 -RegistrationEnabled null `
 -DefaultRoleId null `
 -DefaultUserRoleId null `
 -DockerPrivilegedMode null `
 -ExpirePwdDays null `
 -DisableAfterAttempts null `
 -DisableAfterDaysInactive null `
 -WarnUserDaysBefore null `
 -SmtpMailFrom null `
 -SmtpServer null `
 -SmtpPort null `
 -SmtpSSL null `
 -SmtpTLS null `
 -SmtpUser null `
 -SmtpPassword null `
 -SmtpPasswordHash null `
 -ProxyHost null `
 -ProxyPort null `
 -ProxyUser null `
 -ProxyPassword null `
 -ProxyPasswordHash null `
 -ProxyDomain null `
 -ProxyWorkstation null `
 -CurrencyProvider null `
 -CurrencyKey null `
 -EnabledZoneTypes null `
 -StatsRetainmentPeriod null
```

- Convert the resource to JSON
```powershell
$ApplianceSettings | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

