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
    from openapi_client.model.update_blueprint_permissions_request_resource_permission_sites_inner import UpdateBlueprintPermissionsRequestResourcePermissionSitesInner
    globals()['UpdateBlueprintPermissionsRequestResourcePermissionSitesInner'] = UpdateBlueprintPermissionsRequestResourcePermissionSitesInner


class CatalogItemTypeWorkflowCreate(ModelNormal):
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
            'WORKFLOW': "workflow",
        },
        ('context',): {
            'INSTANCE': "instance",
            'SERVER': "server",
            'APPLIANCE': "appliance",
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
            'workflow': (UpdateBlueprintPermissionsRequestResourcePermissionSitesInner,),  # noqa: E501
            'name': (str,),  # noqa: E501
            'code': (str, none_type,),  # noqa: E501
            'category': (str, none_type,),  # noqa: E501
            'description': (str,),  # noqa: E501
            'labels': ([str], none_type,),  # noqa: E501
            'type': (str,),  # noqa: E501
            'visibility': (str,),  # noqa: E501
            'layout_code': (str, none_type,),  # noqa: E501
            'icon_path': (str,),  # noqa: E501
            'enabled': (bool,),  # noqa: E501
            'featured': (bool,),  # noqa: E501
            'allow_quantity': (bool,),  # noqa: E501
            'context': (str,),  # noqa: E501
            'workflow_config': (str,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'workflow': 'workflow',  # noqa: E501
        'name': 'name',  # noqa: E501
        'code': 'code',  # noqa: E501
        'category': 'category',  # noqa: E501
        'description': 'description',  # noqa: E501
        'labels': 'labels',  # noqa: E501
        'type': 'type',  # noqa: E501
        'visibility': 'visibility',  # noqa: E501
        'layout_code': 'layoutCode',  # noqa: E501
        'icon_path': 'iconPath',  # noqa: E501
        'enabled': 'enabled',  # noqa: E501
        'featured': 'featured',  # noqa: E501
        'allow_quantity': 'allowQuantity',  # noqa: E501
        'context': 'context',  # noqa: E501
        'workflow_config': 'workflowConfig',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, workflow, *args, **kwargs):  # noqa: E501
        """CatalogItemTypeWorkflowCreate - a model defined in OpenAPI

        Args:
            workflow (UpdateBlueprintPermissionsRequestResourcePermissionSitesInner):

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
            name (str): Catalog Item Type name. [optional]  # noqa: E501
            code (str, none_type): Useful shortcode for provisioning naming schemes and export reference.. [optional]  # noqa: E501
            category (str, none_type): Catalog Item Type category. [optional]  # noqa: E501
            description (str): Catalog Item Type description. [optional]  # noqa: E501
            labels ([str], none_type): Array of label strings, can be used for filtering.. [optional]  # noqa: E501
            type (str): Type, `instance`, `blueprint` or `workflow`. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context.. [optional] if omitted the server will use the default value of "workflow"  # noqa: E501
            visibility (str): Visibility - Set to public to allow all tenants. [optional] if omitted the server will use the default value of "private"  # noqa: E501
            layout_code (str, none_type): Identifier primarily used for Plugin Catalog Item Types. [optional]  # noqa: E501
            icon_path (str): Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png.. [optional]  # noqa: E501
            enabled (bool): Can be used to enable / disable the catalog item type.. [optional] if omitted the server will use the default value of True  # noqa: E501
            featured (bool): Can be used to feature the catalog item type.. [optional] if omitted the server will use the default value of False  # noqa: E501
            allow_quantity (bool): Can users order more than one of this item at a time.. [optional] if omitted the server will use the default value of False  # noqa: E501
            context (str): Context for running the workflow, determines if a target resource must be selected.. [optional]  # noqa: E501
            workflow_config (str): Configuration object that contains settings for the workflow.. [optional]  # noqa: E501
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

        self.workflow = workflow
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
    def __init__(self, workflow, *args, **kwargs):  # noqa: E501
        """CatalogItemTypeWorkflowCreate - a model defined in OpenAPI

        Args:
            workflow (UpdateBlueprintPermissionsRequestResourcePermissionSitesInner):

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
            name (str): Catalog Item Type name. [optional]  # noqa: E501
            code (str, none_type): Useful shortcode for provisioning naming schemes and export reference.. [optional]  # noqa: E501
            category (str, none_type): Catalog Item Type category. [optional]  # noqa: E501
            description (str): Catalog Item Type description. [optional]  # noqa: E501
            labels ([str], none_type): Array of label strings, can be used for filtering.. [optional]  # noqa: E501
            type (str): Type, `instance`, `blueprint` or `workflow`. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context.. [optional] if omitted the server will use the default value of "workflow"  # noqa: E501
            visibility (str): Visibility - Set to public to allow all tenants. [optional] if omitted the server will use the default value of "private"  # noqa: E501
            layout_code (str, none_type): Identifier primarily used for Plugin Catalog Item Types. [optional]  # noqa: E501
            icon_path (str): Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png.. [optional]  # noqa: E501
            enabled (bool): Can be used to enable / disable the catalog item type.. [optional] if omitted the server will use the default value of True  # noqa: E501
            featured (bool): Can be used to feature the catalog item type.. [optional] if omitted the server will use the default value of False  # noqa: E501
            allow_quantity (bool): Can users order more than one of this item at a time.. [optional] if omitted the server will use the default value of False  # noqa: E501
            context (str): Context for running the workflow, determines if a target resource must be selected.. [optional]  # noqa: E501
            workflow_config (str): Configuration object that contains settings for the workflow.. [optional]  # noqa: E501
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

        self.workflow = workflow
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
