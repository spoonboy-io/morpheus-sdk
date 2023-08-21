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

/**
 * ApiZonesIdZone
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiZonesIdZone {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_LOCATION = "location";
  @SerializedName(SERIALIZED_NAME_LOCATION)
  private String location;

  /**
   * private or public
   */
  @JsonAdapter(VisibilityEnum.Adapter.class)
  public enum VisibilityEnum {
    PRIVATE("private"),
    
    PUBLIC("public");

    private String value;

    VisibilityEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static VisibilityEnum fromValue(String value) {
      for (VisibilityEnum b : VisibilityEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<VisibilityEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final VisibilityEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public VisibilityEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return VisibilityEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private VisibilityEnum visibility = VisibilityEnum.PRIVATE;

  public static final String SERIALIZED_NAME_ZONE_TYPE = "zoneType";
  @SerializedName(SERIALIZED_NAME_ZONE_TYPE)
  private String zoneType = "standard";

  public static final String SERIALIZED_NAME_GROUP_ID = "groupId";
  @SerializedName(SERIALIZED_NAME_GROUP_ID)
  private Long groupId;

  public static final String SERIALIZED_NAME_ACCOUNT_ID = "accountId";
  @SerializedName(SERIALIZED_NAME_ACCOUNT_ID)
  private Long accountId;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled = true;

  public static final String SERIALIZED_NAME_AUTO_RECOVER_POWER_STATE = "autoRecoverPowerState";
  @SerializedName(SERIALIZED_NAME_AUTO_RECOVER_POWER_STATE)
  private Boolean autoRecoverPowerState = false;

  public static final String SERIALIZED_NAME_SCALE_PRIORITY = "scalePriority";
  @SerializedName(SERIALIZED_NAME_SCALE_PRIORITY)
  private Long scalePriority = 1l;

  public static final String SERIALIZED_NAME_LINKED_ACCOUNT_ID = "linkedAccountId";
  @SerializedName(SERIALIZED_NAME_LINKED_ACCOUNT_ID)
  private Long linkedAccountId;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Object config;

  public static final String SERIALIZED_NAME_SECURITY_MODE = "securityMode";
  @SerializedName(SERIALIZED_NAME_SECURITY_MODE)
  private String securityMode = "off";

  public static final String SERIALIZED_NAME_DEFAULT_CLOUD_LOGOS = "defaultCloudLogos";
  @SerializedName(SERIALIZED_NAME_DEFAULT_CLOUD_LOGOS)
  private Boolean defaultCloudLogos;

  public static final String SERIALIZED_NAME_CREDENTIAL = "credential";
  @SerializedName(SERIALIZED_NAME_CREDENTIAL)
  private Object credential = {"type":"local"};


  public ApiZonesIdZone name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * A unique name scoped to your account for the cloud
   * @return name
  **/
  @ApiModelProperty(example = "My Cloud", required = true, value = "A unique name scoped to your account for the cloud")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ApiZonesIdZone description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Optional description field if you want to put more info there
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Optional description field if you want to put more info there")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ApiZonesIdZone code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Optional code for use with policies
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "mycloud", value = "Optional code for use with policies")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public ApiZonesIdZone location(String location) {
    
    this.location = location;
    return this;
  }

   /**
   * Optional location for your cloud
   * @return location
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "US East", value = "Optional location for your cloud")

  public String getLocation() {
    return location;
  }


  public void setLocation(String location) {
    this.location = location;
  }


  public ApiZonesIdZone visibility(VisibilityEnum visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * private or public
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "private or public")

  public VisibilityEnum getVisibility() {
    return visibility;
  }


  public void setVisibility(VisibilityEnum visibility) {
    this.visibility = visibility;
  }


  public ApiZonesIdZone zoneType(String zoneType) {
    
    this.zoneType = zoneType;
    return this;
  }

   /**
   * Map containing code or id of the cloud type
   * @return zoneType
  **/
  @ApiModelProperty(required = true, value = "Map containing code or id of the cloud type")

  public String getZoneType() {
    return zoneType;
  }


  public void setZoneType(String zoneType) {
    this.zoneType = zoneType;
  }


  public ApiZonesIdZone groupId(Long groupId) {
    
    this.groupId = groupId;
    return this;
  }

   /**
   * Specifies which Server group this cloud should be assigned to
   * @return groupId
  **/
  @ApiModelProperty(example = "3", required = true, value = "Specifies which Server group this cloud should be assigned to")

  public Long getGroupId() {
    return groupId;
  }


  public void setGroupId(Long groupId) {
    this.groupId = groupId;
  }


  public ApiZonesIdZone accountId(Long accountId) {
    
    this.accountId = accountId;
    return this;
  }

   /**
   * Specifies which Tenant this cloud should be assigned to
   * @return accountId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "1", value = "Specifies which Tenant this cloud should be assigned to")

  public Long getAccountId() {
    return accountId;
  }


  public void setAccountId(Long accountId) {
    this.accountId = accountId;
  }


  public ApiZonesIdZone enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Can be used to disable the cloud
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to disable the cloud")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public ApiZonesIdZone autoRecoverPowerState(Boolean autoRecoverPowerState) {
    
    this.autoRecoverPowerState = autoRecoverPowerState;
    return this;
  }

   /**
   * Automatically Power on VMs
   * @return autoRecoverPowerState
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Automatically Power on VMs")

  public Boolean getAutoRecoverPowerState() {
    return autoRecoverPowerState;
  }


  public void setAutoRecoverPowerState(Boolean autoRecoverPowerState) {
    this.autoRecoverPowerState = autoRecoverPowerState;
  }


  public ApiZonesIdZone scalePriority(Long scalePriority) {
    
    this.scalePriority = scalePriority;
    return this;
  }

   /**
   * Scale Priority
   * @return scalePriority
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Scale Priority")

  public Long getScalePriority() {
    return scalePriority;
  }


  public void setScalePriority(Long scalePriority) {
    this.scalePriority = scalePriority;
  }


  public ApiZonesIdZone linkedAccountId(Long linkedAccountId) {
    
    this.linkedAccountId = linkedAccountId;
    return this;
  }

   /**
   * Linked Account ID (enter commercial ID to get costing for AWS Govcloud)
   * @return linkedAccountId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Linked Account ID (enter commercial ID to get costing for AWS Govcloud)")

  public Long getLinkedAccountId() {
    return linkedAccountId;
  }


  public void setLinkedAccountId(Long linkedAccountId) {
    this.linkedAccountId = linkedAccountId;
  }


  public ApiZonesIdZone config(Object config) {
    
    this.config = config;
    return this;
  }

   /**
   * Map containing zone configuration settings. See the section on specific zone types for details.
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Map containing zone configuration settings. See the section on specific zone types for details.")

  public Object getConfig() {
    return config;
  }


  public void setConfig(Object config) {
    this.config = config;
  }


  public ApiZonesIdZone securityMode(String securityMode) {
    
    this.securityMode = securityMode;
    return this;
  }

   /**
   * host firewall. &#x60;off&#x60; or &#x60;internal&#x60;. a.k.a. \&quot;local firewall\&quot;
   * @return securityMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "host firewall. `off` or `internal`. a.k.a. \"local firewall\"")

  public String getSecurityMode() {
    return securityMode;
  }


  public void setSecurityMode(String securityMode) {
    this.securityMode = securityMode;
  }


  public ApiZonesIdZone defaultCloudLogos(Boolean defaultCloudLogos) {
    
    this.defaultCloudLogos = defaultCloudLogos;
    return this;
  }

   /**
   * Can be used to clear any custom logo and darkLogo, reverting to the defaults for the cloud type
   * @return defaultCloudLogos
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to clear any custom logo and darkLogo, reverting to the defaults for the cloud type")

  public Boolean getDefaultCloudLogos() {
    return defaultCloudLogos;
  }


  public void setDefaultCloudLogos(Boolean defaultCloudLogos) {
    this.defaultCloudLogos = defaultCloudLogos;
  }


  public ApiZonesIdZone credential(Object credential) {
    
    this.credential = credential;
    return this;
  }

   /**
   * Map containing Credential ID. &#x60;local&#x60; means use the values set in the local cloud config instead of associating a credential.
   * @return credential
  **/
  @ApiModelProperty(example = "{\"id\":1}", required = true, value = "Map containing Credential ID. `local` means use the values set in the local cloud config instead of associating a credential.")

  public Object getCredential() {
    return credential;
  }


  public void setCredential(Object credential) {
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
    ApiZonesIdZone apiZonesIdZone = (ApiZonesIdZone) o;
    return Objects.equals(this.name, apiZonesIdZone.name) &&
        Objects.equals(this.description, apiZonesIdZone.description) &&
        Objects.equals(this.code, apiZonesIdZone.code) &&
        Objects.equals(this.location, apiZonesIdZone.location) &&
        Objects.equals(this.visibility, apiZonesIdZone.visibility) &&
        Objects.equals(this.zoneType, apiZonesIdZone.zoneType) &&
        Objects.equals(this.groupId, apiZonesIdZone.groupId) &&
        Objects.equals(this.accountId, apiZonesIdZone.accountId) &&
        Objects.equals(this.enabled, apiZonesIdZone.enabled) &&
        Objects.equals(this.autoRecoverPowerState, apiZonesIdZone.autoRecoverPowerState) &&
        Objects.equals(this.scalePriority, apiZonesIdZone.scalePriority) &&
        Objects.equals(this.linkedAccountId, apiZonesIdZone.linkedAccountId) &&
        Objects.equals(this.config, apiZonesIdZone.config) &&
        Objects.equals(this.securityMode, apiZonesIdZone.securityMode) &&
        Objects.equals(this.defaultCloudLogos, apiZonesIdZone.defaultCloudLogos) &&
        Objects.equals(this.credential, apiZonesIdZone.credential);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, code, location, visibility, zoneType, groupId, accountId, enabled, autoRecoverPowerState, scalePriority, linkedAccountId, config, securityMode, defaultCloudLogos, credential);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiZonesIdZone {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    location: ").append(toIndentedString(location)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    zoneType: ").append(toIndentedString(zoneType)).append("\n");
    sb.append("    groupId: ").append(toIndentedString(groupId)).append("\n");
    sb.append("    accountId: ").append(toIndentedString(accountId)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    autoRecoverPowerState: ").append(toIndentedString(autoRecoverPowerState)).append("\n");
    sb.append("    scalePriority: ").append(toIndentedString(scalePriority)).append("\n");
    sb.append("    linkedAccountId: ").append(toIndentedString(linkedAccountId)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    securityMode: ").append(toIndentedString(securityMode)).append("\n");
    sb.append("    defaultCloudLogos: ").append(toIndentedString(defaultCloudLogos)).append("\n");
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

