package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;

@Canonical
class ApplianceSettingsUpdate {
    /* Appliance URL */
    String applianceUrl
    /* Internal Appliance URL (PXE) */
    String internalApplianceUrl
    /* API Allowed Origins */
    String corsAllowed
    /* Registration enabled (true, false) */
    Boolean registrationEnabled
    /* Default tenant role ID */
    Long defaultRoleId
    /* Default user role ID */
    Long defaultUserRoleId
    /* Docker privileged mode (true, false) */
    Boolean dockerPrivilegedMode
    /* Min Password Length */
    String passwordMinLength
    /* Min Password Uppercase */
    String passwordMinUpperCase
    /* Min Password Numbers */
    String passwordMinNumbers
    /* Min Password Symbols */
    String passwordMinSymbols
    /* User Browser Session Timeout (Minutes) */
    String userBrowserSessionTimeout
    /* User Browser Session Warning (Minutes) */
    String userBrowserSessionWarning
    /* Expire password after days. Setting to 0 disabled this feature */
    Long expirePwdDays
    /* Disable user after number of attempts. Set to 0 to disable this feature */
    Long disableAfterAttempts
    /* Disable user if inactive for specified days. Set to 0 to disable this feature */
    Long disableAfterDaysInactive
    /* Send warning email number of days in advance before deactivating. Set to 0 to disable this feature */
    Long warnUserDaysBefore
    /* From email address */
    String smtpMailFrom
    /* SMTP server / host */
    String smtpServer
    /* SMTP port */
    Long smtpPort
    /* Use SSL for SMTP connection */
    Boolean smtpSSL
    /* Use TLS for SMTP connections */
    Boolean smtpTLS
    /* SMTP username */
    String smtpUser
    /* SMTP password */
    String smtpPassword
    /* Proxy host */
    String proxyHost
    /* Proxy port */
    String proxyPort
    /* Proxy username */
    String proxyUser
    /* Proxy password */
    String proxyPassword
    /* Proxy domain */
    String proxyDomain
    /* Proxy workstation */
    String proxyWorkstation
    /* Currency provider */
    String currencyProvider
    /* Currency provider API key */
    String currencyKey
    /* Set all cloud types enabled status on, overrides enableZoneTypes and disableZoneTypes parameters */
    Boolean enableAllZoneTypes
    /* List of cloud type IDs to set enabled status on */
    List<Long> enableZoneTypes
    /* List of cloud type IDs to set enabled status off */
    List<Long> disableZoneTypes
    /* Set all cloud types enabled status off, can be used in conjunction with enableZoneTypes */
    Boolean disableAllZoneTypes
}
