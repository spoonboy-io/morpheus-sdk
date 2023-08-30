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

.PARAMETER VdiApp
No description available.
.OUTPUTS

AddVDIApps200ResponseAnyOf<PSCustomObject>
#>

function Initialize-AddVDIApps200ResponseAnyOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${VdiApp}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddVDIApps200ResponseAnyOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "vdiApp" = ${VdiApp}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddVDIApps200ResponseAnyOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddVDIApps200ResponseAnyOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddVDIApps200ResponseAnyOf<PSCustomObject>
#>
function ConvertFrom-JsonToAddVDIApps200ResponseAnyOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddVDIApps200ResponseAnyOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddVDIApps200ResponseAnyOf
        $AllProperties = ("vdiApp")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "vdiApp"))) { #optional property not found
            $VdiApp = $null
        } else {
            $VdiApp = $JsonParameters.PSobject.Properties["vdiApp"].value
        }

        $PSO = [PSCustomObject]@{
            "vdiApp" = ${VdiApp}
        }

        return $PSO
    }

}
