# ApplianceSettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **string** | Appliance URL | [optional] 
**InternalApplianceUrl** | Pointer to **NullableString** | Internal Appliance URL (PXE) | [optional] 
**CorsAllowed** | Pointer to **NullableString** | API Allowed Origins | [optional] 
**RegistrationEnabled** | Pointer to **bool** | Registration enabled (true, false) | [optional] 
**DefaultRoleId** | Pointer to **int64** | Default tenant role ID | [optional] 
**DefaultUserRoleId** | Pointer to **int64** | Default user role ID | [optional] 
**DockerPrivilegedMode** | Pointer to **bool** | Docker privileged mode (true, false) | [optional] 
**PasswordMinLength** | Pointer to **string** | Min Password Length | [optional] 
**PasswordMinUpperCase** | Pointer to **string** | Min Password Uppercase | [optional] 
**PasswordMinNumbers** | Pointer to **string** | Min Password Numbers | [optional] 
**PasswordMinSymbols** | Pointer to **string** | Min Password Symbols | [optional] 
**UserBrowserSessionTimeout** | Pointer to **string** | User Browser Session Timeout (Minutes) | [optional] 
**UserBrowserSessionWarning** | Pointer to **string** | User Browser Session Warning (Minutes) | [optional] 
**ExpirePwdDays** | Pointer to **int64** | Expire password after days. Setting to 0 disabled this feature | [optional] 
**DisableAfterAttempts** | Pointer to **int64** | Disable user after number of attempts. Set to 0 to disable this feature | [optional] 
**DisableAfterDaysInactive** | Pointer to **int64** | Disable user if inactive for specified days. Set to 0 to disable this feature | [optional] 
**WarnUserDaysBefore** | Pointer to **int64** | Send warning email number of days in advance before deactivating. Set to 0 to disable this feature | [optional] 
**SmtpMailFrom** | Pointer to **string** | From email address | [optional] 
**SmtpServer** | Pointer to **string** | SMTP server / host | [optional] 
**SmtpPort** | Pointer to **int64** | SMTP port | [optional] 
**SmtpSSL** | Pointer to **bool** | Use SSL for SMTP connection | [optional] 
**SmtpTLS** | Pointer to **bool** | Use TLS for SMTP connections | [optional] 
**SmtpUser** | Pointer to **string** | SMTP username | [optional] 
**SmtpPassword** | Pointer to **string** | SMTP password | [optional] 
**ProxyHost** | Pointer to **NullableString** | Proxy host | [optional] 
**ProxyPort** | Pointer to **NullableString** | Proxy port | [optional] 
**ProxyUser** | Pointer to **string** | Proxy username | [optional] 
**ProxyPassword** | Pointer to **string** | Proxy password | [optional] 
**ProxyDomain** | Pointer to **NullableString** | Proxy domain | [optional] 
**ProxyWorkstation** | Pointer to **NullableString** | Proxy workstation | [optional] 
**CurrencyProvider** | Pointer to **string** | Currency provider | [optional] 
**CurrencyKey** | Pointer to **NullableString** | Currency provider API key | [optional] 
**EnableAllZoneTypes** | Pointer to **bool** | Set all cloud types enabled status on, overrides enableZoneTypes and disableZoneTypes parameters | [optional] 
**EnableZoneTypes** | Pointer to **[]int64** | List of cloud type IDs to set enabled status on | [optional] 
**DisableZoneTypes** | Pointer to **[]int64** | List of cloud type IDs to set enabled status off | [optional] 
**DisableAllZoneTypes** | Pointer to **bool** | Set all cloud types enabled status off, can be used in conjunction with enableZoneTypes | [optional] 

## Methods

### NewApplianceSettingsUpdate

`func NewApplianceSettingsUpdate() *ApplianceSettingsUpdate`

NewApplianceSettingsUpdate instantiates a new ApplianceSettingsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceSettingsUpdateWithDefaults

`func NewApplianceSettingsUpdateWithDefaults() *ApplianceSettingsUpdate`

NewApplianceSettingsUpdateWithDefaults instantiates a new ApplianceSettingsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *ApplianceSettingsUpdate) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ApplianceSettingsUpdate) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ApplianceSettingsUpdate) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ApplianceSettingsUpdate) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetInternalApplianceUrl

`func (o *ApplianceSettingsUpdate) GetInternalApplianceUrl() string`

GetInternalApplianceUrl returns the InternalApplianceUrl field if non-nil, zero value otherwise.

### GetInternalApplianceUrlOk

`func (o *ApplianceSettingsUpdate) GetInternalApplianceUrlOk() (*string, bool)`

GetInternalApplianceUrlOk returns a tuple with the InternalApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalApplianceUrl

`func (o *ApplianceSettingsUpdate) SetInternalApplianceUrl(v string)`

SetInternalApplianceUrl sets InternalApplianceUrl field to given value.

### HasInternalApplianceUrl

`func (o *ApplianceSettingsUpdate) HasInternalApplianceUrl() bool`

HasInternalApplianceUrl returns a boolean if a field has been set.

### SetInternalApplianceUrlNil

`func (o *ApplianceSettingsUpdate) SetInternalApplianceUrlNil(b bool)`

 SetInternalApplianceUrlNil sets the value for InternalApplianceUrl to be an explicit nil

### UnsetInternalApplianceUrl
`func (o *ApplianceSettingsUpdate) UnsetInternalApplianceUrl()`

UnsetInternalApplianceUrl ensures that no value is present for InternalApplianceUrl, not even an explicit nil
### GetCorsAllowed

`func (o *ApplianceSettingsUpdate) GetCorsAllowed() string`

GetCorsAllowed returns the CorsAllowed field if non-nil, zero value otherwise.

### GetCorsAllowedOk

`func (o *ApplianceSettingsUpdate) GetCorsAllowedOk() (*string, bool)`

GetCorsAllowedOk returns a tuple with the CorsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowed

`func (o *ApplianceSettingsUpdate) SetCorsAllowed(v string)`

SetCorsAllowed sets CorsAllowed field to given value.

### HasCorsAllowed

`func (o *ApplianceSettingsUpdate) HasCorsAllowed() bool`

HasCorsAllowed returns a boolean if a field has been set.

### SetCorsAllowedNil

`func (o *ApplianceSettingsUpdate) SetCorsAllowedNil(b bool)`

 SetCorsAllowedNil sets the value for CorsAllowed to be an explicit nil

### UnsetCorsAllowed
`func (o *ApplianceSettingsUpdate) UnsetCorsAllowed()`

UnsetCorsAllowed ensures that no value is present for CorsAllowed, not even an explicit nil
### GetRegistrationEnabled

`func (o *ApplianceSettingsUpdate) GetRegistrationEnabled() bool`

GetRegistrationEnabled returns the RegistrationEnabled field if non-nil, zero value otherwise.

### GetRegistrationEnabledOk

`func (o *ApplianceSettingsUpdate) GetRegistrationEnabledOk() (*bool, bool)`

GetRegistrationEnabledOk returns a tuple with the RegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEnabled

`func (o *ApplianceSettingsUpdate) SetRegistrationEnabled(v bool)`

SetRegistrationEnabled sets RegistrationEnabled field to given value.

### HasRegistrationEnabled

`func (o *ApplianceSettingsUpdate) HasRegistrationEnabled() bool`

HasRegistrationEnabled returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *ApplianceSettingsUpdate) GetDefaultRoleId() int64`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *ApplianceSettingsUpdate) GetDefaultRoleIdOk() (*int64, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *ApplianceSettingsUpdate) SetDefaultRoleId(v int64)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *ApplianceSettingsUpdate) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetDefaultUserRoleId

`func (o *ApplianceSettingsUpdate) GetDefaultUserRoleId() int64`

GetDefaultUserRoleId returns the DefaultUserRoleId field if non-nil, zero value otherwise.

### GetDefaultUserRoleIdOk

`func (o *ApplianceSettingsUpdate) GetDefaultUserRoleIdOk() (*int64, bool)`

GetDefaultUserRoleIdOk returns a tuple with the DefaultUserRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserRoleId

`func (o *ApplianceSettingsUpdate) SetDefaultUserRoleId(v int64)`

SetDefaultUserRoleId sets DefaultUserRoleId field to given value.

### HasDefaultUserRoleId

`func (o *ApplianceSettingsUpdate) HasDefaultUserRoleId() bool`

HasDefaultUserRoleId returns a boolean if a field has been set.

### GetDockerPrivilegedMode

`func (o *ApplianceSettingsUpdate) GetDockerPrivilegedMode() bool`

GetDockerPrivilegedMode returns the DockerPrivilegedMode field if non-nil, zero value otherwise.

### GetDockerPrivilegedModeOk

`func (o *ApplianceSettingsUpdate) GetDockerPrivilegedModeOk() (*bool, bool)`

GetDockerPrivilegedModeOk returns a tuple with the DockerPrivilegedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerPrivilegedMode

`func (o *ApplianceSettingsUpdate) SetDockerPrivilegedMode(v bool)`

SetDockerPrivilegedMode sets DockerPrivilegedMode field to given value.

### HasDockerPrivilegedMode

`func (o *ApplianceSettingsUpdate) HasDockerPrivilegedMode() bool`

HasDockerPrivilegedMode returns a boolean if a field has been set.

### GetPasswordMinLength

`func (o *ApplianceSettingsUpdate) GetPasswordMinLength() string`

GetPasswordMinLength returns the PasswordMinLength field if non-nil, zero value otherwise.

### GetPasswordMinLengthOk

`func (o *ApplianceSettingsUpdate) GetPasswordMinLengthOk() (*string, bool)`

GetPasswordMinLengthOk returns a tuple with the PasswordMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinLength

`func (o *ApplianceSettingsUpdate) SetPasswordMinLength(v string)`

SetPasswordMinLength sets PasswordMinLength field to given value.

### HasPasswordMinLength

`func (o *ApplianceSettingsUpdate) HasPasswordMinLength() bool`

HasPasswordMinLength returns a boolean if a field has been set.

### GetPasswordMinUpperCase

`func (o *ApplianceSettingsUpdate) GetPasswordMinUpperCase() string`

GetPasswordMinUpperCase returns the PasswordMinUpperCase field if non-nil, zero value otherwise.

### GetPasswordMinUpperCaseOk

`func (o *ApplianceSettingsUpdate) GetPasswordMinUpperCaseOk() (*string, bool)`

GetPasswordMinUpperCaseOk returns a tuple with the PasswordMinUpperCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinUpperCase

`func (o *ApplianceSettingsUpdate) SetPasswordMinUpperCase(v string)`

SetPasswordMinUpperCase sets PasswordMinUpperCase field to given value.

### HasPasswordMinUpperCase

`func (o *ApplianceSettingsUpdate) HasPasswordMinUpperCase() bool`

HasPasswordMinUpperCase returns a boolean if a field has been set.

### GetPasswordMinNumbers

`func (o *ApplianceSettingsUpdate) GetPasswordMinNumbers() string`

GetPasswordMinNumbers returns the PasswordMinNumbers field if non-nil, zero value otherwise.

### GetPasswordMinNumbersOk

`func (o *ApplianceSettingsUpdate) GetPasswordMinNumbersOk() (*string, bool)`

GetPasswordMinNumbersOk returns a tuple with the PasswordMinNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinNumbers

`func (o *ApplianceSettingsUpdate) SetPasswordMinNumbers(v string)`

SetPasswordMinNumbers sets PasswordMinNumbers field to given value.

### HasPasswordMinNumbers

`func (o *ApplianceSettingsUpdate) HasPasswordMinNumbers() bool`

HasPasswordMinNumbers returns a boolean if a field has been set.

### GetPasswordMinSymbols

`func (o *ApplianceSettingsUpdate) GetPasswordMinSymbols() string`

GetPasswordMinSymbols returns the PasswordMinSymbols field if non-nil, zero value otherwise.

### GetPasswordMinSymbolsOk

`func (o *ApplianceSettingsUpdate) GetPasswordMinSymbolsOk() (*string, bool)`

GetPasswordMinSymbolsOk returns a tuple with the PasswordMinSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinSymbols

`func (o *ApplianceSettingsUpdate) SetPasswordMinSymbols(v string)`

SetPasswordMinSymbols sets PasswordMinSymbols field to given value.

### HasPasswordMinSymbols

`func (o *ApplianceSettingsUpdate) HasPasswordMinSymbols() bool`

HasPasswordMinSymbols returns a boolean if a field has been set.

### GetUserBrowserSessionTimeout

`func (o *ApplianceSettingsUpdate) GetUserBrowserSessionTimeout() string`

GetUserBrowserSessionTimeout returns the UserBrowserSessionTimeout field if non-nil, zero value otherwise.

### GetUserBrowserSessionTimeoutOk

`func (o *ApplianceSettingsUpdate) GetUserBrowserSessionTimeoutOk() (*string, bool)`

GetUserBrowserSessionTimeoutOk returns a tuple with the UserBrowserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBrowserSessionTimeout

`func (o *ApplianceSettingsUpdate) SetUserBrowserSessionTimeout(v string)`

SetUserBrowserSessionTimeout sets UserBrowserSessionTimeout field to given value.

### HasUserBrowserSessionTimeout

`func (o *ApplianceSettingsUpdate) HasUserBrowserSessionTimeout() bool`

HasUserBrowserSessionTimeout returns a boolean if a field has been set.

### GetUserBrowserSessionWarning

`func (o *ApplianceSettingsUpdate) GetUserBrowserSessionWarning() string`

GetUserBrowserSessionWarning returns the UserBrowserSessionWarning field if non-nil, zero value otherwise.

### GetUserBrowserSessionWarningOk

`func (o *ApplianceSettingsUpdate) GetUserBrowserSessionWarningOk() (*string, bool)`

GetUserBrowserSessionWarningOk returns a tuple with the UserBrowserSessionWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBrowserSessionWarning

`func (o *ApplianceSettingsUpdate) SetUserBrowserSessionWarning(v string)`

SetUserBrowserSessionWarning sets UserBrowserSessionWarning field to given value.

### HasUserBrowserSessionWarning

`func (o *ApplianceSettingsUpdate) HasUserBrowserSessionWarning() bool`

HasUserBrowserSessionWarning returns a boolean if a field has been set.

### GetExpirePwdDays

`func (o *ApplianceSettingsUpdate) GetExpirePwdDays() int64`

GetExpirePwdDays returns the ExpirePwdDays field if non-nil, zero value otherwise.

### GetExpirePwdDaysOk

`func (o *ApplianceSettingsUpdate) GetExpirePwdDaysOk() (*int64, bool)`

GetExpirePwdDaysOk returns a tuple with the ExpirePwdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirePwdDays

`func (o *ApplianceSettingsUpdate) SetExpirePwdDays(v int64)`

SetExpirePwdDays sets ExpirePwdDays field to given value.

### HasExpirePwdDays

`func (o *ApplianceSettingsUpdate) HasExpirePwdDays() bool`

HasExpirePwdDays returns a boolean if a field has been set.

### GetDisableAfterAttempts

`func (o *ApplianceSettingsUpdate) GetDisableAfterAttempts() int64`

GetDisableAfterAttempts returns the DisableAfterAttempts field if non-nil, zero value otherwise.

### GetDisableAfterAttemptsOk

`func (o *ApplianceSettingsUpdate) GetDisableAfterAttemptsOk() (*int64, bool)`

GetDisableAfterAttemptsOk returns a tuple with the DisableAfterAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAfterAttempts

`func (o *ApplianceSettingsUpdate) SetDisableAfterAttempts(v int64)`

SetDisableAfterAttempts sets DisableAfterAttempts field to given value.

### HasDisableAfterAttempts

`func (o *ApplianceSettingsUpdate) HasDisableAfterAttempts() bool`

HasDisableAfterAttempts returns a boolean if a field has been set.

### GetDisableAfterDaysInactive

`func (o *ApplianceSettingsUpdate) GetDisableAfterDaysInactive() int64`

GetDisableAfterDaysInactive returns the DisableAfterDaysInactive field if non-nil, zero value otherwise.

### GetDisableAfterDaysInactiveOk

`func (o *ApplianceSettingsUpdate) GetDisableAfterDaysInactiveOk() (*int64, bool)`

GetDisableAfterDaysInactiveOk returns a tuple with the DisableAfterDaysInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAfterDaysInactive

`func (o *ApplianceSettingsUpdate) SetDisableAfterDaysInactive(v int64)`

SetDisableAfterDaysInactive sets DisableAfterDaysInactive field to given value.

### HasDisableAfterDaysInactive

`func (o *ApplianceSettingsUpdate) HasDisableAfterDaysInactive() bool`

HasDisableAfterDaysInactive returns a boolean if a field has been set.

### GetWarnUserDaysBefore

`func (o *ApplianceSettingsUpdate) GetWarnUserDaysBefore() int64`

GetWarnUserDaysBefore returns the WarnUserDaysBefore field if non-nil, zero value otherwise.

### GetWarnUserDaysBeforeOk

`func (o *ApplianceSettingsUpdate) GetWarnUserDaysBeforeOk() (*int64, bool)`

GetWarnUserDaysBeforeOk returns a tuple with the WarnUserDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnUserDaysBefore

`func (o *ApplianceSettingsUpdate) SetWarnUserDaysBefore(v int64)`

SetWarnUserDaysBefore sets WarnUserDaysBefore field to given value.

### HasWarnUserDaysBefore

`func (o *ApplianceSettingsUpdate) HasWarnUserDaysBefore() bool`

HasWarnUserDaysBefore returns a boolean if a field has been set.

### GetSmtpMailFrom

`func (o *ApplianceSettingsUpdate) GetSmtpMailFrom() string`

GetSmtpMailFrom returns the SmtpMailFrom field if non-nil, zero value otherwise.

### GetSmtpMailFromOk

`func (o *ApplianceSettingsUpdate) GetSmtpMailFromOk() (*string, bool)`

GetSmtpMailFromOk returns a tuple with the SmtpMailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpMailFrom

`func (o *ApplianceSettingsUpdate) SetSmtpMailFrom(v string)`

SetSmtpMailFrom sets SmtpMailFrom field to given value.

### HasSmtpMailFrom

`func (o *ApplianceSettingsUpdate) HasSmtpMailFrom() bool`

HasSmtpMailFrom returns a boolean if a field has been set.

### GetSmtpServer

`func (o *ApplianceSettingsUpdate) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *ApplianceSettingsUpdate) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *ApplianceSettingsUpdate) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *ApplianceSettingsUpdate) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetSmtpPort

`func (o *ApplianceSettingsUpdate) GetSmtpPort() int64`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *ApplianceSettingsUpdate) GetSmtpPortOk() (*int64, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *ApplianceSettingsUpdate) SetSmtpPort(v int64)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *ApplianceSettingsUpdate) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpSSL

`func (o *ApplianceSettingsUpdate) GetSmtpSSL() bool`

GetSmtpSSL returns the SmtpSSL field if non-nil, zero value otherwise.

### GetSmtpSSLOk

`func (o *ApplianceSettingsUpdate) GetSmtpSSLOk() (*bool, bool)`

GetSmtpSSLOk returns a tuple with the SmtpSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSSL

`func (o *ApplianceSettingsUpdate) SetSmtpSSL(v bool)`

SetSmtpSSL sets SmtpSSL field to given value.

### HasSmtpSSL

`func (o *ApplianceSettingsUpdate) HasSmtpSSL() bool`

HasSmtpSSL returns a boolean if a field has been set.

### GetSmtpTLS

`func (o *ApplianceSettingsUpdate) GetSmtpTLS() bool`

GetSmtpTLS returns the SmtpTLS field if non-nil, zero value otherwise.

### GetSmtpTLSOk

`func (o *ApplianceSettingsUpdate) GetSmtpTLSOk() (*bool, bool)`

GetSmtpTLSOk returns a tuple with the SmtpTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpTLS

`func (o *ApplianceSettingsUpdate) SetSmtpTLS(v bool)`

SetSmtpTLS sets SmtpTLS field to given value.

### HasSmtpTLS

`func (o *ApplianceSettingsUpdate) HasSmtpTLS() bool`

HasSmtpTLS returns a boolean if a field has been set.

### GetSmtpUser

`func (o *ApplianceSettingsUpdate) GetSmtpUser() string`

GetSmtpUser returns the SmtpUser field if non-nil, zero value otherwise.

### GetSmtpUserOk

`func (o *ApplianceSettingsUpdate) GetSmtpUserOk() (*string, bool)`

GetSmtpUserOk returns a tuple with the SmtpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpUser

`func (o *ApplianceSettingsUpdate) SetSmtpUser(v string)`

SetSmtpUser sets SmtpUser field to given value.

### HasSmtpUser

`func (o *ApplianceSettingsUpdate) HasSmtpUser() bool`

HasSmtpUser returns a boolean if a field has been set.

### GetSmtpPassword

`func (o *ApplianceSettingsUpdate) GetSmtpPassword() string`

GetSmtpPassword returns the SmtpPassword field if non-nil, zero value otherwise.

### GetSmtpPasswordOk

`func (o *ApplianceSettingsUpdate) GetSmtpPasswordOk() (*string, bool)`

GetSmtpPasswordOk returns a tuple with the SmtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPassword

`func (o *ApplianceSettingsUpdate) SetSmtpPassword(v string)`

SetSmtpPassword sets SmtpPassword field to given value.

### HasSmtpPassword

`func (o *ApplianceSettingsUpdate) HasSmtpPassword() bool`

HasSmtpPassword returns a boolean if a field has been set.

### GetProxyHost

`func (o *ApplianceSettingsUpdate) GetProxyHost() string`

GetProxyHost returns the ProxyHost field if non-nil, zero value otherwise.

### GetProxyHostOk

`func (o *ApplianceSettingsUpdate) GetProxyHostOk() (*string, bool)`

GetProxyHostOk returns a tuple with the ProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHost

`func (o *ApplianceSettingsUpdate) SetProxyHost(v string)`

SetProxyHost sets ProxyHost field to given value.

### HasProxyHost

`func (o *ApplianceSettingsUpdate) HasProxyHost() bool`

HasProxyHost returns a boolean if a field has been set.

### SetProxyHostNil

`func (o *ApplianceSettingsUpdate) SetProxyHostNil(b bool)`

 SetProxyHostNil sets the value for ProxyHost to be an explicit nil

### UnsetProxyHost
`func (o *ApplianceSettingsUpdate) UnsetProxyHost()`

UnsetProxyHost ensures that no value is present for ProxyHost, not even an explicit nil
### GetProxyPort

`func (o *ApplianceSettingsUpdate) GetProxyPort() string`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *ApplianceSettingsUpdate) GetProxyPortOk() (*string, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *ApplianceSettingsUpdate) SetProxyPort(v string)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *ApplianceSettingsUpdate) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### SetProxyPortNil

`func (o *ApplianceSettingsUpdate) SetProxyPortNil(b bool)`

 SetProxyPortNil sets the value for ProxyPort to be an explicit nil

### UnsetProxyPort
`func (o *ApplianceSettingsUpdate) UnsetProxyPort()`

UnsetProxyPort ensures that no value is present for ProxyPort, not even an explicit nil
### GetProxyUser

`func (o *ApplianceSettingsUpdate) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *ApplianceSettingsUpdate) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *ApplianceSettingsUpdate) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *ApplianceSettingsUpdate) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### GetProxyPassword

`func (o *ApplianceSettingsUpdate) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *ApplianceSettingsUpdate) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *ApplianceSettingsUpdate) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *ApplianceSettingsUpdate) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetProxyDomain

`func (o *ApplianceSettingsUpdate) GetProxyDomain() string`

GetProxyDomain returns the ProxyDomain field if non-nil, zero value otherwise.

### GetProxyDomainOk

`func (o *ApplianceSettingsUpdate) GetProxyDomainOk() (*string, bool)`

GetProxyDomainOk returns a tuple with the ProxyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyDomain

`func (o *ApplianceSettingsUpdate) SetProxyDomain(v string)`

SetProxyDomain sets ProxyDomain field to given value.

### HasProxyDomain

`func (o *ApplianceSettingsUpdate) HasProxyDomain() bool`

HasProxyDomain returns a boolean if a field has been set.

### SetProxyDomainNil

`func (o *ApplianceSettingsUpdate) SetProxyDomainNil(b bool)`

 SetProxyDomainNil sets the value for ProxyDomain to be an explicit nil

### UnsetProxyDomain
`func (o *ApplianceSettingsUpdate) UnsetProxyDomain()`

UnsetProxyDomain ensures that no value is present for ProxyDomain, not even an explicit nil
### GetProxyWorkstation

`func (o *ApplianceSettingsUpdate) GetProxyWorkstation() string`

GetProxyWorkstation returns the ProxyWorkstation field if non-nil, zero value otherwise.

### GetProxyWorkstationOk

`func (o *ApplianceSettingsUpdate) GetProxyWorkstationOk() (*string, bool)`

GetProxyWorkstationOk returns a tuple with the ProxyWorkstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyWorkstation

`func (o *ApplianceSettingsUpdate) SetProxyWorkstation(v string)`

SetProxyWorkstation sets ProxyWorkstation field to given value.

### HasProxyWorkstation

`func (o *ApplianceSettingsUpdate) HasProxyWorkstation() bool`

HasProxyWorkstation returns a boolean if a field has been set.

### SetProxyWorkstationNil

`func (o *ApplianceSettingsUpdate) SetProxyWorkstationNil(b bool)`

 SetProxyWorkstationNil sets the value for ProxyWorkstation to be an explicit nil

### UnsetProxyWorkstation
`func (o *ApplianceSettingsUpdate) UnsetProxyWorkstation()`

UnsetProxyWorkstation ensures that no value is present for ProxyWorkstation, not even an explicit nil
### GetCurrencyProvider

`func (o *ApplianceSettingsUpdate) GetCurrencyProvider() string`

GetCurrencyProvider returns the CurrencyProvider field if non-nil, zero value otherwise.

### GetCurrencyProviderOk

`func (o *ApplianceSettingsUpdate) GetCurrencyProviderOk() (*string, bool)`

GetCurrencyProviderOk returns a tuple with the CurrencyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyProvider

`func (o *ApplianceSettingsUpdate) SetCurrencyProvider(v string)`

SetCurrencyProvider sets CurrencyProvider field to given value.

### HasCurrencyProvider

`func (o *ApplianceSettingsUpdate) HasCurrencyProvider() bool`

HasCurrencyProvider returns a boolean if a field has been set.

### GetCurrencyKey

`func (o *ApplianceSettingsUpdate) GetCurrencyKey() string`

GetCurrencyKey returns the CurrencyKey field if non-nil, zero value otherwise.

### GetCurrencyKeyOk

`func (o *ApplianceSettingsUpdate) GetCurrencyKeyOk() (*string, bool)`

GetCurrencyKeyOk returns a tuple with the CurrencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyKey

`func (o *ApplianceSettingsUpdate) SetCurrencyKey(v string)`

SetCurrencyKey sets CurrencyKey field to given value.

### HasCurrencyKey

`func (o *ApplianceSettingsUpdate) HasCurrencyKey() bool`

HasCurrencyKey returns a boolean if a field has been set.

### SetCurrencyKeyNil

`func (o *ApplianceSettingsUpdate) SetCurrencyKeyNil(b bool)`

 SetCurrencyKeyNil sets the value for CurrencyKey to be an explicit nil

### UnsetCurrencyKey
`func (o *ApplianceSettingsUpdate) UnsetCurrencyKey()`

UnsetCurrencyKey ensures that no value is present for CurrencyKey, not even an explicit nil
### GetEnableAllZoneTypes

`func (o *ApplianceSettingsUpdate) GetEnableAllZoneTypes() bool`

GetEnableAllZoneTypes returns the EnableAllZoneTypes field if non-nil, zero value otherwise.

### GetEnableAllZoneTypesOk

`func (o *ApplianceSettingsUpdate) GetEnableAllZoneTypesOk() (*bool, bool)`

GetEnableAllZoneTypesOk returns a tuple with the EnableAllZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllZoneTypes

`func (o *ApplianceSettingsUpdate) SetEnableAllZoneTypes(v bool)`

SetEnableAllZoneTypes sets EnableAllZoneTypes field to given value.

### HasEnableAllZoneTypes

`func (o *ApplianceSettingsUpdate) HasEnableAllZoneTypes() bool`

HasEnableAllZoneTypes returns a boolean if a field has been set.

### GetEnableZoneTypes

`func (o *ApplianceSettingsUpdate) GetEnableZoneTypes() []int64`

GetEnableZoneTypes returns the EnableZoneTypes field if non-nil, zero value otherwise.

### GetEnableZoneTypesOk

`func (o *ApplianceSettingsUpdate) GetEnableZoneTypesOk() (*[]int64, bool)`

GetEnableZoneTypesOk returns a tuple with the EnableZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableZoneTypes

`func (o *ApplianceSettingsUpdate) SetEnableZoneTypes(v []int64)`

SetEnableZoneTypes sets EnableZoneTypes field to given value.

### HasEnableZoneTypes

`func (o *ApplianceSettingsUpdate) HasEnableZoneTypes() bool`

HasEnableZoneTypes returns a boolean if a field has been set.

### GetDisableZoneTypes

`func (o *ApplianceSettingsUpdate) GetDisableZoneTypes() []int64`

GetDisableZoneTypes returns the DisableZoneTypes field if non-nil, zero value otherwise.

### GetDisableZoneTypesOk

`func (o *ApplianceSettingsUpdate) GetDisableZoneTypesOk() (*[]int64, bool)`

GetDisableZoneTypesOk returns a tuple with the DisableZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableZoneTypes

`func (o *ApplianceSettingsUpdate) SetDisableZoneTypes(v []int64)`

SetDisableZoneTypes sets DisableZoneTypes field to given value.

### HasDisableZoneTypes

`func (o *ApplianceSettingsUpdate) HasDisableZoneTypes() bool`

HasDisableZoneTypes returns a boolean if a field has been set.

### GetDisableAllZoneTypes

`func (o *ApplianceSettingsUpdate) GetDisableAllZoneTypes() bool`

GetDisableAllZoneTypes returns the DisableAllZoneTypes field if non-nil, zero value otherwise.

### GetDisableAllZoneTypesOk

`func (o *ApplianceSettingsUpdate) GetDisableAllZoneTypesOk() (*bool, bool)`

GetDisableAllZoneTypesOk returns a tuple with the DisableAllZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAllZoneTypes

`func (o *ApplianceSettingsUpdate) SetDisableAllZoneTypes(v bool)`

SetDisableAllZoneTypes sets DisableAllZoneTypes field to given value.

### HasDisableAllZoneTypes

`func (o *ApplianceSettingsUpdate) HasDisableAllZoneTypes() bool`

HasDisableAllZoneTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


