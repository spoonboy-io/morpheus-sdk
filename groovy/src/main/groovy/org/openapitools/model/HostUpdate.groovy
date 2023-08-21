package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiServersIdMakeManagedServerTags;
import org.openapitools.model.InstanceUpdateInstanceRemoveTags;

@Canonical
class HostUpdate {
    /* Unique name scoped to your account for the server. */
    String name
    /* Optional description field. */
    String description
    /* SSH Username */
    String sshUsername
    /* SSH Password */
    String sshPassword
    /* Power schedule ID. */
    Long powerScheduleType
    
    List<String> labels = new ArrayList<String>()
    /* Metadata tags, Array of objects having a name and value. */
    List<ApiServersIdMakeManagedServerTags> tags = new ArrayList<ApiServersIdMakeManagedServerTags>()
    /* Add or update value of Metadata tags, Array of objects having a name and value. */
    List<ApiServersIdMakeManagedServerTags> addTags = new ArrayList<ApiServersIdMakeManagedServerTags>()
    /* Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. */
    List<InstanceUpdateInstanceRemoveTags> removeTags = new ArrayList<InstanceUpdateInstanceRemoveTags>()
    /* The Type of guest console this server provides such as disabled, vnc, rdp, ssh */
    String guestConsoleType
    /* The optional guest console username if you don't want to use the user defaults */
    String guestConsoleUsername
    /* The optional guest console password if not using the accessing users creds */
    String guestConsolePassword
    /* The port the guest console is being accessed from */
    String guestConsolePort
    /* Can turn off guest console preferences on server in favor of hypervisor console */
    Boolean guestConsolePreferred = true
}
