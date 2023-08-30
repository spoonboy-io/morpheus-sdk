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

.PARAMETER ListItems
No description available.
.OUTPUTS

GetOptionListItems200Response<PSCustomObject>
#>

function Initialize-GetOptionListItems200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${ListItems}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetOptionListItems200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "listItems" = ${ListItems}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetOptionListItems200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetOptionListItems200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetOptionListItems200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetOptionListItems200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetOptionListItems200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetOptionListItems200Response
        $AllProperties = ("listItems")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "listItems"))) { #optional property not found
            $ListItems = $null
        } else {
            $ListItems = $JsonParameters.PSobject.Properties["listItems"].value
        }

        $PSO = [PSCustomObject]@{
            "listItems" = ${ListItems}
        }

        return $PSO
    }

}

