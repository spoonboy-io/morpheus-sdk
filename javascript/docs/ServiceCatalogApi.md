# MorpheusApi.ServiceCatalogApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCatalogCart**](ServiceCatalogApi.md#addCatalogCart) | **POST** /api/catalog/checkout | Checkout Catalog Cart
[**addCatalogCartItem**](ServiceCatalogApi.md#addCatalogCartItem) | **PUT** /api/catalog/cart/items | Add Catalog Item to Cart
[**addCatalogOrder**](ServiceCatalogApi.md#addCatalogOrder) | **POST** /api/catalog/orders | Place Catalog Order
[**deleteCatalogCart**](ServiceCatalogApi.md#deleteCatalogCart) | **DELETE** /api/catalog/cart | Clear Catalog Cart
[**deleteCatalogCartItem**](ServiceCatalogApi.md#deleteCatalogCartItem) | **DELETE** /api/catalog/cart/items/{id} | Remove a Catalog Item From Cart
[**deleteCatalogItem**](ServiceCatalogApi.md#deleteCatalogItem) | **DELETE** /api/catalog/items/{id} | Delete a Catalog Inventory Item
[**getCatalogItem**](ServiceCatalogApi.md#getCatalogItem) | **GET** /api/catalog/items/{id} | Get a Specific Catalog Inventory Item
[**getCatalogType**](ServiceCatalogApi.md#getCatalogType) | **GET** /api/catalog/types/{id} | Get a Specific Catalog Type
[**listCatalogCart**](ServiceCatalogApi.md#listCatalogCart) | **GET** /api/catalog/cart | Get Catalog Cart
[**listCatalogItems**](ServiceCatalogApi.md#listCatalogItems) | **GET** /api/catalog/items | List Catalog Inventory Items
[**listCatalogTypes**](ServiceCatalogApi.md#listCatalogTypes) | **GET** /api/catalog/types | List Catalog Types



## addCatalogCart

> Object addCatalogCart()

Checkout Catalog Cart

Use this command to checkout, finalizing your cart and placing an order. This converts each item in the cart to an inventory item, changing the status from IN_CART to ORDERED and potentially starts the provisioning process for each item.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
apiInstance.addCatalogCart((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## addCatalogCartItem

> Object addCatalogCartItem(opts)

Add Catalog Item to Cart

Use this command to add an item to your service catalog cart.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
let opts = {
  'validate': false, // Boolean | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory
  'catalogCartItemCreate': {"item":{"type":{"name":"example"},"config":{"appName":"My App"}}} // CatalogCartItemCreate | 
};
apiInstance.addCatalogCartItem(opts, (error, data, response) => {
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
 **validate** | **Boolean**| Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory | [optional] [default to false]
 **catalogCartItemCreate** | [**CatalogCartItemCreate**](CatalogCartItemCreate.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addCatalogOrder

> Object addCatalogOrder(opts)

Place Catalog Order

This will place an order for the specified items, adding items to the inventory right away, without using the cart.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
let opts = {
  'validate': false, // Boolean | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory
  'inlineObject227': new MorpheusApi.InlineObject227() // InlineObject227 | 
};
apiInstance.addCatalogOrder(opts, (error, data, response) => {
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
 **validate** | **Boolean**| Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory | [optional] [default to false]
 **inlineObject227** | [**InlineObject227**](InlineObject227.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteCatalogCart

> Model200Success deleteCatalogCart()

Clear Catalog Cart

Use this command to empty your cart, deleting all the items in it.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
apiInstance.deleteCatalogCart((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteCatalogCartItem

> Model200Success deleteCatalogCartItem(id)

Remove a Catalog Item From Cart

Will remove a catalog item that is currently in the cart.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteCatalogCartItem(id, (error, data, response) => {
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


## deleteCatalogItem

> Model200Success deleteCatalogItem(id, opts)

Delete a Catalog Inventory Item

Will delete a catalog inventory item, which by default will deprovision any associated any instances and servers.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'preserveVolumes': on, // String | Preserve Volumes
  'keepBackups': on, // String | Preserve copy of backups
  'releaseFloatingIps': off, // String | Release Floating IPs
  'releaseEIPs': off, // String | Alias for releaseFloatingIps
  'removeInstances': off, // String | Remove Instances. Only applies to type `blueprint` (Apps)
  'force': on // String | Force Delete
};
apiInstance.deleteCatalogItem(id, opts, (error, data, response) => {
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
 **preserveVolumes** | **String**| Preserve Volumes | [optional] [default to &#39;off&#39;]
 **keepBackups** | **String**| Preserve copy of backups | [optional] [default to &#39;off&#39;]
 **releaseFloatingIps** | **String**| Release Floating IPs | [optional] [default to &#39;on&#39;]
 **releaseEIPs** | **String**| Alias for releaseFloatingIps | [optional] [default to &#39;on&#39;]
 **removeInstances** | **String**| Remove Instances. Only applies to type &#x60;blueprint&#x60; (Apps) | [optional] [default to &#39;on&#39;]
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getCatalogItem

> InlineResponse200140 getCatalogItem(id)

Get a Specific Catalog Inventory Item

This endpoint retrieves a specific catalog inventory item.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCatalogItem(id, (error, data, response) => {
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

[**InlineResponse200140**](InlineResponse200140.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getCatalogType

> Object getCatalogType(id)

Get a Specific Catalog Type

This endpoint retrieves a specific catalog item type. This also returns an array of associated optionTypes that are used to configure the catalog item.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCatalogType(id, (error, data, response) => {
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


## listCatalogCart

> InlineResponse200139 listCatalogCart()

Get Catalog Cart

This endpoint retrieves the current catalog cart and all the items in it.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
apiInstance.listCatalogCart((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse200139**](InlineResponse200139.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCatalogItems

> Object listCatalogItems(opts)

List Catalog Inventory Items

This endpoint retrieves a list of the catalog inventory items.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listCatalogItems(opts, (error, data, response) => {
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

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCatalogTypes

> Object listCatalogTypes(opts)

List Catalog Types

This endpoint retrieves the types available for ordering.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServiceCatalogApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'featured': false // Boolean | Filter by featured
};
apiInstance.listCatalogTypes(opts, (error, data, response) => {
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
 **featured** | **Boolean**| Filter by featured | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

