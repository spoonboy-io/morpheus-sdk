package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.UserSettingsUpdateDefaultCloud;
import org.openapitools.model.UserSettingsUpdateDefaultGroup;
import org.openapitools.model.UserSettingsUpdateDefaultPersona;

@Canonical
class UserSettingsUpdate {
    /* Username */
    String username
    /* Email */
    String email
    /* First Name */
    String firstName
    /* Last Name */
    String lastName
    /* Change your password */
    String password
    /* Linux Username */
    String linuxUsername
    /* Linux Password */
    String linuxPassword
    /* Linux Key Pair ID */
    Long linuxKeyPairId
    /* Windows Username */
    String windowsUsername
    /* Windows Password */
    String windowsPassword
    /* Receive Notifications (true or false) */
    Boolean receiveNotifications
    
    UserSettingsUpdateDefaultGroup defaultGroup
    
    UserSettingsUpdateDefaultCloud defaultCloud
    
    UserSettingsUpdateDefaultPersona defaultPersona
}
