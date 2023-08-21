#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'BillingInstanceUsages' {
    Context 'BillingInstanceUsages' {
        It 'Initialize-BillingInstanceUsages' {
            # a simple test to create an object
            #$NewObject = Initialize-BillingInstanceUsages -Name "TEST_VALUE" -InstanceName "TEST_VALUE" -ZoneName "TEST_VALUE" -AccountName "TEST_VALUE" -Volumes "TEST_VALUE" -MaxMemory "TEST_VALUE" -MaxCpu "TEST_VALUE" -MaxCores "TEST_VALUE" -ServerExternalId "TEST_VALUE" -ServerInternalId "TEST_VALUE" -PlanName "TEST_VALUE" -HourlyPrice "TEST_VALUE" -HourlyCost "TEST_VALUE" -Currency "TEST_VALUE" -PricesUsed "TEST_VALUE" -Cost "TEST_VALUE" -Price "TEST_VALUE" -CreatedByUser "TEST_VALUE" -CreatedByUserId "TEST_VALUE" -SiteId "TEST_VALUE" -SiteName "TEST_VALUE" -SiteUUID "TEST_VALUE" -SiteCode "TEST_VALUE" -StartDate "TEST_VALUE" -EndDate "TEST_VALUE" -Status "TEST_VALUE" -Tags "TEST_VALUE" -ApplicablePrices "TEST_VALUE" -ServicePlanId "TEST_VALUE" -ServicePlanName "TEST_VALUE" -ResourcePoolId "TEST_VALUE" -ResourcePoolName "TEST_VALUE"
            #$NewObject | Should -BeOfType BillingInstanceUsages
            #$NewObject.property | Should -Be 0
        }
    }
}
