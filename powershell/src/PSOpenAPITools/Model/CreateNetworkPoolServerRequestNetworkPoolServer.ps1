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

Payload for creating a new Network Pool Server

.PARAMETER Json

JSON object

.OUTPUTS

CreateNetworkPoolServerRequestNetworkPoolServer<PSCustomObject>
#>
function ConvertFrom-JsonToCreateNetworkPoolServerRequestNetworkPoolServer {
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
            # try to match NetworkPoolServerCreateBluecat defined in the anyOf schemas
            try {
                $matchInstance = ConvertFrom-JsonToNetworkPoolServerCreateBluecat $Json

                foreach($property in $matchInstance.PsObject.Properties) {
                    if ($null -ne $property.Value) {
                        $matchType = "NetworkPoolServerCreateBluecat"
                        $match++
                        break
                    }
                }
            } catch {
                # fail to match the schema defined in anyOf, proceed to the next one
                Write-Debug "Failed to match 'NetworkPoolServerCreateBluecat' defined in anyOf (CreateNetworkPoolServerRequestNetworkPoolServer). Proceeding to the next one if any."
            }
        }

        if ($match -ne 0) { # no match yet
            # try to match NetworkPoolServerCreateInfoblox defined in the anyOf schemas
            try {
                $matchInstance = ConvertFrom-JsonToNetworkPoolServerCreateInfoblox $Json

                foreach($property in $matchInstance.PsObject.Properties) {
                    if ($null -ne $property.Value) {
                        $matchType = "NetworkPoolServerCreateInfoblox"
                        $match++
                        break
                    }
                }
            } catch {
                # fail to match the schema defined in anyOf, proceed to the next one
                Write-Debug "Failed to match 'NetworkPoolServerCreateInfoblox' defined in anyOf (CreateNetworkPoolServerRequestNetworkPoolServer). Proceeding to the next one if any."
            }
        }

        if ($match -ne 0) { # no match yet
            # try to match NetworkPoolServerCreatePhpIpam defined in the anyOf schemas
            try {
                $matchInstance = ConvertFrom-JsonToNetworkPoolServerCreatePhpIpam $Json

                foreach($property in $matchInstance.PsObject.Properties) {
                    if ($null -ne $property.Value) {
                        $matchType = "NetworkPoolServerCreatePhpIpam"
                        $match++
                        break
                    }
                }
            } catch {
                # fail to match the schema defined in anyOf, proceed to the next one
                Write-Debug "Failed to match 'NetworkPoolServerCreatePhpIpam' defined in anyOf (CreateNetworkPoolServerRequestNetworkPoolServer). Proceeding to the next one if any."
            }
        }

        if ($match -ne 0) { # no match yet
            # try to match NetworkPoolServerCreateSolarWinds defined in the anyOf schemas
            try {
                $matchInstance = ConvertFrom-JsonToNetworkPoolServerCreateSolarWinds $Json

                foreach($property in $matchInstance.PsObject.Properties) {
                    if ($null -ne $property.Value) {
                        $matchType = "NetworkPoolServerCreateSolarWinds"
                        $match++
                        break
                    }
                }
            } catch {
                # fail to match the schema defined in anyOf, proceed to the next one
                Write-Debug "Failed to match 'NetworkPoolServerCreateSolarWinds' defined in anyOf (CreateNetworkPoolServerRequestNetworkPoolServer). Proceeding to the next one if any."
            }
        }

        if ($match -eq 1) {
            return [PSCustomObject]@{
                "ActualType" = ${matchType}
                "ActualInstance" = ${matchInstance}
                "anyOfSchemas" = @("NetworkPoolServerCreateBluecat", "NetworkPoolServerCreateInfoblox", "NetworkPoolServerCreatePhpIpam", "NetworkPoolServerCreateSolarWinds")
            }
        } else {
            throw "Error! The JSON payload doesn't matches any type defined in anyOf schemas ([NetworkPoolServerCreateBluecat, NetworkPoolServerCreateInfoblox, NetworkPoolServerCreatePhpIpam, NetworkPoolServerCreateSolarWinds]). JSON Payload: $($Json)"
        }
    }
}

