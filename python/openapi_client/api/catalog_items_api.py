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
from openapi_client.model.add_catalog_item_type200_response import AddCatalogItemType200Response
from openapi_client.model.add_catalog_item_type_request import AddCatalogItemTypeRequest
from openapi_client.model.add_catalog_item_type_request1 import AddCatalogItemTypeRequest1
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_catalog_item_type200_response import GetCatalogItemType200Response
from openapi_client.model.list_catalog_item_types200_response import ListCatalogItemTypes200Response
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_catalog_item_type200_response import UpdateCatalogItemType200Response
from openapi_client.model.update_catalog_item_type_request import UpdateCatalogItemTypeRequest


class CatalogItemsApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.add_catalog_item_type_endpoint = _Endpoint(
            settings={
                'response_type': (AddCatalogItemType200Response,),
                'auth': [
                    'bearerAuth'
                ],
                'endpoint_path': '/api/catalog-item-types',
                'operation_id': 'add_catalog_item_type',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'add_catalog_item_type_request',
                ],
                'required': [],
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
                    'add_catalog_item_type_request':
                        (AddCatalogItemTypeRequest,),
                },
                'attribute_map': {
                },
                'location_map': {
                    'add_catalog_item_type_request': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'application/json',
                    'multipart/form-data'
                ]
            },
            api_client=api_client
        )
        self.get_catalog_item_type_endpoint = _Endpoint(
            settings={
                'response_type': (GetCatalogItemType200Response,),
                'auth': [
                    'bearerAuth'
                ],
                'endpoint_path': '/api/catalog-item-types/{id}',
                'operation_id': 'get_catalog_item_type',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'id',
                ],
                'required': [
                    'id',
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
                    'id':
                        (int,),
                },
                'attribute_map': {
                    'id': 'id',
                },
                'location_map': {
                    'id': 'path',
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
        self.list_catalog_item_types_endpoint = _Endpoint(
            settings={
                'response_type': (ListCatalogItemTypes200Response,),
                'auth': [
                    'bearerAuth'
                ],
                'endpoint_path': '/api/catalog-item-types',
                'operation_id': 'list_catalog_item_types',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'max',
                    'offset',
                    'sort',
                    'direction',
                    'phrase',
                    'name',
                    'description',
                    'enabled',
                    'featured',
                    'labels',
                    'all_labels',
                    'code',
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
                    'max':
                        (int,),
                    'offset':
                        (int,),
                    'sort':
                        (str,),
                    'direction':
                        (str,),
                    'phrase':
                        (str,),
                    'name':
                        (str,),
                    'description':
                        (str,),
                    'enabled':
                        (bool,),
                    'featured':
                        (bool,),
                    'labels':
                        (str,),
                    'all_labels':
                        (str,),
                    'code':
                        (str,),
                },
                'attribute_map': {
                    'max': 'max',
                    'offset': 'offset',
                    'sort': 'sort',
                    'direction': 'direction',
                    'phrase': 'phrase',
                    'name': 'name',
                    'description': 'description',
                    'enabled': 'enabled',
                    'featured': 'featured',
                    'labels': 'labels',
                    'all_labels': 'allLabels',
                    'code': 'code',
                },
                'location_map': {
                    'max': 'query',
                    'offset': 'query',
                    'sort': 'query',
                    'direction': 'query',
                    'phrase': 'query',
                    'name': 'query',
                    'description': 'query',
                    'enabled': 'query',
                    'featured': 'query',
                    'labels': 'query',
                    'all_labels': 'query',
                    'code': 'query',
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
        self.remove_catalog_item_type_endpoint = _Endpoint(
            settings={
                'response_type': (Model200Success,),
                'auth': [
                    'bearerAuth'
                ],
                'endpoint_path': '/api/catalog-item-types/{id}',
                'operation_id': 'remove_catalog_item_type',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'id',
                ],
                'required': [
                    'id',
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
                    'id':
                        (int,),
                },
                'attribute_map': {
                    'id': 'id',
                },
                'location_map': {
                    'id': 'path',
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
        self.update_catalog_item_type_endpoint = _Endpoint(
            settings={
                'response_type': (UpdateCatalogItemType200Response,),
                'auth': [
                    'bearerAuth'
                ],
                'endpoint_path': '/api/catalog-item-types/{id}',
                'operation_id': 'update_catalog_item_type',
                'http_method': 'PUT',
                'servers': None,
            },
            params_map={
                'all': [
                    'id',
                    'update_catalog_item_type_request',
                ],
                'required': [
                    'id',
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
                    'id':
                        (int,),
                    'update_catalog_item_type_request':
                        (UpdateCatalogItemTypeRequest,),
                },
                'attribute_map': {
                    'id': 'id',
                },
                'location_map': {
                    'id': 'path',
                    'update_catalog_item_type_request': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'application/json',
                    'multipart/form-data'
                ]
            },
            api_client=api_client
        )
        self.update_catalog_item_type_logo_endpoint = _Endpoint(
            settings={
                'response_type': (UpdateCatalogItemType200Response,),
                'auth': [
                    'bearerAuth'
                ],
                'endpoint_path': '/api/catalog-item-types/{id}/update-logo',
                'operation_id': 'update_catalog_item_type_logo',
                'http_method': 'PUT',
                'servers': None,
            },
            params_map={
                'all': [
                    'id',
                    'catalog_item_type_logo',
                    'catalog_item_type_dark_logo',
                ],
                'required': [
                    'id',
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
                    'id':
                        (int,),
                    'catalog_item_type_logo':
                        (file_type,),
                    'catalog_item_type_dark_logo':
                        (file_type,),
                },
                'attribute_map': {
                    'id': 'id',
                    'catalog_item_type_logo': 'catalogItemType.logo',
                    'catalog_item_type_dark_logo': 'catalogItemType.darkLogo',
                },
                'location_map': {
                    'id': 'path',
                    'catalog_item_type_logo': 'form',
                    'catalog_item_type_dark_logo': 'form',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'multipart/form-data'
                ]
            },
            api_client=api_client
        )

    def add_catalog_item_type(
        self,
        **kwargs
    ):
        """Create a Catalog Item Type  # noqa: E501

        Use this command to create a catalog item type.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.add_catalog_item_type(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            add_catalog_item_type_request (AddCatalogItemTypeRequest): [optional]
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
            AddCatalogItemType200Response
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
        return self.add_catalog_item_type_endpoint.call_with_http_info(**kwargs)

    def get_catalog_item_type(
        self,
        id,
        **kwargs
    ):
        """Get a Specific Catalog Item Type  # noqa: E501

        This endpoint retrieves a specific category item type.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_catalog_item_type(id, async_req=True)
        >>> result = thread.get()

        Args:
            id (int): Morpheus ID of the Object being referenced

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
            GetCatalogItemType200Response
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
        kwargs['id'] = \
            id
        return self.get_catalog_item_type_endpoint.call_with_http_info(**kwargs)

    def list_catalog_item_types(
        self,
        **kwargs
    ):
        """Get All Catalog Item Types  # noqa: E501

        This endpoint retrieves all catalog item types.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.list_catalog_item_types(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            max (int): Maximum number of records to return, -1 can be used to fetch all records. [optional] if omitted the server will use the default value of 25
            offset (int): Offset records, the number of records to skip, for paginating requests. [optional] if omitted the server will use the default value of 0
            sort (str): Sort order, the name of the property to sort by. [optional] if omitted the server will use the default value of "name"
            direction (str): Sort direction, use 'desc' to reverse sort. [optional] if omitted the server will use the default value of "asc"
            phrase (str): Search phrase for partial matches on name or description. [optional]
            name (str): Filter by name, wildcard may be specified as %, eg. example-%. [optional]
            description (str): Filter by description, wildcard may be specified as %. eg. `example-%`. [optional]
            enabled (bool): Filter by enabled. [optional]
            featured (bool): Filter by featured. [optional]
            labels (str): Filter by label(s), matches records that contain any of the specified labels. [optional]
            all_labels (str): Filter by label(s), matches records that contain all of the specified labels. [optional]
            code (str): If specified will return an exact match on code. [optional]
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
            ListCatalogItemTypes200Response
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
        return self.list_catalog_item_types_endpoint.call_with_http_info(**kwargs)

    def remove_catalog_item_type(
        self,
        id,
        **kwargs
    ):
        """Delete a Catalog Item Type  # noqa: E501

        Will delete a catalog item type.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.remove_catalog_item_type(id, async_req=True)
        >>> result = thread.get()

        Args:
            id (int): Morpheus ID of the Object being referenced

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
        kwargs['id'] = \
            id
        return self.remove_catalog_item_type_endpoint.call_with_http_info(**kwargs)

    def update_catalog_item_type(
        self,
        id,
        **kwargs
    ):
        """Update a Catalog Item Type  # noqa: E501

        Use this command to update an existing catalog item type.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.update_catalog_item_type(id, async_req=True)
        >>> result = thread.get()

        Args:
            id (int): Morpheus ID of the Object being referenced

        Keyword Args:
            update_catalog_item_type_request (UpdateCatalogItemTypeRequest): [optional]
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
            UpdateCatalogItemType200Response
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
        kwargs['id'] = \
            id
        return self.update_catalog_item_type_endpoint.call_with_http_info(**kwargs)

    def update_catalog_item_type_logo(
        self,
        id,
        **kwargs
    ):
        """Update Logo For Catalog Item Type  # noqa: E501

        Use this command to update the logo and dark logo images for an existing catalog item type. This endpoint expects multipart form data as the request format, not JSON.  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.update_catalog_item_type_logo(id, async_req=True)
        >>> result = thread.get()

        Args:
            id (int): Morpheus ID of the Object being referenced

        Keyword Args:
            catalog_item_type_logo (file_type): Logo File png,jpg,svg. [optional]
            catalog_item_type_dark_logo (file_type): Dark Logo File png,jpg,svg. [optional]
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
            UpdateCatalogItemType200Response
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
        kwargs['id'] = \
            id
        return self.update_catalog_item_type_logo_endpoint.call_with_http_info(**kwargs)
