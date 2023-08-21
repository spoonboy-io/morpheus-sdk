# # ApplianceSettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**appliance_url** | **string** | Appliance URL | [optional]
**internal_appliance_url** | **string** | Internal Appliance URL (PXE) | [optional]
**cors_allowed** | **string** | API Allowed Origins | [optional]
**registration_enabled** | **bool** | Registration enabled (true, false) | [optional]
**default_role_id** | **int** | Default tenant role ID | [optional]
**default_user_role_id** | **int** | Default user role ID | [optional]
**docker_privileged_mode** | **bool** | Docker privileged mode (true, false) | [optional]
**password_min_length** | **string** | Min Password Length | [optional]
**password_min_upper_case** | **string** | Min Password Uppercase | [optional]
**password_min_numbers** | **string** | Min Password Numbers | [optional]
**password_min_symbols** | **string** | Min Password Symbols | [optional]
**user_browser_session_timeout** | **string** | User Browser Session Timeout (Minutes) | [optional]
**user_browser_session_warning** | **string** | User Browser Session Warning (Minutes) | [optional]
**expire_pwd_days** | **int** | Expire password after days. Setting to 0 disabled this feature | [optional]
**disable_after_attempts** | **int** | Disable user after number of attempts. Set to 0 to disable this feature | [optional]
**disable_after_days_inactive** | **int** | Disable user if inactive for specified days. Set to 0 to disable this feature | [optional]
**warn_user_days_before** | **int** | Send warning email number of days in advance before deactivating. Set to 0 to disable this feature | [optional]
**smtp_mail_from** | **string** | From email address | [optional]
**smtp_server** | **string** | SMTP server / host | [optional]
**smtp_port** | **int** | SMTP port | [optional]
**smtp_ssl** | **bool** | Use SSL for SMTP connection | [optional]
**smtp_tls** | **bool** | Use TLS for SMTP connections | [optional]
**smtp_user** | **string** | SMTP username | [optional]
**smtp_password** | **string** | SMTP password | [optional]
**proxy_host** | **string** | Proxy host | [optional]
**proxy_port** | **string** | Proxy port | [optional]
**proxy_user** | **string** | Proxy username | [optional]
**proxy_password** | **string** | Proxy password | [optional]
**proxy_domain** | **string** | Proxy domain | [optional]
**proxy_workstation** | **string** | Proxy workstation | [optional]
**currency_provider** | **string** | Currency provider | [optional]
**currency_key** | **string** | Currency provider API key | [optional]
**enable_all_zone_types** | **bool** | Set all cloud types enabled status on, overrides enableZoneTypes and disableZoneTypes parameters | [optional]
**enable_zone_types** | **int[]** | List of cloud type IDs to set enabled status on | [optional]
**disable_zone_types** | **int[]** | List of cloud type IDs to set enabled status off | [optional]
**disable_all_zone_types** | **bool** | Set all cloud types enabled status off, can be used in conjunction with enableZoneTypes | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
