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
VDI App name
.PARAMETER Description
Description
.PARAMETER IconPath
Icon Path. A relative location of an icon image
.PARAMETER LaunchPrefix
The RDS App Name Prefix.  Must start with '||'
.OUTPUTS

AddVDIAppsRequestVdiAppOneOf<PSCustomObject>
#>

function Initialize-AddVDIAppsRequestVdiAppOneOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.IO.FileInfo]
        ${IconPath},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LaunchPrefix}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddVDIAppsRequestVdiAppOneOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "iconPath" = ${IconPath}
            "launchPrefix" = ${LaunchPrefix}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddVDIAppsRequestVdiAppOneOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddVDIAppsRequestVdiAppOneOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddVDIAppsRequestVdiAppOneOf<PSCustomObject>
#>
function ConvertFrom-JsonToAddVDIAppsRequestVdiAppOneOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddVDIAppsRequestVdiAppOneOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddVDIAppsRequestVdiAppOneOf
        $AllProperties = ("name", "description", "iconPath", "launchPrefix")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "iconPath"))) { #optional property not found
            $IconPath = $null
        } else {
            $IconPath = $JsonParameters.PSobject.Properties["iconPath"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "launchPrefix"))) { #optional property not found
            $LaunchPrefix = $null
        } else {
            $LaunchPrefix = $JsonParameters.PSobject.Properties["launchPrefix"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "iconPath" = ${IconPath}
            "launchPrefix" = ${LaunchPrefix}
        }

        return $PSO
    }

}

