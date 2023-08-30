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

.PARAMETER Role
No description available.
.OUTPUTS

AddRolesRequest<PSCustomObject>
#>

function Initialize-AddRolesRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Role}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddRolesRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Role) {
            throw "invalid value for 'Role', 'Role' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "role" = ${Role}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddRolesRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddRolesRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddRolesRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddRolesRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddRolesRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddRolesRequest
        $AllProperties = ("role")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'role' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "role"))) {
            throw "Error! JSON cannot be serialized due to the required property 'role' missing."
        } else {
            $Role = $JsonParameters.PSobject.Properties["role"].value
        }

        $PSO = [PSCustomObject]@{
            "role" = ${Role}
        }

        return $PSO
    }

}
