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
A name for the security package
.PARAMETER Labels
Array of label strings, can be used for filtering.
.PARAMETER Type
Security Package Type Code
.PARAMETER Description
A description for the security package
.PARAMETER Url
URL to download the security package zip file from
.PARAMETER Enabled
Can be used to disable the security package
.OUTPUTS

ApiSecurityPackagesSecurityPackage<PSCustomObject>
#>

function Initialize-ApiSecurityPackagesSecurityPackage {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${Labels},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Type} = "scap-package",
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Url},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled} = $true
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiSecurityPackagesSecurityPackage' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if (!$Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }

        if (!$Url) {
            throw "invalid value for 'Url', 'Url' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "labels" = ${Labels}
            "type" = ${Type}
            "description" = ${Description}
            "url" = ${Url}
            "enabled" = ${Enabled}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiSecurityPackagesSecurityPackage<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiSecurityPackagesSecurityPackage<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiSecurityPackagesSecurityPackage<PSCustomObject>
#>
function ConvertFrom-JsonToApiSecurityPackagesSecurityPackage {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiSecurityPackagesSecurityPackage' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiSecurityPackagesSecurityPackage
        $AllProperties = ("name", "labels", "type", "description", "url", "enabled")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property `name` missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property `name` missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "url"))) {
            throw "Error! JSON cannot be serialized due to the required property `url` missing."
        } else {
            $Url = $JsonParameters.PSobject.Properties["url"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "labels" = ${Labels}
            "type" = ${Type}
            "description" = ${Description}
            "url" = ${Url}
            "enabled" = ${Enabled}
        }

        return $PSO
    }

}

