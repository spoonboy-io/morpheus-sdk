

# CredentialOauth2ConfigConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**grantType** | [**GrantTypeEnum**](#GrantTypeEnum) | OAuth 2.0 grant type | 
**accessTokenUrl** | **String** | Token endpoint | 
**clientId** | **String** | Client ID |  [optional]
**clientSecret** | **String** | Client Secret |  [optional]
**scope** | **String** | Scope |  [optional]
**clientAuth** | [**ClientAuthEnum**](#ClientAuthEnum) | Client Authentication Method | 



## Enum: GrantTypeEnum

Name | Value
---- | -----
PASSWORD | &quot;password&quot;
CLIENT_CREDENTIALS | &quot;client_credentials&quot;



## Enum: ClientAuthEnum

Name | Value
---- | -----
BODY | &quot;body&quot;
BASIC_AUTH | &quot;basic-auth&quot;



