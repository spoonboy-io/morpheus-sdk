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
.PARAMETER Category
No description available.
.PARAMETER Content
No description available.
.OUTPUTS

ApiWikiPagesPage<PSCustomObject>
#>

function Initialize-ApiWikiPagesPage {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Category},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Content}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiWikiPagesPage' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if (!$Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }

        if (!$Category) {
            throw "invalid value for 'Category', 'Category' cannot be null."
        }

        if (!$Content) {
            throw "invalid value for 'Content', 'Content' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "category" = ${Category}
            "content" = ${Content}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiWikiPagesPage<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiWikiPagesPage<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiWikiPagesPage<PSCustomObject>
#>
function ConvertFrom-JsonToApiWikiPagesPage {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiWikiPagesPage' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiWikiPagesPage
        $AllProperties = ("name", "category", "content")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "category"))) {
            throw "Error! JSON cannot be serialized due to the required property `category` missing."
        } else {
            $Category = $JsonParameters.PSobject.Properties["category"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "content"))) {
            throw "Error! JSON cannot be serialized due to the required property `content` missing."
        } else {
            $Content = $JsonParameters.PSobject.Properties["content"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "category" = ${Category}
            "content" = ${Content}
        }

        return $PSO
    }

}
