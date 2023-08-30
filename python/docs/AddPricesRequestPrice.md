# AddPricesRequestPrice


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Price name | 
**code** | **str** | Price code, must be unique | 
**price_type** | **str** | Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software * &#x60;load_balancer&#x60; - Load Balancer * &#x60;load_balancer_virtual_server&#x60; - Load Balancer Virtual Server  | 
**price_unit** | **str** | The unit of pricing | 
**incur_charges** | **str** | Indicates when to incur charge | 
**currency** | **str** | ISO Currency code | 
**cost** | **float** | Cost | 
**account** | [**AddPricesRequestPriceAccount**](AddPricesRequestPriceAccount.md) |  | [optional] 
**markup_type** | **str** | Price adjustment type | [optional] 
**markup** | **float** | Amount for &#x60;fixed&#x60; price adjustment type | [optional] 
**markup_percent** | **float** | Percent for &#x60;percent&#x60; price adjustment type | [optional] 
**custom_price** | **float** | Custom price for &#x60;custom&#x60; price adjustment type | [optional] 
**platform** | **str** | Platform.  Required for &#x60;platform&#x60; price type | [optional] 
**software** | **str** | Software.  Required for software price type | [optional] 
**volume_type** | [**AddPricesRequestPriceVolumeType**](AddPricesRequestPriceVolumeType.md) |  | [optional] 
**datastore** | [**AddPricesRequestPriceDatastore**](AddPricesRequestPriceDatastore.md) |  | [optional] 
**cross_cloud_apply** | **bool** | Apply price across clouds, optional true/false flag for datastore price type | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


