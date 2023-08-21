# LogData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeCode** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Ts** | Pointer to **time.Time** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**LogSignature** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SignatureVerified** | Pointer to **bool** |  | [optional] 

## Methods

### NewLogData

`func NewLogData() *LogData`

NewLogData instantiates a new LogData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogDataWithDefaults

`func NewLogDataWithDefaults() *LogData`

NewLogDataWithDefaults instantiates a new LogData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeCode

`func (o *LogData) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *LogData) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *LogData) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *LogData) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetMessage

`func (o *LogData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLevel

`func (o *LogData) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LogData) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LogData) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *LogData) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetTs

`func (o *LogData) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *LogData) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *LogData) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *LogData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetSourceType

`func (o *LogData) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *LogData) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *LogData) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *LogData) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTitle

`func (o *LogData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LogData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LogData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LogData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLogSignature

`func (o *LogData) GetLogSignature() string`

GetLogSignature returns the LogSignature field if non-nil, zero value otherwise.

### GetLogSignatureOk

`func (o *LogData) GetLogSignatureOk() (*string, bool)`

GetLogSignatureOk returns a tuple with the LogSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSignature

`func (o *LogData) SetLogSignature(v string)`

SetLogSignature sets LogSignature field to given value.

### HasLogSignature

`func (o *LogData) HasLogSignature() bool`

HasLogSignature returns a boolean if a field has been set.

### GetObjectId

`func (o *LogData) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *LogData) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *LogData) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *LogData) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetSeq

`func (o *LogData) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *LogData) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *LogData) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *LogData) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetId

`func (o *LogData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSignatureVerified

`func (o *LogData) GetSignatureVerified() bool`

GetSignatureVerified returns the SignatureVerified field if non-nil, zero value otherwise.

### GetSignatureVerifiedOk

`func (o *LogData) GetSignatureVerifiedOk() (*bool, bool)`

GetSignatureVerifiedOk returns a tuple with the SignatureVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVerified

`func (o *LogData) SetSignatureVerified(v bool)`

SetSignatureVerified sets SignatureVerified field to given value.

### HasSignatureVerified

`func (o *LogData) HasSignatureVerified() bool`

HasSignatureVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


