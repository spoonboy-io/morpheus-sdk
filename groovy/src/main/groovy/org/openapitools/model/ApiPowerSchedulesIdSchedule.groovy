package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiPowerSchedulesIdSchedule {
    /* A name for the power schedule */
    String name
    /* A description for the power schedule */
    String description
    /* Type of schedule `power` on or `power off` */
    String scheduleType
    /* Time Zone eg. America/New_York, Europe/Amsterdam, etc. */
    String scheduleTimezone = "UTC"
    /* Is the power schedule enabled */
    Boolean enabled = true
    /* Monday Start time of the day in 24-hour format */
    String mondayOnTime = "00:00"
    /* Monday Off time of the day in 24-hour format */
    String mondayOffTime = "24:00"
    /* Tuesday Start time of the day in 24-hour format */
    String tuesdayOnTime = "00:00"
    /* Tuesday Off time of the day in 24-hour format */
    String tuesdayOffTime = "24:00"
    /* Wednesday Start time of the day in 24-hour format */
    String wednesdayOnTime = "00:00"
    /* Wednesday Off time of the day in 24-hour format */
    String wednesdayOffTime = "24:00"
    /* Thursday Start time of the day in 24-hour format */
    String thursdayOnTime = "00:00"
    /* Thursday Off time of the day in 24-hour format */
    String thursdayOffTime = "24:00"
    /* Friday Start time of the day in 24-hour format */
    String fridayOnTime = "00:00"
    /* Friday Off time of the day in 24-hour format */
    String fridayOffTime = "24:00"
    /* Saturday Start time of the day in 24-hour format */
    String saturdayOnTime = "00:00"
    /* Saturday Off time of the day in 24-hour format */
    String saturdayOffTime = "24:00"
    /* Sunday Start time of the day in 24-hour format */
    String sundayOnTime = "00:00"
    /* Sunday Off time of the day in 24-hour format */
    String sundayOffTime = "24:00"
}
