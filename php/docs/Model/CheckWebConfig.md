# # CheckWebConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**web_method** | **string** | HTTP method to use for testing |
**web_url** | **string** | Web URL you wish to use to run a check on |
**ignore_ssl** | **bool** | Ignore SSL Errors | [optional] [default to false]
**check_user** | **string** | If you want to use HTTP Basic Authentication, populate this field with the username | [optional]
**check_password** | **string** | If you want to use HTTP basic Authentication, populate this field with the password | [optional]
**text_check_on** | **string** | Set value to &#x60;on&#x60; if you want to turn on text matching | [optional]
**web_text_match** | **string** | Set the string you want to look for in the page source | [optional]
**tunnel_on** | **string** | Set to on to turn on tunneling | [optional]
**ssh_host** | **string** | Hostname or IP address of the proxy host | [optional]
**ssh_port** | **int** | Port for SSH on the proxy host, defaults to 22 | [optional]
**ssh_user** | **string** | SSH user on the proxy host to login as | [optional]
**ssh_password** | **string** | Password for user, if not using key based authentication | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
