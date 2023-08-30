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
A name for the blueprint
.PARAMETER Image
Path to display image. Defaults to an internal Morpheus image.
.PARAMETER Type
Blueprint Type
.PARAMETER Labels
Array of label strings, can be used for filtering.
.PARAMETER Arm
No description available.
.OUTPUTS

BlueprintARMCreate<PSCustomObject>
#>

function Initialize-BlueprintARMCreate {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Image},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("arm")]
        [String]
        ${Type},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${Labels},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Arm}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BlueprintARMCreate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }

        if ($null -eq $Type) {
            throw "invalid value for 'Type', 'Type' cannot be null."
        }

        if ($null -eq $Arm) {
            throw "invalid value for 'Arm', 'Arm' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "image" = ${Image}
            "type" = ${Type}
            "labels" = ${Labels}
            "arm" = ${Arm}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BlueprintARMCreate<PSCustomObject>

.DESCRIPTION

Convert from JSON to BlueprintARMCreate<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BlueprintARMCreate<PSCustomObject>
#>
function ConvertFrom-JsonToBlueprintARMCreate {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BlueprintARMCreate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BlueprintARMCreate
        $AllProperties = ("name", "image", "type", "labels", "arm")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) {
            throw "Error! JSON cannot be serialized due to the required property 'type' missing."
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "arm"))) {
            throw "Error! JSON cannot be serialized due to the required property 'arm' missing."
        } else {
            $Arm = $JsonParameters.PSobject.Properties["arm"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "image"))) { #optional property not found
            $Image = $null
        } else {
            $Image = $JsonParameters.PSobject.Properties["image"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "labels"))) { #optional property not found
            $Labels = $null
        } else {
            $Labels = $JsonParameters.PSobject.Properties["labels"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "image" = ${Image}
            "type" = ${Type}
            "labels" = ${Labels}
            "arm" = ${Arm}
        }

        return $PSO
    }

}

