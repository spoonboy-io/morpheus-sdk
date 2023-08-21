package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.RoleAppTemplatePermissions;
import org.openapitools.model.RoleFeaturePermissions;
import org.openapitools.model.RoleInstanceTypePermissions;
import org.openapitools.model.RoleRole;
import org.openapitools.model.RoleSites;

@Canonical
class Role {
    
    RoleRole role
    
    List<RoleFeaturePermissions> featurePermissions = new ArrayList<RoleFeaturePermissions>()
    
    String globalSiteAccess
    
    List<RoleSites> sites = new ArrayList<RoleSites>()
    
    String globalZoneAccess
    
    List<RoleSites> zones = new ArrayList<RoleSites>()
    
    String globalInstanceTypeAccess
    
    List<RoleInstanceTypePermissions> instanceTypePermissions = new ArrayList<RoleInstanceTypePermissions>()
    
    String globalAppTemplateAccess
    
    List<RoleAppTemplatePermissions> appTemplatePermissions = new ArrayList<RoleAppTemplatePermissions>()
    
    String globalCatalogItemTypeAccess
    
    List<RoleSites> catalogItemTypePermissions = new ArrayList<RoleSites>()
    
    String globalPersonaAccess
    
    List<RoleInstanceTypePermissions> personaPermissions = new ArrayList<RoleInstanceTypePermissions>()
    
    String globalVdiPoolAccess
    
    List<RoleSites> vdiPoolPermissions = new ArrayList<RoleSites>()
    
    String globalReportTypeAccess
    
    List<RoleInstanceTypePermissions> reportTypePermissions = new ArrayList<RoleInstanceTypePermissions>()
    
    String globalTaskAccess
    
    List<RoleAppTemplatePermissions> taskPermissions = new ArrayList<RoleAppTemplatePermissions>()
    
    String globalTaskSetAccess
    
    List<RoleAppTemplatePermissions> taskSetPermissions = new ArrayList<RoleAppTemplatePermissions>()
}
