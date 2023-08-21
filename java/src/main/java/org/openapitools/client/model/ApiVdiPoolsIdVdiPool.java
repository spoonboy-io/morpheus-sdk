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
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.ApiVdiPoolsIdVdiPoolConfig;

/**
 * ApiVdiPoolsIdVdiPool
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiVdiPoolsIdVdiPool {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_OWNER = "owner";
  @SerializedName(SERIALIZED_NAME_OWNER)
  private Long owner;

  public static final String SERIALIZED_NAME_MIN_IDLE = "minIdle";
  @SerializedName(SERIALIZED_NAME_MIN_IDLE)
  private BigDecimal minIdle;

  public static final String SERIALIZED_NAME_INITIAL_POOL_SIZE = "initialPoolSize";
  @SerializedName(SERIALIZED_NAME_INITIAL_POOL_SIZE)
  private BigDecimal initialPoolSize;

  public static final String SERIALIZED_NAME_MAX_IDLE = "maxIdle";
  @SerializedName(SERIALIZED_NAME_MAX_IDLE)
  private BigDecimal maxIdle;

  public static final String SERIALIZED_NAME_MAX_POOL_SIZE = "maxPoolSize";
  @SerializedName(SERIALIZED_NAME_MAX_POOL_SIZE)
  private BigDecimal maxPoolSize;

  public static final String SERIALIZED_NAME_ALLOCATION_TIMEOUT_MINUTES = "allocationTimeoutMinutes";
  @SerializedName(SERIALIZED_NAME_ALLOCATION_TIMEOUT_MINUTES)
  private BigDecimal allocationTimeoutMinutes;

  public static final String SERIALIZED_NAME_PERSISTENT_USER = "persistentUser";
  @SerializedName(SERIALIZED_NAME_PERSISTENT_USER)
  private Boolean persistentUser = false;

  public static final String SERIALIZED_NAME_RECYCLABLE = "recyclable";
  @SerializedName(SERIALIZED_NAME_RECYCLABLE)
  private Boolean recyclable = false;

  public static final String SERIALIZED_NAME_ALLOW_COPY = "allowCopy";
  @SerializedName(SERIALIZED_NAME_ALLOW_COPY)
  private Boolean allowCopy = false;

  public static final String SERIALIZED_NAME_ALLOW_PRINTER = "allowPrinter";
  @SerializedName(SERIALIZED_NAME_ALLOW_PRINTER)
  private Boolean allowPrinter = false;

  public static final String SERIALIZED_NAME_ALLOW_FILESHARE = "allowFileshare";
  @SerializedName(SERIALIZED_NAME_ALLOW_FILESHARE)
  private Boolean allowFileshare = false;

  public static final String SERIALIZED_NAME_ALLOW_HYPERVISOR_CONSOLE = "allowHypervisorConsole";
  @SerializedName(SERIALIZED_NAME_ALLOW_HYPERVISOR_CONSOLE)
  private Boolean allowHypervisorConsole = false;

  public static final String SERIALIZED_NAME_AUTO_CREATE_LOCAL_USER_ON_RESERVATION = "autoCreateLocalUserOnReservation";
  @SerializedName(SERIALIZED_NAME_AUTO_CREATE_LOCAL_USER_ON_RESERVATION)
  private Boolean autoCreateLocalUserOnReservation = false;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled = true;

  public static final String SERIALIZED_NAME_ICON_PATH = "iconPath";
  @SerializedName(SERIALIZED_NAME_ICON_PATH)
  private String iconPath;

  public static final String SERIALIZED_NAME_APPS = "apps";
  @SerializedName(SERIALIZED_NAME_APPS)
  private List<Long> apps = null;

  public static final String SERIALIZED_NAME_GATEWAY = "gateway";
  @SerializedName(SERIALIZED_NAME_GATEWAY)
  private Long gateway;

  public static final String SERIALIZED_NAME_INSTANCE_CONFIG = "instanceConfig";
  @SerializedName(SERIALIZED_NAME_INSTANCE_CONFIG)
  private String instanceConfig;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private ApiVdiPoolsIdVdiPoolConfig config;

  public static final String SERIALIZED_NAME_GUEST_CONSOLE_JUMP_HOST = "guestConsoleJumpHost";
  @SerializedName(SERIALIZED_NAME_GUEST_CONSOLE_JUMP_HOST)
  private String guestConsoleJumpHost;

  public static final String SERIALIZED_NAME_GUEST_CONSOLE_JUMP_PORT = "guestConsoleJumpPort";
  @SerializedName(SERIALIZED_NAME_GUEST_CONSOLE_JUMP_PORT)
  private Long guestConsoleJumpPort;

  public static final String SERIALIZED_NAME_GUEST_CONSOLE_JUMP_USERNAME = "guestConsoleJumpUsername";
  @SerializedName(SERIALIZED_NAME_GUEST_CONSOLE_JUMP_USERNAME)
  private String guestConsoleJumpUsername;

  public static final String SERIALIZED_NAME_GUEST_CONSOLE_JUMP_PASSWORD = "guestConsoleJumpPassword";
  @SerializedName(SERIALIZED_NAME_GUEST_CONSOLE_JUMP_PASSWORD)
  private String guestConsoleJumpPassword;

  public static final String SERIALIZED_NAME_GUEST_CONSOLE_JUMP_KEYPAIR = "guestConsoleJumpKeypair";
  @SerializedName(SERIALIZED_NAME_GUEST_CONSOLE_JUMP_KEYPAIR)
  private Long guestConsoleJumpKeypair;


  public ApiVdiPoolsIdVdiPool name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Virtual Desktop name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Virtual Desktop name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ApiVdiPoolsIdVdiPool description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Virtual Desktop description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Virtual Desktop description")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ApiVdiPoolsIdVdiPool owner(Long owner) {
    
    this.owner = owner;
    return this;
  }

   /**
   * Owner (User) ID
   * @return owner
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "21", value = "Owner (User) ID")

  public Long getOwner() {
    return owner;
  }


  public void setOwner(Long owner) {
    this.owner = owner;
  }


  public ApiVdiPoolsIdVdiPool minIdle(BigDecimal minIdle) {
    
    this.minIdle = minIdle;
    return this;
  }

   /**
   * Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby. 
   * @return minIdle
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "5", value = "Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby. ")

  public BigDecimal getMinIdle() {
    return minIdle;
  }


  public void setMinIdle(BigDecimal minIdle) {
    this.minIdle = minIdle;
  }


  public ApiVdiPoolsIdVdiPool initialPoolSize(BigDecimal initialPoolSize) {
    
    this.initialPoolSize = initialPoolSize;
    return this;
  }

   /**
   * The initial size of the pool to be allocated on creation
   * @return initialPoolSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "10", value = "The initial size of the pool to be allocated on creation")

  public BigDecimal getInitialPoolSize() {
    return initialPoolSize;
  }


  public void setInitialPoolSize(BigDecimal initialPoolSize) {
    this.initialPoolSize = initialPoolSize;
  }


  public ApiVdiPoolsIdVdiPool maxIdle(BigDecimal maxIdle) {
    
    this.maxIdle = maxIdle;
    return this;
  }

   /**
   * Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances 
   * @return maxIdle
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "2", value = "Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances ")

  public BigDecimal getMaxIdle() {
    return maxIdle;
  }


  public void setMaxIdle(BigDecimal maxIdle) {
    this.maxIdle = maxIdle;
  }


  public ApiVdiPoolsIdVdiPool maxPoolSize(BigDecimal maxPoolSize) {
    
    this.maxPoolSize = maxPoolSize;
    return this;
  }

   /**
   * Max limit on number of allocations and instances within the pool. 
   * @return maxPoolSize
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "50", value = "Max limit on number of allocations and instances within the pool. ")

  public BigDecimal getMaxPoolSize() {
    return maxPoolSize;
  }


  public void setMaxPoolSize(BigDecimal maxPoolSize) {
    this.maxPoolSize = maxPoolSize;
  }


  public ApiVdiPoolsIdVdiPool allocationTimeoutMinutes(BigDecimal allocationTimeoutMinutes) {
    
    this.allocationTimeoutMinutes = allocationTimeoutMinutes;
    return this;
  }

   /**
   * Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence. 
   * @return allocationTimeoutMinutes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "20", value = "Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence. ")

  public BigDecimal getAllocationTimeoutMinutes() {
    return allocationTimeoutMinutes;
  }


  public void setAllocationTimeoutMinutes(BigDecimal allocationTimeoutMinutes) {
    this.allocationTimeoutMinutes = allocationTimeoutMinutes;
  }


  public ApiVdiPoolsIdVdiPool persistentUser(Boolean persistentUser) {
    
    this.persistentUser = persistentUser;
    return this;
  }

   /**
   * Persistent Desktop Pool
   * @return persistentUser
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Persistent Desktop Pool")

  public Boolean getPersistentUser() {
    return persistentUser;
  }


  public void setPersistentUser(Boolean persistentUser) {
    this.persistentUser = persistentUser;
  }


  public ApiVdiPoolsIdVdiPool recyclable(Boolean recyclable) {
    
    this.recyclable = recyclable;
    return this;
  }

   /**
   * Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD)
   * @return recyclable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD)")

  public Boolean getRecyclable() {
    return recyclable;
  }


  public void setRecyclable(Boolean recyclable) {
    this.recyclable = recyclable;
  }


  public ApiVdiPoolsIdVdiPool allowCopy(Boolean allowCopy) {
    
    this.allowCopy = allowCopy;
    return this;
  }

   /**
   * Allow copy from desktop
   * @return allowCopy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Allow copy from desktop")

  public Boolean getAllowCopy() {
    return allowCopy;
  }


  public void setAllowCopy(Boolean allowCopy) {
    this.allowCopy = allowCopy;
  }


  public ApiVdiPoolsIdVdiPool allowPrinter(Boolean allowPrinter) {
    
    this.allowPrinter = allowPrinter;
    return this;
  }

   /**
   * Allow local printers from Desktop
   * @return allowPrinter
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Allow local printers from Desktop")

  public Boolean getAllowPrinter() {
    return allowPrinter;
  }


  public void setAllowPrinter(Boolean allowPrinter) {
    this.allowPrinter = allowPrinter;
  }


  public ApiVdiPoolsIdVdiPool allowFileshare(Boolean allowFileshare) {
    
    this.allowFileshare = allowFileshare;
    return this;
  }

   /**
   * Allow File Share
   * @return allowFileshare
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Allow File Share")

  public Boolean getAllowFileshare() {
    return allowFileshare;
  }


  public void setAllowFileshare(Boolean allowFileshare) {
    this.allowFileshare = allowFileshare;
  }


  public ApiVdiPoolsIdVdiPool allowHypervisorConsole(Boolean allowHypervisorConsole) {
    
    this.allowHypervisorConsole = allowHypervisorConsole;
    return this;
  }

   /**
   * Allow hypervisor console
   * @return allowHypervisorConsole
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Allow hypervisor console")

  public Boolean getAllowHypervisorConsole() {
    return allowHypervisorConsole;
  }


  public void setAllowHypervisorConsole(Boolean allowHypervisorConsole) {
    this.allowHypervisorConsole = allowHypervisorConsole;
  }


  public ApiVdiPoolsIdVdiPool autoCreateLocalUserOnReservation(Boolean autoCreateLocalUserOnReservation) {
    
    this.autoCreateLocalUserOnReservation = autoCreateLocalUserOnReservation;
    return this;
  }

   /**
   * Auto create local user upon reservation
   * @return autoCreateLocalUserOnReservation
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Auto create local user upon reservation")

  public Boolean getAutoCreateLocalUserOnReservation() {
    return autoCreateLocalUserOnReservation;
  }


  public void setAutoCreateLocalUserOnReservation(Boolean autoCreateLocalUserOnReservation) {
    this.autoCreateLocalUserOnReservation = autoCreateLocalUserOnReservation;
  }


  public ApiVdiPoolsIdVdiPool enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Can be used to enable or disable the VDI pool
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to enable or disable the VDI pool")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public ApiVdiPoolsIdVdiPool iconPath(String iconPath) {
    
    this.iconPath = iconPath;
    return this;
  }

   /**
   * The relative location of an icon image
   * @return iconPath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "/assets/containers-png/windows.png", value = "The relative location of an icon image")

  public String getIconPath() {
    return iconPath;
  }


  public void setIconPath(String iconPath) {
    this.iconPath = iconPath;
  }


  public ApiVdiPoolsIdVdiPool apps(List<Long> apps) {
    
    this.apps = apps;
    return this;
  }

  public ApiVdiPoolsIdVdiPool addAppsItem(Long appsItem) {
    if (this.apps == null) {
      this.apps = new ArrayList<Long>();
    }
    this.apps.add(appsItem);
    return this;
  }

   /**
   * Array of VDI App IDs
   * @return apps
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of VDI App IDs")

  public List<Long> getApps() {
    return apps;
  }


  public void setApps(List<Long> apps) {
    this.apps = apps;
  }


  public ApiVdiPoolsIdVdiPool gateway(Long gateway) {
    
    this.gateway = gateway;
    return this;
  }

   /**
   * VDI Gateway ID
   * @return gateway
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "VDI Gateway ID")

  public Long getGateway() {
    return gateway;
  }


  public void setGateway(Long gateway) {
    this.gateway = gateway;
  }


  public ApiVdiPoolsIdVdiPool instanceConfig(String instanceConfig) {
    
    this.instanceConfig = instanceConfig;
    return this;
  }

   /**
   * Instance Config JSON. Passing as a string will preserve property order.  See &#x60;config&#x60; object for required values.
   * @return instanceConfig
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Instance Config JSON. Passing as a string will preserve property order.  See `config` object for required values.")

  public String getInstanceConfig() {
    return instanceConfig;
  }


  public void setInstanceConfig(String instanceConfig) {
    this.instanceConfig = instanceConfig;
  }


  public ApiVdiPoolsIdVdiPool config(ApiVdiPoolsIdVdiPoolConfig config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiVdiPoolsIdVdiPoolConfig getConfig() {
    return config;
  }


  public void setConfig(ApiVdiPoolsIdVdiPoolConfig config) {
    this.config = config;
  }


  public ApiVdiPoolsIdVdiPool guestConsoleJumpHost(String guestConsoleJumpHost) {
    
    this.guestConsoleJumpHost = guestConsoleJumpHost;
    return this;
  }

   /**
   * Guest Console Jump Host
   * @return guestConsoleJumpHost
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Guest Console Jump Host")

  public String getGuestConsoleJumpHost() {
    return guestConsoleJumpHost;
  }


  public void setGuestConsoleJumpHost(String guestConsoleJumpHost) {
    this.guestConsoleJumpHost = guestConsoleJumpHost;
  }


  public ApiVdiPoolsIdVdiPool guestConsoleJumpPort(Long guestConsoleJumpPort) {
    
    this.guestConsoleJumpPort = guestConsoleJumpPort;
    return this;
  }

   /**
   * Guest Console Jump Port
   * @return guestConsoleJumpPort
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Guest Console Jump Port")

  public Long getGuestConsoleJumpPort() {
    return guestConsoleJumpPort;
  }


  public void setGuestConsoleJumpPort(Long guestConsoleJumpPort) {
    this.guestConsoleJumpPort = guestConsoleJumpPort;
  }


  public ApiVdiPoolsIdVdiPool guestConsoleJumpUsername(String guestConsoleJumpUsername) {
    
    this.guestConsoleJumpUsername = guestConsoleJumpUsername;
    return this;
  }

   /**
   * Guest Console Jump Username
   * @return guestConsoleJumpUsername
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Guest Console Jump Username")

  public String getGuestConsoleJumpUsername() {
    return guestConsoleJumpUsername;
  }


  public void setGuestConsoleJumpUsername(String guestConsoleJumpUsername) {
    this.guestConsoleJumpUsername = guestConsoleJumpUsername;
  }


  public ApiVdiPoolsIdVdiPool guestConsoleJumpPassword(String guestConsoleJumpPassword) {
    
    this.guestConsoleJumpPassword = guestConsoleJumpPassword;
    return this;
  }

   /**
   * Guest Console Jump Password
   * @return guestConsoleJumpPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Guest Console Jump Password")

  public String getGuestConsoleJumpPassword() {
    return guestConsoleJumpPassword;
  }


  public void setGuestConsoleJumpPassword(String guestConsoleJumpPassword) {
    this.guestConsoleJumpPassword = guestConsoleJumpPassword;
  }


  public ApiVdiPoolsIdVdiPool guestConsoleJumpKeypair(Long guestConsoleJumpKeypair) {
    
    this.guestConsoleJumpKeypair = guestConsoleJumpKeypair;
    return this;
  }

   /**
   * Guest Console Jump Key Pair. see &#x60;Key Pair&#x60;
   * @return guestConsoleJumpKeypair
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Guest Console Jump Key Pair. see `Key Pair`")

  public Long getGuestConsoleJumpKeypair() {
    return guestConsoleJumpKeypair;
  }


  public void setGuestConsoleJumpKeypair(Long guestConsoleJumpKeypair) {
    this.guestConsoleJumpKeypair = guestConsoleJumpKeypair;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiVdiPoolsIdVdiPool apiVdiPoolsIdVdiPool = (ApiVdiPoolsIdVdiPool) o;
    return Objects.equals(this.name, apiVdiPoolsIdVdiPool.name) &&
        Objects.equals(this.description, apiVdiPoolsIdVdiPool.description) &&
        Objects.equals(this.owner, apiVdiPoolsIdVdiPool.owner) &&
        Objects.equals(this.minIdle, apiVdiPoolsIdVdiPool.minIdle) &&
        Objects.equals(this.initialPoolSize, apiVdiPoolsIdVdiPool.initialPoolSize) &&
        Objects.equals(this.maxIdle, apiVdiPoolsIdVdiPool.maxIdle) &&
        Objects.equals(this.maxPoolSize, apiVdiPoolsIdVdiPool.maxPoolSize) &&
        Objects.equals(this.allocationTimeoutMinutes, apiVdiPoolsIdVdiPool.allocationTimeoutMinutes) &&
        Objects.equals(this.persistentUser, apiVdiPoolsIdVdiPool.persistentUser) &&
        Objects.equals(this.recyclable, apiVdiPoolsIdVdiPool.recyclable) &&
        Objects.equals(this.allowCopy, apiVdiPoolsIdVdiPool.allowCopy) &&
        Objects.equals(this.allowPrinter, apiVdiPoolsIdVdiPool.allowPrinter) &&
        Objects.equals(this.allowFileshare, apiVdiPoolsIdVdiPool.allowFileshare) &&
        Objects.equals(this.allowHypervisorConsole, apiVdiPoolsIdVdiPool.allowHypervisorConsole) &&
        Objects.equals(this.autoCreateLocalUserOnReservation, apiVdiPoolsIdVdiPool.autoCreateLocalUserOnReservation) &&
        Objects.equals(this.enabled, apiVdiPoolsIdVdiPool.enabled) &&
        Objects.equals(this.iconPath, apiVdiPoolsIdVdiPool.iconPath) &&
        Objects.equals(this.apps, apiVdiPoolsIdVdiPool.apps) &&
        Objects.equals(this.gateway, apiVdiPoolsIdVdiPool.gateway) &&
        Objects.equals(this.instanceConfig, apiVdiPoolsIdVdiPool.instanceConfig) &&
        Objects.equals(this.config, apiVdiPoolsIdVdiPool.config) &&
        Objects.equals(this.guestConsoleJumpHost, apiVdiPoolsIdVdiPool.guestConsoleJumpHost) &&
        Objects.equals(this.guestConsoleJumpPort, apiVdiPoolsIdVdiPool.guestConsoleJumpPort) &&
        Objects.equals(this.guestConsoleJumpUsername, apiVdiPoolsIdVdiPool.guestConsoleJumpUsername) &&
        Objects.equals(this.guestConsoleJumpPassword, apiVdiPoolsIdVdiPool.guestConsoleJumpPassword) &&
        Objects.equals(this.guestConsoleJumpKeypair, apiVdiPoolsIdVdiPool.guestConsoleJumpKeypair);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, owner, minIdle, initialPoolSize, maxIdle, maxPoolSize, allocationTimeoutMinutes, persistentUser, recyclable, allowCopy, allowPrinter, allowFileshare, allowHypervisorConsole, autoCreateLocalUserOnReservation, enabled, iconPath, apps, gateway, instanceConfig, config, guestConsoleJumpHost, guestConsoleJumpPort, guestConsoleJumpUsername, guestConsoleJumpPassword, guestConsoleJumpKeypair);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiVdiPoolsIdVdiPool {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    owner: ").append(toIndentedString(owner)).append("\n");
    sb.append("    minIdle: ").append(toIndentedString(minIdle)).append("\n");
    sb.append("    initialPoolSize: ").append(toIndentedString(initialPoolSize)).append("\n");
    sb.append("    maxIdle: ").append(toIndentedString(maxIdle)).append("\n");
    sb.append("    maxPoolSize: ").append(toIndentedString(maxPoolSize)).append("\n");
    sb.append("    allocationTimeoutMinutes: ").append(toIndentedString(allocationTimeoutMinutes)).append("\n");
    sb.append("    persistentUser: ").append(toIndentedString(persistentUser)).append("\n");
    sb.append("    recyclable: ").append(toIndentedString(recyclable)).append("\n");
    sb.append("    allowCopy: ").append(toIndentedString(allowCopy)).append("\n");
    sb.append("    allowPrinter: ").append(toIndentedString(allowPrinter)).append("\n");
    sb.append("    allowFileshare: ").append(toIndentedString(allowFileshare)).append("\n");
    sb.append("    allowHypervisorConsole: ").append(toIndentedString(allowHypervisorConsole)).append("\n");
    sb.append("    autoCreateLocalUserOnReservation: ").append(toIndentedString(autoCreateLocalUserOnReservation)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    iconPath: ").append(toIndentedString(iconPath)).append("\n");
    sb.append("    apps: ").append(toIndentedString(apps)).append("\n");
    sb.append("    gateway: ").append(toIndentedString(gateway)).append("\n");
    sb.append("    instanceConfig: ").append(toIndentedString(instanceConfig)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    guestConsoleJumpHost: ").append(toIndentedString(guestConsoleJumpHost)).append("\n");
    sb.append("    guestConsoleJumpPort: ").append(toIndentedString(guestConsoleJumpPort)).append("\n");
    sb.append("    guestConsoleJumpUsername: ").append(toIndentedString(guestConsoleJumpUsername)).append("\n");
    sb.append("    guestConsoleJumpPassword: ").append(toIndentedString(guestConsoleJumpPassword)).append("\n");
    sb.append("    guestConsoleJumpKeypair: ").append(toIndentedString(guestConsoleJumpKeypair)).append("\n");
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
