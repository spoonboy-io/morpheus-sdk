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
The id for the LV configuration being created
.PARAMETER RootVolume
If set to false then a non-root LV will be created
.PARAMETER Name
Name/type of the LV being created
.PARAMETER Size
Size of the LV to be created in GBs  Default is from the service plan 
.PARAMETER SizeId
Can be used to select pre-existing LV choices from Morpheus
.PARAMETER StorageType
Identifier for LV type
.PARAMETER DatastoreId
The ID of the specific datastore. Auto selection can be specified as auto or `autoCluster` (for clusters).
.OUTPUTS

ClusterServerCreateVolumesInner<PSCustomObject>
#>

function Initialize-ClusterServerCreateVolumesInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id} = -1,
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${RootVolume} = $true,
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name} = "root",
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Size},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SizeId},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${StorageType},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DatastoreId}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterServerCreateVolumesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "rootVolume" = ${RootVolume}
            "name" = ${Name}
            "size" = ${Size}
            "sizeId" = ${SizeId}
            "storageType" = ${StorageType}
            "datastoreId" = ${DatastoreId}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterServerCreateVolumesInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterServerCreateVolumesInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterServerCreateVolumesInner<PSCustomObject>
#>
function ConvertFrom-JsonToClusterServerCreateVolumesInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterServerCreateVolumesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterServerCreateVolumesInner
        $AllProperties = ("id", "rootVolume", "name", "size", "sizeId", "storageType", "datastoreId")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "datastoreId"))) {
            throw "Error! JSON cannot be serialized due to the required property 'datastoreId' missing."
        } else {
            $DatastoreId = $JsonParameters.PSobject.Properties["datastoreId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "id"))) { #optional property not found
            $Id = $null
        } else {
            $Id = $JsonParameters.PSobject.Properties["id"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "rootVolume"))) { #optional property not found
            $RootVolume = $null
        } else {
            $RootVolume = $JsonParameters.PSobject.Properties["rootVolume"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "size"))) { #optional property not found
            $Size = $null
        } else {
            $Size = $JsonParameters.PSobject.Properties["size"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sizeId"))) { #optional property not found
            $SizeId = $null
        } else {
            $SizeId = $JsonParameters.PSobject.Properties["sizeId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "storageType"))) { #optional property not found
            $StorageType = $null
        } else {
            $StorageType = $JsonParameters.PSobject.Properties["storageType"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "rootVolume" = ${RootVolume}
            "name" = ${Name}
            "size" = ${Size}
            "sizeId" = ${SizeId}
            "storageType" = ${StorageType}
            "datastoreId" = ${DatastoreId}
        }

        return $PSO
    }

}
