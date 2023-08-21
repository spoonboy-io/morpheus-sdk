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
.OUTPUTS

ApiZonesZoneIdFoldersIdFolderTenantPermissions<PSCustomObject>
#>

function Initialize-ApiZonesZoneIdFoldersIdFolderTenantPermissions {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${Accounts}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiZonesZoneIdFoldersIdFolderTenantPermissions' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "accounts" = ${Accounts}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiZonesZoneIdFoldersIdFolderTenantPermissions<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiZonesZoneIdFoldersIdFolderTenantPermissions<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiZonesZoneIdFoldersIdFolderTenantPermissions<PSCustomObject>
#>
function ConvertFrom-JsonToApiZonesZoneIdFoldersIdFolderTenantPermissions {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiZonesZoneIdFoldersIdFolderTenantPermissions' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiZonesZoneIdFoldersIdFolderTenantPermissions
        $AllProperties = ("accounts")
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

        $PSO = [PSCustomObject]@{
            "accounts" = ${Accounts}
        }

        return $PSO
    }

}
