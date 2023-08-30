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

.PARAMETER CatalogItemType
No description available.
.OUTPUTS

AddCatalogItemType200ResponseAllOf<PSCustomObject>
#>

function Initialize-AddCatalogItemType200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CatalogItemType}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddCatalogItemType200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "catalogItemType" = ${CatalogItemType}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddCatalogItemType200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddCatalogItemType200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddCatalogItemType200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToAddCatalogItemType200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddCatalogItemType200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddCatalogItemType200ResponseAllOf
        $AllProperties = ("catalogItemType")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "catalogItemType"))) { #optional property not found
            $CatalogItemType = $null
        } else {
            $CatalogItemType = $JsonParameters.PSobject.Properties["catalogItemType"].value
        }

        $PSO = [PSCustomObject]@{
            "catalogItemType" = ${CatalogItemType}
        }

        return $PSO
    }

}
