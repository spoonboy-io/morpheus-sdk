#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER GrantType
OAuth 2.0 grant type
.PARAMETER AccessTokenUrl
Token endpoint
.PARAMETER ClientId
Client ID
.PARAMETER ClientSecret
Client Secret
.PARAMETER Scope
Scope
.PARAMETER ClientAuth
Client Authentication Method
.OUTPUTS

CredentialOauth2ConfigConfig<PSCustomObject>
#>

function Initialize-CredentialOauth2ConfigConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("password", "client_credentials")]
        [String]
        ${GrantType},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${AccessTokenUrl},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ClientId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ClientSecret},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Scope},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("body", "basic-auth")]
        [String]
        ${ClientAuth}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CredentialOauth2ConfigConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $GrantType) {
            throw "invalid value for 'GrantType', 'GrantType' cannot be null."
        }

        if ($null -eq $AccessTokenUrl) {
            throw "invalid value for 'AccessTokenUrl', 'AccessTokenUrl' cannot be null."
        }

        if ($null -eq $ClientAuth) {
            throw "invalid value for 'ClientAuth', 'ClientAuth' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "grantType" = ${GrantType}
            "accessTokenUrl" = ${AccessTokenUrl}
            "clientId" = ${ClientId}
            "clientSecret" = ${ClientSecret}
            "scope" = ${Scope}
            "clientAuth" = ${ClientAuth}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CredentialOauth2ConfigConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to CredentialOauth2ConfigConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CredentialOauth2ConfigConfig<PSCustomObject>
#>
function ConvertFrom-JsonToCredentialOauth2ConfigConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CredentialOauth2ConfigConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CredentialOauth2ConfigConfig
        $AllProperties = ("grantType", "accessTokenUrl", "clientId", "clientSecret", "scope", "clientAuth")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'grantType' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "grantType"))) {
            throw "Error! JSON cannot be serialized due to the required property 'grantType' missing."
        } else {
            $GrantType = $JsonParameters.PSobject.Properties["grantType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "accessTokenUrl"))) {
            throw "Error! JSON cannot be serialized due to the required property 'accessTokenUrl' missing."
        } else {
            $AccessTokenUrl = $JsonParameters.PSobject.Properties["accessTokenUrl"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clientAuth"))) {
            throw "Error! JSON cannot be serialized due to the required property 'clientAuth' missing."
        } else {
            $ClientAuth = $JsonParameters.PSobject.Properties["clientAuth"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clientId"))) { #optional property not found
            $ClientId = $null
        } else {
            $ClientId = $JsonParameters.PSobject.Properties["clientId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clientSecret"))) { #optional property not found
            $ClientSecret = $null
        } else {
            $ClientSecret = $JsonParameters.PSobject.Properties["clientSecret"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scope"))) { #optional property not found
            $Scope = $null
        } else {
            $Scope = $JsonParameters.PSobject.Properties["scope"].value
        }

        $PSO = [PSCustomObject]@{
            "grantType" = ${GrantType}
            "accessTokenUrl" = ${AccessTokenUrl}
            "clientId" = ${ClientId}
            "clientSecret" = ${ClientSecret}
            "scope" = ${Scope}
            "clientAuth" = ${ClientAuth}
        }

        return $PSO
    }

}

