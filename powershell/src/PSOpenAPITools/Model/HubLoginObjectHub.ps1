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

.PARAMETER Email
Email for existing Morpheus Hub user
.PARAMETER Password
Password for existing Morpheus Hub user
.OUTPUTS

HubLoginObjectHub<PSCustomObject>
#>

function Initialize-HubLoginObjectHub {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Email},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Password}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => HubLoginObjectHub' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Email) {
            throw "invalid value for 'Email', 'Email' cannot be null."
        }

        if ($null -eq $Password) {
            throw "invalid value for 'Password', 'Password' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "email" = ${Email}
            "password" = ${Password}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to HubLoginObjectHub<PSCustomObject>

.DESCRIPTION

Convert from JSON to HubLoginObjectHub<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

HubLoginObjectHub<PSCustomObject>
#>
function ConvertFrom-JsonToHubLoginObjectHub {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => HubLoginObjectHub' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in HubLoginObjectHub
        $AllProperties = ("email", "password")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'email' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "email"))) {
            throw "Error! JSON cannot be serialized due to the required property 'email' missing."
        } else {
            $Email = $JsonParameters.PSobject.Properties["email"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "password"))) {
            throw "Error! JSON cannot be serialized due to the required property 'password' missing."
        } else {
            $Password = $JsonParameters.PSobject.Properties["password"].value
        }

        $PSO = [PSCustomObject]@{
            "email" = ${Email}
            "password" = ${Password}
        }

        return $PSO
    }

}

