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
    from openapi_client.model.add_vdi_pools_request_vdi_pool_one_of import AddVDIPoolsRequestVdiPoolOneOf
    from openapi_client.model.add_vdi_pools_request_vdi_pool_one_of1 import AddVDIPoolsRequestVdiPoolOneOf1
    from openapi_client.model.add_vdi_pools_request_vdi_pool_one_of_config import AddVDIPoolsRequestVdiPoolOneOfConfig
    globals()['AddVDIPoolsRequestVdiPoolOneOf'] = AddVDIPoolsRequestVdiPoolOneOf
    globals()['AddVDIPoolsRequestVdiPoolOneOf1'] = AddVDIPoolsRequestVdiPoolOneOf1
    globals()['AddVDIPoolsRequestVdiPoolOneOfConfig'] = AddVDIPoolsRequestVdiPoolOneOfConfig


class AddVDIPoolsRequestVdiPool(ModelComposed):
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
            'description': (str,),  # noqa: E501
            'owner': (int,),  # noqa: E501
            'min_idle': (float,),  # noqa: E501
            'initial_pool_size': (float,),  # noqa: E501
            'max_idle': (float,),  # noqa: E501
            'allocation_timeout_minutes': (float,),  # noqa: E501
            'persistent_user': (bool,),  # noqa: E501
            'recyclable': (bool,),  # noqa: E501
            'allow_copy': (bool,),  # noqa: E501
            'allow_printer': (bool,),  # noqa: E501
            'allow_fileshare': (bool,),  # noqa: E501
            'allow_hypervisor_console': (bool,),  # noqa: E501
            'auto_create_local_user_on_reservation': (bool,),  # noqa: E501
            'enabled': (bool,),  # noqa: E501
            'icon_path': (str,),  # noqa: E501
            'apps': ([int],),  # noqa: E501
            'gateway': (int,),  # noqa: E501
            'guest_console_jump_host': (str,),  # noqa: E501
            'guest_console_jump_port': (int,),  # noqa: E501
            'guest_console_jump_username': (str,),  # noqa: E501
            'guest_console_jump_password': (str,),  # noqa: E501
            'guest_console_jump_keypair': (int,),  # noqa: E501
            'name': (str,),  # noqa: E501
            'max_pool_size': (float,),  # noqa: E501
            'instance_config': (str,),  # noqa: E501
            'config': (AddVDIPoolsRequestVdiPoolOneOfConfig,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'description': 'description',  # noqa: E501
        'owner': 'owner',  # noqa: E501
        'min_idle': 'minIdle',  # noqa: E501
        'initial_pool_size': 'initialPoolSize',  # noqa: E501
        'max_idle': 'maxIdle',  # noqa: E501
        'allocation_timeout_minutes': 'allocationTimeoutMinutes',  # noqa: E501
        'persistent_user': 'persistentUser',  # noqa: E501
        'recyclable': 'recyclable',  # noqa: E501
        'allow_copy': 'allowCopy',  # noqa: E501
        'allow_printer': 'allowPrinter',  # noqa: E501
        'allow_fileshare': 'allowFileshare',  # noqa: E501
        'allow_hypervisor_console': 'allowHypervisorConsole',  # noqa: E501
        'auto_create_local_user_on_reservation': 'autoCreateLocalUserOnReservation',  # noqa: E501
        'enabled': 'enabled',  # noqa: E501
        'icon_path': 'iconPath',  # noqa: E501
        'apps': 'apps',  # noqa: E501
        'gateway': 'gateway',  # noqa: E501
        'guest_console_jump_host': 'guestConsoleJumpHost',  # noqa: E501
        'guest_console_jump_port': 'guestConsoleJumpPort',  # noqa: E501
        'guest_console_jump_username': 'guestConsoleJumpUsername',  # noqa: E501
        'guest_console_jump_password': 'guestConsoleJumpPassword',  # noqa: E501
        'guest_console_jump_keypair': 'guestConsoleJumpKeypair',  # noqa: E501
        'name': 'name',  # noqa: E501
        'max_pool_size': 'maxPoolSize',  # noqa: E501
        'instance_config': 'instanceConfig',  # noqa: E501
        'config': 'config',  # noqa: E501
    }

    read_only_vars = {
    }

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, *args, **kwargs):  # noqa: E501
        """AddVDIPoolsRequestVdiPool - a model defined in OpenAPI

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
            description (str): Virtual Desktop description. [optional]  # noqa: E501
            owner (int): Owner (User) ID. [optional]  # noqa: E501
            min_idle (float): Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby. . [optional]  # noqa: E501
            initial_pool_size (float): The initial size of the pool to be allocated on creation. [optional]  # noqa: E501
            max_idle (float): Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances . [optional]  # noqa: E501
            allocation_timeout_minutes (float): Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence. . [optional]  # noqa: E501
            persistent_user (bool): Persistent Desktop Pool. [optional] if omitted the server will use the default value of False  # noqa: E501
            recyclable (bool): Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD). [optional] if omitted the server will use the default value of False  # noqa: E501
            allow_copy (bool): Allow copy from desktop. [optional] if omitted the server will use the default value of False  # noqa: E501
            allow_printer (bool): Allow local printers from Desktop. [optional] if omitted the server will use the default value of False  # noqa: E501
            allow_fileshare (bool): Allow File Share. [optional] if omitted the server will use the default value of False  # noqa: E501
            allow_hypervisor_console (bool): Allow hypervisor console. [optional] if omitted the server will use the default value of False  # noqa: E501
            auto_create_local_user_on_reservation (bool): Auto create local user upon reservation. [optional] if omitted the server will use the default value of False  # noqa: E501
            enabled (bool): Can be used to enable or disable the VDI pool. [optional] if omitted the server will use the default value of True  # noqa: E501
            icon_path (str): The relative location of an icon image. [optional]  # noqa: E501
            apps ([int]): Array of VDI App IDs. [optional]  # noqa: E501
            gateway (int): VDI Gateway ID. [optional]  # noqa: E501
            guest_console_jump_host (str): Guest Console Jump Host. [optional]  # noqa: E501
            guest_console_jump_port (int): Guest Console Jump Port. [optional]  # noqa: E501
            guest_console_jump_username (str): Guest Console Jump Username. [optional]  # noqa: E501
            guest_console_jump_password (str): Guest Console Jump Password. [optional]  # noqa: E501
            guest_console_jump_keypair (int): Guest Console Jump Key Pair. see `Key Pair`. [optional]  # noqa: E501
            name (str): Virtual Desktop name. [optional]  # noqa: E501
            max_pool_size (float): Max limit on number of allocations and instances within the pool. . [optional]  # noqa: E501
            instance_config (str): Instance Config JSON. Passing as a string will preserve property order.  See `config` object for required values.. [optional]  # noqa: E501
            config (AddVDIPoolsRequestVdiPoolOneOfConfig): [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', False)
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

        constant_args = {
            '_check_type': _check_type,
            '_path_to_item': _path_to_item,
            '_spec_property_naming': _spec_property_naming,
            '_configuration': _configuration,
            '_visited_composed_classes': self._visited_composed_classes,
        }
        composed_info = validate_get_composed_info(
            constant_args, kwargs, self)
        self._composed_instances = composed_info[0]
        self._var_name_to_model_instances = composed_info[1]
        self._additional_properties_model_instances = composed_info[2]
        discarded_args = composed_info[3]

        for var_name, var_value in kwargs.items():
            if var_name in discarded_args and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self._additional_properties_model_instances:
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
        '_composed_instances',
        '_var_name_to_model_instances',
        '_additional_properties_model_instances',
    ])

    @convert_js_args_to_python_args
    def __init__(self, *args, **kwargs):  # noqa: E501
        """AddVDIPoolsRequestVdiPool - a model defined in OpenAPI

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
            description (str): Virtual Desktop description. [optional]  # noqa: E501
            owner (int): Owner (User) ID. [optional]  # noqa: E501
            min_idle (float): Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby. . [optional]  # noqa: E501
            initial_pool_size (float): The initial size of the pool to be allocated on creation. [optional]  # noqa: E501
            max_idle (float): Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances . [optional]  # noqa: E501
            allocation_timeout_minutes (float): Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence. . [optional]  # noqa: E501
            persistent_user (bool): Persistent Desktop Pool. [optional] if omitted the server will use the default value of False  # noqa: E501
            recyclable (bool): Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD). [optional] if omitted the server will use the default value of False  # noqa: E501
            allow_copy (bool): Allow copy from desktop. [optional] if omitted the server will use the default value of False  # noqa: E501
            allow_printer (bool): Allow local printers from Desktop. [optional] if omitted the server will use the default value of False  # noqa: E501
            allow_fileshare (bool): Allow File Share. [optional] if omitted the server will use the default value of False  # noqa: E501
            allow_hypervisor_console (bool): Allow hypervisor console. [optional] if omitted the server will use the default value of False  # noqa: E501
            auto_create_local_user_on_reservation (bool): Auto create local user upon reservation. [optional] if omitted the server will use the default value of False  # noqa: E501
            enabled (bool): Can be used to enable or disable the VDI pool. [optional] if omitted the server will use the default value of True  # noqa: E501
            icon_path (str): The relative location of an icon image. [optional]  # noqa: E501
            apps ([int]): Array of VDI App IDs. [optional]  # noqa: E501
            gateway (int): VDI Gateway ID. [optional]  # noqa: E501
            guest_console_jump_host (str): Guest Console Jump Host. [optional]  # noqa: E501
            guest_console_jump_port (int): Guest Console Jump Port. [optional]  # noqa: E501
            guest_console_jump_username (str): Guest Console Jump Username. [optional]  # noqa: E501
            guest_console_jump_password (str): Guest Console Jump Password. [optional]  # noqa: E501
            guest_console_jump_keypair (int): Guest Console Jump Key Pair. see `Key Pair`. [optional]  # noqa: E501
            name (str): Virtual Desktop name. [optional]  # noqa: E501
            max_pool_size (float): Max limit on number of allocations and instances within the pool. . [optional]  # noqa: E501
            instance_config (str): Instance Config JSON. Passing as a string will preserve property order.  See `config` object for required values.. [optional]  # noqa: E501
            config (AddVDIPoolsRequestVdiPoolOneOfConfig): [optional]  # noqa: E501
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

        constant_args = {
            '_check_type': _check_type,
            '_path_to_item': _path_to_item,
            '_spec_property_naming': _spec_property_naming,
            '_configuration': _configuration,
            '_visited_composed_classes': self._visited_composed_classes,
        }
        composed_info = validate_get_composed_info(
            constant_args, kwargs, self)
        self._composed_instances = composed_info[0]
        self._var_name_to_model_instances = composed_info[1]
        self._additional_properties_model_instances = composed_info[2]
        discarded_args = composed_info[3]

        for var_name, var_value in kwargs.items():
            if var_name in discarded_args and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self._additional_properties_model_instances:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
            if var_name in self.read_only_vars:
                raise ApiAttributeError(f"`{var_name}` is a read-only attribute. Use `from_openapi_data` to instantiate "
                                     f"class with read only attributes.")

    @cached_property
    def _composed_schemas():
        # we need this here to make our import statements work
        # we must store _composed_schemas in here so the code is only run
        # when we invoke this method. If we kept this at the class
        # level we would get an error because the class level
        # code would be run when this module is imported, and these composed
        # classes don't exist yet because their module has not finished
        # loading
        lazy_import()
        return {
          'anyOf': [
          ],
          'allOf': [
          ],
          'oneOf': [
              AddVDIPoolsRequestVdiPoolOneOf,
              AddVDIPoolsRequestVdiPoolOneOf1,
          ],
        }
