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
import org.openapitools.client.model.Creds;
import org.openapitools.client.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.threeten.bp.OffsetDateTime;

/**
 * IntegrationCypher
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class IntegrationCypher {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  /**
   * Gets or Sets type
   */
  @JsonAdapter(TypeEnum.Adapter.class)
  public enum TypeEnum {
    CYPHER("cypher");

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

  public static final String SERIALIZED_NAME_INTEGRATION_TYPE = "integrationType";
  @SerializedName(SERIALIZED_NAME_INTEGRATION_TYPE)
  private InlineResponse20079LoadBalancerMonitorLoadBalancerType integrationType;

  public static final String SERIALIZED_NAME_URL = "url";
  @SerializedName(SERIALIZED_NAME_URL)
  private String url;

  public static final String SERIALIZED_NAME_TOKEN = "token";
  @SerializedName(SERIALIZED_NAME_TOKEN)
  private String token;

  public static final String SERIALIZED_NAME_TOKEN_HASH = "tokenHash";
  @SerializedName(SERIALIZED_NAME_TOKEN_HASH)
  private String tokenHash;

  public static final String SERIALIZED_NAME_IS_PLUGIN = "isPlugin";
  @SerializedName(SERIALIZED_NAME_IS_PLUGIN)
  private Boolean isPlugin;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Object config;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_STATUS_DATE = "statusDate";
  @SerializedName(SERIALIZED_NAME_STATUS_DATE)
  private OffsetDateTime statusDate;

  public static final String SERIALIZED_NAME_STATUS_MESSAGE = "statusMessage";
  @SerializedName(SERIALIZED_NAME_STATUS_MESSAGE)
  private String statusMessage;

  public static final String SERIALIZED_NAME_LAST_SYNC = "lastSync";
  @SerializedName(SERIALIZED_NAME_LAST_SYNC)
  private String lastSync;

  public static final String SERIALIZED_NAME_LAST_SYNC_DURATION = "lastSyncDuration";
  @SerializedName(SERIALIZED_NAME_LAST_SYNC_DURATION)
  private String lastSyncDuration;

  public static final String SERIALIZED_NAME_CREDENTIAL = "credential";
  @SerializedName(SERIALIZED_NAME_CREDENTIAL)
  private Creds credential;


  public IntegrationCypher id(Long id) {
    
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


  public IntegrationCypher name(String name) {
    
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


  public IntegrationCypher enabled(Boolean enabled) {
    
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


  public IntegrationCypher type(TypeEnum type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public TypeEnum getType() {
    return type;
  }


  public void setType(TypeEnum type) {
    this.type = type;
  }


  public IntegrationCypher integrationType(InlineResponse20079LoadBalancerMonitorLoadBalancerType integrationType) {
    
    this.integrationType = integrationType;
    return this;
  }

   /**
   * Get integrationType
   * @return integrationType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20079LoadBalancerMonitorLoadBalancerType getIntegrationType() {
    return integrationType;
  }


  public void setIntegrationType(InlineResponse20079LoadBalancerMonitorLoadBalancerType integrationType) {
    this.integrationType = integrationType;
  }


  public IntegrationCypher url(String url) {
    
    this.url = url;
    return this;
  }

   /**
   * Get url
   * @return url
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUrl() {
    return url;
  }


  public void setUrl(String url) {
    this.url = url;
  }


  public IntegrationCypher token(String token) {
    
    this.token = token;
    return this;
  }

   /**
   * Get token
   * @return token
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getToken() {
    return token;
  }


  public void setToken(String token) {
    this.token = token;
  }


  public IntegrationCypher tokenHash(String tokenHash) {
    
    this.tokenHash = tokenHash;
    return this;
  }

   /**
   * Get tokenHash
   * @return tokenHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTokenHash() {
    return tokenHash;
  }


  public void setTokenHash(String tokenHash) {
    this.tokenHash = tokenHash;
  }


  public IntegrationCypher isPlugin(Boolean isPlugin) {
    
    this.isPlugin = isPlugin;
    return this;
  }

   /**
   * Get isPlugin
   * @return isPlugin
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsPlugin() {
    return isPlugin;
  }


  public void setIsPlugin(Boolean isPlugin) {
    this.isPlugin = isPlugin;
  }


  public IntegrationCypher config(Object config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getConfig() {
    return config;
  }


  public void setConfig(Object config) {
    this.config = config;
  }


  public IntegrationCypher status(String status) {
    
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


  public IntegrationCypher statusDate(OffsetDateTime statusDate) {
    
    this.statusDate = statusDate;
    return this;
  }

   /**
   * Get statusDate
   * @return statusDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getStatusDate() {
    return statusDate;
  }


  public void setStatusDate(OffsetDateTime statusDate) {
    this.statusDate = statusDate;
  }


  public IntegrationCypher statusMessage(String statusMessage) {
    
    this.statusMessage = statusMessage;
    return this;
  }

   /**
   * Get statusMessage
   * @return statusMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatusMessage() {
    return statusMessage;
  }


  public void setStatusMessage(String statusMessage) {
    this.statusMessage = statusMessage;
  }


  public IntegrationCypher lastSync(String lastSync) {
    
    this.lastSync = lastSync;
    return this;
  }

   /**
   * Get lastSync
   * @return lastSync
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLastSync() {
    return lastSync;
  }


  public void setLastSync(String lastSync) {
    this.lastSync = lastSync;
  }


  public IntegrationCypher lastSyncDuration(String lastSyncDuration) {
    
    this.lastSyncDuration = lastSyncDuration;
    return this;
  }

   /**
   * Get lastSyncDuration
   * @return lastSyncDuration
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLastSyncDuration() {
    return lastSyncDuration;
  }


  public void setLastSyncDuration(String lastSyncDuration) {
    this.lastSyncDuration = lastSyncDuration;
  }


  public IntegrationCypher credential(Creds credential) {
    
    this.credential = credential;
    return this;
  }

   /**
   * Get credential
   * @return credential
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Creds getCredential() {
    return credential;
  }


  public void setCredential(Creds credential) {
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
    IntegrationCypher integrationCypher = (IntegrationCypher) o;
    return Objects.equals(this.id, integrationCypher.id) &&
        Objects.equals(this.name, integrationCypher.name) &&
        Objects.equals(this.enabled, integrationCypher.enabled) &&
        Objects.equals(this.type, integrationCypher.type) &&
        Objects.equals(this.integrationType, integrationCypher.integrationType) &&
        Objects.equals(this.url, integrationCypher.url) &&
        Objects.equals(this.token, integrationCypher.token) &&
        Objects.equals(this.tokenHash, integrationCypher.tokenHash) &&
        Objects.equals(this.isPlugin, integrationCypher.isPlugin) &&
        Objects.equals(this.config, integrationCypher.config) &&
        Objects.equals(this.status, integrationCypher.status) &&
        Objects.equals(this.statusDate, integrationCypher.statusDate) &&
        Objects.equals(this.statusMessage, integrationCypher.statusMessage) &&
        Objects.equals(this.lastSync, integrationCypher.lastSync) &&
        Objects.equals(this.lastSyncDuration, integrationCypher.lastSyncDuration) &&
        Objects.equals(this.credential, integrationCypher.credential);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, enabled, type, integrationType, url, token, tokenHash, isPlugin, config, status, statusDate, statusMessage, lastSync, lastSyncDuration, credential);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class IntegrationCypher {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    integrationType: ").append(toIndentedString(integrationType)).append("\n");
    sb.append("    url: ").append(toIndentedString(url)).append("\n");
    sb.append("    token: ").append(toIndentedString(token)).append("\n");
    sb.append("    tokenHash: ").append(toIndentedString(tokenHash)).append("\n");
    sb.append("    isPlugin: ").append(toIndentedString(isPlugin)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    statusDate: ").append(toIndentedString(statusDate)).append("\n");
    sb.append("    statusMessage: ").append(toIndentedString(statusMessage)).append("\n");
    sb.append("    lastSync: ").append(toIndentedString(lastSync)).append("\n");
    sb.append("    lastSyncDuration: ").append(toIndentedString(lastSyncDuration)).append("\n");
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

