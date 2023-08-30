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

.PARAMETER Account
No description available.
.OUTPUTS

AddTenantRequest<PSCustomObject>
#>

function Initialize-AddTenantRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Account}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddTenantRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Account) {
            throw "invalid value for 'Account', 'Account' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "account" = ${Account}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddTenantRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddTenantRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddTenantRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddTenantRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddTenantRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddTenantRequest
        $AllProperties = ("account")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'account' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "account"))) {
            throw "Error! JSON cannot be serialized due to the required property 'account' missing."
        } else {
            $Account = $JsonParameters.PSobject.Properties["account"].value
        }

        $PSO = [PSCustomObject]@{
            "account" = ${Account}
        }

        return $PSO
    }

}
