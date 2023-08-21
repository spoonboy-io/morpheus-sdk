<?php
/**
 * IdentitySourcesSAMLConfigConfig
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
 * IdentitySourcesSAMLConfigConfig Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class IdentitySourcesSAMLConfigConfig implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'identitySourcesSAMLConfig_config';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'role_attribute_name' => 'string',
        'required_attribute_value' => 'string',
        'given_name_attribute' => 'string',
        'surname_attribute' => 'string',
        'logout_url' => 'string',
        'do_not_include_saml_request' => 'bool',
        'public_key' => 'string',
        'email_attribute' => 'string',
        'url' => 'string',
        'do_not_validate_signature' => 'bool',
        'do_not_validate_status_code' => 'bool',
        'do_not_validate_destination' => 'bool',
        'do_not_validate_issue_instants' => 'bool',
        'do_not_validate_assertions' => 'bool',
        'do_not_validate_auth_statements' => 'bool',
        'do_not_validate_subject' => 'bool',
        'do_not_validate_conditions' => 'bool',
        'do_not_validate_audiences' => 'bool',
        'do_not_validate_subject_recipients' => 'bool',
        'saml_signature_mode' => 'string',
        'force_authn' => 'bool'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'role_attribute_name' => null,
        'required_attribute_value' => null,
        'given_name_attribute' => null,
        'surname_attribute' => null,
        'logout_url' => null,
        'do_not_include_saml_request' => null,
        'public_key' => null,
        'email_attribute' => null,
        'url' => null,
        'do_not_validate_signature' => null,
        'do_not_validate_status_code' => null,
        'do_not_validate_destination' => null,
        'do_not_validate_issue_instants' => null,
        'do_not_validate_assertions' => null,
        'do_not_validate_auth_statements' => null,
        'do_not_validate_subject' => null,
        'do_not_validate_conditions' => null,
        'do_not_validate_audiences' => null,
        'do_not_validate_subject_recipients' => null,
        'saml_signature_mode' => null,
        'force_authn' => null
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
        'role_attribute_name' => 'roleAttributeName',
        'required_attribute_value' => 'requiredAttributeValue',
        'given_name_attribute' => 'givenNameAttribute',
        'surname_attribute' => 'surnameAttribute',
        'logout_url' => 'logoutUrl',
        'do_not_include_saml_request' => 'doNotIncludeSAMLRequest',
        'public_key' => 'publicKey',
        'email_attribute' => 'emailAttribute',
        'url' => 'url',
        'do_not_validate_signature' => 'doNotValidateSignature',
        'do_not_validate_status_code' => 'doNotValidateStatusCode',
        'do_not_validate_destination' => 'doNotValidateDestination',
        'do_not_validate_issue_instants' => 'doNotValidateIssueInstants',
        'do_not_validate_assertions' => 'doNotValidateAssertions',
        'do_not_validate_auth_statements' => 'doNotValidateAuthStatements',
        'do_not_validate_subject' => 'doNotValidateSubject',
        'do_not_validate_conditions' => 'doNotValidateConditions',
        'do_not_validate_audiences' => 'doNotValidateAudiences',
        'do_not_validate_subject_recipients' => 'doNotValidateSubjectRecipients',
        'saml_signature_mode' => 'SAMLSignatureMode',
        'force_authn' => 'forceAuthn'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'role_attribute_name' => 'setRoleAttributeName',
        'required_attribute_value' => 'setRequiredAttributeValue',
        'given_name_attribute' => 'setGivenNameAttribute',
        'surname_attribute' => 'setSurnameAttribute',
        'logout_url' => 'setLogoutUrl',
        'do_not_include_saml_request' => 'setDoNotIncludeSamlRequest',
        'public_key' => 'setPublicKey',
        'email_attribute' => 'setEmailAttribute',
        'url' => 'setUrl',
        'do_not_validate_signature' => 'setDoNotValidateSignature',
        'do_not_validate_status_code' => 'setDoNotValidateStatusCode',
        'do_not_validate_destination' => 'setDoNotValidateDestination',
        'do_not_validate_issue_instants' => 'setDoNotValidateIssueInstants',
        'do_not_validate_assertions' => 'setDoNotValidateAssertions',
        'do_not_validate_auth_statements' => 'setDoNotValidateAuthStatements',
        'do_not_validate_subject' => 'setDoNotValidateSubject',
        'do_not_validate_conditions' => 'setDoNotValidateConditions',
        'do_not_validate_audiences' => 'setDoNotValidateAudiences',
        'do_not_validate_subject_recipients' => 'setDoNotValidateSubjectRecipients',
        'saml_signature_mode' => 'setSamlSignatureMode',
        'force_authn' => 'setForceAuthn'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'role_attribute_name' => 'getRoleAttributeName',
        'required_attribute_value' => 'getRequiredAttributeValue',
        'given_name_attribute' => 'getGivenNameAttribute',
        'surname_attribute' => 'getSurnameAttribute',
        'logout_url' => 'getLogoutUrl',
        'do_not_include_saml_request' => 'getDoNotIncludeSamlRequest',
        'public_key' => 'getPublicKey',
        'email_attribute' => 'getEmailAttribute',
        'url' => 'getUrl',
        'do_not_validate_signature' => 'getDoNotValidateSignature',
        'do_not_validate_status_code' => 'getDoNotValidateStatusCode',
        'do_not_validate_destination' => 'getDoNotValidateDestination',
        'do_not_validate_issue_instants' => 'getDoNotValidateIssueInstants',
        'do_not_validate_assertions' => 'getDoNotValidateAssertions',
        'do_not_validate_auth_statements' => 'getDoNotValidateAuthStatements',
        'do_not_validate_subject' => 'getDoNotValidateSubject',
        'do_not_validate_conditions' => 'getDoNotValidateConditions',
        'do_not_validate_audiences' => 'getDoNotValidateAudiences',
        'do_not_validate_subject_recipients' => 'getDoNotValidateSubjectRecipients',
        'saml_signature_mode' => 'getSamlSignatureMode',
        'force_authn' => 'getForceAuthn'
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
        $this->container['role_attribute_name'] = $data['role_attribute_name'] ?? null;
        $this->container['required_attribute_value'] = $data['required_attribute_value'] ?? null;
        $this->container['given_name_attribute'] = $data['given_name_attribute'] ?? null;
        $this->container['surname_attribute'] = $data['surname_attribute'] ?? null;
        $this->container['logout_url'] = $data['logout_url'] ?? null;
        $this->container['do_not_include_saml_request'] = $data['do_not_include_saml_request'] ?? null;
        $this->container['public_key'] = $data['public_key'] ?? null;
        $this->container['email_attribute'] = $data['email_attribute'] ?? null;
        $this->container['url'] = $data['url'] ?? null;
        $this->container['do_not_validate_signature'] = $data['do_not_validate_signature'] ?? null;
        $this->container['do_not_validate_status_code'] = $data['do_not_validate_status_code'] ?? null;
        $this->container['do_not_validate_destination'] = $data['do_not_validate_destination'] ?? null;
        $this->container['do_not_validate_issue_instants'] = $data['do_not_validate_issue_instants'] ?? null;
        $this->container['do_not_validate_assertions'] = $data['do_not_validate_assertions'] ?? null;
        $this->container['do_not_validate_auth_statements'] = $data['do_not_validate_auth_statements'] ?? null;
        $this->container['do_not_validate_subject'] = $data['do_not_validate_subject'] ?? null;
        $this->container['do_not_validate_conditions'] = $data['do_not_validate_conditions'] ?? null;
        $this->container['do_not_validate_audiences'] = $data['do_not_validate_audiences'] ?? null;
        $this->container['do_not_validate_subject_recipients'] = $data['do_not_validate_subject_recipients'] ?? null;
        $this->container['saml_signature_mode'] = $data['saml_signature_mode'] ?? null;
        $this->container['force_authn'] = $data['force_authn'] ?? null;
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
     * Gets do_not_validate_auth_statements
     *
     * @return bool|null
     */
    public function getDoNotValidateAuthStatements()
    {
        return $this->container['do_not_validate_auth_statements'];
    }

    /**
     * Sets do_not_validate_auth_statements
     *
     * @param bool|null $do_not_validate_auth_statements do_not_validate_auth_statements
     *
     * @return self
     */
    public function setDoNotValidateAuthStatements($do_not_validate_auth_statements)
    {
        $this->container['do_not_validate_auth_statements'] = $do_not_validate_auth_statements;

        return $this;
    }

    /**
     * Gets do_not_validate_subject
     *
     * @return bool|null
     */
    public function getDoNotValidateSubject()
    {
        return $this->container['do_not_validate_subject'];
    }

    /**
     * Sets do_not_validate_subject
     *
     * @param bool|null $do_not_validate_subject do_not_validate_subject
     *
     * @return self
     */
    public function setDoNotValidateSubject($do_not_validate_subject)
    {
        $this->container['do_not_validate_subject'] = $do_not_validate_subject;

        return $this;
    }

    /**
     * Gets do_not_validate_conditions
     *
     * @return bool|null
     */
    public function getDoNotValidateConditions()
    {
        return $this->container['do_not_validate_conditions'];
    }

    /**
     * Sets do_not_validate_conditions
     *
     * @param bool|null $do_not_validate_conditions do_not_validate_conditions
     *
     * @return self
     */
    public function setDoNotValidateConditions($do_not_validate_conditions)
    {
        $this->container['do_not_validate_conditions'] = $do_not_validate_conditions;

        return $this;
    }

    /**
     * Gets do_not_validate_audiences
     *
     * @return bool|null
     */
    public function getDoNotValidateAudiences()
    {
        return $this->container['do_not_validate_audiences'];
    }

    /**
     * Sets do_not_validate_audiences
     *
     * @param bool|null $do_not_validate_audiences do_not_validate_audiences
     *
     * @return self
     */
    public function setDoNotValidateAudiences($do_not_validate_audiences)
    {
        $this->container['do_not_validate_audiences'] = $do_not_validate_audiences;

        return $this;
    }

    /**
     * Gets do_not_validate_subject_recipients
     *
     * @return bool|null
     */
    public function getDoNotValidateSubjectRecipients()
    {
        return $this->container['do_not_validate_subject_recipients'];
    }

    /**
     * Sets do_not_validate_subject_recipients
     *
     * @param bool|null $do_not_validate_subject_recipients do_not_validate_subject_recipients
     *
     * @return self
     */
    public function setDoNotValidateSubjectRecipients($do_not_validate_subject_recipients)
    {
        $this->container['do_not_validate_subject_recipients'] = $do_not_validate_subject_recipients;

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
     * Gets force_authn
     *
     * @return bool|null
     */
    public function getForceAuthn()
    {
        return $this->container['force_authn'];
    }

    /**
     * Sets force_authn
     *
     * @param bool|null $force_authn force_authn
     *
     * @return self
     */
    public function setForceAuthn($force_authn)
    {
        $this->container['force_authn'] = $force_authn;

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


