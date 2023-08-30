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
.PARAMETER Labels
No description available.
.PARAMETER Environment
No description available.
.PARAMETER AccountId
No description available.
.PARAMETER Account
No description available.
.PARAMETER Owner
No description available.
.PARAMETER SiteId
No description available.
.PARAMETER Group
No description available.
.PARAMETER Blueprint
No description available.
.PARAMETER Type
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.PARAMETER RemovalDate
No description available.
.PARAMETER AppContext
No description available.
.PARAMETER Status
No description available.
.PARAMETER AppStatus
No description available.
.PARAMETER InstanceCount
No description available.
.PARAMETER ContainerCount
No description available.
.PARAMETER AppTiers
No description available.
.PARAMETER Instances
No description available.
.OUTPUTS

AppCreateResponse<PSCustomObject>
#>

function Initialize-AppCreateResponse {
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
        [String[]]
        ${Labels},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Environment},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${AccountId},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Account},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Owner},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${SiteId},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Group},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Blueprint},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Type},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${RemovalDate},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${AppContext},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Status},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${AppStatus},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${InstanceCount},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ContainerCount},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${AppTiers},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Instances}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AppCreateResponse' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "labels" = ${Labels}
            "environment" = ${Environment}
            "accountId" = ${AccountId}
            "account" = ${Account}
            "owner" = ${Owner}
            "siteId" = ${SiteId}
            "group" = ${Group}
            "blueprint" = ${Blueprint}
            "type" = ${Type}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "removalDate" = ${RemovalDate}
            "appContext" = ${AppContext}
            "status" = ${Status}
            "appStatus" = ${AppStatus}
            "instanceCount" = ${InstanceCount}
            "containerCount" = ${ContainerCount}
            "appTiers" = ${AppTiers}
            "instances" = ${Instances}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AppCreateResponse<PSCustomObject>

.DESCRIPTION

Convert from JSON to AppCreateResponse<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AppCreateResponse<PSCustomObject>
#>
function ConvertFrom-JsonToAppCreateResponse {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AppCreateResponse' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AppCreateResponse
        $AllProperties = ("id", "name", "description", "labels", "environment", "accountId", "account", "owner", "siteId", "group", "blueprint", "type", "dateCreated", "lastUpdated", "removalDate", "appContext", "status", "appStatus", "instanceCount", "containerCount", "appTiers", "instances")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "labels"))) { #optional property not found
            $Labels = $null
        } else {
            $Labels = $JsonParameters.PSobject.Properties["labels"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "environment"))) { #optional property not found
            $Environment = $null
        } else {
            $Environment = $JsonParameters.PSobject.Properties["environment"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "accountId"))) { #optional property not found
            $AccountId = $null
        } else {
            $AccountId = $JsonParameters.PSobject.Properties["accountId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "account"))) { #optional property not found
            $Account = $null
        } else {
            $Account = $JsonParameters.PSobject.Properties["account"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "owner"))) { #optional property not found
            $Owner = $null
        } else {
            $Owner = $JsonParameters.PSobject.Properties["owner"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "siteId"))) { #optional property not found
            $SiteId = $null
        } else {
            $SiteId = $JsonParameters.PSobject.Properties["siteId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "group"))) { #optional property not found
            $Group = $null
        } else {
            $Group = $JsonParameters.PSobject.Properties["group"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "blueprint"))) { #optional property not found
            $Blueprint = $null
        } else {
            $Blueprint = $JsonParameters.PSobject.Properties["blueprint"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "removalDate"))) { #optional property not found
            $RemovalDate = $null
        } else {
            $RemovalDate = $JsonParameters.PSobject.Properties["removalDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "appContext"))) { #optional property not found
            $AppContext = $null
        } else {
            $AppContext = $JsonParameters.PSobject.Properties["appContext"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "appStatus"))) { #optional property not found
            $AppStatus = $null
        } else {
            $AppStatus = $JsonParameters.PSobject.Properties["appStatus"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceCount"))) { #optional property not found
            $InstanceCount = $null
        } else {
            $InstanceCount = $JsonParameters.PSobject.Properties["instanceCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "containerCount"))) { #optional property not found
            $ContainerCount = $null
        } else {
            $ContainerCount = $JsonParameters.PSobject.Properties["containerCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "appTiers"))) { #optional property not found
            $AppTiers = $null
        } else {
            $AppTiers = $JsonParameters.PSobject.Properties["appTiers"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instances"))) { #optional property not found
            $Instances = $null
        } else {
            $Instances = $JsonParameters.PSobject.Properties["instances"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "labels" = ${Labels}
            "environment" = ${Environment}
            "accountId" = ${AccountId}
            "account" = ${Account}
            "owner" = ${Owner}
            "siteId" = ${SiteId}
            "group" = ${Group}
            "blueprint" = ${Blueprint}
            "type" = ${Type}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "removalDate" = ${RemovalDate}
            "appContext" = ${AppContext}
            "status" = ${Status}
            "appStatus" = ${AppStatus}
            "instanceCount" = ${InstanceCount}
            "containerCount" = ${ContainerCount}
            "appTiers" = ${AppTiers}
            "instances" = ${Instances}
        }

        return $PSO
    }

}

