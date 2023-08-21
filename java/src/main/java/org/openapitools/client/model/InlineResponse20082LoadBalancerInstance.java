/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import org.openapitools.client.model.InlineResponse20079LoadBalancerMonitorLoadBalancer;
import org.openapitools.client.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.threeten.bp.OffsetDateTime;

/**
 * InlineResponse20082LoadBalancerInstance
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20082LoadBalancerInstance {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_LOAD_BALANCER = "loadBalancer";
  @SerializedName(SERIALIZED_NAME_LOAD_BALANCER)
  private InlineResponse20079LoadBalancerMonitorLoadBalancer loadBalancer;

  public static final String SERIALIZED_NAME_INSTANCE = "instance";
  @SerializedName(SERIALIZED_NAME_INSTANCE)
  private String instance;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_INTERNAL_ID = "internalId";
  @SerializedName(SERIALIZED_NAME_INTERNAL_ID)
  private String internalId;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active;

  public static final String SERIALIZED_NAME_STICKY = "sticky";
  @SerializedName(SERIALIZED_NAME_STICKY)
  private Boolean sticky;

  public static final String SERIALIZED_NAME_SSL_ENABLED = "sslEnabled";
  @SerializedName(SERIALIZED_NAME_SSL_ENABLED)
  private String sslEnabled;

  public static final String SERIALIZED_NAME_EXTERNAL_ADDRESS = "externalAddress";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ADDRESS)
  private Boolean externalAddress;

  public static final String SERIALIZED_NAME_BACKEND_PORT = "backendPort";
  @SerializedName(SERIALIZED_NAME_BACKEND_PORT)
  private String backendPort;

  public static final String SERIALIZED_NAME_VIP_TYPE = "vipType";
  @SerializedName(SERIALIZED_NAME_VIP_TYPE)
  private String vipType;

  public static final String SERIALIZED_NAME_VIP_ADDRESS = "vipAddress";
  @SerializedName(SERIALIZED_NAME_VIP_ADDRESS)
  private String vipAddress;

  public static final String SERIALIZED_NAME_VIP_HOSTNAME = "vipHostname";
  @SerializedName(SERIALIZED_NAME_VIP_HOSTNAME)
  private String vipHostname;

  public static final String SERIALIZED_NAME_VIP_PROTOCOL = "vipProtocol";
  @SerializedName(SERIALIZED_NAME_VIP_PROTOCOL)
  private String vipProtocol;

  public static final String SERIALIZED_NAME_VIP_SCHEME = "vipScheme";
  @SerializedName(SERIALIZED_NAME_VIP_SCHEME)
  private String vipScheme;

  public static final String SERIALIZED_NAME_VIP_MODE = "vipMode";
  @SerializedName(SERIALIZED_NAME_VIP_MODE)
  private String vipMode;

  public static final String SERIALIZED_NAME_VIP_NAME = "vipName";
  @SerializedName(SERIALIZED_NAME_VIP_NAME)
  private String vipName;

  public static final String SERIALIZED_NAME_VIP_PORT = "vipPort";
  @SerializedName(SERIALIZED_NAME_VIP_PORT)
  private Long vipPort;

  public static final String SERIALIZED_NAME_VIP_STICKY = "vipSticky";
  @SerializedName(SERIALIZED_NAME_VIP_STICKY)
  private String vipSticky;

  public static final String SERIALIZED_NAME_VIP_BALANCE = "vipBalance";
  @SerializedName(SERIALIZED_NAME_VIP_BALANCE)
  private String vipBalance;

  public static final String SERIALIZED_NAME_SERVICE_PORT = "servicePort";
  @SerializedName(SERIALIZED_NAME_SERVICE_PORT)
  private String servicePort;

  public static final String SERIALIZED_NAME_SOURCE_ADDRESS = "sourceAddress";
  @SerializedName(SERIALIZED_NAME_SOURCE_ADDRESS)
  private String sourceAddress;

  public static final String SERIALIZED_NAME_SSL_CERT = "sslCert";
  @SerializedName(SERIALIZED_NAME_SSL_CERT)
  private InlineResponse20082LoadBalancerInstanceSslCert sslCert;

  public static final String SERIALIZED_NAME_SSL_MODE = "sslMode";
  @SerializedName(SERIALIZED_NAME_SSL_MODE)
  private String sslMode;

  public static final String SERIALIZED_NAME_SSL_REDIRECT_MODE = "sslRedirectMode";
  @SerializedName(SERIALIZED_NAME_SSL_REDIRECT_MODE)
  private String sslRedirectMode;

  public static final String SERIALIZED_NAME_VIP_SHARED = "vipShared";
  @SerializedName(SERIALIZED_NAME_VIP_SHARED)
  private Boolean vipShared;

  public static final String SERIALIZED_NAME_VIP_DIRECT_ADDRESS = "vipDirectAddress";
  @SerializedName(SERIALIZED_NAME_VIP_DIRECT_ADDRESS)
  private String vipDirectAddress;

  public static final String SERIALIZED_NAME_SERVER_NAME = "serverName";
  @SerializedName(SERIALIZED_NAME_SERVER_NAME)
  private String serverName;

  public static final String SERIALIZED_NAME_POOL_NAME = "poolName";
  @SerializedName(SERIALIZED_NAME_POOL_NAME)
  private String poolName;

  public static final String SERIALIZED_NAME_REMOVING = "removing";
  @SerializedName(SERIALIZED_NAME_REMOVING)
  private Boolean removing;

  public static final String SERIALIZED_NAME_VIP_SOURCE = "vipSource";
  @SerializedName(SERIALIZED_NAME_VIP_SOURCE)
  private String vipSource;

  public static final String SERIALIZED_NAME_EXTRA_CONFIG = "extraConfig";
  @SerializedName(SERIALIZED_NAME_EXTRA_CONFIG)
  private String extraConfig;

  public static final String SERIALIZED_NAME_SERVICE_ACCESS = "serviceAccess";
  @SerializedName(SERIALIZED_NAME_SERVICE_ACCESS)
  private String serviceAccess;

  public static final String SERIALIZED_NAME_NETWORK_ID = "networkId";
  @SerializedName(SERIALIZED_NAME_NETWORK_ID)
  private String networkId;

  public static final String SERIALIZED_NAME_SUBNET_ID = "subnetId";
  @SerializedName(SERIALIZED_NAME_SUBNET_ID)
  private String subnetId;

  public static final String SERIALIZED_NAME_EXTERNAL_PORT_ID = "externalPortId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_PORT_ID)
  private String externalPortId;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_VIP_STATUS = "vipStatus";
  @SerializedName(SERIALIZED_NAME_VIP_STATUS)
  private String vipStatus;


  public InlineResponse20082LoadBalancerInstance id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public InlineResponse20082LoadBalancerInstance loadBalancer(InlineResponse20079LoadBalancerMonitorLoadBalancer loadBalancer) {
    
    this.loadBalancer = loadBalancer;
    return this;
  }

   /**
   * Get loadBalancer
   * @return loadBalancer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20079LoadBalancerMonitorLoadBalancer getLoadBalancer() {
    return loadBalancer;
  }


  public void setLoadBalancer(InlineResponse20079LoadBalancerMonitorLoadBalancer loadBalancer) {
    this.loadBalancer = loadBalancer;
  }


  public InlineResponse20082LoadBalancerInstance instance(String instance) {
    
    this.instance = instance;
    return this;
  }

   /**
   * Get instance
   * @return instance
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInstance() {
    return instance;
  }


  public void setInstance(String instance) {
    this.instance = instance;
  }


  public InlineResponse20082LoadBalancerInstance description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public InlineResponse20082LoadBalancerInstance internalId(String internalId) {
    
    this.internalId = internalId;
    return this;
  }

   /**
   * Get internalId
   * @return internalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInternalId() {
    return internalId;
  }


  public void setInternalId(String internalId) {
    this.internalId = internalId;
  }


  public InlineResponse20082LoadBalancerInstance externalId(String externalId) {
    
    this.externalId = externalId;
    return this;
  }

   /**
   * Get externalId
   * @return externalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalId() {
    return externalId;
  }


  public void setExternalId(String externalId) {
    this.externalId = externalId;
  }


  public InlineResponse20082LoadBalancerInstance dateCreated(OffsetDateTime dateCreated) {
    
    this.dateCreated = dateCreated;
    return this;
  }

   /**
   * Get dateCreated
   * @return dateCreated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getDateCreated() {
    return dateCreated;
  }


  public void setDateCreated(OffsetDateTime dateCreated) {
    this.dateCreated = dateCreated;
  }


  public InlineResponse20082LoadBalancerInstance lastUpdated(OffsetDateTime lastUpdated) {
    
    this.lastUpdated = lastUpdated;
    return this;
  }

   /**
   * Get lastUpdated
   * @return lastUpdated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastUpdated() {
    return lastUpdated;
  }


  public void setLastUpdated(OffsetDateTime lastUpdated) {
    this.lastUpdated = lastUpdated;
  }


  public InlineResponse20082LoadBalancerInstance active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Get active
   * @return active
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public InlineResponse20082LoadBalancerInstance sticky(Boolean sticky) {
    
    this.sticky = sticky;
    return this;
  }

   /**
   * Get sticky
   * @return sticky
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getSticky() {
    return sticky;
  }


  public void setSticky(Boolean sticky) {
    this.sticky = sticky;
  }


  public InlineResponse20082LoadBalancerInstance sslEnabled(String sslEnabled) {
    
    this.sslEnabled = sslEnabled;
    return this;
  }

   /**
   * Get sslEnabled
   * @return sslEnabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSslEnabled() {
    return sslEnabled;
  }


  public void setSslEnabled(String sslEnabled) {
    this.sslEnabled = sslEnabled;
  }


  public InlineResponse20082LoadBalancerInstance externalAddress(Boolean externalAddress) {
    
    this.externalAddress = externalAddress;
    return this;
  }

   /**
   * Get externalAddress
   * @return externalAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getExternalAddress() {
    return externalAddress;
  }


  public void setExternalAddress(Boolean externalAddress) {
    this.externalAddress = externalAddress;
  }


  public InlineResponse20082LoadBalancerInstance backendPort(String backendPort) {
    
    this.backendPort = backendPort;
    return this;
  }

   /**
   * Get backendPort
   * @return backendPort
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBackendPort() {
    return backendPort;
  }


  public void setBackendPort(String backendPort) {
    this.backendPort = backendPort;
  }


  public InlineResponse20082LoadBalancerInstance vipType(String vipType) {
    
    this.vipType = vipType;
    return this;
  }

   /**
   * Get vipType
   * @return vipType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipType() {
    return vipType;
  }


  public void setVipType(String vipType) {
    this.vipType = vipType;
  }


  public InlineResponse20082LoadBalancerInstance vipAddress(String vipAddress) {
    
    this.vipAddress = vipAddress;
    return this;
  }

   /**
   * Get vipAddress
   * @return vipAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipAddress() {
    return vipAddress;
  }


  public void setVipAddress(String vipAddress) {
    this.vipAddress = vipAddress;
  }


  public InlineResponse20082LoadBalancerInstance vipHostname(String vipHostname) {
    
    this.vipHostname = vipHostname;
    return this;
  }

   /**
   * Get vipHostname
   * @return vipHostname
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipHostname() {
    return vipHostname;
  }


  public void setVipHostname(String vipHostname) {
    this.vipHostname = vipHostname;
  }


  public InlineResponse20082LoadBalancerInstance vipProtocol(String vipProtocol) {
    
    this.vipProtocol = vipProtocol;
    return this;
  }

   /**
   * Get vipProtocol
   * @return vipProtocol
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipProtocol() {
    return vipProtocol;
  }


  public void setVipProtocol(String vipProtocol) {
    this.vipProtocol = vipProtocol;
  }


  public InlineResponse20082LoadBalancerInstance vipScheme(String vipScheme) {
    
    this.vipScheme = vipScheme;
    return this;
  }

   /**
   * Get vipScheme
   * @return vipScheme
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipScheme() {
    return vipScheme;
  }


  public void setVipScheme(String vipScheme) {
    this.vipScheme = vipScheme;
  }


  public InlineResponse20082LoadBalancerInstance vipMode(String vipMode) {
    
    this.vipMode = vipMode;
    return this;
  }

   /**
   * Get vipMode
   * @return vipMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipMode() {
    return vipMode;
  }


  public void setVipMode(String vipMode) {
    this.vipMode = vipMode;
  }


  public InlineResponse20082LoadBalancerInstance vipName(String vipName) {
    
    this.vipName = vipName;
    return this;
  }

   /**
   * Get vipName
   * @return vipName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipName() {
    return vipName;
  }


  public void setVipName(String vipName) {
    this.vipName = vipName;
  }


  public InlineResponse20082LoadBalancerInstance vipPort(Long vipPort) {
    
    this.vipPort = vipPort;
    return this;
  }

   /**
   * Get vipPort
   * @return vipPort
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getVipPort() {
    return vipPort;
  }


  public void setVipPort(Long vipPort) {
    this.vipPort = vipPort;
  }


  public InlineResponse20082LoadBalancerInstance vipSticky(String vipSticky) {
    
    this.vipSticky = vipSticky;
    return this;
  }

   /**
   * Get vipSticky
   * @return vipSticky
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipSticky() {
    return vipSticky;
  }


  public void setVipSticky(String vipSticky) {
    this.vipSticky = vipSticky;
  }


  public InlineResponse20082LoadBalancerInstance vipBalance(String vipBalance) {
    
    this.vipBalance = vipBalance;
    return this;
  }

   /**
   * Get vipBalance
   * @return vipBalance
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipBalance() {
    return vipBalance;
  }


  public void setVipBalance(String vipBalance) {
    this.vipBalance = vipBalance;
  }


  public InlineResponse20082LoadBalancerInstance servicePort(String servicePort) {
    
    this.servicePort = servicePort;
    return this;
  }

   /**
   * Get servicePort
   * @return servicePort
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServicePort() {
    return servicePort;
  }


  public void setServicePort(String servicePort) {
    this.servicePort = servicePort;
  }


  public InlineResponse20082LoadBalancerInstance sourceAddress(String sourceAddress) {
    
    this.sourceAddress = sourceAddress;
    return this;
  }

   /**
   * Get sourceAddress
   * @return sourceAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceAddress() {
    return sourceAddress;
  }


  public void setSourceAddress(String sourceAddress) {
    this.sourceAddress = sourceAddress;
  }


  public InlineResponse20082LoadBalancerInstance sslCert(InlineResponse20082LoadBalancerInstanceSslCert sslCert) {
    
    this.sslCert = sslCert;
    return this;
  }

   /**
   * Get sslCert
   * @return sslCert
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getSslCert() {
    return sslCert;
  }


  public void setSslCert(InlineResponse20082LoadBalancerInstanceSslCert sslCert) {
    this.sslCert = sslCert;
  }


  public InlineResponse20082LoadBalancerInstance sslMode(String sslMode) {
    
    this.sslMode = sslMode;
    return this;
  }

   /**
   * Get sslMode
   * @return sslMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSslMode() {
    return sslMode;
  }


  public void setSslMode(String sslMode) {
    this.sslMode = sslMode;
  }


  public InlineResponse20082LoadBalancerInstance sslRedirectMode(String sslRedirectMode) {
    
    this.sslRedirectMode = sslRedirectMode;
    return this;
  }

   /**
   * Get sslRedirectMode
   * @return sslRedirectMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSslRedirectMode() {
    return sslRedirectMode;
  }


  public void setSslRedirectMode(String sslRedirectMode) {
    this.sslRedirectMode = sslRedirectMode;
  }


  public InlineResponse20082LoadBalancerInstance vipShared(Boolean vipShared) {
    
    this.vipShared = vipShared;
    return this;
  }

   /**
   * Get vipShared
   * @return vipShared
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getVipShared() {
    return vipShared;
  }


  public void setVipShared(Boolean vipShared) {
    this.vipShared = vipShared;
  }


  public InlineResponse20082LoadBalancerInstance vipDirectAddress(String vipDirectAddress) {
    
    this.vipDirectAddress = vipDirectAddress;
    return this;
  }

   /**
   * Get vipDirectAddress
   * @return vipDirectAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipDirectAddress() {
    return vipDirectAddress;
  }


  public void setVipDirectAddress(String vipDirectAddress) {
    this.vipDirectAddress = vipDirectAddress;
  }


  public InlineResponse20082LoadBalancerInstance serverName(String serverName) {
    
    this.serverName = serverName;
    return this;
  }

   /**
   * Get serverName
   * @return serverName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServerName() {
    return serverName;
  }


  public void setServerName(String serverName) {
    this.serverName = serverName;
  }


  public InlineResponse20082LoadBalancerInstance poolName(String poolName) {
    
    this.poolName = poolName;
    return this;
  }

   /**
   * Get poolName
   * @return poolName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPoolName() {
    return poolName;
  }


  public void setPoolName(String poolName) {
    this.poolName = poolName;
  }


  public InlineResponse20082LoadBalancerInstance removing(Boolean removing) {
    
    this.removing = removing;
    return this;
  }

   /**
   * Get removing
   * @return removing
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRemoving() {
    return removing;
  }


  public void setRemoving(Boolean removing) {
    this.removing = removing;
  }


  public InlineResponse20082LoadBalancerInstance vipSource(String vipSource) {
    
    this.vipSource = vipSource;
    return this;
  }

   /**
   * Get vipSource
   * @return vipSource
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipSource() {
    return vipSource;
  }


  public void setVipSource(String vipSource) {
    this.vipSource = vipSource;
  }


  public InlineResponse20082LoadBalancerInstance extraConfig(String extraConfig) {
    
    this.extraConfig = extraConfig;
    return this;
  }

   /**
   * Get extraConfig
   * @return extraConfig
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExtraConfig() {
    return extraConfig;
  }


  public void setExtraConfig(String extraConfig) {
    this.extraConfig = extraConfig;
  }


  public InlineResponse20082LoadBalancerInstance serviceAccess(String serviceAccess) {
    
    this.serviceAccess = serviceAccess;
    return this;
  }

   /**
   * Get serviceAccess
   * @return serviceAccess
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServiceAccess() {
    return serviceAccess;
  }


  public void setServiceAccess(String serviceAccess) {
    this.serviceAccess = serviceAccess;
  }


  public InlineResponse20082LoadBalancerInstance networkId(String networkId) {
    
    this.networkId = networkId;
    return this;
  }

   /**
   * Get networkId
   * @return networkId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNetworkId() {
    return networkId;
  }


  public void setNetworkId(String networkId) {
    this.networkId = networkId;
  }


  public InlineResponse20082LoadBalancerInstance subnetId(String subnetId) {
    
    this.subnetId = subnetId;
    return this;
  }

   /**
   * Get subnetId
   * @return subnetId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSubnetId() {
    return subnetId;
  }


  public void setSubnetId(String subnetId) {
    this.subnetId = subnetId;
  }


  public InlineResponse20082LoadBalancerInstance externalPortId(String externalPortId) {
    
    this.externalPortId = externalPortId;
    return this;
  }

   /**
   * Get externalPortId
   * @return externalPortId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalPortId() {
    return externalPortId;
  }


  public void setExternalPortId(String externalPortId) {
    this.externalPortId = externalPortId;
  }


  public InlineResponse20082LoadBalancerInstance status(String status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatus() {
    return status;
  }


  public void setStatus(String status) {
    this.status = status;
  }


  public InlineResponse20082LoadBalancerInstance vipStatus(String vipStatus) {
    
    this.vipStatus = vipStatus;
    return this;
  }

   /**
   * Get vipStatus
   * @return vipStatus
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVipStatus() {
    return vipStatus;
  }


  public void setVipStatus(String vipStatus) {
    this.vipStatus = vipStatus;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse20082LoadBalancerInstance inlineResponse20082LoadBalancerInstance = (InlineResponse20082LoadBalancerInstance) o;
    return Objects.equals(this.id, inlineResponse20082LoadBalancerInstance.id) &&
        Objects.equals(this.loadBalancer, inlineResponse20082LoadBalancerInstance.loadBalancer) &&
        Objects.equals(this.instance, inlineResponse20082LoadBalancerInstance.instance) &&
        Objects.equals(this.description, inlineResponse20082LoadBalancerInstance.description) &&
        Objects.equals(this.internalId, inlineResponse20082LoadBalancerInstance.internalId) &&
        Objects.equals(this.externalId, inlineResponse20082LoadBalancerInstance.externalId) &&
        Objects.equals(this.dateCreated, inlineResponse20082LoadBalancerInstance.dateCreated) &&
        Objects.equals(this.lastUpdated, inlineResponse20082LoadBalancerInstance.lastUpdated) &&
        Objects.equals(this.active, inlineResponse20082LoadBalancerInstance.active) &&
        Objects.equals(this.sticky, inlineResponse20082LoadBalancerInstance.sticky) &&
        Objects.equals(this.sslEnabled, inlineResponse20082LoadBalancerInstance.sslEnabled) &&
        Objects.equals(this.externalAddress, inlineResponse20082LoadBalancerInstance.externalAddress) &&
        Objects.equals(this.backendPort, inlineResponse20082LoadBalancerInstance.backendPort) &&
        Objects.equals(this.vipType, inlineResponse20082LoadBalancerInstance.vipType) &&
        Objects.equals(this.vipAddress, inlineResponse20082LoadBalancerInstance.vipAddress) &&
        Objects.equals(this.vipHostname, inlineResponse20082LoadBalancerInstance.vipHostname) &&
        Objects.equals(this.vipProtocol, inlineResponse20082LoadBalancerInstance.vipProtocol) &&
        Objects.equals(this.vipScheme, inlineResponse20082LoadBalancerInstance.vipScheme) &&
        Objects.equals(this.vipMode, inlineResponse20082LoadBalancerInstance.vipMode) &&
        Objects.equals(this.vipName, inlineResponse20082LoadBalancerInstance.vipName) &&
        Objects.equals(this.vipPort, inlineResponse20082LoadBalancerInstance.vipPort) &&
        Objects.equals(this.vipSticky, inlineResponse20082LoadBalancerInstance.vipSticky) &&
        Objects.equals(this.vipBalance, inlineResponse20082LoadBalancerInstance.vipBalance) &&
        Objects.equals(this.servicePort, inlineResponse20082LoadBalancerInstance.servicePort) &&
        Objects.equals(this.sourceAddress, inlineResponse20082LoadBalancerInstance.sourceAddress) &&
        Objects.equals(this.sslCert, inlineResponse20082LoadBalancerInstance.sslCert) &&
        Objects.equals(this.sslMode, inlineResponse20082LoadBalancerInstance.sslMode) &&
        Objects.equals(this.sslRedirectMode, inlineResponse20082LoadBalancerInstance.sslRedirectMode) &&
        Objects.equals(this.vipShared, inlineResponse20082LoadBalancerInstance.vipShared) &&
        Objects.equals(this.vipDirectAddress, inlineResponse20082LoadBalancerInstance.vipDirectAddress) &&
        Objects.equals(this.serverName, inlineResponse20082LoadBalancerInstance.serverName) &&
        Objects.equals(this.poolName, inlineResponse20082LoadBalancerInstance.poolName) &&
        Objects.equals(this.removing, inlineResponse20082LoadBalancerInstance.removing) &&
        Objects.equals(this.vipSource, inlineResponse20082LoadBalancerInstance.vipSource) &&
        Objects.equals(this.extraConfig, inlineResponse20082LoadBalancerInstance.extraConfig) &&
        Objects.equals(this.serviceAccess, inlineResponse20082LoadBalancerInstance.serviceAccess) &&
        Objects.equals(this.networkId, inlineResponse20082LoadBalancerInstance.networkId) &&
        Objects.equals(this.subnetId, inlineResponse20082LoadBalancerInstance.subnetId) &&
        Objects.equals(this.externalPortId, inlineResponse20082LoadBalancerInstance.externalPortId) &&
        Objects.equals(this.status, inlineResponse20082LoadBalancerInstance.status) &&
        Objects.equals(this.vipStatus, inlineResponse20082LoadBalancerInstance.vipStatus);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, loadBalancer, instance, description, internalId, externalId, dateCreated, lastUpdated, active, sticky, sslEnabled, externalAddress, backendPort, vipType, vipAddress, vipHostname, vipProtocol, vipScheme, vipMode, vipName, vipPort, vipSticky, vipBalance, servicePort, sourceAddress, sslCert, sslMode, sslRedirectMode, vipShared, vipDirectAddress, serverName, poolName, removing, vipSource, extraConfig, serviceAccess, networkId, subnetId, externalPortId, status, vipStatus);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20082LoadBalancerInstance {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    loadBalancer: ").append(toIndentedString(loadBalancer)).append("\n");
    sb.append("    instance: ").append(toIndentedString(instance)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    internalId: ").append(toIndentedString(internalId)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    sticky: ").append(toIndentedString(sticky)).append("\n");
    sb.append("    sslEnabled: ").append(toIndentedString(sslEnabled)).append("\n");
    sb.append("    externalAddress: ").append(toIndentedString(externalAddress)).append("\n");
    sb.append("    backendPort: ").append(toIndentedString(backendPort)).append("\n");
    sb.append("    vipType: ").append(toIndentedString(vipType)).append("\n");
    sb.append("    vipAddress: ").append(toIndentedString(vipAddress)).append("\n");
    sb.append("    vipHostname: ").append(toIndentedString(vipHostname)).append("\n");
    sb.append("    vipProtocol: ").append(toIndentedString(vipProtocol)).append("\n");
    sb.append("    vipScheme: ").append(toIndentedString(vipScheme)).append("\n");
    sb.append("    vipMode: ").append(toIndentedString(vipMode)).append("\n");
    sb.append("    vipName: ").append(toIndentedString(vipName)).append("\n");
    sb.append("    vipPort: ").append(toIndentedString(vipPort)).append("\n");
    sb.append("    vipSticky: ").append(toIndentedString(vipSticky)).append("\n");
    sb.append("    vipBalance: ").append(toIndentedString(vipBalance)).append("\n");
    sb.append("    servicePort: ").append(toIndentedString(servicePort)).append("\n");
    sb.append("    sourceAddress: ").append(toIndentedString(sourceAddress)).append("\n");
    sb.append("    sslCert: ").append(toIndentedString(sslCert)).append("\n");
    sb.append("    sslMode: ").append(toIndentedString(sslMode)).append("\n");
    sb.append("    sslRedirectMode: ").append(toIndentedString(sslRedirectMode)).append("\n");
    sb.append("    vipShared: ").append(toIndentedString(vipShared)).append("\n");
    sb.append("    vipDirectAddress: ").append(toIndentedString(vipDirectAddress)).append("\n");
    sb.append("    serverName: ").append(toIndentedString(serverName)).append("\n");
    sb.append("    poolName: ").append(toIndentedString(poolName)).append("\n");
    sb.append("    removing: ").append(toIndentedString(removing)).append("\n");
    sb.append("    vipSource: ").append(toIndentedString(vipSource)).append("\n");
    sb.append("    extraConfig: ").append(toIndentedString(extraConfig)).append("\n");
    sb.append("    serviceAccess: ").append(toIndentedString(serviceAccess)).append("\n");
    sb.append("    networkId: ").append(toIndentedString(networkId)).append("\n");
    sb.append("    subnetId: ").append(toIndentedString(subnetId)).append("\n");
    sb.append("    externalPortId: ").append(toIndentedString(externalPortId)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    vipStatus: ").append(toIndentedString(vipStatus)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

