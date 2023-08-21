# # InlineObject47

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**mode** | **string** | Refresh Mode. Run the &#x60;daily&#x60; or &#x60;costing&#x60; job instead of the default &#x60;hourly&#x60; refresh job. | [optional]
**rebuild** | **string** | Rebuild. Pass &#x60;true&#x60; to purge existing invoices for the period before refreshing. Only applies to mode &#x60;costing&#x60;. | [optional]
**period** | **string** | Period. Invoice billing period to refresh in the format &#x60;YYYYMM&#x60;. The default period is the current month. Only applies to mode &#x60;costing&#x60;. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
