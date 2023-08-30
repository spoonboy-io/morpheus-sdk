#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'Budgets' {
    Context 'Budgets' {
        It 'Initialize-Budgets' {
            # a simple test to create an object
            #$NewObject = Initialize-Budgets -Id "TEST_VALUE" -Name "TEST_VALUE" -Description "TEST_VALUE" -Account "TEST_VALUE" -Enabled "TEST_VALUE" -RefScope "TEST_VALUE" -RefType "TEST_VALUE" -RefId "TEST_VALUE" -RefName "TEST_VALUE" -Period "TEST_VALUE" -Year "TEST_VALUE" -ResourceType "TEST_VALUE" -Timezone "TEST_VALUE" -StartDate "TEST_VALUE" -EndDate "TEST_VALUE" -Interval "TEST_VALUE" -Costs "TEST_VALUE" -IsFiscal "TEST_VALUE" -AverageCost "TEST_VALUE" -TotalCost "TEST_VALUE" -Currency "TEST_VALUE" -Rollover "TEST_VALUE" -WarningLimit "TEST_VALUE" -OverLimit "TEST_VALUE" -ExternalId "TEST_VALUE" -InternalId "TEST_VALUE" -CreatedById "TEST_VALUE" -CreatedByName "TEST_VALUE" -UpdatedById "TEST_VALUE" -UpdatedByName "TEST_VALUE" -DateCreated "TEST_VALUE" -LastUpdated "TEST_VALUE"
            #$NewObject | Should -BeOfType Budgets
            #$NewObject.property | Should -Be 0
        }
    }
}
