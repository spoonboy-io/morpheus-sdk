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

AddVDIPools200Response<PSCustomObject>
#>
function ConvertFrom-JsonToAddVDIPools200Response {
    [CmdletBinding()]
    Param (
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        $match = 0
        $matchType = $null
        $matchInstance = $null

        if ($match -ne 0) { # no match yet
            # try to match AddVDIPools200ResponseAnyOf defined in the anyOf schemas
            try {
                $matchInstance = ConvertFrom-JsonToAddVDIPools200ResponseAnyOf $Json

                foreach($property in $matchInstance.PsObject.Properties) {
                    if ($null -ne $property.Value) {
                        $matchType = "AddVDIPools200ResponseAnyOf"
                        $match++
                        break
                    }
                }
            } catch {
                # fail to match the schema defined in anyOf, proceed to the next one
                Write-Debug "Failed to match 'AddVDIPools200ResponseAnyOf' defined in anyOf (AddVDIPools200Response). Proceeding to the next one if any."
            }
        }

        if ($match -ne 0) { # no match yet
            # try to match Model200Success defined in the anyOf schemas
            try {
                $matchInstance = ConvertFrom-JsonToModel200Success $Json

                foreach($property in $matchInstance.PsObject.Properties) {
                    if ($null -ne $property.Value) {
                        $matchType = "Model200Success"
                        $match++
                        break
                    }
                }
            } catch {
                # fail to match the schema defined in anyOf, proceed to the next one
                Write-Debug "Failed to match 'Model200Success' defined in anyOf (AddVDIPools200Response). Proceeding to the next one if any."
            }
        }

        if ($match -eq 1) {
            return [PSCustomObject]@{
                "ActualType" = ${matchType}
                "ActualInstance" = ${matchInstance}
                "anyOfSchemas" = @("AddVDIPools200ResponseAnyOf", "Model200Success")
            }
        } else {
            throw "Error! The JSON payload doesn't matches any type defined in anyOf schemas ([AddVDIPools200ResponseAnyOf, Model200Success]). JSON Payload: $($Json)"
        }
    }
}

