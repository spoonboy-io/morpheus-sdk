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

.PARAMETER Name
A unique name for the workflow
.PARAMETER Description
A description of the workflow
.PARAMETER Labels
An array of Category labels for filtering
.PARAMETER Type
Workflow type
.PARAMETER Visibility
private or public
.PARAMETER OptionTypes
List of option type IDs for use with operational workflow configuration.
.PARAMETER Tasks
No description available.
.OUTPUTS

AddWorkflowsRequestTaskSet<PSCustomObject>
#>

function Initialize-AddWorkflowsRequestTaskSet {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${Labels},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("provision", "operation")]
        [String]
        ${Type} = "provision",
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("private", "public")]
        [String]
        ${Visibility} = "private",
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${OptionTypes},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Tasks}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddWorkflowsRequestTaskSet' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "labels" = ${Labels}
            "type" = ${Type}
            "visibility" = ${Visibility}
            "optionTypes" = ${OptionTypes}
            "tasks" = ${Tasks}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddWorkflowsRequestTaskSet<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddWorkflowsRequestTaskSet<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddWorkflowsRequestTaskSet<PSCustomObject>
#>
function ConvertFrom-JsonToAddWorkflowsRequestTaskSet {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddWorkflowsRequestTaskSet' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddWorkflowsRequestTaskSet
        $AllProperties = ("name", "description", "labels", "type", "visibility", "optionTypes", "tasks")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'name' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property 'name' missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "labels"))) { #optional property not found
            $Labels = $null
        } else {
            $Labels = $JsonParameters.PSobject.Properties["labels"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visibility"))) { #optional property not found
            $Visibility = $null
        } else {
            $Visibility = $JsonParameters.PSobject.Properties["visibility"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "optionTypes"))) { #optional property not found
            $OptionTypes = $null
        } else {
            $OptionTypes = $JsonParameters.PSobject.Properties["optionTypes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tasks"))) { #optional property not found
            $Tasks = $null
        } else {
            $Tasks = $JsonParameters.PSobject.Properties["tasks"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "labels" = ${Labels}
            "type" = ${Type}
            "visibility" = ${Visibility}
            "optionTypes" = ${OptionTypes}
            "tasks" = ${Tasks}
        }

        return $PSO
    }

}

