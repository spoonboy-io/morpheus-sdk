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
    from openapi_client.model.list_deploys200_response_all_of_app_deploys_inner_instance import ListDeploys200ResponseAllOfAppDeploysInnerInstance
    from openapi_client.model.list_load_balancer_monitors200_response_all_of_load_balancer_monitors_inner_load_balancer_type import ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType
    globals()['ListDeploys200ResponseAllOfAppDeploysInnerInstance'] = ListDeploys200ResponseAllOfAppDeploysInnerInstance
    globals()['ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType'] = ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType


class GuidanceVmwareSizingResourceInterfacesInner(ModelNormal):
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
            'id': (int,),  # noqa: E501
            'ref_type': (str, none_type,),  # noqa: E501
            'ref_id': (str, none_type,),  # noqa: E501
            'name': (str,),  # noqa: E501
            'internal_id': (str,),  # noqa: E501
            'external_id': (str,),  # noqa: E501
            'unique_id': (str, none_type,),  # noqa: E501
            'public_ip_address': (str,),  # noqa: E501
            'public_ipv6_address': (str, none_type,),  # noqa: E501
            'ip_address': (str,),  # noqa: E501
            'ipv6_address': (str,),  # noqa: E501
            'ip_subnet': (str, none_type,),  # noqa: E501
            'ipv6_subnet': (str, none_type,),  # noqa: E501
            'description': (str, none_type,),  # noqa: E501
            'dhcp': (bool,),  # noqa: E501
            'active': (bool,),  # noqa: E501
            'pool_assigned': (bool,),  # noqa: E501
            'primary_interface': (bool,),  # noqa: E501
            'network': (ListDeploys200ResponseAllOfAppDeploysInnerInstance,),  # noqa: E501
            'subnet': (str, none_type,),  # noqa: E501
            'network_group': (str, none_type,),  # noqa: E501
            'network_position': (str, none_type,),  # noqa: E501
            'network_pool': (ListDeploys200ResponseAllOfAppDeploysInnerInstance,),  # noqa: E501
            'network_domain': (str, none_type,),  # noqa: E501
            'type': (ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType,),  # noqa: E501
            'ip_mode': (str,),  # noqa: E501
            'mac_address': (str,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'id': 'id',  # noqa: E501
        'ref_type': 'refType',  # noqa: E501
        'ref_id': 'refId',  # noqa: E501
        'name': 'name',  # noqa: E501
        'internal_id': 'internalId',  # noqa: E501
        'external_id': 'externalId',  # noqa: E501
        'unique_id': 'uniqueId',  # noqa: E501
        'public_ip_address': 'publicIpAddress',  # noqa: E501
        'public_ipv6_address': 'publicIpv6Address',  # noqa: E501
        'ip_address': 'ipAddress',  # noqa: E501
        'ipv6_address': 'ipv6Address',  # noqa: E501
        'ip_subnet': 'ipSubnet',  # noqa: E501
        'ipv6_subnet': 'ipv6Subnet',  # noqa: E501
        'description': 'description',  # noqa: E501
        'dhcp': 'dhcp',  # noqa: E501
        'active': 'active',  # noqa: E501
        'pool_assigned': 'poolAssigned',  # noqa: E501
        'primary_interface': 'primaryInterface',  # noqa: E501
        'network': 'network',  # noqa: E501
        'subnet': 'subnet',  # noqa: E501
        'network_group': 'networkGroup',  # noqa: E501
        'network_position': 'networkPosition',  # noqa: E501
        'network_pool': 'networkPool',  # noqa: E501
        'network_domain': 'networkDomain',  # noqa: E501
        'type': 'type',  # noqa: E501
        'ip_mode': 'ipMode',  # noqa: E501
        'mac_address': 'macAddress',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, *args, **kwargs):  # noqa: E501
        """GuidanceVmwareSizingResourceInterfacesInner - a model defined in OpenAPI

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
            id (int): [optional]  # noqa: E501
            ref_type (str, none_type): [optional]  # noqa: E501
            ref_id (str, none_type): [optional]  # noqa: E501
            name (str): [optional]  # noqa: E501
            internal_id (str): [optional]  # noqa: E501
            external_id (str): [optional]  # noqa: E501
            unique_id (str, none_type): [optional]  # noqa: E501
            public_ip_address (str): [optional]  # noqa: E501
            public_ipv6_address (str, none_type): [optional]  # noqa: E501
            ip_address (str): [optional]  # noqa: E501
            ipv6_address (str): [optional]  # noqa: E501
            ip_subnet (str, none_type): [optional]  # noqa: E501
            ipv6_subnet (str, none_type): [optional]  # noqa: E501
            description (str, none_type): [optional]  # noqa: E501
            dhcp (bool): [optional]  # noqa: E501
            active (bool): [optional]  # noqa: E501
            pool_assigned (bool): [optional]  # noqa: E501
            primary_interface (bool): [optional]  # noqa: E501
            network (ListDeploys200ResponseAllOfAppDeploysInnerInstance): [optional]  # noqa: E501
            subnet (str, none_type): [optional]  # noqa: E501
            network_group (str, none_type): [optional]  # noqa: E501
            network_position (str, none_type): [optional]  # noqa: E501
            network_pool (ListDeploys200ResponseAllOfAppDeploysInnerInstance): [optional]  # noqa: E501
            network_domain (str, none_type): [optional]  # noqa: E501
            type (ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType): [optional]  # noqa: E501
            ip_mode (str): [optional]  # noqa: E501
            mac_address (str): [optional]  # noqa: E501
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
    def __init__(self, *args, **kwargs):  # noqa: E501
        """GuidanceVmwareSizingResourceInterfacesInner - a model defined in OpenAPI

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
            id (int): [optional]  # noqa: E501
            ref_type (str, none_type): [optional]  # noqa: E501
            ref_id (str, none_type): [optional]  # noqa: E501
            name (str): [optional]  # noqa: E501
            internal_id (str): [optional]  # noqa: E501
            external_id (str): [optional]  # noqa: E501
            unique_id (str, none_type): [optional]  # noqa: E501
            public_ip_address (str): [optional]  # noqa: E501
            public_ipv6_address (str, none_type): [optional]  # noqa: E501
            ip_address (str): [optional]  # noqa: E501
            ipv6_address (str): [optional]  # noqa: E501
            ip_subnet (str, none_type): [optional]  # noqa: E501
            ipv6_subnet (str, none_type): [optional]  # noqa: E501
            description (str, none_type): [optional]  # noqa: E501
            dhcp (bool): [optional]  # noqa: E501
            active (bool): [optional]  # noqa: E501
            pool_assigned (bool): [optional]  # noqa: E501
            primary_interface (bool): [optional]  # noqa: E501
            network (ListDeploys200ResponseAllOfAppDeploysInnerInstance): [optional]  # noqa: E501
            subnet (str, none_type): [optional]  # noqa: E501
            network_group (str, none_type): [optional]  # noqa: E501
            network_position (str, none_type): [optional]  # noqa: E501
            network_pool (ListDeploys200ResponseAllOfAppDeploysInnerInstance): [optional]  # noqa: E501
            network_domain (str, none_type): [optional]  # noqa: E501
            type (ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType): [optional]  # noqa: E501
            ip_mode (str): [optional]  # noqa: E501
            mac_address (str): [optional]  # noqa: E501
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
