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
Code Repository ID, required for type `repository`. Use `/api/options/codeRepositories` to see available repositories.
.OUTPUTS

AddTasksRequestTaskFileRepository<PSCustomObject>
#>

function Initialize-AddTasksRequestTaskFileRepository {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddTasksRequestTaskFileRepository' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddTasksRequestTaskFileRepository<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddTasksRequestTaskFileRepository<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddTasksRequestTaskFileRepository<PSCustomObject>
#>
function ConvertFrom-JsonToAddTasksRequestTaskFileRepository {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddTasksRequestTaskFileRepository' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddTasksRequestTaskFileRepository
        $AllProperties = ("id")
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

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
        }

        return $PSO
    }

}

