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
import org.openapitools.client.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.client.model.NetworkServerGroupMember;
import org.openapitools.client.model.Permissions;
import org.openapitools.client.model.Tag;

/**
 * InlineResponse200117Group
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse200117Group {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_INTERNAL_ID = "internalId";
  @SerializedName(SERIALIZED_NAME_INTERNAL_ID)
  private String internalId;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account;

  public static final String SERIALIZED_NAME_OWNER = "owner";
  @SerializedName(SERIALIZED_NAME_OWNER)
  private ApiBlueprintsIdUpdatePermissionsResourcePermissionSites owner;

  public static final String SERIALIZED_NAME_NETWORK_SERVER = "networkServer";
  @SerializedName(SERIALIZED_NAME_NETWORK_SERVER)
  private ApiBlueprintsIdUpdatePermissionsResourcePermissionSites networkServer;

  public static final String SERIALIZED_NAME_PERMISSIONS = "permissions";
  @SerializedName(SERIALIZED_NAME_PERMISSIONS)
  private Permissions permissions;

  public static final String SERIALIZED_NAME_TAGS = "tags";
  @SerializedName(SERIALIZED_NAME_TAGS)
  private List<Tag> tags = null;

  public static final String SERIALIZED_NAME_MEMBERS = "members";
  @SerializedName(SERIALIZED_NAME_MEMBERS)
  private List<NetworkServerGroupMember> members = null;


  public InlineResponse200117Group id(Long id) {
    
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


  public InlineResponse200117Group name(String name) {
    
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


  public InlineResponse200117Group description(String description) {
    
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


  public InlineResponse200117Group internalId(String internalId) {
    
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


  public InlineResponse200117Group externalId(String externalId) {
    
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


  public InlineResponse200117Group visibility(String visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * Get visibility
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVisibility() {
    return visibility;
  }


  public void setVisibility(String visibility) {
    this.visibility = visibility;
  }


  public InlineResponse200117Group account(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account) {
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiBlueprintsIdUpdatePermissionsResourcePermissionSites getAccount() {
    return account;
  }


  public void setAccount(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account) {
    this.account = account;
  }


  public InlineResponse200117Group owner(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites owner) {
    
    this.owner = owner;
    return this;
  }

   /**
   * Get owner
   * @return owner
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiBlueprintsIdUpdatePermissionsResourcePermissionSites getOwner() {
    return owner;
  }


  public void setOwner(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites owner) {
    this.owner = owner;
  }


  public InlineResponse200117Group networkServer(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites networkServer) {
    
    this.networkServer = networkServer;
    return this;
  }

   /**
   * Get networkServer
   * @return networkServer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiBlueprintsIdUpdatePermissionsResourcePermissionSites getNetworkServer() {
    return networkServer;
  }


  public void setNetworkServer(ApiBlueprintsIdUpdatePermissionsResourcePermissionSites networkServer) {
    this.networkServer = networkServer;
  }


  public InlineResponse200117Group permissions(Permissions permissions) {
    
    this.permissions = permissions;
    return this;
  }

   /**
   * Get permissions
   * @return permissions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Permissions getPermissions() {
    return permissions;
  }


  public void setPermissions(Permissions permissions) {
    this.permissions = permissions;
  }


  public InlineResponse200117Group tags(List<Tag> tags) {
    
    this.tags = tags;
    return this;
  }

  public InlineResponse200117Group addTagsItem(Tag tagsItem) {
    if (this.tags == null) {
      this.tags = new ArrayList<Tag>();
    }
    this.tags.add(tagsItem);
    return this;
  }

   /**
   * Get tags
   * @return tags
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Tag> getTags() {
    return tags;
  }


  public void setTags(List<Tag> tags) {
    this.tags = tags;
  }


  public InlineResponse200117Group members(List<NetworkServerGroupMember> members) {
    
    this.members = members;
    return this;
  }

  public InlineResponse200117Group addMembersItem(NetworkServerGroupMember membersItem) {
    if (this.members == null) {
      this.members = new ArrayList<NetworkServerGroupMember>();
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

  public List<NetworkServerGroupMember> getMembers() {
    return members;
  }


  public void setMembers(List<NetworkServerGroupMember> members) {
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
    InlineResponse200117Group inlineResponse200117Group = (InlineResponse200117Group) o;
    return Objects.equals(this.id, inlineResponse200117Group.id) &&
        Objects.equals(this.name, inlineResponse200117Group.name) &&
        Objects.equals(this.description, inlineResponse200117Group.description) &&
        Objects.equals(this.internalId, inlineResponse200117Group.internalId) &&
        Objects.equals(this.externalId, inlineResponse200117Group.externalId) &&
        Objects.equals(this.visibility, inlineResponse200117Group.visibility) &&
        Objects.equals(this.account, inlineResponse200117Group.account) &&
        Objects.equals(this.owner, inlineResponse200117Group.owner) &&
        Objects.equals(this.networkServer, inlineResponse200117Group.networkServer) &&
        Objects.equals(this.permissions, inlineResponse200117Group.permissions) &&
        Objects.equals(this.tags, inlineResponse200117Group.tags) &&
        Objects.equals(this.members, inlineResponse200117Group.members);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, description, internalId, externalId, visibility, account, owner, networkServer, permissions, tags, members);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse200117Group {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    internalId: ").append(toIndentedString(internalId)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    owner: ").append(toIndentedString(owner)).append("\n");
    sb.append("    networkServer: ").append(toIndentedString(networkServer)).append("\n");
    sb.append("    permissions: ").append(toIndentedString(permissions)).append("\n");
    sb.append("    tags: ").append(toIndentedString(tags)).append("\n");
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

