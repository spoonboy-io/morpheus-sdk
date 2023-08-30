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
No description available.
.PARAMETER InstanceName
No description available.
.PARAMETER ZoneName
No description available.
.PARAMETER AccountName
No description available.
.PARAMETER Volumes
No description available.
.PARAMETER MaxMemory
No description available.
.PARAMETER MaxCpu
No description available.
.PARAMETER MaxCores
No description available.
.PARAMETER ServerExternalId
No description available.
.PARAMETER ServerInternalId
No description available.
.PARAMETER PlanName
No description available.
.PARAMETER HourlyPrice
No description available.
.PARAMETER HourlyCost
No description available.
.PARAMETER Currency
No description available.
.PARAMETER PricesUsed
No description available.
.PARAMETER Cost
No description available.
.PARAMETER Price
No description available.
.PARAMETER CreatedByUser
No description available.
.PARAMETER CreatedByUserId
No description available.
.PARAMETER SiteId
No description available.
.PARAMETER SiteName
No description available.
.PARAMETER SiteUUID
No description available.
.PARAMETER SiteCode
No description available.
.PARAMETER StartDate
No description available.
.PARAMETER EndDate
No description available.
.PARAMETER Status
No description available.
.PARAMETER Tags
No description available.
.PARAMETER ApplicablePrices
No description available.
.PARAMETER ServicePlanId
No description available.
.PARAMETER ServicePlanName
No description available.
.PARAMETER ResourcePoolId
No description available.
.PARAMETER ResourcePoolName
No description available.
.OUTPUTS

BillingInstanceContainersInnerUsagesInner<PSCustomObject>
#>

function Initialize-BillingInstanceContainersInnerUsagesInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InstanceName},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ZoneName},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${AccountName},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Volumes},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxMemory},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${MaxCpu},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxCores},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServerExternalId},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServerInternalId},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PlanName},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${HourlyPrice},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${HourlyCost},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Currency},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${PricesUsed},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Cost},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Price},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CreatedByUser},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${CreatedByUserId},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${SiteId},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SiteName},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SiteUUID},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SiteCode},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${StartDate},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${EndDate},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Status},
        [Parameter(Position = 26, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Tags},
        [Parameter(Position = 27, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${ApplicablePrices},
        [Parameter(Position = 28, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ServicePlanId},
        [Parameter(Position = 29, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServicePlanName},
        [Parameter(Position = 30, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ResourcePoolId},
        [Parameter(Position = 31, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ResourcePoolName}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BillingInstanceContainersInnerUsagesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "instanceName" = ${InstanceName}
            "zoneName" = ${ZoneName}
            "accountName" = ${AccountName}
            "volumes" = ${Volumes}
            "maxMemory" = ${MaxMemory}
            "maxCpu" = ${MaxCpu}
            "maxCores" = ${MaxCores}
            "serverExternalId" = ${ServerExternalId}
            "serverInternalId" = ${ServerInternalId}
            "planName" = ${PlanName}
            "hourlyPrice" = ${HourlyPrice}
            "hourlyCost" = ${HourlyCost}
            "currency" = ${Currency}
            "pricesUsed" = ${PricesUsed}
            "cost" = ${Cost}
            "price" = ${Price}
            "createdByUser" = ${CreatedByUser}
            "createdByUserId" = ${CreatedByUserId}
            "siteId" = ${SiteId}
            "siteName" = ${SiteName}
            "siteUUID" = ${SiteUUID}
            "siteCode" = ${SiteCode}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "status" = ${Status}
            "tags" = ${Tags}
            "applicablePrices" = ${ApplicablePrices}
            "servicePlanId" = ${ServicePlanId}
            "servicePlanName" = ${ServicePlanName}
            "resourcePoolId" = ${ResourcePoolId}
            "resourcePoolName" = ${ResourcePoolName}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BillingInstanceContainersInnerUsagesInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to BillingInstanceContainersInnerUsagesInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BillingInstanceContainersInnerUsagesInner<PSCustomObject>
#>
function ConvertFrom-JsonToBillingInstanceContainersInnerUsagesInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BillingInstanceContainersInnerUsagesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BillingInstanceContainersInnerUsagesInner
        $AllProperties = ("name", "instanceName", "zoneName", "accountName", "volumes", "maxMemory", "maxCpu", "maxCores", "serverExternalId", "serverInternalId", "planName", "hourlyPrice", "hourlyCost", "currency", "pricesUsed", "cost", "price", "createdByUser", "createdByUserId", "siteId", "siteName", "siteUUID", "siteCode", "startDate", "endDate", "status", "tags", "applicablePrices", "servicePlanId", "servicePlanName", "resourcePoolId", "resourcePoolName")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceName"))) { #optional property not found
            $InstanceName = $null
        } else {
            $InstanceName = $JsonParameters.PSobject.Properties["instanceName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "zoneName"))) { #optional property not found
            $ZoneName = $null
        } else {
            $ZoneName = $JsonParameters.PSobject.Properties["zoneName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "accountName"))) { #optional property not found
            $AccountName = $null
        } else {
            $AccountName = $JsonParameters.PSobject.Properties["accountName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "volumes"))) { #optional property not found
            $Volumes = $null
        } else {
            $Volumes = $JsonParameters.PSobject.Properties["volumes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxMemory"))) { #optional property not found
            $MaxMemory = $null
        } else {
            $MaxMemory = $JsonParameters.PSobject.Properties["maxMemory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxCpu"))) { #optional property not found
            $MaxCpu = $null
        } else {
            $MaxCpu = $JsonParameters.PSobject.Properties["maxCpu"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxCores"))) { #optional property not found
            $MaxCores = $null
        } else {
            $MaxCores = $JsonParameters.PSobject.Properties["maxCores"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serverExternalId"))) { #optional property not found
            $ServerExternalId = $null
        } else {
            $ServerExternalId = $JsonParameters.PSobject.Properties["serverExternalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serverInternalId"))) { #optional property not found
            $ServerInternalId = $null
        } else {
            $ServerInternalId = $JsonParameters.PSobject.Properties["serverInternalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "planName"))) { #optional property not found
            $PlanName = $null
        } else {
            $PlanName = $JsonParameters.PSobject.Properties["planName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hourlyPrice"))) { #optional property not found
            $HourlyPrice = $null
        } else {
            $HourlyPrice = $JsonParameters.PSobject.Properties["hourlyPrice"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hourlyCost"))) { #optional property not found
            $HourlyCost = $null
        } else {
            $HourlyCost = $JsonParameters.PSobject.Properties["hourlyCost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "currency"))) { #optional property not found
            $Currency = $null
        } else {
            $Currency = $JsonParameters.PSobject.Properties["currency"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "pricesUsed"))) { #optional property not found
            $PricesUsed = $null
        } else {
            $PricesUsed = $JsonParameters.PSobject.Properties["pricesUsed"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cost"))) { #optional property not found
            $Cost = $null
        } else {
            $Cost = $JsonParameters.PSobject.Properties["cost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "price"))) { #optional property not found
            $Price = $null
        } else {
            $Price = $JsonParameters.PSobject.Properties["price"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createdByUser"))) { #optional property not found
            $CreatedByUser = $null
        } else {
            $CreatedByUser = $JsonParameters.PSobject.Properties["createdByUser"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createdByUserId"))) { #optional property not found
            $CreatedByUserId = $null
        } else {
            $CreatedByUserId = $JsonParameters.PSobject.Properties["createdByUserId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "siteId"))) { #optional property not found
            $SiteId = $null
        } else {
            $SiteId = $JsonParameters.PSobject.Properties["siteId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "siteName"))) { #optional property not found
            $SiteName = $null
        } else {
            $SiteName = $JsonParameters.PSobject.Properties["siteName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "siteUUID"))) { #optional property not found
            $SiteUUID = $null
        } else {
            $SiteUUID = $JsonParameters.PSobject.Properties["siteUUID"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "siteCode"))) { #optional property not found
            $SiteCode = $null
        } else {
            $SiteCode = $JsonParameters.PSobject.Properties["siteCode"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tags"))) { #optional property not found
            $Tags = $null
        } else {
            $Tags = $JsonParameters.PSobject.Properties["tags"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "applicablePrices"))) { #optional property not found
            $ApplicablePrices = $null
        } else {
            $ApplicablePrices = $JsonParameters.PSobject.Properties["applicablePrices"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "servicePlanId"))) { #optional property not found
            $ServicePlanId = $null
        } else {
            $ServicePlanId = $JsonParameters.PSobject.Properties["servicePlanId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "servicePlanName"))) { #optional property not found
            $ServicePlanName = $null
        } else {
            $ServicePlanName = $JsonParameters.PSobject.Properties["servicePlanName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "resourcePoolId"))) { #optional property not found
            $ResourcePoolId = $null
        } else {
            $ResourcePoolId = $JsonParameters.PSobject.Properties["resourcePoolId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "resourcePoolName"))) { #optional property not found
            $ResourcePoolName = $null
        } else {
            $ResourcePoolName = $JsonParameters.PSobject.Properties["resourcePoolName"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "instanceName" = ${InstanceName}
            "zoneName" = ${ZoneName}
            "accountName" = ${AccountName}
            "volumes" = ${Volumes}
            "maxMemory" = ${MaxMemory}
            "maxCpu" = ${MaxCpu}
            "maxCores" = ${MaxCores}
            "serverExternalId" = ${ServerExternalId}
            "serverInternalId" = ${ServerInternalId}
            "planName" = ${PlanName}
            "hourlyPrice" = ${HourlyPrice}
            "hourlyCost" = ${HourlyCost}
            "currency" = ${Currency}
            "pricesUsed" = ${PricesUsed}
            "cost" = ${Cost}
            "price" = ${Price}
            "createdByUser" = ${CreatedByUser}
            "createdByUserId" = ${CreatedByUserId}
            "siteId" = ${SiteId}
            "siteName" = ${SiteName}
            "siteUUID" = ${SiteUUID}
            "siteCode" = ${SiteCode}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "status" = ${Status}
            "tags" = ${Tags}
            "applicablePrices" = ${ApplicablePrices}
            "servicePlanId" = ${ServicePlanId}
            "servicePlanName" = ${ServicePlanName}
            "resourcePoolId" = ${ResourcePoolId}
            "resourcePoolName" = ${ResourcePoolName}
        }

        return $PSO
    }

}
