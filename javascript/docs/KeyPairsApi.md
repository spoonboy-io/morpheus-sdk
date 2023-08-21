# MorpheusApi.KeyPairsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addKeyPairs**](KeyPairsApi.md#addKeyPairs) | **POST** /api/key-pairs | Creates a Key Pair
[**generateKeyPairs**](KeyPairsApi.md#generateKeyPairs) | **POST** /api/key-pairs/generate | Generates a Key Pair
[**getKeyPairs**](KeyPairsApi.md#getKeyPairs) | **GET** /api/key-pairs/{id} | Retrieves a Specific Key Pair
[**removeKeyPairs**](KeyPairsApi.md#removeKeyPairs) | **DELETE** /api/key-pairs/{id} | Deletes a Key Pair



## addKeyPairs

> Object addKeyPairs(opts)

Creates a Key Pair

Creates a Key Pair.  Private keys are typically optional. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.KeyPairsApi();
let opts = {
  'inlineObject105': new MorpheusApi.InlineObject105() // InlineObject105 | 
};
apiInstance.addKeyPairs(opts, (error, data, response) => {
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
 **inlineObject105** | [**InlineObject105**](InlineObject105.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## generateKeyPairs

> InlineResponse20066 generateKeyPairs(opts)

Generates a Key Pair

Generates a Key Pair. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.KeyPairsApi();
let opts = {
  'inlineObject106': new MorpheusApi.InlineObject106() // InlineObject106 | 
};
apiInstance.generateKeyPairs(opts, (error, data, response) => {
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
 **inlineObject106** | [**InlineObject106**](InlineObject106.md)|  | [optional] 

### Return type

[**InlineResponse20066**](InlineResponse20066.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getKeyPairs

> InlineResponse20067 getKeyPairs(id)

Retrieves a Specific Key Pair

Retrieves a specific key pair. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.KeyPairsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getKeyPairs(id, (error, data, response) => {
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

[**InlineResponse20067**](InlineResponse20067.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeKeyPairs

> Model200Success removeKeyPairs(id)

Deletes a Key Pair

Deletes a specified key pair. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.KeyPairsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeKeyPairs(id, (error, data, response) => {
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

