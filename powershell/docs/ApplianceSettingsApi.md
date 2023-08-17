# PSOpenAPITools.PSOpenAPITools/Api.ApplianceSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Invoke-ListApplianceSettings**](ApplianceSettingsApi.md#Invoke-ListApplianceSettings) | **GET** /api/appliance-settings | Get Appliance Settings
[**Invoke-Reindex**](ApplianceSettingsApi.md#Invoke-Reindex) | **POST** /api/appliance-settings/reindex | Reindex Search
[**Set-ApplianceSettingsMaintenanceMode**](ApplianceSettingsApi.md#Set-ApplianceSettingsMaintenanceMode) | **POST** /api/appliance-settings/maintenance | Toggle Maintenance Mode
[**Update-ApplianceSettings**](ApplianceSettingsApi.md#Update-ApplianceSettings) | **PUT** /api/appliance-settings | Update Appliance Settings


<a id="Invoke-ListApplianceSettings"></a>
# **Invoke-ListApplianceSettings**
> ListApplianceSettings200Response Invoke-ListApplianceSettings<br>

Get Appliance Settings

This endpoint retrieves appliance settings.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration


# Get Appliance Settings
try {
    $Result = Invoke-ListApplianceSettings
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListApplianceSettings: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**ListApplianceSettings200Response**](ListApplianceSettings200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-Reindex"></a>
# **Invoke-Reindex**
> Model200Success Invoke-Reindex<br>

Reindex Search

Reindex all searchable data

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration


# Reindex Search
try {
    $Result = Invoke-Reindex
} catch {
    Write-Host ("Exception occurred when calling Invoke-Reindex: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Model200Success**](Model200Success.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Set-ApplianceSettingsMaintenanceMode"></a>
# **Set-ApplianceSettingsMaintenanceMode**
> Model200Success Set-ApplianceSettingsMaintenanceMode<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Enabled] <System.Nullable[Boolean]><br>

Toggle Maintenance Mode

This endpoint allows toggling the appliance maintenance mode.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Enabled = $true # Boolean | Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. (optional)

# Toggle Maintenance Mode
try {
    $Result = Set-ApplianceSettingsMaintenanceMode -Enabled $Enabled
} catch {
    Write-Host ("Exception occurred when calling Set-ApplianceSettingsMaintenanceMode: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Enabled** | **Boolean**| Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. | [optional] 

### Return type

[**Model200Success**](Model200Success.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Update-ApplianceSettings"></a>
# **Update-ApplianceSettings**
> Model200Success Update-ApplianceSettings<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-UpdateApplianceSettingsRequest] <PSCustomObject><br>

Update Appliance Settings

Update Appliance Settings

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$ApplianceSettingsUpdate = Initialize-ApplianceSettingsUpdate -ApplianceUrl "MyApplianceUrl" -InternalApplianceUrl "MyInternalApplianceUrl" -CorsAllowed "MyCorsAllowed" -RegistrationEnabled $false -DefaultRoleId 0 -DefaultUserRoleId 0 -DockerPrivilegedMode $false -PasswordMinLength "MyPasswordMinLength" -PasswordMinUpperCase "MyPasswordMinUpperCase" -PasswordMinNumbers "MyPasswordMinNumbers" -PasswordMinSymbols "MyPasswordMinSymbols" -UserBrowserSessionTimeout "MyUserBrowserSessionTimeout" -UserBrowserSessionWarning "MyUserBrowserSessionWarning" -ExpirePwdDays 0 -DisableAfterAttempts 0 -DisableAfterDaysInactive 0 -WarnUserDaysBefore 0 -SmtpMailFrom "MySmtpMailFrom" -SmtpServer "MySmtpServer" -SmtpPort 0 -SmtpSSL $false -SmtpTLS $false -SmtpUser "MySmtpUser" -SmtpPassword "MySmtpPassword" -ProxyHost "MyProxyHost" -ProxyPort "MyProxyPort" -ProxyUser "MyProxyUser" -ProxyPassword "MyProxyPassword" -ProxyDomain "MyProxyDomain" -ProxyWorkstation "MyProxyWorkstation" -CurrencyProvider "MyCurrencyProvider" -CurrencyKey "MyCurrencyKey" -EnableAllZoneTypes $false -EnableZoneTypes 0 -DisableZoneTypes 0 -DisableAllZoneTypes $false
$UpdateApplianceSettingsRequest = Initialize-UpdateApplianceSettingsRequest -ApplianceSettings $ApplianceSettingsUpdate # UpdateApplianceSettingsRequest |  (optional)

# Update Appliance Settings
try {
    $Result = Update-ApplianceSettings -UpdateApplianceSettingsRequest $UpdateApplianceSettingsRequest
} catch {
    Write-Host ("Exception occurred when calling Update-ApplianceSettings: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **UpdateApplianceSettingsRequest** | [**UpdateApplianceSettingsRequest**](UpdateApplianceSettingsRequest.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

