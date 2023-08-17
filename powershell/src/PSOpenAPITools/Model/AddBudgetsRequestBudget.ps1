#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.1.1
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
.PARAMETER Description
No description available.
.PARAMETER Scope
No description available.
.PARAMETER Period
No description available.
.PARAMETER Year
No description available.
.PARAMETER StartDate
No description available.
.PARAMETER EndDate
No description available.
.PARAMETER Interval
No description available.
.PARAMETER ScopeTenantId
The Tenant ID to scope to, for use with ``scope``=tenant 
.PARAMETER ScopeGroupId
The Tenant ID to scope to, for use with ``scope``=group  
.PARAMETER ScopeCloudId
The Tenant ID to scope to, for use with ``scope``=cloud 
.PARAMETER ScopeUserId
The Tenant ID to scope to, for use with ``scope``=user 
.PARAMETER Costs
No description available.
.PARAMETER Enabled
No description available.
.OUTPUTS

AddBudgetsRequestBudget<PSCustomObject>
#>

function Initialize-AddBudgetsRequestBudget {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("account", "group", "cloud", "user")]
        [String]
        ${Scope} = "account",
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Period} = "year",
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Year},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${StartDate},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${EndDate},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("year", "quarter", "month")]
        [String]
        ${Interval} = "year",
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ScopeTenantId},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ScopeGroupId},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ScopeCloudId},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ScopeUserId},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${Costs},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled} = $true
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddBudgetsRequestBudget' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "scope" = ${Scope}
            "period" = ${Period}
            "year" = ${Year}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "interval" = ${Interval}
            "scopeTenantId" = ${ScopeTenantId}
            "scopeGroupId" = ${ScopeGroupId}
            "scopeCloudId" = ${ScopeCloudId}
            "scopeUserId" = ${ScopeUserId}
            "costs" = ${Costs}
            "enabled" = ${Enabled}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddBudgetsRequestBudget<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddBudgetsRequestBudget<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddBudgetsRequestBudget<PSCustomObject>
#>
function ConvertFrom-JsonToAddBudgetsRequestBudget {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddBudgetsRequestBudget' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddBudgetsRequestBudget
        $AllProperties = ("name", "description", "scope", "period", "year", "startDate", "endDate", "interval", "scopeTenantId", "scopeGroupId", "scopeCloudId", "scopeUserId", "costs", "enabled")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'name' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property 'name' missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scope"))) { #optional property not found
            $Scope = $null
        } else {
            $Scope = $JsonParameters.PSobject.Properties["scope"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scopeTenantId"))) { #optional property not found
            $ScopeTenantId = $null
        } else {
            $ScopeTenantId = $JsonParameters.PSobject.Properties["scopeTenantId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scopeGroupId"))) { #optional property not found
            $ScopeGroupId = $null
        } else {
            $ScopeGroupId = $JsonParameters.PSobject.Properties["scopeGroupId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scopeCloudId"))) { #optional property not found
            $ScopeCloudId = $null
        } else {
            $ScopeCloudId = $JsonParameters.PSobject.Properties["scopeCloudId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scopeUserId"))) { #optional property not found
            $ScopeUserId = $null
        } else {
            $ScopeUserId = $JsonParameters.PSobject.Properties["scopeUserId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "costs"))) { #optional property not found
            $Costs = $null
        } else {
            $Costs = $JsonParameters.PSobject.Properties["costs"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "scope" = ${Scope}
            "period" = ${Period}
            "year" = ${Year}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "interval" = ${Interval}
            "scopeTenantId" = ${ScopeTenantId}
            "scopeGroupId" = ${ScopeGroupId}
            "scopeCloudId" = ${ScopeCloudId}
            "scopeUserId" = ${ScopeUserId}
            "costs" = ${Costs}
            "enabled" = ${Enabled}
        }

        return $PSO
    }

}

