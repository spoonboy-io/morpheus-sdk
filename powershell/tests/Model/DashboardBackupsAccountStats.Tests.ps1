#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'DashboardBackupsAccountStats' {
    Context 'DashboardBackupsAccountStats' {
        It 'Initialize-DashboardBackupsAccountStats' {
            # a simple test to create an object
            #$NewObject = Initialize-DashboardBackupsAccountStats -TotalSizeByDay "TEST_VALUE" -TotalSizeByDay7Days "TEST_VALUE" -FormattedTotalSize "TEST_VALUE" -BackupCount "TEST_VALUE" -TotalSize "TEST_VALUE" -Success "TEST_VALUE" -Failed "TEST_VALUE" -TotalCompleted "TEST_VALUE" -LastSevenDays "TEST_VALUE" -AvgSize "TEST_VALUE" -FailedRate "TEST_VALUE" -SuccessRate "TEST_VALUE" -NextFireTotal "TEST_VALUE" -BackupDayCount "TEST_VALUE" -BackupDayCountTotal "TEST_VALUE"
            #$NewObject | Should -BeOfType DashboardBackupsAccountStats
            #$NewObject.property | Should -Be 0
        }
    }
}
