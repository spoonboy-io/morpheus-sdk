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
.PARAMETER ExternalId
No description available.
.PARAMETER IacId
No description available.
.PARAMETER Zone
No description available.
.PARAMETER ZonePool
No description available.
.PARAMETER Status
No description available.
.PARAMETER Priority
No description available.
.PARAMETER GroupLayer
No description available.
.PARAMETER Rules
No description available.
.OUTPUTS

GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup<PSCustomObject>
#>

function Initialize-GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup {
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
        [String]
        ${ExternalId},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${IacId},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Zone},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ZonePool},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Status},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Priority},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GroupLayer},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Rules}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "externalId" = ${ExternalId}
            "iacId" = ${IacId}
            "zone" = ${Zone}
            "zonePool" = ${ZonePool}
            "status" = ${Status}
            "priority" = ${Priority}
            "groupLayer" = ${GroupLayer}
            "rules" = ${Rules}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkRouterFirewallRuleGroup200ResponseRuleGroup {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup
        $AllProperties = ("id", "name", "description", "externalId", "iacId", "zone", "zonePool", "status", "priority", "groupLayer", "rules")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalId"))) { #optional property not found
            $ExternalId = $null
        } else {
            $ExternalId = $JsonParameters.PSobject.Properties["externalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "iacId"))) { #optional property not found
            $IacId = $null
        } else {
            $IacId = $JsonParameters.PSobject.Properties["iacId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "zone"))) { #optional property not found
            $Zone = $null
        } else {
            $Zone = $JsonParameters.PSobject.Properties["zone"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "zonePool"))) { #optional property not found
            $ZonePool = $null
        } else {
            $ZonePool = $JsonParameters.PSobject.Properties["zonePool"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "priority"))) { #optional property not found
            $Priority = $null
        } else {
            $Priority = $JsonParameters.PSobject.Properties["priority"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "groupLayer"))) { #optional property not found
            $GroupLayer = $null
        } else {
            $GroupLayer = $JsonParameters.PSobject.Properties["groupLayer"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "rules"))) { #optional property not found
            $Rules = $null
        } else {
            $Rules = $JsonParameters.PSobject.Properties["rules"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "externalId" = ${ExternalId}
            "iacId" = ${IacId}
            "zone" = ${Zone}
            "zonePool" = ${ZonePool}
            "status" = ${Status}
            "priority" = ${Priority}
            "groupLayer" = ${GroupLayer}
            "rules" = ${Rules}
        }

        return $PSO
    }

}

