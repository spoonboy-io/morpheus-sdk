# openapi_client.AutomationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_execute_schedules**](AutomationApi.md#add_execute_schedules) | **POST** /api/execute-schedules | Creates a Execute Schedule
[**add_power_schedule_instances**](AutomationApi.md#add_power_schedule_instances) | **PUT** /api/power-schedules/{id}/add-instances | Add Instances to a Power Schedule
[**add_power_schedule_servers**](AutomationApi.md#add_power_schedule_servers) | **PUT** /api/power-schedules/{id}/add-servers | Add Servers to a Power Schedule
[**add_power_schedules**](AutomationApi.md#add_power_schedules) | **POST** /api/power-schedules | Creates a Power Schedule
[**add_scale_thresholds**](AutomationApi.md#add_scale_thresholds) | **POST** /api/scale-thresholds | Creates a Scale Threshold
[**add_tasks**](AutomationApi.md#add_tasks) | **POST** /api/tasks | Creates a Task
[**add_workflows**](AutomationApi.md#add_workflows) | **POST** /api/task-sets | Creates a Workflow
[**delete_power_schedule_instances**](AutomationApi.md#delete_power_schedule_instances) | **PUT** /api/power-schedules/{id}/remove-instances | Remove Instances from a Power Schedule
[**delete_power_schedule_servers**](AutomationApi.md#delete_power_schedule_servers) | **PUT** /api/power-schedules/{id}/remove-servers | Remove Servers from a Power Schedule
[**execute_execution_request**](AutomationApi.md#execute_execution_request) | **POST** /api/execution-request/execute | Executes an Execution Request
[**execute_tasks**](AutomationApi.md#execute_tasks) | **POST** /api/tasks/{id}/execute | Executes a Task
[**execute_workflows**](AutomationApi.md#execute_workflows) | **POST** /api/task-sets/{id}/execute | Executes a Workflow
[**get_execute_schedules**](AutomationApi.md#get_execute_schedules) | **GET** /api/execute-schedules/{id} | Retrieves a Specific Execute Schedule
[**get_execution_request**](AutomationApi.md#get_execution_request) | **GET** /api/execution-request/{uniqueId} | Retrieves a Specific Execution Request
[**get_power_schedules**](AutomationApi.md#get_power_schedules) | **GET** /api/power-schedules/{id} | Retrieves a Specific Power Schedule
[**get_scale_thresholds**](AutomationApi.md#get_scale_thresholds) | **GET** /api/scale-thresholds/{id} | Retrieves a Specific Scale Threshold
[**get_task_types**](AutomationApi.md#get_task_types) | **GET** /api/task-types/{id} | Retrieves a Specific Task Type
[**get_tasks**](AutomationApi.md#get_tasks) | **GET** /api/tasks/{id} | Retrieves a Specific Task
[**get_workflows**](AutomationApi.md#get_workflows) | **GET** /api/task-sets/{id} | Retrieves a Specific Workflow
[**list_execute_schedules**](AutomationApi.md#list_execute_schedules) | **GET** /api/execute-schedules | Retrieves all Execute Schedules
[**list_power_schedules**](AutomationApi.md#list_power_schedules) | **GET** /api/power-schedules | Retrieves all Power Schedules
[**list_scale_thresholds**](AutomationApi.md#list_scale_thresholds) | **GET** /api/scale-thresholds | Retrieves all Scale Thresholds
[**list_task_types**](AutomationApi.md#list_task_types) | **GET** /api/task-types | Retrieves all Task Types
[**list_tasks**](AutomationApi.md#list_tasks) | **GET** /api/tasks | Retrieves all Tasks
[**list_workflows**](AutomationApi.md#list_workflows) | **GET** /api/task-sets | Retrieves all Workflows
[**remove_execute_schedules**](AutomationApi.md#remove_execute_schedules) | **DELETE** /api/execute-schedules/{id} | Deletes a Execute Schedule
[**remove_power_schedules**](AutomationApi.md#remove_power_schedules) | **DELETE** /api/power-schedules/{id} | Deletes a Power Schedule
[**remove_scale_thresholds**](AutomationApi.md#remove_scale_thresholds) | **DELETE** /api/scale-thresholds/{id} | Deletes a Scale Threshold
[**remove_tasks**](AutomationApi.md#remove_tasks) | **DELETE** /api/tasks/{id} | Deletes a Task
[**remove_workflows**](AutomationApi.md#remove_workflows) | **DELETE** /api/task-sets/{id} | Deletes a Workflow
[**update_execute_schedules**](AutomationApi.md#update_execute_schedules) | **PUT** /api/execute-schedules/{id} | Updates a Execute Schedule
[**update_power_schedules**](AutomationApi.md#update_power_schedules) | **PUT** /api/power-schedules/{id} | Updates a Power Schedule
[**update_scale_thresholds**](AutomationApi.md#update_scale_thresholds) | **PUT** /api/scale-thresholds/{id} | Updates a Scale Threshold
[**update_tasks**](AutomationApi.md#update_tasks) | **PUT** /api/tasks/{id} | Updates a Task
[**update_workflows**](AutomationApi.md#update_workflows) | **PUT** /api/task-sets/{id} | Updates a Workflow


# **add_execute_schedules**
> AddExecuteSchedules200Response add_execute_schedules()

Creates a Execute Schedule

Creates a execute schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.add_execute_schedules200_response import AddExecuteSchedules200Response
from openapi_client.model.add_execute_schedules_request import AddExecuteSchedulesRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    add_execute_schedules_request = AddExecuteSchedulesRequest(
        schedule=AddExecuteSchedulesRequestSchedule(
            name="Sample Execution",
            description="description_example",
            schedule_type="execute",
            schedule_timezone="UTC",
            cron="0 0 * * *",
            enabled=True,
        ),
    ) # AddExecuteSchedulesRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Execute Schedule
        api_response = api_instance.add_execute_schedules(add_execute_schedules_request=add_execute_schedules_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->add_execute_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_execute_schedules_request** | [**AddExecuteSchedulesRequest**](AddExecuteSchedulesRequest.md)|  | [optional]

### Return type

[**AddExecuteSchedules200Response**](AddExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_power_schedule_instances**
> Model200SuccessExpanded add_power_schedule_instances(id)

Add Instances to a Power Schedule

Add Instances to a Power Schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.model200_success_expanded import Model200SuccessExpanded
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_power_schedule_instances_request import AddPowerScheduleInstancesRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_power_schedule_instances_request = AddPowerScheduleInstancesRequest(
        instances=[
            1,
        ],
    ) # AddPowerScheduleInstancesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Add Instances to a Power Schedule
        api_response = api_instance.add_power_schedule_instances(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->add_power_schedule_instances: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Add Instances to a Power Schedule
        api_response = api_instance.add_power_schedule_instances(id, add_power_schedule_instances_request=add_power_schedule_instances_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->add_power_schedule_instances: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_power_schedule_instances_request** | [**AddPowerScheduleInstancesRequest**](AddPowerScheduleInstancesRequest.md)|  | [optional]

### Return type

[**Model200SuccessExpanded**](Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_power_schedule_servers**
> Model200SuccessExpanded add_power_schedule_servers(id)

Add Servers to a Power Schedule

Add Servers to a Power Schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.model200_success_expanded import Model200SuccessExpanded
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_power_schedule_servers_request import AddPowerScheduleServersRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_power_schedule_servers_request = AddPowerScheduleServersRequest(
        servers=[
            1,
        ],
    ) # AddPowerScheduleServersRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Add Servers to a Power Schedule
        api_response = api_instance.add_power_schedule_servers(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->add_power_schedule_servers: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Add Servers to a Power Schedule
        api_response = api_instance.add_power_schedule_servers(id, add_power_schedule_servers_request=add_power_schedule_servers_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->add_power_schedule_servers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_power_schedule_servers_request** | [**AddPowerScheduleServersRequest**](AddPowerScheduleServersRequest.md)|  | [optional]

### Return type

[**Model200SuccessExpanded**](Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_power_schedules**
> AddPowerSchedules200Response add_power_schedules()

Creates a Power Schedule

Creates a power schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.add_power_schedules_request import AddPowerSchedulesRequest
from openapi_client.model.add_power_schedules200_response import AddPowerSchedules200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    add_power_schedules_request = AddPowerSchedulesRequest(
        schedule=AddPowerSchedulesRequestSchedule(
            name="Sample Threshold",
            description="description_example",
            schedule_type="power",
            schedule_timezone="UTC",
            enabled=True,
            monday_on_time="00:00",
            monday_off_time="24:00",
            tuesday_on_time="00:00",
            tuesday_off_time="24:00",
            wednesday_on_time="00:00",
            wednesday_off_time="24:00",
            thursday_on_time="00:00",
            thursday_off_time="24:00",
            friday_on_time="00:00",
            friday_off_time="24:00",
            saturday_on_time="00:00",
            saturday_off_time="24:00",
            sunday_on_time="00:00",
            sunday_off_time="24:00",
        ),
    ) # AddPowerSchedulesRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Power Schedule
        api_response = api_instance.add_power_schedules(add_power_schedules_request=add_power_schedules_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->add_power_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_power_schedules_request** | [**AddPowerSchedulesRequest**](AddPowerSchedulesRequest.md)|  | [optional]

### Return type

[**AddPowerSchedules200Response**](AddPowerSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_scale_thresholds**
> AddScaleThresholds200Response add_scale_thresholds()

Creates a Scale Threshold

Creates a scale threshold. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.add_scale_thresholds_request import AddScaleThresholdsRequest
from openapi_client.model.add_scale_thresholds200_response import AddScaleThresholds200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    add_scale_thresholds_request = AddScaleThresholdsRequest(
        scale_threshold=AddScaleThresholdsRequestScaleThreshold(
            name="Sample Threshold",
            auto_up=False,
            auto_down=False,
            min_count=1,
            max_count=3,
            cpu_enabled=False,
            min_cpu=0,
            max_cpu=0,
            memory_enabled=False,
            min_memory=0,
            max_memory=0,
            disk_enabled=False,
            min_disk=0,
            max_disk=0,
        ),
    ) # AddScaleThresholdsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Scale Threshold
        api_response = api_instance.add_scale_thresholds(add_scale_thresholds_request=add_scale_thresholds_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->add_scale_thresholds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_scale_thresholds_request** | [**AddScaleThresholdsRequest**](AddScaleThresholdsRequest.md)|  | [optional]

### Return type

[**AddScaleThresholds200Response**](AddScaleThresholds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_tasks**
> AddTasks200Response add_tasks()

Creates a Task

Creates a task. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.add_tasks200_response import AddTasks200Response
from openapi_client.model.add_tasks_request import AddTasksRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    add_tasks_request = AddTasksRequest(
        task=AddTasksRequestTask(
            name="Sample Task",
            code="aTaskCode",
            visibility="private",
            task_type=AddTasksRequestTaskTaskType(
                code="ansibleTask",
            ),
            labels=[
                "labels_example",
            ],
            task_options={},
            result_type="exitCode",
            execute_target="local",
            retryable=False,
            retry_count=1,
            retry_delay_seconds=3.14,
            file=AddTasksRequestTaskFile(
                source_type="local",
                content="content_example",
                content_path="content_path_example",
                content_ref="content_ref_example",
                repository=AddTasksRequestTaskFileRepository(
                    id=1,
                ),
            ),
            credential=AddTasksRequestTaskCredential(None),
        ),
    ) # AddTasksRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Task
        api_response = api_instance.add_tasks(add_tasks_request=add_tasks_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->add_tasks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_tasks_request** | [**AddTasksRequest**](AddTasksRequest.md)|  | [optional]

### Return type

[**AddTasks200Response**](AddTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_workflows**
> AddWorkflows200Response add_workflows()

Creates a Workflow

Creates a workflow. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_workflows_request import AddWorkflowsRequest
from openapi_client.model.add_workflows200_response import AddWorkflows200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    add_workflows_request = AddWorkflowsRequest(
        task_set=AddWorkflowsRequestTaskSet(
            name="Sample Workflow",
            description="description_example",
            labels=[
                "labels_example",
            ],
            type="provision",
            visibility="private",
            option_types=[
                1,
            ],
            tasks=AddWorkflowsRequestTaskSetTasks(
                task_id=1,
                task_phase="provision",
            ),
        ),
    ) # AddWorkflowsRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Creates a Workflow
        api_response = api_instance.add_workflows(add_workflows_request=add_workflows_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->add_workflows: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **add_workflows_request** | [**AddWorkflowsRequest**](AddWorkflowsRequest.md)|  | [optional]

### Return type

[**AddWorkflows200Response**](AddWorkflows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_power_schedule_instances**
> Model200SuccessExpanded delete_power_schedule_instances(id)

Remove Instances from a Power Schedule

Remove Instances from a Power Schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.model200_success_expanded import Model200SuccessExpanded
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_power_schedule_instances_request import AddPowerScheduleInstancesRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_power_schedule_instances_request = AddPowerScheduleInstancesRequest(
        instances=[
            1,
        ],
    ) # AddPowerScheduleInstancesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Remove Instances from a Power Schedule
        api_response = api_instance.delete_power_schedule_instances(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->delete_power_schedule_instances: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Remove Instances from a Power Schedule
        api_response = api_instance.delete_power_schedule_instances(id, add_power_schedule_instances_request=add_power_schedule_instances_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->delete_power_schedule_instances: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_power_schedule_instances_request** | [**AddPowerScheduleInstancesRequest**](AddPowerScheduleInstancesRequest.md)|  | [optional]

### Return type

[**Model200SuccessExpanded**](Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_power_schedule_servers**
> Model200SuccessExpanded delete_power_schedule_servers(id)

Remove Servers from a Power Schedule

Remove Servers from a Power Schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.model200_success_expanded import Model200SuccessExpanded
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_power_schedule_servers_request import AddPowerScheduleServersRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    add_power_schedule_servers_request = AddPowerScheduleServersRequest(
        servers=[
            1,
        ],
    ) # AddPowerScheduleServersRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Remove Servers from a Power Schedule
        api_response = api_instance.delete_power_schedule_servers(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->delete_power_schedule_servers: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Remove Servers from a Power Schedule
        api_response = api_instance.delete_power_schedule_servers(id, add_power_schedule_servers_request=add_power_schedule_servers_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->delete_power_schedule_servers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **add_power_schedule_servers_request** | [**AddPowerScheduleServersRequest**](AddPowerScheduleServersRequest.md)|  | [optional]

### Return type

[**Model200SuccessExpanded**](Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **execute_execution_request**
> ExecuteExecutionRequest200Response execute_execution_request()

Executes an Execution Request

Provides API interfaces for executing an arbitrary script or command on an instance, container or host. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.execute_execution_request_request import ExecuteExecutionRequestRequest
from openapi_client.model.execute_execution_request200_response import ExecuteExecutionRequest200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    instance_id = 94 # int | The Instance ID for Filtering (optional)
    container_id = 135 # int | The Container ID for Filtering (optional)
    server_id = 97 # int | The Server ID for Filtering (optional)
    execute_execution_request_request = ExecuteExecutionRequestRequest(
        script="uname -a",
    ) # ExecuteExecutionRequestRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Executes an Execution Request
        api_response = api_instance.execute_execution_request(instance_id=instance_id, container_id=container_id, server_id=server_id, execute_execution_request_request=execute_execution_request_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->execute_execution_request: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **container_id** | **int**| The Container ID for Filtering | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **execute_execution_request_request** | [**ExecuteExecutionRequestRequest**](ExecuteExecutionRequestRequest.md)|  | [optional]

### Return type

[**ExecuteExecutionRequest200Response**](ExecuteExecutionRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Execution Request Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **execute_tasks**
> ExecuteTasks200Response execute_tasks(id)

Executes a Task

Executes a task. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.execute_tasks200_response import ExecuteTasks200Response
from openapi_client.model.execute_tasks_request import ExecuteTasksRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    execute_tasks_request = ExecuteTasksRequest(
        job=ExecuteTasksRequestJob(
            name="name_example",
            target_type="appliance",
            instances=[
                1,
            ],
            servers=[
                1,
            ],
            instance_label="instance_label_example",
            server_label="server_label_example",
            custom_options={},
            custom_config="custom_config_example",
        ),
    ) # ExecuteTasksRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Executes a Task
        api_response = api_instance.execute_tasks(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->execute_tasks: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Executes a Task
        api_response = api_instance.execute_tasks(id, execute_tasks_request=execute_tasks_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->execute_tasks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **execute_tasks_request** | [**ExecuteTasksRequest**](ExecuteTasksRequest.md)|  | [optional]

### Return type

[**ExecuteTasks200Response**](ExecuteTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **execute_workflows**
> ExecuteTasks200Response execute_workflows(id)

Executes a Workflow

Executes a workflow. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.execute_tasks200_response import ExecuteTasks200Response
from openapi_client.model.execute_tasks_request import ExecuteTasksRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    execute_tasks_request = ExecuteTasksRequest(
        job=ExecuteTasksRequestJob(
            name="name_example",
            target_type="appliance",
            instances=[
                1,
            ],
            servers=[
                1,
            ],
            instance_label="instance_label_example",
            server_label="server_label_example",
            custom_options={},
            custom_config="custom_config_example",
        ),
    ) # ExecuteTasksRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Executes a Workflow
        api_response = api_instance.execute_workflows(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->execute_workflows: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Executes a Workflow
        api_response = api_instance.execute_workflows(id, execute_tasks_request=execute_tasks_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->execute_workflows: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **execute_tasks_request** | [**ExecuteTasksRequest**](ExecuteTasksRequest.md)|  | [optional]

### Return type

[**ExecuteTasks200Response**](ExecuteTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_execute_schedules**
> GetExecuteSchedules200Response get_execute_schedules(id)

Retrieves a Specific Execute Schedule

Retrieves a specific execute schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.get_execute_schedules200_response import GetExecuteSchedules200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Execute Schedule
        api_response = api_instance.get_execute_schedules(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->get_execute_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetExecuteSchedules200Response**](GetExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_execution_request**
> GetExecutionRequest200Response get_execution_request(unique_id)

Retrieves a Specific Execution Request

Retrieves a specific execution request. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.get_execution_request200_response import GetExecutionRequest200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    unique_id = "c56f75d0-a59a-4566-92e3-4dc716c5d076" # str | The Unique ID of the execution request

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Execution Request
        api_response = api_instance.get_execution_request(unique_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->get_execution_request: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unique_id** | **str**| The Unique ID of the execution request |

### Return type

[**GetExecutionRequest200Response**](GetExecutionRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_power_schedules**
> GetPowerSchedules200Response get_power_schedules(id)

Retrieves a Specific Power Schedule

Retrieves a specific power schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_power_schedules200_response import GetPowerSchedules200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Power Schedule
        api_response = api_instance.get_power_schedules(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->get_power_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetPowerSchedules200Response**](GetPowerSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_scale_thresholds**
> GetScaleThresholds200Response get_scale_thresholds(id)

Retrieves a Specific Scale Threshold

Retrieves a specific scale threshold. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_scale_thresholds200_response import GetScaleThresholds200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Scale Threshold
        api_response = api_instance.get_scale_thresholds(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->get_scale_thresholds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetScaleThresholds200Response**](GetScaleThresholds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_task_types**
> GetTaskTypes200Response get_task_types(id)

Retrieves a Specific Task Type

Retrieves a specific task type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_task_types200_response import GetTaskTypes200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Task Type
        api_response = api_instance.get_task_types(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->get_task_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetTaskTypes200Response**](GetTaskTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_tasks**
> GetTasks200Response get_tasks(id)

Retrieves a Specific Task

Retrieves a specific task. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.get_tasks200_response import GetTasks200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Task
        api_response = api_instance.get_tasks(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->get_tasks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetTasks200Response**](GetTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_workflows**
> GetWorkflows200Response get_workflows(id)

Retrieves a Specific Workflow

Retrieves a specific workflow. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_workflows200_response import GetWorkflows200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Workflow
        api_response = api_instance.get_workflows(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->get_workflows: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetWorkflows200Response**](GetWorkflows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_execute_schedules**
> ListExecuteSchedules200Response list_execute_schedules()

Retrieves all Execute Schedules

Retrieves all execute schedules. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.list_execute_schedules200_response import ListExecuteSchedules200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Execute Schedules
        api_response = api_instance.list_execute_schedules(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->list_execute_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**ListExecuteSchedules200Response**](ListExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_power_schedules**
> ListPowerSchedules200Response list_power_schedules()

Retrieves all Power Schedules

Retrieves all power schedules. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.list_power_schedules200_response import ListPowerSchedules200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Power Schedules
        api_response = api_instance.list_power_schedules(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->list_power_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**ListPowerSchedules200Response**](ListPowerSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_scale_thresholds**
> ListScaleThresholds200Response list_scale_thresholds()

Retrieves all Scale Thresholds

Retrieves all scale thresholds. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.list_scale_thresholds200_response import ListScaleThresholds200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Scale Thresholds
        api_response = api_instance.list_scale_thresholds(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->list_scale_thresholds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**ListScaleThresholds200Response**](ListScaleThresholds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_task_types**
> ListTaskTypes200Response list_task_types()

Retrieves all Task Types

A Task Type is a type of automation task. Each type defines its own set of options to be configured for each task. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.list_task_types200_response import ListTaskTypes200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Task Types
        api_response = api_instance.list_task_types(name=name, code=code)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->list_task_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **str**| If specified will return an exact match on code | [optional]

### Return type

[**ListTaskTypes200Response**](ListTaskTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_tasks**
> ListTasks200Response list_tasks()

Retrieves all Tasks

Retrieves all tasks. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.list_tasks200_response import ListTasks200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    type = "ansibleTask" # str | If specified will return all tasks by `task type` code. Refer to `Task Types` API for up to date listings.  (optional)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Tasks
        api_response = api_instance.list_tasks(type=type, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->list_tasks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type** | **str**| If specified will return all tasks by &#x60;task type&#x60; code. Refer to &#x60;Task Types&#x60; API for up to date listings.  | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListTasks200Response**](ListTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_workflows**
> ListWorkflows200Response list_workflows()

Retrieves all Workflows

Retrieves all workflows. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.list_workflows200_response import ListWorkflows200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)
    type = "provision" # str | Filter by Workflow Type (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Retrieves all Workflows
        api_response = api_instance.list_workflows(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, labels=labels, all_labels=all_labels, type=type)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->list_workflows: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **type** | **str**| Filter by Workflow Type | [optional]

### Return type

[**ListWorkflows200Response**](ListWorkflows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_execute_schedules**
> Model200Success remove_execute_schedules(id)

Deletes a Execute Schedule

Deletes a specified execute schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Execute Schedule
        api_response = api_instance.remove_execute_schedules(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->remove_execute_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_power_schedules**
> Model200Success remove_power_schedules(id)

Deletes a Power Schedule

Deletes a specified power schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Power Schedule
        api_response = api_instance.remove_power_schedules(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->remove_power_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_scale_thresholds**
> Model200Success remove_scale_thresholds(id)

Deletes a Scale Threshold

Deletes a specified scale threshold. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Scale Threshold
        api_response = api_instance.remove_scale_thresholds(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->remove_scale_thresholds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_tasks**
> Model200Success remove_tasks(id)

Deletes a Task

Deletes a specified task. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Task
        api_response = api_instance.remove_tasks(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->remove_tasks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_workflows**
> Model200Success remove_workflows(id)

Deletes a Workflow

Deletes a specified workflow. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Deletes a Workflow
        api_response = api_instance.remove_workflows(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->remove_workflows: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_execute_schedules**
> AddExecuteSchedules200Response update_execute_schedules(id)

Updates a Execute Schedule

Updates a execute schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.add_execute_schedules200_response import AddExecuteSchedules200Response
from openapi_client.model.update_execute_schedules_request import UpdateExecuteSchedulesRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_execute_schedules_request = UpdateExecuteSchedulesRequest(
        schedule=UpdateExecuteSchedulesRequestSchedule(
            name="Sample Execution",
            description="description_example",
            schedule_type="execute",
            schedule_timezone="UTC",
            cron="0 0 * * *",
            enabled=True,
        ),
    ) # UpdateExecuteSchedulesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Execute Schedule
        api_response = api_instance.update_execute_schedules(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->update_execute_schedules: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Execute Schedule
        api_response = api_instance.update_execute_schedules(id, update_execute_schedules_request=update_execute_schedules_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->update_execute_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_execute_schedules_request** | [**UpdateExecuteSchedulesRequest**](UpdateExecuteSchedulesRequest.md)|  | [optional]

### Return type

[**AddExecuteSchedules200Response**](AddExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_power_schedules**
> AddPowerSchedules200Response update_power_schedules(id)

Updates a Power Schedule

Updates a power schedule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.add_power_schedules200_response import AddPowerSchedules200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_power_schedules_request import UpdatePowerSchedulesRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_power_schedules_request = UpdatePowerSchedulesRequest(
        schedule=UpdatePowerSchedulesRequestSchedule(
            name="Sample Threshold",
            description="description_example",
            schedule_type="power",
            schedule_timezone="UTC",
            enabled=True,
            monday_on_time="00:00",
            monday_off_time="24:00",
            tuesday_on_time="00:00",
            tuesday_off_time="24:00",
            wednesday_on_time="00:00",
            wednesday_off_time="24:00",
            thursday_on_time="00:00",
            thursday_off_time="24:00",
            friday_on_time="00:00",
            friday_off_time="24:00",
            saturday_on_time="00:00",
            saturday_off_time="24:00",
            sunday_on_time="00:00",
            sunday_off_time="24:00",
        ),
    ) # UpdatePowerSchedulesRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Power Schedule
        api_response = api_instance.update_power_schedules(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->update_power_schedules: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Power Schedule
        api_response = api_instance.update_power_schedules(id, update_power_schedules_request=update_power_schedules_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->update_power_schedules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_power_schedules_request** | [**UpdatePowerSchedulesRequest**](UpdatePowerSchedulesRequest.md)|  | [optional]

### Return type

[**AddPowerSchedules200Response**](AddPowerSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_scale_thresholds**
> AddScaleThresholds200Response update_scale_thresholds(id)

Updates a Scale Threshold

Updates a scale threshold. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.add_scale_thresholds200_response import AddScaleThresholds200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_scale_thresholds_request import UpdateScaleThresholdsRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_scale_thresholds_request = UpdateScaleThresholdsRequest(
        scale_threshold=UpdateScaleThresholdsRequestScaleThreshold(
            name="Sample Threshold",
            auto_up=False,
            auto_down=False,
            min_count=1,
            max_count=3,
            cpu_enabled=False,
            min_cpu=0,
            max_cpu=0,
            memory_enabled=False,
            min_memory=0,
            max_memory=0,
            disk_enabled=False,
            min_disk=0,
            max_disk=0,
        ),
    ) # UpdateScaleThresholdsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Scale Threshold
        api_response = api_instance.update_scale_thresholds(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->update_scale_thresholds: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Scale Threshold
        api_response = api_instance.update_scale_thresholds(id, update_scale_thresholds_request=update_scale_thresholds_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->update_scale_thresholds: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_scale_thresholds_request** | [**UpdateScaleThresholdsRequest**](UpdateScaleThresholdsRequest.md)|  | [optional]

### Return type

[**AddScaleThresholds200Response**](AddScaleThresholds200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_tasks**
> AddTasks200Response update_tasks(id)

Updates a Task

Updates a task. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.add_tasks200_response import AddTasks200Response
from openapi_client.model.update_tasks_request import UpdateTasksRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_tasks_request = UpdateTasksRequest(
        task=UpdateTasksRequestTask(
            name="Sample Task",
            code="aTaskCode",
            visibility="private",
            task_type=AddTasksRequestTaskTaskType(
                code="ansibleTask",
            ),
            labels=[
                "labels_example",
            ],
            task_options={},
            result_type="exitCode",
            execute_target="local",
            retryable=False,
            retry_count=1,
            retry_delay_seconds=3.14,
            file=AddTasksRequestTaskFile(
                source_type="local",
                content="content_example",
                content_path="content_path_example",
                content_ref="content_ref_example",
                repository=AddTasksRequestTaskFileRepository(
                    id=1,
                ),
            ),
            credential=UpdateTasksRequestTaskCredential(None),
        ),
    ) # UpdateTasksRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Task
        api_response = api_instance.update_tasks(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->update_tasks: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Task
        api_response = api_instance.update_tasks(id, update_tasks_request=update_tasks_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->update_tasks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_tasks_request** | [**UpdateTasksRequest**](UpdateTasksRequest.md)|  | [optional]

### Return type

[**AddTasks200Response**](AddTasks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_workflows**
> AddWorkflows200Response update_workflows(id)

Updates a Workflow

Updates a workflow. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import automation_api
from openapi_client.model.update_workflows_request import UpdateWorkflowsRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.add_workflows200_response import AddWorkflows200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = automation_api.AutomationApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_workflows_request = UpdateWorkflowsRequest(
        task_set=UpdateWorkflowsRequestTaskSet(
            name="Sample Workflow",
            description="description_example",
            labels=[
                "labels_example",
            ],
            type="provision",
            option_types=[
                1,
            ],
            tasks=AddWorkflowsRequestTaskSetTasks(
                task_id=1,
                task_phase="provision",
            ),
        ),
    ) # UpdateWorkflowsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Updates a Workflow
        api_response = api_instance.update_workflows(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->update_workflows: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Updates a Workflow
        api_response = api_instance.update_workflows(id, update_workflows_request=update_workflows_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling AutomationApi->update_workflows: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_workflows_request** | [**UpdateWorkflowsRequest**](UpdateWorkflowsRequest.md)|  | [optional]

### Return type

[**AddWorkflows200Response**](AddWorkflows200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

