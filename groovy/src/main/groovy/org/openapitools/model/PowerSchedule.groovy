package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class PowerSchedule {
    
    Long id
    
    String name
    
    String description
    
    String visibility
    
    Boolean enabled
    
    String scheduleType
    
    String scheduleTimezone
    
    Long mondayOn
    
    String mondayOnTime
    
    Long mondayOff
    
    String mondayOffTime
    
    Long tuesdayOn
    
    String tuesdayOnTime
    
    Long tuesdayOff
    
    String tuesdayOffTime
    
    Long wednesdayOn
    
    String wednesdayOnTime
    
    Long wednesdayOff
    
    String wednesdayOffTime
    
    Long thursdayOn
    
    String thursdayOnTime
    
    Long thursdayOff
    
    String thursdayOffTime
    
    Long fridayOn
    
    String fridayOnTime
    
    Long fridayOff
    
    String fridayOffTime
    
    Long saturdayOn
    
    String saturdayOnTime
    
    Long saturdayOff
    
    String saturdayOffTime
    
    Long sundayOn
    
    String sundayOnTime
    
    Long sundayOff
    
    String sundayOffTime
    
    BigDecimal totalMonthlyHoursSaved
    
    Date dateCreated
    
    Date lastUpdated
}
