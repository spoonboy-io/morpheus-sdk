# openapi_client.InvoicesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_invoice_line_items**](InvoicesApi.md#get_invoice_line_items) | **GET** /api/invoice-line-items/{id} | Get a Specific Invoice Line Item
[**get_invoices**](InvoicesApi.md#get_invoices) | **GET** /api/invoices/{id} | Get a Specific Invoice
[**list_invoice_line_items**](InvoicesApi.md#list_invoice_line_items) | **GET** /api/invoice-line-items | List All Invoice Line Items
[**list_invoices**](InvoicesApi.md#list_invoices) | **GET** /api/invoices | List All Invoices
[**update_invoices**](InvoicesApi.md#update_invoices) | **PUT** /api/invoices/{id} | Update Invoice Tags


# **get_invoice_line_items**
> GetInvoiceLineItems200Response get_invoice_line_items(id)

Get a Specific Invoice Line Item

Get details about a specific invoice line item.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import invoices_api
from openapi_client.model.get_invoice_line_items200_response import GetInvoiceLineItems200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = invoices_api.InvoicesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Invoice Line Item
        api_response = api_instance.get_invoice_line_items(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InvoicesApi->get_invoice_line_items: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetInvoiceLineItems200Response**](GetInvoiceLineItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice Line Item Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_invoices**
> GetInvoices200Response get_invoices(id)

Get a Specific Invoice

Get details about a specific invoice.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import invoices_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_invoices200_response import GetInvoices200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = invoices_api.InvoicesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Invoice
        api_response = api_instance.get_invoices(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InvoicesApi->get_invoices: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetInvoices200Response**](GetInvoices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_invoice_line_items**
> ListInvoiceLineItems200Response list_invoice_line_items()

List All Invoice Line Items

This endpoint retrieves a list of all invoice line items for the specified parameters.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import invoices_api
from openapi_client.model.list_invoice_line_items200_response import ListInvoiceLineItems200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = invoices_api.InvoicesApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    period = "201901" # str | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  (optional)
    ref_type = "ComputeSite" # str | If specified will return an exact match on refType.  (optional)
    ref_id = 3 # int | If specified will return an exact match on refId (optional)
    zone_id = 3 # int | The Zone ID for Filtering (optional)
    site_id = 7 # int | The Site ID for Filtering (optional)
    instance_id = 94 # int | The Instance ID for Filtering (optional)
    container_id = 135 # int | The Container ID for Filtering (optional)
    server_id = 97 # int | The Server ID for Filtering (optional)
    user_id = 6 # int | Filter by User ID (optional)
    project_id = 1 # int | The Project ID for Filtering (optional)
    active = False # bool | True or False flag for Active (optional)
    account_id = 3 # int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)
    include_totals = False # bool | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called `invoiceTotals`.  (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Invoice Line Items
        api_response = api_instance.list_invoice_line_items(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, start_date=start_date, end_date=end_date, period=period, ref_type=ref_type, ref_id=ref_id, zone_id=zone_id, site_id=site_id, instance_id=instance_id, container_id=container_id, server_id=server_id, user_id=user_id, project_id=project_id, active=active, account_id=account_id, include_totals=include_totals)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InvoicesApi->list_invoice_line_items: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **period** | **str**| Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  | [optional]
 **ref_type** | **str**| If specified will return an exact match on refType.  | [optional]
 **ref_id** | **int**| If specified will return an exact match on refId | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **site_id** | **int**| The Site ID for Filtering | [optional]
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **container_id** | **int**| The Container ID for Filtering | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **user_id** | **int**| Filter by User ID | [optional]
 **project_id** | **int**| The Project ID for Filtering | [optional]
 **active** | **bool**| True or False flag for Active | [optional]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **include_totals** | **bool**| Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called &#x60;invoiceTotals&#x60;.  | [optional] if omitted the server will use the default value of False

### Return type

[**ListInvoiceLineItems200Response**](ListInvoiceLineItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Invoice Line Items |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_invoices**
> ListInvoices200Response list_invoices()

List All Invoices

This endpoint retrieves a list of invoices for the specified parameters.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import invoices_api
from openapi_client.model.list_invoices200_response import ListInvoices200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = invoices_api.InvoicesApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "refName" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "refName"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    start_date = "2019-01-01" # str | Filter by startDate greater than or equal to a specified date (optional)
    end_date = "2020-01-01" # str | Filter by endDate less than or equal to a specified date (optional)
    period = "201901" # str | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  (optional)
    ref_type = "ComputeSite" # str | If specified will return an exact match on refType.  (optional)
    ref_id = 3 # int | If specified will return an exact match on refId (optional)
    ref_status = "provisioned" # str | If specified, will filter on the associated StorageVolume status. This is only applicable whn `refType=StorageVolume`.  (optional)
    zone_id = 3 # int | The Zone ID for Filtering (optional)
    site_id = 7 # int | The Site ID for Filtering (optional)
    instance_id = 94 # int | The Instance ID for Filtering (optional)
    container_id = 135 # int | The Container ID for Filtering (optional)
    server_id = 97 # int | The Server ID for Filtering (optional)
    user_id = 6 # int | Filter by User ID (optional)
    project_id = 1 # int | The Project ID for Filtering (optional)
    active = False # bool | True or False flag for Active (optional)
    account_id = 3 # int | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. (optional)
    include_line_items = False # bool | Pass true to include the list of `lineItems` for each invoice. Only `lineItemCount` is returned by default.  (optional) if omitted the server will use the default value of False
    include_totals = False # bool | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called `invoiceTotals`.  (optional) if omitted the server will use the default value of False
    tags = "tags.env=qa&tags.env=test" # str | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List All Invoices
        api_response = api_instance.list_invoices(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, start_date=start_date, end_date=end_date, period=period, ref_type=ref_type, ref_id=ref_id, ref_status=ref_status, zone_id=zone_id, site_id=site_id, instance_id=instance_id, container_id=container_id, server_id=server_id, user_id=user_id, project_id=project_id, active=active, account_id=account_id, include_line_items=include_line_items, include_totals=include_totals, tags=tags)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InvoicesApi->list_invoices: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "refName"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **start_date** | **str**| Filter by startDate greater than or equal to a specified date | [optional]
 **end_date** | **str**| Filter by endDate less than or equal to a specified date | [optional]
 **period** | **str**| Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  | [optional]
 **ref_type** | **str**| If specified will return an exact match on refType.  | [optional]
 **ref_id** | **int**| If specified will return an exact match on refId | [optional]
 **ref_status** | **str**| If specified, will filter on the associated StorageVolume status. This is only applicable whn &#x60;refType&#x3D;StorageVolume&#x60;.  | [optional]
 **zone_id** | **int**| The Zone ID for Filtering | [optional]
 **site_id** | **int**| The Site ID for Filtering | [optional]
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **container_id** | **int**| The Container ID for Filtering | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **user_id** | **int**| Filter by User ID | [optional]
 **project_id** | **int**| The Project ID for Filtering | [optional]
 **active** | **bool**| True or False flag for Active | [optional]
 **account_id** | **int**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **include_line_items** | **bool**| Pass true to include the list of &#x60;lineItems&#x60; for each invoice. Only &#x60;lineItemCount&#x60; is returned by default.  | [optional] if omitted the server will use the default value of False
 **include_totals** | **bool**| Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called &#x60;invoiceTotals&#x60;.  | [optional] if omitted the server will use the default value of False
 **tags** | **str**| Filter by tags (metadata). This allows filtering by a tag name and value(s)  | [optional]

### Return type

[**ListInvoices200Response**](ListInvoices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of Invoices |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_invoices**
> UpdateInvoices200Response update_invoices(id)

Update Invoice Tags

Update, Add, or Remove invoice tag(s).

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import invoices_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_invoices200_response import UpdateInvoices200Response
from openapi_client.model.update_invoices_request import UpdateInvoicesRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = invoices_api.InvoicesApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_invoices_request = UpdateInvoicesRequest(
        invoice=UpdateInvoicesRequestInvoice(
            tags=[{"name":"hello","value":"world"},{"name":"foo","value":"bar"}],
            add_tags=[{"name":"hello","value":"world"},{"name":"foo","value":"bar"}],
            remove_tags=[{"name":"hello","value":"world"},{"name":"foo","value":"bar"}],
        ),
    ) # UpdateInvoicesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Invoice Tags
        api_response = api_instance.update_invoices(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InvoicesApi->update_invoices: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Invoice Tags
        api_response = api_instance.update_invoices(id, update_invoices_request=update_invoices_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling InvoicesApi->update_invoices: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_invoices_request** | [**UpdateInvoicesRequest**](UpdateInvoicesRequest.md)|  | [optional]

### Return type

[**UpdateInvoices200Response**](UpdateInvoices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Invoice Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

