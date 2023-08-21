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

/**
 * NetworkServerGroupMember
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class NetworkServerGroupMember {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_MEMBER_NAME = "memberName";
  @SerializedName(SERIALIZED_NAME_MEMBER_NAME)
  private String memberName;

  public static final String SERIALIZED_NAME_MEMBER_TYPE = "memberType";
  @SerializedName(SERIALIZED_NAME_MEMBER_TYPE)
  private String memberType;

  public static final String SERIALIZED_NAME_MEMBER_VALUE = "memberValue";
  @SerializedName(SERIALIZED_NAME_MEMBER_VALUE)
  private String memberValue;

  public static final String SERIALIZED_NAME_MEMBER_EXPRESSION = "memberExpression";
  @SerializedName(SERIALIZED_NAME_MEMBER_EXPRESSION)
  private String memberExpression;

  public static final String SERIALIZED_NAME_DISPLAY_ORDER = "displayOrder";
  @SerializedName(SERIALIZED_NAME_DISPLAY_ORDER)
  private Long displayOrder;

  public static final String SERIALIZED_NAME_INTERNAL_ID = "internalId";
  @SerializedName(SERIALIZED_NAME_INTERNAL_ID)
  private String internalId;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_MEMBERS = "members";
  @SerializedName(SERIALIZED_NAME_MEMBERS)
  private List<Object> members = null;


  public NetworkServerGroupMember id(Long id) {
    
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


  public NetworkServerGroupMember category(String category) {
    
    this.category = category;
    return this;
  }

   /**
   * Get category
   * @return category
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCategory() {
    return category;
  }


  public void setCategory(String category) {
    this.category = category;
  }


  public NetworkServerGroupMember type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public NetworkServerGroupMember memberName(String memberName) {
    
    this.memberName = memberName;
    return this;
  }

   /**
   * Get memberName
   * @return memberName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMemberName() {
    return memberName;
  }


  public void setMemberName(String memberName) {
    this.memberName = memberName;
  }


  public NetworkServerGroupMember memberType(String memberType) {
    
    this.memberType = memberType;
    return this;
  }

   /**
   * Get memberType
   * @return memberType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMemberType() {
    return memberType;
  }


  public void setMemberType(String memberType) {
    this.memberType = memberType;
  }


  public NetworkServerGroupMember memberValue(String memberValue) {
    
    this.memberValue = memberValue;
    return this;
  }

   /**
   * Get memberValue
   * @return memberValue
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMemberValue() {
    return memberValue;
  }


  public void setMemberValue(String memberValue) {
    this.memberValue = memberValue;
  }


  public NetworkServerGroupMember memberExpression(String memberExpression) {
    
    this.memberExpression = memberExpression;
    return this;
  }

   /**
   * Get memberExpression
   * @return memberExpression
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMemberExpression() {
    return memberExpression;
  }


  public void setMemberExpression(String memberExpression) {
    this.memberExpression = memberExpression;
  }


  public NetworkServerGroupMember displayOrder(Long displayOrder) {
    
    this.displayOrder = displayOrder;
    return this;
  }

   /**
   * Get displayOrder
   * @return displayOrder
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getDisplayOrder() {
    return displayOrder;
  }


  public void setDisplayOrder(Long displayOrder) {
    this.displayOrder = displayOrder;
  }


  public NetworkServerGroupMember internalId(String internalId) {
    
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


  public NetworkServerGroupMember externalId(String externalId) {
    
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


  public NetworkServerGroupMember members(List<Object> members) {
    
    this.members = members;
    return this;
  }

  public NetworkServerGroupMember addMembersItem(Object membersItem) {
    if (this.members == null) {
      this.members = new ArrayList<Object>();
    }
    this.members.add(membersItem);
    return this;
  }

   /**
   * Get members
   * @return members
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getMembers() {
    return members;
  }


  public void setMembers(List<Object> members) {
    this.members = members;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NetworkServerGroupMember networkServerGroupMember = (NetworkServerGroupMember) o;
    return Objects.equals(this.id, networkServerGroupMember.id) &&
        Objects.equals(this.category, networkServerGroupMember.category) &&
        Objects.equals(this.type, networkServerGroupMember.type) &&
        Objects.equals(this.memberName, networkServerGroupMember.memberName) &&
        Objects.equals(this.memberType, networkServerGroupMember.memberType) &&
        Objects.equals(this.memberValue, networkServerGroupMember.memberValue) &&
        Objects.equals(this.memberExpression, networkServerGroupMember.memberExpression) &&
        Objects.equals(this.displayOrder, networkServerGroupMember.displayOrder) &&
        Objects.equals(this.internalId, networkServerGroupMember.internalId) &&
        Objects.equals(this.externalId, networkServerGroupMember.externalId) &&
        Objects.equals(this.members, networkServerGroupMember.members);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, category, type, memberName, memberType, memberValue, memberExpression, displayOrder, internalId, externalId, members);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NetworkServerGroupMember {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    memberName: ").append(toIndentedString(memberName)).append("\n");
    sb.append("    memberType: ").append(toIndentedString(memberType)).append("\n");
    sb.append("    memberValue: ").append(toIndentedString(memberValue)).append("\n");
    sb.append("    memberExpression: ").append(toIndentedString(memberExpression)).append("\n");
    sb.append("    displayOrder: ").append(toIndentedString(displayOrder)).append("\n");
    sb.append("    internalId: ").append(toIndentedString(internalId)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    members: ").append(toIndentedString(members)).append("\n");
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
