# PSOpenAPITools.PSOpenAPITools/Api.CatalogItemsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add-CatalogItemType**](CatalogItemsApi.md#Add-CatalogItemType) | **POST** /api/catalog-item-types | Create a Catalog Item Type
[**Get-CatalogItemType**](CatalogItemsApi.md#Get-CatalogItemType) | **GET** /api/catalog-item-types/{id} | Get a Specific Catalog Item Type
[**Invoke-ListCatalogItemTypes**](CatalogItemsApi.md#Invoke-ListCatalogItemTypes) | **GET** /api/catalog-item-types | Get All Catalog Item Types
[**Remove-CatalogItemType**](CatalogItemsApi.md#Remove-CatalogItemType) | **DELETE** /api/catalog-item-types/{id} | Delete a Catalog Item Type
[**Update-CatalogItemType**](CatalogItemsApi.md#Update-CatalogItemType) | **PUT** /api/catalog-item-types/{id} | Update a Catalog Item Type
[**Update-CatalogItemTypeLogo**](CatalogItemsApi.md#Update-CatalogItemTypeLogo) | **PUT** /api/catalog-item-types/{id}/update-logo | Update Logo For Catalog Item Type


<a id="Add-CatalogItemType"></a>
# **Add-CatalogItemType**
> AddCatalogItemType200Response Add-CatalogItemType<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddCatalogItemTypeRequest] <PSCustomObject><br>

Create a Catalog Item Type

Use this command to create a catalog item type.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$CatalogItemTypeBlueprintCreateBlueprint = Initialize-CatalogItemTypeBlueprintCreateBlueprint -Id 0 -Name "MyName"
$UpdateBlueprintPermissionsRequestResourcePermissionSitesInner = Initialize-UpdateBlueprintPermissionsRequestResourcePermissionSitesInner -Id 0
$AddCatalogItemTypeRequestCatalogItemType = Initialize-AddCatalogItemTypeRequestCatalogItemType -Name "MyName" -Code "MyCode" -Category "MyCategory" -Description "MyDescription" -Labels "MyLabels" -Type "instance" -Visibility "MyVisibility" -LayoutCode "MyLayoutCode" -IconPath "MyIconPath" -Enabled $false -Featured $false -AllowQuantity $false -Config  -OptionTypes 0 -Content "MyContent" -Blueprint $CatalogItemTypeBlueprintCreateBlueprint -AppSpec "MyAppSpec" -Workflow $UpdateBlueprintPermissionsRequestResourcePermissionSitesInner -Context "instance" -WorkflowConfig "MyWorkflowConfig"

$AddCatalogItemTypeRequest = Initialize-AddCatalogItemTypeRequest -CatalogItemType $AddCatalogItemTypeRequestCatalogItemType # AddCatalogItemTypeRequest |  (optional)

# Create a Catalog Item Type
try {
    $Result = Add-CatalogItemType -AddCatalogItemTypeRequest $AddCatalogItemTypeRequest
} catch {
    Write-Host ("Exception occurred when calling Add-CatalogItemType: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AddCatalogItemTypeRequest** | [**AddCatalogItemTypeRequest**](AddCatalogItemTypeRequest.md)|  | [optional] 

### Return type

[**AddCatalogItemType200Response**](AddCatalogItemType200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-CatalogItemType"></a>
# **Get-CatalogItemType**
> GetCatalogItemType200Response Get-CatalogItemType<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Get a Specific Catalog Item Type

This endpoint retrieves a specific category item type.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Get a Specific Catalog Item Type
try {
    $Result = Get-CatalogItemType -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-CatalogItemType: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetCatalogItemType200Response**](GetCatalogItemType200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListCatalogItemTypes"></a>
# **Invoke-ListCatalogItemTypes**
> ListCatalogItemTypes200Response Invoke-ListCatalogItemTypes<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Max] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Offset] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Sort] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Direction] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Phrase] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Name] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Description] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Enabled] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Featured] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Labels] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AllLabels] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Code] <String><br>

Get All Catalog Item Types

This endpoint retrieves all catalog item types.

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
$Description = "The desription of my cool object" # String | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)
$Enabled = $false # Boolean | Filter by enabled (optional)
$Featured = $false # Boolean | Filter by featured (optional)
$Labels = "MyLabels" # String | Filter by label(s), matches records that contain any of the specified labels (optional)
$AllLabels = "MyAllLabels" # String | Filter by label(s), matches records that contain all of the specified labels (optional)
$Code = "azr" # String | If specified will return an exact match on code (optional)

# Get All Catalog Item Types
try {
    $Result = Invoke-ListCatalogItemTypes -Max $Max -Offset $Offset -Sort $Sort -Direction $Direction -Phrase $Phrase -Name $Name -Description $Description -Enabled $Enabled -Featured $Featured -Labels $Labels -AllLabels $AllLabels -Code $Code
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListCatalogItemTypes: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
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
 **Description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional] 
 **Enabled** | **Boolean**| Filter by enabled | [optional] 
 **Featured** | **Boolean**| Filter by featured | [optional] 
 **Labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **AllLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 
 **Code** | **String**| If specified will return an exact match on code | [optional] 

### Return type

[**ListCatalogItemTypes200Response**](ListCatalogItemTypes200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Remove-CatalogItemType"></a>
# **Remove-CatalogItemType**
> Model200Success Remove-CatalogItemType<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Delete a Catalog Item Type

Will delete a catalog item type.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Delete a Catalog Item Type
try {
    $Result = Remove-CatalogItemType -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Remove-CatalogItemType: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
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

<a id="Update-CatalogItemType"></a>
# **Update-CatalogItemType**
> UpdateCatalogItemType200Response Update-CatalogItemType<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-UpdateCatalogItemTypeRequest] <PSCustomObject><br>

Update a Catalog Item Type

Use this command to update an existing catalog item type.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$CatalogItemTypeBlueprintCreateBlueprint = Initialize-CatalogItemTypeBlueprintCreateBlueprint -Id 0 -Name "MyName"
$UpdateBlueprintPermissionsRequestResourcePermissionSitesInner = Initialize-UpdateBlueprintPermissionsRequestResourcePermissionSitesInner -Id 0
$UpdateCatalogItemTypeRequestCatalogItemType = Initialize-UpdateCatalogItemTypeRequestCatalogItemType -Name "MyName" -Code "MyCode" -Category "MyCategory" -Description "MyDescription" -Labels "MyLabels" -Type "instance" -Visibility "MyVisibility" -LayoutCode "MyLayoutCode" -IconPath "MyIconPath" -Enabled $false -Featured $false -AllowQuantity $false -Config  -OptionTypes 0 -Blueprint $CatalogItemTypeBlueprintCreateBlueprint -AppSpec "MyAppSpec" -Workflow $UpdateBlueprintPermissionsRequestResourcePermissionSitesInner -Context "MyContext" -WorkflowConfig "MyWorkflowConfig"

$UpdateCatalogItemTypeRequest = Initialize-UpdateCatalogItemTypeRequest -CatalogItemType $UpdateCatalogItemTypeRequestCatalogItemType # UpdateCatalogItemTypeRequest |  (optional)

# Update a Catalog Item Type
try {
    $Result = Update-CatalogItemType -Id $Id -UpdateCatalogItemTypeRequest $UpdateCatalogItemTypeRequest
} catch {
    Write-Host ("Exception occurred when calling Update-CatalogItemType: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **UpdateCatalogItemTypeRequest** | [**UpdateCatalogItemTypeRequest**](UpdateCatalogItemTypeRequest.md)|  | [optional] 

### Return type

[**UpdateCatalogItemType200Response**](UpdateCatalogItemType200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Update-CatalogItemTypeLogo"></a>
# **Update-CatalogItemTypeLogo**
> UpdateCatalogItemType200Response Update-CatalogItemTypeLogo<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-CatalogItemTypeLogo] <System.IO.FileInfo><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-CatalogItemTypeDarkLogo] <System.IO.FileInfo><br>

Update Logo For Catalog Item Type

Use this command to update the logo and dark logo images for an existing catalog item type. This endpoint expects multipart form data as the request format, not JSON.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$CatalogItemTypeLogo =  # System.IO.FileInfo | Logo File png,jpg,svg (optional)
$CatalogItemTypeDarkLogo =  # System.IO.FileInfo | Dark Logo File png,jpg,svg (optional)

# Update Logo For Catalog Item Type
try {
    $Result = Update-CatalogItemTypeLogo -Id $Id -CatalogItemTypeLogo $CatalogItemTypeLogo -CatalogItemTypeDarkLogo $CatalogItemTypeDarkLogo
} catch {
    Write-Host ("Exception occurred when calling Update-CatalogItemTypeLogo: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **CatalogItemTypeLogo** | **System.IO.FileInfo****System.IO.FileInfo**| Logo File png,jpg,svg | [optional] 
 **CatalogItemTypeDarkLogo** | **System.IO.FileInfo****System.IO.FileInfo**| Dark Logo File png,jpg,svg | [optional] 

### Return type

[**UpdateCatalogItemType200Response**](UpdateCatalogItemType200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

