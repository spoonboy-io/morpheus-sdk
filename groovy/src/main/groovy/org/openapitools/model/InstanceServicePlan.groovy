package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InstanceServicePlanAutoOptions;
import org.openapitools.model.InstanceServicePlanDatastores;
import org.openapitools.model.InstanceServicePlanStorageType;

@Canonical
class InstanceServicePlan {
    
    Integer id
    
    String name
    
    Integer value
    
    String code
    
    Long maxStorage
    
    Integer maxMemory
    
    Integer maxCpu
    
    Integer maxCores
    
    Boolean customCpu
    
    Boolean customMaxMemory
    
    Boolean customMaxStorage
    
    Boolean customMaxDataStorage
    
    Boolean customCoresPerSocket
    
    Integer coresPerSocket
    
    List<InstanceServicePlanStorageType> storageTypes = new ArrayList<InstanceServicePlanStorageType>()
    
    List<InstanceServicePlanStorageType> rootStorageTypes = new ArrayList<InstanceServicePlanStorageType>()
    
    Boolean addVolumes
    
    Boolean customizeVolume
    
    Boolean rootDiskCustomizable
    
    Boolean noDisks
    
    Boolean hasDatastore
    
    Integer minDisk
    
    String maxDisk
    
    Boolean lvmSupported
    
    InstanceServicePlanDatastores datastores
    
    Boolean supportsAutoDatastore
    
    List<InstanceServicePlanAutoOptions> autoOptions = new ArrayList<InstanceServicePlanAutoOptions>()
    
    List<Object> cpuOptions = new ArrayList<Object>()
    
    List<Object> coreOptions = new ArrayList<Object>()
    
    List<Object> memoryOptions = new ArrayList<Object>()
    
    Object rootCustomSizeOptions
    
    Object customSizeOptions
    
    Boolean customCores
    
    String maxDisks
    
    String memorySizeType
}
