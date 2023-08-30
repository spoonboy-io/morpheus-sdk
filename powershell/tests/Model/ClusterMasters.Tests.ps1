#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'ClusterMasters' {
    Context 'ClusterMasters' {
        It 'Initialize-ClusterMasters' {
            # a simple test to create an object
            #$NewObject = Initialize-ClusterMasters -Id "TEST_VALUE" -Uuid "TEST_VALUE" -ExternalId "TEST_VALUE" -InternalId "TEST_VALUE" -ExternalUniqueId "TEST_VALUE" -Name "TEST_VALUE" -ExternalName "TEST_VALUE" -Hostname "TEST_VALUE" -AccountId "TEST_VALUE" -Account "TEST_VALUE" -Owner "TEST_VALUE" -Zone "TEST_VALUE" -Plan "TEST_VALUE" -ComputeServerType "TEST_VALUE" -Visibility "TEST_VALUE" -Description "TEST_VALUE" -ZoneId "TEST_VALUE" -SiteId "TEST_VALUE" -ResourcePoolId "TEST_VALUE" -FolderId "TEST_VALUE" -SshHost "TEST_VALUE" -SshPort "TEST_VALUE" -ExternalIp "TEST_VALUE" -InternalIp "TEST_VALUE" -VolumeId "TEST_VALUE" -Platform "TEST_VALUE" -PlatformVersion "TEST_VALUE" -SshUsername "TEST_VALUE" -SshPassword "TEST_VALUE" -SshPasswordHash "TEST_VALUE" -OsDevice "TEST_VALUE" -OsType "TEST_VALUE" -DataDevice "TEST_VALUE" -LvmEnabled "TEST_VALUE" -ApiKey "TEST_VALUE" -SoftwareRaid "TEST_VALUE" -DateCreated "TEST_VALUE" -LastUpdated "TEST_VALUE" -Stats "TEST_VALUE" -Status "TEST_VALUE" -StatusMessage "TEST_VALUE" -ErrorMessage "TEST_VALUE" -StatusDate "TEST_VALUE" -StatusPercent "TEST_VALUE" -StatusEta "TEST_VALUE" -PowerState "TEST_VALUE" -AgentInstalled "TEST_VALUE" -LastAgentUpdate "TEST_VALUE" -AgentVersion "TEST_VALUE" -MaxCores "TEST_VALUE" -CoresPerSocket "TEST_VALUE" -MaxMemory "TEST_VALUE" -MaxStorage "TEST_VALUE" -MaxCpu "TEST_VALUE" -HourlyPrice "TEST_VALUE" -SourceImage "TEST_VALUE" -ServerOs "TEST_VALUE" -Volumes "TEST_VALUE" -Controllers "TEST_VALUE" -Interfaces "TEST_VALUE" -Labels "TEST_VALUE" -Tags "TEST_VALUE" -Enabled "TEST_VALUE" -TagCompliant "TEST_VALUE" -Containers "TEST_VALUE" -GuestConsolePreferred "TEST_VALUE" -GuestConsoleType "TEST_VALUE" -GuestConsoleUsername "TEST_VALUE" -GuestConsolePassword "TEST_VALUE" -GuestConsolePasswordHash "TEST_VALUE" -GuestConsolePort "TEST_VALUE"
            #$NewObject | Should -BeOfType ClusterMasters
            #$NewObject.property | Should -Be 0
        }
    }
}
