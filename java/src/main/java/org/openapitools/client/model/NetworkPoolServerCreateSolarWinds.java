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
import org.openapitools.client.model.NetworkPoolServerCreateBluecatConfig;
import org.openapitools.client.model.NetworkPoolServerCreateBluecatCredential;

/**
 * NetworkPoolServerCreateSolarWinds
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkPoolServerCreateSolarWinds {
  /**
   * Type Code (SolarWinds)
   */
  @JsonAdapter(TypeEnum.Adapter.class)
  public enum TypeEnum {
    SOLARWINDSIPAM("solarwindsipam");

    private String value;

    TypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static TypeEnum fromValue(String value) {
      for (TypeEnum b : TypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<TypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final TypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public TypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return TypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private TypeEnum type;

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
  private Boolean ignoreSsl = true;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private NetworkPoolServerCreateBluecatConfig config;

  public static final String SERIALIZED_NAME_CREDENTIAL = "credential";
  @SerializedName(SERIALIZED_NAME_CREDENTIAL)
  private NetworkPoolServerCreateBluecatCredential credential;


  public NetworkPoolServerCreateSolarWinds type(TypeEnum type) {
    
    this.type = type;
    return this;
  }

   /**
   * Type Code (SolarWinds)
   * @return type
  **/
  @ApiModelProperty(required = true, value = "Type Code (SolarWinds)")

  public TypeEnum getType() {
    return type;
  }


  public void setType(TypeEnum type) {
    this.type = type;
  }


  public NetworkPoolServerCreateSolarWinds name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Name
   * @return name
  **/
  @ApiModelProperty(required = true, value = "Name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public NetworkPoolServerCreateSolarWinds enabled(Boolean enabled) {
    
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


  public NetworkPoolServerCreateSolarWinds serviceUrl(String serviceUrl) {
    
    this.serviceUrl = serviceUrl;
    return this;
  }

   /**
   * URL
   * @return serviceUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "https://solarwinds-server/wapi/v2.9", required = true, value = "URL")

  public String getServiceUrl() {
    return serviceUrl;
  }


  public void setServiceUrl(String serviceUrl) {
    this.serviceUrl = serviceUrl;
  }


  public NetworkPoolServerCreateSolarWinds serviceUsername(String serviceUsername) {
    
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


  public NetworkPoolServerCreateSolarWinds servicePassword(String servicePassword) {
    
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


  public NetworkPoolServerCreateSolarWinds serviceThrottleRate(Long serviceThrottleRate) {
    
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


  public NetworkPoolServerCreateSolarWinds ignoreSsl(Boolean ignoreSsl) {
    
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


  public NetworkPoolServerCreateSolarWinds config(NetworkPoolServerCreateBluecatConfig config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public NetworkPoolServerCreateBluecatConfig getConfig() {
    return config;
  }


  public void setConfig(NetworkPoolServerCreateBluecatConfig config) {
    this.config = config;
  }


  public NetworkPoolServerCreateSolarWinds credential(NetworkPoolServerCreateBluecatCredential credential) {
    
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
    NetworkPoolServerCreateSolarWinds networkPoolServerCreateSolarWinds = (NetworkPoolServerCreateSolarWinds) o;
    return Objects.equals(this.type, networkPoolServerCreateSolarWinds.type) &&
        Objects.equals(this.name, networkPoolServerCreateSolarWinds.name) &&
        Objects.equals(this.enabled, networkPoolServerCreateSolarWinds.enabled) &&
        Objects.equals(this.serviceUrl, networkPoolServerCreateSolarWinds.serviceUrl) &&
        Objects.equals(this.serviceUsername, networkPoolServerCreateSolarWinds.serviceUsername) &&
        Objects.equals(this.servicePassword, networkPoolServerCreateSolarWinds.servicePassword) &&
        Objects.equals(this.serviceThrottleRate, networkPoolServerCreateSolarWinds.serviceThrottleRate) &&
        Objects.equals(this.ignoreSsl, networkPoolServerCreateSolarWinds.ignoreSsl) &&
        Objects.equals(this.config, networkPoolServerCreateSolarWinds.config) &&
        Objects.equals(this.credential, networkPoolServerCreateSolarWinds.credential);
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, name, enabled, serviceUrl, serviceUsername, servicePassword, serviceThrottleRate, ignoreSsl, config, credential);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkPoolServerCreateSolarWinds {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    serviceUrl: ").append(toIndentedString(serviceUrl)).append("\n");
    sb.append("    serviceUsername: ").append(toIndentedString(serviceUsername)).append("\n");
    sb.append("    servicePassword: ").append(toIndentedString(servicePassword)).append("\n");
    sb.append("    serviceThrottleRate: ").append(toIndentedString(serviceThrottleRate)).append("\n");
    sb.append("    ignoreSsl: ").append(toIndentedString(ignoreSsl)).append("\n");
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

