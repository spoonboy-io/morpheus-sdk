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
.PARAMETER Account
No description available.
.PARAMETER Active
No description available.
.PARAMETER ApiKey
No description available.
.PARAMETER Availability
No description available.
.PARAMETER CheckAgent
No description available.
.PARAMETER CheckInterval
No description available.
.PARAMETER CheckSpec
No description available.
.PARAMETER CheckType
No description available.
.PARAMETER Config
No description available.
.PARAMETER Container
No description available.
.PARAMETER CreateIncident
No description available.
.PARAMETER Muted
No description available.
.PARAMETER CreatedBy
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER Description
No description available.
.PARAMETER EndDate
No description available.
.PARAMETER Health
No description available.
.PARAMETER InUptime
No description available.
.PARAMETER LastBoxStats
No description available.
.PARAMETER LastCheckStatus
No description available.
.PARAMETER LastError
No description available.
.PARAMETER LastErrorDate
No description available.
.PARAMETER LastMessage
No description available.
.PARAMETER LastMetric
No description available.
.PARAMETER LastRunDate
No description available.
.PARAMETER LastStats
No description available.
.PARAMETER LastSuccessDate
No description available.
.PARAMETER LastTimer
No description available.
.PARAMETER LastUpdated
No description available.
.PARAMETER LastWarningDate
No description available.
.PARAMETER Name
No description available.
.PARAMETER NextRunDate
No description available.
.PARAMETER OutageTime
No description available.
.PARAMETER Severity
No description available.
.PARAMETER StartDate
No description available.
.PARAMETER Deleted
No description available.
.OUTPUTS

Check<PSCustomObject>
#>

function Initialize-Check {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Account},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Active},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ApiKey},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Availability},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CheckAgent},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${CheckInterval},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CheckSpec},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CheckType},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Container},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${CreateIncident},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Muted},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CreatedBy},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${EndDate},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Health},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${InUptime},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LastBoxStats},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LastCheckStatus},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LastError},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastErrorDate},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LastMessage},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LastMetric},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastRunDate},
        [Parameter(Position = 26, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LastStats},
        [Parameter(Position = 27, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastSuccessDate},
        [Parameter(Position = 28, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${LastTimer},
        [Parameter(Position = 29, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated},
        [Parameter(Position = 30, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastWarningDate},
        [Parameter(Position = 31, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 32, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${NextRunDate},
        [Parameter(Position = 33, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${OutageTime},
        [Parameter(Position = 34, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Severity},
        [Parameter(Position = 35, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${StartDate},
        [Parameter(Position = 36, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Deleted}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => Check' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "account" = ${Account}
            "active" = ${Active}
            "apiKey" = ${ApiKey}
            "availability" = ${Availability}
            "checkAgent" = ${CheckAgent}
            "checkInterval" = ${CheckInterval}
            "checkSpec" = ${CheckSpec}
            "checkType" = ${CheckType}
            "config" = ${Config}
            "container" = ${Container}
            "createIncident" = ${CreateIncident}
            "muted" = ${Muted}
            "createdBy" = ${CreatedBy}
            "dateCreated" = ${DateCreated}
            "description" = ${Description}
            "endDate" = ${EndDate}
            "health" = ${Health}
            "inUptime" = ${InUptime}
            "lastBoxStats" = ${LastBoxStats}
            "lastCheckStatus" = ${LastCheckStatus}
            "lastError" = ${LastError}
            "lastErrorDate" = ${LastErrorDate}
            "lastMessage" = ${LastMessage}
            "lastMetric" = ${LastMetric}
            "lastRunDate" = ${LastRunDate}
            "lastStats" = ${LastStats}
            "lastSuccessDate" = ${LastSuccessDate}
            "lastTimer" = ${LastTimer}
            "lastUpdated" = ${LastUpdated}
            "lastWarningDate" = ${LastWarningDate}
            "name" = ${Name}
            "nextRunDate" = ${NextRunDate}
            "outageTime" = ${OutageTime}
            "severity" = ${Severity}
            "startDate" = ${StartDate}
            "deleted" = ${Deleted}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to Check<PSCustomObject>

.DESCRIPTION

Convert from JSON to Check<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

Check<PSCustomObject>
#>
function ConvertFrom-JsonToCheck {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => Check' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in Check
        $AllProperties = ("id", "account", "active", "apiKey", "availability", "checkAgent", "checkInterval", "checkSpec", "checkType", "config", "container", "createIncident", "muted", "createdBy", "dateCreated", "description", "endDate", "health", "inUptime", "lastBoxStats", "lastCheckStatus", "lastError", "lastErrorDate", "lastMessage", "lastMetric", "lastRunDate", "lastStats", "lastSuccessDate", "lastTimer", "lastUpdated", "lastWarningDate", "name", "nextRunDate", "outageTime", "severity", "startDate", "deleted")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "account"))) { #optional property not found
            $Account = $null
        } else {
            $Account = $JsonParameters.PSobject.Properties["account"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "active"))) { #optional property not found
            $Active = $null
        } else {
            $Active = $JsonParameters.PSobject.Properties["active"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "apiKey"))) { #optional property not found
            $ApiKey = $null
        } else {
            $ApiKey = $JsonParameters.PSobject.Properties["apiKey"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "availability"))) { #optional property not found
            $Availability = $null
        } else {
            $Availability = $JsonParameters.PSobject.Properties["availability"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "checkAgent"))) { #optional property not found
            $CheckAgent = $null
        } else {
            $CheckAgent = $JsonParameters.PSobject.Properties["checkAgent"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "checkInterval"))) { #optional property not found
            $CheckInterval = $null
        } else {
            $CheckInterval = $JsonParameters.PSobject.Properties["checkInterval"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "checkSpec"))) { #optional property not found
            $CheckSpec = $null
        } else {
            $CheckSpec = $JsonParameters.PSobject.Properties["checkSpec"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "checkType"))) { #optional property not found
            $CheckType = $null
        } else {
            $CheckType = $JsonParameters.PSobject.Properties["checkType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) { #optional property not found
            $Config = $null
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "container"))) { #optional property not found
            $Container = $null
        } else {
            $Container = $JsonParameters.PSobject.Properties["container"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createIncident"))) { #optional property not found
            $CreateIncident = $null
        } else {
            $CreateIncident = $JsonParameters.PSobject.Properties["createIncident"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "muted"))) { #optional property not found
            $Muted = $null
        } else {
            $Muted = $JsonParameters.PSobject.Properties["muted"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createdBy"))) { #optional property not found
            $CreatedBy = $null
        } else {
            $CreatedBy = $JsonParameters.PSobject.Properties["createdBy"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dateCreated"))) { #optional property not found
            $DateCreated = $null
        } else {
            $DateCreated = $JsonParameters.PSobject.Properties["dateCreated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "endDate"))) { #optional property not found
            $EndDate = $null
        } else {
            $EndDate = $JsonParameters.PSobject.Properties["endDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "health"))) { #optional property not found
            $Health = $null
        } else {
            $Health = $JsonParameters.PSobject.Properties["health"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "inUptime"))) { #optional property not found
            $InUptime = $null
        } else {
            $InUptime = $JsonParameters.PSobject.Properties["inUptime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastBoxStats"))) { #optional property not found
            $LastBoxStats = $null
        } else {
            $LastBoxStats = $JsonParameters.PSobject.Properties["lastBoxStats"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastCheckStatus"))) { #optional property not found
            $LastCheckStatus = $null
        } else {
            $LastCheckStatus = $JsonParameters.PSobject.Properties["lastCheckStatus"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastError"))) { #optional property not found
            $LastError = $null
        } else {
            $LastError = $JsonParameters.PSobject.Properties["lastError"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastErrorDate"))) { #optional property not found
            $LastErrorDate = $null
        } else {
            $LastErrorDate = $JsonParameters.PSobject.Properties["lastErrorDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastMessage"))) { #optional property not found
            $LastMessage = $null
        } else {
            $LastMessage = $JsonParameters.PSobject.Properties["lastMessage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastMetric"))) { #optional property not found
            $LastMetric = $null
        } else {
            $LastMetric = $JsonParameters.PSobject.Properties["lastMetric"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastRunDate"))) { #optional property not found
            $LastRunDate = $null
        } else {
            $LastRunDate = $JsonParameters.PSobject.Properties["lastRunDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastStats"))) { #optional property not found
            $LastStats = $null
        } else {
            $LastStats = $JsonParameters.PSobject.Properties["lastStats"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastSuccessDate"))) { #optional property not found
            $LastSuccessDate = $null
        } else {
            $LastSuccessDate = $JsonParameters.PSobject.Properties["lastSuccessDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastTimer"))) { #optional property not found
            $LastTimer = $null
        } else {
            $LastTimer = $JsonParameters.PSobject.Properties["lastTimer"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastUpdated"))) { #optional property not found
            $LastUpdated = $null
        } else {
            $LastUpdated = $JsonParameters.PSobject.Properties["lastUpdated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastWarningDate"))) { #optional property not found
            $LastWarningDate = $null
        } else {
            $LastWarningDate = $JsonParameters.PSobject.Properties["lastWarningDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "nextRunDate"))) { #optional property not found
            $NextRunDate = $null
        } else {
            $NextRunDate = $JsonParameters.PSobject.Properties["nextRunDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "outageTime"))) { #optional property not found
            $OutageTime = $null
        } else {
            $OutageTime = $JsonParameters.PSobject.Properties["outageTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "severity"))) { #optional property not found
            $Severity = $null
        } else {
            $Severity = $JsonParameters.PSobject.Properties["severity"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "startDate"))) { #optional property not found
            $StartDate = $null
        } else {
            $StartDate = $JsonParameters.PSobject.Properties["startDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "deleted"))) { #optional property not found
            $Deleted = $null
        } else {
            $Deleted = $JsonParameters.PSobject.Properties["deleted"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "account" = ${Account}
            "active" = ${Active}
            "apiKey" = ${ApiKey}
            "availability" = ${Availability}
            "checkAgent" = ${CheckAgent}
            "checkInterval" = ${CheckInterval}
            "checkSpec" = ${CheckSpec}
            "checkType" = ${CheckType}
            "config" = ${Config}
            "container" = ${Container}
            "createIncident" = ${CreateIncident}
            "muted" = ${Muted}
            "createdBy" = ${CreatedBy}
            "dateCreated" = ${DateCreated}
            "description" = ${Description}
            "endDate" = ${EndDate}
            "health" = ${Health}
            "inUptime" = ${InUptime}
            "lastBoxStats" = ${LastBoxStats}
            "lastCheckStatus" = ${LastCheckStatus}
            "lastError" = ${LastError}
            "lastErrorDate" = ${LastErrorDate}
            "lastMessage" = ${LastMessage}
            "lastMetric" = ${LastMetric}
            "lastRunDate" = ${LastRunDate}
            "lastStats" = ${LastStats}
            "lastSuccessDate" = ${LastSuccessDate}
            "lastTimer" = ${LastTimer}
            "lastUpdated" = ${LastUpdated}
            "lastWarningDate" = ${LastWarningDate}
            "name" = ${Name}
            "nextRunDate" = ${NextRunDate}
            "outageTime" = ${OutageTime}
            "severity" = ${Severity}
            "startDate" = ${StartDate}
            "deleted" = ${Deleted}
        }

        return $PSO
    }

}

