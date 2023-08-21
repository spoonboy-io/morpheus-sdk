package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OneOfobjectobject;
import org.openapitools.model.OneOfobjectstring;
import org.openapitools.model.VirtualImageCreateStorageProvider;
import org.openapitools.model.VirtualImageCreateTags;

@Canonical
class VirtualImageCreate {
    /* A name for the virtual image */
    String name
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Code of image type. eg. vmware, ami, etc. */
    String imageType
    
    VirtualImageCreateStorageProvider storageProvider
    /* Cloud Init Enabled? */
    Boolean isCloudInit = false
    /* Cloud-Init User Data, a bash script */
    String userData
    /* Install Agent? */
    Boolean installAgent = false
    /* SSH Username */
    String sshUsername
    /* SSH Password */
    String sshPassword
    /* SSH Key */
    String sshKey
    /* A Map containing the id of the OS Type. This can also be passed as a string (code or name) instead. */
    OneOfobjectstring osType = null
    /* private or public */
    String visibility = "private"
    
    List<Long> accounts = new ArrayList<Long>()
    /* Auto Join Domain? */
    Boolean isAutoJoinDomain = false
    /* VirtIO Drivers Loaded? */
    Boolean virtioSupported = true
    /* VM Tools Installed? */
    Boolean vmToolsInstalled = true
    /* Force Guest Customization? */
    Boolean isForceCustomization = false
    /* Trial Version */
    Boolean trialVersion = false
    /* Sysprep Enabled? */
    Boolean isSysprep = false
    /* Map of configuration properties, varies by image type. */
    OneOfobjectobject config = null
    /* Metadata tags, Array of objects having a name and value */
    List<VirtualImageCreateTags> tags = new ArrayList<VirtualImageCreateTags>()
}
