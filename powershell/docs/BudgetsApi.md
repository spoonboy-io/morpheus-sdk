# PSOpenAPITools.PSOpenAPITools/Api.BudgetsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add-Budgets**](BudgetsApi.md#Add-Budgets) | **POST** /api/budgets | Creates a Budget
[**Get-Budgets**](BudgetsApi.md#Get-Budgets) | **GET** /api/budgets/{id} | Retrieves a Specific Budget
[**Invoke-ListBudgets**](BudgetsApi.md#Invoke-ListBudgets) | **GET** /api/budgets | Retrieves all Budgets
[**Remove-Budgets**](BudgetsApi.md#Remove-Budgets) | **DELETE** /api/budgets/{id} | Deletes a Budget
[**Update-Budgets**](BudgetsApi.md#Update-Budgets) | **PUT** /api/budgets/{id} | Updates a Budget


<a id="Add-Budgets"></a>
# **Add-Budgets**
> AddBudgets200Response Add-Budgets<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddBudgetsRequest] <PSCustomObject><br>

Creates a Budget

Creates a budget. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$AddBudgetsRequestBudget = Initialize-AddBudgetsRequestBudget -Name "MyName" -Description "MyDescription" -Scope "account" -Period "MyPeriod" -Year 2020 -StartDate (Get-Date) -EndDate (Get-Date) -Interval "year" -ScopeTenantId 0 -ScopeGroupId 0 -ScopeCloudId 0 -ScopeUserId 0 -Costs 0 -Enabled $false
$AddBudgetsRequest = Initialize-AddBudgetsRequest -Budget $AddBudgetsRequestBudget # AddBudgetsRequest |  (optional)

# Creates a Budget
try {
    $Result = Add-Budgets -AddBudgetsRequest $AddBudgetsRequest
} catch {
    Write-Host ("Exception occurred when calling Add-Budgets: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AddBudgetsRequest** | [**AddBudgetsRequest**](AddBudgetsRequest.md)|  | [optional] 

### Return type

[**AddBudgets200Response**](AddBudgets200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-Budgets"></a>
# **Get-Budgets**
> GetBudgets200Response Get-Budgets<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Retrieves a Specific Budget

Retrieves a specific budget. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Retrieves a Specific Budget
try {
    $Result = Get-Budgets -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-Budgets: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetBudgets200Response**](GetBudgets200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListBudgets"></a>
# **Invoke-ListBudgets**
> ListBudgets200Response Invoke-ListBudgets<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Max] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Offset] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Sort] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Direction] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Phrase] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Name] <String><br>

Retrieves all Budgets

Retrieves all budgets. 

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

# Retrieves all Budgets
try {
    $Result = Invoke-ListBudgets -Max $Max -Offset $Offset -Sort $Sort -Direction $Direction -Phrase $Phrase -Name $Name
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListBudgets: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
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

[**ListBudgets200Response**](ListBudgets200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Remove-Budgets"></a>
# **Remove-Budgets**
> Model200Success Remove-Budgets<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Deletes a Budget

Deletes a specified Budget. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Deletes a Budget
try {
    $Result = Remove-Budgets -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Remove-Budgets: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
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

<a id="Update-Budgets"></a>
# **Update-Budgets**
> AddBudgets200Response Update-Budgets<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddBudgetsRequest] <PSCustomObject><br>

Updates a Budget

Updates a budget. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$AddBudgetsRequestBudget = Initialize-AddBudgetsRequestBudget -Name "MyName" -Description "MyDescription" -Scope "account" -Period "MyPeriod" -Year 2020 -StartDate (Get-Date) -EndDate (Get-Date) -Interval "year" -ScopeTenantId 0 -ScopeGroupId 0 -ScopeCloudId 0 -ScopeUserId 0 -Costs 0 -Enabled $false
$AddBudgetsRequest = Initialize-AddBudgetsRequest -Budget $AddBudgetsRequestBudget # AddBudgetsRequest |  (optional)

# Updates a Budget
try {
    $Result = Update-Budgets -Id $Id -AddBudgetsRequest $AddBudgetsRequest
} catch {
    Write-Host ("Exception occurred when calling Update-Budgets: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **AddBudgetsRequest** | [**AddBudgetsRequest**](AddBudgetsRequest.md)|  | [optional] 

### Return type

[**AddBudgets200Response**](AddBudgets200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

