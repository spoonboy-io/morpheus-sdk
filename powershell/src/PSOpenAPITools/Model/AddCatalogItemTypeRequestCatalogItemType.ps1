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

.PARAMETER Json

JSON object

.OUTPUTS

AddCatalogItemTypeRequestCatalogItemType<PSCustomObject>
#>
function ConvertFrom-JsonToAddCatalogItemTypeRequestCatalogItemType {
    [CmdletBinding()]
    Param (
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        $match = 0
        $matchType = $null
        $matchInstance = $null

        # try to match CatalogItemTypeBlueprintCreate defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToCatalogItemTypeBlueprintCreate $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "CatalogItemTypeBlueprintCreate"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'CatalogItemTypeBlueprintCreate' defined in oneOf (AddCatalogItemTypeRequestCatalogItemType). Proceeding to the next one if any."
        }

        # try to match CatalogItemTypeInstanceCreate defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToCatalogItemTypeInstanceCreate $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "CatalogItemTypeInstanceCreate"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'CatalogItemTypeInstanceCreate' defined in oneOf (AddCatalogItemTypeRequestCatalogItemType). Proceeding to the next one if any."
        }

        # try to match CatalogItemTypeWorkflowCreate defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToCatalogItemTypeWorkflowCreate $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "CatalogItemTypeWorkflowCreate"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'CatalogItemTypeWorkflowCreate' defined in oneOf (AddCatalogItemTypeRequestCatalogItemType). Proceeding to the next one if any."
        }

        if ($match -gt 1) {
            throw "Error! The JSON payload matches more than one type defined in oneOf schemas ([CatalogItemTypeBlueprintCreate, CatalogItemTypeInstanceCreate, CatalogItemTypeWorkflowCreate]). JSON Payload: $($Json)"
        } elseif ($match -eq 1) {
            return [PSCustomObject]@{
                "ActualType" = ${matchType}
                "ActualInstance" = ${matchInstance}
                "OneOfSchemas" = @("CatalogItemTypeBlueprintCreate", "CatalogItemTypeInstanceCreate", "CatalogItemTypeWorkflowCreate")
            }
        } else {
            throw "Error! The JSON payload doesn't matches any type defined in oneOf schemas ([CatalogItemTypeBlueprintCreate, CatalogItemTypeInstanceCreate, CatalogItemTypeWorkflowCreate]). JSON Payload: $($Json)"
        }
    }
}

