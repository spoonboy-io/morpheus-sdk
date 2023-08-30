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
    from openapi_client.model.cluster_containers_available_actions_inner import ClusterContainersAvailableActionsInner
    from openapi_client.model.cluster_containers_container_type import ClusterContainersContainerType
    from openapi_client.model.cluster_containers_container_type_set import ClusterContainersContainerTypeSet
    from openapi_client.model.cluster_containers_plan import ClusterContainersPlan
    from openapi_client.model.cluster_containers_stats import ClusterContainersStats
    from openapi_client.model.list_deploys200_response_all_of_app_deploys_inner_instance import ListDeploys200ResponseAllOfAppDeploysInnerInstance
    globals()['ClusterContainersAvailableActionsInner'] = ClusterContainersAvailableActionsInner
    globals()['ClusterContainersContainerType'] = ClusterContainersContainerType
    globals()['ClusterContainersContainerTypeSet'] = ClusterContainersContainerTypeSet
    globals()['ClusterContainersPlan'] = ClusterContainersPlan
    globals()['ClusterContainersStats'] = ClusterContainersStats
    globals()['ListDeploys200ResponseAllOfAppDeploysInnerInstance'] = ListDeploys200ResponseAllOfAppDeploysInnerInstance


class ClusterContainers(ModelNormal):
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
            'uuid': (str,),  # noqa: E501
            'account_id': (int,),  # noqa: E501
            'instance': (str, none_type,),  # noqa: E501
            'container_type': (ClusterContainersContainerType,),  # noqa: E501
            'container_type_set': (ClusterContainersContainerTypeSet,),  # noqa: E501
            'server': (ListDeploys200ResponseAllOfAppDeploysInnerInstance,),  # noqa: E501
            'cloud': (ListDeploys200ResponseAllOfAppDeploysInnerInstance,),  # noqa: E501
            'name': (str,),  # noqa: E501
            'ip': (str,),  # noqa: E501
            'internal_ip': (str,),  # noqa: E501
            'internal_hostname': (str,),  # noqa: E501
            'external_hostname': (str,),  # noqa: E501
            'external_domain': (str,),  # noqa: E501
            'external_fqdn': (str,),  # noqa: E501
            'ports': ([{str: (bool, date, datetime, dict, float, int, list, str, none_type)}],),  # noqa: E501
            'plan': (ClusterContainersPlan,),  # noqa: E501
            'date_created': (datetime, none_type,),  # noqa: E501
            'last_updated': (datetime,),  # noqa: E501
            'stats_enabled': (bool,),  # noqa: E501
            'status': (str,),  # noqa: E501
            'user_status': (str,),  # noqa: E501
            'environment_prefix': (str, none_type,),  # noqa: E501
            'config_group': (str, none_type,),  # noqa: E501
            'config_id': (str, none_type,),  # noqa: E501
            'config_role': (str, none_type,),  # noqa: E501
            'stats': (ClusterContainersStats,),  # noqa: E501
            'runtime_info': ({str: (bool, date, datetime, dict, float, int, list, str, none_type)},),  # noqa: E501
            'container_version': (str, none_type,),  # noqa: E501
            'repository_image': (str, none_type,),  # noqa: E501
            'plan_category': (str, none_type,),  # noqa: E501
            'hostname': (str, none_type,),  # noqa: E501
            'domain_name': (str, none_type,),  # noqa: E501
            'volume_created': (bool,),  # noqa: E501
            'container_created': (bool,),  # noqa: E501
            'max_storage': (str, none_type,),  # noqa: E501
            'max_memory': (str, none_type,),  # noqa: E501
            'max_cores': (str, none_type,),  # noqa: E501
            'max_cpu': (str, none_type,),  # noqa: E501
            'hourly_price': (float,),  # noqa: E501
            'available_actions': ([ClusterContainersAvailableActionsInner],),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'id': 'id',  # noqa: E501
        'uuid': 'uuid',  # noqa: E501
        'account_id': 'accountId',  # noqa: E501
        'instance': 'instance',  # noqa: E501
        'container_type': 'containerType',  # noqa: E501
        'container_type_set': 'containerTypeSet',  # noqa: E501
        'server': 'server',  # noqa: E501
        'cloud': 'cloud',  # noqa: E501
        'name': 'name',  # noqa: E501
        'ip': 'ip',  # noqa: E501
        'internal_ip': 'internalIp',  # noqa: E501
        'internal_hostname': 'internalHostname',  # noqa: E501
        'external_hostname': 'externalHostname',  # noqa: E501
        'external_domain': 'externalDomain',  # noqa: E501
        'external_fqdn': 'externalFqdn',  # noqa: E501
        'ports': 'ports',  # noqa: E501
        'plan': 'plan',  # noqa: E501
        'date_created': 'dateCreated',  # noqa: E501
        'last_updated': 'lastUpdated',  # noqa: E501
        'stats_enabled': 'statsEnabled',  # noqa: E501
        'status': 'status',  # noqa: E501
        'user_status': 'userStatus',  # noqa: E501
        'environment_prefix': 'environmentPrefix',  # noqa: E501
        'config_group': 'configGroup',  # noqa: E501
        'config_id': 'configId',  # noqa: E501
        'config_role': 'configRole',  # noqa: E501
        'stats': 'stats',  # noqa: E501
        'runtime_info': 'runtimeInfo',  # noqa: E501
        'container_version': 'containerVersion',  # noqa: E501
        'repository_image': 'repositoryImage',  # noqa: E501
        'plan_category': 'planCategory',  # noqa: E501
        'hostname': 'hostname',  # noqa: E501
        'domain_name': 'domainName',  # noqa: E501
        'volume_created': 'volumeCreated',  # noqa: E501
        'container_created': 'containerCreated',  # noqa: E501
        'max_storage': 'maxStorage',  # noqa: E501
        'max_memory': 'maxMemory',  # noqa: E501
        'max_cores': 'maxCores',  # noqa: E501
        'max_cpu': 'maxCpu',  # noqa: E501
        'hourly_price': 'hourlyPrice',  # noqa: E501
        'available_actions': 'availableActions',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, *args, **kwargs):  # noqa: E501
        """ClusterContainers - a model defined in OpenAPI

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
            uuid (str): [optional]  # noqa: E501
            account_id (int): [optional]  # noqa: E501
            instance (str, none_type): [optional]  # noqa: E501
            container_type (ClusterContainersContainerType): [optional]  # noqa: E501
            container_type_set (ClusterContainersContainerTypeSet): [optional]  # noqa: E501
            server (ListDeploys200ResponseAllOfAppDeploysInnerInstance): [optional]  # noqa: E501
            cloud (ListDeploys200ResponseAllOfAppDeploysInnerInstance): [optional]  # noqa: E501
            name (str): [optional]  # noqa: E501
            ip (str): [optional]  # noqa: E501
            internal_ip (str): [optional]  # noqa: E501
            internal_hostname (str): [optional]  # noqa: E501
            external_hostname (str): [optional]  # noqa: E501
            external_domain (str): [optional]  # noqa: E501
            external_fqdn (str): [optional]  # noqa: E501
            ports ([{str: (bool, date, datetime, dict, float, int, list, str, none_type)}]): [optional]  # noqa: E501
            plan (ClusterContainersPlan): [optional]  # noqa: E501
            date_created (datetime, none_type): [optional]  # noqa: E501
            last_updated (datetime): [optional]  # noqa: E501
            stats_enabled (bool): [optional]  # noqa: E501
            status (str): [optional]  # noqa: E501
            user_status (str): [optional]  # noqa: E501
            environment_prefix (str, none_type): [optional]  # noqa: E501
            config_group (str, none_type): [optional]  # noqa: E501
            config_id (str, none_type): [optional]  # noqa: E501
            config_role (str, none_type): [optional]  # noqa: E501
            stats (ClusterContainersStats): [optional]  # noqa: E501
            runtime_info ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}): [optional]  # noqa: E501
            container_version (str, none_type): [optional]  # noqa: E501
            repository_image (str, none_type): [optional]  # noqa: E501
            plan_category (str, none_type): [optional]  # noqa: E501
            hostname (str, none_type): [optional]  # noqa: E501
            domain_name (str, none_type): [optional]  # noqa: E501
            volume_created (bool): [optional]  # noqa: E501
            container_created (bool): [optional]  # noqa: E501
            max_storage (str, none_type): [optional]  # noqa: E501
            max_memory (str, none_type): [optional]  # noqa: E501
            max_cores (str, none_type): [optional]  # noqa: E501
            max_cpu (str, none_type): [optional]  # noqa: E501
            hourly_price (float): [optional]  # noqa: E501
            available_actions ([ClusterContainersAvailableActionsInner]): [optional]  # noqa: E501
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
        """ClusterContainers - a model defined in OpenAPI

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
            uuid (str): [optional]  # noqa: E501
            account_id (int): [optional]  # noqa: E501
            instance (str, none_type): [optional]  # noqa: E501
            container_type (ClusterContainersContainerType): [optional]  # noqa: E501
            container_type_set (ClusterContainersContainerTypeSet): [optional]  # noqa: E501
            server (ListDeploys200ResponseAllOfAppDeploysInnerInstance): [optional]  # noqa: E501
            cloud (ListDeploys200ResponseAllOfAppDeploysInnerInstance): [optional]  # noqa: E501
            name (str): [optional]  # noqa: E501
            ip (str): [optional]  # noqa: E501
            internal_ip (str): [optional]  # noqa: E501
            internal_hostname (str): [optional]  # noqa: E501
            external_hostname (str): [optional]  # noqa: E501
            external_domain (str): [optional]  # noqa: E501
            external_fqdn (str): [optional]  # noqa: E501
            ports ([{str: (bool, date, datetime, dict, float, int, list, str, none_type)}]): [optional]  # noqa: E501
            plan (ClusterContainersPlan): [optional]  # noqa: E501
            date_created (datetime, none_type): [optional]  # noqa: E501
            last_updated (datetime): [optional]  # noqa: E501
            stats_enabled (bool): [optional]  # noqa: E501
            status (str): [optional]  # noqa: E501
            user_status (str): [optional]  # noqa: E501
            environment_prefix (str, none_type): [optional]  # noqa: E501
            config_group (str, none_type): [optional]  # noqa: E501
            config_id (str, none_type): [optional]  # noqa: E501
            config_role (str, none_type): [optional]  # noqa: E501
            stats (ClusterContainersStats): [optional]  # noqa: E501
            runtime_info ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}): [optional]  # noqa: E501
            container_version (str, none_type): [optional]  # noqa: E501
            repository_image (str, none_type): [optional]  # noqa: E501
            plan_category (str, none_type): [optional]  # noqa: E501
            hostname (str, none_type): [optional]  # noqa: E501
            domain_name (str, none_type): [optional]  # noqa: E501
            volume_created (bool): [optional]  # noqa: E501
            container_created (bool): [optional]  # noqa: E501
            max_storage (str, none_type): [optional]  # noqa: E501
            max_memory (str, none_type): [optional]  # noqa: E501
            max_cores (str, none_type): [optional]  # noqa: E501
            max_cpu (str, none_type): [optional]  # noqa: E501
            hourly_price (float): [optional]  # noqa: E501
            available_actions ([ClusterContainersAvailableActionsInner]): [optional]  # noqa: E501
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
