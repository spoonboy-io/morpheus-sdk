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
import org.openapitools.client.model.NetworkPoolServerCreateBluecatCredential;
import org.openapitools.client.model.NetworkPoolServerCreateInfobloxConfig;

/**
 * NetworkPoolServerUpdateInfoblox
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkPoolServerUpdateInfoblox {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled = true;

  public static final String SERIALIZED_NAME_SERVICE_URL = "serviceUrl";
  @SerializedName(SERIALIZED_NAME_SERVICE_URL)
  private String serviceUrl;

  public static final String SERIALIZED_NAME_SERVICE_USERNAME = "serviceUsername";
  @SerializedName(SERIALIZED_NAME_SERVICE_USERNAME)
  private String serviceUsername;

  public static final String SERIALIZED_NAME_SERVICE_PASSWORD = "servicePassword";
  @SerializedName(SERIALIZED_NAME_SERVICE_PASSWORD)
  private String servicePassword;

  public static final String SERIALIZED_NAME_SERVICE_THROTTLE_RATE = "serviceThrottleRate";
  @SerializedName(SERIALIZED_NAME_SERVICE_THROTTLE_RATE)
  private Long serviceThrottleRate = 0l;

  public static final String SERIALIZED_NAME_IGNORE_SSL = "ignoreSsl";
  @SerializedName(SERIALIZED_NAME_IGNORE_SSL)
  private Boolean ignoreSsl;

  public static final String SERIALIZED_NAME_NETWORK_FILTER = "networkFilter";
  @SerializedName(SERIALIZED_NAME_NETWORK_FILTER)
  private String networkFilter;

  public static final String SERIALIZED_NAME_ZONE_FILTER = "zoneFilter";
  @SerializedName(SERIALIZED_NAME_ZONE_FILTER)
  private String zoneFilter;

  public static final String SERIALIZED_NAME_TENANT_MATCH = "tenantMatch";
  @SerializedName(SERIALIZED_NAME_TENANT_MATCH)
  private String tenantMatch;

  /**
   * IP Mode
   */
  @JsonAdapter(ServiceModeEnum.Adapter.class)
  public enum ServiceModeEnum {
    STATIC("static"),
    
    DHCP("dhcp");

    private String value;

    ServiceModeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static ServiceModeEnum fromValue(String value) {
      for (ServiceModeEnum b : ServiceModeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      return null;
    }

    public static class Adapter extends TypeAdapter<ServiceModeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final ServiceModeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public ServiceModeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return ServiceModeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_SERVICE_MODE = "serviceMode";
  @SerializedName(SERIALIZED_NAME_SERVICE_MODE)
  private ServiceModeEnum serviceMode = ServiceModeEnum.STATIC;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private NetworkPoolServerCreateInfobloxConfig config;

  public static final String SERIALIZED_NAME_CREDENTIAL = "credential";
  @SerializedName(SERIALIZED_NAME_CREDENTIAL)
  private NetworkPoolServerCreateBluecatCredential credential;


  public NetworkPoolServerUpdateInfoblox name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public NetworkPoolServerUpdateInfoblox enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Can be used to enable / disable the network pool server.
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to enable / disable the network pool server.")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public NetworkPoolServerUpdateInfoblox serviceUrl(String serviceUrl) {
    
    this.serviceUrl = serviceUrl;
    return this;
  }

   /**
   * URL
   * @return serviceUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "https://infoblox-server/wapi/v2.9", value = "URL")

  public String getServiceUrl() {
    return serviceUrl;
  }


  public void setServiceUrl(String serviceUrl) {
    this.serviceUrl = serviceUrl;
  }


  public NetworkPoolServerUpdateInfoblox serviceUsername(String serviceUsername) {
    
    this.serviceUsername = serviceUsername;
    return this;
  }

   /**
   * Username
   * @return serviceUsername
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Username")

  public String getServiceUsername() {
    return serviceUsername;
  }


  public void setServiceUsername(String serviceUsername) {
    this.serviceUsername = serviceUsername;
  }


  public NetworkPoolServerUpdateInfoblox servicePassword(String servicePassword) {
    
    this.servicePassword = servicePassword;
    return this;
  }

   /**
   * Password
   * @return servicePassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Password")

  public String getServicePassword() {
    return servicePassword;
  }


  public void setServicePassword(String servicePassword) {
    this.servicePassword = servicePassword;
  }


  public NetworkPoolServerUpdateInfoblox serviceThrottleRate(Long serviceThrottleRate) {
    
    this.serviceThrottleRate = serviceThrottleRate;
    return this;
  }

   /**
   * Throttle Rate
   * @return serviceThrottleRate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Throttle Rate")

  public Long getServiceThrottleRate() {
    return serviceThrottleRate;
  }


  public void setServiceThrottleRate(Long serviceThrottleRate) {
    this.serviceThrottleRate = serviceThrottleRate;
  }


  public NetworkPoolServerUpdateInfoblox ignoreSsl(Boolean ignoreSsl) {
    
    this.ignoreSsl = ignoreSsl;
    return this;
  }

   /**
   * Disable SSL SNI Verification
   * @return ignoreSsl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Disable SSL SNI Verification")

  public Boolean getIgnoreSsl() {
    return ignoreSsl;
  }


  public void setIgnoreSsl(Boolean ignoreSsl) {
    this.ignoreSsl = ignoreSsl;
  }


  public NetworkPoolServerUpdateInfoblox networkFilter(String networkFilter) {
    
    this.networkFilter = networkFilter;
    return this;
  }

   /**
   * Network Filter
   * @return networkFilter
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Network Filter")

  public String getNetworkFilter() {
    return networkFilter;
  }


  public void setNetworkFilter(String networkFilter) {
    this.networkFilter = networkFilter;
  }


  public NetworkPoolServerUpdateInfoblox zoneFilter(String zoneFilter) {
    
    this.zoneFilter = zoneFilter;
    return this;
  }

   /**
   * Zone Filter
   * @return zoneFilter
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Zone Filter")

  public String getZoneFilter() {
    return zoneFilter;
  }


  public void setZoneFilter(String zoneFilter) {
    this.zoneFilter = zoneFilter;
  }


  public NetworkPoolServerUpdateInfoblox tenantMatch(String tenantMatch) {
    
    this.tenantMatch = tenantMatch;
    return this;
  }

   /**
   * Tenant Match
   * @return tenantMatch
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Tenant Match")

  public String getTenantMatch() {
    return tenantMatch;
  }


  public void setTenantMatch(String tenantMatch) {
    this.tenantMatch = tenantMatch;
  }


  public NetworkPoolServerUpdateInfoblox serviceMode(ServiceModeEnum serviceMode) {
    
    this.serviceMode = serviceMode;
    return this;
  }

   /**
   * IP Mode
   * @return serviceMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "static", value = "IP Mode")

  public ServiceModeEnum getServiceMode() {
    return serviceMode;
  }


  public void setServiceMode(ServiceModeEnum serviceMode) {
    this.serviceMode = serviceMode;
  }


  public NetworkPoolServerUpdateInfoblox config(NetworkPoolServerCreateInfobloxConfig config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public NetworkPoolServerCreateInfobloxConfig getConfig() {
    return config;
  }


  public void setConfig(NetworkPoolServerCreateInfobloxConfig config) {
    this.config = config;
  }


  public NetworkPoolServerUpdateInfoblox credential(NetworkPoolServerCreateBluecatCredential credential) {
    
    this.credential = credential;
    return this;
  }

   /**
   * Get credential
   * @return credential
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public NetworkPoolServerCreateBluecatCredential getCredential() {
    return credential;
  }


  public void setCredential(NetworkPoolServerCreateBluecatCredential credential) {
    this.credential = credential;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NetworkPoolServerUpdateInfoblox networkPoolServerUpdateInfoblox = (NetworkPoolServerUpdateInfoblox) o;
    return Objects.equals(this.name, networkPoolServerUpdateInfoblox.name) &&
        Objects.equals(this.enabled, networkPoolServerUpdateInfoblox.enabled) &&
        Objects.equals(this.serviceUrl, networkPoolServerUpdateInfoblox.serviceUrl) &&
        Objects.equals(this.serviceUsername, networkPoolServerUpdateInfoblox.serviceUsername) &&
        Objects.equals(this.servicePassword, networkPoolServerUpdateInfoblox.servicePassword) &&
        Objects.equals(this.serviceThrottleRate, networkPoolServerUpdateInfoblox.serviceThrottleRate) &&
        Objects.equals(this.ignoreSsl, networkPoolServerUpdateInfoblox.ignoreSsl) &&
        Objects.equals(this.networkFilter, networkPoolServerUpdateInfoblox.networkFilter) &&
        Objects.equals(this.zoneFilter, networkPoolServerUpdateInfoblox.zoneFilter) &&
        Objects.equals(this.tenantMatch, networkPoolServerUpdateInfoblox.tenantMatch) &&
        Objects.equals(this.serviceMode, networkPoolServerUpdateInfoblox.serviceMode) &&
        Objects.equals(this.config, networkPoolServerUpdateInfoblox.config) &&
        Objects.equals(this.credential, networkPoolServerUpdateInfoblox.credential);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, enabled, serviceUrl, serviceUsername, servicePassword, serviceThrottleRate, ignoreSsl, networkFilter, zoneFilter, tenantMatch, serviceMode, config, credential);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkPoolServerUpdateInfoblox {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    serviceUrl: ").append(toIndentedString(serviceUrl)).append("\n");
    sb.append("    serviceUsername: ").append(toIndentedString(serviceUsername)).append("\n");
    sb.append("    servicePassword: ").append(toIndentedString(servicePassword)).append("\n");
    sb.append("    serviceThrottleRate: ").append(toIndentedString(serviceThrottleRate)).append("\n");
    sb.append("    ignoreSsl: ").append(toIndentedString(ignoreSsl)).append("\n");
    sb.append("    networkFilter: ").append(toIndentedString(networkFilter)).append("\n");
    sb.append("    zoneFilter: ").append(toIndentedString(zoneFilter)).append("\n");
    sb.append("    tenantMatch: ").append(toIndentedString(tenantMatch)).append("\n");
    sb.append("    serviceMode: ").append(toIndentedString(serviceMode)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    credential: ").append(toIndentedString(credential)).append("\n");
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
