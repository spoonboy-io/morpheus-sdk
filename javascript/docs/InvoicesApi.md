# MorpheusApi.InvoicesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getInvoiceLineItems**](InvoicesApi.md#getInvoiceLineItems) | **GET** /api/invoice-line-items/{id} | Get a Specific Invoice Line Item
[**getInvoices**](InvoicesApi.md#getInvoices) | **GET** /api/invoices/{id} | Get a Specific Invoice
[**listInvoiceLineItems**](InvoicesApi.md#listInvoiceLineItems) | **GET** /api/invoice-line-items | List All Invoice Line Items
[**listInvoices**](InvoicesApi.md#listInvoices) | **GET** /api/invoices | List All Invoices
[**updateInvoices**](InvoicesApi.md#updateInvoices) | **PUT** /api/invoices/{id} | Update Invoice Tags



## getInvoiceLineItems

> Object getInvoiceLineItems(id)

Get a Specific Invoice Line Item

Get details about a specific invoice line item.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InvoicesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getInvoiceLineItems(id, (error, data, response) => {
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


## getInvoices

> Object getInvoices(id)

Get a Specific Invoice

Get details about a specific invoice.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InvoicesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getInvoices(id, (error, data, response) => {
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


## listInvoiceLineItems

> Object listInvoiceLineItems(opts)

List All Invoice Line Items

This endpoint retrieves a list of all invoice line items for the specified parameters.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InvoicesApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'period': 201901, // String | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM. 
  'refType': ComputeSite, // String | If specified will return an exact match on refType. 
  'refId': 3, // Number | If specified will return an exact match on refId
  'zoneId': 3, // Number | The Zone ID for Filtering
  'siteId': 7, // Number | The Site ID for Filtering
  'instanceId': 94, // Number | The Instance ID for Filtering
  'containerId': 135, // Number | The Container ID for Filtering
  'serverId': 97, // Number | The Server ID for Filtering
  'userId': 6, // Number | Filter by User ID
  'projectId': 1, // Number | The Project ID for Filtering
  'active': false, // Boolean | True or False flag for Active
  'accountId': 3, // Number | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
  'includeTotals': false // Boolean | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called `invoiceTotals`. 
};
apiInstance.listInvoiceLineItems(opts, (error, data, response) => {
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
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **period** | **String**| Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  | [optional] 
 **refType** | **String**| If specified will return an exact match on refType.  | [optional] 
 **refId** | **Number**| If specified will return an exact match on refId | [optional] 
 **zoneId** | **Number**| The Zone ID for Filtering | [optional] 
 **siteId** | **Number**| The Site ID for Filtering | [optional] 
 **instanceId** | **Number**| The Instance ID for Filtering | [optional] 
 **containerId** | **Number**| The Container ID for Filtering | [optional] 
 **serverId** | **Number**| The Server ID for Filtering | [optional] 
 **userId** | **Number**| Filter by User ID | [optional] 
 **projectId** | **Number**| The Project ID for Filtering | [optional] 
 **active** | **Boolean**| True or False flag for Active | [optional] 
 **accountId** | **Number**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 
 **includeTotals** | **Boolean**| Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called &#x60;invoiceTotals&#x60;.  | [optional] [default to false]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listInvoices

> Object listInvoices(opts)

List All Invoices

This endpoint retrieves a list of invoices for the specified parameters.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InvoicesApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'refName'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'period': 201901, // String | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM. 
  'refType': ComputeSite, // String | If specified will return an exact match on refType. 
  'refId': 3, // Number | If specified will return an exact match on refId
  'refStatus': provisioned, // String | If specified, will filter on the associated StorageVolume status. This is only applicable whn `refType=StorageVolume`. 
  'zoneId': 3, // Number | The Zone ID for Filtering
  'siteId': 7, // Number | The Site ID for Filtering
  'instanceId': 94, // Number | The Instance ID for Filtering
  'containerId': 135, // Number | The Container ID for Filtering
  'serverId': 97, // Number | The Server ID for Filtering
  'userId': 6, // Number | Filter by User ID
  'projectId': 1, // Number | The Project ID for Filtering
  'active': false, // Boolean | True or False flag for Active
  'accountId': 3, // Number | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
  'includeLineItems': false, // Boolean | Pass true to include the list of `lineItems` for each invoice. Only `lineItemCount` is returned by default. 
  'includeTotals': false, // Boolean | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called `invoiceTotals`. 
  'tags': tags.env=qa&tags.env=test // String | Filter by tags (metadata). This allows filtering by a tag name and value(s) 
};
apiInstance.listInvoices(opts, (error, data, response) => {
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
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;refName&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **period** | **String**| Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  | [optional] 
 **refType** | **String**| If specified will return an exact match on refType.  | [optional] 
 **refId** | **Number**| If specified will return an exact match on refId | [optional] 
 **refStatus** | **String**| If specified, will filter on the associated StorageVolume status. This is only applicable whn &#x60;refType&#x3D;StorageVolume&#x60;.  | [optional] 
 **zoneId** | **Number**| The Zone ID for Filtering | [optional] 
 **siteId** | **Number**| The Site ID for Filtering | [optional] 
 **instanceId** | **Number**| The Instance ID for Filtering | [optional] 
 **containerId** | **Number**| The Container ID for Filtering | [optional] 
 **serverId** | **Number**| The Server ID for Filtering | [optional] 
 **userId** | **Number**| Filter by User ID | [optional] 
 **projectId** | **Number**| The Project ID for Filtering | [optional] 
 **active** | **Boolean**| True or False flag for Active | [optional] 
 **accountId** | **Number**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 
 **includeLineItems** | **Boolean**| Pass true to include the list of &#x60;lineItems&#x60; for each invoice. Only &#x60;lineItemCount&#x60; is returned by default.  | [optional] [default to false]
 **includeTotals** | **Boolean**| Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called &#x60;invoiceTotals&#x60;.  | [optional] [default to false]
 **tags** | **String**| Filter by tags (metadata). This allows filtering by a tag name and value(s)  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateInvoices

> Object updateInvoices(id, opts)

Update Invoice Tags

Update, Add, or Remove invoice tag(s).

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.InvoicesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject102': new MorpheusApi.InlineObject102() // InlineObject102 | 
};
apiInstance.updateInvoices(id, opts, (error, data, response) => {
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
 **inlineObject102** | [**InlineObject102**](InlineObject102.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

