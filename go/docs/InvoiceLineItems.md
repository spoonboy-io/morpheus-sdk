# InvoiceLineItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **int64** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**RefName** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**ItemType** | Pointer to **NullableString** |  | [optional] 
**ItemName** | Pointer to **string** |  | [optional] 
**ItemDescription** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **NullableString** |  | [optional] 
**ProductCode** | Pointer to **string** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**ItemSeller** | Pointer to **NullableString** |  | [optional] 
**ItemAction** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**RateId** | Pointer to **string** |  | [optional] 
**RateClass** | Pointer to **NullableString** |  | [optional] 
**RateUnit** | Pointer to **string** |  | [optional] 
**RateTerm** | Pointer to **NullableString** |  | [optional] 
**UsageType** | Pointer to **string** |  | [optional] 
**UsageCategory** | Pointer to **string** |  | [optional] 
**UsageService** | Pointer to **string** |  | [optional] 
**ItemUsage** | Pointer to **float32** |  | [optional] 
**ItemRate** | Pointer to **float32** |  | [optional] 
**ItemCost** | Pointer to **float32** |  | [optional] 
**ItemPrice** | Pointer to **float32** |  | [optional] 
**ItemTax** | Pointer to **int64** |  | [optional] 
**ItemTerm** | Pointer to **NullableString** |  | [optional] 
**TaxType** | Pointer to **NullableString** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ConversionRate** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInvoiceLineItems

`func NewInvoiceLineItems() *InvoiceLineItems`

NewInvoiceLineItems instantiates a new InvoiceLineItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineItemsWithDefaults

`func NewInvoiceLineItemsWithDefaults() *InvoiceLineItems`

NewInvoiceLineItemsWithDefaults instantiates a new InvoiceLineItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceLineItems) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLineItems) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLineItems) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceLineItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *InvoiceLineItems) GetInvoiceId() int64`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceLineItems) GetInvoiceIdOk() (*int64, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceLineItems) SetInvoiceId(v int64)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *InvoiceLineItems) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetRefType

`func (o *InvoiceLineItems) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *InvoiceLineItems) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *InvoiceLineItems) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *InvoiceLineItems) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *InvoiceLineItems) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *InvoiceLineItems) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *InvoiceLineItems) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *InvoiceLineItems) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *InvoiceLineItems) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *InvoiceLineItems) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *InvoiceLineItems) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *InvoiceLineItems) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetStartDate

`func (o *InvoiceLineItems) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvoiceLineItems) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvoiceLineItems) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvoiceLineItems) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InvoiceLineItems) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InvoiceLineItems) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InvoiceLineItems) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InvoiceLineItems) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetItemId

`func (o *InvoiceLineItems) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *InvoiceLineItems) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *InvoiceLineItems) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *InvoiceLineItems) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemType

`func (o *InvoiceLineItems) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *InvoiceLineItems) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *InvoiceLineItems) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *InvoiceLineItems) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *InvoiceLineItems) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *InvoiceLineItems) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetItemName

`func (o *InvoiceLineItems) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *InvoiceLineItems) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *InvoiceLineItems) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *InvoiceLineItems) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetItemDescription

`func (o *InvoiceLineItems) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *InvoiceLineItems) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *InvoiceLineItems) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *InvoiceLineItems) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### GetProductId

`func (o *InvoiceLineItems) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *InvoiceLineItems) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *InvoiceLineItems) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *InvoiceLineItems) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *InvoiceLineItems) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *InvoiceLineItems) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetProductCode

`func (o *InvoiceLineItems) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *InvoiceLineItems) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *InvoiceLineItems) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *InvoiceLineItems) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetProductName

`func (o *InvoiceLineItems) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *InvoiceLineItems) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *InvoiceLineItems) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *InvoiceLineItems) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetItemSeller

`func (o *InvoiceLineItems) GetItemSeller() string`

GetItemSeller returns the ItemSeller field if non-nil, zero value otherwise.

### GetItemSellerOk

`func (o *InvoiceLineItems) GetItemSellerOk() (*string, bool)`

GetItemSellerOk returns a tuple with the ItemSeller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSeller

`func (o *InvoiceLineItems) SetItemSeller(v string)`

SetItemSeller sets ItemSeller field to given value.

### HasItemSeller

`func (o *InvoiceLineItems) HasItemSeller() bool`

HasItemSeller returns a boolean if a field has been set.

### SetItemSellerNil

`func (o *InvoiceLineItems) SetItemSellerNil(b bool)`

 SetItemSellerNil sets the value for ItemSeller to be an explicit nil

### UnsetItemSeller
`func (o *InvoiceLineItems) UnsetItemSeller()`

UnsetItemSeller ensures that no value is present for ItemSeller, not even an explicit nil
### GetItemAction

`func (o *InvoiceLineItems) GetItemAction() string`

GetItemAction returns the ItemAction field if non-nil, zero value otherwise.

### GetItemActionOk

`func (o *InvoiceLineItems) GetItemActionOk() (*string, bool)`

GetItemActionOk returns a tuple with the ItemAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAction

`func (o *InvoiceLineItems) SetItemAction(v string)`

SetItemAction sets ItemAction field to given value.

### HasItemAction

`func (o *InvoiceLineItems) HasItemAction() bool`

HasItemAction returns a boolean if a field has been set.

### SetItemActionNil

`func (o *InvoiceLineItems) SetItemActionNil(b bool)`

 SetItemActionNil sets the value for ItemAction to be an explicit nil

### UnsetItemAction
`func (o *InvoiceLineItems) UnsetItemAction()`

UnsetItemAction ensures that no value is present for ItemAction, not even an explicit nil
### GetExternalId

`func (o *InvoiceLineItems) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InvoiceLineItems) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InvoiceLineItems) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InvoiceLineItems) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRateId

`func (o *InvoiceLineItems) GetRateId() string`

GetRateId returns the RateId field if non-nil, zero value otherwise.

### GetRateIdOk

`func (o *InvoiceLineItems) GetRateIdOk() (*string, bool)`

GetRateIdOk returns a tuple with the RateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateId

`func (o *InvoiceLineItems) SetRateId(v string)`

SetRateId sets RateId field to given value.

### HasRateId

`func (o *InvoiceLineItems) HasRateId() bool`

HasRateId returns a boolean if a field has been set.

### GetRateClass

`func (o *InvoiceLineItems) GetRateClass() string`

GetRateClass returns the RateClass field if non-nil, zero value otherwise.

### GetRateClassOk

`func (o *InvoiceLineItems) GetRateClassOk() (*string, bool)`

GetRateClassOk returns a tuple with the RateClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateClass

`func (o *InvoiceLineItems) SetRateClass(v string)`

SetRateClass sets RateClass field to given value.

### HasRateClass

`func (o *InvoiceLineItems) HasRateClass() bool`

HasRateClass returns a boolean if a field has been set.

### SetRateClassNil

`func (o *InvoiceLineItems) SetRateClassNil(b bool)`

 SetRateClassNil sets the value for RateClass to be an explicit nil

### UnsetRateClass
`func (o *InvoiceLineItems) UnsetRateClass()`

UnsetRateClass ensures that no value is present for RateClass, not even an explicit nil
### GetRateUnit

`func (o *InvoiceLineItems) GetRateUnit() string`

GetRateUnit returns the RateUnit field if non-nil, zero value otherwise.

### GetRateUnitOk

`func (o *InvoiceLineItems) GetRateUnitOk() (*string, bool)`

GetRateUnitOk returns a tuple with the RateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateUnit

`func (o *InvoiceLineItems) SetRateUnit(v string)`

SetRateUnit sets RateUnit field to given value.

### HasRateUnit

`func (o *InvoiceLineItems) HasRateUnit() bool`

HasRateUnit returns a boolean if a field has been set.

### GetRateTerm

`func (o *InvoiceLineItems) GetRateTerm() string`

GetRateTerm returns the RateTerm field if non-nil, zero value otherwise.

### GetRateTermOk

`func (o *InvoiceLineItems) GetRateTermOk() (*string, bool)`

GetRateTermOk returns a tuple with the RateTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateTerm

`func (o *InvoiceLineItems) SetRateTerm(v string)`

SetRateTerm sets RateTerm field to given value.

### HasRateTerm

`func (o *InvoiceLineItems) HasRateTerm() bool`

HasRateTerm returns a boolean if a field has been set.

### SetRateTermNil

`func (o *InvoiceLineItems) SetRateTermNil(b bool)`

 SetRateTermNil sets the value for RateTerm to be an explicit nil

### UnsetRateTerm
`func (o *InvoiceLineItems) UnsetRateTerm()`

UnsetRateTerm ensures that no value is present for RateTerm, not even an explicit nil
### GetUsageType

`func (o *InvoiceLineItems) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *InvoiceLineItems) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *InvoiceLineItems) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *InvoiceLineItems) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetUsageCategory

`func (o *InvoiceLineItems) GetUsageCategory() string`

GetUsageCategory returns the UsageCategory field if non-nil, zero value otherwise.

### GetUsageCategoryOk

`func (o *InvoiceLineItems) GetUsageCategoryOk() (*string, bool)`

GetUsageCategoryOk returns a tuple with the UsageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCategory

`func (o *InvoiceLineItems) SetUsageCategory(v string)`

SetUsageCategory sets UsageCategory field to given value.

### HasUsageCategory

`func (o *InvoiceLineItems) HasUsageCategory() bool`

HasUsageCategory returns a boolean if a field has been set.

### GetUsageService

`func (o *InvoiceLineItems) GetUsageService() string`

GetUsageService returns the UsageService field if non-nil, zero value otherwise.

### GetUsageServiceOk

`func (o *InvoiceLineItems) GetUsageServiceOk() (*string, bool)`

GetUsageServiceOk returns a tuple with the UsageService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageService

`func (o *InvoiceLineItems) SetUsageService(v string)`

SetUsageService sets UsageService field to given value.

### HasUsageService

`func (o *InvoiceLineItems) HasUsageService() bool`

HasUsageService returns a boolean if a field has been set.

### GetItemUsage

`func (o *InvoiceLineItems) GetItemUsage() float32`

GetItemUsage returns the ItemUsage field if non-nil, zero value otherwise.

### GetItemUsageOk

`func (o *InvoiceLineItems) GetItemUsageOk() (*float32, bool)`

GetItemUsageOk returns a tuple with the ItemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUsage

`func (o *InvoiceLineItems) SetItemUsage(v float32)`

SetItemUsage sets ItemUsage field to given value.

### HasItemUsage

`func (o *InvoiceLineItems) HasItemUsage() bool`

HasItemUsage returns a boolean if a field has been set.

### GetItemRate

`func (o *InvoiceLineItems) GetItemRate() float32`

GetItemRate returns the ItemRate field if non-nil, zero value otherwise.

### GetItemRateOk

`func (o *InvoiceLineItems) GetItemRateOk() (*float32, bool)`

GetItemRateOk returns a tuple with the ItemRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRate

`func (o *InvoiceLineItems) SetItemRate(v float32)`

SetItemRate sets ItemRate field to given value.

### HasItemRate

`func (o *InvoiceLineItems) HasItemRate() bool`

HasItemRate returns a boolean if a field has been set.

### GetItemCost

`func (o *InvoiceLineItems) GetItemCost() float32`

GetItemCost returns the ItemCost field if non-nil, zero value otherwise.

### GetItemCostOk

`func (o *InvoiceLineItems) GetItemCostOk() (*float32, bool)`

GetItemCostOk returns a tuple with the ItemCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCost

`func (o *InvoiceLineItems) SetItemCost(v float32)`

SetItemCost sets ItemCost field to given value.

### HasItemCost

`func (o *InvoiceLineItems) HasItemCost() bool`

HasItemCost returns a boolean if a field has been set.

### GetItemPrice

`func (o *InvoiceLineItems) GetItemPrice() float32`

GetItemPrice returns the ItemPrice field if non-nil, zero value otherwise.

### GetItemPriceOk

`func (o *InvoiceLineItems) GetItemPriceOk() (*float32, bool)`

GetItemPriceOk returns a tuple with the ItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrice

`func (o *InvoiceLineItems) SetItemPrice(v float32)`

SetItemPrice sets ItemPrice field to given value.

### HasItemPrice

`func (o *InvoiceLineItems) HasItemPrice() bool`

HasItemPrice returns a boolean if a field has been set.

### GetItemTax

`func (o *InvoiceLineItems) GetItemTax() int64`

GetItemTax returns the ItemTax field if non-nil, zero value otherwise.

### GetItemTaxOk

`func (o *InvoiceLineItems) GetItemTaxOk() (*int64, bool)`

GetItemTaxOk returns a tuple with the ItemTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTax

`func (o *InvoiceLineItems) SetItemTax(v int64)`

SetItemTax sets ItemTax field to given value.

### HasItemTax

`func (o *InvoiceLineItems) HasItemTax() bool`

HasItemTax returns a boolean if a field has been set.

### GetItemTerm

`func (o *InvoiceLineItems) GetItemTerm() string`

GetItemTerm returns the ItemTerm field if non-nil, zero value otherwise.

### GetItemTermOk

`func (o *InvoiceLineItems) GetItemTermOk() (*string, bool)`

GetItemTermOk returns a tuple with the ItemTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTerm

`func (o *InvoiceLineItems) SetItemTerm(v string)`

SetItemTerm sets ItemTerm field to given value.

### HasItemTerm

`func (o *InvoiceLineItems) HasItemTerm() bool`

HasItemTerm returns a boolean if a field has been set.

### SetItemTermNil

`func (o *InvoiceLineItems) SetItemTermNil(b bool)`

 SetItemTermNil sets the value for ItemTerm to be an explicit nil

### UnsetItemTerm
`func (o *InvoiceLineItems) UnsetItemTerm()`

UnsetItemTerm ensures that no value is present for ItemTerm, not even an explicit nil
### GetTaxType

`func (o *InvoiceLineItems) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *InvoiceLineItems) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *InvoiceLineItems) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *InvoiceLineItems) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### SetTaxTypeNil

`func (o *InvoiceLineItems) SetTaxTypeNil(b bool)`

 SetTaxTypeNil sets the value for TaxType to be an explicit nil

### UnsetTaxType
`func (o *InvoiceLineItems) UnsetTaxType()`

UnsetTaxType ensures that no value is present for TaxType, not even an explicit nil
### GetRegionCode

`func (o *InvoiceLineItems) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *InvoiceLineItems) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *InvoiceLineItems) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *InvoiceLineItems) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceLineItems) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceLineItems) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceLineItems) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceLineItems) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetConversionRate

`func (o *InvoiceLineItems) GetConversionRate() int64`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *InvoiceLineItems) GetConversionRateOk() (*int64, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *InvoiceLineItems) SetConversionRate(v int64)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *InvoiceLineItems) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetDateCreated

`func (o *InvoiceLineItems) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InvoiceLineItems) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InvoiceLineItems) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InvoiceLineItems) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InvoiceLineItems) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InvoiceLineItems) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InvoiceLineItems) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InvoiceLineItems) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


