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

AddCloudResourcePoolRequestResourcePoolConfig<PSCustomObject>
#>
function ConvertFrom-JsonToAddCloudResourcePoolRequestResourcePoolConfig {
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
            # try to match AwsResourcePoolConfig defined in the anyOf schemas
            try {
                $matchInstance = ConvertFrom-JsonToAwsResourcePoolConfig $Json

                foreach($property in $matchInstance.PsObject.Properties) {
                    if ($null -ne $property.Value) {
                        $matchType = "AwsResourcePoolConfig"
                        $match++
                        break
                    }
                }
            } catch {
                # fail to match the schema defined in anyOf, proceed to the next one
                Write-Debug "Failed to match 'AwsResourcePoolConfig' defined in anyOf (AddCloudResourcePoolRequestResourcePoolConfig). Proceeding to the next one if any."
            }
        }

        if ($match -ne 0) { # no match yet
            # try to match CloudFoundryResourcePoolConfig defined in the anyOf schemas
            try {
                $matchInstance = ConvertFrom-JsonToCloudFoundryResourcePoolConfig $Json

                foreach($property in $matchInstance.PsObject.Properties) {
                    if ($null -ne $property.Value) {
                        $matchType = "CloudFoundryResourcePoolConfig"
                        $match++
                        break
                    }
                }
            } catch {
                # fail to match the schema defined in anyOf, proceed to the next one
                Write-Debug "Failed to match 'CloudFoundryResourcePoolConfig' defined in anyOf (AddCloudResourcePoolRequestResourcePoolConfig). Proceeding to the next one if any."
            }
        }

        if ($match -eq 1) {
            return [PSCustomObject]@{
                "ActualType" = ${matchType}
                "ActualInstance" = ${matchInstance}
                "anyOfSchemas" = @("AwsResourcePoolConfig", "CloudFoundryResourcePoolConfig")
            }
        } else {
            throw "Error! The JSON payload doesn't matches any type defined in anyOf schemas ([AwsResourcePoolConfig, CloudFoundryResourcePoolConfig]). JSON Payload: $($Json)"
        }
    }
}

