# # ApiPricesPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Price name |
**code** | **string** | Price code, must be unique |
**account** | [**\OpenAPI\Client\Model\ApiPricesPriceAccount**](ApiPricesPriceAccount.md) |  | [optional]
**price_type** | **string** | Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software * &#x60;load_balancer&#x60; - Load Balancer * &#x60;load_balancer_virtual_server&#x60; - Load Balancer Virtual Server |
**price_unit** | **string** | The unit of pricing |
**incur_charges** | **string** | Indicates when to incur charge |
**currency** | **string** | ISO Currency code |
**cost** | **float** | Cost |
**markup_type** | **string** | Price adjustment type | [optional]
**markup** | **float** | Amount for &#x60;fixed&#x60; price adjustment type | [optional]
**markup_percent** | **float** | Percent for &#x60;percent&#x60; price adjustment type | [optional]
**custom_price** | **float** | Custom price for &#x60;custom&#x60; price adjustment type | [optional]
**platform** | **string** | Platform.  Required for &#x60;platform&#x60; price type | [optional]
**software** | **string** | Software.  Required for software price type | [optional]
**volume_type** | [**\OpenAPI\Client\Model\ApiPricesPriceVolumeType**](ApiPricesPriceVolumeType.md) |  | [optional]
**datastore** | [**\OpenAPI\Client\Model\ApiPricesPriceDatastore**](ApiPricesPriceDatastore.md) |  | [optional]
**cross_cloud_apply** | **bool** | Apply price across clouds, optional true/false flag for datastore price type | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
