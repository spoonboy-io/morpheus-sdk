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

.PARAMETER Status
No description available.
.PARAMETER ClusterName
No description available.
.PARAMETER NodeTotal
No description available.
.PARAMETER NodeData
No description available.
.PARAMETER Shards
No description available.
.PARAMETER Primary
No description available.
.PARAMETER Relocating
No description available.
.PARAMETER Initializing
No description available.
.PARAMETER Unassigned
No description available.
.PARAMETER PendingTasks
No description available.
.PARAMETER ActivePercent
No description available.
.OUTPUTS

HealthElasticStats<PSCustomObject>
#>

function Initialize-HealthElasticStats {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Status},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ClusterName},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${NodeTotal},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${NodeData},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Shards},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Primary},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Relocating},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Initializing},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Unassigned},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PendingTasks},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ActivePercent}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => HealthElasticStats' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "status" = ${Status}
            "clusterName" = ${ClusterName}
            "nodeTotal" = ${NodeTotal}
            "nodeData" = ${NodeData}
            "shards" = ${Shards}
            "primary" = ${Primary}
            "relocating" = ${Relocating}
            "initializing" = ${Initializing}
            "unassigned" = ${Unassigned}
            "pendingTasks" = ${PendingTasks}
            "activePercent" = ${ActivePercent}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to HealthElasticStats<PSCustomObject>

.DESCRIPTION

Convert from JSON to HealthElasticStats<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

HealthElasticStats<PSCustomObject>
#>
function ConvertFrom-JsonToHealthElasticStats {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => HealthElasticStats' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in HealthElasticStats
        $AllProperties = ("status", "clusterName", "nodeTotal", "nodeData", "shards", "primary", "relocating", "initializing", "unassigned", "pendingTasks", "activePercent")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clusterName"))) { #optional property not found
            $ClusterName = $null
        } else {
            $ClusterName = $JsonParameters.PSobject.Properties["clusterName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "nodeTotal"))) { #optional property not found
            $NodeTotal = $null
        } else {
            $NodeTotal = $JsonParameters.PSobject.Properties["nodeTotal"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "nodeData"))) { #optional property not found
            $NodeData = $null
        } else {
            $NodeData = $JsonParameters.PSobject.Properties["nodeData"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "shards"))) { #optional property not found
            $Shards = $null
        } else {
            $Shards = $JsonParameters.PSobject.Properties["shards"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "primary"))) { #optional property not found
            $Primary = $null
        } else {
            $Primary = $JsonParameters.PSobject.Properties["primary"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "relocating"))) { #optional property not found
            $Relocating = $null
        } else {
            $Relocating = $JsonParameters.PSobject.Properties["relocating"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "initializing"))) { #optional property not found
            $Initializing = $null
        } else {
            $Initializing = $JsonParameters.PSobject.Properties["initializing"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "unassigned"))) { #optional property not found
            $Unassigned = $null
        } else {
            $Unassigned = $JsonParameters.PSobject.Properties["unassigned"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "pendingTasks"))) { #optional property not found
            $PendingTasks = $null
        } else {
            $PendingTasks = $JsonParameters.PSobject.Properties["pendingTasks"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "activePercent"))) { #optional property not found
            $ActivePercent = $null
        } else {
            $ActivePercent = $JsonParameters.PSobject.Properties["activePercent"].value
        }

        $PSO = [PSCustomObject]@{
            "status" = ${Status}
            "clusterName" = ${ClusterName}
            "nodeTotal" = ${NodeTotal}
            "nodeData" = ${NodeData}
            "shards" = ${Shards}
            "primary" = ${Primary}
            "relocating" = ${Relocating}
            "initializing" = ${Initializing}
            "unassigned" = ${Unassigned}
            "pendingTasks" = ${PendingTasks}
            "activePercent" = ${ActivePercent}
        }

        return $PSO
    }

}
