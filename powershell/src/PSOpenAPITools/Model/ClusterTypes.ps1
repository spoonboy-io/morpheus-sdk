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
.PARAMETER DeployTargetService
No description available.
.PARAMETER ShortName
No description available.
.PARAMETER ProviderType
No description available.
.PARAMETER Code
No description available.
.PARAMETER HostService
No description available.
.PARAMETER Managed
No description available.
.PARAMETER HasMasters
No description available.
.PARAMETER HasWorkers
No description available.
.PARAMETER ViewSet
No description available.
.PARAMETER ImageCode
No description available.
.PARAMETER KubeCtlLocal
No description available.
.PARAMETER HasDatastore
No description available.
.PARAMETER SupportsCloudScaling
No description available.
.PARAMETER Name
No description available.
.PARAMETER HasDefaultDataDisk
No description available.
.PARAMETER CanManage
No description available.
.PARAMETER HasCluster
No description available.
.PARAMETER Description
No description available.
.PARAMETER OptionTypes
No description available.
.PARAMETER ControllerTypes
No description available.
.PARAMETER WorkerTypes
No description available.
.OUTPUTS

ClusterTypes<PSCustomObject>
#>

function Initialize-ClusterTypes {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DeployTargetService},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ShortName},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ProviderType},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Code},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${HostService},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Managed},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${HasMasters},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${HasWorkers},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ViewSet},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ImageCode},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${KubeCtlLocal},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${HasDatastore},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${SupportsCloudScaling},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${HasDefaultDataDisk},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${CanManage},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${HasCluster},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${OptionTypes},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${ControllerTypes},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${WorkerTypes}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterTypes' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "deployTargetService" = ${DeployTargetService}
            "shortName" = ${ShortName}
            "providerType" = ${ProviderType}
            "code" = ${Code}
            "hostService" = ${HostService}
            "managed" = ${Managed}
            "hasMasters" = ${HasMasters}
            "hasWorkers" = ${HasWorkers}
            "viewSet" = ${ViewSet}
            "imageCode" = ${ImageCode}
            "kubeCtlLocal" = ${KubeCtlLocal}
            "hasDatastore" = ${HasDatastore}
            "supportsCloudScaling" = ${SupportsCloudScaling}
            "name" = ${Name}
            "hasDefaultDataDisk" = ${HasDefaultDataDisk}
            "canManage" = ${CanManage}
            "hasCluster" = ${HasCluster}
            "description" = ${Description}
            "optionTypes" = ${OptionTypes}
            "controllerTypes" = ${ControllerTypes}
            "workerTypes" = ${WorkerTypes}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterTypes<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterTypes<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterTypes<PSCustomObject>
#>
function ConvertFrom-JsonToClusterTypes {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterTypes' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterTypes
        $AllProperties = ("id", "deployTargetService", "shortName", "providerType", "code", "hostService", "managed", "hasMasters", "hasWorkers", "viewSet", "imageCode", "kubeCtlLocal", "hasDatastore", "supportsCloudScaling", "name", "hasDefaultDataDisk", "canManage", "hasCluster", "description", "optionTypes", "controllerTypes", "workerTypes")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "deployTargetService"))) { #optional property not found
            $DeployTargetService = $null
        } else {
            $DeployTargetService = $JsonParameters.PSobject.Properties["deployTargetService"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "shortName"))) { #optional property not found
            $ShortName = $null
        } else {
            $ShortName = $JsonParameters.PSobject.Properties["shortName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "providerType"))) { #optional property not found
            $ProviderType = $null
        } else {
            $ProviderType = $JsonParameters.PSobject.Properties["providerType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) { #optional property not found
            $Code = $null
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hostService"))) { #optional property not found
            $HostService = $null
        } else {
            $HostService = $JsonParameters.PSobject.Properties["hostService"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "managed"))) { #optional property not found
            $Managed = $null
        } else {
            $Managed = $JsonParameters.PSobject.Properties["managed"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hasMasters"))) { #optional property not found
            $HasMasters = $null
        } else {
            $HasMasters = $JsonParameters.PSobject.Properties["hasMasters"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hasWorkers"))) { #optional property not found
            $HasWorkers = $null
        } else {
            $HasWorkers = $JsonParameters.PSobject.Properties["hasWorkers"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "viewSet"))) { #optional property not found
            $ViewSet = $null
        } else {
            $ViewSet = $JsonParameters.PSobject.Properties["viewSet"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "imageCode"))) { #optional property not found
            $ImageCode = $null
        } else {
            $ImageCode = $JsonParameters.PSobject.Properties["imageCode"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "kubeCtlLocal"))) { #optional property not found
            $KubeCtlLocal = $null
        } else {
            $KubeCtlLocal = $JsonParameters.PSobject.Properties["kubeCtlLocal"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hasDatastore"))) { #optional property not found
            $HasDatastore = $null
        } else {
            $HasDatastore = $JsonParameters.PSobject.Properties["hasDatastore"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "supportsCloudScaling"))) { #optional property not found
            $SupportsCloudScaling = $null
        } else {
            $SupportsCloudScaling = $JsonParameters.PSobject.Properties["supportsCloudScaling"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hasDefaultDataDisk"))) { #optional property not found
            $HasDefaultDataDisk = $null
        } else {
            $HasDefaultDataDisk = $JsonParameters.PSobject.Properties["hasDefaultDataDisk"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "canManage"))) { #optional property not found
            $CanManage = $null
        } else {
            $CanManage = $JsonParameters.PSobject.Properties["canManage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hasCluster"))) { #optional property not found
            $HasCluster = $null
        } else {
            $HasCluster = $JsonParameters.PSobject.Properties["hasCluster"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "optionTypes"))) { #optional property not found
            $OptionTypes = $null
        } else {
            $OptionTypes = $JsonParameters.PSobject.Properties["optionTypes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "controllerTypes"))) { #optional property not found
            $ControllerTypes = $null
        } else {
            $ControllerTypes = $JsonParameters.PSobject.Properties["controllerTypes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "workerTypes"))) { #optional property not found
            $WorkerTypes = $null
        } else {
            $WorkerTypes = $JsonParameters.PSobject.Properties["workerTypes"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "deployTargetService" = ${DeployTargetService}
            "shortName" = ${ShortName}
            "providerType" = ${ProviderType}
            "code" = ${Code}
            "hostService" = ${HostService}
            "managed" = ${Managed}
            "hasMasters" = ${HasMasters}
            "hasWorkers" = ${HasWorkers}
            "viewSet" = ${ViewSet}
            "imageCode" = ${ImageCode}
            "kubeCtlLocal" = ${KubeCtlLocal}
            "hasDatastore" = ${HasDatastore}
            "supportsCloudScaling" = ${SupportsCloudScaling}
            "name" = ${Name}
            "hasDefaultDataDisk" = ${HasDefaultDataDisk}
            "canManage" = ${CanManage}
            "hasCluster" = ${HasCluster}
            "description" = ${Description}
            "optionTypes" = ${OptionTypes}
            "controllerTypes" = ${ControllerTypes}
            "workerTypes" = ${WorkerTypes}
        }

        return $PSO
    }

}

