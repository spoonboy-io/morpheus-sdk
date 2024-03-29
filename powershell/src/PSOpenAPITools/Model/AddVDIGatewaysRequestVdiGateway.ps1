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

AddVDIGatewaysRequestVdiGateway<PSCustomObject>
#>
function ConvertFrom-JsonToAddVDIGatewaysRequestVdiGateway {
    [CmdletBinding()]
    Param (
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        $match = 0
        $matchType = $null
        $matchInstance = $null

        # try to match AddVDIGatewaysRequestVdiGatewayOneOf defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToAddVDIGatewaysRequestVdiGatewayOneOf $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "AddVDIGatewaysRequestVdiGatewayOneOf"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'AddVDIGatewaysRequestVdiGatewayOneOf' defined in oneOf (AddVDIGatewaysRequestVdiGateway). Proceeding to the next one if any."
        }

        if ($match -gt 1) {
            throw "Error! The JSON payload matches more than one type defined in oneOf schemas ([AddVDIGatewaysRequestVdiGatewayOneOf]). JSON Payload: $($Json)"
        } elseif ($match -eq 1) {
            return [PSCustomObject]@{
                "ActualType" = ${matchType}
                "ActualInstance" = ${matchInstance}
                "OneOfSchemas" = @("AddVDIGatewaysRequestVdiGatewayOneOf")
            }
        } else {
            throw "Error! The JSON payload doesn't matches any type defined in oneOf schemas ([AddVDIGatewaysRequestVdiGatewayOneOf]). JSON Payload: $($Json)"
        }
    }
}

