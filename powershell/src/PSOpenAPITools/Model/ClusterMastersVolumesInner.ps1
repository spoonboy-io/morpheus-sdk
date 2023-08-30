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
.PARAMETER ControllerId
No description available.
.PARAMETER ControllerMountPoint
No description available.
.PARAMETER Resizeable
No description available.
.PARAMETER PlanResizable
No description available.
.PARAMETER RootVolume
No description available.
.PARAMETER UnitNumber
No description available.
.PARAMETER TypeId
No description available.
.PARAMETER ConfigurableIOPS
No description available.
.PARAMETER DatastoreId
No description available.
.PARAMETER MaxStorage
No description available.
.PARAMETER DisplayOrder
No description available.
.PARAMETER MaxIOPS
No description available.
.PARAMETER Uuid
No description available.
.OUTPUTS

ClusterMastersVolumesInner<PSCustomObject>
#>

function Initialize-ClusterMastersVolumesInner {
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
        ${ControllerId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ControllerMountPoint},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Resizeable},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${PlanResizable},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${RootVolume},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${UnitNumber},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${TypeId},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${ConfigurableIOPS},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DatastoreId},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxStorage},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${DisplayOrder},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${MaxIOPS},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Uuid}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterMastersVolumesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "controllerId" = ${ControllerId}
            "controllerMountPoint" = ${ControllerMountPoint}
            "resizeable" = ${Resizeable}
            "planResizable" = ${PlanResizable}
            "rootVolume" = ${RootVolume}
            "unitNumber" = ${UnitNumber}
            "typeId" = ${TypeId}
            "configurableIOPS" = ${ConfigurableIOPS}
            "datastoreId" = ${DatastoreId}
            "maxStorage" = ${MaxStorage}
            "displayOrder" = ${DisplayOrder}
            "maxIOPS" = ${MaxIOPS}
            "uuid" = ${Uuid}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterMastersVolumesInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterMastersVolumesInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterMastersVolumesInner<PSCustomObject>
#>
function ConvertFrom-JsonToClusterMastersVolumesInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterMastersVolumesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterMastersVolumesInner
        $AllProperties = ("id", "name", "controllerId", "controllerMountPoint", "resizeable", "planResizable", "rootVolume", "unitNumber", "typeId", "configurableIOPS", "datastoreId", "maxStorage", "displayOrder", "maxIOPS", "uuid")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "controllerId"))) { #optional property not found
            $ControllerId = $null
        } else {
            $ControllerId = $JsonParameters.PSobject.Properties["controllerId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "controllerMountPoint"))) { #optional property not found
            $ControllerMountPoint = $null
        } else {
            $ControllerMountPoint = $JsonParameters.PSobject.Properties["controllerMountPoint"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "resizeable"))) { #optional property not found
            $Resizeable = $null
        } else {
            $Resizeable = $JsonParameters.PSobject.Properties["resizeable"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "planResizable"))) { #optional property not found
            $PlanResizable = $null
        } else {
            $PlanResizable = $JsonParameters.PSobject.Properties["planResizable"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "rootVolume"))) { #optional property not found
            $RootVolume = $null
        } else {
            $RootVolume = $JsonParameters.PSobject.Properties["rootVolume"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "unitNumber"))) { #optional property not found
            $UnitNumber = $null
        } else {
            $UnitNumber = $JsonParameters.PSobject.Properties["unitNumber"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "typeId"))) { #optional property not found
            $TypeId = $null
        } else {
            $TypeId = $JsonParameters.PSobject.Properties["typeId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "configurableIOPS"))) { #optional property not found
            $ConfigurableIOPS = $null
        } else {
            $ConfigurableIOPS = $JsonParameters.PSobject.Properties["configurableIOPS"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "datastoreId"))) { #optional property not found
            $DatastoreId = $null
        } else {
            $DatastoreId = $JsonParameters.PSobject.Properties["datastoreId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxStorage"))) { #optional property not found
            $MaxStorage = $null
        } else {
            $MaxStorage = $JsonParameters.PSobject.Properties["maxStorage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "displayOrder"))) { #optional property not found
            $DisplayOrder = $null
        } else {
            $DisplayOrder = $JsonParameters.PSobject.Properties["displayOrder"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxIOPS"))) { #optional property not found
            $MaxIOPS = $null
        } else {
            $MaxIOPS = $JsonParameters.PSobject.Properties["maxIOPS"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "uuid"))) { #optional property not found
            $Uuid = $null
        } else {
            $Uuid = $JsonParameters.PSobject.Properties["uuid"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "controllerId" = ${ControllerId}
            "controllerMountPoint" = ${ControllerMountPoint}
            "resizeable" = ${Resizeable}
            "planResizable" = ${PlanResizable}
            "rootVolume" = ${RootVolume}
            "unitNumber" = ${UnitNumber}
            "typeId" = ${TypeId}
            "configurableIOPS" = ${ConfigurableIOPS}
            "datastoreId" = ${DatastoreId}
            "maxStorage" = ${MaxStorage}
            "displayOrder" = ${DisplayOrder}
            "maxIOPS" = ${MaxIOPS}
            "uuid" = ${Uuid}
        }

        return $PSO
    }

}

