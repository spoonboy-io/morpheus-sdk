<?php
/**
 * UserSourceCreateSaml
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
 * UserSourceCreateSaml Class Doc Comment
 *
 * @category Class
 * @description SAML Configuration
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class UserSourceCreateSaml implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'userSourceCreateSaml';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'url' => 'string',
        'do_not_include_saml_request' => 'bool',
        'logout_url' => 'string',
        'saml_signature_mode' => 'string',
        'request509_certificate' => 'string',
        'request_private_key' => 'string',
        'do_not_validate_signature' => 'bool',
        'public_key' => 'string',
        'private_key' => 'string',
        'given_name_attribute' => 'string',
        'surname_attribute' => 'string',
        'role_attribute_name' => 'string',
        'required_attribute_value' => 'string'
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
        'do_not_include_saml_request' => null,
        'logout_url' => null,
        'saml_signature_mode' => null,
        'request509_certificate' => null,
        'request_private_key' => null,
        'do_not_validate_signature' => null,
        'public_key' => null,
        'private_key' => null,
        'given_name_attribute' => null,
        'surname_attribute' => null,
        'role_attribute_name' => null,
        'required_attribute_value' => null
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
        'do_not_include_saml_request' => 'doNotIncludeSAMLRequest',
        'logout_url' => 'logoutUrl',
        'saml_signature_mode' => 'SAMLSignatureMode',
        'request509_certificate' => 'request509Certificate',
        'request_private_key' => 'requestPrivateKey',
        'do_not_validate_signature' => 'doNotValidateSignature',
        'public_key' => 'publicKey',
        'private_key' => 'privateKey',
        'given_name_attribute' => 'givenNameAttribute',
        'surname_attribute' => 'surnameAttribute',
        'role_attribute_name' => 'roleAttributeName',
        'required_attribute_value' => 'requiredAttributeValue'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'url' => 'setUrl',
        'do_not_include_saml_request' => 'setDoNotIncludeSamlRequest',
        'logout_url' => 'setLogoutUrl',
        'saml_signature_mode' => 'setSamlSignatureMode',
        'request509_certificate' => 'setRequest509Certificate',
        'request_private_key' => 'setRequestPrivateKey',
        'do_not_validate_signature' => 'setDoNotValidateSignature',
        'public_key' => 'setPublicKey',
        'private_key' => 'setPrivateKey',
        'given_name_attribute' => 'setGivenNameAttribute',
        'surname_attribute' => 'setSurnameAttribute',
        'role_attribute_name' => 'setRoleAttributeName',
        'required_attribute_value' => 'setRequiredAttributeValue'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'url' => 'getUrl',
        'do_not_include_saml_request' => 'getDoNotIncludeSamlRequest',
        'logout_url' => 'getLogoutUrl',
        'saml_signature_mode' => 'getSamlSignatureMode',
        'request509_certificate' => 'getRequest509Certificate',
        'request_private_key' => 'getRequestPrivateKey',
        'do_not_validate_signature' => 'getDoNotValidateSignature',
        'public_key' => 'getPublicKey',
        'private_key' => 'getPrivateKey',
        'given_name_attribute' => 'getGivenNameAttribute',
        'surname_attribute' => 'getSurnameAttribute',
        'role_attribute_name' => 'getRoleAttributeName',
        'required_attribute_value' => 'getRequiredAttributeValue'
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

    const SAML_SIGNATURE_MODE_NO_SIGNATURE = 'NoSignature';
    const SAML_SIGNATURE_MODE_SELF_SIGNED = 'SelfSigned';
    const SAML_SIGNATURE_MODE_CUSTOM_SIGNATURE = 'CustomSignature';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getSamlSignatureModeAllowableValues()
    {
        return [
            self::SAML_SIGNATURE_MODE_NO_SIGNATURE,
            self::SAML_SIGNATURE_MODE_SELF_SIGNED,
            self::SAML_SIGNATURE_MODE_CUSTOM_SIGNATURE,
        ];
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
        $this->container['do_not_include_saml_request'] = $data['do_not_include_saml_request'] ?? false;
        $this->container['logout_url'] = $data['logout_url'] ?? null;
        $this->container['saml_signature_mode'] = $data['saml_signature_mode'] ?? 'NoSignature';
        $this->container['request509_certificate'] = $data['request509_certificate'] ?? null;
        $this->container['request_private_key'] = $data['request_private_key'] ?? null;
        $this->container['do_not_validate_signature'] = $data['do_not_validate_signature'] ?? true;
        $this->container['public_key'] = $data['public_key'] ?? null;
        $this->container['private_key'] = $data['private_key'] ?? null;
        $this->container['given_name_attribute'] = $data['given_name_attribute'] ?? null;
        $this->container['surname_attribute'] = $data['surname_attribute'] ?? null;
        $this->container['role_attribute_name'] = $data['role_attribute_name'] ?? null;
        $this->container['required_attribute_value'] = $data['required_attribute_value'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        $allowedValues = $this->getSamlSignatureModeAllowableValues();
        if (!is_null($this->container['saml_signature_mode']) && !in_array($this->container['saml_signature_mode'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'saml_signature_mode', must be one of '%s'",
                $this->container['saml_signature_mode'],
                implode("', '", $allowedValues)
            );
        }

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
     * @param string|null $url Login Redirect URL
     *
     * @return self
     */
    public function setUrl($url)
    {
        $this->container['url'] = $url;

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
     * @param bool|null $do_not_include_saml_request Exclude SAML Request Parameter
     *
     * @return self
     */
    public function setDoNotIncludeSamlRequest($do_not_include_saml_request)
    {
        $this->container['do_not_include_saml_request'] = $do_not_include_saml_request;

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
     * @param string|null $logout_url Lougout Post URL
     *
     * @return self
     */
    public function setLogoutUrl($logout_url)
    {
        $this->container['logout_url'] = $logout_url;

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
     * @param string|null $saml_signature_mode SAML Request signing mode
     *
     * @return self
     */
    public function setSamlSignatureMode($saml_signature_mode)
    {
        $allowedValues = $this->getSamlSignatureModeAllowableValues();
        if (!is_null($saml_signature_mode) && !in_array($saml_signature_mode, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'saml_signature_mode', must be one of '%s'",
                    $saml_signature_mode,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['saml_signature_mode'] = $saml_signature_mode;

        return $this;
    }

    /**
     * Gets request509_certificate
     *
     * @return string|null
     */
    public function getRequest509Certificate()
    {
        return $this->container['request509_certificate'];
    }

    /**
     * Sets request509_certificate
     *
     * @param string|null $request509_certificate Only applies when `SAMLSignatureMode=CustomSignature`
     *
     * @return self
     */
    public function setRequest509Certificate($request509_certificate)
    {
        $this->container['request509_certificate'] = $request509_certificate;

        return $this;
    }

    /**
     * Gets request_private_key
     *
     * @return string|null
     */
    public function getRequestPrivateKey()
    {
        return $this->container['request_private_key'];
    }

    /**
     * Sets request_private_key
     *
     * @param string|null $request_private_key RSA Private Key. Only applies when `SAMLSignatureMode=CustomSignature`
     *
     * @return self
     */
    public function setRequestPrivateKey($request_private_key)
    {
        $this->container['request_private_key'] = $request_private_key;

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
     * @param bool|null $do_not_validate_signature SAML Response Signing Flag
     *
     * @return self
     */
    public function setDoNotValidateSignature($do_not_validate_signature)
    {
        $this->container['do_not_validate_signature'] = $do_not_validate_signature;

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
     * @param string|null $public_key Signing Public Key. Only applies when `doNotValidateSignature=true`
     *
     * @return self
     */
    public function setPublicKey($public_key)
    {
        $this->container['public_key'] = $public_key;

        return $this;
    }

    /**
     * Gets private_key
     *
     * @return string|null
     */
    public function getPrivateKey()
    {
        return $this->container['private_key'];
    }

    /**
     * Sets private_key
     *
     * @param string|null $private_key Encryption RSA Private Key
     *
     * @return self
     */
    public function setPrivateKey($private_key)
    {
        $this->container['private_key'] = $private_key;

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
     * @param string|null $given_name_attribute Given Name Attribute Name
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
     * @param string|null $surname_attribute Surname Attribute Name
     *
     * @return self
     */
    public function setSurnameAttribute($surname_attribute)
    {
        $this->container['surname_attribute'] = $surname_attribute;

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
     * @param string|null $role_attribute_name Role Attribute Name
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
     * @param string|null $required_attribute_value Role Attibute Required Value
     *
     * @return self
     */
    public function setRequiredAttributeValue($required_attribute_value)
    {
        $this->container['required_attribute_value'] = $required_attribute_value;

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


