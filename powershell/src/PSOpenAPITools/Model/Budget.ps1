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
.PARAMETER Name
No description available.
.PARAMETER Description
No description available.
.PARAMETER Account
No description available.
.PARAMETER Enabled
No description available.
.PARAMETER RefScope
No description available.
.PARAMETER RefType
No description available.
.PARAMETER RefId
No description available.
.PARAMETER RefName
No description available.
.PARAMETER Period
No description available.
.PARAMETER Year
No description available.
.PARAMETER ResourceType
No description available.
.PARAMETER Timezone
No description available.
.PARAMETER StartDate
No description available.
.PARAMETER EndDate
No description available.
.PARAMETER Interval
No description available.
.PARAMETER Costs
No description available.
.PARAMETER IsFiscal
No description available.
.PARAMETER AverageCost
No description available.
.PARAMETER TotalCost
No description available.
.PARAMETER Currency
No description available.
.PARAMETER Rollover
No description available.
.PARAMETER WarningLimit
No description available.
.PARAMETER OverLimit
No description available.
.PARAMETER ExternalId
No description available.
.PARAMETER InternalId
No description available.
.PARAMETER CreatedById
No description available.
.PARAMETER CreatedByName
No description available.
.PARAMETER UpdatedById
No description available.
.PARAMETER UpdatedByName
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.PARAMETER Stats
No description available.
.OUTPUTS

Budget<PSCustomObject>
#>

function Initialize-Budget {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Account},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RefScope},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RefType},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${RefId},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RefName},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Period},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Year},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ResourceType},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Timezone},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${StartDate},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${EndDate},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Interval},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${Costs},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${IsFiscal},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${AverageCost},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${TotalCost},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Currency},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Rollover},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${WarningLimit},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${OverLimit},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalId},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InternalId},
        [Parameter(Position = 26, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${CreatedById},
        [Parameter(Position = 27, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CreatedByName},
        [Parameter(Position = 28, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${UpdatedById},
        [Parameter(Position = 29, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${UpdatedByName},
        [Parameter(Position = 30, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 31, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated},
        [Parameter(Position = 32, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Stats}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => Budget' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "account" = ${Account}
            "enabled" = ${Enabled}
            "refScope" = ${RefScope}
            "refType" = ${RefType}
            "refId" = ${RefId}
            "refName" = ${RefName}
            "period" = ${Period}
            "year" = ${Year}
            "resourceType" = ${ResourceType}
            "timezone" = ${Timezone}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "interval" = ${Interval}
            "costs" = ${Costs}
            "isFiscal" = ${IsFiscal}
            "averageCost" = ${AverageCost}
            "totalCost" = ${TotalCost}
            "currency" = ${Currency}
            "rollover" = ${Rollover}
            "warningLimit" = ${WarningLimit}
            "overLimit" = ${OverLimit}
            "externalId" = ${ExternalId}
            "internalId" = ${InternalId}
            "createdById" = ${CreatedById}
            "createdByName" = ${CreatedByName}
            "updatedById" = ${UpdatedById}
            "updatedByName" = ${UpdatedByName}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "stats" = ${Stats}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to Budget<PSCustomObject>

.DESCRIPTION

Convert from JSON to Budget<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

Budget<PSCustomObject>
#>
function ConvertFrom-JsonToBudget {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => Budget' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in Budget
        $AllProperties = ("id", "name", "description", "account", "enabled", "refScope", "refType", "refId", "refName", "period", "year", "resourceType", "timezone", "startDate", "endDate", "interval", "costs", "isFiscal", "averageCost", "totalCost", "currency", "rollover", "warningLimit", "overLimit", "externalId", "internalId", "createdById", "createdByName", "updatedById", "updatedByName", "dateCreated", "lastUpdated", "stats")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "account"))) { #optional property not found
            $Account = $null
        } else {
            $Account = $JsonParameters.PSobject.Properties["account"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "refScope"))) { #optional property not found
            $RefScope = $null
        } else {
            $RefScope = $JsonParameters.PSobject.Properties["refScope"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "refType"))) { #optional property not found
            $RefType = $null
        } else {
            $RefType = $JsonParameters.PSobject.Properties["refType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "refId"))) { #optional property not found
            $RefId = $null
        } else {
            $RefId = $JsonParameters.PSobject.Properties["refId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "refName"))) { #optional property not found
            $RefName = $null
        } else {
            $RefName = $JsonParameters.PSobject.Properties["refName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "period"))) { #optional property not found
            $Period = $null
        } else {
            $Period = $JsonParameters.PSobject.Properties["period"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "year"))) { #optional property not found
            $Year = $null
        } else {
            $Year = $JsonParameters.PSobject.Properties["year"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "resourceType"))) { #optional property not found
            $ResourceType = $null
        } else {
            $ResourceType = $JsonParameters.PSobject.Properties["resourceType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "timezone"))) { #optional property not found
            $Timezone = $null
        } else {
            $Timezone = $JsonParameters.PSobject.Properties["timezone"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "interval"))) { #optional property not found
            $Interval = $null
        } else {
            $Interval = $JsonParameters.PSobject.Properties["interval"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "costs"))) { #optional property not found
            $Costs = $null
        } else {
            $Costs = $JsonParameters.PSobject.Properties["costs"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "isFiscal"))) { #optional property not found
            $IsFiscal = $null
        } else {
            $IsFiscal = $JsonParameters.PSobject.Properties["isFiscal"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "averageCost"))) { #optional property not found
            $AverageCost = $null
        } else {
            $AverageCost = $JsonParameters.PSobject.Properties["averageCost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "totalCost"))) { #optional property not found
            $TotalCost = $null
        } else {
            $TotalCost = $JsonParameters.PSobject.Properties["totalCost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "currency"))) { #optional property not found
            $Currency = $null
        } else {
            $Currency = $JsonParameters.PSobject.Properties["currency"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "rollover"))) { #optional property not found
            $Rollover = $null
        } else {
            $Rollover = $JsonParameters.PSobject.Properties["rollover"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "warningLimit"))) { #optional property not found
            $WarningLimit = $null
        } else {
            $WarningLimit = $JsonParameters.PSobject.Properties["warningLimit"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "overLimit"))) { #optional property not found
            $OverLimit = $null
        } else {
            $OverLimit = $JsonParameters.PSobject.Properties["overLimit"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalId"))) { #optional property not found
            $ExternalId = $null
        } else {
            $ExternalId = $JsonParameters.PSobject.Properties["externalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "internalId"))) { #optional property not found
            $InternalId = $null
        } else {
            $InternalId = $JsonParameters.PSobject.Properties["internalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createdById"))) { #optional property not found
            $CreatedById = $null
        } else {
            $CreatedById = $JsonParameters.PSobject.Properties["createdById"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createdByName"))) { #optional property not found
            $CreatedByName = $null
        } else {
            $CreatedByName = $JsonParameters.PSobject.Properties["createdByName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "updatedById"))) { #optional property not found
            $UpdatedById = $null
        } else {
            $UpdatedById = $JsonParameters.PSobject.Properties["updatedById"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "updatedByName"))) { #optional property not found
            $UpdatedByName = $null
        } else {
            $UpdatedByName = $JsonParameters.PSobject.Properties["updatedByName"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "stats"))) { #optional property not found
            $Stats = $null
        } else {
            $Stats = $JsonParameters.PSobject.Properties["stats"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "account" = ${Account}
            "enabled" = ${Enabled}
            "refScope" = ${RefScope}
            "refType" = ${RefType}
            "refId" = ${RefId}
            "refName" = ${RefName}
            "period" = ${Period}
            "year" = ${Year}
            "resourceType" = ${ResourceType}
            "timezone" = ${Timezone}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "interval" = ${Interval}
            "costs" = ${Costs}
            "isFiscal" = ${IsFiscal}
            "averageCost" = ${AverageCost}
            "totalCost" = ${TotalCost}
            "currency" = ${Currency}
            "rollover" = ${Rollover}
            "warningLimit" = ${WarningLimit}
            "overLimit" = ${OverLimit}
            "externalId" = ${ExternalId}
            "internalId" = ${InternalId}
            "createdById" = ${CreatedById}
            "createdByName" = ${CreatedByName}
            "updatedById" = ${UpdatedById}
            "updatedByName" = ${UpdatedByName}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "stats" = ${Stats}
        }

        return $PSO
    }

}
