# ZoneTypeServerTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**NodeType** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Selectable** | Pointer to **bool** |  | [optional] 
**ExternalDelete** | Pointer to **bool** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**ControlPower** | Pointer to **bool** |  | [optional] 
**ControlSuspend** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**HasAgent** | Pointer to **bool** |  | [optional] 
**VmHypervisor** | Pointer to **bool** |  | [optional] 
**ContainerHypervisor** | Pointer to **bool** |  | [optional] 
**BareMetalHost** | Pointer to **bool** |  | [optional] 
**GuestVm** | Pointer to **bool** |  | [optional] 
**HasAutomation** | Pointer to **bool** |  | [optional] 
**ProvisionType** | Pointer to [**ZoneTypeProvisionType**](zoneType_provisionType.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]ZoneTypeOptionTypes**](ZoneTypeOptionTypes.md) |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 

## Methods

### NewZoneTypeServerTypes

`func NewZoneTypeServerTypes() *ZoneTypeServerTypes`

NewZoneTypeServerTypes instantiates a new ZoneTypeServerTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneTypeServerTypesWithDefaults

`func NewZoneTypeServerTypesWithDefaults() *ZoneTypeServerTypes`

NewZoneTypeServerTypesWithDefaults instantiates a new ZoneTypeServerTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneTypeServerTypes) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneTypeServerTypes) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneTypeServerTypes) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneTypeServerTypes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ZoneTypeServerTypes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ZoneTypeServerTypes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ZoneTypeServerTypes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ZoneTypeServerTypes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ZoneTypeServerTypes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneTypeServerTypes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneTypeServerTypes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneTypeServerTypes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ZoneTypeServerTypes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZoneTypeServerTypes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZoneTypeServerTypes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ZoneTypeServerTypes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNodeType

`func (o *ZoneTypeServerTypes) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ZoneTypeServerTypes) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ZoneTypeServerTypes) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ZoneTypeServerTypes) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetPlatform

`func (o *ZoneTypeServerTypes) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ZoneTypeServerTypes) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ZoneTypeServerTypes) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ZoneTypeServerTypes) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetEnabled

`func (o *ZoneTypeServerTypes) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ZoneTypeServerTypes) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ZoneTypeServerTypes) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ZoneTypeServerTypes) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSelectable

`func (o *ZoneTypeServerTypes) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *ZoneTypeServerTypes) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *ZoneTypeServerTypes) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *ZoneTypeServerTypes) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetExternalDelete

`func (o *ZoneTypeServerTypes) GetExternalDelete() bool`

GetExternalDelete returns the ExternalDelete field if non-nil, zero value otherwise.

### GetExternalDeleteOk

`func (o *ZoneTypeServerTypes) GetExternalDeleteOk() (*bool, bool)`

GetExternalDeleteOk returns a tuple with the ExternalDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDelete

`func (o *ZoneTypeServerTypes) SetExternalDelete(v bool)`

SetExternalDelete sets ExternalDelete field to given value.

### HasExternalDelete

`func (o *ZoneTypeServerTypes) HasExternalDelete() bool`

HasExternalDelete returns a boolean if a field has been set.

### GetManaged

`func (o *ZoneTypeServerTypes) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ZoneTypeServerTypes) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ZoneTypeServerTypes) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ZoneTypeServerTypes) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetControlPower

`func (o *ZoneTypeServerTypes) GetControlPower() bool`

GetControlPower returns the ControlPower field if non-nil, zero value otherwise.

### GetControlPowerOk

`func (o *ZoneTypeServerTypes) GetControlPowerOk() (*bool, bool)`

GetControlPowerOk returns a tuple with the ControlPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPower

`func (o *ZoneTypeServerTypes) SetControlPower(v bool)`

SetControlPower sets ControlPower field to given value.

### HasControlPower

`func (o *ZoneTypeServerTypes) HasControlPower() bool`

HasControlPower returns a boolean if a field has been set.

### GetControlSuspend

`func (o *ZoneTypeServerTypes) GetControlSuspend() bool`

GetControlSuspend returns the ControlSuspend field if non-nil, zero value otherwise.

### GetControlSuspendOk

`func (o *ZoneTypeServerTypes) GetControlSuspendOk() (*bool, bool)`

GetControlSuspendOk returns a tuple with the ControlSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlSuspend

`func (o *ZoneTypeServerTypes) SetControlSuspend(v bool)`

SetControlSuspend sets ControlSuspend field to given value.

### HasControlSuspend

`func (o *ZoneTypeServerTypes) HasControlSuspend() bool`

HasControlSuspend returns a boolean if a field has been set.

### GetCreatable

`func (o *ZoneTypeServerTypes) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ZoneTypeServerTypes) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ZoneTypeServerTypes) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ZoneTypeServerTypes) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasAgent

`func (o *ZoneTypeServerTypes) GetHasAgent() bool`

GetHasAgent returns the HasAgent field if non-nil, zero value otherwise.

### GetHasAgentOk

`func (o *ZoneTypeServerTypes) GetHasAgentOk() (*bool, bool)`

GetHasAgentOk returns a tuple with the HasAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAgent

`func (o *ZoneTypeServerTypes) SetHasAgent(v bool)`

SetHasAgent sets HasAgent field to given value.

### HasHasAgent

`func (o *ZoneTypeServerTypes) HasHasAgent() bool`

HasHasAgent returns a boolean if a field has been set.

### GetVmHypervisor

`func (o *ZoneTypeServerTypes) GetVmHypervisor() bool`

GetVmHypervisor returns the VmHypervisor field if non-nil, zero value otherwise.

### GetVmHypervisorOk

`func (o *ZoneTypeServerTypes) GetVmHypervisorOk() (*bool, bool)`

GetVmHypervisorOk returns a tuple with the VmHypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHypervisor

`func (o *ZoneTypeServerTypes) SetVmHypervisor(v bool)`

SetVmHypervisor sets VmHypervisor field to given value.

### HasVmHypervisor

`func (o *ZoneTypeServerTypes) HasVmHypervisor() bool`

HasVmHypervisor returns a boolean if a field has been set.

### GetContainerHypervisor

`func (o *ZoneTypeServerTypes) GetContainerHypervisor() bool`

GetContainerHypervisor returns the ContainerHypervisor field if non-nil, zero value otherwise.

### GetContainerHypervisorOk

`func (o *ZoneTypeServerTypes) GetContainerHypervisorOk() (*bool, bool)`

GetContainerHypervisorOk returns a tuple with the ContainerHypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHypervisor

`func (o *ZoneTypeServerTypes) SetContainerHypervisor(v bool)`

SetContainerHypervisor sets ContainerHypervisor field to given value.

### HasContainerHypervisor

`func (o *ZoneTypeServerTypes) HasContainerHypervisor() bool`

HasContainerHypervisor returns a boolean if a field has been set.

### GetBareMetalHost

`func (o *ZoneTypeServerTypes) GetBareMetalHost() bool`

GetBareMetalHost returns the BareMetalHost field if non-nil, zero value otherwise.

### GetBareMetalHostOk

`func (o *ZoneTypeServerTypes) GetBareMetalHostOk() (*bool, bool)`

GetBareMetalHostOk returns a tuple with the BareMetalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetalHost

`func (o *ZoneTypeServerTypes) SetBareMetalHost(v bool)`

SetBareMetalHost sets BareMetalHost field to given value.

### HasBareMetalHost

`func (o *ZoneTypeServerTypes) HasBareMetalHost() bool`

HasBareMetalHost returns a boolean if a field has been set.

### GetGuestVm

`func (o *ZoneTypeServerTypes) GetGuestVm() bool`

GetGuestVm returns the GuestVm field if non-nil, zero value otherwise.

### GetGuestVmOk

`func (o *ZoneTypeServerTypes) GetGuestVmOk() (*bool, bool)`

GetGuestVmOk returns a tuple with the GuestVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVm

`func (o *ZoneTypeServerTypes) SetGuestVm(v bool)`

SetGuestVm sets GuestVm field to given value.

### HasGuestVm

`func (o *ZoneTypeServerTypes) HasGuestVm() bool`

HasGuestVm returns a boolean if a field has been set.

### GetHasAutomation

`func (o *ZoneTypeServerTypes) GetHasAutomation() bool`

GetHasAutomation returns the HasAutomation field if non-nil, zero value otherwise.

### GetHasAutomationOk

`func (o *ZoneTypeServerTypes) GetHasAutomationOk() (*bool, bool)`

GetHasAutomationOk returns a tuple with the HasAutomation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomation

`func (o *ZoneTypeServerTypes) SetHasAutomation(v bool)`

SetHasAutomation sets HasAutomation field to given value.

### HasHasAutomation

`func (o *ZoneTypeServerTypes) HasHasAutomation() bool`

HasHasAutomation returns a boolean if a field has been set.

### GetProvisionType

`func (o *ZoneTypeServerTypes) GetProvisionType() ZoneTypeProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ZoneTypeServerTypes) GetProvisionTypeOk() (*ZoneTypeProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ZoneTypeServerTypes) SetProvisionType(v ZoneTypeProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ZoneTypeServerTypes) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ZoneTypeServerTypes) GetOptionTypes() []ZoneTypeOptionTypes`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ZoneTypeServerTypes) GetOptionTypesOk() (*[]ZoneTypeOptionTypes, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ZoneTypeServerTypes) SetOptionTypes(v []ZoneTypeOptionTypes)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ZoneTypeServerTypes) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ZoneTypeServerTypes) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ZoneTypeServerTypes) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ZoneTypeServerTypes) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ZoneTypeServerTypes) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


