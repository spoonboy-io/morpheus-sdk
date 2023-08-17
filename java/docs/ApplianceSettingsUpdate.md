

# ApplianceSettingsUpdate


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**applianceUrl** | **String** | Appliance URL |  [optional] |
|**internalApplianceUrl** | **String** | Internal Appliance URL (PXE) |  [optional] |
|**corsAllowed** | **String** | API Allowed Origins |  [optional] |
|**registrationEnabled** | **Boolean** | Registration enabled (true, false) |  [optional] |
|**defaultRoleId** | **Long** | Default tenant role ID |  [optional] |
|**defaultUserRoleId** | **Long** | Default user role ID |  [optional] |
|**dockerPrivilegedMode** | **Boolean** | Docker privileged mode (true, false) |  [optional] |
|**passwordMinLength** | **String** | Min Password Length |  [optional] |
|**passwordMinUpperCase** | **String** | Min Password Uppercase |  [optional] |
|**passwordMinNumbers** | **String** | Min Password Numbers |  [optional] |
|**passwordMinSymbols** | **String** | Min Password Symbols |  [optional] |
|**userBrowserSessionTimeout** | **String** | User Browser Session Timeout (Minutes) |  [optional] |
|**userBrowserSessionWarning** | **String** | User Browser Session Warning (Minutes) |  [optional] |
|**expirePwdDays** | **Long** | Expire password after days. Setting to 0 disabled this feature |  [optional] |
|**disableAfterAttempts** | **Long** | Disable user after number of attempts. Set to 0 to disable this feature |  [optional] |
|**disableAfterDaysInactive** | **Long** | Disable user if inactive for specified days. Set to 0 to disable this feature |  [optional] |
|**warnUserDaysBefore** | **Long** | Send warning email number of days in advance before deactivating. Set to 0 to disable this feature |  [optional] |
|**smtpMailFrom** | **String** | From email address |  [optional] |
|**smtpServer** | **String** | SMTP server / host |  [optional] |
|**smtpPort** | **Long** | SMTP port |  [optional] |
|**smtpSSL** | **Boolean** | Use SSL for SMTP connection |  [optional] |
|**smtpTLS** | **Boolean** | Use TLS for SMTP connections |  [optional] |
|**smtpUser** | **String** | SMTP username |  [optional] |
|**smtpPassword** | **String** | SMTP password |  [optional] |
|**proxyHost** | **String** | Proxy host |  [optional] |
|**proxyPort** | **String** | Proxy port |  [optional] |
|**proxyUser** | **String** | Proxy username |  [optional] |
|**proxyPassword** | **String** | Proxy password |  [optional] |
|**proxyDomain** | **String** | Proxy domain |  [optional] |
|**proxyWorkstation** | **String** | Proxy workstation |  [optional] |
|**currencyProvider** | **String** | Currency provider |  [optional] |
|**currencyKey** | **String** | Currency provider API key |  [optional] |
|**enableAllZoneTypes** | **Boolean** | Set all cloud types enabled status on, overrides enableZoneTypes and disableZoneTypes parameters |  [optional] |
|**enableZoneTypes** | **List&lt;Long&gt;** | List of cloud type IDs to set enabled status on |  [optional] |
|**disableZoneTypes** | **List&lt;Long&gt;** | List of cloud type IDs to set enabled status off |  [optional] |
|**disableAllZoneTypes** | **Boolean** | Set all cloud types enabled status off, can be used in conjunction with enableZoneTypes |  [optional] |



