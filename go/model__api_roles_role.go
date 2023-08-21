/*
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * API version: 6.2.1
 * Contact: dev@morpheusdata.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApiRolesRole Payload for creating a new role
type ApiRolesRole struct {
	// Authority (Name)
	Authority string `json:"authority"`
	// Description
	Description NullableString `json:"description,omitempty"`
	// Role type
	RoleType *string `json:"roleType,omitempty"`
	// Base Role ID. Create the new role with the same permissions and access levels that the specified base role has. If this is not passed, the role is create without any permissions.
	BaseRoleId *int64 `json:"baseRoleId,omitempty"`
	// Multitenant roles are copied to all tenant accounts and kept in sync until a sub-tenant user modifies their copy of the role. *Only available to master tenant*
	Multitenant *bool `json:"multitenant,omitempty"`
	// Multitenant Locked, prevents sub-tenant users from modifying their copy of multienant roles. *Only available to master tenant*
	MultitenantLocked *bool `json:"multitenantLocked,omitempty"`
	DefaultPersona NullableString `json:"defaultPersona,omitempty"`
	// Set the access level for the specified permissions.
	Permissions *[]ApiRolesRolePermissions `json:"permissions,omitempty"`
	// Set the default access level for for groups (sites). Only applies to user roles.
	GlobalSiteAccess *string `json:"globalSiteAccess,omitempty"`
	// Set the access level for the specified groups (sites). Only applies to user roles.
	Sites *[]ApiRolesRoleSites `json:"sites,omitempty"`
	// Set the default access level for for clouds (zones). Only applies to base account (tenant) roles.
	GlobalZoneAccess *string `json:"globalZoneAccess,omitempty"`
	// Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles.
	Zones *[]ApiRolesRoleZones `json:"zones,omitempty"`
	// Set the default access level for for instance types
	GlobalInstanceTypeAccess *string `json:"globalInstanceTypeAccess,omitempty"`
	// Set the access level for the specified instance types
	InstanceTypes *[]ApiRolesRoleInstanceTypes `json:"instanceTypes,omitempty"`
	// Set the default access level for blueprints
	GlobalAppTemplateAccess *string `json:"globalAppTemplateAccess,omitempty"`
	// Set the access level for the specified blueprints (appTemplates)
	AppTemplates *[]ApiRolesRoleAppTemplates `json:"appTemplates,omitempty"`
	// Set the default access level for catalog item types
	GlobalCatalogItemTypeAccess *string `json:"globalCatalogItemTypeAccess,omitempty"`
	// Set the access level for the specified catalog item types
	CatalogItemTypes *[]ApiRolesRoleCatalogItemTypes `json:"catalogItemTypes,omitempty"`
	// Set the default access level for personas
	GlobalPersonaAccess *string `json:"globalPersonaAccess,omitempty"`
	// Set the access level for the specified personas
	Personas *[]ApiRolesRolePersonas `json:"personas,omitempty"`
	// Set the default access level for VDI pools
	GlobalVdiPoolAccess *string `json:"globalVdiPoolAccess,omitempty"`
	// Set the access level for the specified VDI pools
	VdiPools *[]ApiRolesRoleVdiPools `json:"vdiPools,omitempty"`
	// Set the default access level for report types
	GlobalReportTypeAccess *string `json:"globalReportTypeAccess,omitempty"`
	// Set the access level for the specified report types
	ReportTypes *[]ApiRolesRoleReportTypes `json:"reportTypes,omitempty"`
	// Set the default access level for tasks
	GlobalTaskAccess *string `json:"globalTaskAccess,omitempty"`
	// Set the access level for the specified tasks
	Tasks *[]ApiRolesRoleTasks `json:"tasks,omitempty"`
	// Set the default access level for workflows (taskSets)
	GlobalTaskSetAccess *string `json:"globalTaskSetAccess,omitempty"`
	// Set the access level for the specified workflows (taskSets)
	TaskSets *[]ApiRolesRoleTaskSets `json:"taskSets,omitempty"`
	// Set the role owner (tenant) by ID. *Only available to master tenant*
	Owner *int64 `json:"owner,omitempty"`
}

// NewApiRolesRole instantiates a new ApiRolesRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRolesRole(authority string, ) *ApiRolesRole {
	this := ApiRolesRole{}
	this.Authority = authority
	var roleType string = "user"
	this.RoleType = &roleType
	var multitenant bool = false
	this.Multitenant = &multitenant
	var multitenantLocked bool = false
	this.MultitenantLocked = &multitenantLocked
	return &this
}

// NewApiRolesRoleWithDefaults instantiates a new ApiRolesRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRolesRoleWithDefaults() *ApiRolesRole {
	this := ApiRolesRole{}
	var roleType string = "user"
	this.RoleType = &roleType
	var multitenant bool = false
	this.Multitenant = &multitenant
	var multitenantLocked bool = false
	this.MultitenantLocked = &multitenantLocked
	return &this
}

// GetAuthority returns the Authority field value
func (o *ApiRolesRole) GetAuthority() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetAuthorityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Authority, true
}

// SetAuthority sets field value
func (o *ApiRolesRole) SetAuthority(v string) {
	o.Authority = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiRolesRole) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiRolesRole) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiRolesRole) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ApiRolesRole) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ApiRolesRole) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ApiRolesRole) UnsetDescription() {
	o.Description.Unset()
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *ApiRolesRole) GetRoleType() string {
	if o == nil || o.RoleType == nil {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetRoleTypeOk() (*string, bool) {
	if o == nil || o.RoleType == nil {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *ApiRolesRole) HasRoleType() bool {
	if o != nil && o.RoleType != nil {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *ApiRolesRole) SetRoleType(v string) {
	o.RoleType = &v
}

// GetBaseRoleId returns the BaseRoleId field value if set, zero value otherwise.
func (o *ApiRolesRole) GetBaseRoleId() int64 {
	if o == nil || o.BaseRoleId == nil {
		var ret int64
		return ret
	}
	return *o.BaseRoleId
}

// GetBaseRoleIdOk returns a tuple with the BaseRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetBaseRoleIdOk() (*int64, bool) {
	if o == nil || o.BaseRoleId == nil {
		return nil, false
	}
	return o.BaseRoleId, true
}

// HasBaseRoleId returns a boolean if a field has been set.
func (o *ApiRolesRole) HasBaseRoleId() bool {
	if o != nil && o.BaseRoleId != nil {
		return true
	}

	return false
}

// SetBaseRoleId gets a reference to the given int64 and assigns it to the BaseRoleId field.
func (o *ApiRolesRole) SetBaseRoleId(v int64) {
	o.BaseRoleId = &v
}

// GetMultitenant returns the Multitenant field value if set, zero value otherwise.
func (o *ApiRolesRole) GetMultitenant() bool {
	if o == nil || o.Multitenant == nil {
		var ret bool
		return ret
	}
	return *o.Multitenant
}

// GetMultitenantOk returns a tuple with the Multitenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetMultitenantOk() (*bool, bool) {
	if o == nil || o.Multitenant == nil {
		return nil, false
	}
	return o.Multitenant, true
}

// HasMultitenant returns a boolean if a field has been set.
func (o *ApiRolesRole) HasMultitenant() bool {
	if o != nil && o.Multitenant != nil {
		return true
	}

	return false
}

// SetMultitenant gets a reference to the given bool and assigns it to the Multitenant field.
func (o *ApiRolesRole) SetMultitenant(v bool) {
	o.Multitenant = &v
}

// GetMultitenantLocked returns the MultitenantLocked field value if set, zero value otherwise.
func (o *ApiRolesRole) GetMultitenantLocked() bool {
	if o == nil || o.MultitenantLocked == nil {
		var ret bool
		return ret
	}
	return *o.MultitenantLocked
}

// GetMultitenantLockedOk returns a tuple with the MultitenantLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetMultitenantLockedOk() (*bool, bool) {
	if o == nil || o.MultitenantLocked == nil {
		return nil, false
	}
	return o.MultitenantLocked, true
}

// HasMultitenantLocked returns a boolean if a field has been set.
func (o *ApiRolesRole) HasMultitenantLocked() bool {
	if o != nil && o.MultitenantLocked != nil {
		return true
	}

	return false
}

// SetMultitenantLocked gets a reference to the given bool and assigns it to the MultitenantLocked field.
func (o *ApiRolesRole) SetMultitenantLocked(v bool) {
	o.MultitenantLocked = &v
}

// GetDefaultPersona returns the DefaultPersona field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiRolesRole) GetDefaultPersona() string {
	if o == nil || o.DefaultPersona.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultPersona.Get()
}

// GetDefaultPersonaOk returns a tuple with the DefaultPersona field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiRolesRole) GetDefaultPersonaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultPersona.Get(), o.DefaultPersona.IsSet()
}

// HasDefaultPersona returns a boolean if a field has been set.
func (o *ApiRolesRole) HasDefaultPersona() bool {
	if o != nil && o.DefaultPersona.IsSet() {
		return true
	}

	return false
}

// SetDefaultPersona gets a reference to the given NullableString and assigns it to the DefaultPersona field.
func (o *ApiRolesRole) SetDefaultPersona(v string) {
	o.DefaultPersona.Set(&v)
}
// SetDefaultPersonaNil sets the value for DefaultPersona to be an explicit nil
func (o *ApiRolesRole) SetDefaultPersonaNil() {
	o.DefaultPersona.Set(nil)
}

// UnsetDefaultPersona ensures that no value is present for DefaultPersona, not even an explicit nil
func (o *ApiRolesRole) UnsetDefaultPersona() {
	o.DefaultPersona.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ApiRolesRole) GetPermissions() []ApiRolesRolePermissions {
	if o == nil || o.Permissions == nil {
		var ret []ApiRolesRolePermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetPermissionsOk() (*[]ApiRolesRolePermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ApiRolesRole) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []ApiRolesRolePermissions and assigns it to the Permissions field.
func (o *ApiRolesRole) SetPermissions(v []ApiRolesRolePermissions) {
	o.Permissions = &v
}

// GetGlobalSiteAccess returns the GlobalSiteAccess field value if set, zero value otherwise.
func (o *ApiRolesRole) GetGlobalSiteAccess() string {
	if o == nil || o.GlobalSiteAccess == nil {
		var ret string
		return ret
	}
	return *o.GlobalSiteAccess
}

// GetGlobalSiteAccessOk returns a tuple with the GlobalSiteAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetGlobalSiteAccessOk() (*string, bool) {
	if o == nil || o.GlobalSiteAccess == nil {
		return nil, false
	}
	return o.GlobalSiteAccess, true
}

// HasGlobalSiteAccess returns a boolean if a field has been set.
func (o *ApiRolesRole) HasGlobalSiteAccess() bool {
	if o != nil && o.GlobalSiteAccess != nil {
		return true
	}

	return false
}

// SetGlobalSiteAccess gets a reference to the given string and assigns it to the GlobalSiteAccess field.
func (o *ApiRolesRole) SetGlobalSiteAccess(v string) {
	o.GlobalSiteAccess = &v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *ApiRolesRole) GetSites() []ApiRolesRoleSites {
	if o == nil || o.Sites == nil {
		var ret []ApiRolesRoleSites
		return ret
	}
	return *o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetSitesOk() (*[]ApiRolesRoleSites, bool) {
	if o == nil || o.Sites == nil {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *ApiRolesRole) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given []ApiRolesRoleSites and assigns it to the Sites field.
func (o *ApiRolesRole) SetSites(v []ApiRolesRoleSites) {
	o.Sites = &v
}

// GetGlobalZoneAccess returns the GlobalZoneAccess field value if set, zero value otherwise.
func (o *ApiRolesRole) GetGlobalZoneAccess() string {
	if o == nil || o.GlobalZoneAccess == nil {
		var ret string
		return ret
	}
	return *o.GlobalZoneAccess
}

// GetGlobalZoneAccessOk returns a tuple with the GlobalZoneAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetGlobalZoneAccessOk() (*string, bool) {
	if o == nil || o.GlobalZoneAccess == nil {
		return nil, false
	}
	return o.GlobalZoneAccess, true
}

// HasGlobalZoneAccess returns a boolean if a field has been set.
func (o *ApiRolesRole) HasGlobalZoneAccess() bool {
	if o != nil && o.GlobalZoneAccess != nil {
		return true
	}

	return false
}

// SetGlobalZoneAccess gets a reference to the given string and assigns it to the GlobalZoneAccess field.
func (o *ApiRolesRole) SetGlobalZoneAccess(v string) {
	o.GlobalZoneAccess = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *ApiRolesRole) GetZones() []ApiRolesRoleZones {
	if o == nil || o.Zones == nil {
		var ret []ApiRolesRoleZones
		return ret
	}
	return *o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetZonesOk() (*[]ApiRolesRoleZones, bool) {
	if o == nil || o.Zones == nil {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *ApiRolesRole) HasZones() bool {
	if o != nil && o.Zones != nil {
		return true
	}

	return false
}

// SetZones gets a reference to the given []ApiRolesRoleZones and assigns it to the Zones field.
func (o *ApiRolesRole) SetZones(v []ApiRolesRoleZones) {
	o.Zones = &v
}

// GetGlobalInstanceTypeAccess returns the GlobalInstanceTypeAccess field value if set, zero value otherwise.
func (o *ApiRolesRole) GetGlobalInstanceTypeAccess() string {
	if o == nil || o.GlobalInstanceTypeAccess == nil {
		var ret string
		return ret
	}
	return *o.GlobalInstanceTypeAccess
}

// GetGlobalInstanceTypeAccessOk returns a tuple with the GlobalInstanceTypeAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetGlobalInstanceTypeAccessOk() (*string, bool) {
	if o == nil || o.GlobalInstanceTypeAccess == nil {
		return nil, false
	}
	return o.GlobalInstanceTypeAccess, true
}

// HasGlobalInstanceTypeAccess returns a boolean if a field has been set.
func (o *ApiRolesRole) HasGlobalInstanceTypeAccess() bool {
	if o != nil && o.GlobalInstanceTypeAccess != nil {
		return true
	}

	return false
}

// SetGlobalInstanceTypeAccess gets a reference to the given string and assigns it to the GlobalInstanceTypeAccess field.
func (o *ApiRolesRole) SetGlobalInstanceTypeAccess(v string) {
	o.GlobalInstanceTypeAccess = &v
}

// GetInstanceTypes returns the InstanceTypes field value if set, zero value otherwise.
func (o *ApiRolesRole) GetInstanceTypes() []ApiRolesRoleInstanceTypes {
	if o == nil || o.InstanceTypes == nil {
		var ret []ApiRolesRoleInstanceTypes
		return ret
	}
	return *o.InstanceTypes
}

// GetInstanceTypesOk returns a tuple with the InstanceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetInstanceTypesOk() (*[]ApiRolesRoleInstanceTypes, bool) {
	if o == nil || o.InstanceTypes == nil {
		return nil, false
	}
	return o.InstanceTypes, true
}

// HasInstanceTypes returns a boolean if a field has been set.
func (o *ApiRolesRole) HasInstanceTypes() bool {
	if o != nil && o.InstanceTypes != nil {
		return true
	}

	return false
}

// SetInstanceTypes gets a reference to the given []ApiRolesRoleInstanceTypes and assigns it to the InstanceTypes field.
func (o *ApiRolesRole) SetInstanceTypes(v []ApiRolesRoleInstanceTypes) {
	o.InstanceTypes = &v
}

// GetGlobalAppTemplateAccess returns the GlobalAppTemplateAccess field value if set, zero value otherwise.
func (o *ApiRolesRole) GetGlobalAppTemplateAccess() string {
	if o == nil || o.GlobalAppTemplateAccess == nil {
		var ret string
		return ret
	}
	return *o.GlobalAppTemplateAccess
}

// GetGlobalAppTemplateAccessOk returns a tuple with the GlobalAppTemplateAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetGlobalAppTemplateAccessOk() (*string, bool) {
	if o == nil || o.GlobalAppTemplateAccess == nil {
		return nil, false
	}
	return o.GlobalAppTemplateAccess, true
}

// HasGlobalAppTemplateAccess returns a boolean if a field has been set.
func (o *ApiRolesRole) HasGlobalAppTemplateAccess() bool {
	if o != nil && o.GlobalAppTemplateAccess != nil {
		return true
	}

	return false
}

// SetGlobalAppTemplateAccess gets a reference to the given string and assigns it to the GlobalAppTemplateAccess field.
func (o *ApiRolesRole) SetGlobalAppTemplateAccess(v string) {
	o.GlobalAppTemplateAccess = &v
}

// GetAppTemplates returns the AppTemplates field value if set, zero value otherwise.
func (o *ApiRolesRole) GetAppTemplates() []ApiRolesRoleAppTemplates {
	if o == nil || o.AppTemplates == nil {
		var ret []ApiRolesRoleAppTemplates
		return ret
	}
	return *o.AppTemplates
}

// GetAppTemplatesOk returns a tuple with the AppTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetAppTemplatesOk() (*[]ApiRolesRoleAppTemplates, bool) {
	if o == nil || o.AppTemplates == nil {
		return nil, false
	}
	return o.AppTemplates, true
}

// HasAppTemplates returns a boolean if a field has been set.
func (o *ApiRolesRole) HasAppTemplates() bool {
	if o != nil && o.AppTemplates != nil {
		return true
	}

	return false
}

// SetAppTemplates gets a reference to the given []ApiRolesRoleAppTemplates and assigns it to the AppTemplates field.
func (o *ApiRolesRole) SetAppTemplates(v []ApiRolesRoleAppTemplates) {
	o.AppTemplates = &v
}

// GetGlobalCatalogItemTypeAccess returns the GlobalCatalogItemTypeAccess field value if set, zero value otherwise.
func (o *ApiRolesRole) GetGlobalCatalogItemTypeAccess() string {
	if o == nil || o.GlobalCatalogItemTypeAccess == nil {
		var ret string
		return ret
	}
	return *o.GlobalCatalogItemTypeAccess
}

// GetGlobalCatalogItemTypeAccessOk returns a tuple with the GlobalCatalogItemTypeAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetGlobalCatalogItemTypeAccessOk() (*string, bool) {
	if o == nil || o.GlobalCatalogItemTypeAccess == nil {
		return nil, false
	}
	return o.GlobalCatalogItemTypeAccess, true
}

// HasGlobalCatalogItemTypeAccess returns a boolean if a field has been set.
func (o *ApiRolesRole) HasGlobalCatalogItemTypeAccess() bool {
	if o != nil && o.GlobalCatalogItemTypeAccess != nil {
		return true
	}

	return false
}

// SetGlobalCatalogItemTypeAccess gets a reference to the given string and assigns it to the GlobalCatalogItemTypeAccess field.
func (o *ApiRolesRole) SetGlobalCatalogItemTypeAccess(v string) {
	o.GlobalCatalogItemTypeAccess = &v
}

// GetCatalogItemTypes returns the CatalogItemTypes field value if set, zero value otherwise.
func (o *ApiRolesRole) GetCatalogItemTypes() []ApiRolesRoleCatalogItemTypes {
	if o == nil || o.CatalogItemTypes == nil {
		var ret []ApiRolesRoleCatalogItemTypes
		return ret
	}
	return *o.CatalogItemTypes
}

// GetCatalogItemTypesOk returns a tuple with the CatalogItemTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetCatalogItemTypesOk() (*[]ApiRolesRoleCatalogItemTypes, bool) {
	if o == nil || o.CatalogItemTypes == nil {
		return nil, false
	}
	return o.CatalogItemTypes, true
}

// HasCatalogItemTypes returns a boolean if a field has been set.
func (o *ApiRolesRole) HasCatalogItemTypes() bool {
	if o != nil && o.CatalogItemTypes != nil {
		return true
	}

	return false
}

// SetCatalogItemTypes gets a reference to the given []ApiRolesRoleCatalogItemTypes and assigns it to the CatalogItemTypes field.
func (o *ApiRolesRole) SetCatalogItemTypes(v []ApiRolesRoleCatalogItemTypes) {
	o.CatalogItemTypes = &v
}

// GetGlobalPersonaAccess returns the GlobalPersonaAccess field value if set, zero value otherwise.
func (o *ApiRolesRole) GetGlobalPersonaAccess() string {
	if o == nil || o.GlobalPersonaAccess == nil {
		var ret string
		return ret
	}
	return *o.GlobalPersonaAccess
}

// GetGlobalPersonaAccessOk returns a tuple with the GlobalPersonaAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetGlobalPersonaAccessOk() (*string, bool) {
	if o == nil || o.GlobalPersonaAccess == nil {
		return nil, false
	}
	return o.GlobalPersonaAccess, true
}

// HasGlobalPersonaAccess returns a boolean if a field has been set.
func (o *ApiRolesRole) HasGlobalPersonaAccess() bool {
	if o != nil && o.GlobalPersonaAccess != nil {
		return true
	}

	return false
}

// SetGlobalPersonaAccess gets a reference to the given string and assigns it to the GlobalPersonaAccess field.
func (o *ApiRolesRole) SetGlobalPersonaAccess(v string) {
	o.GlobalPersonaAccess = &v
}

// GetPersonas returns the Personas field value if set, zero value otherwise.
func (o *ApiRolesRole) GetPersonas() []ApiRolesRolePersonas {
	if o == nil || o.Personas == nil {
		var ret []ApiRolesRolePersonas
		return ret
	}
	return *o.Personas
}

// GetPersonasOk returns a tuple with the Personas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetPersonasOk() (*[]ApiRolesRolePersonas, bool) {
	if o == nil || o.Personas == nil {
		return nil, false
	}
	return o.Personas, true
}

// HasPersonas returns a boolean if a field has been set.
func (o *ApiRolesRole) HasPersonas() bool {
	if o != nil && o.Personas != nil {
		return true
	}

	return false
}

// SetPersonas gets a reference to the given []ApiRolesRolePersonas and assigns it to the Personas field.
func (o *ApiRolesRole) SetPersonas(v []ApiRolesRolePersonas) {
	o.Personas = &v
}

// GetGlobalVdiPoolAccess returns the GlobalVdiPoolAccess field value if set, zero value otherwise.
func (o *ApiRolesRole) GetGlobalVdiPoolAccess() string {
	if o == nil || o.GlobalVdiPoolAccess == nil {
		var ret string
		return ret
	}
	return *o.GlobalVdiPoolAccess
}

// GetGlobalVdiPoolAccessOk returns a tuple with the GlobalVdiPoolAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetGlobalVdiPoolAccessOk() (*string, bool) {
	if o == nil || o.GlobalVdiPoolAccess == nil {
		return nil, false
	}
	return o.GlobalVdiPoolAccess, true
}

// HasGlobalVdiPoolAccess returns a boolean if a field has been set.
func (o *ApiRolesRole) HasGlobalVdiPoolAccess() bool {
	if o != nil && o.GlobalVdiPoolAccess != nil {
		return true
	}

	return false
}

// SetGlobalVdiPoolAccess gets a reference to the given string and assigns it to the GlobalVdiPoolAccess field.
func (o *ApiRolesRole) SetGlobalVdiPoolAccess(v string) {
	o.GlobalVdiPoolAccess = &v
}

// GetVdiPools returns the VdiPools field value if set, zero value otherwise.
func (o *ApiRolesRole) GetVdiPools() []ApiRolesRoleVdiPools {
	if o == nil || o.VdiPools == nil {
		var ret []ApiRolesRoleVdiPools
		return ret
	}
	return *o.VdiPools
}

// GetVdiPoolsOk returns a tuple with the VdiPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetVdiPoolsOk() (*[]ApiRolesRoleVdiPools, bool) {
	if o == nil || o.VdiPools == nil {
		return nil, false
	}
	return o.VdiPools, true
}

// HasVdiPools returns a boolean if a field has been set.
func (o *ApiRolesRole) HasVdiPools() bool {
	if o != nil && o.VdiPools != nil {
		return true
	}

	return false
}

// SetVdiPools gets a reference to the given []ApiRolesRoleVdiPools and assigns it to the VdiPools field.
func (o *ApiRolesRole) SetVdiPools(v []ApiRolesRoleVdiPools) {
	o.VdiPools = &v
}

// GetGlobalReportTypeAccess returns the GlobalReportTypeAccess field value if set, zero value otherwise.
func (o *ApiRolesRole) GetGlobalReportTypeAccess() string {
	if o == nil || o.GlobalReportTypeAccess == nil {
		var ret string
		return ret
	}
	return *o.GlobalReportTypeAccess
}

// GetGlobalReportTypeAccessOk returns a tuple with the GlobalReportTypeAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetGlobalReportTypeAccessOk() (*string, bool) {
	if o == nil || o.GlobalReportTypeAccess == nil {
		return nil, false
	}
	return o.GlobalReportTypeAccess, true
}

// HasGlobalReportTypeAccess returns a boolean if a field has been set.
func (o *ApiRolesRole) HasGlobalReportTypeAccess() bool {
	if o != nil && o.GlobalReportTypeAccess != nil {
		return true
	}

	return false
}

// SetGlobalReportTypeAccess gets a reference to the given string and assigns it to the GlobalReportTypeAccess field.
func (o *ApiRolesRole) SetGlobalReportTypeAccess(v string) {
	o.GlobalReportTypeAccess = &v
}

// GetReportTypes returns the ReportTypes field value if set, zero value otherwise.
func (o *ApiRolesRole) GetReportTypes() []ApiRolesRoleReportTypes {
	if o == nil || o.ReportTypes == nil {
		var ret []ApiRolesRoleReportTypes
		return ret
	}
	return *o.ReportTypes
}

// GetReportTypesOk returns a tuple with the ReportTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetReportTypesOk() (*[]ApiRolesRoleReportTypes, bool) {
	if o == nil || o.ReportTypes == nil {
		return nil, false
	}
	return o.ReportTypes, true
}

// HasReportTypes returns a boolean if a field has been set.
func (o *ApiRolesRole) HasReportTypes() bool {
	if o != nil && o.ReportTypes != nil {
		return true
	}

	return false
}

// SetReportTypes gets a reference to the given []ApiRolesRoleReportTypes and assigns it to the ReportTypes field.
func (o *ApiRolesRole) SetReportTypes(v []ApiRolesRoleReportTypes) {
	o.ReportTypes = &v
}

// GetGlobalTaskAccess returns the GlobalTaskAccess field value if set, zero value otherwise.
func (o *ApiRolesRole) GetGlobalTaskAccess() string {
	if o == nil || o.GlobalTaskAccess == nil {
		var ret string
		return ret
	}
	return *o.GlobalTaskAccess
}

// GetGlobalTaskAccessOk returns a tuple with the GlobalTaskAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetGlobalTaskAccessOk() (*string, bool) {
	if o == nil || o.GlobalTaskAccess == nil {
		return nil, false
	}
	return o.GlobalTaskAccess, true
}

// HasGlobalTaskAccess returns a boolean if a field has been set.
func (o *ApiRolesRole) HasGlobalTaskAccess() bool {
	if o != nil && o.GlobalTaskAccess != nil {
		return true
	}

	return false
}

// SetGlobalTaskAccess gets a reference to the given string and assigns it to the GlobalTaskAccess field.
func (o *ApiRolesRole) SetGlobalTaskAccess(v string) {
	o.GlobalTaskAccess = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *ApiRolesRole) GetTasks() []ApiRolesRoleTasks {
	if o == nil || o.Tasks == nil {
		var ret []ApiRolesRoleTasks
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetTasksOk() (*[]ApiRolesRoleTasks, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ApiRolesRole) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []ApiRolesRoleTasks and assigns it to the Tasks field.
func (o *ApiRolesRole) SetTasks(v []ApiRolesRoleTasks) {
	o.Tasks = &v
}

// GetGlobalTaskSetAccess returns the GlobalTaskSetAccess field value if set, zero value otherwise.
func (o *ApiRolesRole) GetGlobalTaskSetAccess() string {
	if o == nil || o.GlobalTaskSetAccess == nil {
		var ret string
		return ret
	}
	return *o.GlobalTaskSetAccess
}

// GetGlobalTaskSetAccessOk returns a tuple with the GlobalTaskSetAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetGlobalTaskSetAccessOk() (*string, bool) {
	if o == nil || o.GlobalTaskSetAccess == nil {
		return nil, false
	}
	return o.GlobalTaskSetAccess, true
}

// HasGlobalTaskSetAccess returns a boolean if a field has been set.
func (o *ApiRolesRole) HasGlobalTaskSetAccess() bool {
	if o != nil && o.GlobalTaskSetAccess != nil {
		return true
	}

	return false
}

// SetGlobalTaskSetAccess gets a reference to the given string and assigns it to the GlobalTaskSetAccess field.
func (o *ApiRolesRole) SetGlobalTaskSetAccess(v string) {
	o.GlobalTaskSetAccess = &v
}

// GetTaskSets returns the TaskSets field value if set, zero value otherwise.
func (o *ApiRolesRole) GetTaskSets() []ApiRolesRoleTaskSets {
	if o == nil || o.TaskSets == nil {
		var ret []ApiRolesRoleTaskSets
		return ret
	}
	return *o.TaskSets
}

// GetTaskSetsOk returns a tuple with the TaskSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetTaskSetsOk() (*[]ApiRolesRoleTaskSets, bool) {
	if o == nil || o.TaskSets == nil {
		return nil, false
	}
	return o.TaskSets, true
}

// HasTaskSets returns a boolean if a field has been set.
func (o *ApiRolesRole) HasTaskSets() bool {
	if o != nil && o.TaskSets != nil {
		return true
	}

	return false
}

// SetTaskSets gets a reference to the given []ApiRolesRoleTaskSets and assigns it to the TaskSets field.
func (o *ApiRolesRole) SetTaskSets(v []ApiRolesRoleTaskSets) {
	o.TaskSets = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ApiRolesRole) GetOwner() int64 {
	if o == nil || o.Owner == nil {
		var ret int64
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRole) GetOwnerOk() (*int64, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ApiRolesRole) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given int64 and assigns it to the Owner field.
func (o *ApiRolesRole) SetOwner(v int64) {
	o.Owner = &v
}

func (o ApiRolesRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authority"] = o.Authority
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.RoleType != nil {
		toSerialize["roleType"] = o.RoleType
	}
	if o.BaseRoleId != nil {
		toSerialize["baseRoleId"] = o.BaseRoleId
	}
	if o.Multitenant != nil {
		toSerialize["multitenant"] = o.Multitenant
	}
	if o.MultitenantLocked != nil {
		toSerialize["multitenantLocked"] = o.MultitenantLocked
	}
	if o.DefaultPersona.IsSet() {
		toSerialize["defaultPersona"] = o.DefaultPersona.Get()
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.GlobalSiteAccess != nil {
		toSerialize["globalSiteAccess"] = o.GlobalSiteAccess
	}
	if o.Sites != nil {
		toSerialize["sites"] = o.Sites
	}
	if o.GlobalZoneAccess != nil {
		toSerialize["globalZoneAccess"] = o.GlobalZoneAccess
	}
	if o.Zones != nil {
		toSerialize["zones"] = o.Zones
	}
	if o.GlobalInstanceTypeAccess != nil {
		toSerialize["globalInstanceTypeAccess"] = o.GlobalInstanceTypeAccess
	}
	if o.InstanceTypes != nil {
		toSerialize["instanceTypes"] = o.InstanceTypes
	}
	if o.GlobalAppTemplateAccess != nil {
		toSerialize["globalAppTemplateAccess"] = o.GlobalAppTemplateAccess
	}
	if o.AppTemplates != nil {
		toSerialize["appTemplates"] = o.AppTemplates
	}
	if o.GlobalCatalogItemTypeAccess != nil {
		toSerialize["globalCatalogItemTypeAccess"] = o.GlobalCatalogItemTypeAccess
	}
	if o.CatalogItemTypes != nil {
		toSerialize["catalogItemTypes"] = o.CatalogItemTypes
	}
	if o.GlobalPersonaAccess != nil {
		toSerialize["globalPersonaAccess"] = o.GlobalPersonaAccess
	}
	if o.Personas != nil {
		toSerialize["personas"] = o.Personas
	}
	if o.GlobalVdiPoolAccess != nil {
		toSerialize["globalVdiPoolAccess"] = o.GlobalVdiPoolAccess
	}
	if o.VdiPools != nil {
		toSerialize["vdiPools"] = o.VdiPools
	}
	if o.GlobalReportTypeAccess != nil {
		toSerialize["globalReportTypeAccess"] = o.GlobalReportTypeAccess
	}
	if o.ReportTypes != nil {
		toSerialize["reportTypes"] = o.ReportTypes
	}
	if o.GlobalTaskAccess != nil {
		toSerialize["globalTaskAccess"] = o.GlobalTaskAccess
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if o.GlobalTaskSetAccess != nil {
		toSerialize["globalTaskSetAccess"] = o.GlobalTaskSetAccess
	}
	if o.TaskSets != nil {
		toSerialize["taskSets"] = o.TaskSets
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullableApiRolesRole struct {
	value *ApiRolesRole
	isSet bool
}

func (v NullableApiRolesRole) Get() *ApiRolesRole {
	return v.value
}

func (v *NullableApiRolesRole) Set(val *ApiRolesRole) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRolesRole) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRolesRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRolesRole(val *ApiRolesRole) *NullableApiRolesRole {
	return &NullableApiRolesRole{value: val, isSet: true}
}

func (v NullableApiRolesRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRolesRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

