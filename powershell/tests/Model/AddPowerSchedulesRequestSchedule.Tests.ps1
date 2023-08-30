#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

Describe -tag 'PSOpenAPITools' -name 'AddPowerSchedulesRequestSchedule' {
    Context 'AddPowerSchedulesRequestSchedule' {
        It 'Initialize-AddPowerSchedulesRequestSchedule' {
            # a simple test to create an object
            #$NewObject = Initialize-AddPowerSchedulesRequestSchedule -Name "TEST_VALUE" -Description "TEST_VALUE" -ScheduleType "TEST_VALUE" -ScheduleTimezone "TEST_VALUE" -Enabled "TEST_VALUE" -MondayOnTime "TEST_VALUE" -MondayOffTime "TEST_VALUE" -TuesdayOnTime "TEST_VALUE" -TuesdayOffTime "TEST_VALUE" -WednesdayOnTime "TEST_VALUE" -WednesdayOffTime "TEST_VALUE" -ThursdayOnTime "TEST_VALUE" -ThursdayOffTime "TEST_VALUE" -FridayOnTime "TEST_VALUE" -FridayOffTime "TEST_VALUE" -SaturdayOnTime "TEST_VALUE" -SaturdayOffTime "TEST_VALUE" -SundayOnTime "TEST_VALUE" -SundayOffTime "TEST_VALUE"
            #$NewObject | Should -BeOfType AddPowerSchedulesRequestSchedule
            #$NewObject.property | Should -Be 0
        }
    }
}
