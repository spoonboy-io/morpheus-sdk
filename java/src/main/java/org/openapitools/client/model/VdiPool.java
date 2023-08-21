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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.openapitools.client.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.client.model.VdiPoolConfig;
import org.openapitools.client.model.VdiPoolOwner;
import org.threeten.bp.OffsetDateTime;

/**
 * VdiPool
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class VdiPool {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_MIN_IDLE = "minIdle";
  @SerializedName(SERIALIZED_NAME_MIN_IDLE)
  private Long minIdle;

  public static final String SERIALIZED_NAME_MAX_IDLE = "maxIdle";
  @SerializedName(SERIALIZED_NAME_MAX_IDLE)
  private Long maxIdle;

  public static final String SERIALIZED_NAME_INITIAL_POOL_SIZE = "initialPoolSize";
  @SerializedName(SERIALIZED_NAME_INITIAL_POOL_SIZE)
  private Long initialPoolSize;

  public static final String SERIALIZED_NAME_MAX_POOL_SIZE = "maxPoolSize";
  @SerializedName(SERIALIZED_NAME_MAX_POOL_SIZE)
  private Long maxPoolSize;

  public static final String SERIALIZED_NAME_ALLOCATION_TIMEOUT_MINUTES = "allocationTimeoutMinutes";
  @SerializedName(SERIALIZED_NAME_ALLOCATION_TIMEOUT_MINUTES)
  private Long allocationTimeoutMinutes;

  public static final String SERIALIZED_NAME_PERSISTENT_USER = "persistentUser";
  @SerializedName(SERIALIZED_NAME_PERSISTENT_USER)
  private Boolean persistentUser;

  public static final String SERIALIZED_NAME_RECYCLABLE = "recyclable";
  @SerializedName(SERIALIZED_NAME_RECYCLABLE)
  private Boolean recyclable;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_AUTO_CREATE_LOCAL_USER_ON_RESERVATION = "autoCreateLocalUserOnReservation";
  @SerializedName(SERIALIZED_NAME_AUTO_CREATE_LOCAL_USER_ON_RESERVATION)
  private Boolean autoCreateLocalUserOnReservation;

  public static final String SERIALIZED_NAME_ALLOW_HYPERVISOR_CONSOLE = "allowHypervisorConsole";
  @SerializedName(SERIALIZED_NAME_ALLOW_HYPERVISOR_CONSOLE)
  private Boolean allowHypervisorConsole;

  public static final String SERIALIZED_NAME_ALLOW_COPY = "allowCopy";
  @SerializedName(SERIALIZED_NAME_ALLOW_COPY)
  private Boolean allowCopy;

  public static final String SERIALIZED_NAME_ALLOW_PRINTER = "allowPrinter";
  @SerializedName(SERIALIZED_NAME_ALLOW_PRINTER)
  private Boolean allowPrinter;

  public static final String SERIALIZED_NAME_ALLOW_FILESHARE = "allowFileshare";
  @SerializedName(SERIALIZED_NAME_ALLOW_FILESHARE)
  private Boolean allowFileshare;

  public static final String SERIALIZED_NAME_GUEST_CONSOLE_JUMP_HOST = "guestConsoleJumpHost";
  @SerializedName(SERIALIZED_NAME_GUEST_CONSOLE_JUMP_HOST)
  private String guestConsoleJumpHost;

  public static final String SERIALIZED_NAME_GUEST_CONSOLE_JUMP_PORT = "guestConsoleJumpPort";
  @SerializedName(SERIALIZED_NAME_GUEST_CONSOLE_JUMP_PORT)
  private String guestConsoleJumpPort;

  public static final String SERIALIZED_NAME_GUEST_CONSOLE_JUMP_USERNAME = "guestConsoleJumpUsername";
  @SerializedName(SERIALIZED_NAME_GUEST_CONSOLE_JUMP_USERNAME)
  private String guestConsoleJumpUsername;

  public static final String SERIALIZED_NAME_GUEST_CONSOLE_JUMP_PASSWORD = "guestConsoleJumpPassword";
  @SerializedName(SERIALIZED_NAME_GUEST_CONSOLE_JUMP_PASSWORD)
  private String guestConsoleJumpPassword;

  public static final String SERIALIZED_NAME_GUEST_CONSOLE_JUMP_KEYPAIR = "guestConsoleJumpKeypair";
  @SerializedName(SERIALIZED_NAME_GUEST_CONSOLE_JUMP_KEYPAIR)
  private String guestConsoleJumpKeypair;

  public static final String SERIALIZED_NAME_GATEWAY = "gateway";
  @SerializedName(SERIALIZED_NAME_GATEWAY)
  private InlineResponse20082LoadBalancerInstanceSslCert gateway;

  public static final String SERIALIZED_NAME_ICON_PATH = "iconPath";
  @SerializedName(SERIALIZED_NAME_ICON_PATH)
  private String iconPath;

  public static final String SERIALIZED_NAME_LOGO = "logo";
  @SerializedName(SERIALIZED_NAME_LOGO)
  private String logo;

  public static final String SERIALIZED_NAME_APPS = "apps";
  @SerializedName(SERIALIZED_NAME_APPS)
  private List<InlineResponse20040AppDeployInstance> apps = null;

  public static final String SERIALIZED_NAME_OWNER = "owner";
  @SerializedName(SERIALIZED_NAME_OWNER)
  private VdiPoolOwner owner;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private VdiPoolConfig config;

  public static final String SERIALIZED_NAME_GROUP = "group";
  @SerializedName(SERIALIZED_NAME_GROUP)
  private InlineResponse20082LoadBalancerInstanceSslCert group;

  public static final String SERIALIZED_NAME_CLOUD = "cloud";
  @SerializedName(SERIALIZED_NAME_CLOUD)
  private InlineResponse20082LoadBalancerInstanceSslCert cloud;

  public static final String SERIALIZED_NAME_USED_COUNT = "usedCount";
  @SerializedName(SERIALIZED_NAME_USED_COUNT)
  private Long usedCount;

  public static final String SERIALIZED_NAME_RESERVED_COUNT = "reservedCount";
  @SerializedName(SERIALIZED_NAME_RESERVED_COUNT)
  private Long reservedCount;

  public static final String SERIALIZED_NAME_PREPARING_COUNT = "preparingCount";
  @SerializedName(SERIALIZED_NAME_PREPARING_COUNT)
  private Long preparingCount;

  public static final String SERIALIZED_NAME_IDLE_COUNT = "idleCount";
  @SerializedName(SERIALIZED_NAME_IDLE_COUNT)
  private Long idleCount;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;


  public VdiPool id(Long id) {
    
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


  public VdiPool name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public VdiPool description(String description) {
    
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


  public VdiPool minIdle(Long minIdle) {
    
    this.minIdle = minIdle;
    return this;
  }

   /**
   * Get minIdle
   * @return minIdle
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMinIdle() {
    return minIdle;
  }


  public void setMinIdle(Long minIdle) {
    this.minIdle = minIdle;
  }


  public VdiPool maxIdle(Long maxIdle) {
    
    this.maxIdle = maxIdle;
    return this;
  }

   /**
   * Get maxIdle
   * @return maxIdle
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxIdle() {
    return maxIdle;
  }


  public void setMaxIdle(Long maxIdle) {
    this.maxIdle = maxIdle;
  }


  public VdiPool initialPoolSize(Long initialPoolSize) {
    
    this.initialPoolSize = initialPoolSize;
    return this;
  }

   /**
   * Get initialPoolSize
   * @return initialPoolSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getInitialPoolSize() {
    return initialPoolSize;
  }


  public void setInitialPoolSize(Long initialPoolSize) {
    this.initialPoolSize = initialPoolSize;
  }


  public VdiPool maxPoolSize(Long maxPoolSize) {
    
    this.maxPoolSize = maxPoolSize;
    return this;
  }

   /**
   * Get maxPoolSize
   * @return maxPoolSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxPoolSize() {
    return maxPoolSize;
  }


  public void setMaxPoolSize(Long maxPoolSize) {
    this.maxPoolSize = maxPoolSize;
  }


  public VdiPool allocationTimeoutMinutes(Long allocationTimeoutMinutes) {
    
    this.allocationTimeoutMinutes = allocationTimeoutMinutes;
    return this;
  }

   /**
   * Get allocationTimeoutMinutes
   * @return allocationTimeoutMinutes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getAllocationTimeoutMinutes() {
    return allocationTimeoutMinutes;
  }


  public void setAllocationTimeoutMinutes(Long allocationTimeoutMinutes) {
    this.allocationTimeoutMinutes = allocationTimeoutMinutes;
  }


  public VdiPool persistentUser(Boolean persistentUser) {
    
    this.persistentUser = persistentUser;
    return this;
  }

   /**
   * Get persistentUser
   * @return persistentUser
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getPersistentUser() {
    return persistentUser;
  }


  public void setPersistentUser(Boolean persistentUser) {
    this.persistentUser = persistentUser;
  }


  public VdiPool recyclable(Boolean recyclable) {
    
    this.recyclable = recyclable;
    return this;
  }

   /**
   * Get recyclable
   * @return recyclable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRecyclable() {
    return recyclable;
  }


  public void setRecyclable(Boolean recyclable) {
    this.recyclable = recyclable;
  }


  public VdiPool enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Get enabled
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public VdiPool autoCreateLocalUserOnReservation(Boolean autoCreateLocalUserOnReservation) {
    
    this.autoCreateLocalUserOnReservation = autoCreateLocalUserOnReservation;
    return this;
  }

   /**
   * Get autoCreateLocalUserOnReservation
   * @return autoCreateLocalUserOnReservation
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAutoCreateLocalUserOnReservation() {
    return autoCreateLocalUserOnReservation;
  }


  public void setAutoCreateLocalUserOnReservation(Boolean autoCreateLocalUserOnReservation) {
    this.autoCreateLocalUserOnReservation = autoCreateLocalUserOnReservation;
  }


  public VdiPool allowHypervisorConsole(Boolean allowHypervisorConsole) {
    
    this.allowHypervisorConsole = allowHypervisorConsole;
    return this;
  }

   /**
   * Get allowHypervisorConsole
   * @return allowHypervisorConsole
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAllowHypervisorConsole() {
    return allowHypervisorConsole;
  }


  public void setAllowHypervisorConsole(Boolean allowHypervisorConsole) {
    this.allowHypervisorConsole = allowHypervisorConsole;
  }


  public VdiPool allowCopy(Boolean allowCopy) {
    
    this.allowCopy = allowCopy;
    return this;
  }

   /**
   * Get allowCopy
   * @return allowCopy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAllowCopy() {
    return allowCopy;
  }


  public void setAllowCopy(Boolean allowCopy) {
    this.allowCopy = allowCopy;
  }


  public VdiPool allowPrinter(Boolean allowPrinter) {
    
    this.allowPrinter = allowPrinter;
    return this;
  }

   /**
   * Get allowPrinter
   * @return allowPrinter
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAllowPrinter() {
    return allowPrinter;
  }


  public void setAllowPrinter(Boolean allowPrinter) {
    this.allowPrinter = allowPrinter;
  }


  public VdiPool allowFileshare(Boolean allowFileshare) {
    
    this.allowFileshare = allowFileshare;
    return this;
  }

   /**
   * Get allowFileshare
   * @return allowFileshare
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAllowFileshare() {
    return allowFileshare;
  }


  public void setAllowFileshare(Boolean allowFileshare) {
    this.allowFileshare = allowFileshare;
  }


  public VdiPool guestConsoleJumpHost(String guestConsoleJumpHost) {
    
    this.guestConsoleJumpHost = guestConsoleJumpHost;
    return this;
  }

   /**
   * Get guestConsoleJumpHost
   * @return guestConsoleJumpHost
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGuestConsoleJumpHost() {
    return guestConsoleJumpHost;
  }


  public void setGuestConsoleJumpHost(String guestConsoleJumpHost) {
    this.guestConsoleJumpHost = guestConsoleJumpHost;
  }


  public VdiPool guestConsoleJumpPort(String guestConsoleJumpPort) {
    
    this.guestConsoleJumpPort = guestConsoleJumpPort;
    return this;
  }

   /**
   * Get guestConsoleJumpPort
   * @return guestConsoleJumpPort
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGuestConsoleJumpPort() {
    return guestConsoleJumpPort;
  }


  public void setGuestConsoleJumpPort(String guestConsoleJumpPort) {
    this.guestConsoleJumpPort = guestConsoleJumpPort;
  }


  public VdiPool guestConsoleJumpUsername(String guestConsoleJumpUsername) {
    
    this.guestConsoleJumpUsername = guestConsoleJumpUsername;
    return this;
  }

   /**
   * Get guestConsoleJumpUsername
   * @return guestConsoleJumpUsername
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGuestConsoleJumpUsername() {
    return guestConsoleJumpUsername;
  }


  public void setGuestConsoleJumpUsername(String guestConsoleJumpUsername) {
    this.guestConsoleJumpUsername = guestConsoleJumpUsername;
  }


  public VdiPool guestConsoleJumpPassword(String guestConsoleJumpPassword) {
    
    this.guestConsoleJumpPassword = guestConsoleJumpPassword;
    return this;
  }

   /**
   * Get guestConsoleJumpPassword
   * @return guestConsoleJumpPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGuestConsoleJumpPassword() {
    return guestConsoleJumpPassword;
  }


  public void setGuestConsoleJumpPassword(String guestConsoleJumpPassword) {
    this.guestConsoleJumpPassword = guestConsoleJumpPassword;
  }


  public VdiPool guestConsoleJumpKeypair(String guestConsoleJumpKeypair) {
    
    this.guestConsoleJumpKeypair = guestConsoleJumpKeypair;
    return this;
  }

   /**
   * Get guestConsoleJumpKeypair
   * @return guestConsoleJumpKeypair
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGuestConsoleJumpKeypair() {
    return guestConsoleJumpKeypair;
  }


  public void setGuestConsoleJumpKeypair(String guestConsoleJumpKeypair) {
    this.guestConsoleJumpKeypair = guestConsoleJumpKeypair;
  }


  public VdiPool gateway(InlineResponse20082LoadBalancerInstanceSslCert gateway) {
    
    this.gateway = gateway;
    return this;
  }

   /**
   * Get gateway
   * @return gateway
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getGateway() {
    return gateway;
  }


  public void setGateway(InlineResponse20082LoadBalancerInstanceSslCert gateway) {
    this.gateway = gateway;
  }


  public VdiPool iconPath(String iconPath) {
    
    this.iconPath = iconPath;
    return this;
  }

   /**
   * Get iconPath
   * @return iconPath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIconPath() {
    return iconPath;
  }


  public void setIconPath(String iconPath) {
    this.iconPath = iconPath;
  }


  public VdiPool logo(String logo) {
    
    this.logo = logo;
    return this;
  }

   /**
   * Get logo
   * @return logo
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLogo() {
    return logo;
  }


  public void setLogo(String logo) {
    this.logo = logo;
  }


  public VdiPool apps(List<InlineResponse20040AppDeployInstance> apps) {
    
    this.apps = apps;
    return this;
  }

  public VdiPool addAppsItem(InlineResponse20040AppDeployInstance appsItem) {
    if (this.apps == null) {
      this.apps = new ArrayList<InlineResponse20040AppDeployInstance>();
    }
    this.apps.add(appsItem);
    return this;
  }

   /**
   * Get apps
   * @return apps
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<InlineResponse20040AppDeployInstance> getApps() {
    return apps;
  }


  public void setApps(List<InlineResponse20040AppDeployInstance> apps) {
    this.apps = apps;
  }


  public VdiPool owner(VdiPoolOwner owner) {
    
    this.owner = owner;
    return this;
  }

   /**
   * Get owner
   * @return owner
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public VdiPoolOwner getOwner() {
    return owner;
  }


  public void setOwner(VdiPoolOwner owner) {
    this.owner = owner;
  }


  public VdiPool config(VdiPoolConfig config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public VdiPoolConfig getConfig() {
    return config;
  }


  public void setConfig(VdiPoolConfig config) {
    this.config = config;
  }


  public VdiPool group(InlineResponse20082LoadBalancerInstanceSslCert group) {
    
    this.group = group;
    return this;
  }

   /**
   * Get group
   * @return group
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getGroup() {
    return group;
  }


  public void setGroup(InlineResponse20082LoadBalancerInstanceSslCert group) {
    this.group = group;
  }


  public VdiPool cloud(InlineResponse20082LoadBalancerInstanceSslCert cloud) {
    
    this.cloud = cloud;
    return this;
  }

   /**
   * Get cloud
   * @return cloud
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getCloud() {
    return cloud;
  }


  public void setCloud(InlineResponse20082LoadBalancerInstanceSslCert cloud) {
    this.cloud = cloud;
  }


  public VdiPool usedCount(Long usedCount) {
    
    this.usedCount = usedCount;
    return this;
  }

   /**
   * Get usedCount
   * @return usedCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getUsedCount() {
    return usedCount;
  }


  public void setUsedCount(Long usedCount) {
    this.usedCount = usedCount;
  }


  public VdiPool reservedCount(Long reservedCount) {
    
    this.reservedCount = reservedCount;
    return this;
  }

   /**
   * Get reservedCount
   * @return reservedCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getReservedCount() {
    return reservedCount;
  }


  public void setReservedCount(Long reservedCount) {
    this.reservedCount = reservedCount;
  }


  public VdiPool preparingCount(Long preparingCount) {
    
    this.preparingCount = preparingCount;
    return this;
  }

   /**
   * Get preparingCount
   * @return preparingCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getPreparingCount() {
    return preparingCount;
  }


  public void setPreparingCount(Long preparingCount) {
    this.preparingCount = preparingCount;
  }


  public VdiPool idleCount(Long idleCount) {
    
    this.idleCount = idleCount;
    return this;
  }

   /**
   * Get idleCount
   * @return idleCount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getIdleCount() {
    return idleCount;
  }


  public void setIdleCount(Long idleCount) {
    this.idleCount = idleCount;
  }


  public VdiPool status(String status) {
    
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


  public VdiPool dateCreated(OffsetDateTime dateCreated) {
    
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


  public VdiPool lastUpdated(OffsetDateTime lastUpdated) {
    
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    VdiPool vdiPool = (VdiPool) o;
    return Objects.equals(this.id, vdiPool.id) &&
        Objects.equals(this.name, vdiPool.name) &&
        Objects.equals(this.description, vdiPool.description) &&
        Objects.equals(this.minIdle, vdiPool.minIdle) &&
        Objects.equals(this.maxIdle, vdiPool.maxIdle) &&
        Objects.equals(this.initialPoolSize, vdiPool.initialPoolSize) &&
        Objects.equals(this.maxPoolSize, vdiPool.maxPoolSize) &&
        Objects.equals(this.allocationTimeoutMinutes, vdiPool.allocationTimeoutMinutes) &&
        Objects.equals(this.persistentUser, vdiPool.persistentUser) &&
        Objects.equals(this.recyclable, vdiPool.recyclable) &&
        Objects.equals(this.enabled, vdiPool.enabled) &&
        Objects.equals(this.autoCreateLocalUserOnReservation, vdiPool.autoCreateLocalUserOnReservation) &&
        Objects.equals(this.allowHypervisorConsole, vdiPool.allowHypervisorConsole) &&
        Objects.equals(this.allowCopy, vdiPool.allowCopy) &&
        Objects.equals(this.allowPrinter, vdiPool.allowPrinter) &&
        Objects.equals(this.allowFileshare, vdiPool.allowFileshare) &&
        Objects.equals(this.guestConsoleJumpHost, vdiPool.guestConsoleJumpHost) &&
        Objects.equals(this.guestConsoleJumpPort, vdiPool.guestConsoleJumpPort) &&
        Objects.equals(this.guestConsoleJumpUsername, vdiPool.guestConsoleJumpUsername) &&
        Objects.equals(this.guestConsoleJumpPassword, vdiPool.guestConsoleJumpPassword) &&
        Objects.equals(this.guestConsoleJumpKeypair, vdiPool.guestConsoleJumpKeypair) &&
        Objects.equals(this.gateway, vdiPool.gateway) &&
        Objects.equals(this.iconPath, vdiPool.iconPath) &&
        Objects.equals(this.logo, vdiPool.logo) &&
        Objects.equals(this.apps, vdiPool.apps) &&
        Objects.equals(this.owner, vdiPool.owner) &&
        Objects.equals(this.config, vdiPool.config) &&
        Objects.equals(this.group, vdiPool.group) &&
        Objects.equals(this.cloud, vdiPool.cloud) &&
        Objects.equals(this.usedCount, vdiPool.usedCount) &&
        Objects.equals(this.reservedCount, vdiPool.reservedCount) &&
        Objects.equals(this.preparingCount, vdiPool.preparingCount) &&
        Objects.equals(this.idleCount, vdiPool.idleCount) &&
        Objects.equals(this.status, vdiPool.status) &&
        Objects.equals(this.dateCreated, vdiPool.dateCreated) &&
        Objects.equals(this.lastUpdated, vdiPool.lastUpdated);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, description, minIdle, maxIdle, initialPoolSize, maxPoolSize, allocationTimeoutMinutes, persistentUser, recyclable, enabled, autoCreateLocalUserOnReservation, allowHypervisorConsole, allowCopy, allowPrinter, allowFileshare, guestConsoleJumpHost, guestConsoleJumpPort, guestConsoleJumpUsername, guestConsoleJumpPassword, guestConsoleJumpKeypair, gateway, iconPath, logo, apps, owner, config, group, cloud, usedCount, reservedCount, preparingCount, idleCount, status, dateCreated, lastUpdated);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class VdiPool {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    minIdle: ").append(toIndentedString(minIdle)).append("\n");
    sb.append("    maxIdle: ").append(toIndentedString(maxIdle)).append("\n");
    sb.append("    initialPoolSize: ").append(toIndentedString(initialPoolSize)).append("\n");
    sb.append("    maxPoolSize: ").append(toIndentedString(maxPoolSize)).append("\n");
    sb.append("    allocationTimeoutMinutes: ").append(toIndentedString(allocationTimeoutMinutes)).append("\n");
    sb.append("    persistentUser: ").append(toIndentedString(persistentUser)).append("\n");
    sb.append("    recyclable: ").append(toIndentedString(recyclable)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    autoCreateLocalUserOnReservation: ").append(toIndentedString(autoCreateLocalUserOnReservation)).append("\n");
    sb.append("    allowHypervisorConsole: ").append(toIndentedString(allowHypervisorConsole)).append("\n");
    sb.append("    allowCopy: ").append(toIndentedString(allowCopy)).append("\n");
    sb.append("    allowPrinter: ").append(toIndentedString(allowPrinter)).append("\n");
    sb.append("    allowFileshare: ").append(toIndentedString(allowFileshare)).append("\n");
    sb.append("    guestConsoleJumpHost: ").append(toIndentedString(guestConsoleJumpHost)).append("\n");
    sb.append("    guestConsoleJumpPort: ").append(toIndentedString(guestConsoleJumpPort)).append("\n");
    sb.append("    guestConsoleJumpUsername: ").append(toIndentedString(guestConsoleJumpUsername)).append("\n");
    sb.append("    guestConsoleJumpPassword: ").append(toIndentedString(guestConsoleJumpPassword)).append("\n");
    sb.append("    guestConsoleJumpKeypair: ").append(toIndentedString(guestConsoleJumpKeypair)).append("\n");
    sb.append("    gateway: ").append(toIndentedString(gateway)).append("\n");
    sb.append("    iconPath: ").append(toIndentedString(iconPath)).append("\n");
    sb.append("    logo: ").append(toIndentedString(logo)).append("\n");
    sb.append("    apps: ").append(toIndentedString(apps)).append("\n");
    sb.append("    owner: ").append(toIndentedString(owner)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    group: ").append(toIndentedString(group)).append("\n");
    sb.append("    cloud: ").append(toIndentedString(cloud)).append("\n");
    sb.append("    usedCount: ").append(toIndentedString(usedCount)).append("\n");
    sb.append("    reservedCount: ").append(toIndentedString(reservedCount)).append("\n");
    sb.append("    preparingCount: ").append(toIndentedString(preparingCount)).append("\n");
    sb.append("    idleCount: ").append(toIndentedString(idleCount)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
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
