# ApplianceSettingsUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**appliance_url** | **str** | Appliance URL | [optional] 
**internal_appliance_url** | **str, none_type** | Internal Appliance URL (PXE) | [optional] 
**cors_allowed** | **str, none_type** | API Allowed Origins | [optional] 
**registration_enabled** | **bool** | Registration enabled (true, false) | [optional] 
**default_role_id** | **int** | Default tenant role ID | [optional] 
**default_user_role_id** | **int** | Default user role ID | [optional] 
**docker_privileged_mode** | **bool** | Docker privileged mode (true, false) | [optional] 
**password_min_length** | **str** | Min Password Length | [optional] 
**password_min_upper_case** | **str** | Min Password Uppercase | [optional] 
**password_min_numbers** | **str** | Min Password Numbers | [optional] 
**password_min_symbols** | **str** | Min Password Symbols | [optional] 
**user_browser_session_timeout** | **str** | User Browser Session Timeout (Minutes) | [optional] 
**user_browser_session_warning** | **str** | User Browser Session Warning (Minutes) | [optional] 
**expire_pwd_days** | **int** | Expire password after days. Setting to 0 disabled this feature | [optional] 
**disable_after_attempts** | **int** | Disable user after number of attempts. Set to 0 to disable this feature | [optional] 
**disable_after_days_inactive** | **int** | Disable user if inactive for specified days. Set to 0 to disable this feature | [optional] 
**warn_user_days_before** | **int** | Send warning email number of days in advance before deactivating. Set to 0 to disable this feature | [optional] 
**smtp_mail_from** | **str** | From email address | [optional] 
**smtp_server** | **str** | SMTP server / host | [optional] 
**smtp_port** | **int** | SMTP port | [optional] 
**smtp_ssl** | **bool** | Use SSL for SMTP connection | [optional] 
**smtp_tls** | **bool** | Use TLS for SMTP connections | [optional] 
**smtp_user** | **str** | SMTP username | [optional] 
**smtp_password** | **str** | SMTP password | [optional] 
**proxy_host** | **str, none_type** | Proxy host | [optional] 
**proxy_port** | **str, none_type** | Proxy port | [optional] 
**proxy_user** | **str** | Proxy username | [optional] 
**proxy_password** | **str** | Proxy password | [optional] 
**proxy_domain** | **str, none_type** | Proxy domain | [optional] 
**proxy_workstation** | **str, none_type** | Proxy workstation | [optional] 
**currency_provider** | **str** | Currency provider | [optional] 
**currency_key** | **str, none_type** | Currency provider API key | [optional] 
**enable_all_zone_types** | **bool** | Set all cloud types enabled status on, overrides enableZoneTypes and disableZoneTypes parameters | [optional] 
**enable_zone_types** | **[int]** | List of cloud type IDs to set enabled status on | [optional] 
**disable_zone_types** | **[int]** | List of cloud type IDs to set enabled status off | [optional] 
**disable_all_zone_types** | **bool** | Set all cloud types enabled status off, can be used in conjunction with enableZoneTypes | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


