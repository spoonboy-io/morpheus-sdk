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
.PARAMETER Success
No description available.
.PARAMETER ActivityType
No description available.
.PARAMETER Name
No description available.
.PARAMETER Message
No description available.
.PARAMETER ObjectType
No description available.
.PARAMETER ObjectId
No description available.
.PARAMETER User
No description available.
.PARAMETER Ts
No description available.
.OUTPUTS

ActivityActivityInner<PSCustomObject>
#>

function Initialize-ActivityActivityInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Success},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ActivityType},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Message},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ObjectType},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ObjectId},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${User},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${Ts}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ActivityActivityInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "_id" = ${Id}
            "success" = ${Success}
            "activityType" = ${ActivityType}
            "name" = ${Name}
            "message" = ${Message}
            "objectType" = ${ObjectType}
            "objectId" = ${ObjectId}
            "user" = ${User}
            "ts" = ${Ts}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ActivityActivityInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to ActivityActivityInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ActivityActivityInner<PSCustomObject>
#>
function ConvertFrom-JsonToActivityActivityInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ActivityActivityInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ActivityActivityInner
        $AllProperties = ("_id", "success", "activityType", "name", "message", "objectType", "objectId", "user", "ts")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "_id"))) { #optional property not found
            $Id = $null
        } else {
            $Id = $JsonParameters.PSobject.Properties["_id"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "success"))) { #optional property not found
            $Success = $null
        } else {
            $Success = $JsonParameters.PSobject.Properties["success"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "activityType"))) { #optional property not found
            $ActivityType = $null
        } else {
            $ActivityType = $JsonParameters.PSobject.Properties["activityType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "message"))) { #optional property not found
            $Message = $null
        } else {
            $Message = $JsonParameters.PSobject.Properties["message"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "objectType"))) { #optional property not found
            $ObjectType = $null
        } else {
            $ObjectType = $JsonParameters.PSobject.Properties["objectType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "objectId"))) { #optional property not found
            $ObjectId = $null
        } else {
            $ObjectId = $JsonParameters.PSobject.Properties["objectId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "user"))) { #optional property not found
            $User = $null
        } else {
            $User = $JsonParameters.PSobject.Properties["user"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ts"))) { #optional property not found
            $Ts = $null
        } else {
            $Ts = $JsonParameters.PSobject.Properties["ts"].value
        }

        $PSO = [PSCustomObject]@{
            "_id" = ${Id}
            "success" = ${Success}
            "activityType" = ${ActivityType}
            "name" = ${Name}
            "message" = ${Message}
            "objectType" = ${ObjectType}
            "objectId" = ${ObjectId}
            "user" = ${User}
            "ts" = ${Ts}
        }

        return $PSO
    }

}

