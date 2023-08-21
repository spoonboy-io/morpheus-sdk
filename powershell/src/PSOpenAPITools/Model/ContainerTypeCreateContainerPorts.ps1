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

.PARAMETER Name
No description available.
.PARAMETER Port
No description available.
.PARAMETER LoadBalanceProtocol
No description available.
.OUTPUTS

ContainerTypeCreateContainerPorts<PSCustomObject>
#>

function Initialize-ContainerTypeCreateContainerPorts {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [Int64]
        ${Port},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LoadBalanceProtocol}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ContainerTypeCreateContainerPorts' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if (!$Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }

        if (!$Port) {
            throw "invalid value for 'Port', 'Port' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "port" = ${Port}
            "loadBalanceProtocol" = ${LoadBalanceProtocol}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ContainerTypeCreateContainerPorts<PSCustomObject>

.DESCRIPTION

Convert from JSON to ContainerTypeCreateContainerPorts<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ContainerTypeCreateContainerPorts<PSCustomObject>
#>
function ConvertFrom-JsonToContainerTypeCreateContainerPorts {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ContainerTypeCreateContainerPorts' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ContainerTypeCreateContainerPorts
        $AllProperties = ("name", "port", "loadBalanceProtocol")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property `name` missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property `name` missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "port"))) {
            throw "Error! JSON cannot be serialized due to the required property `port` missing."
        } else {
            $Port = $JsonParameters.PSobject.Properties["port"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "loadBalanceProtocol"))) { #optional property not found
            $LoadBalanceProtocol = $null
        } else {
            $LoadBalanceProtocol = $JsonParameters.PSobject.Properties["loadBalanceProtocol"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "port" = ${Port}
            "loadBalanceProtocol" = ${LoadBalanceProtocol}
        }

        return $PSO
    }

}

