# PSOpenAPITools.PSOpenAPITools/Api.AppsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add-AppInstance**](AppsApi.md#Add-AppInstance) | **POST** /api/apps/{id}/add-instance | Add Existing Instance to App
[**Add-AppUndoDelete**](AppsApi.md#Add-AppUndoDelete) | **PUT** /api/apps/{id}/cancel-removal | Undo Delete of an App
[**Add-Apps**](AppsApi.md#Add-Apps) | **POST** /api/apps | Create an App
[**Invoke-ApplyAppState**](AppsApi.md#Invoke-ApplyAppState) | **POST** /api/apps/{id}/apply | Apply State of an App
[**Invoke-DeleteApp**](AppsApi.md#Invoke-DeleteApp) | **DELETE** /api/apps/{id} | Delete an App
[**Get-App**](AppsApi.md#Get-App) | **GET** /api/apps/{id} | Get a Specific App
[**Get-AppSecurityGroups**](AppsApi.md#Get-AppSecurityGroups) | **GET** /api/apps/{id}/security-groups | Get Security Groups for an App
[**Get-AppState**](AppsApi.md#Get-AppState) | **GET** /api/apps/{id}/state | Get State of an App
[**Invoke-ListApps**](AppsApi.md#Invoke-ListApps) | **GET** /api/apps | Get All Apps
[**Invoke-PrepareAppApply**](AppsApi.md#Invoke-PrepareAppApply) | **GET** /api/apps/{id}/prepare-apply | Prepare To Apply an App
[**Invoke-RefreshAppState**](AppsApi.md#Invoke-RefreshAppState) | **POST** /api/apps/{id}/refresh | Refresh State of an App
[**Remove-AppInstance**](AppsApi.md#Remove-AppInstance) | **POST** /api/apps/{id}/remove-instance | Remove Instance from App
[**Set-AppSecurityGroups**](AppsApi.md#Set-AppSecurityGroups) | **POST** /api/apps/{id}/security-groups | Set Security Groups for an App
[**Update-App**](AppsApi.md#Update-App) | **PUT** /api/apps/{id} | Updating an App
[**Confirm-AppState**](AppsApi.md#Confirm-AppState) | **POST** /api/apps/{id}/validate-apply | Validate Apply State for an App


<a id="Add-AppInstance"></a>
# **Add-AppInstance**
> GetApp200Response Add-AppInstance<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddAppInstanceRequest] <PSCustomObject><br>

Add Existing Instance to App

Add Existing Instance to App

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$AddAppInstanceRequest = Initialize-AddAppInstanceRequest -InstanceId 0 -TierName "MyTierName" # AddAppInstanceRequest |  (optional)

# Add Existing Instance to App
try {
    $Result = Add-AppInstance -Id $Id -AddAppInstanceRequest $AddAppInstanceRequest
} catch {
    Write-Host ("Exception occurred when calling Add-AppInstance: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **AddAppInstanceRequest** | [**AddAppInstanceRequest**](AddAppInstanceRequest.md)|  | [optional] 

### Return type

[**GetApp200Response**](GetApp200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Add-AppUndoDelete"></a>
# **Add-AppUndoDelete**
> GetApp200Response Add-AppUndoDelete<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Undo Delete of an App

This operation will undo the delete of an app that is pending removal.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Undo Delete of an App
try {
    $Result = Add-AppUndoDelete -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Add-AppUndoDelete: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetApp200Response**](GetApp200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Add-Apps"></a>
# **Add-Apps**
> AddApps200Response Add-Apps<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AppCreate] <PSCustomObject><br>

Create an App

Create an App

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$AppCreateBlueprintId = Initialize-AppCreateBlueprintId 
$AppCreateGroup = Initialize-AppCreateGroup -Id 0
$AppCreateDefaultCloud = Initialize-AppCreateDefaultCloud -Id 0
$AppCreate = Initialize-AppCreate -TemplateId 0 -BlueprintId $AppCreateBlueprintId -Name "MyName" -Description "MyDescription" -Labels "MyLabels" -Group $AppCreateGroup -DefaultCloud $AppCreateDefaultCloud -Environment "MyEnvironment" -Tiers # AppCreate |  (optional)

# Create an App
try {
    $Result = Add-Apps -AppCreate $AppCreate
} catch {
    Write-Host ("Exception occurred when calling Add-Apps: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AppCreate** | [**AppCreate**](AppCreate.md)|  | [optional] 

### Return type

[**AddApps200Response**](AddApps200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ApplyAppState"></a>
# **Invoke-ApplyAppState**
> Model200Success Invoke-ApplyAppState<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ApplyAppStateRequest] <PSCustomObject><br>

Apply State of an App

This endpoint provides a way to apply the state of an app. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$ApplyAppStateRequest = Initialize-ApplyAppStateRequest -TemplateParameter # ApplyAppStateRequest |  (optional)

# Apply State of an App
try {
    $Result = Invoke-ApplyAppState -Id $Id -ApplyAppStateRequest $ApplyAppStateRequest
} catch {
    Write-Host ("Exception occurred when calling Invoke-ApplyAppState: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **ApplyAppStateRequest** | [**ApplyAppStateRequest**](ApplyAppStateRequest.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-DeleteApp"></a>
# **Invoke-DeleteApp**
> Model200Success Invoke-DeleteApp<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-RemoveInstances] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-PreserveVolumes] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-KeepBackups] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ReleaseFloatingIps] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ReleaseEIPs] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Force] <String><br>

Delete an App

Will delete an app. Use removeInstances=on to also delete the instances in the app and all associated monitors and backups.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$RemoveInstances = "on" # String | Remove Instances (optional) (default to "off")
$PreserveVolumes = "on" # String | Preserve Volumes (optional) (default to "off")
$KeepBackups = "on" # String | Preserve copy of backups (optional) (default to "off")
$ReleaseFloatingIps = "off" # String | Release Floating IPs (optional) (default to "on")
$ReleaseEIPs = "off" # String | Alias for releaseFloatingIps (optional) (default to "on")
$Force = "on" # String | Force Delete (optional) (default to "off")

# Delete an App
try {
    $Result = Invoke-DeleteApp -Id $Id -RemoveInstances $RemoveInstances -PreserveVolumes $PreserveVolumes -KeepBackups $KeepBackups -ReleaseFloatingIps $ReleaseFloatingIps -ReleaseEIPs $ReleaseEIPs -Force $Force
} catch {
    Write-Host ("Exception occurred when calling Invoke-DeleteApp: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **RemoveInstances** | **String**| Remove Instances | [optional] [default to &quot;off&quot;]
 **PreserveVolumes** | **String**| Preserve Volumes | [optional] [default to &quot;off&quot;]
 **KeepBackups** | **String**| Preserve copy of backups | [optional] [default to &quot;off&quot;]
 **ReleaseFloatingIps** | **String**| Release Floating IPs | [optional] [default to &quot;on&quot;]
 **ReleaseEIPs** | **String**| Alias for releaseFloatingIps | [optional] [default to &quot;on&quot;]
 **Force** | **String**| Force Delete | [optional] [default to &quot;off&quot;]

### Return type

[**Model200Success**](Model200Success.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-App"></a>
# **Get-App**
> GetApp200Response Get-App<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Get a Specific App

This endpoint retrieves a specific app.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Get a Specific App
try {
    $Result = Get-App -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-App: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetApp200Response**](GetApp200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-AppSecurityGroups"></a>
# **Get-AppSecurityGroups**
> GetAppSecurityGroups200Response Get-AppSecurityGroups<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Get Security Groups for an App

This returns a list of all of the security groups applied to an app and whether the firewall is enabled.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Get Security Groups for an App
try {
    $Result = Get-AppSecurityGroups -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-AppSecurityGroups: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetAppSecurityGroups200Response**](GetAppSecurityGroups200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-AppState"></a>
# **Get-AppState**
> GetAppState200Response Get-AppState<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Get State of an App

This endpoint provides a way to view the state of an app. The response includes output and resource planning information from the template provider software such as Terraform. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Get State of an App
try {
    $Result = Get-AppState -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-AppState: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetAppState200Response**](GetAppState200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListApps"></a>
# **Invoke-ListApps**
> ListApps200Response Invoke-ListApps<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Max] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Offset] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Name] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Phrase] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-CreatedBy] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ShowDeleted] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Labels] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AllLabels] <String><br>

Get All Apps

This endpoint retrieves a paginated list of apps. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Max = 789 # Int64 | Maximum number of records to return (optional) (default to 25)
$Offset = 789 # Int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
$Name = "example-%" # String | Filter by name, wildcard may be specified as %, eg. example-% (optional)
$Phrase = "MyPhrase" # String | Search phrase for partial matches on name or description (optional)
$CreatedBy = 63 # Int64 | The User ID for Filtering (optional)
$ShowDeleted = $true # Boolean | If true, includes instances in pending removal status. (optional) (default to $false)
$Labels = "MyLabels" # String | Filter by label(s), matches records that contain any of the specified labels (optional)
$AllLabels = "MyAllLabels" # String | Filter by label(s), matches records that contain all of the specified labels (optional)

# Get All Apps
try {
    $Result = Invoke-ListApps -Max $Max -Offset $Offset -Name $Name -Phrase $Phrase -CreatedBy $CreatedBy -ShowDeleted $ShowDeleted -Labels $Labels -AllLabels $AllLabels
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListApps: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Max** | **Int64**| Maximum number of records to return | [optional] [default to 25]
 **Offset** | **Int64**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **Name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **Phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **CreatedBy** | **Int64**| The User ID for Filtering | [optional] 
 **ShowDeleted** | **Boolean**| If true, includes instances in pending removal status. | [optional] [default to $false]
 **Labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **AllLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

[**ListApps200Response**](ListApps200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-PrepareAppApply"></a>
# **Invoke-PrepareAppApply**
> PrepareAppApply200Response Invoke-PrepareAppApply<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Prepare To Apply an App

This endpoint provides a way to view the current app configuration and templateParameter variables available to apply. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Prepare To Apply an App
try {
    $Result = Invoke-PrepareAppApply -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Invoke-PrepareAppApply: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**PrepareAppApply200Response**](PrepareAppApply200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-RefreshAppState"></a>
# **Invoke-RefreshAppState**
> Model200Success Invoke-RefreshAppState<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Body] <System.Nullable[SystemCollectionsHashtable]><br>

Refresh State of an App

This endpoint provides a way to refresh the state of an app. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.   

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$Body = @{ key_example = ... } # SystemCollectionsHashtable |  (optional)

# Refresh State of an App
try {
    $Result = Invoke-RefreshAppState -Id $Id -Body $Body
} catch {
    Write-Host ("Exception occurred when calling Invoke-RefreshAppState: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **Body** | **SystemCollectionsHashtable**|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Remove-AppInstance"></a>
# **Remove-AppInstance**
> GetApp200Response Remove-AppInstance<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-RemoveAppInstanceRequest] <PSCustomObject><br>

Remove Instance from App

Remove Instance from App

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$RemoveAppInstanceRequest = Initialize-RemoveAppInstanceRequest -InstanceId 0 # RemoveAppInstanceRequest |  (optional)

# Remove Instance from App
try {
    $Result = Remove-AppInstance -Id $Id -RemoveAppInstanceRequest $RemoveAppInstanceRequest
} catch {
    Write-Host ("Exception occurred when calling Remove-AppInstance: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **RemoveAppInstanceRequest** | [**RemoveAppInstanceRequest**](RemoveAppInstanceRequest.md)|  | [optional] 

### Return type

[**GetApp200Response**](GetApp200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Set-AppSecurityGroups"></a>
# **Set-AppSecurityGroups**
> SetAppSecurityGroups200Response Set-AppSecurityGroups<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-SetAppSecurityGroupsRequest] <PSCustomObject><br>

Set Security Groups for an App

Set Security Groups for an App

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$SetAppSecurityGroupsRequest = Initialize-SetAppSecurityGroupsRequest -SecurityGroupIds 0 # SetAppSecurityGroupsRequest |  (optional)

# Set Security Groups for an App
try {
    $Result = Set-AppSecurityGroups -Id $Id -SetAppSecurityGroupsRequest $SetAppSecurityGroupsRequest
} catch {
    Write-Host ("Exception occurred when calling Set-AppSecurityGroups: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **SetAppSecurityGroupsRequest** | [**SetAppSecurityGroupsRequest**](SetAppSecurityGroupsRequest.md)|  | [optional] 

### Return type

[**SetAppSecurityGroups200Response**](SetAppSecurityGroups200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Update-App"></a>
# **Update-App**
> GetApp200Response Update-App<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AppUpdate] <PSCustomObject><br>

Updating an App

This endpoint provides updating of some basic app settings.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$AppUpdate = Initialize-AppUpdate -Name "MyName" -Description "MyDescription" -Labels "MyLabels" -Environment "MyEnvironment" -OwnerId 0 # AppUpdate |  (optional)

# Updating an App
try {
    $Result = Update-App -Id $Id -AppUpdate $AppUpdate
} catch {
    Write-Host ("Exception occurred when calling Update-App: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **AppUpdate** | [**AppUpdate**](AppUpdate.md)|  | [optional] 

### Return type

[**GetApp200Response**](GetApp200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Confirm-AppState"></a>
# **Confirm-AppState**
> ValidateAppState200Response Confirm-AppState<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ApplyAppStateRequest] <PSCustomObject><br>

Validate Apply State for an App

This endpoint provides a way to validate app configuration and templateParameter variables before executing the apply. This only validates the configuration to see any planned changes and it does not actually apply the changes. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$ApplyAppStateRequest = Initialize-ApplyAppStateRequest -TemplateParameter # ApplyAppStateRequest |  (optional)

# Validate Apply State for an App
try {
    $Result = Confirm-AppState -Id $Id -ApplyAppStateRequest $ApplyAppStateRequest
} catch {
    Write-Host ("Exception occurred when calling Confirm-AppState: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **ApplyAppStateRequest** | [**ApplyAppStateRequest**](ApplyAppStateRequest.md)|  | [optional] 

### Return type

[**ValidateAppState200Response**](ValidateAppState200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

