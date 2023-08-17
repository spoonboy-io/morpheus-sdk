# PSOpenAPITools.PSOpenAPITools/Api.AlertsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add-Alerts**](AlertsApi.md#Add-Alerts) | **POST** /api/monitoring/alerts | Create a New Alert
[**Invoke-DeleteAlerts**](AlertsApi.md#Invoke-DeleteAlerts) | **DELETE** /api/monitoring/alerts/{id} | Delete a Specific Alert
[**Get-Alerts**](AlertsApi.md#Get-Alerts) | **GET** /api/monitoring/alerts/{id} | Get a Specific Alert
[**Invoke-ListAlerts**](AlertsApi.md#Invoke-ListAlerts) | **GET** /api/monitoring/alerts | List All Alerts
[**Update-Alerts**](AlertsApi.md#Update-Alerts) | **PUT** /api/monitoring/alerts/{id} | Update Alert


<a id="Add-Alerts"></a>
# **Add-Alerts**
> AddAlerts200Response Add-Alerts<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddAlertsRequest] <PSCustomObject><br>

Create a New Alert

Create a new monitoring alert.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$AddAlertsRequestAlertContactsInner = Initialize-AddAlertsRequestAlertContactsInner -Id 0 -Name "MyName" -Method "MyMethod" -Notify $false -Close $false
$AddAlertsRequestAlert = Initialize-AddAlertsRequestAlert -Name "My Alert" -MinDuration 0 -MinSeverity "info" -Active $false -AllChecks $false -AllGroups $false -AllApps $false -Checks 0 -Groups 0 -Apps 0 -Contacts $AddAlertsRequestAlertContactsInner

$AddAlertsRequest = Initialize-AddAlertsRequest -Alert $AddAlertsRequestAlert # AddAlertsRequest |  (optional)

# Create a New Alert
try {
    $Result = Add-Alerts -AddAlertsRequest $AddAlertsRequest
} catch {
    Write-Host ("Exception occurred when calling Add-Alerts: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AddAlertsRequest** | [**AddAlertsRequest**](AddAlertsRequest.md)|  | [optional] 

### Return type

[**AddAlerts200Response**](AddAlerts200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-DeleteAlerts"></a>
# **Invoke-DeleteAlerts**
> Model200Success Invoke-DeleteAlerts<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Delete a Specific Alert

Delete an existing monitoring alert.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Delete a Specific Alert
try {
    $Result = Invoke-DeleteAlerts -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Invoke-DeleteAlerts: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-Alerts"></a>
# **Get-Alerts**
> GetAlerts200Response Get-Alerts<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Get a Specific Alert

Get details about a specific monitoring alert.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Get a Specific Alert
try {
    $Result = Get-Alerts -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-Alerts: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetAlerts200Response**](GetAlerts200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListAlerts"></a>
# **Invoke-ListAlerts**
> ListAlerts200Response Invoke-ListAlerts<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Max] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Offset] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-LastUpdated] <System.Nullable[System.DateTime]><br>

List All Alerts

Get a list of monitoring alerts.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Max = 789 # Int64 | Maximum number of records to return (optional) (default to 25)
$Offset = 789 # Int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
$LastUpdated = (Get-Date) # System.DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

# List All Alerts
try {
    $Result = Invoke-ListAlerts -Max $Max -Offset $Offset -LastUpdated $LastUpdated
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListAlerts: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Max** | **Int64**| Maximum number of records to return | [optional] [default to 25]
 **Offset** | **Int64**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **LastUpdated** | **System.DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 

### Return type

[**ListAlerts200Response**](ListAlerts200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Update-Alerts"></a>
# **Update-Alerts**
> UpdateAlerts200Response Update-Alerts<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-UpdateAlertsRequest] <PSCustomObject><br>

Update Alert

Update an existing monitoring alert.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$AddAlertsRequestAlertContactsInner = Initialize-AddAlertsRequestAlertContactsInner -Id 0 -Name "MyName" -Method "MyMethod" -Notify $false -Close $false
$UpdateAlertsRequestAlert = Initialize-UpdateAlertsRequestAlert -Name "My Alert" -MinDuration 0 -MinSeverity "info" -Active $false -AllChecks $false -AllGroups $false -AllApps $false -Checks 0 -Groups 0 -Apps 0 -Contacts $AddAlertsRequestAlertContactsInner

$UpdateAlertsRequest = Initialize-UpdateAlertsRequest -Alert $UpdateAlertsRequestAlert # UpdateAlertsRequest |  (optional)

# Update Alert
try {
    $Result = Update-Alerts -Id $Id -UpdateAlertsRequest $UpdateAlertsRequest
} catch {
    Write-Host ("Exception occurred when calling Update-Alerts: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **UpdateAlertsRequest** | [**UpdateAlertsRequest**](UpdateAlertsRequest.md)|  | [optional] 

### Return type

[**UpdateAlerts200Response**](UpdateAlerts200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

