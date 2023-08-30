#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'GuidanceVmwareSizing' {
    Context 'GuidanceVmwareSizing' {
        It 'Initialize-GuidanceVmwareSizing' {
            # a simple test to create an object
            #$NewObject = Initialize-GuidanceVmwareSizing -Id "TEST_VALUE" -DateCreated "TEST_VALUE" -LastUpdated "TEST_VALUE" -ActionCategory "TEST_VALUE" -ActionMessage "TEST_VALUE" -ActionTitle "TEST_VALUE" -ActionType "TEST_VALUE" -ActionValue "TEST_VALUE" -ActionValueType "TEST_VALUE" -ActionPlanId "TEST_VALUE" -StatusMessage "TEST_VALUE" -AccountId "TEST_VALUE" -UserId "TEST_VALUE" -SiteId "TEST_VALUE" -Zone "TEST_VALUE" -State "TEST_VALUE" -StateMessage "TEST_VALUE" -Severity "TEST_VALUE" -Resolved "TEST_VALUE" -ResolvedMessage "TEST_VALUE" -RefType "TEST_VALUE" -RefId "TEST_VALUE" -RefName "TEST_VALUE" -Type "TEST_VALUE" -Savings "TEST_VALUE" -Resource "TEST_VALUE" -PlanBeforeAction "TEST_VALUE" -PlanAfterAction "TEST_VALUE" -Config "TEST_VALUE"
            #$NewObject | Should -BeOfType GuidanceVmwareSizing
            #$NewObject.property | Should -Be 0
        }
    }
}
