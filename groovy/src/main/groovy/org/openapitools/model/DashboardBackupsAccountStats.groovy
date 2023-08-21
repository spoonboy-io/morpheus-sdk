package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.DashboardBackupsAccountStatsFormattedTotalSize;
import org.openapitools.model.DashboardBackupsAccountStatsLastSevenDays;

@Canonical
class DashboardBackupsAccountStats {
    
    List<BigDecimal> totalSizeByDay = new ArrayList<BigDecimal>()
    
    List<BigDecimal> totalSizeByDay7Days = new ArrayList<BigDecimal>()
    
    DashboardBackupsAccountStatsFormattedTotalSize formattedTotalSize
    
    BigDecimal backupCount
    
    BigDecimal totalSize
    
    BigDecimal success
    
    BigDecimal failed
    
    BigDecimal totalCompleted
    
    DashboardBackupsAccountStatsLastSevenDays lastSevenDays
    
    BigDecimal avgSize
    
    BigDecimal failedRate
    
    BigDecimal successRate
    
    BigDecimal nextFireTotal
    
    List<BigDecimal> backupDayCount = new ArrayList<BigDecimal>()
    
    BigDecimal backupDayCountTotal
}
