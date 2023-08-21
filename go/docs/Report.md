# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**ReportType**](report_type.md) |  | [optional] 
**ReportTitle** | Pointer to **NullableString** |  | [optional] 
**FilterTitle** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Config** | Pointer to [**ReportConfig**](report_config.md) |  | [optional] 
**CreatedBy** | Pointer to [**InlineResponse200107NetworkPoolCreatedBy**](inline_response_200_107_networkPool_createdBy.md) |  | [optional] 
**Rows** | Pointer to [**[]ReportRows**](ReportRows.md) |  | [optional] 

## Methods

### NewReport

`func NewReport() *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Report) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Report) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Report) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Report) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Report) GetType() ReportType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Report) GetTypeOk() (*ReportType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Report) SetType(v ReportType)`

SetType sets Type field to given value.

### HasType

`func (o *Report) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReportTitle

`func (o *Report) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *Report) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *Report) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.

### HasReportTitle

`func (o *Report) HasReportTitle() bool`

HasReportTitle returns a boolean if a field has been set.

### SetReportTitleNil

`func (o *Report) SetReportTitleNil(b bool)`

 SetReportTitleNil sets the value for ReportTitle to be an explicit nil

### UnsetReportTitle
`func (o *Report) UnsetReportTitle()`

UnsetReportTitle ensures that no value is present for ReportTitle, not even an explicit nil
### GetFilterTitle

`func (o *Report) GetFilterTitle() string`

GetFilterTitle returns the FilterTitle field if non-nil, zero value otherwise.

### GetFilterTitleOk

`func (o *Report) GetFilterTitleOk() (*string, bool)`

GetFilterTitleOk returns a tuple with the FilterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTitle

`func (o *Report) SetFilterTitle(v string)`

SetFilterTitle sets FilterTitle field to given value.

### HasFilterTitle

`func (o *Report) HasFilterTitle() bool`

HasFilterTitle returns a boolean if a field has been set.

### SetFilterTitleNil

`func (o *Report) SetFilterTitleNil(b bool)`

 SetFilterTitleNil sets the value for FilterTitle to be an explicit nil

### UnsetFilterTitle
`func (o *Report) UnsetFilterTitle()`

UnsetFilterTitle ensures that no value is present for FilterTitle, not even an explicit nil
### GetStatus

`func (o *Report) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Report) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Report) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Report) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *Report) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Report) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Report) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Report) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Report) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Report) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Report) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Report) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStartDate

`func (o *Report) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Report) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Report) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Report) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *Report) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Report) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *Report) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Report) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Report) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Report) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *Report) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Report) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetConfig

`func (o *Report) GetConfig() ReportConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Report) GetConfigOk() (*ReportConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Report) SetConfig(v ReportConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Report) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Report) GetCreatedBy() InlineResponse200107NetworkPoolCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Report) GetCreatedByOk() (*InlineResponse200107NetworkPoolCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Report) SetCreatedBy(v InlineResponse200107NetworkPoolCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Report) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetRows

`func (o *Report) GetRows() []ReportRows`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *Report) GetRowsOk() (*[]ReportRows, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *Report) SetRows(v []ReportRows)`

SetRows sets Rows field to given value.

### HasRows

`func (o *Report) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


