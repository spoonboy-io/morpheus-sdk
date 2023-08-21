

# ApiPricesIdPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Price name |  [optional]
**code** | **String** | Price code, must be unique |  [optional]
**account** | [**ApiPricesPriceAccount**](ApiPricesPriceAccount.md) |  |  [optional]
**priceType** | [**PriceTypeEnum**](#PriceTypeEnum) | Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software * &#x60;load_balancer&#x60; - Load Balancer * &#x60;load_balancer_virtual_server&#x60; - Load Balancer Virtual Server  |  [optional]
**priceUnit** | [**PriceUnitEnum**](#PriceUnitEnum) | The unit of pricing |  [optional]
**incurCharges** | [**IncurChargesEnum**](#IncurChargesEnum) | Indicates when to incur charge |  [optional]
**currency** | **String** | ISO Currency code |  [optional]
**cost** | **BigDecimal** | Cost |  [optional]
**markupType** | [**MarkupTypeEnum**](#MarkupTypeEnum) | Price adjustment type |  [optional]
**markup** | **BigDecimal** | Amount for &#x60;fixed&#x60; price adjustment type |  [optional]
**markupPercent** | **BigDecimal** | Percent for &#x60;percent&#x60; price adjustment type |  [optional]
**customPrice** | **BigDecimal** | Custom price for &#x60;custom&#x60; price adjustment type |  [optional]
**platform** | **String** | Platform.  Required for &#x60;platform&#x60; price type |  [optional]
**software** | **String** | Software.  Required for software price type |  [optional]
**volumeType** | [**ApiPricesPriceVolumeType**](ApiPricesPriceVolumeType.md) |  |  [optional]
**datastore** | [**ApiPricesPriceDatastore**](ApiPricesPriceDatastore.md) |  |  [optional]
**crossCloudApply** | **Boolean** | Apply price across clouds, optional true/false flag for datastore price type |  [optional]



## Enum: PriceTypeEnum

Name | Value
---- | -----
FIXED | &quot;fixed&quot;
COMPUTE | &quot;compute&quot;
MEMORY | &quot;memory&quot;
CORES | &quot;cores&quot;
STORAGE | &quot;storage&quot;
DATASTORE | &quot;datastore&quot;
PLATFORM | &quot;platform&quot;
SOFTWARE | &quot;software&quot;
LOAD_BALANCER | &quot;load_balancer&quot;
LOAD_BALANCER_VIRTUAL_SERVER | &quot;load_balancer_virtual_server&quot;



## Enum: PriceUnitEnum

Name | Value
---- | -----
MINUTE | &quot;minute&quot;
HOUR | &quot;hour&quot;
DAY | &quot;day&quot;
MONTH | &quot;month&quot;
YEAR | &quot;year&quot;
TWO_YEAR | &quot;two year&quot;
THREE_YEAR | &quot;three year&quot;
FOUR_YEAR | &quot;four year&quot;
FIVE_YEAR | &quot;five year&quot;



## Enum: IncurChargesEnum

Name | Value
---- | -----
RUNNING | &quot;running&quot;
STOPPED | &quot;stopped&quot;
ALWAYS | &quot;always&quot;



## Enum: MarkupTypeEnum

Name | Value
---- | -----
FIXED | &quot;fixed&quot;
PERCENT | &quot;percent&quot;
CUSTOM | &quot;custom&quot;



