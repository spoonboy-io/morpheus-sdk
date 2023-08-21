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

.PARAMETER Accounts
Array of tenant account ids that are allowed access
.PARAMETER DefaultTarget
Array of tenant account ids which should use the data store as the Default
.PARAMETER DefaultStore
Array of tenant account ids which should use the data store as the Image Target
.OUTPUTS

ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions<PSCustomObject>
#>

function Initialize-ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${Accounts},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${DefaultTarget},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${DefaultStore}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "accounts" = ${Accounts}
            "defaultTarget" = ${DefaultTarget}
            "defaultStore" = ${DefaultStore}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions<PSCustomObject>
#>
function ConvertFrom-JsonToApiZonesZoneIdDataStoresIdDatastoreTenantPermissions {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions
        $AllProperties = ("accounts", "defaultTarget", "defaultStore")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "accounts"))) { #optional property not found
            $Accounts = $null
        } else {
            $Accounts = $JsonParameters.PSobject.Properties["accounts"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "defaultTarget"))) { #optional property not found
            $DefaultTarget = $null
        } else {
            $DefaultTarget = $JsonParameters.PSobject.Properties["defaultTarget"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "defaultStore"))) { #optional property not found
            $DefaultStore = $null
        } else {
            $DefaultStore = $JsonParameters.PSobject.Properties["defaultStore"].value
        }

        $PSO = [PSCustomObject]@{
            "accounts" = ${Accounts}
            "defaultTarget" = ${DefaultTarget}
            "defaultStore" = ${DefaultStore}
        }

        return $PSO
    }

}

