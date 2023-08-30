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

AddJobsRequestJob<PSCustomObject>
#>
function ConvertFrom-JsonToAddJobsRequestJob {
    [CmdletBinding()]
    Param (
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        $match = 0
        $matchType = $null
        $matchInstance = $null

        # try to match SecurityScanJob defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToSecurityScanJob $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "SecurityScanJob"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'SecurityScanJob' defined in oneOf (AddJobsRequestJob). Proceeding to the next one if any."
        }

        # try to match TaskJobPayload defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskJobPayload $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskJobPayload"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskJobPayload' defined in oneOf (AddJobsRequestJob). Proceeding to the next one if any."
        }

        # try to match WorkflowJobPayload defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToWorkflowJobPayload $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "WorkflowJobPayload"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'WorkflowJobPayload' defined in oneOf (AddJobsRequestJob). Proceeding to the next one if any."
        }

        if ($match -gt 1) {
            throw "Error! The JSON payload matches more than one type defined in oneOf schemas ([SecurityScanJob, TaskJobPayload, WorkflowJobPayload]). JSON Payload: $($Json)"
        } elseif ($match -eq 1) {
            return [PSCustomObject]@{
                "ActualType" = ${matchType}
                "ActualInstance" = ${matchInstance}
                "OneOfSchemas" = @("SecurityScanJob", "TaskJobPayload", "WorkflowJobPayload")
            }
        } else {
            throw "Error! The JSON payload doesn't matches any type defined in oneOf schemas ([SecurityScanJob, TaskJobPayload, WorkflowJobPayload]). JSON Payload: $($Json)"
        }
    }
}

