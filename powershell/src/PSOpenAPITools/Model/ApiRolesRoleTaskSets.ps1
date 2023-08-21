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

.PARAMETER Id
`id` of the workflow (taskSet)
.PARAMETER Access
The new access level.
.OUTPUTS

ApiRolesRoleTaskSets<PSCustomObject>
#>

function Initialize-ApiRolesRoleTaskSets {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [Int32]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("default", "full", "none")]
        [String]
        ${Access}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiRolesRoleTaskSets' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if (!$Id) {
            throw "invalid value for 'Id', 'Id' cannot be null."
        }

        if (!$Access) {
            throw "invalid value for 'Access', 'Access' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "access" = ${Access}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiRolesRoleTaskSets<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiRolesRoleTaskSets<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiRolesRoleTaskSets<PSCustomObject>
#>
function ConvertFrom-JsonToApiRolesRoleTaskSets {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiRolesRoleTaskSets' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiRolesRoleTaskSets
        $AllProperties = ("id", "access")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property `id` missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "id"))) {
            throw "Error! JSON cannot be serialized due to the required property `id` missing."
        } else {
            $Id = $JsonParameters.PSobject.Properties["id"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "access"))) {
            throw "Error! JSON cannot be serialized due to the required property `access` missing."
        } else {
            $Access = $JsonParameters.PSobject.Properties["access"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "access" = ${Access}
        }

        return $PSO
    }

}

