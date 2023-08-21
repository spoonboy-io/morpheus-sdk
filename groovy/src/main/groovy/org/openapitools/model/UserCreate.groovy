package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class UserCreate {
    /* The user's first name (optional) */
    String firstName
    /* The user's last name (optional) */
    String lastName
    /* Username (unique per tenant). */
    String username
    /* Email address */
    String email
    /* Password to apply to the user */
    String password
    /* Array of objects with id of the role(s) to assign to the user. */
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> roles = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
    /* Receive Notifications? */
    Boolean receiveNotifications = true
    /* Linux Username, user settings for provisioning */
    String linuxUsername
    /* Linux Password, user settings for provisioning */
    String linuxPassword
    /* Linux SSH Key, user settings for provisioning */
    Long linuxKeyPairId
    /* Windows Username, user settings for provisioning */
    String windowsUsername
    /* Windows Password, user settings for provisioning */
    String windowsPassword
}
