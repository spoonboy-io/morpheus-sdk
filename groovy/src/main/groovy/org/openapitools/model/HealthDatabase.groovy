package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.HealthDatabaseInnodbStats;
import org.openapitools.model.HealthDatabaseScans;
import org.openapitools.model.HealthDatabaseSlowQueries;
import org.openapitools.model.HealthDatabaseStats;

@Canonical
class HealthDatabase {
    
    Boolean success
    
    List<Object> connectionList = new ArrayList<Object>()
    
    List<String> busyConnections = new ArrayList<String>()
    
    Long maxConnections
    
    Long maxUsedConnections
    
    Long usedConnections
    
    Long abortedConnections
    
    String innodbStatus
    
    HealthDatabaseStats stats
    
    HealthDatabaseScans scans
    
    List<HealthDatabaseSlowQueries> slowQueries = new ArrayList<HealthDatabaseSlowQueries>()
    
    HealthDatabaseInnodbStats innodbStats
    
    BigDecimal scanPercent
    
    String status
}
