# PSOpenAPITools.PSOpenAPITools/Api.ArchivesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add-ArchiveBucket**](ArchivesApi.md#Add-ArchiveBucket) | **POST** /api/archives/buckets | Create an Archive Bucket
[**Add-ArchiveFile**](ArchivesApi.md#Add-ArchiveFile) | **POST** /api/archives/buckets/{bucket}/files/{filepath} | Upload Archive File
[**Add-ArchiveFileLink**](ArchivesApi.md#Add-ArchiveFileLink) | **POST** /api/archives/files/{id}/links | Create an Archive File Link
[**Invoke-DeleteArchiveBucket**](ArchivesApi.md#Invoke-DeleteArchiveBucket) | **DELETE** /api/archives/buckets/{id} | Delete an Archive Bucket
[**Invoke-DeleteArchiveFile**](ArchivesApi.md#Invoke-DeleteArchiveFile) | **DELETE** /api/archives/files/{id} | Delete Archive File
[**Invoke-DeleteArchiveFileLink**](ArchivesApi.md#Invoke-DeleteArchiveFileLink) | **DELETE** /api/archives/files/{id}/links/{linkId} | Delete an Archive File Link
[**Get-ArchiveBucket**](ArchivesApi.md#Get-ArchiveBucket) | **GET** /api/archives/buckets/{id} | Get a Specific Archive Bucket
[**Get-ArchiveFile**](ArchivesApi.md#Get-ArchiveFile) | **GET** /api/archives/download/{bucket}/{filepath} | Download an Archive File
[**Get-ArchiveFileDetail**](ArchivesApi.md#Get-ArchiveFileDetail) | **GET** /api/archives/files/{id} | Get Archive File Details
[**Get-ArchiveFileLinks**](ArchivesApi.md#Get-ArchiveFileLinks) | **GET** /api/archives/files/{id}/links | Get Archive File Links
[**Invoke-ListArchiveBuckets**](ArchivesApi.md#Invoke-ListArchiveBuckets) | **GET** /api/archives/buckets | Get All Archive Buckets
[**Invoke-ListArchiveFiles**](ArchivesApi.md#Invoke-ListArchiveFiles) | **GET** /api/archives/buckets/{bucket}/files/{filepath} | Get All Archive Files
[**Update-ArchiveBucket**](ArchivesApi.md#Update-ArchiveBucket) | **PUT** /api/archives/buckets/{id} | Update an Archive Bucket


<a id="Add-ArchiveBucket"></a>
# **Add-ArchiveBucket**
> AddArchiveBucket200Response Add-ArchiveBucket<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddArchiveBucketRequest] <PSCustomObject><br>

Create an Archive Bucket

Create an Archive Bucket

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$ArchiveBucketCreateStorageProvider = Initialize-ArchiveBucketCreateStorageProvider -Id 0
$UpdateBlueprintPermissionsRequestResourcePermissionSitesInner = Initialize-UpdateBlueprintPermissionsRequestResourcePermissionSitesInner -Id 0
$ArchiveBucketCreate = Initialize-ArchiveBucketCreate -Name "MyName" -Description "MyDescription" -StorageProvider $ArchiveBucketCreateStorageProvider -Visibility "public" -IsPublic $false -Accounts $UpdateBlueprintPermissionsRequestResourcePermissionSitesInner

$AddArchiveBucketRequest = Initialize-AddArchiveBucketRequest -ArchiveBucket $ArchiveBucketCreate # AddArchiveBucketRequest |  (optional)

# Create an Archive Bucket
try {
    $Result = Add-ArchiveBucket -AddArchiveBucketRequest $AddArchiveBucketRequest
} catch {
    Write-Host ("Exception occurred when calling Add-ArchiveBucket: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AddArchiveBucketRequest** | [**AddArchiveBucketRequest**](AddArchiveBucketRequest.md)|  | [optional] 

### Return type

[**AddArchiveBucket200Response**](AddArchiveBucket200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Add-ArchiveFile"></a>
# **Add-ArchiveFile**
> AddArchiveFile200Response Add-ArchiveFile<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Bucket] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Filepath] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Filename] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-File] <System.IO.FileInfo><br>

Upload Archive File

Upload a file to the specified archive bucket and file path.  This will overwrite the file if it already exists. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Bucket = "MyBucket" # String | Bucket name
$Filepath = "/config/environments/" # String | The path to to search for files under. Default is the root directory /. (default to "/")
$Filename = "/path/to/file" # String | Specify a filename for archive file. The base filename of the uploaded file is the default. (optional)
$File =  # System.IO.FileInfo |  (optional)

# Upload Archive File
try {
    $Result = Add-ArchiveFile -Bucket $Bucket -Filepath $Filepath -Filename $Filename -File $File
} catch {
    Write-Host ("Exception occurred when calling Add-ArchiveFile: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Bucket** | **String**| Bucket name | 
 **Filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]
 **Filename** | **String**| Specify a filename for archive file. The base filename of the uploaded file is the default. | [optional] 
 **File** | **System.IO.FileInfo****System.IO.FileInfo**|  | [optional] 

### Return type

[**AddArchiveFile200Response**](AddArchiveFile200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Add-ArchiveFileLink"></a>
# **Add-ArchiveFileLink**
> AddArchiveFileLink200Response Add-ArchiveFileLink<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ExpireSeconds] <System.Nullable[Int64]><br>

Create an Archive File Link

This returns a secret token that can be used to download the file via a public url, without any other authentication or authorization. File links can be set to expire after a certain amount of time.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$ExpireSeconds = 600 # Int64 | Time to live in seconds. 0 means do not expire. (optional) (default to 0)

# Create an Archive File Link
try {
    $Result = Add-ArchiveFileLink -Id $Id -ExpireSeconds $ExpireSeconds
} catch {
    Write-Host ("Exception occurred when calling Add-ArchiveFileLink: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **ExpireSeconds** | **Int64**| Time to live in seconds. 0 means do not expire. | [optional] [default to 0]

### Return type

[**AddArchiveFileLink200Response**](AddArchiveFileLink200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-DeleteArchiveBucket"></a>
# **Invoke-DeleteArchiveBucket**
> Model200Success Invoke-DeleteArchiveBucket<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Delete an Archive Bucket

Will delete an archive bucket from the system and make it no longer usable.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Delete an Archive Bucket
try {
    $Result = Invoke-DeleteArchiveBucket -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Invoke-DeleteArchiveBucket: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
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

<a id="Invoke-DeleteArchiveFile"></a>
# **Invoke-DeleteArchiveFile**
> Model200Success Invoke-DeleteArchiveFile<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Delete Archive File

Permanently delete a file or directory.  Deleting a directory will also delete all the files under it. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Delete Archive File
try {
    $Result = Invoke-DeleteArchiveFile -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Invoke-DeleteArchiveFile: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
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

<a id="Invoke-DeleteArchiveFileLink"></a>
# **Invoke-DeleteArchiveFileLink**
> Model200Success Invoke-DeleteArchiveFileLink<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-LinkId] <Int64><br>

Delete an Archive File Link

This will delete the link from the system, so it can no longer be used.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$LinkId = 789 # Int64 | The ID of the archive file link.

# Delete an Archive File Link
try {
    $Result = Invoke-DeleteArchiveFileLink -Id $Id -LinkId $LinkId
} catch {
    Write-Host ("Exception occurred when calling Invoke-DeleteArchiveFileLink: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **LinkId** | **Int64**| The ID of the archive file link. | 

### Return type

[**Model200Success**](Model200Success.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-ArchiveBucket"></a>
# **Get-ArchiveBucket**
> GetArchiveBucket200Response Get-ArchiveBucket<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Get a Specific Archive Bucket

This endpoint retrieves a specific archive bucket.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Get a Specific Archive Bucket
try {
    $Result = Get-ArchiveBucket -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-ArchiveBucket: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetArchiveBucket200Response**](GetArchiveBucket200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-ArchiveFile"></a>
# **Get-ArchiveFile**
> void Get-ArchiveFile<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Bucket] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Filepath] <String><br>

Download an Archive File

Download the file as an authorized user with access to the bucket.  Downloading a directory will return a .zip file containing all files under it. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Bucket = "MyBucket" # String | Bucket name
$Filepath = "/config/environments/" # String | The path to to search for files under. Default is the root directory /. (default to "/")

# Download an Archive File
try {
    $Result = Get-ArchiveFile -Bucket $Bucket -Filepath $Filepath
} catch {
    Write-Host ("Exception occurred when calling Get-ArchiveFile: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Bucket** | **String**| Bucket name | 
 **Filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]

### Return type

void (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-ArchiveFileDetail"></a>
# **Get-ArchiveFileDetail**
> GetArchiveFileDetail200Response Get-ArchiveFileDetail<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Get Archive File Details

Get details about a specific archive file.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Get Archive File Details
try {
    $Result = Get-ArchiveFileDetail -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-ArchiveFileDetail: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetArchiveFileDetail200Response**](GetArchiveFileDetail200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-ArchiveFileLinks"></a>
# **Get-ArchiveFileLinks**
> GetArchiveFileLinks200Response Get-ArchiveFileLinks<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Get Archive File Links

This endpoint retrieves the links that have been created for the specified file.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Get Archive File Links
try {
    $Result = Get-ArchiveFileLinks -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-ArchiveFileLinks: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetArchiveFileLinks200Response**](GetArchiveFileLinks200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListArchiveBuckets"></a>
# **Invoke-ListArchiveBuckets**
> ListArchiveBuckets200Response Invoke-ListArchiveBuckets<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Name] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Phrase] <String><br>

Get All Archive Buckets

This endpoint retrieves all archive buckets associated with the account.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Name = "example-%" # String | Filter by name, wildcard may be specified as %, eg. example-% (optional)
$Phrase = "MyPhrase" # String | Search phrase for partial matches on name or description (optional)

# Get All Archive Buckets
try {
    $Result = Invoke-ListArchiveBuckets -Name $Name -Phrase $Phrase
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListArchiveBuckets: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **Phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

[**ListArchiveBuckets200Response**](ListArchiveBuckets200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListArchiveFiles"></a>
# **Invoke-ListArchiveFiles**
> ListArchiveFiles200Response Invoke-ListArchiveFiles<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Bucket] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Filepath] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Name] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Phrase] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-FullTree] <System.Nullable[Boolean]><br>

Get All Archive Files

This endpoint retrieves all files in an archive bucket under the specified `filePath`.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Bucket = "MyBucket" # String | Bucket name
$Filepath = "/config/environments/" # String | The path to to search for files under. Default is the root directory /. (default to "/")
$Name = "example-%" # String | Filter by name, wildcard may be specified as %, eg. example-% (optional)
$Phrase = "MyPhrase" # String | Search phrase for partial matches on name or description (optional)
$FullTree = $true # Boolean | Include files under sub directories too. This is always true when searching with phrase. (optional) (default to $false)

# Get All Archive Files
try {
    $Result = Invoke-ListArchiveFiles -Bucket $Bucket -Filepath $Filepath -Name $Name -Phrase $Phrase -FullTree $FullTree
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListArchiveFiles: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Bucket** | **String**| Bucket name | 
 **Filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]
 **Name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **Phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **FullTree** | **Boolean**| Include files under sub directories too. This is always true when searching with phrase. | [optional] [default to $false]

### Return type

[**ListArchiveFiles200Response**](ListArchiveFiles200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Update-ArchiveBucket"></a>
# **Update-ArchiveBucket**
> AddArchiveBucket200Response Update-ArchiveBucket<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-UpdateArchiveBucketRequest] <PSCustomObject><br>

Update an Archive Bucket

Update an Archive Bucket

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$UpdateBlueprintPermissionsRequestResourcePermissionSitesInner = Initialize-UpdateBlueprintPermissionsRequestResourcePermissionSitesInner -Id 0
$ArchiveBucketUpdate = Initialize-ArchiveBucketUpdate -Name "MyName" -Description "MyDescription" -Visibility "public" -IsPublic $false -Accounts $UpdateBlueprintPermissionsRequestResourcePermissionSitesInner

$UpdateArchiveBucketRequest = Initialize-UpdateArchiveBucketRequest -ArchiveBucket $ArchiveBucketUpdate # UpdateArchiveBucketRequest |  (optional)

# Update an Archive Bucket
try {
    $Result = Update-ArchiveBucket -Id $Id -UpdateArchiveBucketRequest $UpdateArchiveBucketRequest
} catch {
    Write-Host ("Exception occurred when calling Update-ArchiveBucket: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **UpdateArchiveBucketRequest** | [**UpdateArchiveBucketRequest**](UpdateArchiveBucketRequest.md)|  | [optional] 

### Return type

[**AddArchiveBucket200Response**](AddArchiveBucket200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

