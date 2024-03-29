# # Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional]
**owner_id** | **int** |  | [optional]
**account** | [**\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional]
**group** | **object** |  | [optional]
**cloud** | [**\OpenAPI\Client\Model\InvoiceCloud**](InvoiceCloud.md) |  | [optional]
**instance** | **object** |  | [optional]
**server** | **string** |  | [optional]
**cluster** | **string** |  | [optional]
**user** | **object** |  | [optional]
**plan** | **object** |  | [optional]
**tags** | **object[]** |  | [optional]
**project** | **string** |  | [optional]
**ref_type** | **string** |  | [optional]
**ref_id** | **int** |  | [optional]
**ref_uuid** | **string** |  | [optional]
**ref_name** | **string** |  | [optional]
**ref_category** | **string** |  | [optional]
**resource_id** | **string** |  | [optional]
**resource_uuid** | **string** |  | [optional]
**resource_type** | **string** |  | [optional]
**resource_name** | **string** |  | [optional]
**resource_external_id** | **string** |  | [optional]
**resource_internal_id** | **string** |  | [optional]
**interval** | **string** |  | [optional]
**period** | **string** |  | [optional]
**estimate** | **bool** |  | [optional]
**summary_invoice** | **bool** |  | [optional]
**active** | **bool** |  | [optional]
**start_date** | [**\DateTime**](\DateTime.md) |  | [optional]
**end_date** | [**\DateTime**](\DateTime.md) |  | [optional]
**ref_start** | [**\DateTime**](\DateTime.md) |  | [optional]
**ref_end** | [**\DateTime**](\DateTime.md) |  | [optional]
**estimated_compute_price** | **float** |  | [optional]
**estimated_compute_cost** | **float** |  | [optional]
**estimated_memory_price** | **float** |  | [optional]
**estimated_memory_cost** | **float** |  | [optional]
**estimated_storage_price** | **float** |  | [optional]
**estimated_storage_cost** | **float** |  | [optional]
**estimated_network_price** | **float** |  | [optional]
**estimated_network_cost** | **float** |  | [optional]
**estimated_license_price** | **float** |  | [optional]
**estimated_license_cost** | **float** |  | [optional]
**estimated_extra_price** | **float** |  | [optional]
**estimated_extra_cost** | **float** |  | [optional]
**estimated_total_price** | **float** |  | [optional]
**estimated_total_cost** | **float** |  | [optional]
**estimated_running_price** | **float** |  | [optional]
**estimated_running_cost** | **float** |  | [optional]
**estimated_currency** | **string** |  | [optional]
**estimated_conversion_rate** | **float** |  | [optional]
**actual_compute_price** | **float** |  | [optional]
**actual_compute_cost** | **float** |  | [optional]
**actual_memory_price** | **float** |  | [optional]
**actual_memory_cost** | **float** |  | [optional]
**actual_storage_price** | **float** |  | [optional]
**actual_storage_cost** | **float** |  | [optional]
**actual_network_price** | **float** |  | [optional]
**actual_network_cost** | **float** |  | [optional]
**actual_license_price** | **float** |  | [optional]
**actual_license_cost** | **float** |  | [optional]
**actual_extra_price** | **float** |  | [optional]
**actual_extra_cost** | **float** |  | [optional]
**actual_total_price** | **float** |  | [optional]
**actual_total_cost** | **float** |  | [optional]
**actual_running_price** | **float** |  | [optional]
**actual_running_cost** | **float** |  | [optional]
**actual_currency** | **string** |  | [optional]
**actual_conversion_rate** | **float** |  | [optional]
**compute_price** | **float** |  | [optional]
**compute_cost** | **float** |  | [optional]
**memory_price** | **float** |  | [optional]
**memory_cost** | **float** |  | [optional]
**storage_price** | **float** |  | [optional]
**storage_cost** | **float** |  | [optional]
**network_price** | **float** |  | [optional]
**network_cost** | **float** |  | [optional]
**license_price** | **float** |  | [optional]
**license_cost** | **float** |  | [optional]
**extra_price** | **float** |  | [optional]
**extra_cost** | **float** |  | [optional]
**total_price** | **float** |  | [optional]
**total_cost** | **float** |  | [optional]
**running_price** | **float** |  | [optional]
**running_cost** | **float** |  | [optional]
**currency** | **string** |  | [optional]
**conversion_rate** | **float** |  | [optional]
**cost_type** | **string** |  | [optional]
**off_time** | **int** |  | [optional]
**power_state** | **string** |  | [optional]
**power_date** | [**\DateTime**](\DateTime.md) |  | [optional]
**running_multiplier** | **float** |  | [optional]
**usage_type** | **string** |  | [optional]
**usage_category** | **string** |  | [optional]
**last_cost_date** | [**\DateTime**](\DateTime.md) |  | [optional]
**last_actual_date** | [**\DateTime**](\DateTime.md) |  | [optional]
**date_created** | [**\DateTime**](\DateTime.md) |  | [optional]
**last_updated** | [**\DateTime**](\DateTime.md) |  | [optional]
**line_item_count** | **int** |  | [optional]
**line_items** | [**\OpenAPI\Client\Model\InvoiceLineItems[]**](InvoiceLineItems.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
