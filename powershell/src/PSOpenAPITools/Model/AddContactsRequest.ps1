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

.PARAMETER Contact
No description available.
.OUTPUTS

AddContactsRequest<PSCustomObject>
#>

function Initialize-AddContactsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Contact}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddContactsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Contact) {
            throw "invalid value for 'Contact', 'Contact' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "contact" = ${Contact}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddContactsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddContactsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddContactsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddContactsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddContactsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddContactsRequest
        $AllProperties = ("contact")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'contact' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "contact"))) {
            throw "Error! JSON cannot be serialized due to the required property 'contact' missing."
        } else {
            $Contact = $JsonParameters.PSobject.Properties["contact"].value
        }

        $PSO = [PSCustomObject]@{
            "contact" = ${Contact}
        }

        return $PSO
    }

}

