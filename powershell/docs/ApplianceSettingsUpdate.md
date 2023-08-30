# ApplianceSettingsUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | **String** | Appliance URL | [optional] 
**InternalApplianceUrl** | **String** | Internal Appliance URL (PXE) | [optional] 
**CorsAllowed** | **String** | API Allowed Origins | [optional] 
**RegistrationEnabled** | **Boolean** | Registration enabled (true, false) | [optional] 
**DefaultRoleId** | **Int64** | Default tenant role ID | [optional] 
**DefaultUserRoleId** | **Int64** | Default user role ID | [optional] 
**DockerPrivilegedMode** | **Boolean** | Docker privileged mode (true, false) | [optional] 
**PasswordMinLength** | **String** | Min Password Length | [optional] 
**PasswordMinUpperCase** | **String** | Min Password Uppercase | [optional] 
**PasswordMinNumbers** | **String** | Min Password Numbers | [optional] 
**PasswordMinSymbols** | **String** | Min Password Symbols | [optional] 
**UserBrowserSessionTimeout** | **String** | User Browser Session Timeout (Minutes) | [optional] 
**UserBrowserSessionWarning** | **String** | User Browser Session Warning (Minutes) | [optional] 
**ExpirePwdDays** | **Int64** | Expire password after days. Setting to 0 disabled this feature | [optional] 
**DisableAfterAttempts** | **Int64** | Disable user after number of attempts. Set to 0 to disable this feature | [optional] 
**DisableAfterDaysInactive** | **Int64** | Disable user if inactive for specified days. Set to 0 to disable this feature | [optional] 
**WarnUserDaysBefore** | **Int64** | Send warning email number of days in advance before deactivating. Set to 0 to disable this feature | [optional] 
**SmtpMailFrom** | **String** | From email address | [optional] 
**SmtpServer** | **String** | SMTP server / host | [optional] 
**SmtpPort** | **Int64** | SMTP port | [optional] 
**SmtpSSL** | **Boolean** | Use SSL for SMTP connection | [optional] 
**SmtpTLS** | **Boolean** | Use TLS for SMTP connections | [optional] 
**SmtpUser** | **String** | SMTP username | [optional] 
**SmtpPassword** | **String** | SMTP password | [optional] 
**ProxyHost** | **String** | Proxy host | [optional] 
**ProxyPort** | **String** | Proxy port | [optional] 
**ProxyUser** | **String** | Proxy username | [optional] 
**ProxyPassword** | **String** | Proxy password | [optional] 
**ProxyDomain** | **String** | Proxy domain | [optional] 
**ProxyWorkstation** | **String** | Proxy workstation | [optional] 
**CurrencyProvider** | **String** | Currency provider | [optional] 
**CurrencyKey** | **String** | Currency provider API key | [optional] 
**EnableAllZoneTypes** | **Boolean** | Set all cloud types enabled status on, overrides enableZoneTypes and disableZoneTypes parameters | [optional] 
**EnableZoneTypes** | **Int64[]** | List of cloud type IDs to set enabled status on | [optional] 
**DisableZoneTypes** | **Int64[]** | List of cloud type IDs to set enabled status off | [optional] 
**DisableAllZoneTypes** | **Boolean** | Set all cloud types enabled status off, can be used in conjunction with enableZoneTypes | [optional] 

## Examples

- Prepare the resource
```powershell
$ApplianceSettingsUpdate = Initialize-PSOpenAPIToolsApplianceSettingsUpdate  -ApplianceUrl null `
 -InternalApplianceUrl null `
 -CorsAllowed null `
 -RegistrationEnabled null `
 -DefaultRoleId null `
 -DefaultUserRoleId null `
 -DockerPrivilegedMode null `
 -PasswordMinLength null `
 -PasswordMinUpperCase null `
 -PasswordMinNumbers null `
 -PasswordMinSymbols null `
 -UserBrowserSessionTimeout null `
 -UserBrowserSessionWarning null `
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
 -ProxyHost null `
 -ProxyPort null `
 -ProxyUser null `
 -ProxyPassword null `
 -ProxyDomain null `
 -ProxyWorkstation null `
 -CurrencyProvider null `
 -CurrencyKey null `
 -EnableAllZoneTypes null `
 -EnableZoneTypes null `
 -DisableZoneTypes null `
 -DisableAllZoneTypes null
```

- Convert the resource to JSON
```powershell
$ApplianceSettingsUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

