# ListBudgets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Budgets** | Pointer to [**[]Budgets**](Budgets.md) |  | [optional] 

## Methods

### NewListBudgets200Response

`func NewListBudgets200Response() *ListBudgets200Response`

NewListBudgets200Response instantiates a new ListBudgets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBudgets200ResponseWithDefaults

`func NewListBudgets200ResponseWithDefaults() *ListBudgets200Response`

NewListBudgets200ResponseWithDefaults instantiates a new ListBudgets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListBudgets200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBudgets200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBudgets200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBudgets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetBudgets

`func (o *ListBudgets200Response) GetBudgets() []Budgets`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *ListBudgets200Response) GetBudgetsOk() (*[]Budgets, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgets

`func (o *ListBudgets200Response) SetBudgets(v []Budgets)`

SetBudgets sets Budgets field to given value.

### HasBudgets

`func (o *ListBudgets200Response) HasBudgets() bool`

HasBudgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


