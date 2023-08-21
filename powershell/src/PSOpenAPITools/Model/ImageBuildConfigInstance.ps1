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

.PARAMETER Layout
No description available.
.PARAMETER Type
No description available.
.PARAMETER UserGroup
No description available.
.OUTPUTS

ImageBuildConfigInstance<PSCustomObject>
#>

function Initialize-ImageBuildConfigInstance {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Layout},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Type},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${UserGroup}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ImageBuildConfigInstance' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "layout" = ${Layout}
            "type" = ${Type}
            "userGroup" = ${UserGroup}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ImageBuildConfigInstance<PSCustomObject>

.DESCRIPTION

Convert from JSON to ImageBuildConfigInstance<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ImageBuildConfigInstance<PSCustomObject>
#>
function ConvertFrom-JsonToImageBuildConfigInstance {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ImageBuildConfigInstance' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ImageBuildConfigInstance
        $AllProperties = ("layout", "type", "userGroup")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "layout"))) { #optional property not found
            $Layout = $null
        } else {
            $Layout = $JsonParameters.PSobject.Properties["layout"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "userGroup"))) { #optional property not found
            $UserGroup = $null
        } else {
            $UserGroup = $JsonParameters.PSobject.Properties["userGroup"].value
        }

        $PSO = [PSCustomObject]@{
            "layout" = ${Layout}
            "type" = ${Type}
            "userGroup" = ${UserGroup}
        }

        return $PSO
    }

}
