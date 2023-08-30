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

.PARAMETER CheckGroup
No description available.
.OUTPUTS

AddCheckGroupsRequest<PSCustomObject>
#>

function Initialize-AddCheckGroupsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CheckGroup}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddCheckGroupsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $CheckGroup) {
            throw "invalid value for 'CheckGroup', 'CheckGroup' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "checkGroup" = ${CheckGroup}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddCheckGroupsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddCheckGroupsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddCheckGroupsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddCheckGroupsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddCheckGroupsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddCheckGroupsRequest
        $AllProperties = ("checkGroup")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'checkGroup' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "checkGroup"))) {
            throw "Error! JSON cannot be serialized due to the required property 'checkGroup' missing."
        } else {
            $CheckGroup = $JsonParameters.PSobject.Properties["checkGroup"].value
        }

        $PSO = [PSCustomObject]@{
            "checkGroup" = ${CheckGroup}
        }

        return $PSO
    }

}

