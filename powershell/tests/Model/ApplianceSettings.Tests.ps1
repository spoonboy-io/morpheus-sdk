#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.1.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'ApplianceSettings' {
    Context 'ApplianceSettings' {
        It 'Initialize-ApplianceSettings' {
            # a simple test to create an object
            #$NewObject = Initialize-ApplianceSettings -ApplianceUrl "TEST_VALUE" -InternalApplianceUrl "TEST_VALUE" -CorsAllowed "TEST_VALUE" -RegistrationEnabled "TEST_VALUE" -DefaultRoleId "TEST_VALUE" -DefaultUserRoleId "TEST_VALUE" -DockerPrivilegedMode "TEST_VALUE" -ExpirePwdDays "TEST_VALUE" -DisableAfterAttempts "TEST_VALUE" -DisableAfterDaysInactive "TEST_VALUE" -WarnUserDaysBefore "TEST_VALUE" -SmtpMailFrom "TEST_VALUE" -SmtpServer "TEST_VALUE" -SmtpPort "TEST_VALUE" -SmtpSSL "TEST_VALUE" -SmtpTLS "TEST_VALUE" -SmtpUser "TEST_VALUE" -SmtpPassword "TEST_VALUE" -SmtpPasswordHash "TEST_VALUE" -ProxyHost "TEST_VALUE" -ProxyPort "TEST_VALUE" -ProxyUser "TEST_VALUE" -ProxyPassword "TEST_VALUE" -ProxyPasswordHash "TEST_VALUE" -ProxyDomain "TEST_VALUE" -ProxyWorkstation "TEST_VALUE" -CurrencyProvider "TEST_VALUE" -CurrencyKey "TEST_VALUE" -EnabledZoneTypes "TEST_VALUE" -StatsRetainmentPeriod "TEST_VALUE"
            #$NewObject | Should -BeOfType ApplianceSettings
            #$NewObject.property | Should -Be 0
        }
    }
}
