# # ServerServicePlans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional]
**name** | **string** |  | [optional]
**value** | **int** |  | [optional]
**code** | **string** |  | [optional]
**max_storage** | **int** |  | [optional]
**max_memory** | **int** |  | [optional]
**max_cpu** | **int** |  | [optional]
**max_cores** | **int** |  | [optional]
**max_data_storage** | **int** |  | [optional]
**custom_cpu** | **bool** |  | [optional]
**custom_max_memory** | **bool** |  | [optional]
**custom_max_storage** | **bool** |  | [optional]
**custom_max_data_storage** | **bool** |  | [optional]
**custom_cores_per_socket** | **bool** |  | [optional]
**cores_per_socket** | **int** |  | [optional]
**storage_types** | **object[]** |  | [optional]
**root_storage_types** | **object[]** |  | [optional]
**add_volumes** | **bool** |  | [optional]
**customize_volume** | **bool** |  | [optional]
**root_disk_customizable** | **bool** |  | [optional]
**host_disk_mode** | **string** |  | [optional]
**has_datastore** | **string** |  | [optional]
**lvm_supported** | **string** |  | [optional]
**min_disk** | **string** |  | [optional]
**max_disk** | **string** |  | [optional]
**datastores** | [**\OpenAPI\Client\Model\ServerServicePlansDatastores**](ServerServicePlansDatastores.md) |  | [optional]
**supports_auto_datastore** | **bool** |  | [optional]
**auto_options** | [**\OpenAPI\Client\Model\InstanceServicePlanAutoOptions[]**](InstanceServicePlanAutoOptions.md) |  | [optional]
**cpu_options** | **object[]** |  | [optional]
**memory_options** | **object[]** |  | [optional]
**root_custom_size_options** | **object** |  | [optional]
**custom_size_options** | **object** |  | [optional]
**custom_cores** | **bool** |  | [optional]
**max_disks** | **string** |  | [optional]
**memory_size_type** | **string** |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
