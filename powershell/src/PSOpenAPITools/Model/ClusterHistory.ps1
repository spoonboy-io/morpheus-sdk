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
.PARAMETER AccountId
No description available.
.PARAMETER UniqueId
No description available.
.PARAMETER ProcessType
No description available.
.PARAMETER DisplayName
No description available.
.PARAMETER Description
No description available.
.PARAMETER SubType
No description available.
.PARAMETER SubId
No description available.
.PARAMETER ZoneId
No description available.
.PARAMETER IntegrationId
No description available.
.PARAMETER AppId
No description available.
.PARAMETER InstanceId
No description available.
.PARAMETER ContainerId
No description available.
.PARAMETER ServerId
No description available.
.PARAMETER ContainerName
No description available.
.PARAMETER Status
No description available.
.PARAMETER Reason
No description available.
.PARAMETER Percent
No description available.
.PARAMETER StatusEta
No description available.
.PARAMETER Message
No description available.
.PARAMETER Output
No description available.
.PARAMETER VarError
No description available.
.PARAMETER StartDate
No description available.
.PARAMETER EndDate
No description available.
.PARAMETER Duration
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.PARAMETER CreatedBy
No description available.
.PARAMETER UpdatedBy
No description available.
.PARAMETER Events
No description available.
.OUTPUTS

ClusterHistory<PSCustomObject>
#>

function Initialize-ClusterHistory {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${AccountId},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${UniqueId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ProcessType},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DisplayName},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SubType},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SubId},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ZoneId},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${IntegrationId},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${AppId},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${InstanceId},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ContainerId},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ServerId},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ContainerName},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Status},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Reason},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Percent},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${StatusEta},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Message},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Output},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${VarError},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${StartDate},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${EndDate},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Duration},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 26, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated},
        [Parameter(Position = 27, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CreatedBy},
        [Parameter(Position = 28, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${UpdatedBy},
        [Parameter(Position = 29, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Events}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterHistory' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "accountId" = ${AccountId}
            "uniqueId" = ${UniqueId}
            "processType" = ${ProcessType}
            "displayName" = ${DisplayName}
            "description" = ${Description}
            "subType" = ${SubType}
            "subId" = ${SubId}
            "zoneId" = ${ZoneId}
            "integrationId" = ${IntegrationId}
            "appId" = ${AppId}
            "instanceId" = ${InstanceId}
            "containerId" = ${ContainerId}
            "serverId" = ${ServerId}
            "containerName" = ${ContainerName}
            "status" = ${Status}
            "reason" = ${Reason}
            "percent" = ${Percent}
            "statusEta" = ${StatusEta}
            "message" = ${Message}
            "output" = ${Output}
            "error" = ${VarError}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "duration" = ${Duration}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "createdBy" = ${CreatedBy}
            "updatedBy" = ${UpdatedBy}
            "events" = ${Events}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterHistory<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterHistory<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterHistory<PSCustomObject>
#>
function ConvertFrom-JsonToClusterHistory {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterHistory' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterHistory
        $AllProperties = ("id", "accountId", "uniqueId", "processType", "displayName", "description", "subType", "subId", "zoneId", "integrationId", "appId", "instanceId", "containerId", "serverId", "containerName", "status", "reason", "percent", "statusEta", "message", "output", "error", "startDate", "endDate", "duration", "dateCreated", "lastUpdated", "createdBy", "updatedBy", "events")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "accountId"))) { #optional property not found
            $AccountId = $null
        } else {
            $AccountId = $JsonParameters.PSobject.Properties["accountId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "uniqueId"))) { #optional property not found
            $UniqueId = $null
        } else {
            $UniqueId = $JsonParameters.PSobject.Properties["uniqueId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "processType"))) { #optional property not found
            $ProcessType = $null
        } else {
            $ProcessType = $JsonParameters.PSobject.Properties["processType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "displayName"))) { #optional property not found
            $DisplayName = $null
        } else {
            $DisplayName = $JsonParameters.PSobject.Properties["displayName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "subType"))) { #optional property not found
            $SubType = $null
        } else {
            $SubType = $JsonParameters.PSobject.Properties["subType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "subId"))) { #optional property not found
            $SubId = $null
        } else {
            $SubId = $JsonParameters.PSobject.Properties["subId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "zoneId"))) { #optional property not found
            $ZoneId = $null
        } else {
            $ZoneId = $JsonParameters.PSobject.Properties["zoneId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "integrationId"))) { #optional property not found
            $IntegrationId = $null
        } else {
            $IntegrationId = $JsonParameters.PSobject.Properties["integrationId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "appId"))) { #optional property not found
            $AppId = $null
        } else {
            $AppId = $JsonParameters.PSobject.Properties["appId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceId"))) { #optional property not found
            $InstanceId = $null
        } else {
            $InstanceId = $JsonParameters.PSobject.Properties["instanceId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "containerId"))) { #optional property not found
            $ContainerId = $null
        } else {
            $ContainerId = $JsonParameters.PSobject.Properties["containerId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serverId"))) { #optional property not found
            $ServerId = $null
        } else {
            $ServerId = $JsonParameters.PSobject.Properties["serverId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "containerName"))) { #optional property not found
            $ContainerName = $null
        } else {
            $ContainerName = $JsonParameters.PSobject.Properties["containerName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "reason"))) { #optional property not found
            $Reason = $null
        } else {
            $Reason = $JsonParameters.PSobject.Properties["reason"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "percent"))) { #optional property not found
            $Percent = $null
        } else {
            $Percent = $JsonParameters.PSobject.Properties["percent"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "statusEta"))) { #optional property not found
            $StatusEta = $null
        } else {
            $StatusEta = $JsonParameters.PSobject.Properties["statusEta"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "message"))) { #optional property not found
            $Message = $null
        } else {
            $Message = $JsonParameters.PSobject.Properties["message"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "output"))) { #optional property not found
            $Output = $null
        } else {
            $Output = $JsonParameters.PSobject.Properties["output"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "error"))) { #optional property not found
            $VarError = $null
        } else {
            $VarError = $JsonParameters.PSobject.Properties["error"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "startDate"))) { #optional property not found
            $StartDate = $null
        } else {
            $StartDate = $JsonParameters.PSobject.Properties["startDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "endDate"))) { #optional property not found
            $EndDate = $null
        } else {
            $EndDate = $JsonParameters.PSobject.Properties["endDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "duration"))) { #optional property not found
            $Duration = $null
        } else {
            $Duration = $JsonParameters.PSobject.Properties["duration"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dateCreated"))) { #optional property not found
            $DateCreated = $null
        } else {
            $DateCreated = $JsonParameters.PSobject.Properties["dateCreated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastUpdated"))) { #optional property not found
            $LastUpdated = $null
        } else {
            $LastUpdated = $JsonParameters.PSobject.Properties["lastUpdated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createdBy"))) { #optional property not found
            $CreatedBy = $null
        } else {
            $CreatedBy = $JsonParameters.PSobject.Properties["createdBy"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "updatedBy"))) { #optional property not found
            $UpdatedBy = $null
        } else {
            $UpdatedBy = $JsonParameters.PSobject.Properties["updatedBy"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "events"))) { #optional property not found
            $Events = $null
        } else {
            $Events = $JsonParameters.PSobject.Properties["events"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "accountId" = ${AccountId}
            "uniqueId" = ${UniqueId}
            "processType" = ${ProcessType}
            "displayName" = ${DisplayName}
            "description" = ${Description}
            "subType" = ${SubType}
            "subId" = ${SubId}
            "zoneId" = ${ZoneId}
            "integrationId" = ${IntegrationId}
            "appId" = ${AppId}
            "instanceId" = ${InstanceId}
            "containerId" = ${ContainerId}
            "serverId" = ${ServerId}
            "containerName" = ${ContainerName}
            "status" = ${Status}
            "reason" = ${Reason}
            "percent" = ${Percent}
            "statusEta" = ${StatusEta}
            "message" = ${Message}
            "output" = ${Output}
            "error" = ${VarError}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "duration" = ${Duration}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "createdBy" = ${CreatedBy}
            "updatedBy" = ${UpdatedBy}
            "events" = ${Events}
        }

        return $PSO
    }

}
