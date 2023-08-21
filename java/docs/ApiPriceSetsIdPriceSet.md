

# ApiPriceSetsIdPriceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Price set name |  [optional]
**code** | **String** | Price set code. Must be unique. |  [optional]
**regionCode** | **String** | Price set region code |  [optional]
**zone** | [**ApiPriceSetsPriceSetZone**](ApiPriceSetsPriceSetZone.md) |  |  [optional]
**zonePool** | [**ApiPriceSetsPriceSetZonePool**](ApiPriceSetsPriceSetZonePool.md) |  |  [optional]
**priceUnit** | [**PriceUnitEnum**](#PriceUnitEnum) | Price Unit |  [optional]
**type** | [**TypeEnum**](#TypeEnum) | Price set type |  [optional]
**prices** | **List&lt;Long&gt;** |  |  [optional]



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



## Enum: TypeEnum

Name | Value
---- | -----
FIXED | &quot;fixed&quot;
COMPUTE_PLUS_STORAGE | &quot;compute_plus_storage&quot;
COMPONENT | &quot;component&quot;
LOAD_BALANCER | &quot;load_balancer&quot;
SNAPSHOT | &quot;snapshot&quot;
VIRTUAL_IMAGE | &quot;virtual_image&quot;
SOFTWARE_OR_SERVICE | &quot;software_or_service&quot;



