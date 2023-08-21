# MorpheusApi.PricesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addPrices**](PricesApi.md#addPrices) | **POST** /api/prices | Creates a Price
[**deactivatePrices**](PricesApi.md#deactivatePrices) | **PUT** /api/prices/{id}/deactivate | Deactivates a Price
[**getPrices**](PricesApi.md#getPrices) | **GET** /api/prices/{id} | Retrieves a Specific Price
[**listPrices**](PricesApi.md#listPrices) | **GET** /api/prices | Retrieves all Prices
[**updatePrices**](PricesApi.md#updatePrices) | **PUT** /api/prices/{id} | Updates a Price



## addPrices

> Object addPrices(opts)

Creates a Price

Creates a price. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PricesApi();
let opts = {
  'inlineObject200': new MorpheusApi.InlineObject200() // InlineObject200 | 
};
apiInstance.addPrices(opts, (error, data, response) => {
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
 **inlineObject200** | [**InlineObject200**](InlineObject200.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deactivatePrices

> Model200Success deactivatePrices(id)

Deactivates a Price

Deactivates a price. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PricesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deactivatePrices(id, (error, data, response) => {
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


## getPrices

> InlineResponse200124 getPrices(id)

Retrieves a Specific Price

Retrieves a specific price. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PricesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getPrices(id, (error, data, response) => {
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

[**InlineResponse200124**](InlineResponse200124.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listPrices

> Object listPrices(opts)

Retrieves all Prices

Retrieves all prices. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PricesApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'includeInactive': true, // Boolean | If true, include inactive prices in the results
  'priceType': "priceType_example", // String | Restricts query to only load only prices with specified priceType. * `fixed` - Everything * `compute` - Memory + CPU * `memory` - Memory * `cores` - Cores * `storage` - Storage * `datastore` - Datastore * `platform` - Platform * `software` - Software 
  'platform': linux, // String | Restricts query to only load only prices with specified platform
  'currency': USD, // String | Restricts query to only load only prices with specified currency
  'type': "type_example" // String | Filter by type code
};
apiInstance.listPrices(opts, (error, data, response) => {
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
 **includeInactive** | **Boolean**| If true, include inactive prices in the results | [optional] 
 **priceType** | **String**| Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software  | [optional] 
 **platform** | **String**| Restricts query to only load only prices with specified platform | [optional] 
 **currency** | **String**| Restricts query to only load only prices with specified currency | [optional] 
 **type** | **String**| Filter by type code | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updatePrices

> Object updatePrices(id, opts)

Updates a Price

Updates a price. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PricesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject201': new MorpheusApi.InlineObject201() // InlineObject201 | 
};
apiInstance.updatePrices(id, opts, (error, data, response) => {
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
 **inlineObject201** | [**InlineObject201**](InlineObject201.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

