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

.PARAMETER ServiceUrl
No description available.
.PARAMETER ServiceHost
No description available.
.PARAMETER ServicePath
No description available.
.PARAMETER ServiceHostname
No description available.
.PARAMETER ServicePort
No description available.
.PARAMETER ServiceUsername
No description available.
.PARAMETER ServicePassword
No description available.
.PARAMETER ServicePasswordHash
No description available.
.PARAMETER ServiceToken
API Token
.PARAMETER ServiceAccess
Kube Config
.PARAMETER ServiceCert
No description available.
.PARAMETER ServiceVersion
No description available.
.OUTPUTS

ClusterApiConfig<PSCustomObject>
#>

function Initialize-ClusterApiConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServiceUrl},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServiceHost},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServicePath},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServiceHostname},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ServicePort},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServiceUsername},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServicePassword},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServicePasswordHash},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServiceToken},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServiceAccess},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServiceCert},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServiceVersion}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterApiConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "serviceUrl" = ${ServiceUrl}
            "serviceHost" = ${ServiceHost}
            "servicePath" = ${ServicePath}
            "serviceHostname" = ${ServiceHostname}
            "servicePort" = ${ServicePort}
            "serviceUsername" = ${ServiceUsername}
            "servicePassword" = ${ServicePassword}
            "servicePasswordHash" = ${ServicePasswordHash}
            "serviceToken" = ${ServiceToken}
            "serviceAccess" = ${ServiceAccess}
            "serviceCert" = ${ServiceCert}
            "serviceVersion" = ${ServiceVersion}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterApiConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterApiConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterApiConfig<PSCustomObject>
#>
function ConvertFrom-JsonToClusterApiConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterApiConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterApiConfig
        $AllProperties = ("serviceUrl", "serviceHost", "servicePath", "serviceHostname", "servicePort", "serviceUsername", "servicePassword", "servicePasswordHash", "serviceToken", "serviceAccess", "serviceCert", "serviceVersion")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serviceUrl"))) { #optional property not found
            $ServiceUrl = $null
        } else {
            $ServiceUrl = $JsonParameters.PSobject.Properties["serviceUrl"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serviceHost"))) { #optional property not found
            $ServiceHost = $null
        } else {
            $ServiceHost = $JsonParameters.PSobject.Properties["serviceHost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "servicePath"))) { #optional property not found
            $ServicePath = $null
        } else {
            $ServicePath = $JsonParameters.PSobject.Properties["servicePath"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serviceHostname"))) { #optional property not found
            $ServiceHostname = $null
        } else {
            $ServiceHostname = $JsonParameters.PSobject.Properties["serviceHostname"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "servicePort"))) { #optional property not found
            $ServicePort = $null
        } else {
            $ServicePort = $JsonParameters.PSobject.Properties["servicePort"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serviceUsername"))) { #optional property not found
            $ServiceUsername = $null
        } else {
            $ServiceUsername = $JsonParameters.PSobject.Properties["serviceUsername"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "servicePassword"))) { #optional property not found
            $ServicePassword = $null
        } else {
            $ServicePassword = $JsonParameters.PSobject.Properties["servicePassword"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "servicePasswordHash"))) { #optional property not found
            $ServicePasswordHash = $null
        } else {
            $ServicePasswordHash = $JsonParameters.PSobject.Properties["servicePasswordHash"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serviceToken"))) { #optional property not found
            $ServiceToken = $null
        } else {
            $ServiceToken = $JsonParameters.PSobject.Properties["serviceToken"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serviceAccess"))) { #optional property not found
            $ServiceAccess = $null
        } else {
            $ServiceAccess = $JsonParameters.PSobject.Properties["serviceAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serviceCert"))) { #optional property not found
            $ServiceCert = $null
        } else {
            $ServiceCert = $JsonParameters.PSobject.Properties["serviceCert"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serviceVersion"))) { #optional property not found
            $ServiceVersion = $null
        } else {
            $ServiceVersion = $JsonParameters.PSobject.Properties["serviceVersion"].value
        }

        $PSO = [PSCustomObject]@{
            "serviceUrl" = ${ServiceUrl}
            "serviceHost" = ${ServiceHost}
            "servicePath" = ${ServicePath}
            "serviceHostname" = ${ServiceHostname}
            "servicePort" = ${ServicePort}
            "serviceUsername" = ${ServiceUsername}
            "servicePassword" = ${ServicePassword}
            "servicePasswordHash" = ${ServicePasswordHash}
            "serviceToken" = ${ServiceToken}
            "serviceAccess" = ${ServiceAccess}
            "serviceCert" = ${ServiceCert}
            "serviceVersion" = ${ServiceVersion}
        }

        return $PSO
    }

}
