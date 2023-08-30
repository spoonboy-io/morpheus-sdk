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

.PARAMETER VdiGateway
No description available.
.OUTPUTS

AddVDIGatewaysRequest<PSCustomObject>
#>

function Initialize-AddVDIGatewaysRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${VdiGateway}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddVDIGatewaysRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $VdiGateway) {
            throw "invalid value for 'VdiGateway', 'VdiGateway' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "vdiGateway" = ${VdiGateway}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddVDIGatewaysRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddVDIGatewaysRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddVDIGatewaysRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddVDIGatewaysRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddVDIGatewaysRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddVDIGatewaysRequest
        $AllProperties = ("vdiGateway")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'vdiGateway' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "vdiGateway"))) {
            throw "Error! JSON cannot be serialized due to the required property 'vdiGateway' missing."
        } else {
            $VdiGateway = $JsonParameters.PSobject.Properties["vdiGateway"].value
        }

        $PSO = [PSCustomObject]@{
            "vdiGateway" = ${VdiGateway}
        }

        return $PSO
    }

}

