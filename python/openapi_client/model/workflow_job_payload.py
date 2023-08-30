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
    from openapi_client.model.workflow_job_payload_schedule_mode import WorkflowJobPayloadScheduleMode
    from openapi_client.model.workflow_job_payload_targets_inner import WorkflowJobPayloadTargetsInner
    from openapi_client.model.workflow_job_payload_task import WorkflowJobPayloadTask
    globals()['WorkflowJobPayloadScheduleMode'] = WorkflowJobPayloadScheduleMode
    globals()['WorkflowJobPayloadTargetsInner'] = WorkflowJobPayloadTargetsInner
    globals()['WorkflowJobPayloadTask'] = WorkflowJobPayloadTask


class WorkflowJobPayload(ModelNormal):
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
        ('target_type',): {
            'APPLIANCE': "appliance",
            'INSTANCE': "instance",
            'INSTANCE-LABEL': "instance-label",
            'SERVER': "server",
            'SERVER-LABEL': "server-label",
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
            'workflow': (WorkflowJobPayloadTask,),  # noqa: E501
            'target_type': (str,),  # noqa: E501
            'schedule_mode': (WorkflowJobPayloadScheduleMode,),  # noqa: E501
            'labels': ([str], none_type,),  # noqa: E501
            'enabled': (bool,),  # noqa: E501
            'task': (WorkflowJobPayloadTask,),  # noqa: E501
            'targets': ([WorkflowJobPayloadTargetsInner], none_type,),  # noqa: E501
            'instance_label': (str,),  # noqa: E501
            'server_label': (str,),  # noqa: E501
            'custom_options': ({str: (bool, date, datetime, dict, float, int, list, str, none_type)},),  # noqa: E501
            'custom_config': (str,),  # noqa: E501
            'date_time': (datetime,),  # noqa: E501
            'run': (bool,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'name': 'name',  # noqa: E501
        'workflow': 'workflow',  # noqa: E501
        'target_type': 'targetType',  # noqa: E501
        'schedule_mode': 'scheduleMode',  # noqa: E501
        'labels': 'labels',  # noqa: E501
        'enabled': 'enabled',  # noqa: E501
        'task': 'task',  # noqa: E501
        'targets': 'targets',  # noqa: E501
        'instance_label': 'instanceLabel',  # noqa: E501
        'server_label': 'serverLabel',  # noqa: E501
        'custom_options': 'customOptions',  # noqa: E501
        'custom_config': 'customConfig',  # noqa: E501
        'date_time': 'dateTime',  # noqa: E501
        'run': 'run',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, name, workflow, target_type, schedule_mode, *args, **kwargs):  # noqa: E501
        """WorkflowJobPayload - a model defined in OpenAPI

        Args:
            name (str): A name for the Job
            workflow (WorkflowJobPayloadTask):
            target_type (str): Target type where job will execute
            schedule_mode (WorkflowJobPayloadScheduleMode):

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
            labels ([str], none_type): Array of label strings, can be used for filtering.. [optional]  # noqa: E501
            enabled (bool): Use this to set enabled state. [optional] if omitted the server will use the default value of True  # noqa: E501
            task (WorkflowJobPayloadTask): [optional]  # noqa: E501
            targets ([WorkflowJobPayloadTargetsInner], none_type): [optional]  # noqa: E501
            instance_label (str): Instance Label. Only applicable if `targetType` is `instance-label`.. [optional]  # noqa: E501
            server_label (str): Server Label. Only applicable if `targetType` is `server-label`.. [optional]  # noqa: E501
            custom_options ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}): Map of options to be used as values in the workflow tasks. These correspond to option types.. [optional]  # noqa: E501
            custom_config (str): Job custom configuration (String in JSON format). [optional]  # noqa: E501
            date_time (datetime): Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is 'dateTime'.. [optional]  # noqa: E501
            run (bool): If true, executes job. [optional]  # noqa: E501
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
        self.workflow = workflow
        self.target_type = target_type
        self.schedule_mode = schedule_mode
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
    def __init__(self, name, workflow, target_type, schedule_mode, *args, **kwargs):  # noqa: E501
        """WorkflowJobPayload - a model defined in OpenAPI

        Args:
            name (str): A name for the Job
            workflow (WorkflowJobPayloadTask):
            target_type (str): Target type where job will execute
            schedule_mode (WorkflowJobPayloadScheduleMode):

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
            labels ([str], none_type): Array of label strings, can be used for filtering.. [optional]  # noqa: E501
            enabled (bool): Use this to set enabled state. [optional] if omitted the server will use the default value of True  # noqa: E501
            task (WorkflowJobPayloadTask): [optional]  # noqa: E501
            targets ([WorkflowJobPayloadTargetsInner], none_type): [optional]  # noqa: E501
            instance_label (str): Instance Label. Only applicable if `targetType` is `instance-label`.. [optional]  # noqa: E501
            server_label (str): Server Label. Only applicable if `targetType` is `server-label`.. [optional]  # noqa: E501
            custom_options ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}): Map of options to be used as values in the workflow tasks. These correspond to option types.. [optional]  # noqa: E501
            custom_config (str): Job custom configuration (String in JSON format). [optional]  # noqa: E501
            date_time (datetime): Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is 'dateTime'.. [optional]  # noqa: E501
            run (bool): If true, executes job. [optional]  # noqa: E501
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
        self.workflow = workflow
        self.target_type = target_type
        self.schedule_mode = schedule_mode
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
