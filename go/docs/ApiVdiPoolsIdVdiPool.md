# ApiVdiPoolsIdVdiPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Virtual Desktop name | [optional] 
**Description** | Pointer to **string** | Virtual Desktop description | [optional] 
**Owner** | Pointer to **int64** | Owner (User) ID | [optional] 
**MinIdle** | Pointer to **float32** | Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby.  | [optional] 
**InitialPoolSize** | Pointer to **float32** | The initial size of the pool to be allocated on creation | [optional] 
**MaxIdle** | Pointer to **float32** | Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances  | [optional] 
**MaxPoolSize** | Pointer to **float32** | Max limit on number of allocations and instances within the pool.  | [optional] 
**AllocationTimeoutMinutes** | Pointer to **float32** | Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence.  | [optional] 
**PersistentUser** | Pointer to **bool** | Persistent Desktop Pool | [optional] [default to false]
**Recyclable** | Pointer to **bool** | Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD) | [optional] [default to false]
**AllowCopy** | Pointer to **bool** | Allow copy from desktop | [optional] [default to false]
**AllowPrinter** | Pointer to **bool** | Allow local printers from Desktop | [optional] [default to false]
**AllowFileshare** | Pointer to **bool** | Allow File Share | [optional] [default to false]
**AllowHypervisorConsole** | Pointer to **bool** | Allow hypervisor console | [optional] [default to false]
**AutoCreateLocalUserOnReservation** | Pointer to **bool** | Auto create local user upon reservation | [optional] [default to false]
**Enabled** | Pointer to **bool** | Can be used to enable or disable the VDI pool | [optional] [default to true]
**IconPath** | Pointer to **string** | The relative location of an icon image | [optional] 
**Apps** | Pointer to **[]int64** | Array of VDI App IDs | [optional] 
**Gateway** | Pointer to **int64** | VDI Gateway ID | [optional] 
**InstanceConfig** | Pointer to **string** | Instance Config JSON. Passing as a string will preserve property order.  See &#x60;config&#x60; object for required values. | [optional] 
**Config** | Pointer to [**ApiVdiPoolsIdVdiPoolConfig**](_api_vdi_pools__id__vdiPool_config.md) |  | [optional] 
**GuestConsoleJumpHost** | Pointer to **string** | Guest Console Jump Host | [optional] 
**GuestConsoleJumpPort** | Pointer to **int64** | Guest Console Jump Port | [optional] 
**GuestConsoleJumpUsername** | Pointer to **string** | Guest Console Jump Username | [optional] 
**GuestConsoleJumpPassword** | Pointer to **string** | Guest Console Jump Password | [optional] 
**GuestConsoleJumpKeypair** | Pointer to **int64** | Guest Console Jump Key Pair. see &#x60;Key Pair&#x60; | [optional] 

## Methods

### NewApiVdiPoolsIdVdiPool

`func NewApiVdiPoolsIdVdiPool() *ApiVdiPoolsIdVdiPool`

NewApiVdiPoolsIdVdiPool instantiates a new ApiVdiPoolsIdVdiPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiVdiPoolsIdVdiPoolWithDefaults

`func NewApiVdiPoolsIdVdiPoolWithDefaults() *ApiVdiPoolsIdVdiPool`

NewApiVdiPoolsIdVdiPoolWithDefaults instantiates a new ApiVdiPoolsIdVdiPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiVdiPoolsIdVdiPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiVdiPoolsIdVdiPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiVdiPoolsIdVdiPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiVdiPoolsIdVdiPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiVdiPoolsIdVdiPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiVdiPoolsIdVdiPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiVdiPoolsIdVdiPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiVdiPoolsIdVdiPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *ApiVdiPoolsIdVdiPool) GetOwner() int64`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ApiVdiPoolsIdVdiPool) GetOwnerOk() (*int64, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ApiVdiPoolsIdVdiPool) SetOwner(v int64)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ApiVdiPoolsIdVdiPool) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetMinIdle

`func (o *ApiVdiPoolsIdVdiPool) GetMinIdle() float32`

GetMinIdle returns the MinIdle field if non-nil, zero value otherwise.

### GetMinIdleOk

`func (o *ApiVdiPoolsIdVdiPool) GetMinIdleOk() (*float32, bool)`

GetMinIdleOk returns a tuple with the MinIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIdle

`func (o *ApiVdiPoolsIdVdiPool) SetMinIdle(v float32)`

SetMinIdle sets MinIdle field to given value.

### HasMinIdle

`func (o *ApiVdiPoolsIdVdiPool) HasMinIdle() bool`

HasMinIdle returns a boolean if a field has been set.

### GetInitialPoolSize

`func (o *ApiVdiPoolsIdVdiPool) GetInitialPoolSize() float32`

GetInitialPoolSize returns the InitialPoolSize field if non-nil, zero value otherwise.

### GetInitialPoolSizeOk

`func (o *ApiVdiPoolsIdVdiPool) GetInitialPoolSizeOk() (*float32, bool)`

GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPoolSize

`func (o *ApiVdiPoolsIdVdiPool) SetInitialPoolSize(v float32)`

SetInitialPoolSize sets InitialPoolSize field to given value.

### HasInitialPoolSize

`func (o *ApiVdiPoolsIdVdiPool) HasInitialPoolSize() bool`

HasInitialPoolSize returns a boolean if a field has been set.

### GetMaxIdle

`func (o *ApiVdiPoolsIdVdiPool) GetMaxIdle() float32`

GetMaxIdle returns the MaxIdle field if non-nil, zero value otherwise.

### GetMaxIdleOk

`func (o *ApiVdiPoolsIdVdiPool) GetMaxIdleOk() (*float32, bool)`

GetMaxIdleOk returns a tuple with the MaxIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIdle

`func (o *ApiVdiPoolsIdVdiPool) SetMaxIdle(v float32)`

SetMaxIdle sets MaxIdle field to given value.

### HasMaxIdle

`func (o *ApiVdiPoolsIdVdiPool) HasMaxIdle() bool`

HasMaxIdle returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *ApiVdiPoolsIdVdiPool) GetMaxPoolSize() float32`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *ApiVdiPoolsIdVdiPool) GetMaxPoolSizeOk() (*float32, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *ApiVdiPoolsIdVdiPool) SetMaxPoolSize(v float32)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *ApiVdiPoolsIdVdiPool) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetAllocationTimeoutMinutes

`func (o *ApiVdiPoolsIdVdiPool) GetAllocationTimeoutMinutes() float32`

GetAllocationTimeoutMinutes returns the AllocationTimeoutMinutes field if non-nil, zero value otherwise.

### GetAllocationTimeoutMinutesOk

`func (o *ApiVdiPoolsIdVdiPool) GetAllocationTimeoutMinutesOk() (*float32, bool)`

GetAllocationTimeoutMinutesOk returns a tuple with the AllocationTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationTimeoutMinutes

`func (o *ApiVdiPoolsIdVdiPool) SetAllocationTimeoutMinutes(v float32)`

SetAllocationTimeoutMinutes sets AllocationTimeoutMinutes field to given value.

### HasAllocationTimeoutMinutes

`func (o *ApiVdiPoolsIdVdiPool) HasAllocationTimeoutMinutes() bool`

HasAllocationTimeoutMinutes returns a boolean if a field has been set.

### GetPersistentUser

`func (o *ApiVdiPoolsIdVdiPool) GetPersistentUser() bool`

GetPersistentUser returns the PersistentUser field if non-nil, zero value otherwise.

### GetPersistentUserOk

`func (o *ApiVdiPoolsIdVdiPool) GetPersistentUserOk() (*bool, bool)`

GetPersistentUserOk returns a tuple with the PersistentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentUser

`func (o *ApiVdiPoolsIdVdiPool) SetPersistentUser(v bool)`

SetPersistentUser sets PersistentUser field to given value.

### HasPersistentUser

`func (o *ApiVdiPoolsIdVdiPool) HasPersistentUser() bool`

HasPersistentUser returns a boolean if a field has been set.

### GetRecyclable

`func (o *ApiVdiPoolsIdVdiPool) GetRecyclable() bool`

GetRecyclable returns the Recyclable field if non-nil, zero value otherwise.

### GetRecyclableOk

`func (o *ApiVdiPoolsIdVdiPool) GetRecyclableOk() (*bool, bool)`

GetRecyclableOk returns a tuple with the Recyclable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclable

`func (o *ApiVdiPoolsIdVdiPool) SetRecyclable(v bool)`

SetRecyclable sets Recyclable field to given value.

### HasRecyclable

`func (o *ApiVdiPoolsIdVdiPool) HasRecyclable() bool`

HasRecyclable returns a boolean if a field has been set.

### GetAllowCopy

`func (o *ApiVdiPoolsIdVdiPool) GetAllowCopy() bool`

GetAllowCopy returns the AllowCopy field if non-nil, zero value otherwise.

### GetAllowCopyOk

`func (o *ApiVdiPoolsIdVdiPool) GetAllowCopyOk() (*bool, bool)`

GetAllowCopyOk returns a tuple with the AllowCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCopy

`func (o *ApiVdiPoolsIdVdiPool) SetAllowCopy(v bool)`

SetAllowCopy sets AllowCopy field to given value.

### HasAllowCopy

`func (o *ApiVdiPoolsIdVdiPool) HasAllowCopy() bool`

HasAllowCopy returns a boolean if a field has been set.

### GetAllowPrinter

`func (o *ApiVdiPoolsIdVdiPool) GetAllowPrinter() bool`

GetAllowPrinter returns the AllowPrinter field if non-nil, zero value otherwise.

### GetAllowPrinterOk

`func (o *ApiVdiPoolsIdVdiPool) GetAllowPrinterOk() (*bool, bool)`

GetAllowPrinterOk returns a tuple with the AllowPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrinter

`func (o *ApiVdiPoolsIdVdiPool) SetAllowPrinter(v bool)`

SetAllowPrinter sets AllowPrinter field to given value.

### HasAllowPrinter

`func (o *ApiVdiPoolsIdVdiPool) HasAllowPrinter() bool`

HasAllowPrinter returns a boolean if a field has been set.

### GetAllowFileshare

`func (o *ApiVdiPoolsIdVdiPool) GetAllowFileshare() bool`

GetAllowFileshare returns the AllowFileshare field if non-nil, zero value otherwise.

### GetAllowFileshareOk

`func (o *ApiVdiPoolsIdVdiPool) GetAllowFileshareOk() (*bool, bool)`

GetAllowFileshareOk returns a tuple with the AllowFileshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFileshare

`func (o *ApiVdiPoolsIdVdiPool) SetAllowFileshare(v bool)`

SetAllowFileshare sets AllowFileshare field to given value.

### HasAllowFileshare

`func (o *ApiVdiPoolsIdVdiPool) HasAllowFileshare() bool`

HasAllowFileshare returns a boolean if a field has been set.

### GetAllowHypervisorConsole

`func (o *ApiVdiPoolsIdVdiPool) GetAllowHypervisorConsole() bool`

GetAllowHypervisorConsole returns the AllowHypervisorConsole field if non-nil, zero value otherwise.

### GetAllowHypervisorConsoleOk

`func (o *ApiVdiPoolsIdVdiPool) GetAllowHypervisorConsoleOk() (*bool, bool)`

GetAllowHypervisorConsoleOk returns a tuple with the AllowHypervisorConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHypervisorConsole

`func (o *ApiVdiPoolsIdVdiPool) SetAllowHypervisorConsole(v bool)`

SetAllowHypervisorConsole sets AllowHypervisorConsole field to given value.

### HasAllowHypervisorConsole

`func (o *ApiVdiPoolsIdVdiPool) HasAllowHypervisorConsole() bool`

HasAllowHypervisorConsole returns a boolean if a field has been set.

### GetAutoCreateLocalUserOnReservation

`func (o *ApiVdiPoolsIdVdiPool) GetAutoCreateLocalUserOnReservation() bool`

GetAutoCreateLocalUserOnReservation returns the AutoCreateLocalUserOnReservation field if non-nil, zero value otherwise.

### GetAutoCreateLocalUserOnReservationOk

`func (o *ApiVdiPoolsIdVdiPool) GetAutoCreateLocalUserOnReservationOk() (*bool, bool)`

GetAutoCreateLocalUserOnReservationOk returns a tuple with the AutoCreateLocalUserOnReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateLocalUserOnReservation

`func (o *ApiVdiPoolsIdVdiPool) SetAutoCreateLocalUserOnReservation(v bool)`

SetAutoCreateLocalUserOnReservation sets AutoCreateLocalUserOnReservation field to given value.

### HasAutoCreateLocalUserOnReservation

`func (o *ApiVdiPoolsIdVdiPool) HasAutoCreateLocalUserOnReservation() bool`

HasAutoCreateLocalUserOnReservation returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiVdiPoolsIdVdiPool) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiVdiPoolsIdVdiPool) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiVdiPoolsIdVdiPool) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiVdiPoolsIdVdiPool) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIconPath

`func (o *ApiVdiPoolsIdVdiPool) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *ApiVdiPoolsIdVdiPool) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *ApiVdiPoolsIdVdiPool) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *ApiVdiPoolsIdVdiPool) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetApps

`func (o *ApiVdiPoolsIdVdiPool) GetApps() []int64`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ApiVdiPoolsIdVdiPool) GetAppsOk() (*[]int64, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ApiVdiPoolsIdVdiPool) SetApps(v []int64)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ApiVdiPoolsIdVdiPool) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetGateway

`func (o *ApiVdiPoolsIdVdiPool) GetGateway() int64`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ApiVdiPoolsIdVdiPool) GetGatewayOk() (*int64, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ApiVdiPoolsIdVdiPool) SetGateway(v int64)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ApiVdiPoolsIdVdiPool) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInstanceConfig

`func (o *ApiVdiPoolsIdVdiPool) GetInstanceConfig() string`

GetInstanceConfig returns the InstanceConfig field if non-nil, zero value otherwise.

### GetInstanceConfigOk

`func (o *ApiVdiPoolsIdVdiPool) GetInstanceConfigOk() (*string, bool)`

GetInstanceConfigOk returns a tuple with the InstanceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceConfig

`func (o *ApiVdiPoolsIdVdiPool) SetInstanceConfig(v string)`

SetInstanceConfig sets InstanceConfig field to given value.

### HasInstanceConfig

`func (o *ApiVdiPoolsIdVdiPool) HasInstanceConfig() bool`

HasInstanceConfig returns a boolean if a field has been set.

### GetConfig

`func (o *ApiVdiPoolsIdVdiPool) GetConfig() ApiVdiPoolsIdVdiPoolConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiVdiPoolsIdVdiPool) GetConfigOk() (*ApiVdiPoolsIdVdiPoolConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiVdiPoolsIdVdiPool) SetConfig(v ApiVdiPoolsIdVdiPoolConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiVdiPoolsIdVdiPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetGuestConsoleJumpHost

`func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpHost() string`

GetGuestConsoleJumpHost returns the GuestConsoleJumpHost field if non-nil, zero value otherwise.

### GetGuestConsoleJumpHostOk

`func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpHostOk() (*string, bool)`

GetGuestConsoleJumpHostOk returns a tuple with the GuestConsoleJumpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpHost

`func (o *ApiVdiPoolsIdVdiPool) SetGuestConsoleJumpHost(v string)`

SetGuestConsoleJumpHost sets GuestConsoleJumpHost field to given value.

### HasGuestConsoleJumpHost

`func (o *ApiVdiPoolsIdVdiPool) HasGuestConsoleJumpHost() bool`

HasGuestConsoleJumpHost returns a boolean if a field has been set.

### GetGuestConsoleJumpPort

`func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpPort() int64`

GetGuestConsoleJumpPort returns the GuestConsoleJumpPort field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPortOk

`func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpPortOk() (*int64, bool)`

GetGuestConsoleJumpPortOk returns a tuple with the GuestConsoleJumpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPort

`func (o *ApiVdiPoolsIdVdiPool) SetGuestConsoleJumpPort(v int64)`

SetGuestConsoleJumpPort sets GuestConsoleJumpPort field to given value.

### HasGuestConsoleJumpPort

`func (o *ApiVdiPoolsIdVdiPool) HasGuestConsoleJumpPort() bool`

HasGuestConsoleJumpPort returns a boolean if a field has been set.

### GetGuestConsoleJumpUsername

`func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpUsername() string`

GetGuestConsoleJumpUsername returns the GuestConsoleJumpUsername field if non-nil, zero value otherwise.

### GetGuestConsoleJumpUsernameOk

`func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpUsernameOk() (*string, bool)`

GetGuestConsoleJumpUsernameOk returns a tuple with the GuestConsoleJumpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpUsername

`func (o *ApiVdiPoolsIdVdiPool) SetGuestConsoleJumpUsername(v string)`

SetGuestConsoleJumpUsername sets GuestConsoleJumpUsername field to given value.

### HasGuestConsoleJumpUsername

`func (o *ApiVdiPoolsIdVdiPool) HasGuestConsoleJumpUsername() bool`

HasGuestConsoleJumpUsername returns a boolean if a field has been set.

### GetGuestConsoleJumpPassword

`func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpPassword() string`

GetGuestConsoleJumpPassword returns the GuestConsoleJumpPassword field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPasswordOk

`func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpPasswordOk() (*string, bool)`

GetGuestConsoleJumpPasswordOk returns a tuple with the GuestConsoleJumpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPassword

`func (o *ApiVdiPoolsIdVdiPool) SetGuestConsoleJumpPassword(v string)`

SetGuestConsoleJumpPassword sets GuestConsoleJumpPassword field to given value.

### HasGuestConsoleJumpPassword

`func (o *ApiVdiPoolsIdVdiPool) HasGuestConsoleJumpPassword() bool`

HasGuestConsoleJumpPassword returns a boolean if a field has been set.

### GetGuestConsoleJumpKeypair

`func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpKeypair() int64`

GetGuestConsoleJumpKeypair returns the GuestConsoleJumpKeypair field if non-nil, zero value otherwise.

### GetGuestConsoleJumpKeypairOk

`func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpKeypairOk() (*int64, bool)`

GetGuestConsoleJumpKeypairOk returns a tuple with the GuestConsoleJumpKeypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpKeypair

`func (o *ApiVdiPoolsIdVdiPool) SetGuestConsoleJumpKeypair(v int64)`

SetGuestConsoleJumpKeypair sets GuestConsoleJumpKeypair field to given value.

### HasGuestConsoleJumpKeypair

`func (o *ApiVdiPoolsIdVdiPool) HasGuestConsoleJumpKeypair() bool`

HasGuestConsoleJumpKeypair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


