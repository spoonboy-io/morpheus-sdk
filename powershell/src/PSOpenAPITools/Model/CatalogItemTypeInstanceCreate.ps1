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
Catalog Item Type name
.PARAMETER Code
Useful shortcode for provisioning naming schemes and export reference.
.PARAMETER Category
Catalog Item Type category
.PARAMETER Description
Catalog Item Type description
.PARAMETER Labels
Array of label strings, can be used for filtering.
.PARAMETER Type
Type, `instance`, `blueprint` or `workflow`. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context.
.PARAMETER Visibility
Visibility - Set to public to allow all tenants
.PARAMETER LayoutCode
Identifier primarily used for Plugin Catalog Item Types
.PARAMETER IconPath
Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png.
.PARAMETER Enabled
Can be used to enable / disable the catalog item type.
.PARAMETER Featured
Can be used to feature the catalog item type.
.PARAMETER AllowQuantity
Can users order more than one of this item at a time.
.PARAMETER Config
No description available.
.PARAMETER OptionTypes
Array of option type IDs. Only applies to type instance and blueprint.
.PARAMETER Content
Documentation content for this Catalog Item. Markdown-formatted text is accepted and displayed appropriately when the item is ordered from the Service Catalog. A new Catalog Item-type Wiki entry will also be added containing this information.
.OUTPUTS

CatalogItemTypeInstanceCreate<PSCustomObject>
#>

function Initialize-CatalogItemTypeInstanceCreate {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Code},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Category},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${Labels},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("instance")]
        [String]
        ${Type},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Visibility} = "private",
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LayoutCode},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${IconPath},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled} = $true,
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Featured} = $false,
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AllowQuantity} = $false,
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${OptionTypes},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Content}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CatalogItemTypeInstanceCreate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if (!$Config) {
            throw "invalid value for 'Config', 'Config' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "code" = ${Code}
            "category" = ${Category}
            "description" = ${Description}
            "labels" = ${Labels}
            "type" = ${Type}
            "visibility" = ${Visibility}
            "layoutCode" = ${LayoutCode}
            "iconPath" = ${IconPath}
            "enabled" = ${Enabled}
            "featured" = ${Featured}
            "allowQuantity" = ${AllowQuantity}
            "config" = ${Config}
            "optionTypes" = ${OptionTypes}
            "content" = ${Content}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CatalogItemTypeInstanceCreate<PSCustomObject>

.DESCRIPTION

Convert from JSON to CatalogItemTypeInstanceCreate<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CatalogItemTypeInstanceCreate<PSCustomObject>
#>
function ConvertFrom-JsonToCatalogItemTypeInstanceCreate {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CatalogItemTypeInstanceCreate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CatalogItemTypeInstanceCreate
        $AllProperties = ("name", "code", "category", "description", "labels", "type", "visibility", "layoutCode", "iconPath", "enabled", "featured", "allowQuantity", "config", "optionTypes", "content")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property `config` missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) {
            throw "Error! JSON cannot be serialized due to the required property `config` missing."
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) { #optional property not found
            $Code = $null
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "category"))) { #optional property not found
            $Category = $null
        } else {
            $Category = $JsonParameters.PSobject.Properties["category"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visibility"))) { #optional property not found
            $Visibility = $null
        } else {
            $Visibility = $JsonParameters.PSobject.Properties["visibility"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "layoutCode"))) { #optional property not found
            $LayoutCode = $null
        } else {
            $LayoutCode = $JsonParameters.PSobject.Properties["layoutCode"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "iconPath"))) { #optional property not found
            $IconPath = $null
        } else {
            $IconPath = $JsonParameters.PSobject.Properties["iconPath"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "optionTypes"))) { #optional property not found
            $OptionTypes = $null
        } else {
            $OptionTypes = $JsonParameters.PSobject.Properties["optionTypes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "content"))) { #optional property not found
            $Content = $null
        } else {
            $Content = $JsonParameters.PSobject.Properties["content"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "code" = ${Code}
            "category" = ${Category}
            "description" = ${Description}
            "labels" = ${Labels}
            "type" = ${Type}
            "visibility" = ${Visibility}
            "layoutCode" = ${LayoutCode}
            "iconPath" = ${IconPath}
            "enabled" = ${Enabled}
            "featured" = ${Featured}
            "allowQuantity" = ${AllowQuantity}
            "config" = ${Config}
            "optionTypes" = ${OptionTypes}
            "content" = ${Content}
        }

        return $PSO
    }

}

