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

.PARAMETER Id
No description available.
.PARAMETER Name
No description available.
.PARAMETER Description
No description available.
.PARAMETER Type
No description available.
.PARAMETER Context
No description available.
.PARAMETER Featured
No description available.
.PARAMETER AllowQuantity
No description available.
.PARAMETER ImagePath
No description available.
.PARAMETER DarkImagePath
No description available.
.OUTPUTS

CatalogTypesInner<PSCustomObject>
#>

function Initialize-CatalogTypesInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Type},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Context},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Featured},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AllowQuantity},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ImagePath},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DarkImagePath}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CatalogTypesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "type" = ${Type}
            "context" = ${Context}
            "featured" = ${Featured}
            "allowQuantity" = ${AllowQuantity}
            "imagePath" = ${ImagePath}
            "darkImagePath" = ${DarkImagePath}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CatalogTypesInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to CatalogTypesInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CatalogTypesInner<PSCustomObject>
#>
function ConvertFrom-JsonToCatalogTypesInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CatalogTypesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CatalogTypesInner
        $AllProperties = ("id", "name", "description", "type", "context", "featured", "allowQuantity", "imagePath", "darkImagePath")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "id"))) { #optional property not found
            $Id = $null
        } else {
            $Id = $JsonParameters.PSobject.Properties["id"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "context"))) { #optional property not found
            $Context = $null
        } else {
            $Context = $JsonParameters.PSobject.Properties["context"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "featured"))) { #optional property not found
            $Featured = $null
        } else {
            $Featured = $JsonParameters.PSobject.Properties["featured"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "allowQuantity"))) { #optional property not found
            $AllowQuantity = $null
        } else {
            $AllowQuantity = $JsonParameters.PSobject.Properties["allowQuantity"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "imagePath"))) { #optional property not found
            $ImagePath = $null
        } else {
            $ImagePath = $JsonParameters.PSobject.Properties["imagePath"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "darkImagePath"))) { #optional property not found
            $DarkImagePath = $null
        } else {
            $DarkImagePath = $JsonParameters.PSobject.Properties["darkImagePath"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "type" = ${Type}
            "context" = ${Context}
            "featured" = ${Featured}
            "allowQuantity" = ${AllowQuantity}
            "imagePath" = ${ImagePath}
            "darkImagePath" = ${DarkImagePath}
        }

        return $PSO
    }

}

