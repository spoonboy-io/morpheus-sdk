# PSOpenAPITools.PSOpenAPITools/Api.BlueprintsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add-Blueprint**](BlueprintsApi.md#Add-Blueprint) | **POST** /api/blueprints | Create a Blueprint
[**Invoke-DeleteBlueprint**](BlueprintsApi.md#Invoke-DeleteBlueprint) | **DELETE** /api/blueprints/{id} | Delete a Blueprint
[**Get-Blueprint**](BlueprintsApi.md#Get-Blueprint) | **GET** /api/blueprints/{id} | Get a Specific Blueprint
[**Invoke-ListBlueprints**](BlueprintsApi.md#Invoke-ListBlueprints) | **GET** /api/blueprints | Get All Blueprints
[**Update-Blueprint**](BlueprintsApi.md#Update-Blueprint) | **PUT** /api/blueprints/{id} | Updating a Blueprint
[**Update-BlueprintImage**](BlueprintsApi.md#Update-BlueprintImage) | **POST** /api/blueprints/{id}/image | Update Blueprint Image
[**Update-BlueprintPermissions**](BlueprintsApi.md#Update-BlueprintPermissions) | **PUT** /api/blueprints/{id}/update-permissions | Update Blueprint Permissions


<a id="Add-Blueprint"></a>
# **Add-Blueprint**
> AddBlueprint200Response Add-Blueprint<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddBlueprintRequest] <PSCustomObject><br>

Create a Blueprint

Create a Blueprint

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$BlueprintARMCreateArmGit = Initialize-BlueprintARMCreateArmGit -RepoId 0 -Path "MyPath" -IntegrationId 0 -Branch "MyBranch"
$BlueprintARMCreateArmInstallAgent = Initialize-BlueprintARMCreateArmInstallAgent 
$BlueprintARMCreateArmCloudInitEnabled = Initialize-BlueprintARMCreateArmCloudInitEnabled 
$BlueprintARMCreateArm = Initialize-BlueprintARMCreateArm -ConfigType "json" -Json "MyJson" -Yaml "MyYaml" -Git $BlueprintARMCreateArmGit -OsType "linux" -InstallAgent $BlueprintARMCreateArmInstallAgent -CloudInitEnabled $BlueprintARMCreateArmCloudInitEnabled

$BlueprintCFTCreateCloudFormationGit = Initialize-BlueprintCFTCreateCloudFormationGit -RepoId 0 -Path "MyPath" -IntegrationId 0 -Branch "MyBranch"
$BlueprintCFTCreateCloudFormation = Initialize-BlueprintCFTCreateCloudFormation -ConfigType "json" -Json "MyJson" -Yaml "MyYaml" -Git $BlueprintCFTCreateCloudFormationGit -IAM $false -CAPABILITYNAMEDIAM $false -CAPABILITYAUTOEXPAND $false -InstallAgent $false -CloudInitEnabled $false

$BlueprintHelmCreateHelmGit = Initialize-BlueprintHelmCreateHelmGit -RepoId 0 -Path "MyPath" -IntegrationId 0 -Branch "MyBranch"
$BlueprintHelmCreateHelm = Initialize-BlueprintHelmCreateHelm -ConfigType "git" -Git $BlueprintHelmCreateHelmGit

$BlueprintKubernetesCreateKubernetesGit = Initialize-BlueprintKubernetesCreateKubernetesGit -RepoId 0 -Path "MyPath" -IntegrationId 0 -Branch "MyBranch"
$BlueprintKubernetesCreateKubernetes = Initialize-BlueprintKubernetesCreateKubernetes -ConfigType "yaml" -Yaml "MyYaml" -Git $BlueprintKubernetesCreateKubernetesGit

$BlueprintKubernetesCreateConfigSpecsInner = Initialize-BlueprintKubernetesCreateConfigSpecsInner -Id 0 -Value 0 -Name "MyName"
$BlueprintTerraformCreateConfig = Initialize-BlueprintTerraformCreateConfig -Specs $BlueprintKubernetesCreateConfigSpecsInner

$BlueprintTerraformCreateTerraformGit = Initialize-BlueprintTerraformCreateTerraformGit -RepoId 0 -Path "MyPath" -IntegrationId 0 -Branch "MyBranch"
$BlueprintTerraformCreateTerraform = Initialize-BlueprintTerraformCreateTerraform -ConfigType "tf" -Json "MyJson" -Tf "MyTf" -Git $BlueprintTerraformCreateTerraformGit -TfvarSecret "MyTfvarSecret"

$AddBlueprintRequest = Initialize-AddBlueprintRequest -Name "MyName" -Image "MyImage" -Type "arm" -Labels "MyLabels" -Arm $BlueprintARMCreateArm -CloudFormation $BlueprintCFTCreateCloudFormation -Helm $BlueprintHelmCreateHelm -Kubernetes $BlueprintKubernetesCreateKubernetes -Config $BlueprintTerraformCreateConfig -Tiers  -Terraform $BlueprintTerraformCreateTerraform # AddBlueprintRequest |  (optional)

# Create a Blueprint
try {
    $Result = Add-Blueprint -AddBlueprintRequest $AddBlueprintRequest
} catch {
    Write-Host ("Exception occurred when calling Add-Blueprint: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AddBlueprintRequest** | [**AddBlueprintRequest**](AddBlueprintRequest.md)|  | [optional] 

### Return type

[**AddBlueprint200Response**](AddBlueprint200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-DeleteBlueprint"></a>
# **Invoke-DeleteBlueprint**
> Model200Success Invoke-DeleteBlueprint<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Delete a Blueprint

Delete a Blueprint

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Delete a Blueprint
try {
    $Result = Invoke-DeleteBlueprint -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Invoke-DeleteBlueprint: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
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

<a id="Get-Blueprint"></a>
# **Get-Blueprint**
> GetBlueprint200Response Get-Blueprint<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Get a Specific Blueprint

This endpoint retrieves a specific blueprint.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Get a Specific Blueprint
try {
    $Result = Get-Blueprint -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-Blueprint: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListBlueprints"></a>
# **Invoke-ListBlueprints**
> ListBlueprints200Response Invoke-ListBlueprints<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Max] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Offset] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Name] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Phrase] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Labels] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AllLabels] <String><br>

Get All Blueprints

This endpoint retrieves all blueprints.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Max = 789 # Int64 | Maximum number of records to return (optional) (default to 25)
$Offset = 789 # Int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
$Name = "example-%" # String | Filter by name, wildcard may be specified as %, eg. example-% (optional)
$Phrase = "MyPhrase" # String | Search phrase for partial matches on name or description (optional)
$Labels = "MyLabels" # String | Filter by label(s), matches records that contain any of the specified labels (optional)
$AllLabels = "MyAllLabels" # String | Filter by label(s), matches records that contain all of the specified labels (optional)

# Get All Blueprints
try {
    $Result = Invoke-ListBlueprints -Max $Max -Offset $Offset -Name $Name -Phrase $Phrase -Labels $Labels -AllLabels $AllLabels
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListBlueprints: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
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
 **Labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **AllLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

[**ListBlueprints200Response**](ListBlueprints200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Update-Blueprint"></a>
# **Update-Blueprint**
> GetBlueprint200Response Update-Blueprint<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddBlueprintRequest] <PSCustomObject><br>

Updating a Blueprint

Update a Blueprint. This overwrites the entire config, so the entire blueprint config should be passed.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$BlueprintARMCreateArmGit = Initialize-BlueprintARMCreateArmGit -RepoId 0 -Path "MyPath" -IntegrationId 0 -Branch "MyBranch"
$BlueprintARMCreateArmInstallAgent = Initialize-BlueprintARMCreateArmInstallAgent 
$BlueprintARMCreateArmCloudInitEnabled = Initialize-BlueprintARMCreateArmCloudInitEnabled 
$BlueprintARMCreateArm = Initialize-BlueprintARMCreateArm -ConfigType "json" -Json "MyJson" -Yaml "MyYaml" -Git $BlueprintARMCreateArmGit -OsType "linux" -InstallAgent $BlueprintARMCreateArmInstallAgent -CloudInitEnabled $BlueprintARMCreateArmCloudInitEnabled

$BlueprintCFTCreateCloudFormationGit = Initialize-BlueprintCFTCreateCloudFormationGit -RepoId 0 -Path "MyPath" -IntegrationId 0 -Branch "MyBranch"
$BlueprintCFTCreateCloudFormation = Initialize-BlueprintCFTCreateCloudFormation -ConfigType "json" -Json "MyJson" -Yaml "MyYaml" -Git $BlueprintCFTCreateCloudFormationGit -IAM $false -CAPABILITYNAMEDIAM $false -CAPABILITYAUTOEXPAND $false -InstallAgent $false -CloudInitEnabled $false

$BlueprintHelmCreateHelmGit = Initialize-BlueprintHelmCreateHelmGit -RepoId 0 -Path "MyPath" -IntegrationId 0 -Branch "MyBranch"
$BlueprintHelmCreateHelm = Initialize-BlueprintHelmCreateHelm -ConfigType "git" -Git $BlueprintHelmCreateHelmGit

$BlueprintKubernetesCreateKubernetesGit = Initialize-BlueprintKubernetesCreateKubernetesGit -RepoId 0 -Path "MyPath" -IntegrationId 0 -Branch "MyBranch"
$BlueprintKubernetesCreateKubernetes = Initialize-BlueprintKubernetesCreateKubernetes -ConfigType "yaml" -Yaml "MyYaml" -Git $BlueprintKubernetesCreateKubernetesGit

$BlueprintKubernetesCreateConfigSpecsInner = Initialize-BlueprintKubernetesCreateConfigSpecsInner -Id 0 -Value 0 -Name "MyName"
$BlueprintTerraformCreateConfig = Initialize-BlueprintTerraformCreateConfig -Specs $BlueprintKubernetesCreateConfigSpecsInner

$BlueprintTerraformCreateTerraformGit = Initialize-BlueprintTerraformCreateTerraformGit -RepoId 0 -Path "MyPath" -IntegrationId 0 -Branch "MyBranch"
$BlueprintTerraformCreateTerraform = Initialize-BlueprintTerraformCreateTerraform -ConfigType "tf" -Json "MyJson" -Tf "MyTf" -Git $BlueprintTerraformCreateTerraformGit -TfvarSecret "MyTfvarSecret"

$AddBlueprintRequest = Initialize-AddBlueprintRequest -Name "MyName" -Image "MyImage" -Type "arm" -Labels "MyLabels" -Arm $BlueprintARMCreateArm -CloudFormation $BlueprintCFTCreateCloudFormation -Helm $BlueprintHelmCreateHelm -Kubernetes $BlueprintKubernetesCreateKubernetes -Config $BlueprintTerraformCreateConfig -Tiers  -Terraform $BlueprintTerraformCreateTerraform # AddBlueprintRequest |  (optional)

# Updating a Blueprint
try {
    $Result = Update-Blueprint -Id $Id -AddBlueprintRequest $AddBlueprintRequest
} catch {
    Write-Host ("Exception occurred when calling Update-Blueprint: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **AddBlueprintRequest** | [**AddBlueprintRequest**](AddBlueprintRequest.md)|  | [optional] 

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Update-BlueprintImage"></a>
# **Update-BlueprintImage**
> GetBlueprint200Response Update-BlueprintImage<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-TemplateImage] <System.IO.FileInfo><br>

Update Blueprint Image

Update Blueprint Image

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$TemplateImage =  # System.IO.FileInfo |  (optional)

# Update Blueprint Image
try {
    $Result = Update-BlueprintImage -Id $Id -TemplateImage $TemplateImage
} catch {
    Write-Host ("Exception occurred when calling Update-BlueprintImage: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **TemplateImage** | **System.IO.FileInfo****System.IO.FileInfo**|  | [optional] 

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Update-BlueprintPermissions"></a>
# **Update-BlueprintPermissions**
> GetBlueprint200Response Update-BlueprintPermissions<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-UpdateBlueprintPermissionsRequest] <PSCustomObject><br>

Update Blueprint Permissions

Update Blueprint Permissions

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$UpdateBlueprintPermissionsRequestResourcePermissionSitesInner = Initialize-UpdateBlueprintPermissionsRequestResourcePermissionSitesInner -Id 0
$UpdateBlueprintPermissionsRequestResourcePermission = Initialize-UpdateBlueprintPermissionsRequestResourcePermission -All $false -Sites $UpdateBlueprintPermissionsRequestResourcePermissionSitesInner -OwnerId 0

$UpdateBlueprintPermissionsRequest = Initialize-UpdateBlueprintPermissionsRequest -ResourcePermission $UpdateBlueprintPermissionsRequestResourcePermission # UpdateBlueprintPermissionsRequest |  (optional)

# Update Blueprint Permissions
try {
    $Result = Update-BlueprintPermissions -Id $Id -UpdateBlueprintPermissionsRequest $UpdateBlueprintPermissionsRequest
} catch {
    Write-Host ("Exception occurred when calling Update-BlueprintPermissions: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **UpdateBlueprintPermissionsRequest** | [**UpdateBlueprintPermissionsRequest**](UpdateBlueprintPermissionsRequest.md)|  | [optional] 

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

