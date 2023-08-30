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
No description available.
.PARAMETER ProcessId
No description available.
.OUTPUTS

ExecuteTasks200ResponseAllOfJobExecution<PSCustomObject>
#>

function Initialize-ExecuteTasks200ResponseAllOfJobExecution {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ProcessId}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ExecuteTasks200ResponseAllOfJobExecution' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "processId" = ${ProcessId}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ExecuteTasks200ResponseAllOfJobExecution<PSCustomObject>

.DESCRIPTION

Convert from JSON to ExecuteTasks200ResponseAllOfJobExecution<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ExecuteTasks200ResponseAllOfJobExecution<PSCustomObject>
#>
function ConvertFrom-JsonToExecuteTasks200ResponseAllOfJobExecution {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ExecuteTasks200ResponseAllOfJobExecution' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ExecuteTasks200ResponseAllOfJobExecution
        $AllProperties = ("id", "processId")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "id"))) { #optional property not found
            $Id = $null
        } else {
            $Id = $JsonParameters.PSobject.Properties["id"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "processId"))) { #optional property not found
            $ProcessId = $null
        } else {
            $ProcessId = $JsonParameters.PSobject.Properties["processId"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "processId" = ${ProcessId}
        }

        return $PSO
    }

}

