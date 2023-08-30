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

.PARAMETER Approval
No description available.
.OUTPUTS

GetApprovals200Response<PSCustomObject>
#>

function Initialize-GetApprovals200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Approval}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetApprovals200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "approval" = ${Approval}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetApprovals200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetApprovals200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetApprovals200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetApprovals200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetApprovals200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetApprovals200Response
        $AllProperties = ("approval")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "approval"))) { #optional property not found
            $Approval = $null
        } else {
            $Approval = $JsonParameters.PSobject.Properties["approval"].value
        }

        $PSO = [PSCustomObject]@{
            "approval" = ${Approval}
        }

        return $PSO
    }

}

