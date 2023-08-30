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
A unique name for the app
.PARAMETER Description
Description
.PARAMETER Labels
Array of label strings, can be used for filtering.
.PARAMETER Environment
Environment code (appContext)
.PARAMETER OwnerId
User ID, can be used to change app owner. This also changes the owner for each instance in the app.
.OUTPUTS

AppUpdate<PSCustomObject>
#>

function Initialize-AppUpdate {
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
        [String]
        ${Environment},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${OwnerId}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AppUpdate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "labels" = ${Labels}
            "environment" = ${Environment}
            "ownerId" = ${OwnerId}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AppUpdate<PSCustomObject>

.DESCRIPTION

Convert from JSON to AppUpdate<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AppUpdate<PSCustomObject>
#>
function ConvertFrom-JsonToAppUpdate {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AppUpdate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AppUpdate
        $AllProperties = ("name", "description", "labels", "environment", "ownerId")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "environment"))) { #optional property not found
            $Environment = $null
        } else {
            $Environment = $JsonParameters.PSobject.Properties["environment"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ownerId"))) { #optional property not found
            $OwnerId = $null
        } else {
            $OwnerId = $JsonParameters.PSobject.Properties["ownerId"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "labels" = ${Labels}
            "environment" = ${Environment}
            "ownerId" = ${OwnerId}
        }

        return $PSO
    }

}

