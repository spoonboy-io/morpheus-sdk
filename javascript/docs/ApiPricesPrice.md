# MorpheusApi.ApiPricesPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Price name | 
**code** | **String** | Price code, must be unique | 
**account** | [**ApiPricesPriceAccount**](ApiPricesPriceAccount.md) |  | [optional] 
**priceType** | **String** | Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software * &#x60;load_balancer&#x60; - Load Balancer * &#x60;load_balancer_virtual_server&#x60; - Load Balancer Virtual Server  | 
**priceUnit** | **String** | The unit of pricing | 
**incurCharges** | **String** | Indicates when to incur charge | 
**currency** | **String** | ISO Currency code | 
**cost** | **Number** | Cost | 
**markupType** | **String** | Price adjustment type | [optional] 
**markup** | **Number** | Amount for &#x60;fixed&#x60; price adjustment type | [optional] 
**markupPercent** | **Number** | Percent for &#x60;percent&#x60; price adjustment type | [optional] 
**customPrice** | **Number** | Custom price for &#x60;custom&#x60; price adjustment type | [optional] 
**platform** | **String** | Platform.  Required for &#x60;platform&#x60; price type | [optional] 
**software** | **String** | Software.  Required for software price type | [optional] 
**volumeType** | [**ApiPricesPriceVolumeType**](ApiPricesPriceVolumeType.md) |  | [optional] 
**datastore** | [**ApiPricesPriceDatastore**](ApiPricesPriceDatastore.md) |  | [optional] 
**crossCloudApply** | **Boolean** | Apply price across clouds, optional true/false flag for datastore price type | [optional] 



## Enum: PriceTypeEnum


* `fixed` (value: `"fixed"`)

* `compute` (value: `"compute"`)

* `memory` (value: `"memory"`)

* `cores` (value: `"cores"`)

* `storage` (value: `"storage"`)

* `datastore` (value: `"datastore"`)

* `platform` (value: `"platform"`)

* `software` (value: `"software"`)

* `load_balancer` (value: `"load_balancer"`)

* `load_balancer_virtual_server` (value: `"load_balancer_virtual_server"`)





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





## Enum: IncurChargesEnum


* `running` (value: `"running"`)

* `stopped` (value: `"stopped"`)

* `always` (value: `"always"`)





## Enum: MarkupTypeEnum


* `fixed` (value: `"fixed"`)

* `percent` (value: `"percent"`)

* `custom` (value: `"custom"`)




