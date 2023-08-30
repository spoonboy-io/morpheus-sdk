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

.PARAMETER Username
Specified as ""username"" or ""tenantId\username"" with proper HTML encoding and used in conjunction with `password`. 
.OUTPUTS

ForgotPasswordRequest<PSCustomObject>
#>

function Initialize-ForgotPasswordRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Username}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ForgotPasswordRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Username) {
            throw "invalid value for 'Username', 'Username' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "username" = ${Username}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ForgotPasswordRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to ForgotPasswordRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ForgotPasswordRequest<PSCustomObject>
#>
function ConvertFrom-JsonToForgotPasswordRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ForgotPasswordRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ForgotPasswordRequest
        $AllProperties = ("username")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'username' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "username"))) {
            throw "Error! JSON cannot be serialized due to the required property 'username' missing."
        } else {
            $Username = $JsonParameters.PSobject.Properties["username"].value
        }

        $PSO = [PSCustomObject]@{
            "username" = ${Username}
        }

        return $PSO
    }

}
