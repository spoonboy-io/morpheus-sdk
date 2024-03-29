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
 * SecurityGroupRules
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class SecurityGroupRules {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_RULE_TYPE = "ruleType";
  @SerializedName(SERIALIZED_NAME_RULE_TYPE)
  private String ruleType;

  public static final String SERIALIZED_NAME_CUSTOM_RULE = "customRule";
  @SerializedName(SERIALIZED_NAME_CUSTOM_RULE)
  private Boolean customRule;

  public static final String SERIALIZED_NAME_INSTANCE_TYPE_ID = "instanceTypeId";
  @SerializedName(SERIALIZED_NAME_INSTANCE_TYPE_ID)
  private String instanceTypeId;

  public static final String SERIALIZED_NAME_DIRECTION = "direction";
  @SerializedName(SERIALIZED_NAME_DIRECTION)
  private String direction;

  public static final String SERIALIZED_NAME_POLICY = "policy";
  @SerializedName(SERIALIZED_NAME_POLICY)
  private String policy;

  public static final String SERIALIZED_NAME_SOURCE_TYPE = "sourceType";
  @SerializedName(SERIALIZED_NAME_SOURCE_TYPE)
  private String sourceType;

  public static final String SERIALIZED_NAME_SOURCE = "source";
  @SerializedName(SERIALIZED_NAME_SOURCE)
  private String source;

  public static final String SERIALIZED_NAME_SOURCE_GROUP = "sourceGroup";
  @SerializedName(SERIALIZED_NAME_SOURCE_GROUP)
  private String sourceGroup;

  public static final String SERIALIZED_NAME_SOURCE_TIER = "sourceTier";
  @SerializedName(SERIALIZED_NAME_SOURCE_TIER)
  private String sourceTier;

  public static final String SERIALIZED_NAME_PORT_RANGE = "portRange";
  @SerializedName(SERIALIZED_NAME_PORT_RANGE)
  private String portRange;

  public static final String SERIALIZED_NAME_PROTOCOL = "protocol";
  @SerializedName(SERIALIZED_NAME_PROTOCOL)
  private String protocol;

  public static final String SERIALIZED_NAME_DESTINATION_TYPE = "destinationType";
  @SerializedName(SERIALIZED_NAME_DESTINATION_TYPE)
  private String destinationType;

  public static final String SERIALIZED_NAME_DESTINATION = "destination";
  @SerializedName(SERIALIZED_NAME_DESTINATION)
  private String destination;

  public static final String SERIALIZED_NAME_DESTINATION_GROUP = "destinationGroup";
  @SerializedName(SERIALIZED_NAME_DESTINATION_GROUP)
  private String destinationGroup;

  public static final String SERIALIZED_NAME_DESTINATION_TIER = "destinationTier";
  @SerializedName(SERIALIZED_NAME_DESTINATION_TIER)
  private String destinationTier;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private String enabled;


  public SecurityGroupRules id(Long id) {
    
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


  public SecurityGroupRules name(String name) {
    
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


  public SecurityGroupRules ruleType(String ruleType) {
    
    this.ruleType = ruleType;
    return this;
  }

   /**
   * Get ruleType
   * @return ruleType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRuleType() {
    return ruleType;
  }


  public void setRuleType(String ruleType) {
    this.ruleType = ruleType;
  }


  public SecurityGroupRules customRule(Boolean customRule) {
    
    this.customRule = customRule;
    return this;
  }

   /**
   * Get customRule
   * @return customRule
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getCustomRule() {
    return customRule;
  }


  public void setCustomRule(Boolean customRule) {
    this.customRule = customRule;
  }


  public SecurityGroupRules instanceTypeId(String instanceTypeId) {
    
    this.instanceTypeId = instanceTypeId;
    return this;
  }

   /**
   * Get instanceTypeId
   * @return instanceTypeId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInstanceTypeId() {
    return instanceTypeId;
  }


  public void setInstanceTypeId(String instanceTypeId) {
    this.instanceTypeId = instanceTypeId;
  }


  public SecurityGroupRules direction(String direction) {
    
    this.direction = direction;
    return this;
  }

   /**
   * Get direction
   * @return direction
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDirection() {
    return direction;
  }


  public void setDirection(String direction) {
    this.direction = direction;
  }


  public SecurityGroupRules policy(String policy) {
    
    this.policy = policy;
    return this;
  }

   /**
   * Get policy
   * @return policy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPolicy() {
    return policy;
  }


  public void setPolicy(String policy) {
    this.policy = policy;
  }


  public SecurityGroupRules sourceType(String sourceType) {
    
    this.sourceType = sourceType;
    return this;
  }

   /**
   * Get sourceType
   * @return sourceType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceType() {
    return sourceType;
  }


  public void setSourceType(String sourceType) {
    this.sourceType = sourceType;
  }


  public SecurityGroupRules source(String source) {
    
    this.source = source;
    return this;
  }

   /**
   * Get source
   * @return source
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSource() {
    return source;
  }


  public void setSource(String source) {
    this.source = source;
  }


  public SecurityGroupRules sourceGroup(String sourceGroup) {
    
    this.sourceGroup = sourceGroup;
    return this;
  }

   /**
   * Get sourceGroup
   * @return sourceGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceGroup() {
    return sourceGroup;
  }


  public void setSourceGroup(String sourceGroup) {
    this.sourceGroup = sourceGroup;
  }


  public SecurityGroupRules sourceTier(String sourceTier) {
    
    this.sourceTier = sourceTier;
    return this;
  }

   /**
   * Get sourceTier
   * @return sourceTier
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceTier() {
    return sourceTier;
  }


  public void setSourceTier(String sourceTier) {
    this.sourceTier = sourceTier;
  }


  public SecurityGroupRules portRange(String portRange) {
    
    this.portRange = portRange;
    return this;
  }

   /**
   * Get portRange
   * @return portRange
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPortRange() {
    return portRange;
  }


  public void setPortRange(String portRange) {
    this.portRange = portRange;
  }


  public SecurityGroupRules protocol(String protocol) {
    
    this.protocol = protocol;
    return this;
  }

   /**
   * Get protocol
   * @return protocol
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getProtocol() {
    return protocol;
  }


  public void setProtocol(String protocol) {
    this.protocol = protocol;
  }


  public SecurityGroupRules destinationType(String destinationType) {
    
    this.destinationType = destinationType;
    return this;
  }

   /**
   * Get destinationType
   * @return destinationType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDestinationType() {
    return destinationType;
  }


  public void setDestinationType(String destinationType) {
    this.destinationType = destinationType;
  }


  public SecurityGroupRules destination(String destination) {
    
    this.destination = destination;
    return this;
  }

   /**
   * Get destination
   * @return destination
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDestination() {
    return destination;
  }


  public void setDestination(String destination) {
    this.destination = destination;
  }


  public SecurityGroupRules destinationGroup(String destinationGroup) {
    
    this.destinationGroup = destinationGroup;
    return this;
  }

   /**
   * Get destinationGroup
   * @return destinationGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDestinationGroup() {
    return destinationGroup;
  }


  public void setDestinationGroup(String destinationGroup) {
    this.destinationGroup = destinationGroup;
  }


  public SecurityGroupRules destinationTier(String destinationTier) {
    
    this.destinationTier = destinationTier;
    return this;
  }

   /**
   * Get destinationTier
   * @return destinationTier
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDestinationTier() {
    return destinationTier;
  }


  public void setDestinationTier(String destinationTier) {
    this.destinationTier = destinationTier;
  }


  public SecurityGroupRules externalId(String externalId) {
    
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


  public SecurityGroupRules enabled(String enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Get enabled
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEnabled() {
    return enabled;
  }


  public void setEnabled(String enabled) {
    this.enabled = enabled;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SecurityGroupRules securityGroupRules = (SecurityGroupRules) o;
    return Objects.equals(this.id, securityGroupRules.id) &&
        Objects.equals(this.name, securityGroupRules.name) &&
        Objects.equals(this.ruleType, securityGroupRules.ruleType) &&
        Objects.equals(this.customRule, securityGroupRules.customRule) &&
        Objects.equals(this.instanceTypeId, securityGroupRules.instanceTypeId) &&
        Objects.equals(this.direction, securityGroupRules.direction) &&
        Objects.equals(this.policy, securityGroupRules.policy) &&
        Objects.equals(this.sourceType, securityGroupRules.sourceType) &&
        Objects.equals(this.source, securityGroupRules.source) &&
        Objects.equals(this.sourceGroup, securityGroupRules.sourceGroup) &&
        Objects.equals(this.sourceTier, securityGroupRules.sourceTier) &&
        Objects.equals(this.portRange, securityGroupRules.portRange) &&
        Objects.equals(this.protocol, securityGroupRules.protocol) &&
        Objects.equals(this.destinationType, securityGroupRules.destinationType) &&
        Objects.equals(this.destination, securityGroupRules.destination) &&
        Objects.equals(this.destinationGroup, securityGroupRules.destinationGroup) &&
        Objects.equals(this.destinationTier, securityGroupRules.destinationTier) &&
        Objects.equals(this.externalId, securityGroupRules.externalId) &&
        Objects.equals(this.enabled, securityGroupRules.enabled);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, ruleType, customRule, instanceTypeId, direction, policy, sourceType, source, sourceGroup, sourceTier, portRange, protocol, destinationType, destination, destinationGroup, destinationTier, externalId, enabled);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SecurityGroupRules {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    ruleType: ").append(toIndentedString(ruleType)).append("\n");
    sb.append("    customRule: ").append(toIndentedString(customRule)).append("\n");
    sb.append("    instanceTypeId: ").append(toIndentedString(instanceTypeId)).append("\n");
    sb.append("    direction: ").append(toIndentedString(direction)).append("\n");
    sb.append("    policy: ").append(toIndentedString(policy)).append("\n");
    sb.append("    sourceType: ").append(toIndentedString(sourceType)).append("\n");
    sb.append("    source: ").append(toIndentedString(source)).append("\n");
    sb.append("    sourceGroup: ").append(toIndentedString(sourceGroup)).append("\n");
    sb.append("    sourceTier: ").append(toIndentedString(sourceTier)).append("\n");
    sb.append("    portRange: ").append(toIndentedString(portRange)).append("\n");
    sb.append("    protocol: ").append(toIndentedString(protocol)).append("\n");
    sb.append("    destinationType: ").append(toIndentedString(destinationType)).append("\n");
    sb.append("    destination: ").append(toIndentedString(destination)).append("\n");
    sb.append("    destinationGroup: ").append(toIndentedString(destinationGroup)).append("\n");
    sb.append("    destinationTier: ").append(toIndentedString(destinationTier)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
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

