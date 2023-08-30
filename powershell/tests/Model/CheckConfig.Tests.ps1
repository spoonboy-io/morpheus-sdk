#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'CheckConfig' {
    Context 'CheckConfig' {
        It 'Initialize-CheckConfig' {
            # a simple test to create an object
            #$NewObject = Initialize-CheckConfig -WebMethod "TEST_VALUE" -WebUrl "TEST_VALUE" -IgnoreSSL "TEST_VALUE" -CheckUser "TEST_VALUE" -CheckPassword "TEST_VALUE" -TextCheckOn "TEST_VALUE" -WebTextMatch "TEST_VALUE" -TunnelOn "TEST_VALUE" -SshHost "TEST_VALUE" -SshPort "TEST_VALUE" -SshUser "TEST_VALUE" -SshPassword "TEST_VALUE" -DbHost "TEST_VALUE" -DbPort "TEST_VALUE" -DbUser "TEST_VALUE" -DbPassword "TEST_VALUE" -DbPasswordHash "TEST_VALUE" -DbName "TEST_VALUE" -DbQuery "TEST_VALUE" -CheckOperator "TEST_VALUE" -CheckResult "TEST_VALUE" -CheckPasswordHash "TEST_VALUE" -EsHost "TEST_VALUE" -EsPort "TEST_VALUE" -VarHost "TEST_VALUE" -Port "TEST_VALUE" -Send "TEST_VALUE" -ResponseMatch "TEST_VALUE" -ContainerName "TEST_VALUE" -ExternalId "TEST_VALUE"
            #$NewObject | Should -BeOfType CheckConfig
            #$NewObject.property | Should -Be 0
        }
    }
}