

# UserSourceCreateCustomExternal

Custom External Configuration
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**loginUrl** | **String** | External Login URL |  [optional]
**doNotIncludeSAMLRequest** | **Boolean** | Do not include SAMLRequest |  [optional]
**logout** | **String** | External Logout URL |  [optional]
**encryptionAlgo** | [**EncryptionAlgoEnum**](#EncryptionAlgoEnum) | Encryption Algorithm |  [optional]
**encryptionKey** | **String** | Encryption Key |  [optional]



## Enum: EncryptionAlgoEnum

Name | Value
---- | -----
NONE | &quot;NONE&quot;
AES | &quot;AES&quot;
DES | &quot;DES&quot;
DESEDE | &quot;DESede&quot;
HMACSHA1 | &quot;HmacSHA1&quot;
HMACSHA256 | &quot;HmacSHA256&quot;



