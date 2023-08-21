<?php
/**
 * IdentitySourcesAzureADConfigConfig
 *
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.0.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * IdentitySourcesAzureADConfigConfig Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class IdentitySourcesAzureADConfigConfig implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'identitySourcesAzureADConfig_config';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'url' => 'string',
        'logout_url' => 'string',
        'do_not_include_saml_request' => 'bool',
        'saml_signature_mode' => 'string',
        'do_not_validate_signature' => 'bool',
        'do_not_validate_status_code' => 'bool',
        'do_not_validate_destination' => 'bool',
        'do_not_validate_issue_instants' => 'bool',
        'do_not_validate_assertions' => 'bool',
        'given_name_attribute' => 'string',
        'surname_attribute' => 'string',
        'email_attribute' => 'string',
        'required_attribute_value' => 'string',
        'role_attribute_name' => 'string',
        'azure_tenant_id' => 'string',
        'azure_app_id' => 'string',
        'azure_app_secret' => 'string',
        'role_link_attribute_name' => 'string',
        'public_key' => 'string',
        'azure_app_secret_hash' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'url' => null,
        'logout_url' => null,
        'do_not_include_saml_request' => null,
        'saml_signature_mode' => null,
        'do_not_validate_signature' => null,
        'do_not_validate_status_code' => null,
        'do_not_validate_destination' => null,
        'do_not_validate_issue_instants' => null,
        'do_not_validate_assertions' => null,
        'given_name_attribute' => null,
        'surname_attribute' => null,
        'email_attribute' => null,
        'required_attribute_value' => null,
        'role_attribute_name' => null,
        'azure_tenant_id' => null,
        'azure_app_id' => null,
        'azure_app_secret' => null,
        'role_link_attribute_name' => null,
        'public_key' => null,
        'azure_app_secret_hash' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'url' => 'url',
        'logout_url' => 'logoutUrl',
        'do_not_include_saml_request' => 'doNotIncludeSAMLRequest',
        'saml_signature_mode' => 'SAMLSignatureMode',
        'do_not_validate_signature' => 'doNotValidateSignature',
        'do_not_validate_status_code' => 'doNotValidateStatusCode',
        'do_not_validate_destination' => 'doNotValidateDestination',
        'do_not_validate_issue_instants' => 'doNotValidateIssueInstants',
        'do_not_validate_assertions' => 'doNotValidateAssertions',
        'given_name_attribute' => 'givenNameAttribute',
        'surname_attribute' => 'surnameAttribute',
        'email_attribute' => 'emailAttribute',
        'required_attribute_value' => 'requiredAttributeValue',
        'role_attribute_name' => 'roleAttributeName',
        'azure_tenant_id' => 'azureTenantId',
        'azure_app_id' => 'azureAppId',
        'azure_app_secret' => 'azureAppSecret',
        'role_link_attribute_name' => 'roleLinkAttributeName',
        'public_key' => 'publicKey',
        'azure_app_secret_hash' => 'azureAppSecretHash'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'url' => 'setUrl',
        'logout_url' => 'setLogoutUrl',
        'do_not_include_saml_request' => 'setDoNotIncludeSamlRequest',
        'saml_signature_mode' => 'setSamlSignatureMode',
        'do_not_validate_signature' => 'setDoNotValidateSignature',
        'do_not_validate_status_code' => 'setDoNotValidateStatusCode',
        'do_not_validate_destination' => 'setDoNotValidateDestination',
        'do_not_validate_issue_instants' => 'setDoNotValidateIssueInstants',
        'do_not_validate_assertions' => 'setDoNotValidateAssertions',
        'given_name_attribute' => 'setGivenNameAttribute',
        'surname_attribute' => 'setSurnameAttribute',
        'email_attribute' => 'setEmailAttribute',
        'required_attribute_value' => 'setRequiredAttributeValue',
        'role_attribute_name' => 'setRoleAttributeName',
        'azure_tenant_id' => 'setAzureTenantId',
        'azure_app_id' => 'setAzureAppId',
        'azure_app_secret' => 'setAzureAppSecret',
        'role_link_attribute_name' => 'setRoleLinkAttributeName',
        'public_key' => 'setPublicKey',
        'azure_app_secret_hash' => 'setAzureAppSecretHash'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'url' => 'getUrl',
        'logout_url' => 'getLogoutUrl',
        'do_not_include_saml_request' => 'getDoNotIncludeSamlRequest',
        'saml_signature_mode' => 'getSamlSignatureMode',
        'do_not_validate_signature' => 'getDoNotValidateSignature',
        'do_not_validate_status_code' => 'getDoNotValidateStatusCode',
        'do_not_validate_destination' => 'getDoNotValidateDestination',
        'do_not_validate_issue_instants' => 'getDoNotValidateIssueInstants',
        'do_not_validate_assertions' => 'getDoNotValidateAssertions',
        'given_name_attribute' => 'getGivenNameAttribute',
        'surname_attribute' => 'getSurnameAttribute',
        'email_attribute' => 'getEmailAttribute',
        'required_attribute_value' => 'getRequiredAttributeValue',
        'role_attribute_name' => 'getRoleAttributeName',
        'azure_tenant_id' => 'getAzureTenantId',
        'azure_app_id' => 'getAzureAppId',
        'azure_app_secret' => 'getAzureAppSecret',
        'role_link_attribute_name' => 'getRoleLinkAttributeName',
        'public_key' => 'getPublicKey',
        'azure_app_secret_hash' => 'getAzureAppSecretHash'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    

    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['url'] = $data['url'] ?? null;
        $this->container['logout_url'] = $data['logout_url'] ?? null;
        $this->container['do_not_include_saml_request'] = $data['do_not_include_saml_request'] ?? null;
        $this->container['saml_signature_mode'] = $data['saml_signature_mode'] ?? null;
        $this->container['do_not_validate_signature'] = $data['do_not_validate_signature'] ?? null;
        $this->container['do_not_validate_status_code'] = $data['do_not_validate_status_code'] ?? null;
        $this->container['do_not_validate_destination'] = $data['do_not_validate_destination'] ?? null;
        $this->container['do_not_validate_issue_instants'] = $data['do_not_validate_issue_instants'] ?? null;
        $this->container['do_not_validate_assertions'] = $data['do_not_validate_assertions'] ?? null;
        $this->container['given_name_attribute'] = $data['given_name_attribute'] ?? null;
        $this->container['surname_attribute'] = $data['surname_attribute'] ?? null;
        $this->container['email_attribute'] = $data['email_attribute'] ?? null;
        $this->container['required_attribute_value'] = $data['required_attribute_value'] ?? null;
        $this->container['role_attribute_name'] = $data['role_attribute_name'] ?? null;
        $this->container['azure_tenant_id'] = $data['azure_tenant_id'] ?? null;
        $this->container['azure_app_id'] = $data['azure_app_id'] ?? null;
        $this->container['azure_app_secret'] = $data['azure_app_secret'] ?? null;
        $this->container['role_link_attribute_name'] = $data['role_link_attribute_name'] ?? null;
        $this->container['public_key'] = $data['public_key'] ?? null;
        $this->container['azure_app_secret_hash'] = $data['azure_app_secret_hash'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets url
     *
     * @return string|null
     */
    public function getUrl()
    {
        return $this->container['url'];
    }

    /**
     * Sets url
     *
     * @param string|null $url url
     *
     * @return self
     */
    public function setUrl($url)
    {
        $this->container['url'] = $url;

        return $this;
    }

    /**
     * Gets logout_url
     *
     * @return string|null
     */
    public function getLogoutUrl()
    {
        return $this->container['logout_url'];
    }

    /**
     * Sets logout_url
     *
     * @param string|null $logout_url logout_url
     *
     * @return self
     */
    public function setLogoutUrl($logout_url)
    {
        $this->container['logout_url'] = $logout_url;

        return $this;
    }

    /**
     * Gets do_not_include_saml_request
     *
     * @return bool|null
     */
    public function getDoNotIncludeSamlRequest()
    {
        return $this->container['do_not_include_saml_request'];
    }

    /**
     * Sets do_not_include_saml_request
     *
     * @param bool|null $do_not_include_saml_request do_not_include_saml_request
     *
     * @return self
     */
    public function setDoNotIncludeSamlRequest($do_not_include_saml_request)
    {
        $this->container['do_not_include_saml_request'] = $do_not_include_saml_request;

        return $this;
    }

    /**
     * Gets saml_signature_mode
     *
     * @return string|null
     */
    public function getSamlSignatureMode()
    {
        return $this->container['saml_signature_mode'];
    }

    /**
     * Sets saml_signature_mode
     *
     * @param string|null $saml_signature_mode saml_signature_mode
     *
     * @return self
     */
    public function setSamlSignatureMode($saml_signature_mode)
    {
        $this->container['saml_signature_mode'] = $saml_signature_mode;

        return $this;
    }

    /**
     * Gets do_not_validate_signature
     *
     * @return bool|null
     */
    public function getDoNotValidateSignature()
    {
        return $this->container['do_not_validate_signature'];
    }

    /**
     * Sets do_not_validate_signature
     *
     * @param bool|null $do_not_validate_signature do_not_validate_signature
     *
     * @return self
     */
    public function setDoNotValidateSignature($do_not_validate_signature)
    {
        $this->container['do_not_validate_signature'] = $do_not_validate_signature;

        return $this;
    }

    /**
     * Gets do_not_validate_status_code
     *
     * @return bool|null
     */
    public function getDoNotValidateStatusCode()
    {
        return $this->container['do_not_validate_status_code'];
    }

    /**
     * Sets do_not_validate_status_code
     *
     * @param bool|null $do_not_validate_status_code do_not_validate_status_code
     *
     * @return self
     */
    public function setDoNotValidateStatusCode($do_not_validate_status_code)
    {
        $this->container['do_not_validate_status_code'] = $do_not_validate_status_code;

        return $this;
    }

    /**
     * Gets do_not_validate_destination
     *
     * @return bool|null
     */
    public function getDoNotValidateDestination()
    {
        return $this->container['do_not_validate_destination'];
    }

    /**
     * Sets do_not_validate_destination
     *
     * @param bool|null $do_not_validate_destination do_not_validate_destination
     *
     * @return self
     */
    public function setDoNotValidateDestination($do_not_validate_destination)
    {
        $this->container['do_not_validate_destination'] = $do_not_validate_destination;

        return $this;
    }

    /**
     * Gets do_not_validate_issue_instants
     *
     * @return bool|null
     */
    public function getDoNotValidateIssueInstants()
    {
        return $this->container['do_not_validate_issue_instants'];
    }

    /**
     * Sets do_not_validate_issue_instants
     *
     * @param bool|null $do_not_validate_issue_instants do_not_validate_issue_instants
     *
     * @return self
     */
    public function setDoNotValidateIssueInstants($do_not_validate_issue_instants)
    {
        $this->container['do_not_validate_issue_instants'] = $do_not_validate_issue_instants;

        return $this;
    }

    /**
     * Gets do_not_validate_assertions
     *
     * @return bool|null
     */
    public function getDoNotValidateAssertions()
    {
        return $this->container['do_not_validate_assertions'];
    }

    /**
     * Sets do_not_validate_assertions
     *
     * @param bool|null $do_not_validate_assertions do_not_validate_assertions
     *
     * @return self
     */
    public function setDoNotValidateAssertions($do_not_validate_assertions)
    {
        $this->container['do_not_validate_assertions'] = $do_not_validate_assertions;

        return $this;
    }

    /**
     * Gets given_name_attribute
     *
     * @return string|null
     */
    public function getGivenNameAttribute()
    {
        return $this->container['given_name_attribute'];
    }

    /**
     * Sets given_name_attribute
     *
     * @param string|null $given_name_attribute given_name_attribute
     *
     * @return self
     */
    public function setGivenNameAttribute($given_name_attribute)
    {
        $this->container['given_name_attribute'] = $given_name_attribute;

        return $this;
    }

    /**
     * Gets surname_attribute
     *
     * @return string|null
     */
    public function getSurnameAttribute()
    {
        return $this->container['surname_attribute'];
    }

    /**
     * Sets surname_attribute
     *
     * @param string|null $surname_attribute surname_attribute
     *
     * @return self
     */
    public function setSurnameAttribute($surname_attribute)
    {
        $this->container['surname_attribute'] = $surname_attribute;

        return $this;
    }

    /**
     * Gets email_attribute
     *
     * @return string|null
     */
    public function getEmailAttribute()
    {
        return $this->container['email_attribute'];
    }

    /**
     * Sets email_attribute
     *
     * @param string|null $email_attribute email_attribute
     *
     * @return self
     */
    public function setEmailAttribute($email_attribute)
    {
        $this->container['email_attribute'] = $email_attribute;

        return $this;
    }

    /**
     * Gets required_attribute_value
     *
     * @return string|null
     */
    public function getRequiredAttributeValue()
    {
        return $this->container['required_attribute_value'];
    }

    /**
     * Sets required_attribute_value
     *
     * @param string|null $required_attribute_value required_attribute_value
     *
     * @return self
     */
    public function setRequiredAttributeValue($required_attribute_value)
    {
        $this->container['required_attribute_value'] = $required_attribute_value;

        return $this;
    }

    /**
     * Gets role_attribute_name
     *
     * @return string|null
     */
    public function getRoleAttributeName()
    {
        return $this->container['role_attribute_name'];
    }

    /**
     * Sets role_attribute_name
     *
     * @param string|null $role_attribute_name role_attribute_name
     *
     * @return self
     */
    public function setRoleAttributeName($role_attribute_name)
    {
        $this->container['role_attribute_name'] = $role_attribute_name;

        return $this;
    }

    /**
     * Gets azure_tenant_id
     *
     * @return string|null
     */
    public function getAzureTenantId()
    {
        return $this->container['azure_tenant_id'];
    }

    /**
     * Sets azure_tenant_id
     *
     * @param string|null $azure_tenant_id azure_tenant_id
     *
     * @return self
     */
    public function setAzureTenantId($azure_tenant_id)
    {
        $this->container['azure_tenant_id'] = $azure_tenant_id;

        return $this;
    }

    /**
     * Gets azure_app_id
     *
     * @return string|null
     */
    public function getAzureAppId()
    {
        return $this->container['azure_app_id'];
    }

    /**
     * Sets azure_app_id
     *
     * @param string|null $azure_app_id azure_app_id
     *
     * @return self
     */
    public function setAzureAppId($azure_app_id)
    {
        $this->container['azure_app_id'] = $azure_app_id;

        return $this;
    }

    /**
     * Gets azure_app_secret
     *
     * @return string|null
     */
    public function getAzureAppSecret()
    {
        return $this->container['azure_app_secret'];
    }

    /**
     * Sets azure_app_secret
     *
     * @param string|null $azure_app_secret azure_app_secret
     *
     * @return self
     */
    public function setAzureAppSecret($azure_app_secret)
    {
        $this->container['azure_app_secret'] = $azure_app_secret;

        return $this;
    }

    /**
     * Gets role_link_attribute_name
     *
     * @return string|null
     */
    public function getRoleLinkAttributeName()
    {
        return $this->container['role_link_attribute_name'];
    }

    /**
     * Sets role_link_attribute_name
     *
     * @param string|null $role_link_attribute_name role_link_attribute_name
     *
     * @return self
     */
    public function setRoleLinkAttributeName($role_link_attribute_name)
    {
        $this->container['role_link_attribute_name'] = $role_link_attribute_name;

        return $this;
    }

    /**
     * Gets public_key
     *
     * @return string|null
     */
    public function getPublicKey()
    {
        return $this->container['public_key'];
    }

    /**
     * Sets public_key
     *
     * @param string|null $public_key public_key
     *
     * @return self
     */
    public function setPublicKey($public_key)
    {
        $this->container['public_key'] = $public_key;

        return $this;
    }

    /**
     * Gets azure_app_secret_hash
     *
     * @return string|null
     */
    public function getAzureAppSecretHash()
    {
        return $this->container['azure_app_secret_hash'];
    }

    /**
     * Sets azure_app_secret_hash
     *
     * @param string|null $azure_app_secret_hash azure_app_secret_hash
     *
     * @return self
     */
    public function setAzureAppSecretHash($azure_app_secret_hash)
    {
        $this->container['azure_app_secret_hash'] = $azure_app_secret_hash;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


