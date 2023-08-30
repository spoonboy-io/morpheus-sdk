"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from openapi_client.model_utils import (  # noqa: F401
    ApiTypeError,
    ModelComposed,
    ModelNormal,
    ModelSimple,
    cached_property,
    change_keys_js_to_python,
    convert_js_args_to_python_args,
    date,
    datetime,
    file_type,
    none_type,
    validate_get_composed_info,
    OpenApiModel
)
from openapi_client.exceptions import ApiAttributeError


def lazy_import():
    from openapi_client.model.option_type_list_create_config import OptionTypeListCreateConfig
    from openapi_client.model.option_type_list_create_credential import OptionTypeListCreateCredential
    globals()['OptionTypeListCreateConfig'] = OptionTypeListCreateConfig
    globals()['OptionTypeListCreateCredential'] = OptionTypeListCreateCredential


class OptionTypeListCreate(ModelNormal):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.

    Attributes:
      allowed_values (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          with a capitalized key describing the allowed value and an allowed
          value. These dicts store the allowed enum values.
      attribute_map (dict): The key is attribute name
          and the value is json key in definition.
      discriminator_value_class_map (dict): A dict to go from the discriminator
          variable value to the discriminator class name.
      validations (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          that stores validations for max_length, min_length, max_items,
          min_items, exclusive_maximum, inclusive_maximum, exclusive_minimum,
          inclusive_minimum, and regex.
      additional_properties_type (tuple): A tuple of classes accepted
          as additional properties values.
    """

    allowed_values = {
        ('type',): {
            'REST': "rest",
            'API': "api",
            'LDAP': "ldap",
            'MANUAL': "manual",
        },
        ('visibility',): {
            'PRIVATE': "private",
            'PUBLIC': "public",
        },
        ('source_method',): {
            'GET': "GET",
            'POST': "POST",
        },
    }

    validations = {
    }

    @cached_property
    def additional_properties_type():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded
        """
        lazy_import()
        return (bool, date, datetime, dict, float, int, list, str, none_type,)  # noqa: E501

    _nullable = False

    @cached_property
    def openapi_types():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded

        Returns
            openapi_types (dict): The key is attribute name
                and the value is attribute type.
        """
        lazy_import()
        return {
            'name': (str,),  # noqa: E501
            'description': (str, none_type,),  # noqa: E501
            'labels': ([str], none_type,),  # noqa: E501
            'type': (str,),  # noqa: E501
            'source_url': (str,),  # noqa: E501
            'visibility': (str,),  # noqa: E501
            'source_method': (str,),  # noqa: E501
            'api_type': (str, none_type,),  # noqa: E501
            'ignore_ssl_errors': (bool,),  # noqa: E501
            'real_time': (bool,),  # noqa: E501
            'credential': (OptionTypeListCreateCredential,),  # noqa: E501
            'service_username': (str, none_type,),  # noqa: E501
            'service_password': (str, none_type,),  # noqa: E501
            'initial_dataset': (str, none_type,),  # noqa: E501
            'translation_script': (str, none_type,),  # noqa: E501
            'request_script': (str, none_type,),  # noqa: E501
            'config': (OptionTypeListCreateConfig,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'name': 'name',  # noqa: E501
        'description': 'description',  # noqa: E501
        'labels': 'labels',  # noqa: E501
        'type': 'type',  # noqa: E501
        'source_url': 'sourceUrl',  # noqa: E501
        'visibility': 'visibility',  # noqa: E501
        'source_method': 'sourceMethod',  # noqa: E501
        'api_type': 'apiType',  # noqa: E501
        'ignore_ssl_errors': 'ignoreSSLErrors',  # noqa: E501
        'real_time': 'realTime',  # noqa: E501
        'credential': 'credential',  # noqa: E501
        'service_username': 'serviceUsername',  # noqa: E501
        'service_password': 'servicePassword',  # noqa: E501
        'initial_dataset': 'initialDataset',  # noqa: E501
        'translation_script': 'translationScript',  # noqa: E501
        'request_script': 'requestScript',  # noqa: E501
        'config': 'config',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, name, *args, **kwargs):  # noqa: E501
        """OptionTypeListCreate - a model defined in OpenAPI

        Args:
            name (str): Name

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            description (str, none_type): Description. [optional]  # noqa: E501
            labels ([str], none_type): Array of label strings, can be used for filtering.. [optional]  # noqa: E501
            type (str): Option List Type eg. `rest`, `api`, `ldap` or `manual`.. [optional] if omitted the server will use the default value of "rest"  # noqa: E501
            source_url (str): Source URL, the http(s) URL to request data from. Required when type is rest.. [optional]  # noqa: E501
            visibility (str): Visibility. [optional] if omitted the server will use the default value of "private"  # noqa: E501
            source_method (str): Source Method, the HTTP method.. [optional] if omitted the server will use the default value of "GET"  # noqa: E501
            api_type (str, none_type): Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api.. [optional]  # noqa: E501
            ignore_ssl_errors (bool): Ignore SSL Errors.. [optional] if omitted the server will use the default value of False  # noqa: E501
            real_time (bool): Real Time.. [optional] if omitted the server will use the default value of False  # noqa: E501
            credential (OptionTypeListCreateCredential): [optional]  # noqa: E501
            service_username (str, none_type): Username for authenticating with Basic Auth when type is rest or ldap username.. [optional]  # noqa: E501
            service_password (str, none_type): Password for authenticating with Basic Auth when type is rest or ldap password.. [optional]  # noqa: E501
            initial_dataset (str, none_type): Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties 'name', and 'value'. Required when type is manual.. [optional]  # noqa: E501
            translation_script (str, none_type): Translation Script. Create a js script to translate the result data object into an Array containing objects with properties 'name' and 'value'. The input data is provided as data and the result should be put on the global variable results.. [optional]  # noqa: E501
            request_script (str, none_type): Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties 'name' and 'value' for a get. The input data is provided as data and the result should be put on the global variable results.. [optional]  # noqa: E501
            config (OptionTypeListCreateConfig): [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', True)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        self = super(OpenApiModel, cls).__new__(cls)

        if args:
            for arg in args:
                if isinstance(arg, dict):
                    kwargs.update(arg)
                else:
                    raise ApiTypeError(
                        "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                            args,
                            self.__class__.__name__,
                        ),
                        path_to_item=_path_to_item,
                        valid_classes=(self.__class__,),
                    )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        self.name = name
        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
        return self

    required_properties = set([
        '_data_store',
        '_check_type',
        '_spec_property_naming',
        '_path_to_item',
        '_configuration',
        '_visited_composed_classes',
    ])

    @convert_js_args_to_python_args
    def __init__(self, name, *args, **kwargs):  # noqa: E501
        """OptionTypeListCreate - a model defined in OpenAPI

        Args:
            name (str): Name

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            description (str, none_type): Description. [optional]  # noqa: E501
            labels ([str], none_type): Array of label strings, can be used for filtering.. [optional]  # noqa: E501
            type (str): Option List Type eg. `rest`, `api`, `ldap` or `manual`.. [optional] if omitted the server will use the default value of "rest"  # noqa: E501
            source_url (str): Source URL, the http(s) URL to request data from. Required when type is rest.. [optional]  # noqa: E501
            visibility (str): Visibility. [optional] if omitted the server will use the default value of "private"  # noqa: E501
            source_method (str): Source Method, the HTTP method.. [optional] if omitted the server will use the default value of "GET"  # noqa: E501
            api_type (str, none_type): Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api.. [optional]  # noqa: E501
            ignore_ssl_errors (bool): Ignore SSL Errors.. [optional] if omitted the server will use the default value of False  # noqa: E501
            real_time (bool): Real Time.. [optional] if omitted the server will use the default value of False  # noqa: E501
            credential (OptionTypeListCreateCredential): [optional]  # noqa: E501
            service_username (str, none_type): Username for authenticating with Basic Auth when type is rest or ldap username.. [optional]  # noqa: E501
            service_password (str, none_type): Password for authenticating with Basic Auth when type is rest or ldap password.. [optional]  # noqa: E501
            initial_dataset (str, none_type): Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties 'name', and 'value'. Required when type is manual.. [optional]  # noqa: E501
            translation_script (str, none_type): Translation Script. Create a js script to translate the result data object into an Array containing objects with properties 'name' and 'value'. The input data is provided as data and the result should be put on the global variable results.. [optional]  # noqa: E501
            request_script (str, none_type): Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties 'name' and 'value' for a get. The input data is provided as data and the result should be put on the global variable results.. [optional]  # noqa: E501
            config (OptionTypeListCreateConfig): [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', False)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        if args:
            for arg in args:
                if isinstance(arg, dict):
                    kwargs.update(arg)
                else:
                    raise ApiTypeError(
                        "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                            args,
                            self.__class__.__name__,
                        ),
                        path_to_item=_path_to_item,
                        valid_classes=(self.__class__,),
                    )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        self.name = name
        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
            if var_name in self.read_only_vars:
                raise ApiAttributeError(f"`{var_name}` is a read-only attribute. Use `from_openapi_data` to instantiate "
                                     f"class with read only attributes.")