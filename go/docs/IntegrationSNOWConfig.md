# IntegrationSNOWConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncidentAccess** | Pointer to **bool** |  | [optional] 
**RequestAccess** | Pointer to **bool** |  | [optional] 
**ServiceNowCMDBBusinessObject** | Pointer to **string** |  | [optional] 
**ServiceNowCustomCmdbMapping** | Pointer to **string** |  | [optional] 
**ServiceNowCmdbClassMapping** | Pointer to [**[]IntegrationSNOWConfigServiceNowCmdbClassMapping**](IntegrationSNOWConfigServiceNowCmdbClassMapping.md) |  | [optional] 
**WebServiceImportUrl** | Pointer to **string** |  | [optional] 
**WebServiceImportSysId** | Pointer to **string** |  | [optional] 
**WebServiceOperationUrl** | Pointer to **string** |  | [optional] 
**PreparedForSync** | Pointer to **bool** |  | [optional] 

## Methods

### NewIntegrationSNOWConfig

`func NewIntegrationSNOWConfig() *IntegrationSNOWConfig`

NewIntegrationSNOWConfig instantiates a new IntegrationSNOWConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationSNOWConfigWithDefaults

`func NewIntegrationSNOWConfigWithDefaults() *IntegrationSNOWConfig`

NewIntegrationSNOWConfigWithDefaults instantiates a new IntegrationSNOWConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidentAccess

`func (o *IntegrationSNOWConfig) GetIncidentAccess() bool`

GetIncidentAccess returns the IncidentAccess field if non-nil, zero value otherwise.

### GetIncidentAccessOk

`func (o *IntegrationSNOWConfig) GetIncidentAccessOk() (*bool, bool)`

GetIncidentAccessOk returns a tuple with the IncidentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentAccess

`func (o *IntegrationSNOWConfig) SetIncidentAccess(v bool)`

SetIncidentAccess sets IncidentAccess field to given value.

### HasIncidentAccess

`func (o *IntegrationSNOWConfig) HasIncidentAccess() bool`

HasIncidentAccess returns a boolean if a field has been set.

### GetRequestAccess

`func (o *IntegrationSNOWConfig) GetRequestAccess() bool`

GetRequestAccess returns the RequestAccess field if non-nil, zero value otherwise.

### GetRequestAccessOk

`func (o *IntegrationSNOWConfig) GetRequestAccessOk() (*bool, bool)`

GetRequestAccessOk returns a tuple with the RequestAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAccess

`func (o *IntegrationSNOWConfig) SetRequestAccess(v bool)`

SetRequestAccess sets RequestAccess field to given value.

### HasRequestAccess

`func (o *IntegrationSNOWConfig) HasRequestAccess() bool`

HasRequestAccess returns a boolean if a field has been set.

### GetServiceNowCMDBBusinessObject

`func (o *IntegrationSNOWConfig) GetServiceNowCMDBBusinessObject() string`

GetServiceNowCMDBBusinessObject returns the ServiceNowCMDBBusinessObject field if non-nil, zero value otherwise.

### GetServiceNowCMDBBusinessObjectOk

`func (o *IntegrationSNOWConfig) GetServiceNowCMDBBusinessObjectOk() (*string, bool)`

GetServiceNowCMDBBusinessObjectOk returns a tuple with the ServiceNowCMDBBusinessObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowCMDBBusinessObject

`func (o *IntegrationSNOWConfig) SetServiceNowCMDBBusinessObject(v string)`

SetServiceNowCMDBBusinessObject sets ServiceNowCMDBBusinessObject field to given value.

### HasServiceNowCMDBBusinessObject

`func (o *IntegrationSNOWConfig) HasServiceNowCMDBBusinessObject() bool`

HasServiceNowCMDBBusinessObject returns a boolean if a field has been set.

### GetServiceNowCustomCmdbMapping

`func (o *IntegrationSNOWConfig) GetServiceNowCustomCmdbMapping() string`

GetServiceNowCustomCmdbMapping returns the ServiceNowCustomCmdbMapping field if non-nil, zero value otherwise.

### GetServiceNowCustomCmdbMappingOk

`func (o *IntegrationSNOWConfig) GetServiceNowCustomCmdbMappingOk() (*string, bool)`

GetServiceNowCustomCmdbMappingOk returns a tuple with the ServiceNowCustomCmdbMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowCustomCmdbMapping

`func (o *IntegrationSNOWConfig) SetServiceNowCustomCmdbMapping(v string)`

SetServiceNowCustomCmdbMapping sets ServiceNowCustomCmdbMapping field to given value.

### HasServiceNowCustomCmdbMapping

`func (o *IntegrationSNOWConfig) HasServiceNowCustomCmdbMapping() bool`

HasServiceNowCustomCmdbMapping returns a boolean if a field has been set.

### GetServiceNowCmdbClassMapping

`func (o *IntegrationSNOWConfig) GetServiceNowCmdbClassMapping() []IntegrationSNOWConfigServiceNowCmdbClassMapping`

GetServiceNowCmdbClassMapping returns the ServiceNowCmdbClassMapping field if non-nil, zero value otherwise.

### GetServiceNowCmdbClassMappingOk

`func (o *IntegrationSNOWConfig) GetServiceNowCmdbClassMappingOk() (*[]IntegrationSNOWConfigServiceNowCmdbClassMapping, bool)`

GetServiceNowCmdbClassMappingOk returns a tuple with the ServiceNowCmdbClassMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowCmdbClassMapping

`func (o *IntegrationSNOWConfig) SetServiceNowCmdbClassMapping(v []IntegrationSNOWConfigServiceNowCmdbClassMapping)`

SetServiceNowCmdbClassMapping sets ServiceNowCmdbClassMapping field to given value.

### HasServiceNowCmdbClassMapping

`func (o *IntegrationSNOWConfig) HasServiceNowCmdbClassMapping() bool`

HasServiceNowCmdbClassMapping returns a boolean if a field has been set.

### GetWebServiceImportUrl

`func (o *IntegrationSNOWConfig) GetWebServiceImportUrl() string`

GetWebServiceImportUrl returns the WebServiceImportUrl field if non-nil, zero value otherwise.

### GetWebServiceImportUrlOk

`func (o *IntegrationSNOWConfig) GetWebServiceImportUrlOk() (*string, bool)`

GetWebServiceImportUrlOk returns a tuple with the WebServiceImportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceImportUrl

`func (o *IntegrationSNOWConfig) SetWebServiceImportUrl(v string)`

SetWebServiceImportUrl sets WebServiceImportUrl field to given value.

### HasWebServiceImportUrl

`func (o *IntegrationSNOWConfig) HasWebServiceImportUrl() bool`

HasWebServiceImportUrl returns a boolean if a field has been set.

### GetWebServiceImportSysId

`func (o *IntegrationSNOWConfig) GetWebServiceImportSysId() string`

GetWebServiceImportSysId returns the WebServiceImportSysId field if non-nil, zero value otherwise.

### GetWebServiceImportSysIdOk

`func (o *IntegrationSNOWConfig) GetWebServiceImportSysIdOk() (*string, bool)`

GetWebServiceImportSysIdOk returns a tuple with the WebServiceImportSysId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceImportSysId

`func (o *IntegrationSNOWConfig) SetWebServiceImportSysId(v string)`

SetWebServiceImportSysId sets WebServiceImportSysId field to given value.

### HasWebServiceImportSysId

`func (o *IntegrationSNOWConfig) HasWebServiceImportSysId() bool`

HasWebServiceImportSysId returns a boolean if a field has been set.

### GetWebServiceOperationUrl

`func (o *IntegrationSNOWConfig) GetWebServiceOperationUrl() string`

GetWebServiceOperationUrl returns the WebServiceOperationUrl field if non-nil, zero value otherwise.

### GetWebServiceOperationUrlOk

`func (o *IntegrationSNOWConfig) GetWebServiceOperationUrlOk() (*string, bool)`

GetWebServiceOperationUrlOk returns a tuple with the WebServiceOperationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServiceOperationUrl

`func (o *IntegrationSNOWConfig) SetWebServiceOperationUrl(v string)`

SetWebServiceOperationUrl sets WebServiceOperationUrl field to given value.

### HasWebServiceOperationUrl

`func (o *IntegrationSNOWConfig) HasWebServiceOperationUrl() bool`

HasWebServiceOperationUrl returns a boolean if a field has been set.

### GetPreparedForSync

`func (o *IntegrationSNOWConfig) GetPreparedForSync() bool`

GetPreparedForSync returns the PreparedForSync field if non-nil, zero value otherwise.

### GetPreparedForSyncOk

`func (o *IntegrationSNOWConfig) GetPreparedForSyncOk() (*bool, bool)`

GetPreparedForSyncOk returns a tuple with the PreparedForSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparedForSync

`func (o *IntegrationSNOWConfig) SetPreparedForSync(v bool)`

SetPreparedForSync sets PreparedForSync field to given value.

### HasPreparedForSync

`func (o *IntegrationSNOWConfig) HasPreparedForSync() bool`

HasPreparedForSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


