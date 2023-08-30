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

.PARAMETER License
No description available.
.OUTPUTS

GetProvisioningLicense200Response<PSCustomObject>
#>

function Initialize-GetProvisioningLicense200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${License}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetProvisioningLicense200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "license" = ${License}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetProvisioningLicense200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetProvisioningLicense200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetProvisioningLicense200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetProvisioningLicense200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetProvisioningLicense200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetProvisioningLicense200Response
        $AllProperties = ("license")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "license"))) { #optional property not found
            $License = $null
        } else {
            $License = $JsonParameters.PSobject.Properties["license"].value
        }

        $PSO = [PSCustomObject]@{
            "license" = ${License}
        }

        return $PSO
    }

}
