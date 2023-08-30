#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'ClusterLayoutComputeServersInner' {
    Context 'ClusterLayoutComputeServersInner' {
        It 'Initialize-ClusterLayoutComputeServersInner' {
            # a simple test to create an object
            #$NewObject = Initialize-ClusterLayoutComputeServersInner -Id "TEST_VALUE" -PriorityOrder "TEST_VALUE" -NodeCount "TEST_VALUE" -NodeType "TEST_VALUE" -MinNodeCount "TEST_VALUE" -MaxNodeCount "TEST_VALUE" -DynamicCount "TEST_VALUE" -InstallContainerRuntime "TEST_VALUE" -InstallStorageRuntime "TEST_VALUE" -Name "TEST_VALUE" -Code "TEST_VALUE" -Category "TEST_VALUE" -Config "TEST_VALUE" -ContainerType "TEST_VALUE" -ComputeServerType "TEST_VALUE" -ProvisionService "TEST_VALUE" -PlanCategory "TEST_VALUE" -NamePrefix "TEST_VALUE" -NameSuffix "TEST_VALUE" -ForceNameIndex "TEST_VALUE" -LoadBalance "TEST_VALUE"
            #$NewObject | Should -BeOfType ClusterLayoutComputeServersInner
            #$NewObject.property | Should -Be 0
        }
    }
}
