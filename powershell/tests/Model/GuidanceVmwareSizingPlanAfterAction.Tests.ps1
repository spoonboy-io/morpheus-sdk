#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'GuidanceVmwareSizingPlanAfterAction' {
    Context 'GuidanceVmwareSizingPlanAfterAction' {
        It 'Initialize-GuidanceVmwareSizingPlanAfterAction' {
            # a simple test to create an object
            #$NewObject = Initialize-GuidanceVmwareSizingPlanAfterAction -Id "TEST_VALUE" -Name "TEST_VALUE" -Code "TEST_VALUE" -Active "TEST_VALUE" -SortOrder "TEST_VALUE" -Description "TEST_VALUE" -MaxStorage "TEST_VALUE" -MaxMemory "TEST_VALUE" -MaxCpu "TEST_VALUE" -MaxCores "TEST_VALUE" -MaxDisks "TEST_VALUE" -CoresPerSocket "TEST_VALUE" -CustomCpu "TEST_VALUE" -CustomCores "TEST_VALUE" -CustomMaxStorage "TEST_VALUE" -CustomMaxDataStorage "TEST_VALUE" -CustomMaxMemory "TEST_VALUE" -AddVolumes "TEST_VALUE" -MemoryOptionSource "TEST_VALUE" -CpuOptionSource "TEST_VALUE" -DateCreated "TEST_VALUE" -LastUpdated "TEST_VALUE" -RegionCode "TEST_VALUE" -Visibility "TEST_VALUE" -Editable "TEST_VALUE" -ProvisionType "TEST_VALUE" -Tenants "TEST_VALUE" -PriceSets "TEST_VALUE" -Config "TEST_VALUE"
            #$NewObject | Should -BeOfType GuidanceVmwareSizingPlanAfterAction
            #$NewObject.property | Should -Be 0
        }
    }
}
