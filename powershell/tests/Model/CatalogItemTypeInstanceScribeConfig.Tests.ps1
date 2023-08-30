#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'CatalogItemTypeInstanceScribeConfig' {
    Context 'CatalogItemTypeInstanceScribeConfig' {
        It 'Initialize-CatalogItemTypeInstanceScribeConfig' {
            # a simple test to create an object
            #$NewObject = Initialize-CatalogItemTypeInstanceScribeConfig -ResourcePoolId "TEST_VALUE" -AvailabilityOptions "TEST_VALUE" -AvailabilitySet "TEST_VALUE" -AvailabilityZone "TEST_VALUE" -AzurefloatingIp "TEST_VALUE" -BootDiagnostics "TEST_VALUE" -OsGuestDiagnostics "TEST_VALUE" -DiagnosticsStorageAccount "TEST_VALUE" -CreateUser "TEST_VALUE" -NoAgent "TEST_VALUE" -HostId "TEST_VALUE" -SmbiosAssetTag "TEST_VALUE" -NestedVirtualization "TEST_VALUE" -VmwareFolderId "TEST_VALUE" -GoogleZoneId "TEST_VALUE" -ExternalIpType "TEST_VALUE" -NetworkTags "TEST_VALUE" -ServiceAccount "TEST_VALUE" -AccessScope "TEST_VALUE" -IsEC2 "TEST_VALUE" -AvailabilityId "TEST_VALUE" -SecurityId "TEST_VALUE" -PublicIpType "TEST_VALUE" -InstanceProfile "TEST_VALUE" -KmsKeyId "TEST_VALUE"
            #$NewObject | Should -BeOfType CatalogItemTypeInstanceScribeConfig
            #$NewObject.property | Should -Be 0
        }
    }
}