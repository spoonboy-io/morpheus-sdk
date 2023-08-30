# openapi_client.LibraryApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_file_template**](LibraryApi.md#add_file_template) | **POST** /api/library/container-templates | Create a File Template
[**add_instance_type**](LibraryApi.md#add_instance_type) | **POST** /api/library/instance-types | Create an Instance Type
[**add_layout**](LibraryApi.md#add_layout) | **POST** /api/library/instance-types/{instanceTypeId}/layouts | Create a Layout
[**add_node_type**](LibraryApi.md#add_node_type) | **POST** /api/library/container-types | Create a Node Type
[**add_option_list**](LibraryApi.md#add_option_list) | **POST** /api/library/option-type-lists | Create an Option List
[**add_option_type**](LibraryApi.md#add_option_type) | **POST** /api/library/option-types | Create an Input
[**add_script**](LibraryApi.md#add_script) | **POST** /api/library/container-scripts | Create a Script
[**add_spec_template**](LibraryApi.md#add_spec_template) | **POST** /api/library/spec-templates | Create a Spec Template
[**add_virtual_image**](LibraryApi.md#add_virtual_image) | **POST** /api/virtual-images | Create a Virtual Image
[**add_virtual_image_file**](LibraryApi.md#add_virtual_image_file) | **POST** /api/virtual-images/{virtualImageId}/upload | Upload Virtual Image File
[**delete_file_template**](LibraryApi.md#delete_file_template) | **DELETE** /api/library/container-templates/{id} | Delete a File Template
[**delete_instance_type**](LibraryApi.md#delete_instance_type) | **DELETE** /api/library/instance-types/{instanceTypeId} | Delete an Instance Type
[**delete_layout**](LibraryApi.md#delete_layout) | **DELETE** /api/library/layouts/{id} | Delete a Layout
[**delete_node_type**](LibraryApi.md#delete_node_type) | **DELETE** /api/library/container-types/{id} | Delete a Node Type
[**delete_option_list**](LibraryApi.md#delete_option_list) | **DELETE** /api/library/option-type-lists/{id} | Delete an Option List
[**delete_option_type**](LibraryApi.md#delete_option_type) | **DELETE** /api/library/option-types/{id} | Delete an Input
[**delete_script**](LibraryApi.md#delete_script) | **DELETE** /api/library/container-scripts/{id} | Delete a Script
[**delete_spec_template**](LibraryApi.md#delete_spec_template) | **DELETE** /api/library/spec-templates/{id} | Delete a Spec Template
[**get_file_template**](LibraryApi.md#get_file_template) | **GET** /api/library/container-templates/{id} | Get a Specific File Template
[**get_input**](LibraryApi.md#get_input) | **GET** /api/library/option-types/{id} | Get A Specific Input
[**get_instance_type**](LibraryApi.md#get_instance_type) | **GET** /api/library/instance-types/{instanceTypeId} | Get a Specific Instance Type
[**get_layout**](LibraryApi.md#get_layout) | **GET** /api/library/layouts/{id} | Get a Specific Layout
[**get_node_type**](LibraryApi.md#get_node_type) | **GET** /api/library/container-types/{id} | Get a Specific Node Type
[**get_option_list**](LibraryApi.md#get_option_list) | **GET** /api/library/option-type-lists/{id} | Get a Specific Option List
[**get_option_list_items**](LibraryApi.md#get_option_list_items) | **GET** /api/library/option-type-lists/{id}/items | List Items for a Specific Option List
[**get_script**](LibraryApi.md#get_script) | **GET** /api/library/container-scripts/{id} | Get a Specific Script
[**get_security_package_type**](LibraryApi.md#get_security_package_type) | **GET** /api/security-package-types/{id} | Retrieves a Specific Security Package Type
[**get_spec_template**](LibraryApi.md#get_spec_template) | **GET** /api/library/spec-templates/{id} | Get a Specific Spec Template
[**get_virtual_image**](LibraryApi.md#get_virtual_image) | **GET** /api/virtual-images/{virtualImageId} | Get a Specific Virtual Image
[**list_file_templates**](LibraryApi.md#list_file_templates) | **GET** /api/library/container-templates | Get All File Templates
[**list_inputs**](LibraryApi.md#list_inputs) | **GET** /api/library/option-types | Get All Inputs
[**list_instance_types**](LibraryApi.md#list_instance_types) | **GET** /api/library/instance-types | Get All Instance Types
[**list_layouts**](LibraryApi.md#list_layouts) | **GET** /api/library/layouts | Get All Layouts
[**list_layouts_for_instance_type**](LibraryApi.md#list_layouts_for_instance_type) | **GET** /api/library/instance-types/{instanceTypeId}/layouts | Get All Layouts For an Instance Type
[**list_node_types**](LibraryApi.md#list_node_types) | **GET** /api/library/container-types | Get All Node Types
[**list_option_lists**](LibraryApi.md#list_option_lists) | **GET** /api/library/option-type-lists | Get All Option Lists
[**list_scripts**](LibraryApi.md#list_scripts) | **GET** /api/library/container-scripts | Get All Scripts
[**list_security_package_types**](LibraryApi.md#list_security_package_types) | **GET** /api/security-package-types | Retrieves all Security Package Types
[**list_spec_templates**](LibraryApi.md#list_spec_templates) | **GET** /api/library/spec-templates | Get All Spec Templates
[**list_virtual_image_locations**](LibraryApi.md#list_virtual_image_locations) | **GET** /api/virtual-images/{virtualImageId}/locations | Get a List of Virtual Image Locations
[**list_virtual_images**](LibraryApi.md#list_virtual_images) | **GET** /api/virtual-images | Get List of Virtual Images
[**remove_security_scans**](LibraryApi.md#remove_security_scans) | **DELETE** /api/security-scans/{id} | Deletes a Security Scan
[**remove_virtual_image**](LibraryApi.md#remove_virtual_image) | **DELETE** /api/virtual-images/{virtualImageId} | Delete a Virtual Image
[**remove_virtual_image_file**](LibraryApi.md#remove_virtual_image_file) | **DELETE** /api/virtual-images/{virtualImageId}/files | Remove Virtual Image File
[**remove_virtual_image_location**](LibraryApi.md#remove_virtual_image_location) | **DELETE** /api/virtual-images/{virtualImageId}/locations/{id} | Delete a Virtual Image Location
[**set_instance_type_featured**](LibraryApi.md#set_instance_type_featured) | **PUT** /api/library/instance-types/{instanceTypeId}/toggle-featured | Toggle Featured For Instance Type
[**update_file_template**](LibraryApi.md#update_file_template) | **PUT** /api/library/container-templates/{id} | Update a File Template
[**update_instance_type**](LibraryApi.md#update_instance_type) | **PUT** /api/library/instance-types/{instanceTypeId} | Update an Instance Type
[**update_instance_type_logo**](LibraryApi.md#update_instance_type_logo) | **POST** /api/library/instance-types/{instanceTypeId}/update-logo | Update Logo For Instance Type
[**update_layout**](LibraryApi.md#update_layout) | **PUT** /api/library/layouts/{id} | Update a Layout
[**update_layout_permissions**](LibraryApi.md#update_layout_permissions) | **PUT** /api/library/layouts/{id}/permissions | Update Layout Permissions
[**update_node_type**](LibraryApi.md#update_node_type) | **PUT** /api/library/container-types/{id} | Update a Node Type
[**update_option_list**](LibraryApi.md#update_option_list) | **PUT** /api/library/option-type-lists/{id} | Update an Option List
[**update_option_type**](LibraryApi.md#update_option_type) | **PUT** /api/library/option-types/{id} | Update an Input
[**update_script**](LibraryApi.md#update_script) | **PUT** /api/library/container-scripts/{id} | Update a Script
[**update_spec_template**](LibraryApi.md#update_spec_template) | **PUT** /api/library/spec-templates/{id} | Update a Spec Template
[**update_virtual_image**](LibraryApi.md#update_virtual_image) | **PUT** /api/virtual-images/{virtualImageId} | Update a Virtual Image


# **add_file_template**
> SuccessId add_file_template()

Create a File Template

Use this command to create a file template.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.add_file_template_request import AddFileTemplateRequest
from openapi_client.model.success_id import SuccessId
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
    api_instance = library_api.LibraryApi(api_client)
    add_file_template_request = AddFileTemplateRequest(
        container_template=FileTemplateCreate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            file_name="file_name_example",
            file_path="file_path_example",
            category="category_example",
            template_phase="template_phase_example",
            template="template_example",
            file_owner=1,
            setting_name="setting_name_example",
            setting_category="setting_category_example",
        ),
    ) # AddFileTemplateRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a File Template
        api_response = api_instance.add_file_template(add_file_template_request=add_file_template_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_file_template: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_file_template_request** | [**AddFileTemplateRequest**](AddFileTemplateRequest.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_instance_type**
> Model200Success add_instance_type()

Create an Instance Type

Use this command to create an instance type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.add_instance_type_request import AddInstanceTypeRequest
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    add_instance_type_request = AddInstanceTypeRequest(
        instance_type=InstanceTypeCreate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            description="description_example",
            code="code_example",
            category="category_example",
            visibility="private",
            featured=True,
            has_settings=True,
            has_auto_scale=True,
            has_deployment=True,
            environment_prefix="environment_prefix_example",
            environment_variables=[
                ClusterLayoutCreateEnvironmentVariablesInner(
                    name="name_example",
                    value="value_example",
                    masked=False,
                    export=False,
                ),
            ],
            price_sets=[
                InstanceTypeCreatePriceSetsInner(
                    id=1,
                ),
            ],
            option_types=[
                1,
            ],
        ),
    ) # AddInstanceTypeRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an Instance Type
        api_response = api_instance.add_instance_type(add_instance_type_request=add_instance_type_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_instance_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_instance_type_request** | [**AddInstanceTypeRequest**](AddInstanceTypeRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_layout**
> SuccessId add_layout(instance_type_id)

Create a Layout

Use this command to create a layout.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_layout_request import AddLayoutRequest
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
    api_instance = library_api.LibraryApi(api_client)
    instance_type_id = 2 # float | The ID of the instance type
    add_layout_request = AddLayoutRequest(
        instance_type_layout=InstanceTypeLayoutCreate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            instance_version="instance_version_example",
            description="description_example",
            creatable=True,
            provision_type_code="provision_type_code_example",
            memory_requirement="memory_requirement_example",
            has_auto_scale=False,
            supports_convert_to_managed=False,
            container_types=[
                1,
            ],
            option_types=[
                1,
            ],
            spec_templates=[
                1,
            ],
            environment_variables=[
                ClusterLayoutCreateEnvironmentVariablesInner(
                    name="name_example",
                    value="value_example",
                    masked=False,
                    export=False,
                ),
            ],
            price_sets=[
                InstanceTypeCreatePriceSetsInner(
                    id=1,
                ),
            ],
            permissions=UpdateLayoutPermissionsRequestInstanceTypeLayoutPermissions(
                resource_permissions=UpdateLayoutPermissionsRequestInstanceTypeLayoutPermissionsResourcePermissions(
                    all=True,
                    sites=[
                        UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                            id=1,
                        ),
                    ],
                ),
            ),
        ),
    ) # AddLayoutRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Layout
        api_response = api_instance.add_layout(instance_type_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_layout: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Layout
        api_response = api_instance.add_layout(instance_type_id, add_layout_request=add_layout_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_layout: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |
 **add_layout_request** | [**AddLayoutRequest**](AddLayoutRequest.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_node_type**
> AddNodeType200Response add_node_type()

Create a Node Type

Use this command to create a node type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.add_node_type200_response import AddNodeType200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_node_type_request import AddNodeTypeRequest
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
    api_instance = library_api.LibraryApi(api_client)
    add_node_type_request = AddNodeTypeRequest(
        container_type=ContainerTypeCreate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            short_name="short_name_example",
            description="description_example",
            container_version="container_version_example",
            provision_type_code="provision_type_code_example",
            scripts=[
                1,
            ],
            templates=[
                1,
            ],
            virtual_image_id=1,
            stat_type_code="stat_type_code_example",
            log_type_code="log_type_code_example",
            server_type="server_type_example",
            container_ports=[
                ContainerTypeCreateContainerPortsInner(
                    name="name_example",
                    port=1,
                    load_balance_protocol="load_balance_protocol_example",
                ),
            ],
            environment_variables=[
                ClusterLayoutCreateEnvironmentVariablesInner(
                    name="name_example",
                    value="value_example",
                    masked=False,
                    export=False,
                ),
            ],
            config={},
        ),
    ) # AddNodeTypeRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Node Type
        api_response = api_instance.add_node_type(add_node_type_request=add_node_type_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_node_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_node_type_request** | [**AddNodeTypeRequest**](AddNodeTypeRequest.md)|  | [optional]

### Return type

[**AddNodeType200Response**](AddNodeType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_option_list**
> Model200Success add_option_list()

Create an Option List

Use this command to create an option list.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.add_option_list_request import AddOptionListRequest
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    add_option_list_request = AddOptionListRequest(
        option_type_list=OptionTypeListCreate(
            name="name_example",
            description="description_example",
            labels=[
                "labels_example",
            ],
            type="rest",
            source_url="source_url_example",
            visibility="private",
            source_method="GET",
            api_type="api_type_example",
            ignore_ssl_errors=False,
            real_time=False,
            credential=OptionTypeListCreateCredential(
                type="type_example",
                id=1,
            ),
            service_username="service_username_example",
            service_password="service_password_example",
            initial_dataset="initial_dataset_example",
            translation_script="translation_script_example",
            request_script="request_script_example",
            config=OptionTypeListCreateConfig(
                source_headers=[
                    OptionTypeListCreateConfigSourceHeadersInner(
                        name="name_example",
                        value="value_example",
                        masked=False,
                    ),
                ],
            ),
        ),
    ) # AddOptionListRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an Option List
        api_response = api_instance.add_option_list(add_option_list_request=add_option_list_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_option_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_option_list_request** | [**AddOptionListRequest**](AddOptionListRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_option_type**
> SuccessId add_option_type()

Create an Input

Use this command to create an option type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.add_option_type_request import AddOptionTypeRequest
from openapi_client.model.success_id import SuccessId
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
    api_instance = library_api.LibraryApi(api_client)
    add_option_type_request = AddOptionTypeRequest(
        option_type=OptionTypeCreate(
            name="name_example",
            description="description_example",
            labels=[
                "labels_example",
            ],
            field_name="field_name_example",
            type="text",
            field_label="field_label_example",
            placeholder="placeholder_example",
            verify_pattern="verify_pattern_example",
            default_value="default_value_example",
            required=False,
            export_meta=False,
            editable=False,
            option_list=OptionTypeCreateOptionList(
                id=1,
            ),
        ),
    ) # AddOptionTypeRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create an Input
        api_response = api_instance.add_option_type(add_option_type_request=add_option_type_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_option_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_option_type_request** | [**AddOptionTypeRequest**](AddOptionTypeRequest.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_script**
> ScriptSuccessId add_script()

Create a Script

Use this command to create a script.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.script_success_id import ScriptSuccessId
from openapi_client.model.add_script_request import AddScriptRequest
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
    api_instance = library_api.LibraryApi(api_client)
    add_script_request = AddScriptRequest(
        container_script=ScriptCreate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            category="category_example",
            script_version="1",
            script_phase="script_phase_example",
            script_type="bash",
            script="script_example",
            run_as_user="run_as_user_example",
            sudo_user=False,
        ),
    ) # AddScriptRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Script
        api_response = api_instance.add_script(add_script_request=add_script_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_script: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_script_request** | [**AddScriptRequest**](AddScriptRequest.md)|  | [optional]

### Return type

[**ScriptSuccessId**](ScriptSuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_spec_template**
> SuccessId add_spec_template()

Create a Spec Template

Use this command to create a spec template.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.add_spec_template_request import AddSpecTemplateRequest
from openapi_client.model.success_id import SuccessId
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
    api_instance = library_api.LibraryApi(api_client)
    add_spec_template_request = AddSpecTemplateRequest(
        spec_template=SpecTemplateCreate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            type=SpecTemplateCreateType(
                code="code_example",
            ),
            file=SpecTemplateCreateFile(
                source_type="local",
                content="content_example",
                content_path="content_path_example",
                content_ref="content_ref_example",
                repository=SpecTemplateCreateFileRepository(
                    id=1,
                ),
            ),
            config=SpecTemplateCreateConfig(
                cloudformation=SpecTemplateCreateConfigCloudformation(
                    iam=SpecTemplateCreateConfigCloudformationIAM(None),
                    capability_named_iam=SpecTemplateCreateConfigCloudformationCAPABILITYNAMEDIAM(None),
                    capability_auto_expand=SpecTemplateCreateConfigCloudformationCAPABILITYAUTOEXPAND(None),
                ),
            ),
        ),
    ) # AddSpecTemplateRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Spec Template
        api_response = api_instance.add_spec_template(add_spec_template_request=add_spec_template_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_spec_template: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_spec_template_request** | [**AddSpecTemplateRequest**](AddSpecTemplateRequest.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_virtual_image**
> AddVirtualImage200Response add_virtual_image()

Create a Virtual Image

This endpoint creates a new virtual image, without any files yet.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.add_virtual_image_request import AddVirtualImageRequest
from openapi_client.model.add_virtual_image200_response import AddVirtualImage200Response
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
    api_instance = library_api.LibraryApi(api_client)
    add_virtual_image_request = AddVirtualImageRequest(
        virtual_image=VirtualImageCreate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            image_type="image_type_example",
            storage_provider=VirtualImageCreateStorageProvider(
                id=1,
            ),
            is_cloud_init=False,
            user_data="user_data_example",
            install_agent=False,
            ssh_username="ssh_username_example",
            ssh_password="ssh_password_example",
            ssh_key="ssh_key_example",
            os_type=VirtualImageCreateOsType(None),
            visibility="private",
            accounts=[
                1,
            ],
            is_auto_join_domain=False,
            virtio_supported=True,
            vm_tools_installed=True,
            is_force_customization=False,
            trial_version=False,
            is_sysprep=False,
            config=VirtualImageCreateConfig(None),
            tags=[
                VirtualImageCreateTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
        ),
    ) # AddVirtualImageRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Virtual Image
        api_response = api_instance.add_virtual_image(add_virtual_image_request=add_virtual_image_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_virtual_image: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_virtual_image_request** | [**AddVirtualImageRequest**](AddVirtualImageRequest.md)|  | [optional]

### Return type

[**AddVirtualImage200Response**](AddVirtualImage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_virtual_image_file**
> Model200Success add_virtual_image_file(virtual_image_id)

Upload Virtual Image File

This will upload the file and associate it to the Virtual Image.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    virtual_image_id = 4 # float | Virtual Image ID
    filename = "testimage.ovf" # str | The name of the file (optional)
    url = "https://example.com/testimage.ovf" # str | Download the file from a remote url. This can be used instead of uploading a local file. (optional)
    body = open('/path/to/file', 'rb') # file_type |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Upload Virtual Image File
        api_response = api_instance.add_virtual_image_file(virtual_image_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_virtual_image_file: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Upload Virtual Image File
        api_response = api_instance.add_virtual_image_file(virtual_image_id, filename=filename, url=url, body=body)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->add_virtual_image_file: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |
 **filename** | **str**| The name of the file | [optional]
 **url** | **str**| Download the file from a remote url. This can be used instead of uploading a local file. | [optional]
 **body** | **file_type**|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_file_template**
> Model200Success delete_file_template(id)

Delete a File Template

Will delete a file template

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a File Template
        api_response = api_instance.delete_file_template(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->delete_file_template: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_instance_type**
> Model200Success delete_instance_type(instance_type_id)

Delete an Instance Type

Will delete an instance type

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    instance_type_id = 2 # float | The ID of the instance type

    # example passing only required values which don't have defaults set
    try:
        # Delete an Instance Type
        api_response = api_instance.delete_instance_type(instance_type_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->delete_instance_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_layout**
> Model200Success delete_layout(id)

Delete a Layout

Will delete a layout

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Layout
        api_response = api_instance.delete_layout(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->delete_layout: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_node_type**
> Model200Success delete_node_type(id)

Delete a Node Type

Will delete a node type

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Node Type
        api_response = api_instance.delete_node_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->delete_node_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_option_list**
> Model200Success delete_option_list(id)

Delete an Option List

Will delete an option list.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete an Option List
        api_response = api_instance.delete_option_list(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->delete_option_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_option_type**
> Model200Success delete_option_type(id)

Delete an Input

Will delete an option type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete an Input
        api_response = api_instance.delete_option_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->delete_option_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_script**
> Model200Success delete_script(id)

Delete a Script

Will delete a script

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Script
        api_response = api_instance.delete_script(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->delete_script: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_spec_template**
> Model200Success delete_spec_template(id)

Delete a Spec Template

Will delete a spec template

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Spec Template
        api_response = api_instance.delete_spec_template(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->delete_spec_template: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_file_template**
> GetFileTemplate200Response get_file_template(id)

Get a Specific File Template

This endpoint retrieves a specific file template.  The value of template will be masked as ************ for system owned file templates. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.get_file_template200_response import GetFileTemplate200Response
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific File Template
        api_response = api_instance.get_file_template(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_file_template: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetFileTemplate200Response**](GetFileTemplate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_input**
> GetInput200Response get_input(id)

Get A Specific Input

This endpoint retrieves a specific option type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.get_input200_response import GetInput200Response
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get A Specific Input
        api_response = api_instance.get_input(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_input: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetInput200Response**](GetInput200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_instance_type**
> GetInstanceTypeProvisioning200ResponseAllOf get_instance_type(instance_type_id)

Get a Specific Instance Type

This endpoint retrieves a specific instance type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_instance_type_provisioning200_response_all_of import GetInstanceTypeProvisioning200ResponseAllOf
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
    api_instance = library_api.LibraryApi(api_client)
    instance_type_id = 2 # float | The ID of the instance type

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Instance Type
        api_response = api_instance.get_instance_type(instance_type_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_instance_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |

### Return type

[**GetInstanceTypeProvisioning200ResponseAllOf**](GetInstanceTypeProvisioning200ResponseAllOf.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_layout**
> GetLayout200Response get_layout(id)

Get a Specific Layout

This endpoint retrieves a specific layout.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.get_layout200_response import GetLayout200Response
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Layout
        api_response = api_instance.get_layout(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_layout: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetLayout200Response**](GetLayout200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_node_type**
> AddNodeType200ResponseAllOf get_node_type(id)

Get a Specific Node Type

This endpoint retrieves a specific node type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.add_node_type200_response_all_of import AddNodeType200ResponseAllOf
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Node Type
        api_response = api_instance.get_node_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_node_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**AddNodeType200ResponseAllOf**](AddNodeType200ResponseAllOf.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_option_list**
> ListOptionLists200ResponseAllOf get_option_list(id)

Get a Specific Option List

This endpoint retrieves a specific option list.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_option_lists200_response_all_of import ListOptionLists200ResponseAllOf
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Option List
        api_response = api_instance.get_option_list(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_option_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**ListOptionLists200ResponseAllOf**](ListOptionLists200ResponseAllOf.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_option_list_items**
> GetOptionListItems200Response get_option_list_items(id)

List Items for a Specific Option List

This endpoint retrieves the items for a specific option list.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.get_option_list_items200_response import GetOptionListItems200Response
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # List Items for a Specific Option List
        api_response = api_instance.get_option_list_items(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_option_list_items: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetOptionListItems200Response**](GetOptionListItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_script**
> GetScript200Response get_script(id)

Get a Specific Script

This endpoint retrieves a specific script.  The value of script will be masked as ************ for system owned scripts. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_script200_response import GetScript200Response
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Script
        api_response = api_instance.get_script(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_script: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetScript200Response**](GetScript200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_security_package_type**
> GetSecurityPackageType200Response get_security_package_type(id)

Retrieves a Specific Security Package Type

Retrieves a specific security package type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_security_package_type200_response import GetSecurityPackageType200Response
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Security Package Type
        api_response = api_instance.get_security_package_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_security_package_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetSecurityPackageType200Response**](GetSecurityPackageType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_spec_template**
> GetSpecTemplate200Response get_spec_template(id)

Get a Specific Spec Template

This endpoint retrieves a specific spec template.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.get_spec_template200_response import GetSpecTemplate200Response
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Spec Template
        api_response = api_instance.get_spec_template(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_spec_template: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetSpecTemplate200Response**](GetSpecTemplate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_virtual_image**
> AddVirtualImage200ResponseAllOf get_virtual_image(virtual_image_id)

Get a Specific Virtual Image

This endpoint retrieves a specific virtual image and its files.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.add_virtual_image200_response_all_of import AddVirtualImage200ResponseAllOf
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
    api_instance = library_api.LibraryApi(api_client)
    virtual_image_id = 4 # float | Virtual Image ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Virtual Image
        api_response = api_instance.get_virtual_image(virtual_image_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->get_virtual_image: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |

### Return type

[**AddVirtualImage200ResponseAllOf**](AddVirtualImage200ResponseAllOf.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Virtual Image |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_file_templates**
> ListFileTemplates200Response list_file_templates()

Get All File Templates

This endpoint retrieves all file templates.  The value of template will be masked as ************ for system owned file templates. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_file_templates200_response import ListFileTemplates200Response
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
    api_instance = library_api.LibraryApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)
    file_name = "testtext.txt" # str | Filename filter, restricts query to only load file template matching fileName specified (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All File Templates
        api_response = api_instance.list_file_templates(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, labels=labels, all_labels=all_labels, file_name=file_name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_file_templates: %s\n" % e)
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
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **file_name** | **str**| Filename filter, restricts query to only load file template matching fileName specified | [optional]

### Return type

[**ListFileTemplates200Response**](ListFileTemplates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_inputs**
> ListInputs200Response list_inputs()

Get All Inputs

This endpoint retrieves all option types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.list_inputs200_response import ListInputs200Response
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
    api_instance = library_api.LibraryApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)
    field_name = "cloudName" # str | Field Name filter, restricts query to only load type matching fieldName specified (optional)
    field_context = "config.customOptions" # str | Field Context filter, restricts query to only load type matching fieldContext specified (optional)
    field_label = "DB Version" # str | Field Label filter, restricts query to only load type matching fieldLabel specified (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Inputs
        api_response = api_instance.list_inputs(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code, labels=labels, all_labels=all_labels, field_name=field_name, field_context=field_context, field_label=field_label)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_inputs: %s\n" % e)
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
 **code** | **str**| If specified will return an exact match on code | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **field_name** | **str**| Field Name filter, restricts query to only load type matching fieldName specified | [optional]
 **field_context** | **str**| Field Context filter, restricts query to only load type matching fieldContext specified | [optional]
 **field_label** | **str**| Field Label filter, restricts query to only load type matching fieldLabel specified | [optional]

### Return type

[**ListInputs200Response**](ListInputs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_instance_types**
> ListInstanceTypesProvisioning200Response list_instance_types()

Get All Instance Types

This endpoint retrieves all instance types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.list_instance_types_provisioning200_response import ListInstanceTypesProvisioning200Response
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
    api_instance = library_api.LibraryApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    featured = False # bool | Filter by featured (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)
    details = True # bool | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Instance Types
        api_response = api_instance.list_instance_types(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code, featured=featured, labels=labels, all_labels=all_labels, details=details)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_instance_types: %s\n" % e)
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
 **code** | **str**| If specified will return an exact match on code | [optional]
 **featured** | **bool**| Filter by featured | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **details** | **bool**| Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. | [optional]

### Return type

[**ListInstanceTypesProvisioning200Response**](ListInstanceTypesProvisioning200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_layouts**
> ListLayoutsForInstanceType200Response list_layouts()

Get All Layouts

This endpoint retrieves all layouts. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.list_layouts_for_instance_type200_response import ListLayoutsForInstanceType200Response
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
    api_instance = library_api.LibraryApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    provision_type = "AKS" # str | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Layouts
        api_response = api_instance.list_layouts(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code, provision_type=provision_type, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_layouts: %s\n" % e)
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
 **code** | **str**| If specified will return an exact match on code | [optional]
 **provision_type** | **str**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListLayoutsForInstanceType200Response**](ListLayoutsForInstanceType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_layouts_for_instance_type**
> ListLayoutsForInstanceType200Response list_layouts_for_instance_type(instance_type_id)

Get All Layouts For an Instance Type

This endpoint retrieves all layouts for a specific instance type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.list_layouts_for_instance_type200_response import ListLayoutsForInstanceType200Response
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
    api_instance = library_api.LibraryApi(api_client)
    instance_type_id = 2 # float | The ID of the instance type
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    provision_type = "AKS" # str | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get All Layouts For an Instance Type
        api_response = api_instance.list_layouts_for_instance_type(instance_type_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_layouts_for_instance_type: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Layouts For an Instance Type
        api_response = api_instance.list_layouts_for_instance_type(instance_type_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code, provision_type=provision_type, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_layouts_for_instance_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **str**| If specified will return an exact match on code | [optional]
 **provision_type** | **str**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListLayoutsForInstanceType200Response**](ListLayoutsForInstanceType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_node_types**
> ListNodeTypes200Response list_node_types()

Get All Node Types

This endpoint retrieves all node types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.list_node_types200_response import ListNodeTypes200Response
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
    api_instance = library_api.LibraryApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    provision_type = "AKS" # str | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Node Types
        api_response = api_instance.list_node_types(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code, provision_type=provision_type, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_node_types: %s\n" % e)
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
 **code** | **str**| If specified will return an exact match on code | [optional]
 **provision_type** | **str**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListNodeTypes200Response**](ListNodeTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_option_lists**
> ListOptionLists200Response list_option_lists()

Get All Option Lists

This endpoint retrieves all option lists.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.list_option_lists200_response import ListOptionLists200Response
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
    api_instance = library_api.LibraryApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Option Lists
        api_response = api_instance.list_option_lists(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_option_lists: %s\n" % e)
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
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListOptionLists200Response**](ListOptionLists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_scripts**
> ListScripts200Response list_scripts()

Get All Scripts

This endpoint retrieves all scripts.  The value of script will be masked as ************ for system owned scripts. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_scripts200_response import ListScripts200Response
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
    api_instance = library_api.LibraryApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)
    script_type = "scriptType_example" # str | Script type code filter, restricts query to only load scripts of specified type (optional)
    script_phase = "scriptPhase_example" # str | Script phase filter, restricts query to only load scripts of specified phase (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Scripts
        api_response = api_instance.list_scripts(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, labels=labels, all_labels=all_labels, script_type=script_type, script_phase=script_phase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_scripts: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **script_type** | **str**| Script type code filter, restricts query to only load scripts of specified type | [optional]
 **script_phase** | **str**| Script phase filter, restricts query to only load scripts of specified phase | [optional]

### Return type

[**ListScripts200Response**](ListScripts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_security_package_types**
> ListSecurityPackageTypes200Response list_security_package_types()

Retrieves all Security Package Types

Retrieves all Security Package Types 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.list_security_package_types200_response import ListSecurityPackageTypes200Response
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
    api_instance = library_api.LibraryApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Retrieves all Security Package Types
        api_response = api_instance.list_security_package_types()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_security_package_types: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**ListSecurityPackageTypes200Response**](ListSecurityPackageTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_spec_templates**
> ListSpecTemplates200Response list_spec_templates()

Get All Spec Templates

This endpoint retrieves all spec templates.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_spec_templates200_response import ListSpecTemplates200Response
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
    api_instance = library_api.LibraryApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Spec Templates
        api_response = api_instance.list_spec_templates(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_spec_templates: %s\n" % e)
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
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListSpecTemplates200Response**](ListSpecTemplates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_virtual_image_locations**
> ListVirtualImageLocations200Response list_virtual_image_locations(virtual_image_id)

Get a List of Virtual Image Locations

This endpoint retrieves a specific virtual image and its files.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_virtual_image_locations200_response import ListVirtualImageLocations200Response
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
    api_instance = library_api.LibraryApi(api_client)
    virtual_image_id = 4 # float | Virtual Image ID

    # example passing only required values which don't have defaults set
    try:
        # Get a List of Virtual Image Locations
        api_response = api_instance.list_virtual_image_locations(virtual_image_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_virtual_image_locations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |

### Return type

[**ListVirtualImageLocations200Response**](ListVirtualImageLocations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_virtual_images**
> ListVirtualImages200Response list_virtual_images()

Get List of Virtual Images

This endpoint retrieves a list of virtual images for the specified filter.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.list_virtual_images200_response import ListVirtualImages200Response
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
    api_instance = library_api.LibraryApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    last_updated = dateutil_parser('2019-03-06T17:52:29+0000') # datetime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    filter_type = "System" # str | Filter by type, \"User\", \"System\", \"Synced\", or \"All\" (optional) if omitted the server will use the default value of "User"
    image_type = "vmware" # str | Filter by image type code, \"vmware\", \"ami\", etc (optional)
    tags = "tags.env=qa&tags.env=test" # str | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get List of Virtual Images
        api_response = api_instance.list_virtual_images(max=max, offset=offset, name=name, phrase=phrase, last_updated=last_updated, filter_type=filter_type, image_type=image_type, tags=tags, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->list_virtual_images: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **last_updated** | **datetime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **filter_type** | **str**| Filter by type, \&quot;User\&quot;, \&quot;System\&quot;, \&quot;Synced\&quot;, or \&quot;All\&quot; | [optional] if omitted the server will use the default value of "User"
 **image_type** | **str**| Filter by image type code, \&quot;vmware\&quot;, \&quot;ami\&quot;, etc | [optional]
 **tags** | **str**| Filter by tags (metadata). This allows filtering by a tag name and value(s)  | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListVirtualImages200Response**](ListVirtualImages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_security_scans**
> Model200Success remove_security_scans(id)

Deletes a Security Scan

Deletes a specified security scan. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Security Scan
        api_response = api_instance.remove_security_scans(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->remove_security_scans: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_virtual_image**
> Model200Success remove_virtual_image(virtual_image_id)

Delete a Virtual Image

Will delete a virtual image and any associated files, use removeFromCloud=true to also delete image locations from all clouds.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    virtual_image_id = 4 # float | Virtual Image ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Virtual Image
        api_response = api_instance.remove_virtual_image(virtual_image_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->remove_virtual_image: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_virtual_image_file**
> Model200Success remove_virtual_image_file(virtual_image_id)

Remove Virtual Image File

Remove Virtual Image File

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    virtual_image_id = 4 # float | Virtual Image ID
    filename = "testimage.ovf" # str | The name of the file (optional)

    # example passing only required values which don't have defaults set
    try:
        # Remove Virtual Image File
        api_response = api_instance.remove_virtual_image_file(virtual_image_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->remove_virtual_image_file: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Remove Virtual Image File
        api_response = api_instance.remove_virtual_image_file(virtual_image_id, filename=filename)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->remove_virtual_image_file: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |
 **filename** | **str**| The name of the file | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_virtual_image_location**
> Model200Success remove_virtual_image_location(virtual_image_id, id)

Delete a Virtual Image Location

Will delete a virtual image location, use removeFromCloud=true to all also delete image locations from all clouds as well.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    virtual_image_id = 4 # float | Virtual Image ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Virtual Image Location
        api_response = api_instance.remove_virtual_image_location(virtual_image_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->remove_virtual_image_location: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **set_instance_type_featured**
> Model200Success set_instance_type_featured(instance_type_id)

Toggle Featured For Instance Type

Use this command to toggle the featured flag for an existing instance type. This will change the value from false to true, or from true to false. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    instance_type_id = 2 # float | The ID of the instance type

    # example passing only required values which don't have defaults set
    try:
        # Toggle Featured For Instance Type
        api_response = api_instance.set_instance_type_featured(instance_type_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->set_instance_type_featured: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_file_template**
> Model200Success update_file_template(id)

Update a File Template

Use this command to update an existing file template.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_file_template_request import UpdateFileTemplateRequest
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_file_template_request = UpdateFileTemplateRequest(
        container_template=FileTemplateUpdate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            file_name="file_name_example",
            file_path="file_path_example",
            category="category_example",
            template_phase="template_phase_example",
            template="template_example",
            file_owner=1,
            setting_name="setting_name_example",
            setting_category="setting_category_example",
        ),
    ) # UpdateFileTemplateRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a File Template
        api_response = api_instance.update_file_template(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_file_template: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a File Template
        api_response = api_instance.update_file_template(id, update_file_template_request=update_file_template_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_file_template: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_file_template_request** | [**UpdateFileTemplateRequest**](UpdateFileTemplateRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_instance_type**
> Model200Success update_instance_type(instance_type_id)

Update an Instance Type

Use this command to update an existing instance type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_instance_type_request import UpdateInstanceTypeRequest
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
    api_instance = library_api.LibraryApi(api_client)
    instance_type_id = 2 # float | The ID of the instance type
    update_instance_type_request = UpdateInstanceTypeRequest(
        instance_type=InstanceTypeUpdate(
            name="name_example",
            description="description_example",
            labels=[
                "labels_example",
            ],
            code="code_example",
            category="category_example",
            visibility="private",
            featured=True,
            has_settings=True,
            has_auto_scale=True,
            has_deployment=True,
            environment_prefix="environment_prefix_example",
            environment_variables=[
                ClusterLayoutCreateEnvironmentVariablesInner(
                    name="name_example",
                    value="value_example",
                    masked=False,
                    export=False,
                ),
            ],
            price_sets=[
                InstanceTypeCreatePriceSetsInner(
                    id=1,
                ),
            ],
            option_types=[
                1,
            ],
        ),
    ) # UpdateInstanceTypeRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an Instance Type
        api_response = api_instance.update_instance_type(instance_type_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_instance_type: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an Instance Type
        api_response = api_instance.update_instance_type(instance_type_id, update_instance_type_request=update_instance_type_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_instance_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |
 **update_instance_type_request** | [**UpdateInstanceTypeRequest**](UpdateInstanceTypeRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_instance_type_logo**
> Model200Success update_instance_type_logo(instance_type_id)

Update Logo For Instance Type

Use this command to update the logo and dark logo images for an existing instance type. This endpoint expects multipart form data as the request format, not JSON. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    instance_type_id = 2 # float | The ID of the instance type
    logo = open('/path/to/file', 'rb') # file_type | Logo File png,jpg,svg (optional)
    dark_logo = open('/path/to/file', 'rb') # file_type | Dark Logo File png,jpg,svg (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Logo For Instance Type
        api_response = api_instance.update_instance_type_logo(instance_type_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_instance_type_logo: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Logo For Instance Type
        api_response = api_instance.update_instance_type_logo(instance_type_id, logo=logo, dark_logo=dark_logo)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_instance_type_logo: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |
 **logo** | **file_type**| Logo File png,jpg,svg | [optional]
 **dark_logo** | **file_type**| Dark Logo File png,jpg,svg | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_layout**
> Model200Success update_layout(id)

Update a Layout

Use this command to update an existing layout.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.update_layout_request import UpdateLayoutRequest
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_layout_request = UpdateLayoutRequest(
        instance_type_layout=InstanceTypeLayoutUpdate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            instance_version="instance_version_example",
            description="description_example",
            creatable=True,
            provision_type_code="provision_type_code_example",
            memory_requirement="memory_requirement_example",
            has_auto_scale=False,
            supports_convert_to_managed=False,
            container_types=[
                1,
            ],
            option_types=[
                1,
            ],
            spec_templates=[
                1,
            ],
            environment_variables=[
                ClusterLayoutCreateEnvironmentVariablesInner(
                    name="name_example",
                    value="value_example",
                    masked=False,
                    export=False,
                ),
            ],
            price_sets=[
                InstanceTypeCreatePriceSetsInner(
                    id=1,
                ),
            ],
            permissions=UpdateLayoutPermissionsRequestInstanceTypeLayoutPermissions(
                resource_permissions=UpdateLayoutPermissionsRequestInstanceTypeLayoutPermissionsResourcePermissions(
                    all=True,
                    sites=[
                        UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                            id=1,
                        ),
                    ],
                ),
            ),
        ),
    ) # UpdateLayoutRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Layout
        api_response = api_instance.update_layout(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_layout: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Layout
        api_response = api_instance.update_layout(id, update_layout_request=update_layout_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_layout: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_layout_request** | [**UpdateLayoutRequest**](UpdateLayoutRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_layout_permissions**
> Model200Success update_layout_permissions(id)

Update Layout Permissions

Use this command to update permissions for an existing layout.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.update_layout_permissions_request import UpdateLayoutPermissionsRequest
from openapi_client.model.model200_success import Model200Success
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_layout_permissions_request = UpdateLayoutPermissionsRequest(
        instance_type_layout=UpdateLayoutPermissionsRequestInstanceTypeLayout(
            permissions=UpdateLayoutPermissionsRequestInstanceTypeLayoutPermissions(
                resource_permissions=UpdateLayoutPermissionsRequestInstanceTypeLayoutPermissionsResourcePermissions(
                    all=True,
                    sites=[
                        UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                            id=1,
                        ),
                    ],
                ),
            ),
        ),
    ) # UpdateLayoutPermissionsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Layout Permissions
        api_response = api_instance.update_layout_permissions(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_layout_permissions: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Layout Permissions
        api_response = api_instance.update_layout_permissions(id, update_layout_permissions_request=update_layout_permissions_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_layout_permissions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_layout_permissions_request** | [**UpdateLayoutPermissionsRequest**](UpdateLayoutPermissionsRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_node_type**
> Model200Success update_node_type(id)

Update a Node Type

Use this command to update an existing node type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_node_type_request import UpdateNodeTypeRequest
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_node_type_request = UpdateNodeTypeRequest(
        container_type=ContainerTypeUpdate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            short_name="short_name_example",
            description="description_example",
            container_version="container_version_example",
            provision_type_code="provision_type_code_example",
            scripts=[
                1,
            ],
            templates=[
                1,
            ],
            virtual_image_id=1,
            stat_type_code="stat_type_code_example",
            log_type_code="log_type_code_example",
            server_type="server_type_example",
            container_ports=[
                ContainerTypeCreateContainerPortsInner(
                    name="name_example",
                    port=1,
                    load_balance_protocol="load_balance_protocol_example",
                ),
            ],
            environment_variables=[
                ClusterLayoutCreateEnvironmentVariablesInner(
                    name="name_example",
                    value="value_example",
                    masked=False,
                    export=False,
                ),
            ],
            config={},
        ),
    ) # UpdateNodeTypeRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Node Type
        api_response = api_instance.update_node_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_node_type: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Node Type
        api_response = api_instance.update_node_type(id, update_node_type_request=update_node_type_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_node_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_node_type_request** | [**UpdateNodeTypeRequest**](UpdateNodeTypeRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_option_list**
> Model200Success update_option_list(id)

Update an Option List

Use this command to update an existing option list.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_option_list_request import UpdateOptionListRequest
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_option_list_request = UpdateOptionListRequest(
        option_type_list=OptionTypeListUpdate(
            name="name_example",
            description="description_example",
            labels=[
                "labels_example",
            ],
            type="rest",
            source_url="source_url_example",
            visibility="private",
            source_method="GET",
            api_type="api_type_example",
            ignore_ssl_errors=False,
            real_time=False,
            credential=OptionTypeListCreateCredential(
                type="type_example",
                id=1,
            ),
            service_username="service_username_example",
            service_password="service_password_example",
            initial_dataset="initial_dataset_example",
            translation_script="translation_script_example",
            request_script="request_script_example",
            config=OptionTypeListCreateConfig(
                source_headers=[
                    OptionTypeListCreateConfigSourceHeadersInner(
                        name="name_example",
                        value="value_example",
                        masked=False,
                    ),
                ],
            ),
        ),
    ) # UpdateOptionListRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an Option List
        api_response = api_instance.update_option_list(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_option_list: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an Option List
        api_response = api_instance.update_option_list(id, update_option_list_request=update_option_list_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_option_list: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_option_list_request** | [**UpdateOptionListRequest**](UpdateOptionListRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_option_type**
> Model200Success update_option_type(id)

Update an Input

Use this command to update an existing option type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_option_type_request import UpdateOptionTypeRequest
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_option_type_request = UpdateOptionTypeRequest(
        option_type=OptionTypeUpdate(
            name="name_example",
            description="description_example",
            labels=[
                "labels_example",
            ],
            field_name="field_name_example",
            type="text",
            field_label="field_label_example",
            placeholder="placeholder_example",
            verify_pattern="verify_pattern_example",
            default_value="default_value_example",
            required=False,
            export_meta=False,
            editable=False,
            option_list=OptionTypeCreateOptionList(
                id=1,
            ),
        ),
    ) # UpdateOptionTypeRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update an Input
        api_response = api_instance.update_option_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_option_type: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update an Input
        api_response = api_instance.update_option_type(id, update_option_type_request=update_option_type_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_option_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_option_type_request** | [**UpdateOptionTypeRequest**](UpdateOptionTypeRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_script**
> ScriptSuccessId update_script(id)

Update a Script

Use this command to update an existing script.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.script_success_id import ScriptSuccessId
from openapi_client.model.update_script_request import UpdateScriptRequest
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_script_request = UpdateScriptRequest(
        container_script=ScriptUpdate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            category="category_example",
            script_version="1",
            script_phase="script_phase_example",
            script_type="bash",
            script="script_example",
            run_as_user="run_as_user_example",
            sudo_user=False,
        ),
    ) # UpdateScriptRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Script
        api_response = api_instance.update_script(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_script: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Script
        api_response = api_instance.update_script(id, update_script_request=update_script_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_script: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_script_request** | [**UpdateScriptRequest**](UpdateScriptRequest.md)|  | [optional]

### Return type

[**ScriptSuccessId**](ScriptSuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_spec_template**
> Model200Success update_spec_template(id)

Update a Spec Template

Use this command to update an existing spec template.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_spec_template_request import UpdateSpecTemplateRequest
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
    api_instance = library_api.LibraryApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_spec_template_request = UpdateSpecTemplateRequest(
        spec_template=SpecTemplateUpdate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            type=SpecTemplateUpdateType(
                code="code_example",
            ),
            file=SpecTemplateUpdateFile(
                source_type="local",
                content="content_example",
                content_path="content_path_example",
                content_ref="content_ref_example",
                repository=SpecTemplateUpdateFileRepository(
                    id=1,
                ),
            ),
            config=SpecTemplateUpdateConfig(
                cloudformation=SpecTemplateUpdateConfigCloudformation(
                    iam=SpecTemplateCreateConfigCloudformationIAM(None),
                    capability_named_iam=SpecTemplateCreateConfigCloudformationCAPABILITYNAMEDIAM(None),
                    capability_auto_expand=SpecTemplateCreateConfigCloudformationCAPABILITYAUTOEXPAND(None),
                ),
            ),
        ),
    ) # UpdateSpecTemplateRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Spec Template
        api_response = api_instance.update_spec_template(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_spec_template: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Spec Template
        api_response = api_instance.update_spec_template(id, update_spec_template_request=update_spec_template_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_spec_template: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_spec_template_request** | [**UpdateSpecTemplateRequest**](UpdateSpecTemplateRequest.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_virtual_image**
> AddVirtualImage200Response update_virtual_image(virtual_image_id)

Update a Virtual Image

This endpoint updates an existing virtual image.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import library_api
from openapi_client.model.update_virtual_image_request import UpdateVirtualImageRequest
from openapi_client.model.add_virtual_image200_response import AddVirtualImage200Response
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
    api_instance = library_api.LibraryApi(api_client)
    virtual_image_id = 4 # float | Virtual Image ID
    update_virtual_image_request = UpdateVirtualImageRequest(
        virtual_image=VirtualImageUpdate(
            name="name_example",
            labels=[
                "labels_example",
            ],
            image_type="image_type_example",
            storage_provider=VirtualImageCreateStorageProvider(
                id=1,
            ),
            is_cloud_init=False,
            user_data="user_data_example",
            install_agent=False,
            ssh_username="ssh_username_example",
            ssh_password="ssh_password_example",
            ssh_key="ssh_key_example",
            os_type=VirtualImageCreateOsType(None),
            visibility="private",
            accounts=[
                1,
            ],
            is_auto_join_domain=False,
            virtio_supported=True,
            vm_tools_installed=True,
            is_force_customization=False,
            trial_version=False,
            is_sysprep=False,
            config=VirtualImageCreateConfig(None),
            tags=[
                VirtualImageCreateTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
            add_tags=[
                VirtualImageCreateTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
            remove_tags=[
                VirtualImageUpdateRemoveTagsInner(
                    name="name_example",
                    value="value_example",
                ),
            ],
        ),
    ) # UpdateVirtualImageRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Virtual Image
        api_response = api_instance.update_virtual_image(virtual_image_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_virtual_image: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Virtual Image
        api_response = api_instance.update_virtual_image(virtual_image_id, update_virtual_image_request=update_virtual_image_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling LibraryApi->update_virtual_image: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |
 **update_virtual_image_request** | [**UpdateVirtualImageRequest**](UpdateVirtualImageRequest.md)|  | [optional]

### Return type

[**AddVirtualImage200Response**](AddVirtualImage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

