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
import org.openapitools.client.model.ReferenceObject;

/**
 * ApiUsersIdUser
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiUsersIdUser {
  public static final String SERIALIZED_NAME_FIRST_NAME = "firstName";
  @SerializedName(SERIALIZED_NAME_FIRST_NAME)
  private String firstName;

  public static final String SERIALIZED_NAME_LAST_NAME = "lastName";
  @SerializedName(SERIALIZED_NAME_LAST_NAME)
  private String lastName;

  public static final String SERIALIZED_NAME_USERNAME = "username";
  @SerializedName(SERIALIZED_NAME_USERNAME)
  private String username;

  public static final String SERIALIZED_NAME_EMAIL = "email";
  @SerializedName(SERIALIZED_NAME_EMAIL)
  private String email;

  public static final String SERIALIZED_NAME_PASSWORD = "password";
  @SerializedName(SERIALIZED_NAME_PASSWORD)
  private String password;

  public static final String SERIALIZED_NAME_ROLES = "roles";
  @SerializedName(SERIALIZED_NAME_ROLES)
  private List<ReferenceObject> roles = null;


  public ApiUsersIdUser firstName(String firstName) {
    
    this.firstName = firstName;
    return this;
  }

   /**
   * First Name
   * @return firstName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "First Name")

  public String getFirstName() {
    return firstName;
  }


  public void setFirstName(String firstName) {
    this.firstName = firstName;
  }


  public ApiUsersIdUser lastName(String lastName) {
    
    this.lastName = lastName;
    return this;
  }

   /**
   * Last Name
   * @return lastName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Last Name")

  public String getLastName() {
    return lastName;
  }


  public void setLastName(String lastName) {
    this.lastName = lastName;
  }


  public ApiUsersIdUser username(String username) {
    
    this.username = username;
    return this;
  }

   /**
   * Username (unique per tenant).
   * @return username
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Username (unique per tenant).")

  public String getUsername() {
    return username;
  }


  public void setUsername(String username) {
    this.username = username;
  }


  public ApiUsersIdUser email(String email) {
    
    this.email = email;
    return this;
  }

   /**
   * Email address
   * @return email
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Email address")

  public String getEmail() {
    return email;
  }


  public void setEmail(String email) {
    this.email = email;
  }


  public ApiUsersIdUser password(String password) {
    
    this.password = password;
    return this;
  }

   /**
   * Password
   * @return password
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Password")

  public String getPassword() {
    return password;
  }


  public void setPassword(String password) {
    this.password = password;
  }


  public ApiUsersIdUser roles(List<ReferenceObject> roles) {
    
    this.roles = roles;
    return this;
  }

  public ApiUsersIdUser addRolesItem(ReferenceObject rolesItem) {
    if (this.roles == null) {
      this.roles = new ArrayList<ReferenceObject>();
    }
    this.roles.add(rolesItem);
    return this;
  }

   /**
   * List of Roles
   * @return roles
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "List of Roles")

  public List<ReferenceObject> getRoles() {
    return roles;
  }


  public void setRoles(List<ReferenceObject> roles) {
    this.roles = roles;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiUsersIdUser apiUsersIdUser = (ApiUsersIdUser) o;
    return Objects.equals(this.firstName, apiUsersIdUser.firstName) &&
        Objects.equals(this.lastName, apiUsersIdUser.lastName) &&
        Objects.equals(this.username, apiUsersIdUser.username) &&
        Objects.equals(this.email, apiUsersIdUser.email) &&
        Objects.equals(this.password, apiUsersIdUser.password) &&
        Objects.equals(this.roles, apiUsersIdUser.roles);
  }

  @Override
  public int hashCode() {
    return Objects.hash(firstName, lastName, username, email, password, roles);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiUsersIdUser {\n");
    sb.append("    firstName: ").append(toIndentedString(firstName)).append("\n");
    sb.append("    lastName: ").append(toIndentedString(lastName)).append("\n");
    sb.append("    username: ").append(toIndentedString(username)).append("\n");
    sb.append("    email: ").append(toIndentedString(email)).append("\n");
    sb.append("    password: ").append(toIndentedString(password)).append("\n");
    sb.append("    roles: ").append(toIndentedString(roles)).append("\n");
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

