# ApiRolesRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | **string** | Authority (Name) | 
**Description** | Pointer to **NullableString** | Description | [optional] 
**RoleType** | Pointer to **string** | Role type | [optional] [default to "user"]
**BaseRoleId** | Pointer to **int64** | Base Role ID. Create the new role with the same permissions and access levels that the specified base role has. If this is not passed, the role is create without any permissions. | [optional] 
**Multitenant** | Pointer to **bool** | Multitenant roles are copied to all tenant accounts and kept in sync until a sub-tenant user modifies their copy of the role. *Only available to master tenant* | [optional] [default to false]
**MultitenantLocked** | Pointer to **bool** | Multitenant Locked, prevents sub-tenant users from modifying their copy of multienant roles. *Only available to master tenant* | [optional] [default to false]
**DefaultPersona** | Pointer to **NullableString** |  | [optional] 
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
**Owner** | Pointer to **int64** | Set the role owner (tenant) by ID. *Only available to master tenant* | [optional] 

## Methods

### NewApiRolesRole

`func NewApiRolesRole(authority string, ) *ApiRolesRole`

NewApiRolesRole instantiates a new ApiRolesRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRolesRoleWithDefaults

`func NewApiRolesRoleWithDefaults() *ApiRolesRole`

NewApiRolesRoleWithDefaults instantiates a new ApiRolesRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *ApiRolesRole) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *ApiRolesRole) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *ApiRolesRole) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetDescription

`func (o *ApiRolesRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiRolesRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiRolesRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiRolesRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiRolesRole) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiRolesRole) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRoleType

`func (o *ApiRolesRole) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *ApiRolesRole) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *ApiRolesRole) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *ApiRolesRole) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetBaseRoleId

`func (o *ApiRolesRole) GetBaseRoleId() int64`

GetBaseRoleId returns the BaseRoleId field if non-nil, zero value otherwise.

### GetBaseRoleIdOk

`func (o *ApiRolesRole) GetBaseRoleIdOk() (*int64, bool)`

GetBaseRoleIdOk returns a tuple with the BaseRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRoleId

`func (o *ApiRolesRole) SetBaseRoleId(v int64)`

SetBaseRoleId sets BaseRoleId field to given value.

### HasBaseRoleId

`func (o *ApiRolesRole) HasBaseRoleId() bool`

HasBaseRoleId returns a boolean if a field has been set.

### GetMultitenant

`func (o *ApiRolesRole) GetMultitenant() bool`

GetMultitenant returns the Multitenant field if non-nil, zero value otherwise.

### GetMultitenantOk

`func (o *ApiRolesRole) GetMultitenantOk() (*bool, bool)`

GetMultitenantOk returns a tuple with the Multitenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenant

`func (o *ApiRolesRole) SetMultitenant(v bool)`

SetMultitenant sets Multitenant field to given value.

### HasMultitenant

`func (o *ApiRolesRole) HasMultitenant() bool`

HasMultitenant returns a boolean if a field has been set.

### GetMultitenantLocked

`func (o *ApiRolesRole) GetMultitenantLocked() bool`

GetMultitenantLocked returns the MultitenantLocked field if non-nil, zero value otherwise.

### GetMultitenantLockedOk

`func (o *ApiRolesRole) GetMultitenantLockedOk() (*bool, bool)`

GetMultitenantLockedOk returns a tuple with the MultitenantLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenantLocked

`func (o *ApiRolesRole) SetMultitenantLocked(v bool)`

SetMultitenantLocked sets MultitenantLocked field to given value.

### HasMultitenantLocked

`func (o *ApiRolesRole) HasMultitenantLocked() bool`

HasMultitenantLocked returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *ApiRolesRole) GetDefaultPersona() string`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *ApiRolesRole) GetDefaultPersonaOk() (*string, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *ApiRolesRole) SetDefaultPersona(v string)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *ApiRolesRole) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### SetDefaultPersonaNil

`func (o *ApiRolesRole) SetDefaultPersonaNil(b bool)`

 SetDefaultPersonaNil sets the value for DefaultPersona to be an explicit nil

### UnsetDefaultPersona
`func (o *ApiRolesRole) UnsetDefaultPersona()`

UnsetDefaultPersona ensures that no value is present for DefaultPersona, not even an explicit nil
### GetPermissions

`func (o *ApiRolesRole) GetPermissions() []ApiRolesRolePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiRolesRole) GetPermissionsOk() (*[]ApiRolesRolePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiRolesRole) SetPermissions(v []ApiRolesRolePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ApiRolesRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetGlobalSiteAccess

`func (o *ApiRolesRole) GetGlobalSiteAccess() string`

GetGlobalSiteAccess returns the GlobalSiteAccess field if non-nil, zero value otherwise.

### GetGlobalSiteAccessOk

`func (o *ApiRolesRole) GetGlobalSiteAccessOk() (*string, bool)`

GetGlobalSiteAccessOk returns a tuple with the GlobalSiteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSiteAccess

`func (o *ApiRolesRole) SetGlobalSiteAccess(v string)`

SetGlobalSiteAccess sets GlobalSiteAccess field to given value.

### HasGlobalSiteAccess

`func (o *ApiRolesRole) HasGlobalSiteAccess() bool`

HasGlobalSiteAccess returns a boolean if a field has been set.

### GetSites

`func (o *ApiRolesRole) GetSites() []ApiRolesRoleSites`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ApiRolesRole) GetSitesOk() (*[]ApiRolesRoleSites, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ApiRolesRole) SetSites(v []ApiRolesRoleSites)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ApiRolesRole) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetGlobalZoneAccess

`func (o *ApiRolesRole) GetGlobalZoneAccess() string`

GetGlobalZoneAccess returns the GlobalZoneAccess field if non-nil, zero value otherwise.

### GetGlobalZoneAccessOk

`func (o *ApiRolesRole) GetGlobalZoneAccessOk() (*string, bool)`

GetGlobalZoneAccessOk returns a tuple with the GlobalZoneAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalZoneAccess

`func (o *ApiRolesRole) SetGlobalZoneAccess(v string)`

SetGlobalZoneAccess sets GlobalZoneAccess field to given value.

### HasGlobalZoneAccess

`func (o *ApiRolesRole) HasGlobalZoneAccess() bool`

HasGlobalZoneAccess returns a boolean if a field has been set.

### GetZones

`func (o *ApiRolesRole) GetZones() []ApiRolesRoleZones`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ApiRolesRole) GetZonesOk() (*[]ApiRolesRoleZones, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ApiRolesRole) SetZones(v []ApiRolesRoleZones)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ApiRolesRole) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetGlobalInstanceTypeAccess

`func (o *ApiRolesRole) GetGlobalInstanceTypeAccess() string`

GetGlobalInstanceTypeAccess returns the GlobalInstanceTypeAccess field if non-nil, zero value otherwise.

### GetGlobalInstanceTypeAccessOk

`func (o *ApiRolesRole) GetGlobalInstanceTypeAccessOk() (*string, bool)`

GetGlobalInstanceTypeAccessOk returns a tuple with the GlobalInstanceTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalInstanceTypeAccess

`func (o *ApiRolesRole) SetGlobalInstanceTypeAccess(v string)`

SetGlobalInstanceTypeAccess sets GlobalInstanceTypeAccess field to given value.

### HasGlobalInstanceTypeAccess

`func (o *ApiRolesRole) HasGlobalInstanceTypeAccess() bool`

HasGlobalInstanceTypeAccess returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *ApiRolesRole) GetInstanceTypes() []ApiRolesRoleInstanceTypes`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *ApiRolesRole) GetInstanceTypesOk() (*[]ApiRolesRoleInstanceTypes, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *ApiRolesRole) SetInstanceTypes(v []ApiRolesRoleInstanceTypes)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *ApiRolesRole) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetGlobalAppTemplateAccess

`func (o *ApiRolesRole) GetGlobalAppTemplateAccess() string`

GetGlobalAppTemplateAccess returns the GlobalAppTemplateAccess field if non-nil, zero value otherwise.

### GetGlobalAppTemplateAccessOk

`func (o *ApiRolesRole) GetGlobalAppTemplateAccessOk() (*string, bool)`

GetGlobalAppTemplateAccessOk returns a tuple with the GlobalAppTemplateAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAppTemplateAccess

`func (o *ApiRolesRole) SetGlobalAppTemplateAccess(v string)`

SetGlobalAppTemplateAccess sets GlobalAppTemplateAccess field to given value.

### HasGlobalAppTemplateAccess

`func (o *ApiRolesRole) HasGlobalAppTemplateAccess() bool`

HasGlobalAppTemplateAccess returns a boolean if a field has been set.

### GetAppTemplates

`func (o *ApiRolesRole) GetAppTemplates() []ApiRolesRoleAppTemplates`

GetAppTemplates returns the AppTemplates field if non-nil, zero value otherwise.

### GetAppTemplatesOk

`func (o *ApiRolesRole) GetAppTemplatesOk() (*[]ApiRolesRoleAppTemplates, bool)`

GetAppTemplatesOk returns a tuple with the AppTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplates

`func (o *ApiRolesRole) SetAppTemplates(v []ApiRolesRoleAppTemplates)`

SetAppTemplates sets AppTemplates field to given value.

### HasAppTemplates

`func (o *ApiRolesRole) HasAppTemplates() bool`

HasAppTemplates returns a boolean if a field has been set.

### GetGlobalCatalogItemTypeAccess

`func (o *ApiRolesRole) GetGlobalCatalogItemTypeAccess() string`

GetGlobalCatalogItemTypeAccess returns the GlobalCatalogItemTypeAccess field if non-nil, zero value otherwise.

### GetGlobalCatalogItemTypeAccessOk

`func (o *ApiRolesRole) GetGlobalCatalogItemTypeAccessOk() (*string, bool)`

GetGlobalCatalogItemTypeAccessOk returns a tuple with the GlobalCatalogItemTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCatalogItemTypeAccess

`func (o *ApiRolesRole) SetGlobalCatalogItemTypeAccess(v string)`

SetGlobalCatalogItemTypeAccess sets GlobalCatalogItemTypeAccess field to given value.

### HasGlobalCatalogItemTypeAccess

`func (o *ApiRolesRole) HasGlobalCatalogItemTypeAccess() bool`

HasGlobalCatalogItemTypeAccess returns a boolean if a field has been set.

### GetCatalogItemTypes

`func (o *ApiRolesRole) GetCatalogItemTypes() []ApiRolesRoleCatalogItemTypes`

GetCatalogItemTypes returns the CatalogItemTypes field if non-nil, zero value otherwise.

### GetCatalogItemTypesOk

`func (o *ApiRolesRole) GetCatalogItemTypesOk() (*[]ApiRolesRoleCatalogItemTypes, bool)`

GetCatalogItemTypesOk returns a tuple with the CatalogItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypes

`func (o *ApiRolesRole) SetCatalogItemTypes(v []ApiRolesRoleCatalogItemTypes)`

SetCatalogItemTypes sets CatalogItemTypes field to given value.

### HasCatalogItemTypes

`func (o *ApiRolesRole) HasCatalogItemTypes() bool`

HasCatalogItemTypes returns a boolean if a field has been set.

### GetGlobalPersonaAccess

`func (o *ApiRolesRole) GetGlobalPersonaAccess() string`

GetGlobalPersonaAccess returns the GlobalPersonaAccess field if non-nil, zero value otherwise.

### GetGlobalPersonaAccessOk

`func (o *ApiRolesRole) GetGlobalPersonaAccessOk() (*string, bool)`

GetGlobalPersonaAccessOk returns a tuple with the GlobalPersonaAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPersonaAccess

`func (o *ApiRolesRole) SetGlobalPersonaAccess(v string)`

SetGlobalPersonaAccess sets GlobalPersonaAccess field to given value.

### HasGlobalPersonaAccess

`func (o *ApiRolesRole) HasGlobalPersonaAccess() bool`

HasGlobalPersonaAccess returns a boolean if a field has been set.

### GetPersonas

`func (o *ApiRolesRole) GetPersonas() []ApiRolesRolePersonas`

GetPersonas returns the Personas field if non-nil, zero value otherwise.

### GetPersonasOk

`func (o *ApiRolesRole) GetPersonasOk() (*[]ApiRolesRolePersonas, bool)`

GetPersonasOk returns a tuple with the Personas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonas

`func (o *ApiRolesRole) SetPersonas(v []ApiRolesRolePersonas)`

SetPersonas sets Personas field to given value.

### HasPersonas

`func (o *ApiRolesRole) HasPersonas() bool`

HasPersonas returns a boolean if a field has been set.

### GetGlobalVdiPoolAccess

`func (o *ApiRolesRole) GetGlobalVdiPoolAccess() string`

GetGlobalVdiPoolAccess returns the GlobalVdiPoolAccess field if non-nil, zero value otherwise.

### GetGlobalVdiPoolAccessOk

`func (o *ApiRolesRole) GetGlobalVdiPoolAccessOk() (*string, bool)`

GetGlobalVdiPoolAccessOk returns a tuple with the GlobalVdiPoolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalVdiPoolAccess

`func (o *ApiRolesRole) SetGlobalVdiPoolAccess(v string)`

SetGlobalVdiPoolAccess sets GlobalVdiPoolAccess field to given value.

### HasGlobalVdiPoolAccess

`func (o *ApiRolesRole) HasGlobalVdiPoolAccess() bool`

HasGlobalVdiPoolAccess returns a boolean if a field has been set.

### GetVdiPools

`func (o *ApiRolesRole) GetVdiPools() []ApiRolesRoleVdiPools`

GetVdiPools returns the VdiPools field if non-nil, zero value otherwise.

### GetVdiPoolsOk

`func (o *ApiRolesRole) GetVdiPoolsOk() (*[]ApiRolesRoleVdiPools, bool)`

GetVdiPoolsOk returns a tuple with the VdiPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPools

`func (o *ApiRolesRole) SetVdiPools(v []ApiRolesRoleVdiPools)`

SetVdiPools sets VdiPools field to given value.

### HasVdiPools

`func (o *ApiRolesRole) HasVdiPools() bool`

HasVdiPools returns a boolean if a field has been set.

### GetGlobalReportTypeAccess

`func (o *ApiRolesRole) GetGlobalReportTypeAccess() string`

GetGlobalReportTypeAccess returns the GlobalReportTypeAccess field if non-nil, zero value otherwise.

### GetGlobalReportTypeAccessOk

`func (o *ApiRolesRole) GetGlobalReportTypeAccessOk() (*string, bool)`

GetGlobalReportTypeAccessOk returns a tuple with the GlobalReportTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalReportTypeAccess

`func (o *ApiRolesRole) SetGlobalReportTypeAccess(v string)`

SetGlobalReportTypeAccess sets GlobalReportTypeAccess field to given value.

### HasGlobalReportTypeAccess

`func (o *ApiRolesRole) HasGlobalReportTypeAccess() bool`

HasGlobalReportTypeAccess returns a boolean if a field has been set.

### GetReportTypes

`func (o *ApiRolesRole) GetReportTypes() []ApiRolesRoleReportTypes`

GetReportTypes returns the ReportTypes field if non-nil, zero value otherwise.

### GetReportTypesOk

`func (o *ApiRolesRole) GetReportTypesOk() (*[]ApiRolesRoleReportTypes, bool)`

GetReportTypesOk returns a tuple with the ReportTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypes

`func (o *ApiRolesRole) SetReportTypes(v []ApiRolesRoleReportTypes)`

SetReportTypes sets ReportTypes field to given value.

### HasReportTypes

`func (o *ApiRolesRole) HasReportTypes() bool`

HasReportTypes returns a boolean if a field has been set.

### GetGlobalTaskAccess

`func (o *ApiRolesRole) GetGlobalTaskAccess() string`

GetGlobalTaskAccess returns the GlobalTaskAccess field if non-nil, zero value otherwise.

### GetGlobalTaskAccessOk

`func (o *ApiRolesRole) GetGlobalTaskAccessOk() (*string, bool)`

GetGlobalTaskAccessOk returns a tuple with the GlobalTaskAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskAccess

`func (o *ApiRolesRole) SetGlobalTaskAccess(v string)`

SetGlobalTaskAccess sets GlobalTaskAccess field to given value.

### HasGlobalTaskAccess

`func (o *ApiRolesRole) HasGlobalTaskAccess() bool`

HasGlobalTaskAccess returns a boolean if a field has been set.

### GetTasks

`func (o *ApiRolesRole) GetTasks() []ApiRolesRoleTasks`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ApiRolesRole) GetTasksOk() (*[]ApiRolesRoleTasks, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ApiRolesRole) SetTasks(v []ApiRolesRoleTasks)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ApiRolesRole) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetGlobalTaskSetAccess

`func (o *ApiRolesRole) GetGlobalTaskSetAccess() string`

GetGlobalTaskSetAccess returns the GlobalTaskSetAccess field if non-nil, zero value otherwise.

### GetGlobalTaskSetAccessOk

`func (o *ApiRolesRole) GetGlobalTaskSetAccessOk() (*string, bool)`

GetGlobalTaskSetAccessOk returns a tuple with the GlobalTaskSetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskSetAccess

`func (o *ApiRolesRole) SetGlobalTaskSetAccess(v string)`

SetGlobalTaskSetAccess sets GlobalTaskSetAccess field to given value.

### HasGlobalTaskSetAccess

`func (o *ApiRolesRole) HasGlobalTaskSetAccess() bool`

HasGlobalTaskSetAccess returns a boolean if a field has been set.

### GetTaskSets

`func (o *ApiRolesRole) GetTaskSets() []ApiRolesRoleTaskSets`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *ApiRolesRole) GetTaskSetsOk() (*[]ApiRolesRoleTaskSets, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *ApiRolesRole) SetTaskSets(v []ApiRolesRoleTaskSets)`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *ApiRolesRole) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### GetOwner

`func (o *ApiRolesRole) GetOwner() int64`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ApiRolesRole) GetOwnerOk() (*int64, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ApiRolesRole) SetOwner(v int64)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ApiRolesRole) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


