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

.PARAMETER RoleAttributeName
No description available.
.PARAMETER RequiredAttributeValue
No description available.
.PARAMETER GivenNameAttribute
No description available.
.PARAMETER SurnameAttribute
No description available.
.PARAMETER LogoutUrl
No description available.
.PARAMETER DoNotIncludeSAMLRequest
No description available.
.PARAMETER PublicKey
No description available.
.PARAMETER EmailAttribute
No description available.
.PARAMETER Url
No description available.
.PARAMETER DoNotValidateSignature
No description available.
.PARAMETER DoNotValidateStatusCode
No description available.
.PARAMETER DoNotValidateDestination
No description available.
.PARAMETER DoNotValidateIssueInstants
No description available.
.PARAMETER DoNotValidateAssertions
No description available.
.PARAMETER DoNotValidateAuthStatements
No description available.
.PARAMETER DoNotValidateSubject
No description available.
.PARAMETER DoNotValidateConditions
No description available.
.PARAMETER DoNotValidateAudiences
No description available.
.PARAMETER DoNotValidateSubjectRecipients
No description available.
.PARAMETER SAMLSignatureMode
No description available.
.PARAMETER ForceAuthn
No description available.
.OUTPUTS

IdentitySourcesSAMLConfigConfig<PSCustomObject>
#>

function Initialize-IdentitySourcesSAMLConfigConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RoleAttributeName},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RequiredAttributeValue},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GivenNameAttribute},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SurnameAttribute},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LogoutUrl},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotIncludeSAMLRequest},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PublicKey},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${EmailAttribute},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Url},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotValidateSignature},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotValidateStatusCode},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotValidateDestination},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotValidateIssueInstants},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotValidateAssertions},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotValidateAuthStatements},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotValidateSubject},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotValidateConditions},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotValidateAudiences},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${DoNotValidateSubjectRecipients},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SAMLSignatureMode},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${ForceAuthn}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => IdentitySourcesSAMLConfigConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "roleAttributeName" = ${RoleAttributeName}
            "requiredAttributeValue" = ${RequiredAttributeValue}
            "givenNameAttribute" = ${GivenNameAttribute}
            "surnameAttribute" = ${SurnameAttribute}
            "logoutUrl" = ${LogoutUrl}
            "doNotIncludeSAMLRequest" = ${DoNotIncludeSAMLRequest}
            "publicKey" = ${PublicKey}
            "emailAttribute" = ${EmailAttribute}
            "url" = ${Url}
            "doNotValidateSignature" = ${DoNotValidateSignature}
            "doNotValidateStatusCode" = ${DoNotValidateStatusCode}
            "doNotValidateDestination" = ${DoNotValidateDestination}
            "doNotValidateIssueInstants" = ${DoNotValidateIssueInstants}
            "doNotValidateAssertions" = ${DoNotValidateAssertions}
            "doNotValidateAuthStatements" = ${DoNotValidateAuthStatements}
            "doNotValidateSubject" = ${DoNotValidateSubject}
            "doNotValidateConditions" = ${DoNotValidateConditions}
            "doNotValidateAudiences" = ${DoNotValidateAudiences}
            "doNotValidateSubjectRecipients" = ${DoNotValidateSubjectRecipients}
            "SAMLSignatureMode" = ${SAMLSignatureMode}
            "forceAuthn" = ${ForceAuthn}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to IdentitySourcesSAMLConfigConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to IdentitySourcesSAMLConfigConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

IdentitySourcesSAMLConfigConfig<PSCustomObject>
#>
function ConvertFrom-JsonToIdentitySourcesSAMLConfigConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => IdentitySourcesSAMLConfigConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in IdentitySourcesSAMLConfigConfig
        $AllProperties = ("roleAttributeName", "requiredAttributeValue", "givenNameAttribute", "surnameAttribute", "logoutUrl", "doNotIncludeSAMLRequest", "publicKey", "emailAttribute", "url", "doNotValidateSignature", "doNotValidateStatusCode", "doNotValidateDestination", "doNotValidateIssueInstants", "doNotValidateAssertions", "doNotValidateAuthStatements", "doNotValidateSubject", "doNotValidateConditions", "doNotValidateAudiences", "doNotValidateSubjectRecipients", "SAMLSignatureMode", "forceAuthn")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "roleAttributeName"))) { #optional property not found
            $RoleAttributeName = $null
        } else {
            $RoleAttributeName = $JsonParameters.PSobject.Properties["roleAttributeName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "requiredAttributeValue"))) { #optional property not found
            $RequiredAttributeValue = $null
        } else {
            $RequiredAttributeValue = $JsonParameters.PSobject.Properties["requiredAttributeValue"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "givenNameAttribute"))) { #optional property not found
            $GivenNameAttribute = $null
        } else {
            $GivenNameAttribute = $JsonParameters.PSobject.Properties["givenNameAttribute"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "surnameAttribute"))) { #optional property not found
            $SurnameAttribute = $null
        } else {
            $SurnameAttribute = $JsonParameters.PSobject.Properties["surnameAttribute"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "logoutUrl"))) { #optional property not found
            $LogoutUrl = $null
        } else {
            $LogoutUrl = $JsonParameters.PSobject.Properties["logoutUrl"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotIncludeSAMLRequest"))) { #optional property not found
            $DoNotIncludeSAMLRequest = $null
        } else {
            $DoNotIncludeSAMLRequest = $JsonParameters.PSobject.Properties["doNotIncludeSAMLRequest"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "publicKey"))) { #optional property not found
            $PublicKey = $null
        } else {
            $PublicKey = $JsonParameters.PSobject.Properties["publicKey"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "emailAttribute"))) { #optional property not found
            $EmailAttribute = $null
        } else {
            $EmailAttribute = $JsonParameters.PSobject.Properties["emailAttribute"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "url"))) { #optional property not found
            $Url = $null
        } else {
            $Url = $JsonParameters.PSobject.Properties["url"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotValidateSignature"))) { #optional property not found
            $DoNotValidateSignature = $null
        } else {
            $DoNotValidateSignature = $JsonParameters.PSobject.Properties["doNotValidateSignature"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotValidateStatusCode"))) { #optional property not found
            $DoNotValidateStatusCode = $null
        } else {
            $DoNotValidateStatusCode = $JsonParameters.PSobject.Properties["doNotValidateStatusCode"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotValidateDestination"))) { #optional property not found
            $DoNotValidateDestination = $null
        } else {
            $DoNotValidateDestination = $JsonParameters.PSobject.Properties["doNotValidateDestination"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotValidateIssueInstants"))) { #optional property not found
            $DoNotValidateIssueInstants = $null
        } else {
            $DoNotValidateIssueInstants = $JsonParameters.PSobject.Properties["doNotValidateIssueInstants"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotValidateAssertions"))) { #optional property not found
            $DoNotValidateAssertions = $null
        } else {
            $DoNotValidateAssertions = $JsonParameters.PSobject.Properties["doNotValidateAssertions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotValidateAuthStatements"))) { #optional property not found
            $DoNotValidateAuthStatements = $null
        } else {
            $DoNotValidateAuthStatements = $JsonParameters.PSobject.Properties["doNotValidateAuthStatements"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotValidateSubject"))) { #optional property not found
            $DoNotValidateSubject = $null
        } else {
            $DoNotValidateSubject = $JsonParameters.PSobject.Properties["doNotValidateSubject"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotValidateConditions"))) { #optional property not found
            $DoNotValidateConditions = $null
        } else {
            $DoNotValidateConditions = $JsonParameters.PSobject.Properties["doNotValidateConditions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotValidateAudiences"))) { #optional property not found
            $DoNotValidateAudiences = $null
        } else {
            $DoNotValidateAudiences = $JsonParameters.PSobject.Properties["doNotValidateAudiences"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "doNotValidateSubjectRecipients"))) { #optional property not found
            $DoNotValidateSubjectRecipients = $null
        } else {
            $DoNotValidateSubjectRecipients = $JsonParameters.PSobject.Properties["doNotValidateSubjectRecipients"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "SAMLSignatureMode"))) { #optional property not found
            $SAMLSignatureMode = $null
        } else {
            $SAMLSignatureMode = $JsonParameters.PSobject.Properties["SAMLSignatureMode"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "forceAuthn"))) { #optional property not found
            $ForceAuthn = $null
        } else {
            $ForceAuthn = $JsonParameters.PSobject.Properties["forceAuthn"].value
        }

        $PSO = [PSCustomObject]@{
            "roleAttributeName" = ${RoleAttributeName}
            "requiredAttributeValue" = ${RequiredAttributeValue}
            "givenNameAttribute" = ${GivenNameAttribute}
            "surnameAttribute" = ${SurnameAttribute}
            "logoutUrl" = ${LogoutUrl}
            "doNotIncludeSAMLRequest" = ${DoNotIncludeSAMLRequest}
            "publicKey" = ${PublicKey}
            "emailAttribute" = ${EmailAttribute}
            "url" = ${Url}
            "doNotValidateSignature" = ${DoNotValidateSignature}
            "doNotValidateStatusCode" = ${DoNotValidateStatusCode}
            "doNotValidateDestination" = ${DoNotValidateDestination}
            "doNotValidateIssueInstants" = ${DoNotValidateIssueInstants}
            "doNotValidateAssertions" = ${DoNotValidateAssertions}
            "doNotValidateAuthStatements" = ${DoNotValidateAuthStatements}
            "doNotValidateSubject" = ${DoNotValidateSubject}
            "doNotValidateConditions" = ${DoNotValidateConditions}
            "doNotValidateAudiences" = ${DoNotValidateAudiences}
            "doNotValidateSubjectRecipients" = ${DoNotValidateSubjectRecipients}
            "SAMLSignatureMode" = ${SAMLSignatureMode}
            "forceAuthn" = ${ForceAuthn}
        }

        return $PSO
    }

}

