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

.PARAMETER ClientSecret
No description available.
.PARAMETER ClientId
No description available.
.PARAMETER ClientAuth
No description available.
.PARAMETER Scope
No description available.
.PARAMETER GrantType
No description available.
.PARAMETER AccessTokenUrl
No description available.
.PARAMETER ClientSecretHash
No description available.
.OUTPUTS

CredentialConfig<PSCustomObject>
#>

function Initialize-CredentialConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ClientSecret},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ClientId},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ClientAuth},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Scope},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GrantType},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${AccessTokenUrl},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ClientSecretHash}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CredentialConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "clientSecret" = ${ClientSecret}
            "clientId" = ${ClientId}
            "clientAuth" = ${ClientAuth}
            "scope" = ${Scope}
            "grantType" = ${GrantType}
            "accessTokenUrl" = ${AccessTokenUrl}
            "clientSecretHash" = ${ClientSecretHash}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CredentialConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to CredentialConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CredentialConfig<PSCustomObject>
#>
function ConvertFrom-JsonToCredentialConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CredentialConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CredentialConfig
        $AllProperties = ("clientSecret", "clientId", "clientAuth", "scope", "grantType", "accessTokenUrl", "clientSecretHash")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clientSecret"))) { #optional property not found
            $ClientSecret = $null
        } else {
            $ClientSecret = $JsonParameters.PSobject.Properties["clientSecret"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clientId"))) { #optional property not found
            $ClientId = $null
        } else {
            $ClientId = $JsonParameters.PSobject.Properties["clientId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clientAuth"))) { #optional property not found
            $ClientAuth = $null
        } else {
            $ClientAuth = $JsonParameters.PSobject.Properties["clientAuth"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scope"))) { #optional property not found
            $Scope = $null
        } else {
            $Scope = $JsonParameters.PSobject.Properties["scope"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "grantType"))) { #optional property not found
            $GrantType = $null
        } else {
            $GrantType = $JsonParameters.PSobject.Properties["grantType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "accessTokenUrl"))) { #optional property not found
            $AccessTokenUrl = $null
        } else {
            $AccessTokenUrl = $JsonParameters.PSobject.Properties["accessTokenUrl"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clientSecretHash"))) { #optional property not found
            $ClientSecretHash = $null
        } else {
            $ClientSecretHash = $JsonParameters.PSobject.Properties["clientSecretHash"].value
        }

        $PSO = [PSCustomObject]@{
            "clientSecret" = ${ClientSecret}
            "clientId" = ${ClientId}
            "clientAuth" = ${ClientAuth}
            "scope" = ${Scope}
            "grantType" = ${GrantType}
            "accessTokenUrl" = ${AccessTokenUrl}
            "clientSecretHash" = ${ClientSecretHash}
        }

        return $PSO
    }

}

