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

.PARAMETER Active
Activate `true` or disable `false` the datastore
.PARAMETER Visibility
Setting `private` or `public`
.PARAMETER TenantPermissions
No description available.
.PARAMETER ResourcePermissions
No description available.
.OUTPUTS

ApiZonesZoneIdDataStoresIdDatastore<PSCustomObject>
#>

function Initialize-ApiZonesZoneIdDataStoresIdDatastore {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Active},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("public", "private")]
        [String]
        ${Visibility} = "private",
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${TenantPermissions},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ResourcePermissions}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiZonesZoneIdDataStoresIdDatastore' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "active" = ${Active}
            "visibility" = ${Visibility}
            "tenantPermissions" = ${TenantPermissions}
            "resourcePermissions" = ${ResourcePermissions}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiZonesZoneIdDataStoresIdDatastore<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiZonesZoneIdDataStoresIdDatastore<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiZonesZoneIdDataStoresIdDatastore<PSCustomObject>
#>
function ConvertFrom-JsonToApiZonesZoneIdDataStoresIdDatastore {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiZonesZoneIdDataStoresIdDatastore' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiZonesZoneIdDataStoresIdDatastore
        $AllProperties = ("active", "visibility", "tenantPermissions", "resourcePermissions")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "active"))) { #optional property not found
            $Active = $null
        } else {
            $Active = $JsonParameters.PSobject.Properties["active"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visibility"))) { #optional property not found
            $Visibility = $null
        } else {
            $Visibility = $JsonParameters.PSobject.Properties["visibility"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tenantPermissions"))) { #optional property not found
            $TenantPermissions = $null
        } else {
            $TenantPermissions = $JsonParameters.PSobject.Properties["tenantPermissions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "resourcePermissions"))) { #optional property not found
            $ResourcePermissions = $null
        } else {
            $ResourcePermissions = $JsonParameters.PSobject.Properties["resourcePermissions"].value
        }

        $PSO = [PSCustomObject]@{
            "active" = ${Active}
            "visibility" = ${Visibility}
            "tenantPermissions" = ${TenantPermissions}
            "resourcePermissions" = ${ResourcePermissions}
        }

        return $PSO
    }

}

