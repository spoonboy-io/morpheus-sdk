# PSOpenAPITools.PSOpenAPITools/Api.BillingApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get-BillingAccount**](BillingApi.md#Get-BillingAccount) | **GET** /api/billing/account/{id} | This endpoint will retrieve a specific account by id if the user has permission to access it
[**Get-BillingInstancesIdentifier**](BillingApi.md#Get-BillingInstancesIdentifier) | **GET** /api/billing/instances/{identifier} | Retrieves billing information for an instance in the requestor&#39;s account. Use instanceUUID whenever possible.
[**Get-BillingServersIdentifier**](BillingApi.md#Get-BillingServersIdentifier) | **GET** /api/billing/servers/{identifier} | Retrieves billing information for a specific server (container host) in the requestor&#39;s account. Use refUUID whenever possible.
[**Get-BillingZoneIdentifier**](BillingApi.md#Get-BillingZoneIdentifier) | **GET** /api/billing/zones/{identifier} | Retrieves billing information for a specific zone in the requestor&#39;s account. Use zoneUUID whenever possible.
[**Invoke-ListBillingAccount**](BillingApi.md#Invoke-ListBillingAccount) | **GET** /api/billing/account | Retrieves billing information for the requesting user&#39;s account.
[**Invoke-ListBillingInstances**](BillingApi.md#Invoke-ListBillingInstances) | **GET** /api/billing/instances | Retrieves billing information for all instances on the requestor&#39;s account.
[**Invoke-ListBillingServers**](BillingApi.md#Invoke-ListBillingServers) | **GET** /api/billing/servers | Retrieves billing information for all servers (container hosts) on the requestor&#39;s account.
[**Invoke-ListBillingZone**](BillingApi.md#Invoke-ListBillingZone) | **GET** /api/billing/zones | Retrieves billing information for all zones on the requestor&#39;s account.


<a id="Get-BillingAccount"></a>
# **Get-BillingAccount**
> ListBillingAccount200Response Get-BillingAccount<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Id] <Int64><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-StartDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-EndDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeUsages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-MaxUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-OffsetUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeComputeServers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeInstances] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeDiscoveredServers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeLoadBalancers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeVirtualImages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeSnapshots] <System.Nullable[Boolean]><br>

This endpoint will retrieve a specific account by id if the user has permission to access it

Will retrieve billing information for a specific tenant, if it is the current account or a sub account of the requesting user's account. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Id = 1 # Int64 | Morpheus ID of the Object being referenced
$StartDate = "2019-01-01" # String | Filter by startDate greater than or equal to a specified date (optional)
$EndDate = "2020-01-01" # String | Filter by endDate less than or equal to a specified date (optional)
$IncludeUsages = $true # Boolean | Optional ability to suppress the usage records (optional) (default to $true)
$MaxUsages = 789 # Int64 | Optional ability to limit the usages returned (optional)
$OffsetUsages = 789 # Int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
$IncludeComputeServers = $true # Boolean | Optional ability to exclude compute servers (optional) (default to $true)
$IncludeInstances = $true # Boolean | Optional ability to exclude instances (optional) (default to $true)
$IncludeDiscoveredServers = $true # Boolean | Optional ability to exclude discovered servers (optional) (default to $true)
$IncludeLoadBalancers = $true # Boolean | Optional ability to exclude load balancers (optional) (default to $true)
$IncludeVirtualImages = $true # Boolean | Optional ability to exclude virtual images (optional) (default to $true)
$IncludeSnapshots = $true # Boolean | Optional ability to exclude snapshots (optional) (default to $true)

# This endpoint will retrieve a specific account by id if the user has permission to access it
try {
    $Result = Get-BillingAccount -Id $Id -StartDate $StartDate -EndDate $EndDate -IncludeUsages $IncludeUsages -MaxUsages $MaxUsages -OffsetUsages $OffsetUsages -IncludeComputeServers $IncludeComputeServers -IncludeInstances $IncludeInstances -IncludeDiscoveredServers $IncludeDiscoveredServers -IncludeLoadBalancers $IncludeLoadBalancers -IncludeVirtualImages $IncludeVirtualImages -IncludeSnapshots $IncludeSnapshots
} catch {
    Write-Host ("Exception occurred when calling Get-BillingAccount: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Id** | **Int64**| Morpheus ID of the Object being referenced | 
 **StartDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **EndDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **IncludeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to $true]
 **MaxUsages** | **Int64**| Optional ability to limit the usages returned | [optional] 
 **OffsetUsages** | **Int64**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **IncludeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to $true]
 **IncludeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to $true]
 **IncludeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to $true]
 **IncludeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to $true]
 **IncludeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to $true]
 **IncludeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to $true]

### Return type

[**ListBillingAccount200Response**](ListBillingAccount200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-BillingInstancesIdentifier"></a>
# **Get-BillingInstancesIdentifier**
> GetBillingInstancesIdentifier200Response Get-BillingInstancesIdentifier<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Identifier] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-StartDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-EndDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeUsages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-MaxUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-OffsetUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeTenants] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AccountId] <System.Nullable[Int64]><br>

Retrieves billing information for an instance in the requestor's account. Use instanceUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Identifier = "MyIdentifier" # String | Morpheus UUID or ID of the Object being created or referenced
$StartDate = "2019-01-01" # String | Filter by startDate greater than or equal to a specified date (optional)
$EndDate = "2020-01-01" # String | Filter by endDate less than or equal to a specified date (optional)
$IncludeUsages = $true # Boolean | Optional ability to suppress the usage records (optional) (default to $true)
$MaxUsages = 789 # Int64 | Optional ability to limit the usages returned (optional)
$OffsetUsages = 789 # Int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
$IncludeTenants = $true # Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) (default to $false)
$AccountId = 3 # Int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

# Retrieves billing information for an instance in the requestor's account. Use instanceUUID whenever possible.
try {
    $Result = Get-BillingInstancesIdentifier -Identifier $Identifier -StartDate $StartDate -EndDate $EndDate -IncludeUsages $IncludeUsages -MaxUsages $MaxUsages -OffsetUsages $OffsetUsages -IncludeTenants $IncludeTenants -AccountId $AccountId
} catch {
    Write-Host ("Exception occurred when calling Get-BillingInstancesIdentifier: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Identifier** | **String**| Morpheus UUID or ID of the Object being created or referenced | 
 **StartDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **EndDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **IncludeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to $true]
 **MaxUsages** | **Int64**| Optional ability to limit the usages returned | [optional] 
 **OffsetUsages** | **Int64**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **IncludeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to $false]
 **AccountId** | **Int64**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 

### Return type

[**GetBillingInstancesIdentifier200Response**](GetBillingInstancesIdentifier200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-BillingServersIdentifier"></a>
# **Get-BillingServersIdentifier**
> GetBillingServersIdentifier200Response Get-BillingServersIdentifier<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Identifier] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-StartDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-EndDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeUsages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-MaxUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-OffsetUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeTenants] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AccountId] <System.Nullable[Int64]><br>

Retrieves billing information for a specific server (container host) in the requestor's account. Use refUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Identifier = "MyIdentifier" # String | Morpheus UUID or ID of the Object being created or referenced
$StartDate = "2019-01-01" # String | Filter by startDate greater than or equal to a specified date (optional)
$EndDate = "2020-01-01" # String | Filter by endDate less than or equal to a specified date (optional)
$IncludeUsages = $true # Boolean | Optional ability to suppress the usage records (optional) (default to $true)
$MaxUsages = 789 # Int64 | Optional ability to limit the usages returned (optional)
$OffsetUsages = 789 # Int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
$IncludeTenants = $true # Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) (default to $false)
$AccountId = 3 # Int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

# Retrieves billing information for a specific server (container host) in the requestor's account. Use refUUID whenever possible.
try {
    $Result = Get-BillingServersIdentifier -Identifier $Identifier -StartDate $StartDate -EndDate $EndDate -IncludeUsages $IncludeUsages -MaxUsages $MaxUsages -OffsetUsages $OffsetUsages -IncludeTenants $IncludeTenants -AccountId $AccountId
} catch {
    Write-Host ("Exception occurred when calling Get-BillingServersIdentifier: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Identifier** | **String**| Morpheus UUID or ID of the Object being created or referenced | 
 **StartDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **EndDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **IncludeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to $true]
 **MaxUsages** | **Int64**| Optional ability to limit the usages returned | [optional] 
 **OffsetUsages** | **Int64**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **IncludeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to $false]
 **AccountId** | **Int64**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 

### Return type

[**GetBillingServersIdentifier200Response**](GetBillingServersIdentifier200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Get-BillingZoneIdentifier"></a>
# **Get-BillingZoneIdentifier**
> GetBillingZoneIdentifier200Response Get-BillingZoneIdentifier<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-Identifier] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-StartDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-EndDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeUsages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-MaxUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-OffsetUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeComputeServers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeInstances] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeDiscoveredServers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeLoadBalancers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeVirtualImages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeSnapshots] <System.Nullable[Boolean]><br>

Retrieves billing information for a specific zone in the requestor's account. Use zoneUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$Identifier = "MyIdentifier" # String | Morpheus UUID or ID of the Object being created or referenced
$StartDate = "2019-01-01" # String | Filter by startDate greater than or equal to a specified date (optional)
$EndDate = "2020-01-01" # String | Filter by endDate less than or equal to a specified date (optional)
$IncludeUsages = $true # Boolean | Optional ability to suppress the usage records (optional) (default to $true)
$MaxUsages = 789 # Int64 | Optional ability to limit the usages returned (optional)
$OffsetUsages = 789 # Int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
$IncludeComputeServers = $true # Boolean | Optional ability to exclude compute servers (optional) (default to $true)
$IncludeInstances = $true # Boolean | Optional ability to exclude instances (optional) (default to $true)
$IncludeDiscoveredServers = $true # Boolean | Optional ability to exclude discovered servers (optional) (default to $true)
$IncludeLoadBalancers = $true # Boolean | Optional ability to exclude load balancers (optional) (default to $true)
$IncludeVirtualImages = $true # Boolean | Optional ability to exclude virtual images (optional) (default to $true)
$IncludeSnapshots = $true # Boolean | Optional ability to exclude snapshots (optional) (default to $true)

# Retrieves billing information for a specific zone in the requestor's account. Use zoneUUID whenever possible.
try {
    $Result = Get-BillingZoneIdentifier -Identifier $Identifier -StartDate $StartDate -EndDate $EndDate -IncludeUsages $IncludeUsages -MaxUsages $MaxUsages -OffsetUsages $OffsetUsages -IncludeComputeServers $IncludeComputeServers -IncludeInstances $IncludeInstances -IncludeDiscoveredServers $IncludeDiscoveredServers -IncludeLoadBalancers $IncludeLoadBalancers -IncludeVirtualImages $IncludeVirtualImages -IncludeSnapshots $IncludeSnapshots
} catch {
    Write-Host ("Exception occurred when calling Get-BillingZoneIdentifier: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Identifier** | **String**| Morpheus UUID or ID of the Object being created or referenced | 
 **StartDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **EndDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **IncludeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to $true]
 **MaxUsages** | **Int64**| Optional ability to limit the usages returned | [optional] 
 **OffsetUsages** | **Int64**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **IncludeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to $true]
 **IncludeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to $true]
 **IncludeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to $true]
 **IncludeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to $true]
 **IncludeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to $true]
 **IncludeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to $true]

### Return type

[**GetBillingZoneIdentifier200Response**](GetBillingZoneIdentifier200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListBillingAccount"></a>
# **Invoke-ListBillingAccount**
> ListBillingAccount200Response Invoke-ListBillingAccount<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-StartDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-EndDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeUsages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-MaxUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-OffsetUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeComputeServers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeInstances] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeDiscoveredServers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeLoadBalancers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeVirtualImages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeSnapshots] <System.Nullable[Boolean]><br>

Retrieves billing information for the requesting user's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$StartDate = "2019-01-01" # String | Filter by startDate greater than or equal to a specified date (optional)
$EndDate = "2020-01-01" # String | Filter by endDate less than or equal to a specified date (optional)
$IncludeUsages = $true # Boolean | Optional ability to suppress the usage records (optional) (default to $true)
$MaxUsages = 789 # Int64 | Optional ability to limit the usages returned (optional)
$OffsetUsages = 789 # Int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
$IncludeComputeServers = $true # Boolean | Optional ability to exclude compute servers (optional) (default to $true)
$IncludeInstances = $true # Boolean | Optional ability to exclude instances (optional) (default to $true)
$IncludeDiscoveredServers = $true # Boolean | Optional ability to exclude discovered servers (optional) (default to $true)
$IncludeLoadBalancers = $true # Boolean | Optional ability to exclude load balancers (optional) (default to $true)
$IncludeVirtualImages = $true # Boolean | Optional ability to exclude virtual images (optional) (default to $true)
$IncludeSnapshots = $true # Boolean | Optional ability to exclude snapshots (optional) (default to $true)

# Retrieves billing information for the requesting user's account.
try {
    $Result = Invoke-ListBillingAccount -StartDate $StartDate -EndDate $EndDate -IncludeUsages $IncludeUsages -MaxUsages $MaxUsages -OffsetUsages $OffsetUsages -IncludeComputeServers $IncludeComputeServers -IncludeInstances $IncludeInstances -IncludeDiscoveredServers $IncludeDiscoveredServers -IncludeLoadBalancers $IncludeLoadBalancers -IncludeVirtualImages $IncludeVirtualImages -IncludeSnapshots $IncludeSnapshots
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListBillingAccount: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **StartDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **EndDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **IncludeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to $true]
 **MaxUsages** | **Int64**| Optional ability to limit the usages returned | [optional] 
 **OffsetUsages** | **Int64**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **IncludeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to $true]
 **IncludeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to $true]
 **IncludeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to $true]
 **IncludeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to $true]
 **IncludeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to $true]
 **IncludeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to $true]

### Return type

[**ListBillingAccount200Response**](ListBillingAccount200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListBillingInstances"></a>
# **Invoke-ListBillingInstances**
> ListBillingInstances200Response Invoke-ListBillingInstances<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-StartDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-EndDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeUsages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-MaxUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-OffsetUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeTenants] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AccountId] <System.Nullable[Int64]><br>

Retrieves billing information for all instances on the requestor's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$StartDate = "2019-01-01" # String | Filter by startDate greater than or equal to a specified date (optional)
$EndDate = "2020-01-01" # String | Filter by endDate less than or equal to a specified date (optional)
$IncludeUsages = $true # Boolean | Optional ability to suppress the usage records (optional) (default to $true)
$MaxUsages = 789 # Int64 | Optional ability to limit the usages returned (optional)
$OffsetUsages = 789 # Int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
$IncludeTenants = $true # Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) (default to $false)
$AccountId = 3 # Int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

# Retrieves billing information for all instances on the requestor's account.
try {
    $Result = Invoke-ListBillingInstances -StartDate $StartDate -EndDate $EndDate -IncludeUsages $IncludeUsages -MaxUsages $MaxUsages -OffsetUsages $OffsetUsages -IncludeTenants $IncludeTenants -AccountId $AccountId
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListBillingInstances: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **StartDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **EndDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **IncludeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to $true]
 **MaxUsages** | **Int64**| Optional ability to limit the usages returned | [optional] 
 **OffsetUsages** | **Int64**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **IncludeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to $false]
 **AccountId** | **Int64**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 

### Return type

[**ListBillingInstances200Response**](ListBillingInstances200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListBillingServers"></a>
# **Invoke-ListBillingServers**
> ListBillingServers200Response Invoke-ListBillingServers<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-StartDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-EndDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeUsages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-MaxUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-OffsetUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeTenants] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-AccountId] <System.Nullable[Int64]><br>

Retrieves billing information for all servers (container hosts) on the requestor's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$StartDate = "2019-01-01" # String | Filter by startDate greater than or equal to a specified date (optional)
$EndDate = "2020-01-01" # String | Filter by endDate less than or equal to a specified date (optional)
$IncludeUsages = $true # Boolean | Optional ability to suppress the usage records (optional) (default to $true)
$MaxUsages = 789 # Int64 | Optional ability to limit the usages returned (optional)
$OffsetUsages = 789 # Int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
$IncludeTenants = $true # Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user (optional) (default to $false)
$AccountId = 3 # Int64 | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)

# Retrieves billing information for all servers (container hosts) on the requestor's account.
try {
    $Result = Invoke-ListBillingServers -StartDate $StartDate -EndDate $EndDate -IncludeUsages $IncludeUsages -MaxUsages $MaxUsages -OffsetUsages $OffsetUsages -IncludeTenants $IncludeTenants -AccountId $AccountId
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListBillingServers: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **StartDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **EndDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **IncludeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to $true]
 **MaxUsages** | **Int64**| Optional ability to limit the usages returned | [optional] 
 **OffsetUsages** | **Int64**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **IncludeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to $false]
 **AccountId** | **Int64**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 

### Return type

[**ListBillingServers200Response**](ListBillingServers200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a id="Invoke-ListBillingZone"></a>
# **Invoke-ListBillingZone**
> ListBillingZone200Response Invoke-ListBillingZone<br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-StartDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-EndDate] <String><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeUsages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-MaxUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-OffsetUsages] <System.Nullable[Int64]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeComputeServers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeInstances] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeDiscoveredServers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeLoadBalancers] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeVirtualImages] <System.Nullable[Boolean]><br>
> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[-IncludeSnapshots] <System.Nullable[Boolean]><br>

Retrieves billing information for all zones on the requestor's account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```powershell
# general setting of the PowerShell module, e.g. base URL, authentication, etc
$Configuration = Get-Configuration

$StartDate = "2019-01-01" # String | Filter by startDate greater than or equal to a specified date (optional)
$EndDate = "2020-01-01" # String | Filter by endDate less than or equal to a specified date (optional)
$IncludeUsages = $true # Boolean | Optional ability to suppress the usage records (optional) (default to $true)
$MaxUsages = 789 # Int64 | Optional ability to limit the usages returned (optional)
$OffsetUsages = 789 # Int64 | Optional ability to offset the usages returned, for use with maxUsages to paginate (optional)
$IncludeComputeServers = $true # Boolean | Optional ability to exclude compute servers (optional) (default to $true)
$IncludeInstances = $true # Boolean | Optional ability to exclude instances (optional) (default to $true)
$IncludeDiscoveredServers = $true # Boolean | Optional ability to exclude discovered servers (optional) (default to $true)
$IncludeLoadBalancers = $true # Boolean | Optional ability to exclude load balancers (optional) (default to $true)
$IncludeVirtualImages = $true # Boolean | Optional ability to exclude virtual images (optional) (default to $true)
$IncludeSnapshots = $true # Boolean | Optional ability to exclude snapshots (optional) (default to $true)

# Retrieves billing information for all zones on the requestor's account.
try {
    $Result = Invoke-ListBillingZone -StartDate $StartDate -EndDate $EndDate -IncludeUsages $IncludeUsages -MaxUsages $MaxUsages -OffsetUsages $OffsetUsages -IncludeComputeServers $IncludeComputeServers -IncludeInstances $IncludeInstances -IncludeDiscoveredServers $IncludeDiscoveredServers -IncludeLoadBalancers $IncludeLoadBalancers -IncludeVirtualImages $IncludeVirtualImages -IncludeSnapshots $IncludeSnapshots
} catch {
    Write-Host ("Exception occurred when calling Invoke-ListBillingZone: {0}" -f ($_.ErrorDetails | ConvertFrom-Json))
    Write-Host ("Response headers: {0}" -f ($_.Exception.Response.Headers | ConvertTo-Json))
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **StartDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **EndDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **IncludeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to $true]
 **MaxUsages** | **Int64**| Optional ability to limit the usages returned | [optional] 
 **OffsetUsages** | **Int64**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **IncludeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to $true]
 **IncludeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to $true]
 **IncludeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to $true]
 **IncludeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to $true]
 **IncludeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to $true]
 **IncludeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to $true]

### Return type

[**ListBillingZone200Response**](ListBillingZone200Response.md) (PSCustomObject)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

