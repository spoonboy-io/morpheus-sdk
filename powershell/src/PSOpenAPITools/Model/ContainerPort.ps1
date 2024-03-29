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
.PARAMETER Index
No description available.
.PARAMETER External
No description available.
.PARAMETER Internal
No description available.
.PARAMETER DisplayName
No description available.
.PARAMETER PrimaryPort
No description available.
.PARAMETER Export
No description available.
.PARAMETER Visible
No description available.
.PARAMETER ExportName
No description available.
.PARAMETER LoadBalanceProtocol
No description available.
.PARAMETER LoadBalance
No description available.
.PARAMETER Protocol
No description available.
.PARAMETER Link
No description available.
.PARAMETER ExternalIp
No description available.
.PARAMETER InternalIp
No description available.
.OUTPUTS

ContainerPort<PSCustomObject>
#>

function Initialize-ContainerPort {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Index},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${External},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Internal},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DisplayName},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${PrimaryPort},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Export},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Visible},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExportName},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LoadBalanceProtocol},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${LoadBalance},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Protocol},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Link},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalIp},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InternalIp}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ContainerPort' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "index" = ${Index}
            "external" = ${External}
            "internal" = ${Internal}
            "displayName" = ${DisplayName}
            "primaryPort" = ${PrimaryPort}
            "export" = ${Export}
            "visible" = ${Visible}
            "exportName" = ${ExportName}
            "loadBalanceProtocol" = ${LoadBalanceProtocol}
            "loadBalance" = ${LoadBalance}
            "protocol" = ${Protocol}
            "link" = ${Link}
            "externalIp" = ${ExternalIp}
            "internalIp" = ${InternalIp}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ContainerPort<PSCustomObject>

.DESCRIPTION

Convert from JSON to ContainerPort<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ContainerPort<PSCustomObject>
#>
function ConvertFrom-JsonToContainerPort {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ContainerPort' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ContainerPort
        $AllProperties = ("id", "index", "external", "internal", "displayName", "primaryPort", "export", "visible", "exportName", "loadBalanceProtocol", "loadBalance", "protocol", "link", "externalIp", "internalIp")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "index"))) { #optional property not found
            $Index = $null
        } else {
            $Index = $JsonParameters.PSobject.Properties["index"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "external"))) { #optional property not found
            $External = $null
        } else {
            $External = $JsonParameters.PSobject.Properties["external"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "internal"))) { #optional property not found
            $Internal = $null
        } else {
            $Internal = $JsonParameters.PSobject.Properties["internal"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "displayName"))) { #optional property not found
            $DisplayName = $null
        } else {
            $DisplayName = $JsonParameters.PSobject.Properties["displayName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "primaryPort"))) { #optional property not found
            $PrimaryPort = $null
        } else {
            $PrimaryPort = $JsonParameters.PSobject.Properties["primaryPort"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "export"))) { #optional property not found
            $Export = $null
        } else {
            $Export = $JsonParameters.PSobject.Properties["export"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visible"))) { #optional property not found
            $Visible = $null
        } else {
            $Visible = $JsonParameters.PSobject.Properties["visible"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "exportName"))) { #optional property not found
            $ExportName = $null
        } else {
            $ExportName = $JsonParameters.PSobject.Properties["exportName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "loadBalanceProtocol"))) { #optional property not found
            $LoadBalanceProtocol = $null
        } else {
            $LoadBalanceProtocol = $JsonParameters.PSobject.Properties["loadBalanceProtocol"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "loadBalance"))) { #optional property not found
            $LoadBalance = $null
        } else {
            $LoadBalance = $JsonParameters.PSobject.Properties["loadBalance"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "protocol"))) { #optional property not found
            $Protocol = $null
        } else {
            $Protocol = $JsonParameters.PSobject.Properties["protocol"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "link"))) { #optional property not found
            $Link = $null
        } else {
            $Link = $JsonParameters.PSobject.Properties["link"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalIp"))) { #optional property not found
            $ExternalIp = $null
        } else {
            $ExternalIp = $JsonParameters.PSobject.Properties["externalIp"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "internalIp"))) { #optional property not found
            $InternalIp = $null
        } else {
            $InternalIp = $JsonParameters.PSobject.Properties["internalIp"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "index" = ${Index}
            "external" = ${External}
            "internal" = ${Internal}
            "displayName" = ${DisplayName}
            "primaryPort" = ${PrimaryPort}
            "export" = ${Export}
            "visible" = ${Visible}
            "exportName" = ${ExportName}
            "loadBalanceProtocol" = ${LoadBalanceProtocol}
            "loadBalance" = ${LoadBalance}
            "protocol" = ${Protocol}
            "link" = ${Link}
            "externalIp" = ${ExternalIp}
            "internalIp" = ${InternalIp}
        }

        return $PSO
    }

}

