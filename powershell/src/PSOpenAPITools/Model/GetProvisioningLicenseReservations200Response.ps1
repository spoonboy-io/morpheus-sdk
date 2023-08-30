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

.PARAMETER Licenses
No description available.
.OUTPUTS

GetProvisioningLicenseReservations200Response<PSCustomObject>
#>

function Initialize-GetProvisioningLicenseReservations200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Licenses}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetProvisioningLicenseReservations200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "licenses" = ${Licenses}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetProvisioningLicenseReservations200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetProvisioningLicenseReservations200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetProvisioningLicenseReservations200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetProvisioningLicenseReservations200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetProvisioningLicenseReservations200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetProvisioningLicenseReservations200Response
        $AllProperties = ("licenses")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "licenses"))) { #optional property not found
            $Licenses = $null
        } else {
            $Licenses = $JsonParameters.PSobject.Properties["licenses"].value
        }

        $PSO = [PSCustomObject]@{
            "licenses" = ${Licenses}
        }

        return $PSO
    }

}

