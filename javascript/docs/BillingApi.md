# MorpheusApi.BillingApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getBillingAccount**](BillingApi.md#getBillingAccount) | **GET** /api/billing/account/{id} | This endpoint will retrieve a specific account by id if the user has permission to access it
[**getBillingInstancesIdentifier**](BillingApi.md#getBillingInstancesIdentifier) | **GET** /api/billing/instances/{identifier} | Retrieves billing information for an instance in the requestor&#39;s account. Use instanceUUID whenever possible.
[**getBillingServersIdentifier**](BillingApi.md#getBillingServersIdentifier) | **GET** /api/billing/servers/{identifier} | Retrieves billing information for a specific server (container host) in the requestor&#39;s account. Use refUUID whenever possible.
[**getBillingZoneIdentifier**](BillingApi.md#getBillingZoneIdentifier) | **GET** /api/billing/zones/{identifier} | Retrieves billing information for a specific zone in the requestor&#39;s account. Use zoneUUID whenever possible.
[**listBillingAccount**](BillingApi.md#listBillingAccount) | **GET** /api/billing/account | Retrieves billing information for the requesting user&#39;s account.
[**listBillingInstances**](BillingApi.md#listBillingInstances) | **GET** /api/billing/instances | Retrieves billing information for all instances on the requestor&#39;s account.
[**listBillingServers**](BillingApi.md#listBillingServers) | **GET** /api/billing/servers | Retrieves billing information for all servers (container hosts) on the requestor&#39;s account.
[**listBillingZone**](BillingApi.md#listBillingZone) | **GET** /api/billing/zones | Retrieves billing information for all zones on the requestor&#39;s account.



## getBillingAccount

> Object getBillingAccount(id, opts)

This endpoint will retrieve a specific account by id if the user has permission to access it

Will retrieve billing information for a specific tenant, if it is the current account or a sub account of the requesting user&#39;s account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BillingApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'includeUsages': true, // Boolean | Optional ability to suppress the usage records
  'maxUsages': 789, // Number | Optional ability to limit the usages returned
  'offsetUsages': 789, // Number | Optional ability to offset the usages returned, for use with maxUsages to paginate
  'includeComputeServers': true, // Boolean | Optional ability to exclude compute servers
  'includeInstances': true, // Boolean | Optional ability to exclude instances
  'includeDiscoveredServers': true, // Boolean | Optional ability to exclude discovered servers
  'includeLoadBalancers': true, // Boolean | Optional ability to exclude load balancers
  'includeVirtualImages': true, // Boolean | Optional ability to exclude virtual images
  'includeSnapshots': true // Boolean | Optional ability to exclude snapshots
};
apiInstance.getBillingAccount(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true]
 **maxUsages** | **Number**| Optional ability to limit the usages returned | [optional] 
 **offsetUsages** | **Number**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **includeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to true]
 **includeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to true]
 **includeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to true]
 **includeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to true]
 **includeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to true]
 **includeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to true]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getBillingInstancesIdentifier

> Object getBillingInstancesIdentifier(identifier, opts)

Retrieves billing information for an instance in the requestor&#39;s account. Use instanceUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BillingApi();
let identifier = "identifier_example"; // String | Morpheus UUID or ID of the Object being created or referenced
let opts = {
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'includeUsages': true, // Boolean | Optional ability to suppress the usage records
  'maxUsages': 789, // Number | Optional ability to limit the usages returned
  'offsetUsages': 789, // Number | Optional ability to offset the usages returned, for use with maxUsages to paginate
  'includeTenants': false, // Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user
  'accountId': 3 // Number | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
};
apiInstance.getBillingInstancesIdentifier(identifier, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **String**| Morpheus UUID or ID of the Object being created or referenced | 
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true]
 **maxUsages** | **Number**| Optional ability to limit the usages returned | [optional] 
 **offsetUsages** | **Number**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **includeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false]
 **accountId** | **Number**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getBillingServersIdentifier

> Object getBillingServersIdentifier(identifier, opts)

Retrieves billing information for a specific server (container host) in the requestor&#39;s account. Use refUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BillingApi();
let identifier = "identifier_example"; // String | Morpheus UUID or ID of the Object being created or referenced
let opts = {
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'includeUsages': true, // Boolean | Optional ability to suppress the usage records
  'maxUsages': 789, // Number | Optional ability to limit the usages returned
  'offsetUsages': 789, // Number | Optional ability to offset the usages returned, for use with maxUsages to paginate
  'includeTenants': false, // Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user
  'accountId': 3 // Number | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
};
apiInstance.getBillingServersIdentifier(identifier, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **String**| Morpheus UUID or ID of the Object being created or referenced | 
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true]
 **maxUsages** | **Number**| Optional ability to limit the usages returned | [optional] 
 **offsetUsages** | **Number**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **includeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false]
 **accountId** | **Number**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getBillingZoneIdentifier

> Object getBillingZoneIdentifier(identifier, opts)

Retrieves billing information for a specific zone in the requestor&#39;s account. Use zoneUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BillingApi();
let identifier = "identifier_example"; // String | Morpheus UUID or ID of the Object being created or referenced
let opts = {
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'includeUsages': true, // Boolean | Optional ability to suppress the usage records
  'maxUsages': 789, // Number | Optional ability to limit the usages returned
  'offsetUsages': 789, // Number | Optional ability to offset the usages returned, for use with maxUsages to paginate
  'includeComputeServers': true, // Boolean | Optional ability to exclude compute servers
  'includeInstances': true, // Boolean | Optional ability to exclude instances
  'includeDiscoveredServers': true, // Boolean | Optional ability to exclude discovered servers
  'includeLoadBalancers': true, // Boolean | Optional ability to exclude load balancers
  'includeVirtualImages': true, // Boolean | Optional ability to exclude virtual images
  'includeSnapshots': true // Boolean | Optional ability to exclude snapshots
};
apiInstance.getBillingZoneIdentifier(identifier, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **String**| Morpheus UUID or ID of the Object being created or referenced | 
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true]
 **maxUsages** | **Number**| Optional ability to limit the usages returned | [optional] 
 **offsetUsages** | **Number**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **includeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to true]
 **includeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to true]
 **includeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to true]
 **includeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to true]
 **includeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to true]
 **includeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to true]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listBillingAccount

> Object listBillingAccount(opts)

Retrieves billing information for the requesting user&#39;s account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BillingApi();
let opts = {
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'includeUsages': true, // Boolean | Optional ability to suppress the usage records
  'maxUsages': 789, // Number | Optional ability to limit the usages returned
  'offsetUsages': 789, // Number | Optional ability to offset the usages returned, for use with maxUsages to paginate
  'includeComputeServers': true, // Boolean | Optional ability to exclude compute servers
  'includeInstances': true, // Boolean | Optional ability to exclude instances
  'includeDiscoveredServers': true, // Boolean | Optional ability to exclude discovered servers
  'includeLoadBalancers': true, // Boolean | Optional ability to exclude load balancers
  'includeVirtualImages': true, // Boolean | Optional ability to exclude virtual images
  'includeSnapshots': true // Boolean | Optional ability to exclude snapshots
};
apiInstance.listBillingAccount(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true]
 **maxUsages** | **Number**| Optional ability to limit the usages returned | [optional] 
 **offsetUsages** | **Number**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **includeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to true]
 **includeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to true]
 **includeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to true]
 **includeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to true]
 **includeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to true]
 **includeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to true]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listBillingInstances

> Object listBillingInstances(opts)

Retrieves billing information for all instances on the requestor&#39;s account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BillingApi();
let opts = {
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'includeUsages': true, // Boolean | Optional ability to suppress the usage records
  'maxUsages': 789, // Number | Optional ability to limit the usages returned
  'offsetUsages': 789, // Number | Optional ability to offset the usages returned, for use with maxUsages to paginate
  'includeTenants': false, // Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user
  'accountId': 3 // Number | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
};
apiInstance.listBillingInstances(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true]
 **maxUsages** | **Number**| Optional ability to limit the usages returned | [optional] 
 **offsetUsages** | **Number**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **includeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false]
 **accountId** | **Number**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listBillingServers

> Object listBillingServers(opts)

Retrieves billing information for all servers (container hosts) on the requestor&#39;s account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BillingApi();
let opts = {
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'includeUsages': true, // Boolean | Optional ability to suppress the usage records
  'maxUsages': 789, // Number | Optional ability to limit the usages returned
  'offsetUsages': 789, // Number | Optional ability to offset the usages returned, for use with maxUsages to paginate
  'includeTenants': false, // Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user
  'accountId': 3 // Number | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
};
apiInstance.listBillingServers(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true]
 **maxUsages** | **Number**| Optional ability to limit the usages returned | [optional] 
 **offsetUsages** | **Number**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **includeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false]
 **accountId** | **Number**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listBillingZone

> Object listBillingZone(opts)

Retrieves billing information for all zones on the requestor&#39;s account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BillingApi();
let opts = {
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'includeUsages': true, // Boolean | Optional ability to suppress the usage records
  'maxUsages': 789, // Number | Optional ability to limit the usages returned
  'offsetUsages': 789, // Number | Optional ability to offset the usages returned, for use with maxUsages to paginate
  'includeComputeServers': true, // Boolean | Optional ability to exclude compute servers
  'includeInstances': true, // Boolean | Optional ability to exclude instances
  'includeDiscoveredServers': true, // Boolean | Optional ability to exclude discovered servers
  'includeLoadBalancers': true, // Boolean | Optional ability to exclude load balancers
  'includeVirtualImages': true, // Boolean | Optional ability to exclude virtual images
  'includeSnapshots': true // Boolean | Optional ability to exclude snapshots
};
apiInstance.listBillingZone(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true]
 **maxUsages** | **Number**| Optional ability to limit the usages returned | [optional] 
 **offsetUsages** | **Number**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] 
 **includeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to true]
 **includeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to true]
 **includeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to true]
 **includeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to true]
 **includeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to true]
 **includeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to true]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

