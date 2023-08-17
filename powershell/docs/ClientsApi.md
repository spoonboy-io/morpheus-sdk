# PSOpenAPITools.PSOpenAPITools/Api.ClientsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add-Client**](ClientsApi.md#Add-Client) | **POST** /api/clients | Create an Oauth Client
[**Get-Clients**](ClientsApi.md#Get-Clients) | **GET** /api/clients/{id} | Retrieves a Specific Oauth Client
[**Invoke-ListClients**](ClientsApi.md#Invoke-ListClients) | **GET** /api/clients | Get All Oauth Clients
[**Remove-Clients**](ClientsApi.md#Remove-Clients) | **DELETE** /api/clients/{id} | Deletes an Oauth Client
[**Update-Clients**](ClientsApi.md#Update-Clients) | **PUT** /api/clients/{id} | Updates an Oauth Client


<a id="Add-Client"></a>
# **Add-Client**
> AddClient200Response Add-Client<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AddClientRequest] <PSCustomObject><br>

Create an Oauth Client

Create a new Oauth Client.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$AddClientRequestClient = Initialize-AddClientRequestClient -ClientId "Test Client" -ClientSecret "thisIsaClientSecretKeyPhrase" -AccessTokenValiditySeconds 43200 -RefreshTokenValiditySeconds 43200
$AddClientRequest = Initialize-AddClientRequest -Client $AddClientRequestClient # AddClientRequest |  (optional)

# Create an Oauth Client
try {
    $Result = Add-Client -AddClientRequest $AddClientRequest
} catch {
    Write-Host ("Exception occurred when calling Add-Client: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AddClientRequest** | [**AddClientRequest**](AddClientRequest.md)|  | [optional] 

### Return type

[**AddClient200Response**](AddClient200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-Clients"></a>
# **Get-Clients**
> GetClients200Response Get-Clients<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Retrieves a Specific Oauth Client

Retrieves a specific oauth client. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Retrieves a Specific Oauth Client
try {
    $Result = Get-Clients -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Get-Clients: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 

### Return type

[**GetClients200Response**](GetClients200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListClients"></a>
# **Invoke-ListClients**
> ListClients200Response Invoke-ListClients<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Max] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Offset] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Sort] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Direction] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Phrase] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-ClientId] <String><br>

Get All Oauth Clients

This endpoint retrieves a paginated list of oauth clients.

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Max = 789 # Int64 | Maximum number of records to return (optional) (default to 25)
$Offset = 789 # Int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
$Sort = "MySort" # String | Sort order, the name of the property to sort by (optional) (default to "clientId")
$Direction = "asc" # String | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
$Phrase = "MyPhrase" # String | Search phrase for partial matches on clientId (optional)
$ClientId = "MyClientId" # String | Search phrase for partial matches on clientId (optional)

# Get All Oauth Clients
try {
    $Result = Invoke-ListClients -Max $Max -Offset $Offset -Sort $Sort -Direction $Direction -Phrase $Phrase -ClientId $ClientId
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListClients: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Max** | **Int64**| Maximum number of records to return | [optional] [default to 25]
 **Offset** | **Int64**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **Sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;clientId&quot;]
 **Direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &quot;asc&quot;]
 **Phrase** | **String**| Search phrase for partial matches on clientId | [optional] 
 **ClientId** | **String**| Search phrase for partial matches on clientId | [optional] 

### Return type

[**ListClients200Response**](ListClients200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Remove-Clients"></a>
# **Remove-Clients**
> Model200Success Remove-Clients<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>

Deletes an Oauth Client

Deletes a specified oauth client. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced

# Deletes an Oauth Client
try {
    $Result = Remove-Clients -Id $Id
} catch {
    Write-Host ("Exception occurred when calling Remove-Clients: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
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

<a id="Update-Clients"></a>
# **Update-Clients**
> AddClient200Response Update-Clients<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-UpdateClientsRequest] <PSCustomObject><br>

Updates an Oauth Client

Updates an oauth client. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$ClientUpdate = Initialize-ClientUpdate -ClientId "MyClientId" -AccessTokenValiditySeconds 0 -RefreshTokenValiditySeconds 0
$UpdateClientsRequest = Initialize-UpdateClientsRequest -Client $ClientUpdate # UpdateClientsRequest |  (optional)

# Updates an Oauth Client
try {
    $Result = Update-Clients -Id $Id -UpdateClientsRequest $UpdateClientsRequest
} catch {
    Write-Host ("Exception occurred when calling Update-Clients: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **UpdateClientsRequest** | [**UpdateClientsRequest**](UpdateClientsRequest.md)|  | [optional] 

### Return type

[**AddClient200Response**](AddClient200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

