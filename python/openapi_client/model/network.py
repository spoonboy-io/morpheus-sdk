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
    from openapi_client.model.get_all_network_floating_ips200_response_all_of_network_floating_ips_inner_network_domain import GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerNetworkDomain
    from openapi_client.model.list_deploys200_response_all_of_app_deploys_inner_instance import ListDeploys200ResponseAllOfAppDeploysInnerInstance
    from openapi_client.model.list_load_balancer_virtual_servers200_response_all_of_load_balancer_instances_inner_ssl_cert import ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert
    from openapi_client.model.network_config import NetworkConfig
    from openapi_client.model.network_network_proxy import NetworkNetworkProxy
    from openapi_client.model.network_owner import NetworkOwner
    from openapi_client.model.network_type import NetworkType
    from openapi_client.model.network_zone import NetworkZone
    globals()['GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerNetworkDomain'] = GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerNetworkDomain
    globals()['ListDeploys200ResponseAllOfAppDeploysInnerInstance'] = ListDeploys200ResponseAllOfAppDeploysInnerInstance
    globals()['ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert'] = ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert
    globals()['NetworkConfig'] = NetworkConfig
    globals()['NetworkNetworkProxy'] = NetworkNetworkProxy
    globals()['NetworkOwner'] = NetworkOwner
    globals()['NetworkType'] = NetworkType
    globals()['NetworkZone'] = NetworkZone


class Network(ModelNormal):
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
            'name': (str,),  # noqa: E501
            'display_name': (str,),  # noqa: E501
            'labels': ([str], none_type,),  # noqa: E501
            'zone': (NetworkZone,),  # noqa: E501
            'type': (NetworkType,),  # noqa: E501
            'owner': (NetworkOwner,),  # noqa: E501
            'code': (str, none_type,),  # noqa: E501
            'ipv4_enabled': (bool,),  # noqa: E501
            'ipv6_enabled': (bool,),  # noqa: E501
            'category': (str, none_type,),  # noqa: E501
            'interface_name': (str, none_type,),  # noqa: E501
            'bridge_name': (str, none_type,),  # noqa: E501
            'bridge_interface': (str, none_type,),  # noqa: E501
            'description': (str, none_type,),  # noqa: E501
            'external_id': (str, none_type,),  # noqa: E501
            'internal_id': (str, none_type,),  # noqa: E501
            'unique_id': (str, none_type,),  # noqa: E501
            'external_type': (str,),  # noqa: E501
            'ref_url': (str, none_type,),  # noqa: E501
            'ref_type': (str,),  # noqa: E501
            'ref_id': (int,),  # noqa: E501
            'vlan_id': (int, none_type,),  # noqa: E501
            'vswitch_name': (str, none_type,),  # noqa: E501
            'dhcp_server': (bool,),  # noqa: E501
            'dhcp_ip': (str, none_type,),  # noqa: E501
            'dhcp_server_ipv6': (bool,),  # noqa: E501
            'gateway': (str, none_type,),  # noqa: E501
            'netmask': (str, none_type,),  # noqa: E501
            'broadcast': (str, none_type,),  # noqa: E501
            'subnet_address': (str, none_type,),  # noqa: E501
            'dns_primary': (str, none_type,),  # noqa: E501
            'dns_secondary': (str, none_type,),  # noqa: E501
            'cidr': (str,),  # noqa: E501
            'gateway_ipv6': (str, none_type,),  # noqa: E501
            'netmask_ipv6': (str, none_type,),  # noqa: E501
            'dns_primary_ipv6': (str, none_type,),  # noqa: E501
            'dns_secondary_ipv6': (str, none_type,),  # noqa: E501
            'cidr_ipv6': (str, none_type,),  # noqa: E501
            'tftp_server': (str, none_type,),  # noqa: E501
            'boot_file': (str, none_type,),  # noqa: E501
            'switch_id': (str, none_type,),  # noqa: E501
            'fabric_id': (str, none_type,),  # noqa: E501
            'network_role': (str, none_type,),  # noqa: E501
            'status': (str,),  # noqa: E501
            'availability_zone': (str, none_type,),  # noqa: E501
            'pool': ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type,),  # noqa: E501
            'pool_ipv6': ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type,),  # noqa: E501
            'network_proxy': (NetworkNetworkProxy,),  # noqa: E501
            'network_domain': (GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerNetworkDomain,),  # noqa: E501
            'search_domains': (str, none_type,),  # noqa: E501
            'prefix_length': (str, none_type,),  # noqa: E501
            'visibility': (str,),  # noqa: E501
            'enable_admin': (bool,),  # noqa: E501
            'active': (bool,),  # noqa: E501
            'default_network': (bool,),  # noqa: E501
            'assign_public_ip': (bool,),  # noqa: E501
            'no_proxy': (str, none_type,),  # noqa: E501
            'appliance_url_proxy_bypass': (bool,),  # noqa: E501
            'zone_pool': (ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert,),  # noqa: E501
            'allow_static_override': (bool,),  # noqa: E501
            'config': (NetworkConfig,),  # noqa: E501
            'tenants': ([ListDeploys200ResponseAllOfAppDeploysInnerInstance],),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'id': 'id',  # noqa: E501
        'name': 'name',  # noqa: E501
        'display_name': 'displayName',  # noqa: E501
        'labels': 'labels',  # noqa: E501
        'zone': 'zone',  # noqa: E501
        'type': 'type',  # noqa: E501
        'owner': 'owner',  # noqa: E501
        'code': 'code',  # noqa: E501
        'ipv4_enabled': 'ipv4Enabled',  # noqa: E501
        'ipv6_enabled': 'ipv6Enabled',  # noqa: E501
        'category': 'category',  # noqa: E501
        'interface_name': 'interfaceName',  # noqa: E501
        'bridge_name': 'bridgeName',  # noqa: E501
        'bridge_interface': 'bridgeInterface',  # noqa: E501
        'description': 'description',  # noqa: E501
        'external_id': 'externalId',  # noqa: E501
        'internal_id': 'internalId',  # noqa: E501
        'unique_id': 'uniqueId',  # noqa: E501
        'external_type': 'externalType',  # noqa: E501
        'ref_url': 'refUrl',  # noqa: E501
        'ref_type': 'refType',  # noqa: E501
        'ref_id': 'refId',  # noqa: E501
        'vlan_id': 'vlanId',  # noqa: E501
        'vswitch_name': 'vswitchName',  # noqa: E501
        'dhcp_server': 'dhcpServer',  # noqa: E501
        'dhcp_ip': 'dhcpIp',  # noqa: E501
        'dhcp_server_ipv6': 'dhcpServerIPv6',  # noqa: E501
        'gateway': 'gateway',  # noqa: E501
        'netmask': 'netmask',  # noqa: E501
        'broadcast': 'broadcast',  # noqa: E501
        'subnet_address': 'subnetAddress',  # noqa: E501
        'dns_primary': 'dnsPrimary',  # noqa: E501
        'dns_secondary': 'dnsSecondary',  # noqa: E501
        'cidr': 'cidr',  # noqa: E501
        'gateway_ipv6': 'gatewayIPv6',  # noqa: E501
        'netmask_ipv6': 'netmaskIPv6',  # noqa: E501
        'dns_primary_ipv6': 'dnsPrimaryIPv6',  # noqa: E501
        'dns_secondary_ipv6': 'dnsSecondaryIPv6',  # noqa: E501
        'cidr_ipv6': 'cidrIPv6',  # noqa: E501
        'tftp_server': 'tftpServer',  # noqa: E501
        'boot_file': 'bootFile',  # noqa: E501
        'switch_id': 'switchId',  # noqa: E501
        'fabric_id': 'fabricId',  # noqa: E501
        'network_role': 'networkRole',  # noqa: E501
        'status': 'status',  # noqa: E501
        'availability_zone': 'availabilityZone',  # noqa: E501
        'pool': 'pool',  # noqa: E501
        'pool_ipv6': 'poolIPv6',  # noqa: E501
        'network_proxy': 'networkProxy',  # noqa: E501
        'network_domain': 'networkDomain',  # noqa: E501
        'search_domains': 'searchDomains',  # noqa: E501
        'prefix_length': 'prefixLength',  # noqa: E501
        'visibility': 'visibility',  # noqa: E501
        'enable_admin': 'enableAdmin',  # noqa: E501
        'active': 'active',  # noqa: E501
        'default_network': 'defaultNetwork',  # noqa: E501
        'assign_public_ip': 'assignPublicIp',  # noqa: E501
        'no_proxy': 'noProxy',  # noqa: E501
        'appliance_url_proxy_bypass': 'applianceUrlProxyBypass',  # noqa: E501
        'zone_pool': 'zonePool',  # noqa: E501
        'allow_static_override': 'allowStaticOverride',  # noqa: E501
        'config': 'config',  # noqa: E501
        'tenants': 'tenants',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, *args, **kwargs):  # noqa: E501
        """Network - a model defined in OpenAPI

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
            id (int): Network ID. [optional]  # noqa: E501
            name (str): Name. [optional]  # noqa: E501
            display_name (str): Network Display Name. [optional]  # noqa: E501
            labels ([str], none_type): [optional]  # noqa: E501
            zone (NetworkZone): [optional]  # noqa: E501
            type (NetworkType): [optional]  # noqa: E501
            owner (NetworkOwner): [optional]  # noqa: E501
            code (str, none_type): Network Code. [optional]  # noqa: E501
            ipv4_enabled (bool): [optional]  # noqa: E501
            ipv6_enabled (bool): [optional]  # noqa: E501
            category (str, none_type): Network Category. [optional]  # noqa: E501
            interface_name (str, none_type): [optional]  # noqa: E501
            bridge_name (str, none_type): [optional]  # noqa: E501
            bridge_interface (str, none_type): [optional]  # noqa: E501
            description (str, none_type): Description. [optional]  # noqa: E501
            external_id (str, none_type): [optional]  # noqa: E501
            internal_id (str, none_type): [optional]  # noqa: E501
            unique_id (str, none_type): [optional]  # noqa: E501
            external_type (str): [optional]  # noqa: E501
            ref_url (str, none_type): [optional]  # noqa: E501
            ref_type (str): [optional]  # noqa: E501
            ref_id (int): [optional]  # noqa: E501
            vlan_id (int, none_type): [optional]  # noqa: E501
            vswitch_name (str, none_type): [optional]  # noqa: E501
            dhcp_server (bool): [optional]  # noqa: E501
            dhcp_ip (str, none_type): [optional]  # noqa: E501
            dhcp_server_ipv6 (bool): [optional]  # noqa: E501
            gateway (str, none_type): Network Gateway. [optional]  # noqa: E501
            netmask (str, none_type): [optional]  # noqa: E501
            broadcast (str, none_type): [optional]  # noqa: E501
            subnet_address (str, none_type): [optional]  # noqa: E501
            dns_primary (str, none_type): Primary DNS Server. [optional]  # noqa: E501
            dns_secondary (str, none_type): Secondary DNS Server. [optional]  # noqa: E501
            cidr (str): Network CIDR. [optional]  # noqa: E501
            gateway_ipv6 (str, none_type): IPv6 Network Gateway. [optional]  # noqa: E501
            netmask_ipv6 (str, none_type): [optional]  # noqa: E501
            dns_primary_ipv6 (str, none_type): Primary IPv6 DNS Server. [optional]  # noqa: E501
            dns_secondary_ipv6 (str, none_type): Secondary IPv6 DNS Server. [optional]  # noqa: E501
            cidr_ipv6 (str, none_type): IPv6 Network CIDR. [optional]  # noqa: E501
            tftp_server (str, none_type): [optional]  # noqa: E501
            boot_file (str, none_type): [optional]  # noqa: E501
            switch_id (str, none_type): [optional]  # noqa: E501
            fabric_id (str, none_type): [optional]  # noqa: E501
            network_role (str, none_type): [optional]  # noqa: E501
            status (str): [optional]  # noqa: E501
            availability_zone (str, none_type): [optional]  # noqa: E501
            pool ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type): [optional]  # noqa: E501
            pool_ipv6 ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type): [optional]  # noqa: E501
            network_proxy (NetworkNetworkProxy): [optional]  # noqa: E501
            network_domain (GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerNetworkDomain): [optional]  # noqa: E501
            search_domains (str, none_type): [optional]  # noqa: E501
            prefix_length (str, none_type): [optional]  # noqa: E501
            visibility (str): [optional]  # noqa: E501
            enable_admin (bool): [optional]  # noqa: E501
            active (bool): [optional]  # noqa: E501
            default_network (bool): [optional]  # noqa: E501
            assign_public_ip (bool): [optional]  # noqa: E501
            no_proxy (str, none_type): [optional]  # noqa: E501
            appliance_url_proxy_bypass (bool): [optional]  # noqa: E501
            zone_pool (ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert): [optional]  # noqa: E501
            allow_static_override (bool): [optional]  # noqa: E501
            config (NetworkConfig): [optional]  # noqa: E501
            tenants ([ListDeploys200ResponseAllOfAppDeploysInnerInstance]): [optional]  # noqa: E501
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
        """Network - a model defined in OpenAPI

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
            id (int): Network ID. [optional]  # noqa: E501
            name (str): Name. [optional]  # noqa: E501
            display_name (str): Network Display Name. [optional]  # noqa: E501
            labels ([str], none_type): [optional]  # noqa: E501
            zone (NetworkZone): [optional]  # noqa: E501
            type (NetworkType): [optional]  # noqa: E501
            owner (NetworkOwner): [optional]  # noqa: E501
            code (str, none_type): Network Code. [optional]  # noqa: E501
            ipv4_enabled (bool): [optional]  # noqa: E501
            ipv6_enabled (bool): [optional]  # noqa: E501
            category (str, none_type): Network Category. [optional]  # noqa: E501
            interface_name (str, none_type): [optional]  # noqa: E501
            bridge_name (str, none_type): [optional]  # noqa: E501
            bridge_interface (str, none_type): [optional]  # noqa: E501
            description (str, none_type): Description. [optional]  # noqa: E501
            external_id (str, none_type): [optional]  # noqa: E501
            internal_id (str, none_type): [optional]  # noqa: E501
            unique_id (str, none_type): [optional]  # noqa: E501
            external_type (str): [optional]  # noqa: E501
            ref_url (str, none_type): [optional]  # noqa: E501
            ref_type (str): [optional]  # noqa: E501
            ref_id (int): [optional]  # noqa: E501
            vlan_id (int, none_type): [optional]  # noqa: E501
            vswitch_name (str, none_type): [optional]  # noqa: E501
            dhcp_server (bool): [optional]  # noqa: E501
            dhcp_ip (str, none_type): [optional]  # noqa: E501
            dhcp_server_ipv6 (bool): [optional]  # noqa: E501
            gateway (str, none_type): Network Gateway. [optional]  # noqa: E501
            netmask (str, none_type): [optional]  # noqa: E501
            broadcast (str, none_type): [optional]  # noqa: E501
            subnet_address (str, none_type): [optional]  # noqa: E501
            dns_primary (str, none_type): Primary DNS Server. [optional]  # noqa: E501
            dns_secondary (str, none_type): Secondary DNS Server. [optional]  # noqa: E501
            cidr (str): Network CIDR. [optional]  # noqa: E501
            gateway_ipv6 (str, none_type): IPv6 Network Gateway. [optional]  # noqa: E501
            netmask_ipv6 (str, none_type): [optional]  # noqa: E501
            dns_primary_ipv6 (str, none_type): Primary IPv6 DNS Server. [optional]  # noqa: E501
            dns_secondary_ipv6 (str, none_type): Secondary IPv6 DNS Server. [optional]  # noqa: E501
            cidr_ipv6 (str, none_type): IPv6 Network CIDR. [optional]  # noqa: E501
            tftp_server (str, none_type): [optional]  # noqa: E501
            boot_file (str, none_type): [optional]  # noqa: E501
            switch_id (str, none_type): [optional]  # noqa: E501
            fabric_id (str, none_type): [optional]  # noqa: E501
            network_role (str, none_type): [optional]  # noqa: E501
            status (str): [optional]  # noqa: E501
            availability_zone (str, none_type): [optional]  # noqa: E501
            pool ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type): [optional]  # noqa: E501
            pool_ipv6 ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type): [optional]  # noqa: E501
            network_proxy (NetworkNetworkProxy): [optional]  # noqa: E501
            network_domain (GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerNetworkDomain): [optional]  # noqa: E501
            search_domains (str, none_type): [optional]  # noqa: E501
            prefix_length (str, none_type): [optional]  # noqa: E501
            visibility (str): [optional]  # noqa: E501
            enable_admin (bool): [optional]  # noqa: E501
            active (bool): [optional]  # noqa: E501
            default_network (bool): [optional]  # noqa: E501
            assign_public_ip (bool): [optional]  # noqa: E501
            no_proxy (str, none_type): [optional]  # noqa: E501
            appliance_url_proxy_bypass (bool): [optional]  # noqa: E501
            zone_pool (ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert): [optional]  # noqa: E501
            allow_static_override (bool): [optional]  # noqa: E501
            config (NetworkConfig): [optional]  # noqa: E501
            tenants ([ListDeploys200ResponseAllOfAppDeploysInnerInstance]): [optional]  # noqa: E501
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