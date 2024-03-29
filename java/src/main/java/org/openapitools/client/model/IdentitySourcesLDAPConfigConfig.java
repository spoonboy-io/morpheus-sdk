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
 * IdentitySourcesLDAPConfigConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class IdentitySourcesLDAPConfigConfig {
  public static final String SERIALIZED_NAME_URL = "url";
  @SerializedName(SERIALIZED_NAME_URL)
  private String url;

  public static final String SERIALIZED_NAME_BINDING_USERNAME = "bindingUsername";
  @SerializedName(SERIALIZED_NAME_BINDING_USERNAME)
  private String bindingUsername;

  public static final String SERIALIZED_NAME_BINDING_PASSWORD = "bindingPassword";
  @SerializedName(SERIALIZED_NAME_BINDING_PASSWORD)
  private String bindingPassword;

  public static final String SERIALIZED_NAME_USER_FQN_EXPRESSION = "userFqnExpression";
  @SerializedName(SERIALIZED_NAME_USER_FQN_EXPRESSION)
  private String userFqnExpression;

  public static final String SERIALIZED_NAME_REQUIRED_ROLE_FQN = "requiredRoleFqn";
  @SerializedName(SERIALIZED_NAME_REQUIRED_ROLE_FQN)
  private String requiredRoleFqn;

  public static final String SERIALIZED_NAME_USERNAME_ATTRIBUTE = "usernameAttribute";
  @SerializedName(SERIALIZED_NAME_USERNAME_ATTRIBUTE)
  private String usernameAttribute;

  public static final String SERIALIZED_NAME_COMMON_NAME_ATTRIBUTE = "commonNameAttribute";
  @SerializedName(SERIALIZED_NAME_COMMON_NAME_ATTRIBUTE)
  private String commonNameAttribute;

  public static final String SERIALIZED_NAME_FIRST_NAME_ATTRIBUTE = "firstNameAttribute";
  @SerializedName(SERIALIZED_NAME_FIRST_NAME_ATTRIBUTE)
  private String firstNameAttribute;

  public static final String SERIALIZED_NAME_LAST_NAME_ATTRIBUTE = "lastNameAttribute";
  @SerializedName(SERIALIZED_NAME_LAST_NAME_ATTRIBUTE)
  private String lastNameAttribute;

  public static final String SERIALIZED_NAME_EMAIL_ATTRIBUTE = "emailAttribute";
  @SerializedName(SERIALIZED_NAME_EMAIL_ATTRIBUTE)
  private String emailAttribute;

  public static final String SERIALIZED_NAME_UNIQUE_MEMBER_ATTRIBUTE = "uniqueMemberAttribute";
  @SerializedName(SERIALIZED_NAME_UNIQUE_MEMBER_ATTRIBUTE)
  private String uniqueMemberAttribute;

  public static final String SERIALIZED_NAME_MEMBER_OF_ATTRIBUTE = "memberOfAttribute";
  @SerializedName(SERIALIZED_NAME_MEMBER_OF_ATTRIBUTE)
  private String memberOfAttribute;

  public static final String SERIALIZED_NAME_BINDING_PASSWORD_HASH = "bindingPasswordHash";
  @SerializedName(SERIALIZED_NAME_BINDING_PASSWORD_HASH)
  private String bindingPasswordHash;


  public IdentitySourcesLDAPConfigConfig url(String url) {
    
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


  public IdentitySourcesLDAPConfigConfig bindingUsername(String bindingUsername) {
    
    this.bindingUsername = bindingUsername;
    return this;
  }

   /**
   * Get bindingUsername
   * @return bindingUsername
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBindingUsername() {
    return bindingUsername;
  }


  public void setBindingUsername(String bindingUsername) {
    this.bindingUsername = bindingUsername;
  }


  public IdentitySourcesLDAPConfigConfig bindingPassword(String bindingPassword) {
    
    this.bindingPassword = bindingPassword;
    return this;
  }

   /**
   * Get bindingPassword
   * @return bindingPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBindingPassword() {
    return bindingPassword;
  }


  public void setBindingPassword(String bindingPassword) {
    this.bindingPassword = bindingPassword;
  }


  public IdentitySourcesLDAPConfigConfig userFqnExpression(String userFqnExpression) {
    
    this.userFqnExpression = userFqnExpression;
    return this;
  }

   /**
   * Get userFqnExpression
   * @return userFqnExpression
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUserFqnExpression() {
    return userFqnExpression;
  }


  public void setUserFqnExpression(String userFqnExpression) {
    this.userFqnExpression = userFqnExpression;
  }


  public IdentitySourcesLDAPConfigConfig requiredRoleFqn(String requiredRoleFqn) {
    
    this.requiredRoleFqn = requiredRoleFqn;
    return this;
  }

   /**
   * Get requiredRoleFqn
   * @return requiredRoleFqn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRequiredRoleFqn() {
    return requiredRoleFqn;
  }


  public void setRequiredRoleFqn(String requiredRoleFqn) {
    this.requiredRoleFqn = requiredRoleFqn;
  }


  public IdentitySourcesLDAPConfigConfig usernameAttribute(String usernameAttribute) {
    
    this.usernameAttribute = usernameAttribute;
    return this;
  }

   /**
   * Get usernameAttribute
   * @return usernameAttribute
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUsernameAttribute() {
    return usernameAttribute;
  }


  public void setUsernameAttribute(String usernameAttribute) {
    this.usernameAttribute = usernameAttribute;
  }


  public IdentitySourcesLDAPConfigConfig commonNameAttribute(String commonNameAttribute) {
    
    this.commonNameAttribute = commonNameAttribute;
    return this;
  }

   /**
   * Get commonNameAttribute
   * @return commonNameAttribute
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCommonNameAttribute() {
    return commonNameAttribute;
  }


  public void setCommonNameAttribute(String commonNameAttribute) {
    this.commonNameAttribute = commonNameAttribute;
  }


  public IdentitySourcesLDAPConfigConfig firstNameAttribute(String firstNameAttribute) {
    
    this.firstNameAttribute = firstNameAttribute;
    return this;
  }

   /**
   * Get firstNameAttribute
   * @return firstNameAttribute
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFirstNameAttribute() {
    return firstNameAttribute;
  }


  public void setFirstNameAttribute(String firstNameAttribute) {
    this.firstNameAttribute = firstNameAttribute;
  }


  public IdentitySourcesLDAPConfigConfig lastNameAttribute(String lastNameAttribute) {
    
    this.lastNameAttribute = lastNameAttribute;
    return this;
  }

   /**
   * Get lastNameAttribute
   * @return lastNameAttribute
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLastNameAttribute() {
    return lastNameAttribute;
  }


  public void setLastNameAttribute(String lastNameAttribute) {
    this.lastNameAttribute = lastNameAttribute;
  }


  public IdentitySourcesLDAPConfigConfig emailAttribute(String emailAttribute) {
    
    this.emailAttribute = emailAttribute;
    return this;
  }

   /**
   * Get emailAttribute
   * @return emailAttribute
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEmailAttribute() {
    return emailAttribute;
  }


  public void setEmailAttribute(String emailAttribute) {
    this.emailAttribute = emailAttribute;
  }


  public IdentitySourcesLDAPConfigConfig uniqueMemberAttribute(String uniqueMemberAttribute) {
    
    this.uniqueMemberAttribute = uniqueMemberAttribute;
    return this;
  }

   /**
   * Get uniqueMemberAttribute
   * @return uniqueMemberAttribute
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUniqueMemberAttribute() {
    return uniqueMemberAttribute;
  }


  public void setUniqueMemberAttribute(String uniqueMemberAttribute) {
    this.uniqueMemberAttribute = uniqueMemberAttribute;
  }


  public IdentitySourcesLDAPConfigConfig memberOfAttribute(String memberOfAttribute) {
    
    this.memberOfAttribute = memberOfAttribute;
    return this;
  }

   /**
   * Get memberOfAttribute
   * @return memberOfAttribute
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMemberOfAttribute() {
    return memberOfAttribute;
  }


  public void setMemberOfAttribute(String memberOfAttribute) {
    this.memberOfAttribute = memberOfAttribute;
  }


  public IdentitySourcesLDAPConfigConfig bindingPasswordHash(String bindingPasswordHash) {
    
    this.bindingPasswordHash = bindingPasswordHash;
    return this;
  }

   /**
   * Get bindingPasswordHash
   * @return bindingPasswordHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getBindingPasswordHash() {
    return bindingPasswordHash;
  }


  public void setBindingPasswordHash(String bindingPasswordHash) {
    this.bindingPasswordHash = bindingPasswordHash;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    IdentitySourcesLDAPConfigConfig identitySourcesLDAPConfigConfig = (IdentitySourcesLDAPConfigConfig) o;
    return Objects.equals(this.url, identitySourcesLDAPConfigConfig.url) &&
        Objects.equals(this.bindingUsername, identitySourcesLDAPConfigConfig.bindingUsername) &&
        Objects.equals(this.bindingPassword, identitySourcesLDAPConfigConfig.bindingPassword) &&
        Objects.equals(this.userFqnExpression, identitySourcesLDAPConfigConfig.userFqnExpression) &&
        Objects.equals(this.requiredRoleFqn, identitySourcesLDAPConfigConfig.requiredRoleFqn) &&
        Objects.equals(this.usernameAttribute, identitySourcesLDAPConfigConfig.usernameAttribute) &&
        Objects.equals(this.commonNameAttribute, identitySourcesLDAPConfigConfig.commonNameAttribute) &&
        Objects.equals(this.firstNameAttribute, identitySourcesLDAPConfigConfig.firstNameAttribute) &&
        Objects.equals(this.lastNameAttribute, identitySourcesLDAPConfigConfig.lastNameAttribute) &&
        Objects.equals(this.emailAttribute, identitySourcesLDAPConfigConfig.emailAttribute) &&
        Objects.equals(this.uniqueMemberAttribute, identitySourcesLDAPConfigConfig.uniqueMemberAttribute) &&
        Objects.equals(this.memberOfAttribute, identitySourcesLDAPConfigConfig.memberOfAttribute) &&
        Objects.equals(this.bindingPasswordHash, identitySourcesLDAPConfigConfig.bindingPasswordHash);
  }

  @Override
  public int hashCode() {
    return Objects.hash(url, bindingUsername, bindingPassword, userFqnExpression, requiredRoleFqn, usernameAttribute, commonNameAttribute, firstNameAttribute, lastNameAttribute, emailAttribute, uniqueMemberAttribute, memberOfAttribute, bindingPasswordHash);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class IdentitySourcesLDAPConfigConfig {\n");
    sb.append("    url: ").append(toIndentedString(url)).append("\n");
    sb.append("    bindingUsername: ").append(toIndentedString(bindingUsername)).append("\n");
    sb.append("    bindingPassword: ").append(toIndentedString(bindingPassword)).append("\n");
    sb.append("    userFqnExpression: ").append(toIndentedString(userFqnExpression)).append("\n");
    sb.append("    requiredRoleFqn: ").append(toIndentedString(requiredRoleFqn)).append("\n");
    sb.append("    usernameAttribute: ").append(toIndentedString(usernameAttribute)).append("\n");
    sb.append("    commonNameAttribute: ").append(toIndentedString(commonNameAttribute)).append("\n");
    sb.append("    firstNameAttribute: ").append(toIndentedString(firstNameAttribute)).append("\n");
    sb.append("    lastNameAttribute: ").append(toIndentedString(lastNameAttribute)).append("\n");
    sb.append("    emailAttribute: ").append(toIndentedString(emailAttribute)).append("\n");
    sb.append("    uniqueMemberAttribute: ").append(toIndentedString(uniqueMemberAttribute)).append("\n");
    sb.append("    memberOfAttribute: ").append(toIndentedString(memberOfAttribute)).append("\n");
    sb.append("    bindingPasswordHash: ").append(toIndentedString(bindingPasswordHash)).append("\n");
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

