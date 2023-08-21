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
A unique name for the task
.PARAMETER Code
A unique code for the task
.PARAMETER Visibility
Visibility
.PARAMETER TaskType
No description available.
.PARAMETER Labels
An array of Category labels for filtering
.PARAMETER TaskOptions
Map of options specific to each `task type`. eg. script
.PARAMETER ResultType
No description available.
.PARAMETER ExecuteTarget
The execution target. eg. local,remote,resource. The default value varies by task type. 
.PARAMETER Retryable
If the task should be retried or not.
.PARAMETER RetryCount
The number of times to retry.
.PARAMETER RetryDelaySeconds
The delay, between retries.
.PARAMETER File
No description available.
.PARAMETER Credential
Map containing Credential ID or the default {""type"": ""local""}  which means use the values set in the local task options username and password instead of associating a credential. 
.OUTPUTS

ApiTasksIdTask<PSCustomObject>
#>

function Initialize-ApiTasksIdTask {
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
        ${Visibility} = "private",
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${TaskType},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${Labels},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${TaskOptions},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("exitCode", "keyValue", "json")]
        [String]
        ${ResultType},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("local", "remote", "resource")]
        [String]
        ${ExecuteTarget},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Retryable} = $false,
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${RetryCount},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${RetryDelaySeconds},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${File},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Credential}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiTasksIdTask' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "code" = ${Code}
            "visibility" = ${Visibility}
            "taskType" = ${TaskType}
            "labels" = ${Labels}
            "taskOptions" = ${TaskOptions}
            "resultType" = ${ResultType}
            "executeTarget" = ${ExecuteTarget}
            "retryable" = ${Retryable}
            "retryCount" = ${RetryCount}
            "retryDelaySeconds" = ${RetryDelaySeconds}
            "file" = ${File}
            "credential" = ${Credential}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiTasksIdTask<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiTasksIdTask<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiTasksIdTask<PSCustomObject>
#>
function ConvertFrom-JsonToApiTasksIdTask {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiTasksIdTask' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiTasksIdTask
        $AllProperties = ("name", "code", "visibility", "taskType", "labels", "taskOptions", "resultType", "executeTarget", "retryable", "retryCount", "retryDelaySeconds", "file", "credential")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visibility"))) { #optional property not found
            $Visibility = $null
        } else {
            $Visibility = $JsonParameters.PSobject.Properties["visibility"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "taskType"))) { #optional property not found
            $TaskType = $null
        } else {
            $TaskType = $JsonParameters.PSobject.Properties["taskType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "labels"))) { #optional property not found
            $Labels = $null
        } else {
            $Labels = $JsonParameters.PSobject.Properties["labels"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "taskOptions"))) { #optional property not found
            $TaskOptions = $null
        } else {
            $TaskOptions = $JsonParameters.PSobject.Properties["taskOptions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "resultType"))) { #optional property not found
            $ResultType = $null
        } else {
            $ResultType = $JsonParameters.PSobject.Properties["resultType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "executeTarget"))) { #optional property not found
            $ExecuteTarget = $null
        } else {
            $ExecuteTarget = $JsonParameters.PSobject.Properties["executeTarget"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "retryable"))) { #optional property not found
            $Retryable = $null
        } else {
            $Retryable = $JsonParameters.PSobject.Properties["retryable"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "retryCount"))) { #optional property not found
            $RetryCount = $null
        } else {
            $RetryCount = $JsonParameters.PSobject.Properties["retryCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "retryDelaySeconds"))) { #optional property not found
            $RetryDelaySeconds = $null
        } else {
            $RetryDelaySeconds = $JsonParameters.PSobject.Properties["retryDelaySeconds"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "file"))) { #optional property not found
            $File = $null
        } else {
            $File = $JsonParameters.PSobject.Properties["file"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "credential"))) { #optional property not found
            $Credential = $null
        } else {
            $Credential = $JsonParameters.PSobject.Properties["credential"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "code" = ${Code}
            "visibility" = ${Visibility}
            "taskType" = ${TaskType}
            "labels" = ${Labels}
            "taskOptions" = ${TaskOptions}
            "resultType" = ${ResultType}
            "executeTarget" = ${ExecuteTarget}
            "retryable" = ${Retryable}
            "retryCount" = ${RetryCount}
            "retryDelaySeconds" = ${RetryDelaySeconds}
            "file" = ${File}
            "credential" = ${Credential}
        }

        return $PSO
    }

}

