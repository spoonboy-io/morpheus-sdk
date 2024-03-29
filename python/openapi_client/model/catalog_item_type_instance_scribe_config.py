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
    from openapi_client.model.instances_config_aws import InstancesConfigAWS
    from openapi_client.model.instances_config_azure import InstancesConfigAzure
    from openapi_client.model.instances_config_gcp import InstancesConfigGCP
    from openapi_client.model.instances_config_vm_ware import InstancesConfigVMWare
    globals()['InstancesConfigAWS'] = InstancesConfigAWS
    globals()['InstancesConfigAzure'] = InstancesConfigAzure
    globals()['InstancesConfigGCP'] = InstancesConfigGCP
    globals()['InstancesConfigVMWare'] = InstancesConfigVMWare


class CatalogItemTypeInstanceScribeConfig(ModelComposed):
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
        ('availability_options',): {
            'ZONE': "zone",
            'SET': "set",
        },
        ('azurefloating_ip',): {
            'ON': "on",
            'OFF': "off",
        },
        ('boot_diagnostics',): {
            'ENABLE': "enable",
            'ENABLE_CUSTOM_STORAGE': "enable_custom_storage",
        },
        ('os_guest_diagnostics',): {
            'ON': "on",
            'OFF': "off",
        },
        ('nested_virtualization',): {
            'ON': "on",
            'OFF': "off",
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
            'resource_pool_id': (str,),  # noqa: E501
            'availability_options': (str,),  # noqa: E501
            'availability_set': (str,),  # noqa: E501
            'availability_zone': (int,),  # noqa: E501
            'azurefloating_ip': (str,),  # noqa: E501
            'boot_diagnostics': (str,),  # noqa: E501
            'os_guest_diagnostics': (str,),  # noqa: E501
            'diagnostics_storage_account': (str,),  # noqa: E501
            'create_user': (bool,),  # noqa: E501
            'no_agent': (bool, none_type,),  # noqa: E501
            'host_id': (str,),  # noqa: E501
            'smbios_asset_tag': (str,),  # noqa: E501
            'nested_virtualization': (str,),  # noqa: E501
            'vmware_folder_id': (str,),  # noqa: E501
            'google_zone_id': (int,),  # noqa: E501
            'external_ip_type': (int,),  # noqa: E501
            'network_tags': (str,),  # noqa: E501
            'service_account': (str,),  # noqa: E501
            'access_scope': (str,),  # noqa: E501
            'is_ec2': (str,),  # noqa: E501
            'availability_id': (str,),  # noqa: E501
            'security_id': (str,),  # noqa: E501
            'public_ip_type': (str,),  # noqa: E501
            'instance_profile': (str,),  # noqa: E501
            'kms_key_id': (str,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'resource_pool_id': 'resourcePoolId',  # noqa: E501
        'availability_options': 'availabilityOptions',  # noqa: E501
        'availability_set': 'availabilitySet',  # noqa: E501
        'availability_zone': 'availabilityZone',  # noqa: E501
        'azurefloating_ip': 'azurefloatingIp',  # noqa: E501
        'boot_diagnostics': 'bootDiagnostics',  # noqa: E501
        'os_guest_diagnostics': 'osGuestDiagnostics',  # noqa: E501
        'diagnostics_storage_account': 'diagnosticsStorageAccount',  # noqa: E501
        'create_user': 'createUser',  # noqa: E501
        'no_agent': 'noAgent',  # noqa: E501
        'host_id': 'hostId',  # noqa: E501
        'smbios_asset_tag': 'smbiosAssetTag',  # noqa: E501
        'nested_virtualization': 'nestedVirtualization',  # noqa: E501
        'vmware_folder_id': 'vmwareFolderId',  # noqa: E501
        'google_zone_id': 'googleZoneId',  # noqa: E501
        'external_ip_type': 'externalIpType',  # noqa: E501
        'network_tags': 'networkTags',  # noqa: E501
        'service_account': 'serviceAccount',  # noqa: E501
        'access_scope': 'accessScope',  # noqa: E501
        'is_ec2': 'isEC2',  # noqa: E501
        'availability_id': 'availabilityId',  # noqa: E501
        'security_id': 'securityId',  # noqa: E501
        'public_ip_type': 'publicIpType',  # noqa: E501
        'instance_profile': 'instanceProfile',  # noqa: E501
        'kms_key_id': 'kmsKeyId',  # noqa: E501
    }

    read_only_vars = {
    }

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, *args, **kwargs):  # noqa: E501
        """CatalogItemTypeInstanceScribeConfig - a model defined in OpenAPI

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
            resource_pool_id (str): id of the resource group to be used, can be prefixed with `pool-`. A resource pool group can be specified instead by prefixing its ID with `poolGroup-`.. [optional]  # noqa: E501
            availability_options (str): Availability Options. [optional]  # noqa: E501
            availability_set (str): Availability Set. [optional]  # noqa: E501
            availability_zone (int): Availability Zone. [optional]  # noqa: E501
            azurefloating_ip (str): Assign Public IP. [optional]  # noqa: E501
            boot_diagnostics (str): Boot Diagnostics. [optional]  # noqa: E501
            os_guest_diagnostics (str): OS Guest Diagnostics. [optional]  # noqa: E501
            diagnostics_storage_account (str): Diagnostics Storage Account. [optional]  # noqa: E501
            create_user (bool): Create User. [optional] if omitted the server will use the default value of True  # noqa: E501
            no_agent (bool, none_type): Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected.. [optional] if omitted the server will use the default value of False  # noqa: E501
            host_id (str): Specific host to deploy to if so desired.. [optional]  # noqa: E501
            smbios_asset_tag (str): Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used.. [optional]  # noqa: E501
            nested_virtualization (str): Enable Nested Virtualization. [optional] if omitted the server will use the default value of "off"  # noqa: E501
            vmware_folder_id (str): VMWare Folder External ID (as a String) or ID (as an Integer or String). [optional]  # noqa: E501
            google_zone_id (int): External ID of the Google zone to use for instance.. [optional]  # noqa: E501
            external_ip_type (int): External IP Type.  `-1` for ephemeral IP.. [optional]  # noqa: E501
            network_tags (str): Network Tags. [optional]  # noqa: E501
            service_account (str): Service Account. [optional]  # noqa: E501
            access_scope (str): Access Scope. [optional]  # noqa: E501
            is_ec2 (str): Amazon Cloud Type. [optional] if omitted the server will use the default value of "false"  # noqa: E501
            availability_id (str): Amazon Zone. [optional]  # noqa: E501
            security_id (str): Security Group. [optional]  # noqa: E501
            public_ip_type (str): Public IP. [optional]  # noqa: E501
            instance_profile (str): IAM Profile. [optional]  # noqa: E501
            kms_key_id (str): KMS Key ID. [optional]  # noqa: E501
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
        """CatalogItemTypeInstanceScribeConfig - a model defined in OpenAPI

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
            resource_pool_id (str): id of the resource group to be used, can be prefixed with `pool-`. A resource pool group can be specified instead by prefixing its ID with `poolGroup-`.. [optional]  # noqa: E501
            availability_options (str): Availability Options. [optional]  # noqa: E501
            availability_set (str): Availability Set. [optional]  # noqa: E501
            availability_zone (int): Availability Zone. [optional]  # noqa: E501
            azurefloating_ip (str): Assign Public IP. [optional]  # noqa: E501
            boot_diagnostics (str): Boot Diagnostics. [optional]  # noqa: E501
            os_guest_diagnostics (str): OS Guest Diagnostics. [optional]  # noqa: E501
            diagnostics_storage_account (str): Diagnostics Storage Account. [optional]  # noqa: E501
            create_user (bool): Create User. [optional] if omitted the server will use the default value of True  # noqa: E501
            no_agent (bool, none_type): Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected.. [optional] if omitted the server will use the default value of False  # noqa: E501
            host_id (str): Specific host to deploy to if so desired.. [optional]  # noqa: E501
            smbios_asset_tag (str): Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used.. [optional]  # noqa: E501
            nested_virtualization (str): Enable Nested Virtualization. [optional] if omitted the server will use the default value of "off"  # noqa: E501
            vmware_folder_id (str): VMWare Folder External ID (as a String) or ID (as an Integer or String). [optional]  # noqa: E501
            google_zone_id (int): External ID of the Google zone to use for instance.. [optional]  # noqa: E501
            external_ip_type (int): External IP Type.  `-1` for ephemeral IP.. [optional]  # noqa: E501
            network_tags (str): Network Tags. [optional]  # noqa: E501
            service_account (str): Service Account. [optional]  # noqa: E501
            access_scope (str): Access Scope. [optional]  # noqa: E501
            is_ec2 (str): Amazon Cloud Type. [optional] if omitted the server will use the default value of "false"  # noqa: E501
            availability_id (str): Amazon Zone. [optional]  # noqa: E501
            security_id (str): Security Group. [optional]  # noqa: E501
            public_ip_type (str): Public IP. [optional]  # noqa: E501
            instance_profile (str): IAM Profile. [optional]  # noqa: E501
            kms_key_id (str): KMS Key ID. [optional]  # noqa: E501
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
              InstancesConfigAWS,
              InstancesConfigAzure,
              InstancesConfigGCP,
              InstancesConfigVMWare,
              {str: (bool, date, datetime, dict, float, int, list, str, none_type)},
          ],
          'allOf': [
          ],
          'oneOf': [
          ],
        }
