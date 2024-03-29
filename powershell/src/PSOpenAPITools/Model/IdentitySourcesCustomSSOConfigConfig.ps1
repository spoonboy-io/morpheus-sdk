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

.PARAMETER Noninteractive
No description available.
.PARAMETER LoginUrl
No description available.
.PARAMETER LogoutUrl
No description available.
.PARAMETER EncryptionAlgo
No description available.
.PARAMETER EncryptionKey
No description available.
.PARAMETER RequiredAttributeValue
No description available.
.OUTPUTS

IdentitySourcesCustomSSOConfigConfig<PSCustomObject>
#>

function Initialize-IdentitySourcesCustomSSOConfigConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Noninteractive},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LoginUrl},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LogoutUrl},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${EncryptionAlgo},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${EncryptionKey},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RequiredAttributeValue}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => IdentitySourcesCustomSSOConfigConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "noninteractive" = ${Noninteractive}
            "loginUrl" = ${LoginUrl}
            "logoutUrl" = ${LogoutUrl}
            "encryptionAlgo" = ${EncryptionAlgo}
            "encryptionKey" = ${EncryptionKey}
            "requiredAttributeValue" = ${RequiredAttributeValue}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to IdentitySourcesCustomSSOConfigConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to IdentitySourcesCustomSSOConfigConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

IdentitySourcesCustomSSOConfigConfig<PSCustomObject>
#>
function ConvertFrom-JsonToIdentitySourcesCustomSSOConfigConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => IdentitySourcesCustomSSOConfigConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in IdentitySourcesCustomSSOConfigConfig
        $AllProperties = ("noninteractive", "loginUrl", "logoutUrl", "encryptionAlgo", "encryptionKey", "requiredAttributeValue")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "noninteractive"))) { #optional property not found
            $Noninteractive = $null
        } else {
            $Noninteractive = $JsonParameters.PSobject.Properties["noninteractive"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "loginUrl"))) { #optional property not found
            $LoginUrl = $null
        } else {
            $LoginUrl = $JsonParameters.PSobject.Properties["loginUrl"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "logoutUrl"))) { #optional property not found
            $LogoutUrl = $null
        } else {
            $LogoutUrl = $JsonParameters.PSobject.Properties["logoutUrl"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "encryptionAlgo"))) { #optional property not found
            $EncryptionAlgo = $null
        } else {
            $EncryptionAlgo = $JsonParameters.PSobject.Properties["encryptionAlgo"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "encryptionKey"))) { #optional property not found
            $EncryptionKey = $null
        } else {
            $EncryptionKey = $JsonParameters.PSobject.Properties["encryptionKey"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "requiredAttributeValue"))) { #optional property not found
            $RequiredAttributeValue = $null
        } else {
            $RequiredAttributeValue = $JsonParameters.PSobject.Properties["requiredAttributeValue"].value
        }

        $PSO = [PSCustomObject]@{
            "noninteractive" = ${Noninteractive}
            "loginUrl" = ${LoginUrl}
            "logoutUrl" = ${LogoutUrl}
            "encryptionAlgo" = ${EncryptionAlgo}
            "encryptionKey" = ${EncryptionKey}
            "requiredAttributeValue" = ${RequiredAttributeValue}
        }

        return $PSO
    }

}

