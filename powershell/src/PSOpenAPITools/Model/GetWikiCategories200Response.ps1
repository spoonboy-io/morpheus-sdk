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

.PARAMETER Categories
No description available.
.OUTPUTS

GetWikiCategories200Response<PSCustomObject>
#>

function Initialize-GetWikiCategories200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Categories}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetWikiCategories200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "categories" = ${Categories}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetWikiCategories200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetWikiCategories200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetWikiCategories200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetWikiCategories200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetWikiCategories200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetWikiCategories200Response
        $AllProperties = ("categories")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "categories"))) { #optional property not found
            $Categories = $null
        } else {
            $Categories = $JsonParameters.PSobject.Properties["categories"].value
        }

        $PSO = [PSCustomObject]@{
            "categories" = ${Categories}
        }

        return $PSO
    }

}

