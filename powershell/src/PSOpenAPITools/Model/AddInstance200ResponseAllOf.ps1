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

AddInstance200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToAddInstance200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        $match = 0
        $matchType = $null
        $matchInstance = $null

        # try to match InstanceCreateSuccess defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToInstanceCreateSuccess $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "InstanceCreateSuccess"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'InstanceCreateSuccess' defined in oneOf (AddInstance200ResponseAllOf). Proceeding to the next one if any."
        }

        # try to match SuccessError defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToSuccessError $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "SuccessError"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'SuccessError' defined in oneOf (AddInstance200ResponseAllOf). Proceeding to the next one if any."
        }

        if ($match -gt 1) {
            throw "Error! The JSON payload matches more than one type defined in oneOf schemas ([InstanceCreateSuccess, SuccessError]). JSON Payload: $($Json)"
        } elseif ($match -eq 1) {
            return [PSCustomObject]@{
                "ActualType" = ${matchType}
                "ActualInstance" = ${matchInstance}
                "OneOfSchemas" = @("InstanceCreateSuccess", "SuccessError")
            }
        } else {
            throw "Error! The JSON payload doesn't matches any type defined in oneOf schemas ([InstanceCreateSuccess, SuccessError]). JSON Payload: $($Json)"
        }
    }
}
