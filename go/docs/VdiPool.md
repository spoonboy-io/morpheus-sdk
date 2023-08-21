# VdiPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**MinIdle** | Pointer to **int64** |  | [optional] 
**MaxIdle** | Pointer to **int64** |  | [optional] 
**InitialPoolSize** | Pointer to **int64** |  | [optional] 
**MaxPoolSize** | Pointer to **int64** |  | [optional] 
**AllocationTimeoutMinutes** | Pointer to **int64** |  | [optional] 
**PersistentUser** | Pointer to **NullableBool** |  | [optional] 
**Recyclable** | Pointer to **NullableBool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AutoCreateLocalUserOnReservation** | Pointer to **bool** |  | [optional] 
**AllowHypervisorConsole** | Pointer to **NullableBool** |  | [optional] 
**AllowCopy** | Pointer to **NullableBool** |  | [optional] 
**AllowPrinter** | Pointer to **NullableBool** |  | [optional] 
**AllowFileshare** | Pointer to **NullableBool** |  | [optional] 
**GuestConsoleJumpHost** | Pointer to **NullableString** |  | [optional] 
**GuestConsoleJumpPort** | Pointer to **NullableString** |  | [optional] 
**GuestConsoleJumpUsername** | Pointer to **NullableString** |  | [optional] 
**GuestConsoleJumpPassword** | Pointer to **NullableString** |  | [optional] 
**GuestConsoleJumpKeypair** | Pointer to **NullableString** |  | [optional] 
**Gateway** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**IconPath** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Apps** | Pointer to [**[]InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Owner** | Pointer to [**VdiPoolOwner**](vdiPool_owner.md) |  | [optional] 
**Config** | Pointer to [**NullableVdiPoolConfig**](vdiPool_config.md) |  | [optional] 
**Group** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**Cloud** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**UsedCount** | Pointer to **int64** |  | [optional] 
**ReservedCount** | Pointer to **int64** |  | [optional] 
**PreparingCount** | Pointer to **int64** |  | [optional] 
**IdleCount** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVdiPool

`func NewVdiPool() *VdiPool`

NewVdiPool instantiates a new VdiPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdiPoolWithDefaults

`func NewVdiPoolWithDefaults() *VdiPool`

NewVdiPoolWithDefaults instantiates a new VdiPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VdiPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VdiPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VdiPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VdiPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VdiPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VdiPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VdiPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VdiPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VdiPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VdiPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VdiPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VdiPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VdiPool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VdiPool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMinIdle

`func (o *VdiPool) GetMinIdle() int64`

GetMinIdle returns the MinIdle field if non-nil, zero value otherwise.

### GetMinIdleOk

`func (o *VdiPool) GetMinIdleOk() (*int64, bool)`

GetMinIdleOk returns a tuple with the MinIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIdle

`func (o *VdiPool) SetMinIdle(v int64)`

SetMinIdle sets MinIdle field to given value.

### HasMinIdle

`func (o *VdiPool) HasMinIdle() bool`

HasMinIdle returns a boolean if a field has been set.

### GetMaxIdle

`func (o *VdiPool) GetMaxIdle() int64`

GetMaxIdle returns the MaxIdle field if non-nil, zero value otherwise.

### GetMaxIdleOk

`func (o *VdiPool) GetMaxIdleOk() (*int64, bool)`

GetMaxIdleOk returns a tuple with the MaxIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIdle

`func (o *VdiPool) SetMaxIdle(v int64)`

SetMaxIdle sets MaxIdle field to given value.

### HasMaxIdle

`func (o *VdiPool) HasMaxIdle() bool`

HasMaxIdle returns a boolean if a field has been set.

### GetInitialPoolSize

`func (o *VdiPool) GetInitialPoolSize() int64`

GetInitialPoolSize returns the InitialPoolSize field if non-nil, zero value otherwise.

### GetInitialPoolSizeOk

`func (o *VdiPool) GetInitialPoolSizeOk() (*int64, bool)`

GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPoolSize

`func (o *VdiPool) SetInitialPoolSize(v int64)`

SetInitialPoolSize sets InitialPoolSize field to given value.

### HasInitialPoolSize

`func (o *VdiPool) HasInitialPoolSize() bool`

HasInitialPoolSize returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *VdiPool) GetMaxPoolSize() int64`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *VdiPool) GetMaxPoolSizeOk() (*int64, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *VdiPool) SetMaxPoolSize(v int64)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *VdiPool) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetAllocationTimeoutMinutes

`func (o *VdiPool) GetAllocationTimeoutMinutes() int64`

GetAllocationTimeoutMinutes returns the AllocationTimeoutMinutes field if non-nil, zero value otherwise.

### GetAllocationTimeoutMinutesOk

`func (o *VdiPool) GetAllocationTimeoutMinutesOk() (*int64, bool)`

GetAllocationTimeoutMinutesOk returns a tuple with the AllocationTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationTimeoutMinutes

`func (o *VdiPool) SetAllocationTimeoutMinutes(v int64)`

SetAllocationTimeoutMinutes sets AllocationTimeoutMinutes field to given value.

### HasAllocationTimeoutMinutes

`func (o *VdiPool) HasAllocationTimeoutMinutes() bool`

HasAllocationTimeoutMinutes returns a boolean if a field has been set.

### GetPersistentUser

`func (o *VdiPool) GetPersistentUser() bool`

GetPersistentUser returns the PersistentUser field if non-nil, zero value otherwise.

### GetPersistentUserOk

`func (o *VdiPool) GetPersistentUserOk() (*bool, bool)`

GetPersistentUserOk returns a tuple with the PersistentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentUser

`func (o *VdiPool) SetPersistentUser(v bool)`

SetPersistentUser sets PersistentUser field to given value.

### HasPersistentUser

`func (o *VdiPool) HasPersistentUser() bool`

HasPersistentUser returns a boolean if a field has been set.

### SetPersistentUserNil

`func (o *VdiPool) SetPersistentUserNil(b bool)`

 SetPersistentUserNil sets the value for PersistentUser to be an explicit nil

### UnsetPersistentUser
`func (o *VdiPool) UnsetPersistentUser()`

UnsetPersistentUser ensures that no value is present for PersistentUser, not even an explicit nil
### GetRecyclable

`func (o *VdiPool) GetRecyclable() bool`

GetRecyclable returns the Recyclable field if non-nil, zero value otherwise.

### GetRecyclableOk

`func (o *VdiPool) GetRecyclableOk() (*bool, bool)`

GetRecyclableOk returns a tuple with the Recyclable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclable

`func (o *VdiPool) SetRecyclable(v bool)`

SetRecyclable sets Recyclable field to given value.

### HasRecyclable

`func (o *VdiPool) HasRecyclable() bool`

HasRecyclable returns a boolean if a field has been set.

### SetRecyclableNil

`func (o *VdiPool) SetRecyclableNil(b bool)`

 SetRecyclableNil sets the value for Recyclable to be an explicit nil

### UnsetRecyclable
`func (o *VdiPool) UnsetRecyclable()`

UnsetRecyclable ensures that no value is present for Recyclable, not even an explicit nil
### GetEnabled

`func (o *VdiPool) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VdiPool) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VdiPool) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VdiPool) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoCreateLocalUserOnReservation

`func (o *VdiPool) GetAutoCreateLocalUserOnReservation() bool`

GetAutoCreateLocalUserOnReservation returns the AutoCreateLocalUserOnReservation field if non-nil, zero value otherwise.

### GetAutoCreateLocalUserOnReservationOk

`func (o *VdiPool) GetAutoCreateLocalUserOnReservationOk() (*bool, bool)`

GetAutoCreateLocalUserOnReservationOk returns a tuple with the AutoCreateLocalUserOnReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateLocalUserOnReservation

`func (o *VdiPool) SetAutoCreateLocalUserOnReservation(v bool)`

SetAutoCreateLocalUserOnReservation sets AutoCreateLocalUserOnReservation field to given value.

### HasAutoCreateLocalUserOnReservation

`func (o *VdiPool) HasAutoCreateLocalUserOnReservation() bool`

HasAutoCreateLocalUserOnReservation returns a boolean if a field has been set.

### GetAllowHypervisorConsole

`func (o *VdiPool) GetAllowHypervisorConsole() bool`

GetAllowHypervisorConsole returns the AllowHypervisorConsole field if non-nil, zero value otherwise.

### GetAllowHypervisorConsoleOk

`func (o *VdiPool) GetAllowHypervisorConsoleOk() (*bool, bool)`

GetAllowHypervisorConsoleOk returns a tuple with the AllowHypervisorConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHypervisorConsole

`func (o *VdiPool) SetAllowHypervisorConsole(v bool)`

SetAllowHypervisorConsole sets AllowHypervisorConsole field to given value.

### HasAllowHypervisorConsole

`func (o *VdiPool) HasAllowHypervisorConsole() bool`

HasAllowHypervisorConsole returns a boolean if a field has been set.

### SetAllowHypervisorConsoleNil

`func (o *VdiPool) SetAllowHypervisorConsoleNil(b bool)`

 SetAllowHypervisorConsoleNil sets the value for AllowHypervisorConsole to be an explicit nil

### UnsetAllowHypervisorConsole
`func (o *VdiPool) UnsetAllowHypervisorConsole()`

UnsetAllowHypervisorConsole ensures that no value is present for AllowHypervisorConsole, not even an explicit nil
### GetAllowCopy

`func (o *VdiPool) GetAllowCopy() bool`

GetAllowCopy returns the AllowCopy field if non-nil, zero value otherwise.

### GetAllowCopyOk

`func (o *VdiPool) GetAllowCopyOk() (*bool, bool)`

GetAllowCopyOk returns a tuple with the AllowCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCopy

`func (o *VdiPool) SetAllowCopy(v bool)`

SetAllowCopy sets AllowCopy field to given value.

### HasAllowCopy

`func (o *VdiPool) HasAllowCopy() bool`

HasAllowCopy returns a boolean if a field has been set.

### SetAllowCopyNil

`func (o *VdiPool) SetAllowCopyNil(b bool)`

 SetAllowCopyNil sets the value for AllowCopy to be an explicit nil

### UnsetAllowCopy
`func (o *VdiPool) UnsetAllowCopy()`

UnsetAllowCopy ensures that no value is present for AllowCopy, not even an explicit nil
### GetAllowPrinter

`func (o *VdiPool) GetAllowPrinter() bool`

GetAllowPrinter returns the AllowPrinter field if non-nil, zero value otherwise.

### GetAllowPrinterOk

`func (o *VdiPool) GetAllowPrinterOk() (*bool, bool)`

GetAllowPrinterOk returns a tuple with the AllowPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrinter

`func (o *VdiPool) SetAllowPrinter(v bool)`

SetAllowPrinter sets AllowPrinter field to given value.

### HasAllowPrinter

`func (o *VdiPool) HasAllowPrinter() bool`

HasAllowPrinter returns a boolean if a field has been set.

### SetAllowPrinterNil

`func (o *VdiPool) SetAllowPrinterNil(b bool)`

 SetAllowPrinterNil sets the value for AllowPrinter to be an explicit nil

### UnsetAllowPrinter
`func (o *VdiPool) UnsetAllowPrinter()`

UnsetAllowPrinter ensures that no value is present for AllowPrinter, not even an explicit nil
### GetAllowFileshare

`func (o *VdiPool) GetAllowFileshare() bool`

GetAllowFileshare returns the AllowFileshare field if non-nil, zero value otherwise.

### GetAllowFileshareOk

`func (o *VdiPool) GetAllowFileshareOk() (*bool, bool)`

GetAllowFileshareOk returns a tuple with the AllowFileshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFileshare

`func (o *VdiPool) SetAllowFileshare(v bool)`

SetAllowFileshare sets AllowFileshare field to given value.

### HasAllowFileshare

`func (o *VdiPool) HasAllowFileshare() bool`

HasAllowFileshare returns a boolean if a field has been set.

### SetAllowFileshareNil

`func (o *VdiPool) SetAllowFileshareNil(b bool)`

 SetAllowFileshareNil sets the value for AllowFileshare to be an explicit nil

### UnsetAllowFileshare
`func (o *VdiPool) UnsetAllowFileshare()`

UnsetAllowFileshare ensures that no value is present for AllowFileshare, not even an explicit nil
### GetGuestConsoleJumpHost

`func (o *VdiPool) GetGuestConsoleJumpHost() string`

GetGuestConsoleJumpHost returns the GuestConsoleJumpHost field if non-nil, zero value otherwise.

### GetGuestConsoleJumpHostOk

`func (o *VdiPool) GetGuestConsoleJumpHostOk() (*string, bool)`

GetGuestConsoleJumpHostOk returns a tuple with the GuestConsoleJumpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpHost

`func (o *VdiPool) SetGuestConsoleJumpHost(v string)`

SetGuestConsoleJumpHost sets GuestConsoleJumpHost field to given value.

### HasGuestConsoleJumpHost

`func (o *VdiPool) HasGuestConsoleJumpHost() bool`

HasGuestConsoleJumpHost returns a boolean if a field has been set.

### SetGuestConsoleJumpHostNil

`func (o *VdiPool) SetGuestConsoleJumpHostNil(b bool)`

 SetGuestConsoleJumpHostNil sets the value for GuestConsoleJumpHost to be an explicit nil

### UnsetGuestConsoleJumpHost
`func (o *VdiPool) UnsetGuestConsoleJumpHost()`

UnsetGuestConsoleJumpHost ensures that no value is present for GuestConsoleJumpHost, not even an explicit nil
### GetGuestConsoleJumpPort

`func (o *VdiPool) GetGuestConsoleJumpPort() string`

GetGuestConsoleJumpPort returns the GuestConsoleJumpPort field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPortOk

`func (o *VdiPool) GetGuestConsoleJumpPortOk() (*string, bool)`

GetGuestConsoleJumpPortOk returns a tuple with the GuestConsoleJumpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPort

`func (o *VdiPool) SetGuestConsoleJumpPort(v string)`

SetGuestConsoleJumpPort sets GuestConsoleJumpPort field to given value.

### HasGuestConsoleJumpPort

`func (o *VdiPool) HasGuestConsoleJumpPort() bool`

HasGuestConsoleJumpPort returns a boolean if a field has been set.

### SetGuestConsoleJumpPortNil

`func (o *VdiPool) SetGuestConsoleJumpPortNil(b bool)`

 SetGuestConsoleJumpPortNil sets the value for GuestConsoleJumpPort to be an explicit nil

### UnsetGuestConsoleJumpPort
`func (o *VdiPool) UnsetGuestConsoleJumpPort()`

UnsetGuestConsoleJumpPort ensures that no value is present for GuestConsoleJumpPort, not even an explicit nil
### GetGuestConsoleJumpUsername

`func (o *VdiPool) GetGuestConsoleJumpUsername() string`

GetGuestConsoleJumpUsername returns the GuestConsoleJumpUsername field if non-nil, zero value otherwise.

### GetGuestConsoleJumpUsernameOk

`func (o *VdiPool) GetGuestConsoleJumpUsernameOk() (*string, bool)`

GetGuestConsoleJumpUsernameOk returns a tuple with the GuestConsoleJumpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpUsername

`func (o *VdiPool) SetGuestConsoleJumpUsername(v string)`

SetGuestConsoleJumpUsername sets GuestConsoleJumpUsername field to given value.

### HasGuestConsoleJumpUsername

`func (o *VdiPool) HasGuestConsoleJumpUsername() bool`

HasGuestConsoleJumpUsername returns a boolean if a field has been set.

### SetGuestConsoleJumpUsernameNil

`func (o *VdiPool) SetGuestConsoleJumpUsernameNil(b bool)`

 SetGuestConsoleJumpUsernameNil sets the value for GuestConsoleJumpUsername to be an explicit nil

### UnsetGuestConsoleJumpUsername
`func (o *VdiPool) UnsetGuestConsoleJumpUsername()`

UnsetGuestConsoleJumpUsername ensures that no value is present for GuestConsoleJumpUsername, not even an explicit nil
### GetGuestConsoleJumpPassword

`func (o *VdiPool) GetGuestConsoleJumpPassword() string`

GetGuestConsoleJumpPassword returns the GuestConsoleJumpPassword field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPasswordOk

`func (o *VdiPool) GetGuestConsoleJumpPasswordOk() (*string, bool)`

GetGuestConsoleJumpPasswordOk returns a tuple with the GuestConsoleJumpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPassword

`func (o *VdiPool) SetGuestConsoleJumpPassword(v string)`

SetGuestConsoleJumpPassword sets GuestConsoleJumpPassword field to given value.

### HasGuestConsoleJumpPassword

`func (o *VdiPool) HasGuestConsoleJumpPassword() bool`

HasGuestConsoleJumpPassword returns a boolean if a field has been set.

### SetGuestConsoleJumpPasswordNil

`func (o *VdiPool) SetGuestConsoleJumpPasswordNil(b bool)`

 SetGuestConsoleJumpPasswordNil sets the value for GuestConsoleJumpPassword to be an explicit nil

### UnsetGuestConsoleJumpPassword
`func (o *VdiPool) UnsetGuestConsoleJumpPassword()`

UnsetGuestConsoleJumpPassword ensures that no value is present for GuestConsoleJumpPassword, not even an explicit nil
### GetGuestConsoleJumpKeypair

`func (o *VdiPool) GetGuestConsoleJumpKeypair() string`

GetGuestConsoleJumpKeypair returns the GuestConsoleJumpKeypair field if non-nil, zero value otherwise.

### GetGuestConsoleJumpKeypairOk

`func (o *VdiPool) GetGuestConsoleJumpKeypairOk() (*string, bool)`

GetGuestConsoleJumpKeypairOk returns a tuple with the GuestConsoleJumpKeypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpKeypair

`func (o *VdiPool) SetGuestConsoleJumpKeypair(v string)`

SetGuestConsoleJumpKeypair sets GuestConsoleJumpKeypair field to given value.

### HasGuestConsoleJumpKeypair

`func (o *VdiPool) HasGuestConsoleJumpKeypair() bool`

HasGuestConsoleJumpKeypair returns a boolean if a field has been set.

### SetGuestConsoleJumpKeypairNil

`func (o *VdiPool) SetGuestConsoleJumpKeypairNil(b bool)`

 SetGuestConsoleJumpKeypairNil sets the value for GuestConsoleJumpKeypair to be an explicit nil

### UnsetGuestConsoleJumpKeypair
`func (o *VdiPool) UnsetGuestConsoleJumpKeypair()`

UnsetGuestConsoleJumpKeypair ensures that no value is present for GuestConsoleJumpKeypair, not even an explicit nil
### GetGateway

`func (o *VdiPool) GetGateway() InlineResponse20082LoadBalancerInstanceSslCert`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *VdiPool) GetGatewayOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *VdiPool) SetGateway(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *VdiPool) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *VdiPool) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *VdiPool) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetIconPath

`func (o *VdiPool) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *VdiPool) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *VdiPool) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *VdiPool) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetLogo

`func (o *VdiPool) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *VdiPool) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *VdiPool) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *VdiPool) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetApps

`func (o *VdiPool) GetApps() []InlineResponse20040AppDeployInstance`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *VdiPool) GetAppsOk() (*[]InlineResponse20040AppDeployInstance, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *VdiPool) SetApps(v []InlineResponse20040AppDeployInstance)`

SetApps sets Apps field to given value.

### HasApps

`func (o *VdiPool) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetOwner

`func (o *VdiPool) GetOwner() VdiPoolOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VdiPool) GetOwnerOk() (*VdiPoolOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VdiPool) SetOwner(v VdiPoolOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VdiPool) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetConfig

`func (o *VdiPool) GetConfig() VdiPoolConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VdiPool) GetConfigOk() (*VdiPoolConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VdiPool) SetConfig(v VdiPoolConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *VdiPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *VdiPool) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *VdiPool) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetGroup

`func (o *VdiPool) GetGroup() InlineResponse20082LoadBalancerInstanceSslCert`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VdiPool) GetGroupOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VdiPool) SetGroup(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *VdiPool) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *VdiPool) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *VdiPool) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetCloud

`func (o *VdiPool) GetCloud() InlineResponse20082LoadBalancerInstanceSslCert`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *VdiPool) GetCloudOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *VdiPool) SetCloud(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *VdiPool) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### SetCloudNil

`func (o *VdiPool) SetCloudNil(b bool)`

 SetCloudNil sets the value for Cloud to be an explicit nil

### UnsetCloud
`func (o *VdiPool) UnsetCloud()`

UnsetCloud ensures that no value is present for Cloud, not even an explicit nil
### GetUsedCount

`func (o *VdiPool) GetUsedCount() int64`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *VdiPool) GetUsedCountOk() (*int64, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *VdiPool) SetUsedCount(v int64)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *VdiPool) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetReservedCount

`func (o *VdiPool) GetReservedCount() int64`

GetReservedCount returns the ReservedCount field if non-nil, zero value otherwise.

### GetReservedCountOk

`func (o *VdiPool) GetReservedCountOk() (*int64, bool)`

GetReservedCountOk returns a tuple with the ReservedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCount

`func (o *VdiPool) SetReservedCount(v int64)`

SetReservedCount sets ReservedCount field to given value.

### HasReservedCount

`func (o *VdiPool) HasReservedCount() bool`

HasReservedCount returns a boolean if a field has been set.

### GetPreparingCount

`func (o *VdiPool) GetPreparingCount() int64`

GetPreparingCount returns the PreparingCount field if non-nil, zero value otherwise.

### GetPreparingCountOk

`func (o *VdiPool) GetPreparingCountOk() (*int64, bool)`

GetPreparingCountOk returns a tuple with the PreparingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparingCount

`func (o *VdiPool) SetPreparingCount(v int64)`

SetPreparingCount sets PreparingCount field to given value.

### HasPreparingCount

`func (o *VdiPool) HasPreparingCount() bool`

HasPreparingCount returns a boolean if a field has been set.

### GetIdleCount

`func (o *VdiPool) GetIdleCount() int64`

GetIdleCount returns the IdleCount field if non-nil, zero value otherwise.

### GetIdleCountOk

`func (o *VdiPool) GetIdleCountOk() (*int64, bool)`

GetIdleCountOk returns a tuple with the IdleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleCount

`func (o *VdiPool) SetIdleCount(v int64)`

SetIdleCount sets IdleCount field to given value.

### HasIdleCount

`func (o *VdiPool) HasIdleCount() bool`

HasIdleCount returns a boolean if a field has been set.

### GetStatus

`func (o *VdiPool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VdiPool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VdiPool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VdiPool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *VdiPool) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VdiPool) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VdiPool) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VdiPool) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *VdiPool) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VdiPool) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VdiPool) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *VdiPool) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


