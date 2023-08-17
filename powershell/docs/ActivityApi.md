# PSOpenAPITools.PSOpenAPITools/Api.ActivityApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Invoke-ListActivity**](ActivityApi.md#Invoke-ListActivity) | **GET** /api/activity | Retrieves Activity


<a id="Invoke-ListActivity"></a>
# **Invoke-ListActivity**
> ListActivity200Response Invoke-ListActivity<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Max] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Offset] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Sort] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Order] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Phrase] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Name] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-UserId] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-TenantId] <System.Nullable[Decimal]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Timeframe] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Start] <System.Nullable[System.DateTime]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-End] <System.Nullable[System.DateTime]><br>

Retrieves Activity

Retrieves activity. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Max = 789 # Int64 | Maximum number of records to return (optional) (default to 25)
$Offset = 789 # Int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
$Sort = "MySort" # String | Sort order, the name of the property to sort by (optional) (default to "name")
$Order = "asc" # String | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
$Phrase = "MyPhrase" # String | Search phrase for partial matches on name or description (optional)
$Name = "example-%" # String | Filter by name, wildcard may be specified as %, eg. example-% (optional)
$UserId = 6 # Int64 | Filter by User ID (optional)
$TenantId = 3 # Decimal | Filter by Tenant ID. Only available to the master account. (optional)
$Timeframe = "MyTimeframe" # String | Filter by a timeframe (ex - today, yesterday, week, month, 3months) (optional) (default to "month")
$Start = (Get-Date) # System.DateTime | Filter by activity on or after a date(time). Default is 1 month prior (optional)
$End = (Get-Date) # System.DateTime | Filter by activity on or before a date(time). Default is current date (optional)

# Retrieves Activity
try {
    $Result = Invoke-ListActivity -Max $Max -Offset $Offset -Sort $Sort -Order $Order -Phrase $Phrase -Name $Name -UserId $UserId -TenantId $TenantId -Timeframe $Timeframe -Start $Start -End $End
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListActivity: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Max** | **Int64**| Maximum number of records to return | [optional] [default to 25]
 **Offset** | **Int64**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **Sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **Order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &quot;asc&quot;]
 **Phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **Name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **UserId** | **Int64**| Filter by User ID | [optional] 
 **TenantId** | **Decimal**| Filter by Tenant ID. Only available to the master account. | [optional] 
 **Timeframe** | **String**| Filter by a timeframe (ex - today, yesterday, week, month, 3months) | [optional] [default to &quot;month&quot;]
 **Start** | **System.DateTime**| Filter by activity on or after a date(time). Default is 1 month prior | [optional] 
 **End** | **System.DateTime**| Filter by activity on or before a date(time). Default is current date | [optional] 

### Return type

[**ListActivity200Response**](ListActivity200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

