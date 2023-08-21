package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.VirtualImageConfig;
import org.openapitools.model.VirtualImageOsType;

@Canonical
class VirtualImage {
    
    Long id
    
    String name
    
    String description
    
    List<String> labels = new ArrayList<String>()
    
    Long ownerId
    
    InlineResponse20040AppDeployInstance tenant
    
    String imageType
    
    Boolean userUploaded
    
    Boolean userDefined
    
    Boolean systemImage
    
    Boolean isCloudInit
    
    String sshUsername
    
    String sshPassword
    
    String sshPasswordHash
    
    String sshKey
    
    VirtualImageOsType osType
    
    Long minRam
    
    Long minRamGB
    
    String minDisk
    
    String minDiskGB
    
    Long rawSize
    
    BigDecimal rawSizeGB
    
    Boolean trialVersion
    
    Boolean virtioSupported
    
    String uefi
    
    Boolean isAutoJoinDomain
    
    Boolean vmToolsInstalled
    
    Boolean installAgent
    
    Boolean isForceCustomization
    
    Boolean isSysprep
    
    Boolean fipsEnabled
    
    String userData
    
    String consoleKeymap
    
    String storageProvider
    
    String externalId
    
    String visibility
    
    List<InlineResponse20040AppDeployInstance> accounts = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    VirtualImageConfig config
    
    List<Object> volumes = new ArrayList<Object>()
    
    List<Object> storageControllers = new ArrayList<Object>()
    
    List<Object> networkInterfaces = new ArrayList<Object>()
    
    List<Object> tags = new ArrayList<Object>()
    
    List<Object> locations = new ArrayList<Object>()
    
    Date dateCreated
    
    Date lastUpdated
    
    String status
}
