package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class ApplianceSettings {
    
    String applianceUrl
    
    String internalApplianceUrl
    
    String corsAllowed
    
    Boolean registrationEnabled
    
    String defaultRoleId
    
    String defaultUserRoleId
    
    Boolean dockerPrivilegedMode
    
    String expirePwdDays
    
    String disableAfterAttempts
    
    String disableAfterDaysInactive
    
    String warnUserDaysBefore
    
    String smtpMailFrom
    
    String smtpServer
    
    String smtpPort
    
    Boolean smtpSSL
    
    Boolean smtpTLS
    
    String smtpUser
    
    String smtpPassword
    
    String smtpPasswordHash
    
    String proxyHost
    
    String proxyPort
    
    String proxyUser
    
    String proxyPassword
    
    String proxyPasswordHash
    
    String proxyDomain
    
    String proxyWorkstation
    
    String currencyProvider
    
    String currencyKey
    
    List<InlineResponse20040AppDeployInstance> enabledZoneTypes = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    Long statsRetainmentPeriod
}
