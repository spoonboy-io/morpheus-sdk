

# UserSourceCreateCustomApi

Custom API Configuration
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**endpoint** | **String** | API Endpoint |  [optional]
**apiStyle** | [**ApiStyleEnum**](#ApiStyleEnum) | API Style |  [optional]
**encryptionAlgo** | [**EncryptionAlgoEnum**](#EncryptionAlgoEnum) | Encryption Algorithm |  [optional]
**encryptionKey** | **String** | Encryption Key |  [optional]



## Enum: ApiStyleEnum

Name | Value
---- | -----
FORM_URL_ENCODED_GET_ | &quot;Form URL Encoded [GET]&quot;
FORM_URL_ENCODED_POST_ | &quot;Form URL Encoded [POST]&quot;
JSON_POST_ | &quot;JSON [POST]&quot;
XML_POST_ | &quot;XML [POST]&quot;
HTTP_BASIC_GET_ | &quot;HTTP Basic [GET]&quot;



## Enum: EncryptionAlgoEnum

Name | Value
---- | -----
NONE | &quot;NONE&quot;
AES | &quot;AES&quot;
DES | &quot;DES&quot;
DESEDE | &quot;DESede&quot;
HMACSHA1 | &quot;HmacSHA1&quot;
HMACSHA256 | &quot;HmacSHA256&quot;



