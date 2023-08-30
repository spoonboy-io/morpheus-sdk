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

.PARAMETER TaskSet
No description available.
.OUTPUTS

AddWorkflowsRequest<PSCustomObject>
#>

function Initialize-AddWorkflowsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${TaskSet}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddWorkflowsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $TaskSet) {
            throw "invalid value for 'TaskSet', 'TaskSet' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "taskSet" = ${TaskSet}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddWorkflowsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddWorkflowsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddWorkflowsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddWorkflowsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddWorkflowsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddWorkflowsRequest
        $AllProperties = ("taskSet")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'taskSet' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "taskSet"))) {
            throw "Error! JSON cannot be serialized due to the required property 'taskSet' missing."
        } else {
            $TaskSet = $JsonParameters.PSobject.Properties["taskSet"].value
        }

        $PSO = [PSCustomObject]@{
            "taskSet" = ${TaskSet}
        }

        return $PSO
    }

}

