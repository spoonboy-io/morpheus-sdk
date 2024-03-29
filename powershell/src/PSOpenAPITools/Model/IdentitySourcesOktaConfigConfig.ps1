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

.PARAMETER Url
No description available.
.PARAMETER AdministratorAPIToken
No description available.
.PARAMETER RequiredGroup
No description available.
.PARAMETER RequiredGroupId
No description available.
.PARAMETER AdministratorAPITokenHash
No description available.
.OUTPUTS

IdentitySourcesOktaConfigConfig<PSCustomObject>
#>

function Initialize-IdentitySourcesOktaConfigConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Url},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${AdministratorAPIToken},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RequiredGroup},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RequiredGroupId},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${AdministratorAPITokenHash}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => IdentitySourcesOktaConfigConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "url" = ${Url}
            "administratorAPIToken" = ${AdministratorAPIToken}
            "requiredGroup" = ${RequiredGroup}
            "requiredGroupId" = ${RequiredGroupId}
            "administratorAPITokenHash" = ${AdministratorAPITokenHash}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to IdentitySourcesOktaConfigConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to IdentitySourcesOktaConfigConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

IdentitySourcesOktaConfigConfig<PSCustomObject>
#>
function ConvertFrom-JsonToIdentitySourcesOktaConfigConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => IdentitySourcesOktaConfigConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in IdentitySourcesOktaConfigConfig
        $AllProperties = ("url", "administratorAPIToken", "requiredGroup", "requiredGroupId", "administratorAPITokenHash")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "url"))) { #optional property not found
            $Url = $null
        } else {
            $Url = $JsonParameters.PSobject.Properties["url"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "administratorAPIToken"))) { #optional property not found
            $AdministratorAPIToken = $null
        } else {
            $AdministratorAPIToken = $JsonParameters.PSobject.Properties["administratorAPIToken"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "requiredGroup"))) { #optional property not found
            $RequiredGroup = $null
        } else {
            $RequiredGroup = $JsonParameters.PSobject.Properties["requiredGroup"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "requiredGroupId"))) { #optional property not found
            $RequiredGroupId = $null
        } else {
            $RequiredGroupId = $JsonParameters.PSobject.Properties["requiredGroupId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "administratorAPITokenHash"))) { #optional property not found
            $AdministratorAPITokenHash = $null
        } else {
            $AdministratorAPITokenHash = $JsonParameters.PSobject.Properties["administratorAPITokenHash"].value
        }

        $PSO = [PSCustomObject]@{
            "url" = ${Url}
            "administratorAPIToken" = ${AdministratorAPIToken}
            "requiredGroup" = ${RequiredGroup}
            "requiredGroupId" = ${RequiredGroupId}
            "administratorAPITokenHash" = ${AdministratorAPITokenHash}
        }

        return $PSO
    }

}

