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

.PARAMETER ImageBuildExecutions
No description available.
.PARAMETER ImageBuildExecutionCount
No description available.
.PARAMETER ImageBuild
No description available.
.OUTPUTS

GetImageBuildExecutions200ResponseAllOf<PSCustomObject>
#>

function Initialize-GetImageBuildExecutions200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${ImageBuildExecutions},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ImageBuildExecutionCount},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ImageBuild}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetImageBuildExecutions200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "imageBuildExecutions" = ${ImageBuildExecutions}
            "imageBuildExecutionCount" = ${ImageBuildExecutionCount}
            "imageBuild" = ${ImageBuild}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetImageBuildExecutions200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetImageBuildExecutions200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetImageBuildExecutions200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToGetImageBuildExecutions200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetImageBuildExecutions200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetImageBuildExecutions200ResponseAllOf
        $AllProperties = ("imageBuildExecutions", "imageBuildExecutionCount", "imageBuild")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "imageBuildExecutions"))) { #optional property not found
            $ImageBuildExecutions = $null
        } else {
            $ImageBuildExecutions = $JsonParameters.PSobject.Properties["imageBuildExecutions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "imageBuildExecutionCount"))) { #optional property not found
            $ImageBuildExecutionCount = $null
        } else {
            $ImageBuildExecutionCount = $JsonParameters.PSobject.Properties["imageBuildExecutionCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "imageBuild"))) { #optional property not found
            $ImageBuild = $null
        } else {
            $ImageBuild = $JsonParameters.PSobject.Properties["imageBuild"].value
        }

        $PSO = [PSCustomObject]@{
            "imageBuildExecutions" = ${ImageBuildExecutions}
            "imageBuildExecutionCount" = ${ImageBuildExecutionCount}
            "imageBuild" = ${ImageBuild}
        }

        return $PSO
    }

}

