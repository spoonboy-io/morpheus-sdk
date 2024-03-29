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
.PARAMETER Code
No description available.
.PARAMETER Description
No description available.
.PARAMETER Priority
No description available.
.PARAMETER RouteType
No description available.
.PARAMETER Source
No description available.
.PARAMETER SourceType
No description available.
.PARAMETER Destination
No description available.
.PARAMETER DestinationType
No description available.
.PARAMETER DefaultRoute
No description available.
.PARAMETER NetworkMtu
No description available.
.PARAMETER ExternalInterface
No description available.
.PARAMETER InternalId
No description available.
.PARAMETER UniqueId
No description available.
.PARAMETER ExternalType
No description available.
.PARAMETER Enabled
No description available.
.PARAMETER Visible
No description available.
.OUTPUTS

GetStaticRoutes200ResponseNetworkRoutesInner<PSCustomObject>
#>

function Initialize-GetStaticRoutes200ResponseNetworkRoutesInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Code},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Priority},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RouteType},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Source},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SourceType},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Destination},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DestinationType},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DefaultRoute},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${NetworkMtu},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalInterface},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InternalId},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${UniqueId},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalType},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Visible}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetStaticRoutes200ResponseNetworkRoutesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "code" = ${Code}
            "description" = ${Description}
            "priority" = ${Priority}
            "routeType" = ${RouteType}
            "source" = ${Source}
            "sourceType" = ${SourceType}
            "destination" = ${Destination}
            "destinationType" = ${DestinationType}
            "defaultRoute" = ${DefaultRoute}
            "networkMtu" = ${NetworkMtu}
            "externalInterface" = ${ExternalInterface}
            "internalId" = ${InternalId}
            "uniqueId" = ${UniqueId}
            "externalType" = ${ExternalType}
            "enabled" = ${Enabled}
            "visible" = ${Visible}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetStaticRoutes200ResponseNetworkRoutesInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetStaticRoutes200ResponseNetworkRoutesInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetStaticRoutes200ResponseNetworkRoutesInner<PSCustomObject>
#>
function ConvertFrom-JsonToGetStaticRoutes200ResponseNetworkRoutesInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetStaticRoutes200ResponseNetworkRoutesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetStaticRoutes200ResponseNetworkRoutesInner
        $AllProperties = ("id", "name", "code", "description", "priority", "routeType", "source", "sourceType", "destination", "destinationType", "defaultRoute", "networkMtu", "externalInterface", "internalId", "uniqueId", "externalType", "enabled", "visible")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) { #optional property not found
            $Code = $null
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "priority"))) { #optional property not found
            $Priority = $null
        } else {
            $Priority = $JsonParameters.PSobject.Properties["priority"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "routeType"))) { #optional property not found
            $RouteType = $null
        } else {
            $RouteType = $JsonParameters.PSobject.Properties["routeType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "source"))) { #optional property not found
            $Source = $null
        } else {
            $Source = $JsonParameters.PSobject.Properties["source"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sourceType"))) { #optional property not found
            $SourceType = $null
        } else {
            $SourceType = $JsonParameters.PSobject.Properties["sourceType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "destination"))) { #optional property not found
            $Destination = $null
        } else {
            $Destination = $JsonParameters.PSobject.Properties["destination"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "destinationType"))) { #optional property not found
            $DestinationType = $null
        } else {
            $DestinationType = $JsonParameters.PSobject.Properties["destinationType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "defaultRoute"))) { #optional property not found
            $DefaultRoute = $null
        } else {
            $DefaultRoute = $JsonParameters.PSobject.Properties["defaultRoute"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkMtu"))) { #optional property not found
            $NetworkMtu = $null
        } else {
            $NetworkMtu = $JsonParameters.PSobject.Properties["networkMtu"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalInterface"))) { #optional property not found
            $ExternalInterface = $null
        } else {
            $ExternalInterface = $JsonParameters.PSobject.Properties["externalInterface"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "internalId"))) { #optional property not found
            $InternalId = $null
        } else {
            $InternalId = $JsonParameters.PSobject.Properties["internalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "uniqueId"))) { #optional property not found
            $UniqueId = $null
        } else {
            $UniqueId = $JsonParameters.PSobject.Properties["uniqueId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalType"))) { #optional property not found
            $ExternalType = $null
        } else {
            $ExternalType = $JsonParameters.PSobject.Properties["externalType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visible"))) { #optional property not found
            $Visible = $null
        } else {
            $Visible = $JsonParameters.PSobject.Properties["visible"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "code" = ${Code}
            "description" = ${Description}
            "priority" = ${Priority}
            "routeType" = ${RouteType}
            "source" = ${Source}
            "sourceType" = ${SourceType}
            "destination" = ${Destination}
            "destinationType" = ${DestinationType}
            "defaultRoute" = ${DefaultRoute}
            "networkMtu" = ${NetworkMtu}
            "externalInterface" = ${ExternalInterface}
            "internalId" = ${InternalId}
            "uniqueId" = ${UniqueId}
            "externalType" = ${ExternalType}
            "enabled" = ${Enabled}
            "visible" = ${Visible}
        }

        return $PSO
    }

}

