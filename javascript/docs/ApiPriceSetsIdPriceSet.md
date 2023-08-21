# MorpheusApi.ApiPriceSetsIdPriceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Price set name | [optional] 
**code** | **String** | Price set code. Must be unique. | [optional] 
**regionCode** | **String** | Price set region code | [optional] 
**zone** | [**ApiPriceSetsPriceSetZone**](ApiPriceSetsPriceSetZone.md) |  | [optional] 
**zonePool** | [**ApiPriceSetsPriceSetZonePool**](ApiPriceSetsPriceSetZonePool.md) |  | [optional] 
**priceUnit** | **String** | Price Unit | [optional] 
**type** | **String** | Price set type | [optional] 
**prices** | **[Number]** |  | [optional] 



## Enum: PriceUnitEnum


* `minute` (value: `"minute"`)

* `hour` (value: `"hour"`)

* `day` (value: `"day"`)

* `month` (value: `"month"`)

* `year` (value: `"year"`)

* `two year` (value: `"two year"`)

* `three year` (value: `"three year"`)

* `four year` (value: `"four year"`)

* `five year` (value: `"five year"`)





## Enum: TypeEnum


* `fixed` (value: `"fixed"`)

* `compute_plus_storage` (value: `"compute_plus_storage"`)

* `component` (value: `"component"`)

* `load_balancer` (value: `"load_balancer"`)

* `snapshot` (value: `"snapshot"`)

* `virtual_image` (value: `"virtual_image"`)

* `software_or_service` (value: `"software_or_service"`)




