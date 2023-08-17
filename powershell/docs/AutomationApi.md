# PSOpenAPITools.PSOpenAPITools/Api.AutomationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add-ExecuteSchedules**](AutomationApi.md#Add-ExecuteSchedules) | **POST** /api/execute-schedules | Creates a Execute Schedule
[**Invoke-ExecuteExecutionRequest**](AutomationApi.md#Invoke-ExecuteExecutionRequest) | **POST** /api/execution-request/execute | Executes an Execution Request
[**Get-ExecuteSchedules**](AutomationApi.md#Get-ExecuteSchedules) | **GET** /api/execute-schedules/{id} | Retrieves a Specific Execute Schedule
[**Get-ExecutionRequest**](AutomationApi.md#Get-ExecutionRequest) | **GET** /api/execution-request/{uniqueId} | Retrieves a Specific Execution Request
[**Invoke-ListExecuteSchedules**](AutomationApi.md#Invoke-ListExecuteSchedules) | **GET** /api/execute-schedules | Retrieves all Execute Schedules
[**Remove-ExecuteSchedules**](AutomationApi.md#Remove-ExecuteSchedules) | **DELETE** /api/execute-schedules/{id} | Deletes a Execute Schedule
[**Update-ExecuteSchedules**](AutomationApi.md#Update-ExecuteSchedules) | **PUT** /api/execute-schedules/{id} | Updates a Execute Schedule


<a id="Add-ExecuteSchedules"></a>
# **Add-ExecuteSchedules**
> AddExecuteSchedules200Response Add-ExecuteSchedules<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddExecuteSchedulesRequest] <PSCustomObject><br>

Creates a Execute Schedule

Creates a execute schedule. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$AddExecuteSchedulesRequestSchedule = Initialize-AddExecuteSchedulesRequestSchedule -Name "Sample Execution" -Description "MyDescription" -ScheduleType "execute" -ScheduleTimezone "MyScheduleTimezone" -Cron "MyCron" -Enabled $false
$AddExecuteSchedulesRequest = Initialize-AddExecuteSchedulesRequest -Schedule $AddExecuteSchedulesRequestSchedule # AddExecuteSchedulesRequest |  (optional)

# Creates a Execute Schedule
try {
    $Result = Add-ExecuteSchedules -AddExecuteSchedulesRequest $AddExecuteSchedulesRequest
} catch {
    Write-Host ("Exception occurred when calling Add-ExecuteSchedules: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AddExecuteSchedulesRequest** | [**AddExecuteSchedulesRequest**](AddExecuteSchedulesRequest.md)|  | [optional] 

### Return type

[**AddExecuteSchedules200Response**](AddExecuteSchedules200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ExecuteExecutionRequest"></a>
# **Invoke-ExecuteExecutionRequest**
> ExecuteExecutionRequest200Response Invoke-ExecuteExecutionRequest<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-InstanceId] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ContainerId] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ServerId] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ExecuteExecutionRequestRequest] <PSCustomObject><br>

Executes an Execution Request

Provides API interfaces for executing an arbitrary script or command on an instance, container or host. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$InstanceId = 94 # Int64 | The Instance ID for Filtering (optional)
$ContainerId = 135 # Int64 | The Container ID for Filtering (optional)
$ServerId = 97 # Int64 | The Server ID for Filtering (optional)
$ExecuteExecutionRequestRequest = Initialize-ExecuteExecutionRequestRequest -Script "uname -a" # ExecuteExecutionRequestRequest |  (optional)

# Executes an Execution Request
try {
    $Result = Invoke-ExecuteExecutionRequest -InstanceId $InstanceId -ContainerId $ContainerId -ServerId $ServerId -ExecuteExecutionRequestRequest $ExecuteExecutionRequestRequest
} catch {
    Write-Host ("Exception occurred when calling Invoke-ExecuteExecutionRequest: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **InstanceId** | **Int64**| The Instance ID for Filtering | [optional] 
 **ContainerId** | **Int64**| The Container ID for Filtering | [optional] 
 **ServerId** | **Int64**| The Server ID for Filtering | [optional] 
 **ExecuteExecutionRequestRequest** | [**ExecuteExecutionRequestRequest**](ExecuteExecutionRequestRequest.md)|  | [optional] 

### Return type

[**ExecuteExecutionRequest200Response**](ExecuteExecutionRequest200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-ExecuteSchedules"></a>
# **Get-ExecuteSchedules**
> GetExecuteSchedules200Response Get-ExecuteSchedules<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Retrieves a Specific Execute Schedule

Retrieves a specific execute schedule. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Retrieves a Specific Execute Schedule
try {
    $Result = Get-ExecuteSchedules -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-ExecuteSchedules: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetExecuteSchedules200Response**](GetExecuteSchedules200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-ExecutionRequest"></a>
# **Get-ExecutionRequest**
> GetExecutionRequest200Response Get-ExecutionRequest<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-UniqueId] <String><br>

Retrieves a Specific Execution Request

Retrieves a specific execution request. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$UniqueId = "c56f75d0-a59a-4566-92e3-4dc716c5d076" # String | The Unique ID of the execution request

# Retrieves a Specific Execution Request
try {
    $Result = Get-ExecutionRequest -UniqueId $UniqueId
} catch {
    Write-Host ("Exception occurred when calling Get-ExecutionRequest: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **UniqueId** | **String**| The Unique ID of the execution request | 

### Return type

[**GetExecutionRequest200Response**](GetExecutionRequest200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListExecuteSchedules"></a>
# **Invoke-ListExecuteSchedules**
> ListExecuteSchedules200Response Invoke-ListExecuteSchedules<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Max] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Offset] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Sort] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Direction] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Phrase] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Name] <String><br>

Retrieves all Execute Schedules

Retrieves all execute schedules. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Max = 789 # Int64 | Maximum number of records to return (optional) (default to 25)
$Offset = 789 # Int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
$Sort = "MySort" # String | Sort order, the name of the property to sort by (optional) (default to "name")
$Direction = "asc" # String | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
$Phrase = "MyPhrase" # String | Search phrase for partial matches on name or description (optional)
$Name = "example-%" # String | Filter by name, wildcard may be specified as %, eg. example-% (optional)

# Retrieves all Execute Schedules
try {
    $Result = Invoke-ListExecuteSchedules -Max $Max -Offset $Offset -Sort $Sort -Direction $Direction -Phrase $Phrase -Name $Name
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListExecuteSchedules: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Max** | **Int64**| Maximum number of records to return | [optional] [default to 25]
 **Offset** | **Int64**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **Sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **Direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &quot;asc&quot;]
 **Phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **Name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 

### Return type

[**ListExecuteSchedules200Response**](ListExecuteSchedules200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Remove-ExecuteSchedules"></a>
# **Remove-ExecuteSchedules**
> Model200Success Remove-ExecuteSchedules<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Deletes a Execute Schedule

Deletes a specified execute schedule. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Deletes a Execute Schedule
try {
    $Result = Remove-ExecuteSchedules -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Remove-ExecuteSchedules: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
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

<a id="Update-ExecuteSchedules"></a>
# **Update-ExecuteSchedules**
> AddExecuteSchedules200Response Update-ExecuteSchedules<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-UpdateExecuteSchedulesRequest] <PSCustomObject><br>

Updates a Execute Schedule

Updates a execute schedule. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$UpdateExecuteSchedulesRequestSchedule = Initialize-UpdateExecuteSchedulesRequestSchedule -Name "Sample Execution" -Description "MyDescription" -ScheduleType "execute" -ScheduleTimezone "MyScheduleTimezone" -Cron "MyCron" -Enabled $false
$UpdateExecuteSchedulesRequest = Initialize-UpdateExecuteSchedulesRequest -Schedule $UpdateExecuteSchedulesRequestSchedule # UpdateExecuteSchedulesRequest |  (optional)

# Updates a Execute Schedule
try {
    $Result = Update-ExecuteSchedules -Id $Id -UpdateExecuteSchedulesRequest $UpdateExecuteSchedulesRequest
} catch {
    Write-Host ("Exception occurred when calling Update-ExecuteSchedules: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **UpdateExecuteSchedulesRequest** | [**UpdateExecuteSchedulesRequest**](UpdateExecuteSchedulesRequest.md)|  | [optional] 

### Return type

[**AddExecuteSchedules200Response**](AddExecuteSchedules200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

