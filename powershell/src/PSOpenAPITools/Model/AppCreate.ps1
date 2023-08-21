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

.PARAMETER TemplateId
No description available.
.PARAMETER BlueprintId
The ID of the Blueprint. Use ""existing"" to create a blank app.
.PARAMETER Name
A unique name for the app
.PARAMETER Description
Description
.PARAMETER Labels
Array of label strings, can be used for filtering.
.PARAMETER Group
No description available.
.PARAMETER DefaultCloud
No description available.
.PARAMETER Environment
Environment code (appContext)
.PARAMETER Tiers
Configuration of app elements
.OUTPUTS

AppCreate<PSCustomObject>
#>

function Initialize-AppCreate {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${TemplateId},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${BlueprintId},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${Labels},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Group},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${DefaultCloud},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Environment},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Tiers}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AppCreate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if (!$BlueprintId) {
            throw "invalid value for 'BlueprintId', 'BlueprintId' cannot be null."
        }

        if (!$Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "templateId" = ${TemplateId}
            "blueprintId" = ${BlueprintId}
            "name" = ${Name}
            "description" = ${Description}
            "labels" = ${Labels}
            "group" = ${Group}
            "defaultCloud" = ${DefaultCloud}
            "environment" = ${Environment}
            "tiers" = ${Tiers}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AppCreate<PSCustomObject>

.DESCRIPTION

Convert from JSON to AppCreate<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AppCreate<PSCustomObject>
#>
function ConvertFrom-JsonToAppCreate {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AppCreate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AppCreate
        $AllProperties = ("templateId", "blueprintId", "name", "description", "labels", "group", "defaultCloud", "environment", "tiers")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property `blueprintId` missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "blueprintId"))) {
            throw "Error! JSON cannot be serialized due to the required property `blueprintId` missing."
        } else {
            $BlueprintId = $JsonParameters.PSobject.Properties["blueprintId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property `name` missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "templateId"))) { #optional property not found
            $TemplateId = $null
        } else {
            $TemplateId = $JsonParameters.PSobject.Properties["templateId"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "group"))) { #optional property not found
            $Group = $null
        } else {
            $Group = $JsonParameters.PSobject.Properties["group"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "defaultCloud"))) { #optional property not found
            $DefaultCloud = $null
        } else {
            $DefaultCloud = $JsonParameters.PSobject.Properties["defaultCloud"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "environment"))) { #optional property not found
            $Environment = $null
        } else {
            $Environment = $JsonParameters.PSobject.Properties["environment"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tiers"))) { #optional property not found
            $Tiers = $null
        } else {
            $Tiers = $JsonParameters.PSobject.Properties["tiers"].value
        }

        $PSO = [PSCustomObject]@{
            "templateId" = ${TemplateId}
            "blueprintId" = ${BlueprintId}
            "name" = ${Name}
            "description" = ${Description}
            "labels" = ${Labels}
            "group" = ${Group}
            "defaultCloud" = ${DefaultCloud}
            "environment" = ${Environment}
            "tiers" = ${Tiers}
        }

        return $PSO
    }

}

