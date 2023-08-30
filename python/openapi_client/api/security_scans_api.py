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
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_security_scans200_response import GetSecurityScans200Response
from openapi_client.model.list_security_scans200_response import ListSecurityScans200Response


class SecurityScansApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.get_security_scans_endpoint = _Endpoint(
            settings={
                'response_type': (GetSecurityScans200Response,),
                'auth': [
                    'bearerAuth'
                ],
                'endpoint_path': '/api/security-scans/{id}',
                'operation_id': 'get_security_scans',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'id',
                    'results',
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
                    'results':
                        (bool,),
                },
                'attribute_map': {
                    'id': 'id',
                    'results': 'results',
                },
                'location_map': {
                    'id': 'path',
                    'results': 'query',
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
        self.list_security_scans_endpoint = _Endpoint(
            settings={
                'response_type': (ListSecurityScans200Response,),
                'auth': [
                    'bearerAuth'
                ],
                'endpoint_path': '/api/security-scans',
                'operation_id': 'list_security_scans',
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
                    'security_package_id',
                    'server_id',
                    'results',
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
                    'security_package_id':
                        (int,),
                    'server_id':
                        (int,),
                    'results':
                        (bool,),
                },
                'attribute_map': {
                    'max': 'max',
                    'offset': 'offset',
                    'sort': 'sort',
                    'direction': 'direction',
                    'phrase': 'phrase',
                    'security_package_id': 'securityPackageId',
                    'server_id': 'serverId',
                    'results': 'results',
                },
                'location_map': {
                    'max': 'query',
                    'offset': 'query',
                    'sort': 'query',
                    'direction': 'query',
                    'phrase': 'query',
                    'security_package_id': 'query',
                    'server_id': 'query',
                    'results': 'query',
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

    def get_security_scans(
        self,
        id,
        **kwargs
    ):
        """Retrieves a Specific Security Scan  # noqa: E501

        Retrieves a specific security scan.   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.get_security_scans(id, async_req=True)
        >>> result = thread.get()

        Args:
            id (int): Morpheus ID of the Object being referenced

        Keyword Args:
            results (bool): Include the `results` object in the response under the security scan. This is a potentially very large object containing the raw results of the scan.. [optional] if omitted the server will use the default value of False
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
            GetSecurityScans200Response
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
        return self.get_security_scans_endpoint.call_with_http_info(**kwargs)

    def list_security_scans(
        self,
        **kwargs
    ):
        """Retrieves all Security Scans  # noqa: E501

        Retrieves all security scans.   # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.list_security_scans(async_req=True)
        >>> result = thread.get()


        Keyword Args:
            max (int): Maximum number of records to return, -1 can be used to fetch all records. [optional] if omitted the server will use the default value of 25
            offset (int): Offset records, the number of records to skip, for paginating requests. [optional] if omitted the server will use the default value of 0
            sort (str): Sort order, the name of the property to sort by. [optional] if omitted the server will use the default value of "scanDate"
            direction (str): Sort direction, use 'desc' to reverse sort. [optional] if omitted the server will use the default value of "desc"
            phrase (str): Search phrase for partial matches on name or description of security package. [optional]
            security_package_id (int): Filter results by security package id(s). This parameter can be passed multiple times to match more than one id.. [optional]
            server_id (int): The Server ID for Filtering. [optional]
            results (bool): Include the `results` object in the response under each security scan. This is a potentially very large object containing the raw results of the scan.. [optional] if omitted the server will use the default value of False
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
            ListSecurityScans200Response
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
        return self.list_security_scans_endpoint.call_with_http_info(**kwargs)

