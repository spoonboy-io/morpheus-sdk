"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from openapi_client.api_client import ApiClient, Endpoint as _Endpoint
from openapi_client.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from openapi_client.model.add_cypher_key200_response import AddCypherKey200Response
from openapi_client.model.add_cypher_key_request import AddCypherKeyRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_cypher_key200_response import GetCypherKey200Response
from openapi_client.model.list_cypher_keys200_response import ListCypherKeys200Response
from openapi_client.model.model200_success import Model200Success


class CypherApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.add_cypher_key_endpoint = _Endpoint(
            settings={
                'response_type': (AddCypherKey200Response,),
                'auth': [
                    'bearerAuth'
                ],
                'endpoint_path': '/api/cypher/{cypherPath}',
                'operation_id': 'add_cypher_key',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'cypher_path',
                    'ttl',
                    'value',
                    'type',
                    'add_cypher_key_request',
                ],
                'required': [
                    'cypher_path',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'cypher_path':
                        (str,),
                    'ttl':
                        (AddCypherKeyTtlParameter,),
                    'value':
                        (str,),
                    'type':
                        (str,),
                    'add_cypher_key_request':
                        (AddCypherKeyRequest,),
                },
                'attribute_map': {
                    'cypher_path': 'cypherPath',
                    'ttl': 'ttl',
                    'value': 'value',
                    'type': 'type',
                },
                'location_map': {
                    'cypher_path': 'path',
                    'ttl': 'query',
                    'value': 'query',
                    'type': 'query',
                    'add_cypher_key_request': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client
        )
        self.get_cypher_key_endpoint = _Endpoint(
            settings={
                'response_type': (GetCypherKey200Response,),
                'auth': [
                    'bearerAuth',
                    'cypherAuth-XCToken',
                    'cypherAuth-XMLease',
                    'cypherAuth-XVToken'
                ],
                'endpoint_path': '/api/cypher/{cypherPath}',
                'operation_id': 'get_cypher_key',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'cypher_path',
                    'lease_token',
                    'sort',
                    'direction',
                ],
                'required': [
                    'cypher_path',
                ],
                'nullable': [
                ],
                'enum': [
                    'direction',
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                    ('direction',): {

                        "ASC": "asc",
                        "DESC": "desc"
                    },
                },
                'openapi_types': {
                    'cypher_path':
                        (str,),
                    'lease_token':
                        (str,),
                    'sort':
                        (str,),
                    'direction':
                        (str,),
                },
                'attribute_map': {
                    'cypher_path': 'cypherPath',
                    'lease_token': 'leaseToken',
                    'sort': 'sort',
                    'direction': 'direction',
                },
                'location_map': {
                    'cypher_path': 'path',
                    'lease_token': 'query',
                    'sort': 'query',
                    'direction': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.list_cypher_keys_endpoint = _Endpoint(
            settings={
                'response_type': (ListCypherKeys200Response,),
                'auth': [
                    'bearerAuth',
                    'cypherAuth-XCToken',
                    'cypherAuth-XMLease',
                    'cypherAuth-XVToken'
                ],
                'endpoint_path': '/api/cypher',
                'operation_id': 'list_cypher_keys',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'lease_token',
                    'list',
                    'phrase',
                    'max',
                    'offset',
                    'sort',
                    'direction',
                ],
                'required': [],
                'nullable': [
                ],
                'enum': [
                    'direction',
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                    ('direction',): {

                        "ASC": "asc",
                        "DESC": "desc"
                    },
                },
                'openapi_types': {
                    'lease_token':
                        (str,),
                    'list':
                        (bool,),
                    'phrase':
                        (str,),
                    'max':
                        (int,),
                    'offset':
                        (int,),
                    'sort':
                        (str,),
                    'direction':
                        (str,),
                },
                'attribute_map': {
                    'lease_token': 'leaseToken',
                    'list': 'list',
                    'phrase': 'phrase',
                    'max': 'max',
                    'offset': 'offset',
                    'sort': 'sort',
                    'direction': 'direction',
                },
                'location_map': {
                    'lease_token': 'query',
                    'list': 'query',
                    'phrase': 'query',
                    'max': 'query',
                    'offset': 'query',
                    'sort': 'query',
                    'direction': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )
        self.remove_cypher_endpoint = _Endpoint(
            settings={
                'response_type': (Model200Success,),
                'auth': [
                    'bearerAuth'
                ],
                'endpoint_path': '/api/cypher/{cypherPath}',
                'operation_id': 'remove_cypher',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'cypher_path',
                ],
                'required': [
                    'cypher_path',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'cypher_path':
                        (str,),
                },
                'attribute_map': {
                    'cypher_path': 'cypherPath',
                },
                'location_map': {
                    'cypher_path': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client
        )

    def add_cypher_key(
        self,
        cypher_path,
        **kwargs
    ):
        """Write a Cypher  # noqa: E501

        This endpoint will create or update a cypher key.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.add_cypher_key(cypher_path, async_req=True)
        >>> result = thread.get()

        Args:
            cypher_path (str): The key includes a mount prefix separated by a /. For example, the key secret/foo uses the secret mount.  Available Mounts  <table>   <tr>     <th>Mount</th>     <th>Description</th>     <th>Example</th>   </tr>   <tr>     <td>password</td>     <td>Generates a secure password of specified character length in the key pattern (or 15) with symbols, numbers, upper case, and lower case letters (i.e. password/15/mypass generates a 15 character password).</td>     <td>password/15/mypass</td>   </tr>   <tr>     <td>tfvars</td>     <td>This is a module to store a tfvars file for terraform.</td>     <td>tfvars/mytfvar</td>   </tr>   <tr>     <td>secret</td>     <td>This is the standard secret module that stores a key/value in encrypted form. Capable of storing entire JSON object or a String.</td>     <td>secret/foo</td>   </tr>   <tr>     <td>uuid</td>     <td>Returns a new UUID by key name when requested and stores the generated UUID by key name for a given lease timeout period.</td>     <td>uuid/autoMac1</td>   </tr>   <tr>     <td>key</td>     <td>Generates a Base 64 encoded AES Key of specified bit length in the key pattern (i.e. key/128/mykey generates a 128-bit key)</td>     <td>key/128/mykey</td>   </tr> </table> 

        Keyword Args:
            ttl (AddCypherKeyTtlParameter): Time to Live. The lease duration in seconds, or a human readable format eg. '15m', 8h, '7d'.  0 means no expiry.. [optional]
            value (str): The secret value to be stored. Only required for certain mounts. Some mounts generate their own value and do not require a value to be passed. eg. `uuid`, `key` and `password`.. [optional]
            type (str): The type of data being stored, `string` or `object`. The data type depends on the cypher mount being used. Most mounts use `string` as their data type, but `secret` uses `object` by default. You can store a string instead by passing `type=string`. This means the `data` value returned by the API will be a string instead of an object.. [optional]
            add_cypher_key_request (AddCypherKeyRequest): [optional]
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            AddCypherKey200Response
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        kwargs['cypher_path'] = \
            cypher_path
        return self.add_cypher_key_endpoint.call_with_http_info(**kwargs)

    def get_cypher_key(
        self,
        cypher_path,
        **kwargs
    ):
        """Read or Create a Cypher Key  # noqa: E501

        This endpoint retrieves a specific cypher key. The value of the key is decrypted and returned as data. It may be a String or an object with many {\"key\":\"value\"} pairs.  The type depends on the cypher mount's capabilities and what type of data was written to the key.  For example the `secret/` mount allows either a string or an object, while the `password/` mount will always store and return a string. This endpoint can also create a key. This only applies to mount types `uuid`, `key`, `password`.  Refer to the `POST` endpoint for more information.   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_cypher_key(cypher_path, async_req=True)
        >>> result = thread.get()

        Args:
            cypher_path (str): The cypher key including the mount prefix.

        Keyword Args:
            lease_token (str): An execution lease token.. [optional]
            sort (str): Sort order, the name of the property to sort by. [optional] if omitted the server will use the default value of "name"
            direction (str): Sort direction, use 'desc' to reverse sort. [optional] if omitted the server will use the default value of "asc"
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            GetCypherKey200Response
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        kwargs['cypher_path'] = \
            cypher_path
        return self.get_cypher_key_endpoint.call_with_http_info(**kwargs)

    def list_cypher_keys(
        self,
        **kwargs
    ):
        """List Cypher Keys  # noqa: E501

        This endpoint retrieves all cypher keys associated with the account, or user.  This method can be used to list keys as well, by passing the query parameter list=true.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.list_cypher_keys(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            lease_token (str): An execution lease token.. [optional]
            list (bool): This endpoint is available via the http method LIST. The GET method can be used to list keys as well, by passing the query parameter list=true.. [optional] if omitted the server will use the default value of False
            phrase (str): Search phrase for partial matches on name or description. [optional]
            max (int): Maximum number of records to return, -1 can be used to fetch all records. [optional] if omitted the server will use the default value of 25
            offset (int): Offset records, the number of records to skip, for paginating requests. [optional] if omitted the server will use the default value of 0
            sort (str): Sort order, the name of the property to sort by. [optional] if omitted the server will use the default value of "name"
            direction (str): Sort direction, use 'desc' to reverse sort. [optional] if omitted the server will use the default value of "asc"
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            ListCypherKeys200Response
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        return self.list_cypher_keys_endpoint.call_with_http_info(**kwargs)

    def remove_cypher(
        self,
        cypher_path,
        **kwargs
    ):
        """Delete a Cypher  # noqa: E501

        Will delete a cypher from the system and make it no longer usable.   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.remove_cypher(cypher_path, async_req=True)
        >>> result = thread.get()

        Args:
            cypher_path (str): The cypher key including the mount prefix.

        Keyword Args:
            _return_http_data_only (bool): response data without head status
                code and headers. Default is True.
            _preload_content (bool): if False, the urllib3.HTTPResponse object
                will be returned without reading/decoding response data.
                Default is True.
            _request_timeout (int/float/tuple): timeout setting for this request. If
                one number provided, it will be total request timeout. It can also
                be a pair (tuple) of (connection, read) timeouts.
                Default is None.
            _check_input_type (bool): specifies if type checking
                should be done one the data sent to the server.
                Default is True.
            _check_return_type (bool): specifies if type checking
                should be done one the data received from the server.
                Default is True.
            _spec_property_naming (bool): True if the variable names in the input data
                are serialized names, as specified in the OpenAPI document.
                False if the variable names in the input data
                are pythonic names, e.g. snake case (default)
            _content_type (str/None): force body content-type.
                Default is None and content-type will be predicted by allowed
                content-types and body.
            _host_index (int/None): specifies the index of the server
                that we want to use.
                Default is read from the configuration.
            _request_auths (list): set to override the auth_settings for an a single
                request; this effectively ignores the authentication
                in the spec for a single request.
                Default is None
            async_req (bool): execute request asynchronously

        Returns:
            Model200Success
                If the method is called asynchronously, returns the request
                thread.
        """
        kwargs['async_req'] = kwargs.get(
            'async_req', False
        )
        kwargs['_return_http_data_only'] = kwargs.get(
            '_return_http_data_only', True
        )
        kwargs['_preload_content'] = kwargs.get(
            '_preload_content', True
        )
        kwargs['_request_timeout'] = kwargs.get(
            '_request_timeout', None
        )
        kwargs['_check_input_type'] = kwargs.get(
            '_check_input_type', True
        )
        kwargs['_check_return_type'] = kwargs.get(
            '_check_return_type', True
        )
        kwargs['_spec_property_naming'] = kwargs.get(
            '_spec_property_naming', False
        )
        kwargs['_content_type'] = kwargs.get(
            '_content_type')
        kwargs['_host_index'] = kwargs.get('_host_index')
        kwargs['_request_auths'] = kwargs.get('_request_auths', None)
        kwargs['cypher_path'] = \
            cypher_path
        return self.remove_cypher_endpoint.call_with_http_info(**kwargs)

