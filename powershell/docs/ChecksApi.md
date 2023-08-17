# PSOpenAPITools.PSOpenAPITools/Api.ChecksApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add-CheckApps**](ChecksApi.md#Add-CheckApps) | **POST** /api/monitoring/apps | Create a New Check App
[**Invoke-ListCheckApps**](ChecksApi.md#Invoke-ListCheckApps) | **GET** /api/monitoring/apps | List All Check Apps


<a id="Add-CheckApps"></a>
# **Add-CheckApps**
> AddCheckApps200Response Add-CheckApps<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddCheckAppsRequest] <PSCustomObject><br>

Create a New Check App

Create a new check app.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$AddCheckAppsRequestMonitorApp = Initialize-AddCheckAppsRequestMonitorApp -Name "My Check App" -Description "My cool description" -InUptime $false -Severity "info" -Active $false -Checks 0 -CheckGroups 0
$AddCheckAppsRequest = Initialize-AddCheckAppsRequest -MonitorApp $AddCheckAppsRequestMonitorApp # AddCheckAppsRequest |  (optional)

# Create a New Check App
try {
    $Result = Add-CheckApps -AddCheckAppsRequest $AddCheckAppsRequest
} catch {
    Write-Host ("Exception occurred when calling Add-CheckApps: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AddCheckAppsRequest** | [**AddCheckAppsRequest**](AddCheckAppsRequest.md)|  | [optional] 

### Return type

[**AddCheckApps200Response**](AddCheckApps200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListCheckApps"></a>
# **Invoke-ListCheckApps**
> ListCheckApps200Response Invoke-ListCheckApps<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Max] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Offset] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Sort] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Name] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Phrase] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Status] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-LastUpdated] <System.Nullable[System.DateTime]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Deleted] <System.Nullable[Boolean]><br>

List All Check Apps

Get a list of check apps.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Max = 789 # Int64 | Maximum number of records to return (optional) (default to 25)
$Offset = 789 # Int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
$Sort = "MySort" # String | Sort order, the name of the property to sort by (optional) (default to "name")
$Name = "example-%" # String | Filter by name, wildcard may be specified as %, eg. example-% (optional)
$Phrase = "MyPhrase" # String | Search phrase for partial matches on name or description (optional)
$Status = "running" # String | The instance status for filtering. (optional)
$LastUpdated = (Get-Date) # System.DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
$Deleted = $true # Boolean | If true, only deleted resources or instances in pending removal status are returned. (optional)

# List All Check Apps
try {
    $Result = Invoke-ListCheckApps -Max $Max -Offset $Offset -Sort $Sort -Name $Name -Phrase $Phrase -Status $Status -LastUpdated $LastUpdated -Deleted $Deleted
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListCheckApps: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Max** | **Int64**| Maximum number of records to return | [optional] [default to 25]
 **Offset** | **Int64**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **Sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **Name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **Phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **Status** | **String**| The instance status for filtering. | [optional] 
 **LastUpdated** | **System.DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **Deleted** | **Boolean**| If true, only deleted resources or instances in pending removal status are returned. | [optional] 

### Return type

[**ListCheckApps200Response**](ListCheckApps200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

