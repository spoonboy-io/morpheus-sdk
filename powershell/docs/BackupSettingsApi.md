# PSOpenAPITools.PSOpenAPITools/Api.BackupSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Invoke-ListBackupSettings**](BackupSettingsApi.md#Invoke-ListBackupSettings) | **GET** /api/backup-settings | Get Backup Settings
[**Update-BackupSettings**](BackupSettingsApi.md#Update-BackupSettings) | **PUT** /api/backup-settings | Update Backup Settings


<a id="Invoke-ListBackupSettings"></a>
# **Invoke-ListBackupSettings**
> ListBackupSettings200Response Invoke-ListBackupSettings<br>

Get Backup Settings

This endpoint retrieves backup settings.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration


# Get Backup Settings
try {
    $Result = Invoke-ListBackupSettings
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListBackupSettings: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**ListBackupSettings200Response**](ListBackupSettings200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Update-BackupSettings"></a>
# **Update-BackupSettings**
> Model200Success Update-BackupSettings<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-UpdateBackupSettingsRequest] <PSCustomObject><br>

Update Backup Settings

Update Backup Settings

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$BackupSettingsUpdateDefaultSchedule = Initialize-BackupSettingsUpdateDefaultSchedule -Id 0
$BackupSettingsUpdateDefaultStorageBucket = Initialize-BackupSettingsUpdateDefaultStorageBucket -Id 0
$BackupSettingsUpdate = Initialize-BackupSettingsUpdate -BackupsEnabled $false -RetentionCount 0 -CreateBackups $false -BackupAppliance $false -UpdateExisting $false -DefaultSchedule $BackupSettingsUpdateDefaultSchedule -ClearDefaultSchedule $false -DefaultStorageBucket $BackupSettingsUpdateDefaultStorageBucket -ClearDefaultStorageBucket $false

$UpdateBackupSettingsRequest = Initialize-UpdateBackupSettingsRequest -BackupSettings $BackupSettingsUpdate # UpdateBackupSettingsRequest |  (optional)

# Update Backup Settings
try {
    $Result = Update-BackupSettings -UpdateBackupSettingsRequest $UpdateBackupSettingsRequest
} catch {
    Write-Host ("Exception occurred when calling Update-BackupSettings: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **UpdateBackupSettingsRequest** | [**UpdateBackupSettingsRequest**](UpdateBackupSettingsRequest.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

