# ApiRolesIdRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to **string** | Authority (Name) | [optional] 
**Description** | Pointer to **NullableString** | Description | [optional] 
**DefaultPersona** | Pointer to **NullableString** | Set the default persona by code. | [optional] 
**Permissions** | Pointer to [**[]ApiRolesRolePermissions**](ApiRolesRolePermissions.md) | Set the access level for the specified permissions. | [optional] 
**GlobalSiteAccess** | Pointer to **string** | Set the default access level for for groups (sites). Only applies to user roles. | [optional] 
**Sites** | Pointer to [**[]ApiRolesRoleSites**](ApiRolesRoleSites.md) | Set the access level for the specified groups (sites). Only applies to user roles. | [optional] 
**GlobalZoneAccess** | Pointer to **string** | Set the default access level for for clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**Zones** | Pointer to [**[]ApiRolesRoleZones**](ApiRolesRoleZones.md) | Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**GlobalInstanceTypeAccess** | Pointer to **string** | Set the default access level for for instance types | [optional] 
**InstanceTypes** | Pointer to [**[]ApiRolesRoleInstanceTypes**](ApiRolesRoleInstanceTypes.md) | Set the access level for the specified instance types | [optional] 
**GlobalAppTemplateAccess** | Pointer to **string** | Set the default access level for blueprints | [optional] 
**AppTemplates** | Pointer to [**[]ApiRolesRoleAppTemplates**](ApiRolesRoleAppTemplates.md) | Set the access level for the specified blueprints (appTemplates) | [optional] 
**GlobalCatalogItemTypeAccess** | Pointer to **string** | Set the default access level for catalog item types | [optional] 
**CatalogItemTypes** | Pointer to [**[]ApiRolesRoleCatalogItemTypes**](ApiRolesRoleCatalogItemTypes.md) | Set the access level for the specified catalog item types | [optional] 
**GlobalPersonaAccess** | Pointer to **string** | Set the default access level for personas | [optional] 
**Personas** | Pointer to [**[]ApiRolesRolePersonas**](ApiRolesRolePersonas.md) | Set the access level for the specified personas | [optional] 
**GlobalVdiPoolAccess** | Pointer to **string** | Set the default access level for VDI pools | [optional] 
**VdiPools** | Pointer to [**[]ApiRolesRoleVdiPools**](ApiRolesRoleVdiPools.md) | Set the access level for the specified VDI pools | [optional] 
**GlobalReportTypeAccess** | Pointer to **string** | Set the default access level for report types | [optional] 
**ReportTypes** | Pointer to [**[]ApiRolesRoleReportTypes**](ApiRolesRoleReportTypes.md) | Set the access level for the specified report types | [optional] 
**GlobalTaskAccess** | Pointer to **string** | Set the default access level for tasks | [optional] 
**Tasks** | Pointer to [**[]ApiRolesRoleTasks**](ApiRolesRoleTasks.md) | Set the access level for the specified tasks | [optional] 
**GlobalTaskSetAccess** | Pointer to **string** | Set the default access level for workflows (taskSets) | [optional] 
**TaskSets** | Pointer to [**[]ApiRolesRoleTaskSets**](ApiRolesRoleTaskSets.md) | Set the access level for the specified workflows (taskSets) | [optional] 

## Methods

### NewApiRolesIdRole

`func NewApiRolesIdRole() *ApiRolesIdRole`

NewApiRolesIdRole instantiates a new ApiRolesIdRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRolesIdRoleWithDefaults

`func NewApiRolesIdRoleWithDefaults() *ApiRolesIdRole`

NewApiRolesIdRoleWithDefaults instantiates a new ApiRolesIdRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *ApiRolesIdRole) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *ApiRolesIdRole) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *ApiRolesIdRole) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *ApiRolesIdRole) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetDescription

`func (o *ApiRolesIdRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiRolesIdRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiRolesIdRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiRolesIdRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiRolesIdRole) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiRolesIdRole) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDefaultPersona

`func (o *ApiRolesIdRole) GetDefaultPersona() string`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *ApiRolesIdRole) GetDefaultPersonaOk() (*string, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *ApiRolesIdRole) SetDefaultPersona(v string)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *ApiRolesIdRole) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### SetDefaultPersonaNil

`func (o *ApiRolesIdRole) SetDefaultPersonaNil(b bool)`

 SetDefaultPersonaNil sets the value for DefaultPersona to be an explicit nil

### UnsetDefaultPersona
`func (o *ApiRolesIdRole) UnsetDefaultPersona()`

UnsetDefaultPersona ensures that no value is present for DefaultPersona, not even an explicit nil
### GetPermissions

`func (o *ApiRolesIdRole) GetPermissions() []ApiRolesRolePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiRolesIdRole) GetPermissionsOk() (*[]ApiRolesRolePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiRolesIdRole) SetPermissions(v []ApiRolesRolePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ApiRolesIdRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetGlobalSiteAccess

`func (o *ApiRolesIdRole) GetGlobalSiteAccess() string`

GetGlobalSiteAccess returns the GlobalSiteAccess field if non-nil, zero value otherwise.

### GetGlobalSiteAccessOk

`func (o *ApiRolesIdRole) GetGlobalSiteAccessOk() (*string, bool)`

GetGlobalSiteAccessOk returns a tuple with the GlobalSiteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSiteAccess

`func (o *ApiRolesIdRole) SetGlobalSiteAccess(v string)`

SetGlobalSiteAccess sets GlobalSiteAccess field to given value.

### HasGlobalSiteAccess

`func (o *ApiRolesIdRole) HasGlobalSiteAccess() bool`

HasGlobalSiteAccess returns a boolean if a field has been set.

### GetSites

`func (o *ApiRolesIdRole) GetSites() []ApiRolesRoleSites`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ApiRolesIdRole) GetSitesOk() (*[]ApiRolesRoleSites, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ApiRolesIdRole) SetSites(v []ApiRolesRoleSites)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ApiRolesIdRole) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetGlobalZoneAccess

`func (o *ApiRolesIdRole) GetGlobalZoneAccess() string`

GetGlobalZoneAccess returns the GlobalZoneAccess field if non-nil, zero value otherwise.

### GetGlobalZoneAccessOk

`func (o *ApiRolesIdRole) GetGlobalZoneAccessOk() (*string, bool)`

GetGlobalZoneAccessOk returns a tuple with the GlobalZoneAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalZoneAccess

`func (o *ApiRolesIdRole) SetGlobalZoneAccess(v string)`

SetGlobalZoneAccess sets GlobalZoneAccess field to given value.

### HasGlobalZoneAccess

`func (o *ApiRolesIdRole) HasGlobalZoneAccess() bool`

HasGlobalZoneAccess returns a boolean if a field has been set.

### GetZones

`func (o *ApiRolesIdRole) GetZones() []ApiRolesRoleZones`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ApiRolesIdRole) GetZonesOk() (*[]ApiRolesRoleZones, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ApiRolesIdRole) SetZones(v []ApiRolesRoleZones)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ApiRolesIdRole) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetGlobalInstanceTypeAccess

`func (o *ApiRolesIdRole) GetGlobalInstanceTypeAccess() string`

GetGlobalInstanceTypeAccess returns the GlobalInstanceTypeAccess field if non-nil, zero value otherwise.

### GetGlobalInstanceTypeAccessOk

`func (o *ApiRolesIdRole) GetGlobalInstanceTypeAccessOk() (*string, bool)`

GetGlobalInstanceTypeAccessOk returns a tuple with the GlobalInstanceTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalInstanceTypeAccess

`func (o *ApiRolesIdRole) SetGlobalInstanceTypeAccess(v string)`

SetGlobalInstanceTypeAccess sets GlobalInstanceTypeAccess field to given value.

### HasGlobalInstanceTypeAccess

`func (o *ApiRolesIdRole) HasGlobalInstanceTypeAccess() bool`

HasGlobalInstanceTypeAccess returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *ApiRolesIdRole) GetInstanceTypes() []ApiRolesRoleInstanceTypes`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *ApiRolesIdRole) GetInstanceTypesOk() (*[]ApiRolesRoleInstanceTypes, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *ApiRolesIdRole) SetInstanceTypes(v []ApiRolesRoleInstanceTypes)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *ApiRolesIdRole) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetGlobalAppTemplateAccess

`func (o *ApiRolesIdRole) GetGlobalAppTemplateAccess() string`

GetGlobalAppTemplateAccess returns the GlobalAppTemplateAccess field if non-nil, zero value otherwise.

### GetGlobalAppTemplateAccessOk

`func (o *ApiRolesIdRole) GetGlobalAppTemplateAccessOk() (*string, bool)`

GetGlobalAppTemplateAccessOk returns a tuple with the GlobalAppTemplateAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAppTemplateAccess

`func (o *ApiRolesIdRole) SetGlobalAppTemplateAccess(v string)`

SetGlobalAppTemplateAccess sets GlobalAppTemplateAccess field to given value.

### HasGlobalAppTemplateAccess

`func (o *ApiRolesIdRole) HasGlobalAppTemplateAccess() bool`

HasGlobalAppTemplateAccess returns a boolean if a field has been set.

### GetAppTemplates

`func (o *ApiRolesIdRole) GetAppTemplates() []ApiRolesRoleAppTemplates`

GetAppTemplates returns the AppTemplates field if non-nil, zero value otherwise.

### GetAppTemplatesOk

`func (o *ApiRolesIdRole) GetAppTemplatesOk() (*[]ApiRolesRoleAppTemplates, bool)`

GetAppTemplatesOk returns a tuple with the AppTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplates

`func (o *ApiRolesIdRole) SetAppTemplates(v []ApiRolesRoleAppTemplates)`

SetAppTemplates sets AppTemplates field to given value.

### HasAppTemplates

`func (o *ApiRolesIdRole) HasAppTemplates() bool`

HasAppTemplates returns a boolean if a field has been set.

### GetGlobalCatalogItemTypeAccess

`func (o *ApiRolesIdRole) GetGlobalCatalogItemTypeAccess() string`

GetGlobalCatalogItemTypeAccess returns the GlobalCatalogItemTypeAccess field if non-nil, zero value otherwise.

### GetGlobalCatalogItemTypeAccessOk

`func (o *ApiRolesIdRole) GetGlobalCatalogItemTypeAccessOk() (*string, bool)`

GetGlobalCatalogItemTypeAccessOk returns a tuple with the GlobalCatalogItemTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCatalogItemTypeAccess

`func (o *ApiRolesIdRole) SetGlobalCatalogItemTypeAccess(v string)`

SetGlobalCatalogItemTypeAccess sets GlobalCatalogItemTypeAccess field to given value.

### HasGlobalCatalogItemTypeAccess

`func (o *ApiRolesIdRole) HasGlobalCatalogItemTypeAccess() bool`

HasGlobalCatalogItemTypeAccess returns a boolean if a field has been set.

### GetCatalogItemTypes

`func (o *ApiRolesIdRole) GetCatalogItemTypes() []ApiRolesRoleCatalogItemTypes`

GetCatalogItemTypes returns the CatalogItemTypes field if non-nil, zero value otherwise.

### GetCatalogItemTypesOk

`func (o *ApiRolesIdRole) GetCatalogItemTypesOk() (*[]ApiRolesRoleCatalogItemTypes, bool)`

GetCatalogItemTypesOk returns a tuple with the CatalogItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypes

`func (o *ApiRolesIdRole) SetCatalogItemTypes(v []ApiRolesRoleCatalogItemTypes)`

SetCatalogItemTypes sets CatalogItemTypes field to given value.

### HasCatalogItemTypes

`func (o *ApiRolesIdRole) HasCatalogItemTypes() bool`

HasCatalogItemTypes returns a boolean if a field has been set.

### GetGlobalPersonaAccess

`func (o *ApiRolesIdRole) GetGlobalPersonaAccess() string`

GetGlobalPersonaAccess returns the GlobalPersonaAccess field if non-nil, zero value otherwise.

### GetGlobalPersonaAccessOk

`func (o *ApiRolesIdRole) GetGlobalPersonaAccessOk() (*string, bool)`

GetGlobalPersonaAccessOk returns a tuple with the GlobalPersonaAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPersonaAccess

`func (o *ApiRolesIdRole) SetGlobalPersonaAccess(v string)`

SetGlobalPersonaAccess sets GlobalPersonaAccess field to given value.

### HasGlobalPersonaAccess

`func (o *ApiRolesIdRole) HasGlobalPersonaAccess() bool`

HasGlobalPersonaAccess returns a boolean if a field has been set.

### GetPersonas

`func (o *ApiRolesIdRole) GetPersonas() []ApiRolesRolePersonas`

GetPersonas returns the Personas field if non-nil, zero value otherwise.

### GetPersonasOk

`func (o *ApiRolesIdRole) GetPersonasOk() (*[]ApiRolesRolePersonas, bool)`

GetPersonasOk returns a tuple with the Personas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonas

`func (o *ApiRolesIdRole) SetPersonas(v []ApiRolesRolePersonas)`

SetPersonas sets Personas field to given value.

### HasPersonas

`func (o *ApiRolesIdRole) HasPersonas() bool`

HasPersonas returns a boolean if a field has been set.

### GetGlobalVdiPoolAccess

`func (o *ApiRolesIdRole) GetGlobalVdiPoolAccess() string`

GetGlobalVdiPoolAccess returns the GlobalVdiPoolAccess field if non-nil, zero value otherwise.

### GetGlobalVdiPoolAccessOk

`func (o *ApiRolesIdRole) GetGlobalVdiPoolAccessOk() (*string, bool)`

GetGlobalVdiPoolAccessOk returns a tuple with the GlobalVdiPoolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalVdiPoolAccess

`func (o *ApiRolesIdRole) SetGlobalVdiPoolAccess(v string)`

SetGlobalVdiPoolAccess sets GlobalVdiPoolAccess field to given value.

### HasGlobalVdiPoolAccess

`func (o *ApiRolesIdRole) HasGlobalVdiPoolAccess() bool`

HasGlobalVdiPoolAccess returns a boolean if a field has been set.

### GetVdiPools

`func (o *ApiRolesIdRole) GetVdiPools() []ApiRolesRoleVdiPools`

GetVdiPools returns the VdiPools field if non-nil, zero value otherwise.

### GetVdiPoolsOk

`func (o *ApiRolesIdRole) GetVdiPoolsOk() (*[]ApiRolesRoleVdiPools, bool)`

GetVdiPoolsOk returns a tuple with the VdiPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPools

`func (o *ApiRolesIdRole) SetVdiPools(v []ApiRolesRoleVdiPools)`

SetVdiPools sets VdiPools field to given value.

### HasVdiPools

`func (o *ApiRolesIdRole) HasVdiPools() bool`

HasVdiPools returns a boolean if a field has been set.

### GetGlobalReportTypeAccess

`func (o *ApiRolesIdRole) GetGlobalReportTypeAccess() string`

GetGlobalReportTypeAccess returns the GlobalReportTypeAccess field if non-nil, zero value otherwise.

### GetGlobalReportTypeAccessOk

`func (o *ApiRolesIdRole) GetGlobalReportTypeAccessOk() (*string, bool)`

GetGlobalReportTypeAccessOk returns a tuple with the GlobalReportTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalReportTypeAccess

`func (o *ApiRolesIdRole) SetGlobalReportTypeAccess(v string)`

SetGlobalReportTypeAccess sets GlobalReportTypeAccess field to given value.

### HasGlobalReportTypeAccess

`func (o *ApiRolesIdRole) HasGlobalReportTypeAccess() bool`

HasGlobalReportTypeAccess returns a boolean if a field has been set.

### GetReportTypes

`func (o *ApiRolesIdRole) GetReportTypes() []ApiRolesRoleReportTypes`

GetReportTypes returns the ReportTypes field if non-nil, zero value otherwise.

### GetReportTypesOk

`func (o *ApiRolesIdRole) GetReportTypesOk() (*[]ApiRolesRoleReportTypes, bool)`

GetReportTypesOk returns a tuple with the ReportTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypes

`func (o *ApiRolesIdRole) SetReportTypes(v []ApiRolesRoleReportTypes)`

SetReportTypes sets ReportTypes field to given value.

### HasReportTypes

`func (o *ApiRolesIdRole) HasReportTypes() bool`

HasReportTypes returns a boolean if a field has been set.

### GetGlobalTaskAccess

`func (o *ApiRolesIdRole) GetGlobalTaskAccess() string`

GetGlobalTaskAccess returns the GlobalTaskAccess field if non-nil, zero value otherwise.

### GetGlobalTaskAccessOk

`func (o *ApiRolesIdRole) GetGlobalTaskAccessOk() (*string, bool)`

GetGlobalTaskAccessOk returns a tuple with the GlobalTaskAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskAccess

`func (o *ApiRolesIdRole) SetGlobalTaskAccess(v string)`

SetGlobalTaskAccess sets GlobalTaskAccess field to given value.

### HasGlobalTaskAccess

`func (o *ApiRolesIdRole) HasGlobalTaskAccess() bool`

HasGlobalTaskAccess returns a boolean if a field has been set.

### GetTasks

`func (o *ApiRolesIdRole) GetTasks() []ApiRolesRoleTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ApiRolesIdRole) GetTasksOk() (*[]ApiRolesRoleTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ApiRolesIdRole) SetTasks(v []ApiRolesRoleTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ApiRolesIdRole) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetGlobalTaskSetAccess

`func (o *ApiRolesIdRole) GetGlobalTaskSetAccess() string`

GetGlobalTaskSetAccess returns the GlobalTaskSetAccess field if non-nil, zero value otherwise.

### GetGlobalTaskSetAccessOk

`func (o *ApiRolesIdRole) GetGlobalTaskSetAccessOk() (*string, bool)`

GetGlobalTaskSetAccessOk returns a tuple with the GlobalTaskSetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskSetAccess

`func (o *ApiRolesIdRole) SetGlobalTaskSetAccess(v string)`

SetGlobalTaskSetAccess sets GlobalTaskSetAccess field to given value.

### HasGlobalTaskSetAccess

`func (o *ApiRolesIdRole) HasGlobalTaskSetAccess() bool`

HasGlobalTaskSetAccess returns a boolean if a field has been set.

### GetTaskSets

`func (o *ApiRolesIdRole) GetTaskSets() []ApiRolesRoleTaskSets`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *ApiRolesIdRole) GetTaskSetsOk() (*[]ApiRolesRoleTaskSets, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *ApiRolesIdRole) SetTaskSets(v []ApiRolesRoleTaskSets)`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *ApiRolesIdRole) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


