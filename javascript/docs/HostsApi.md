# MorpheusApi.HostsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getHost**](HostsApi.md#getHost) | **GET** /api/servers/{id} | Get a Specific Host
[**getHostSnpshots**](HostsApi.md#getHostSnpshots) | **GET** /api/servers/{id}/snapshots | Get list of snapshots for a Host
[**getHostType**](HostsApi.md#getHostType) | **GET** /api/server-types/{id} | Get a Specific Host Type
[**getWikiServer**](HostsApi.md#getWikiServer) | **GET** /api/servers/{id}/wiki | Retrieves a Server Wiki Page
[**listHostTypes**](HostsApi.md#listHostTypes) | **GET** /api/server-types | Host Types
[**listHosts**](HostsApi.md#listHosts) | **GET** /api/servers | Get All Hosts
[**listServerServicePlans**](HostsApi.md#listServerServicePlans) | **GET** /api/servers/service-plans | Get Available Service Plans for a Host
[**removeHost**](HostsApi.md#removeHost) | **DELETE** /api/servers/{id} | Delete a Host
[**restartHost**](HostsApi.md#restartHost) | **PUT** /api/servers/{id}/restart | Restart a Host
[**startHost**](HostsApi.md#startHost) | **PUT** /api/servers/{id}/start | Start a Host
[**stopHost**](HostsApi.md#stopHost) | **PUT** /api/servers/{id}/stop | Stop a Host
[**updateHost**](HostsApi.md#updateHost) | **PUT** /api/servers/{id} | Updating a Host
[**updateHostAssignTenant**](HostsApi.md#updateHostAssignTenant) | **PUT** /api/servers/{id}/assign-account | Assign To Tenant
[**updateHostCloud**](HostsApi.md#updateHostCloud) | **PUT** /api/servers/change-cloud | Change Server Cloud
[**updateHostExecuteWorkflow**](HostsApi.md#updateHostExecuteWorkflow) | **PUT** /api/servers/{id}/workflow | Run Workflow on a Host
[**updateHostInstallAgent**](HostsApi.md#updateHostInstallAgent) | **PUT** /api/servers/{id}/install-agent | Install Agent
[**updateHostManaged**](HostsApi.md#updateHostManaged) | **PUT** /api/servers/{id}/make-managed | Convert To Managed
[**updateHostResize**](HostsApi.md#updateHostResize) | **PUT** /api/servers/{id}/resize | Resize a Host
[**updateHostUpgradeAgent**](HostsApi.md#updateHostUpgradeAgent) | **PUT** /api/servers/{id}/upgrade | Upgrade Agent
[**updateServerNetworkInterface**](HostsApi.md#updateServerNetworkInterface) | **PUT** /api/servers/{id}/networkInterfaces/{networkInterfaceId} | Updating a label for a Server&#39;s Network
[**updateWikiServer**](HostsApi.md#updateWikiServer) | **PUT** /api/servers/{id}/wiki | Update a Server Wiki Page



## getHost

> InlineResponse200137 getHost(id)

Get a Specific Host

This endpoint retrieves a specific host.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getHost(id, (error, data, response) => {
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

### Return type

[**InlineResponse200137**](InlineResponse200137.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getHostSnpshots

> InlineResponse200138 getHostSnpshots(id)

Get list of snapshots for a Host

Get list of snapshots for a Host

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getHostSnpshots(id, (error, data, response) => {
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

### Return type

[**InlineResponse200138**](InlineResponse200138.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getHostType

> InlineResponse20050 getHostType(id)

Get a Specific Host Type

This endpoint will retrieve a specific host type by id

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getHostType(id, (error, data, response) => {
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

### Return type

[**InlineResponse20050**](InlineResponse20050.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getWikiServer

> InlineResponse200168 getWikiServer(id)

Retrieves a Server Wiki Page

This endpoint retrieves a server Wiki page. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getWikiServer(id, (error, data, response) => {
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

### Return type

[**InlineResponse200168**](InlineResponse200168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listHostTypes

> Object listHostTypes(opts)

Host Types

Fetch a paginated list of available host types. This returns the configuration options for each type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'provisionType': "provisionType_example", // String | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings. 
  'zoneType': "zoneType_example", // String | Filter by Cloud Type code.
  'creatable': true // Boolean | Filter by creatable flag. This is whether or not it can be provisioned.
};
apiInstance.listHostTypes(opts, (error, data, response) => {
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
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **provisionType** | **String**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional] 
 **zoneType** | **String**| Filter by Cloud Type code. | [optional] 
 **creatable** | **Boolean**| Filter by creatable flag. This is whether or not it can be provisioned. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listHosts

> Object listHosts(opts)

Get All Hosts

This endpoint retrieves a paginated list of hosts.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'zoneId': 3, // Number | The Zone ID for Filtering
  'siteId': 7, // Number | The Site ID for Filtering
  'clusterId': 135, // Number | The Cluster ID(s) for filtering. Accepts multiple values.
  'managed': false, // Boolean | Filter by managed (true) or unmanaged (false)
  'serverType': vmwareHypervisor, // String | Filter by server type code
  'powerState': off, // String | Filter by power status
  'ip': 192.168.1.45, // String | Filter by IP address
  'vm': false, // Boolean | Filter to show only Virtual Machines (true)
  'vmHypervisor': false, // Boolean | Filter to show only VM Hypervisors (true)
  'bareMetalHost': false, // Boolean | Filter to show only Baremetal Servers
  'status': "status_example", // String | Filter by status
  'agentInstalled': true, // Boolean | Filter by agent installed (true)
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'lastUpdated': 2019-03-06T17:52:29+0000, // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
  'createdBy': 63, // Number | The User ID for Filtering
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example", // String | Filter by label(s), matches records that contain all of the specified labels
  'tags': tags.env=qa&tags.env=test, // String | Filter by tags (metadata). This allows filtering by a tag name and value(s) 
  'metadata': tags.env=qa&tags.env=test, // String | Alias for tags
  'uuid': "uuid_example", // String | Filter by UUID
  'externalId': "externalId_example", // String | Filter by External ID
  'internalId': "internalId_example", // String | Filter by Internal ID
  'externalUniquelId': "externalUniquelId_example" // String | Filter by External Unique ID
};
apiInstance.listHosts(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **zoneId** | **Number**| The Zone ID for Filtering | [optional] 
 **siteId** | **Number**| The Site ID for Filtering | [optional] 
 **clusterId** | **Number**| The Cluster ID(s) for filtering. Accepts multiple values. | [optional] 
 **managed** | **Boolean**| Filter by managed (true) or unmanaged (false) | [optional] 
 **serverType** | **String**| Filter by server type code | [optional] 
 **powerState** | **String**| Filter by power status | [optional] 
 **ip** | **String**| Filter by IP address | [optional] 
 **vm** | **Boolean**| Filter to show only Virtual Machines (true) | [optional] 
 **vmHypervisor** | **Boolean**| Filter to show only VM Hypervisors (true) | [optional] 
 **bareMetalHost** | **Boolean**| Filter to show only Baremetal Servers | [optional] 
 **status** | **String**| Filter by status | [optional] 
 **agentInstalled** | **Boolean**| Filter by agent installed (true) | [optional] 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **createdBy** | **Number**| The User ID for Filtering | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 
 **tags** | **String**| Filter by tags (metadata). This allows filtering by a tag name and value(s)  | [optional] 
 **metadata** | **String**| Alias for tags | [optional] 
 **uuid** | **String**| Filter by UUID | [optional] 
 **externalId** | **String**| Filter by External ID | [optional] 
 **internalId** | **String**| Filter by Internal ID | [optional] 
 **externalUniquelId** | **String**| Filter by External Unique ID | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listServerServicePlans

> InlineResponse200141 listServerServicePlans(zoneId, opts)

Get Available Service Plans for a Host

This endpoint retrieves all the Service Plans available for the specified cloud and host type. It may be used to get the list of available plans when creating a new host or resizing an existing host.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let zoneId = 3; // Number | The Zone ID for Filtering
let opts = {
  'serverTypeId': 5, // Number | The ID of the Host Type
  'siteId': 7 // Number | The Site ID for Filtering
};
apiInstance.listServerServicePlans(zoneId, opts, (error, data, response) => {
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
 **zoneId** | **Number**| The Zone ID for Filtering | 
 **serverTypeId** | **Number**| The ID of the Host Type | [optional] 
 **siteId** | **Number**| The Site ID for Filtering | [optional] 

### Return type

[**InlineResponse200141**](InlineResponse200141.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeHost

> Model200Success removeHost(id, opts)

Delete a Host

Will delete a host asynchronously.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'removeResources': off, // String | Remove Resources
  'removeInstances': on, // String | Remove Instances
  'preserveVolumes': on, // String | Preserve Volumes
  'releaseFloatingIps': off, // String | Release Floating IPs
  'releaseEIPs': off, // String | Alias for releaseFloatingIps
  'force': on // String | Force Delete
};
apiInstance.removeHost(id, opts, (error, data, response) => {
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
 **removeResources** | **String**| Remove Resources | [optional] [default to &#39;on&#39;]
 **removeInstances** | **String**| Remove Instances | [optional] [default to &#39;off&#39;]
 **preserveVolumes** | **String**| Preserve Volumes | [optional] [default to &#39;off&#39;]
 **releaseFloatingIps** | **String**| Release Floating IPs | [optional] [default to &#39;on&#39;]
 **releaseEIPs** | **String**| Alias for releaseFloatingIps | [optional] [default to &#39;on&#39;]
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## restartHost

> Object restartHost(id)

Restart a Host

This will restart a host.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.restartHost(id, (error, data, response) => {
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

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## startHost

> Model200Success startHost(id)

Start a Host

This will start a host.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.startHost(id, (error, data, response) => {
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

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## stopHost

> Model200Success stopHost(id)

Stop a Host

This will stop a host.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.stopHost(id, (error, data, response) => {
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

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateHost

> InlineResponse200137 updateHost(id, opts)

Updating a Host

Updating a Host

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject220': new MorpheusApi.InlineObject220() // InlineObject220 | 
};
apiInstance.updateHost(id, opts, (error, data, response) => {
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
 **inlineObject220** | [**InlineObject220**](InlineObject220.md)|  | [optional] 

### Return type

[**InlineResponse200137**](InlineResponse200137.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateHostAssignTenant

> Object updateHostAssignTenant(id, opts)

Assign To Tenant

This will change the ownership of the host to the specified Tenant account. This is only available to Master Tenant users.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'accountId': 3, // Number | ID of the Tenant
  'inlineObject221': new MorpheusApi.InlineObject221() // InlineObject221 | 
};
apiInstance.updateHostAssignTenant(id, opts, (error, data, response) => {
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
 **accountId** | **Number**| ID of the Tenant | [optional] 
 **inlineObject221** | [**InlineObject221**](InlineObject221.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateHostCloud

> Object updateHostCloud(opts)

Change Server Cloud

This api call is reserved for migrating servers from one cloud to another. This could be due to moving clusters or resource pool scoping of a server without losing the data.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let opts = {
  'inlineObject226': new MorpheusApi.InlineObject226() // InlineObject226 | 
};
apiInstance.updateHostCloud(opts, (error, data, response) => {
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
 **inlineObject226** | [**InlineObject226**](InlineObject226.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateHostExecuteWorkflow

> Model200Success updateHostExecuteWorkflow(id, opts)

Run Workflow on a Host

This will run a provisioning workflow on a host.  For operational workflows, see Execute a Workflow. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'workflowId': 3, // Number | ID of the workflow to execute
  'workflowName': myworkflow, // String | Name of the workflow to execute
  'inlineObject225': new MorpheusApi.InlineObject225() // InlineObject225 | 
};
apiInstance.updateHostExecuteWorkflow(id, opts, (error, data, response) => {
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
 **workflowId** | **Number**| ID of the workflow to execute | [optional] 
 **workflowName** | **String**| Name of the workflow to execute | [optional] 
 **inlineObject225** | [**InlineObject225**](InlineObject225.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateHostInstallAgent

> Object updateHostInstallAgent(id, opts)

Install Agent

This will make the host a managed server, and install the agent.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject222': new MorpheusApi.InlineObject222() // InlineObject222 | 
};
apiInstance.updateHostInstallAgent(id, opts, (error, data, response) => {
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
 **inlineObject222** | [**InlineObject222**](InlineObject222.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateHostManaged

> Object updateHostManaged(id, opts)

Convert To Managed

This will make the host a managed server, and install the agent.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject223': new MorpheusApi.InlineObject223() // InlineObject223 | 
};
apiInstance.updateHostManaged(id, opts, (error, data, response) => {
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
 **inlineObject223** | [**InlineObject223**](InlineObject223.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateHostResize

> Object updateHostResize(id, opts)

Resize a Host

Will resize a host asynchronously. This endpoint also allows for NIC reconfiguration by passing a new array of &#x60;networkInterfaces&#x60;.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject224': new MorpheusApi.InlineObject224() // InlineObject224 | 
};
apiInstance.updateHostResize(id, opts, (error, data, response) => {
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
 **inlineObject224** | [**InlineObject224**](InlineObject224.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateHostUpgradeAgent

> Model200Success updateHostUpgradeAgent(id)

Upgrade Agent

This will upgrade the version of the agent installed on the host.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.updateHostUpgradeAgent(id, (error, data, response) => {
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

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateServerNetworkInterface

> Object updateServerNetworkInterface(id, networkInterfaceId, opts)

Updating a label for a Server&#39;s Network

Updating a Server&#39;s Network&#39;s Label

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let networkInterfaceId = 7; // Number | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
let opts = {
  'networkInterfaceUpdate': {"name":"newLabelName"} // NetworkInterfaceUpdate | 
};
apiInstance.updateServerNetworkInterface(id, networkInterfaceId, opts, (error, data, response) => {
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
 **networkInterfaceId** | **Number**| NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced | 
 **networkInterfaceUpdate** | [**NetworkInterfaceUpdate**](NetworkInterfaceUpdate.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateWikiServer

> Object updateWikiServer(id, opts)

Update a Server Wiki Page

Updates a server Wiki page. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HostsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject271': new MorpheusApi.InlineObject271() // InlineObject271 | 
};
apiInstance.updateWikiServer(id, opts, (error, data, response) => {
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
 **inlineObject271** | [**InlineObject271**](InlineObject271.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

