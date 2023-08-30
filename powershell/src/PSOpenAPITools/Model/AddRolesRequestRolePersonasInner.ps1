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

.PARAMETER Code
`code` of the persona
.PARAMETER Access
The new access level.
.OUTPUTS

AddRolesRequestRolePersonasInner<PSCustomObject>
#>

function Initialize-AddRolesRequestRolePersonasInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("standard", "serviceCatalog", "vdi")]
        [String]
        ${Code},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("default", "full", "none")]
        [String]
        ${Access}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddRolesRequestRolePersonasInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Access) {
            throw "invalid value for 'Access', 'Access' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "code" = ${Code}
            "access" = ${Access}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddRolesRequestRolePersonasInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddRolesRequestRolePersonasInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddRolesRequestRolePersonasInner<PSCustomObject>
#>
function ConvertFrom-JsonToAddRolesRequestRolePersonasInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddRolesRequestRolePersonasInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddRolesRequestRolePersonasInner
        $AllProperties = ("code", "access")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'access' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "access"))) {
            throw "Error! JSON cannot be serialized due to the required property 'access' missing."
        } else {
            $Access = $JsonParameters.PSobject.Properties["access"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) { #optional property not found
            $Code = $null
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        $PSO = [PSCustomObject]@{
            "code" = ${Code}
            "access" = ${Access}
        }

        return $PSO
    }

}

