#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'HealthThreadsBusyThreadsInner' {
    Context 'HealthThreadsBusyThreadsInner' {
        It 'Initialize-HealthThreadsBusyThreadsInner' {
            # a simple test to create an object
            #$NewObject = Initialize-HealthThreadsBusyThreadsInner -ThreadId "TEST_VALUE" -Name "TEST_VALUE" -CpuTime "TEST_VALUE" -BlockedTime "TEST_VALUE" -LockName "TEST_VALUE" -LockOwnerId "TEST_VALUE" -LockOwnerName "TEST_VALUE" -State "TEST_VALUE" -WaitedCount "TEST_VALUE" -WaitedTime "TEST_VALUE" -IsInNative "TEST_VALUE" -IsSuspended "TEST_VALUE" -LockedMonitors "TEST_VALUE" -LockedSynchronizers "TEST_VALUE" -LockInfo "TEST_VALUE" -CurrentLines "TEST_VALUE" -CpuPercent "TEST_VALUE"
            #$NewObject | Should -BeOfType HealthThreadsBusyThreadsInner
            #$NewObject.property | Should -Be 0
        }
    }
}
