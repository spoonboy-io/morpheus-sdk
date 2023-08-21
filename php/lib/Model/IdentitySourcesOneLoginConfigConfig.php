<?php
/**
 * IdentitySourcesOneLoginConfigConfig
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
 * IdentitySourcesOneLoginConfigConfig Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class IdentitySourcesOneLoginConfigConfig implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'identitySourcesOneLoginConfig_config';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'subdomain' => 'string',
        'region' => 'string',
        'client_secret' => 'string',
        'client_id' => 'string',
        'required_role' => 'string',
        'required_role_id' => 'string',
        'client_secret_hash' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'subdomain' => null,
        'region' => null,
        'client_secret' => null,
        'client_id' => null,
        'required_role' => null,
        'required_role_id' => null,
        'client_secret_hash' => null
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
        'subdomain' => 'subdomain',
        'region' => 'region',
        'client_secret' => 'clientSecret',
        'client_id' => 'clientId',
        'required_role' => 'requiredRole',
        'required_role_id' => 'requiredRoleId',
        'client_secret_hash' => 'clientSecretHash'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'subdomain' => 'setSubdomain',
        'region' => 'setRegion',
        'client_secret' => 'setClientSecret',
        'client_id' => 'setClientId',
        'required_role' => 'setRequiredRole',
        'required_role_id' => 'setRequiredRoleId',
        'client_secret_hash' => 'setClientSecretHash'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'subdomain' => 'getSubdomain',
        'region' => 'getRegion',
        'client_secret' => 'getClientSecret',
        'client_id' => 'getClientId',
        'required_role' => 'getRequiredRole',
        'required_role_id' => 'getRequiredRoleId',
        'client_secret_hash' => 'getClientSecretHash'
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
        $this->container['subdomain'] = $data['subdomain'] ?? null;
        $this->container['region'] = $data['region'] ?? null;
        $this->container['client_secret'] = $data['client_secret'] ?? null;
        $this->container['client_id'] = $data['client_id'] ?? null;
        $this->container['required_role'] = $data['required_role'] ?? null;
        $this->container['required_role_id'] = $data['required_role_id'] ?? null;
        $this->container['client_secret_hash'] = $data['client_secret_hash'] ?? null;
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
     * Gets subdomain
     *
     * @return string|null
     */
    public function getSubdomain()
    {
        return $this->container['subdomain'];
    }

    /**
     * Sets subdomain
     *
     * @param string|null $subdomain subdomain
     *
     * @return self
     */
    public function setSubdomain($subdomain)
    {
        $this->container['subdomain'] = $subdomain;

        return $this;
    }

    /**
     * Gets region
     *
     * @return string|null
     */
    public function getRegion()
    {
        return $this->container['region'];
    }

    /**
     * Sets region
     *
     * @param string|null $region region
     *
     * @return self
     */
    public function setRegion($region)
    {
        $this->container['region'] = $region;

        return $this;
    }

    /**
     * Gets client_secret
     *
     * @return string|null
     */
    public function getClientSecret()
    {
        return $this->container['client_secret'];
    }

    /**
     * Sets client_secret
     *
     * @param string|null $client_secret client_secret
     *
     * @return self
     */
    public function setClientSecret($client_secret)
    {
        $this->container['client_secret'] = $client_secret;

        return $this;
    }

    /**
     * Gets client_id
     *
     * @return string|null
     */
    public function getClientId()
    {
        return $this->container['client_id'];
    }

    /**
     * Sets client_id
     *
     * @param string|null $client_id client_id
     *
     * @return self
     */
    public function setClientId($client_id)
    {
        $this->container['client_id'] = $client_id;

        return $this;
    }

    /**
     * Gets required_role
     *
     * @return string|null
     */
    public function getRequiredRole()
    {
        return $this->container['required_role'];
    }

    /**
     * Sets required_role
     *
     * @param string|null $required_role required_role
     *
     * @return self
     */
    public function setRequiredRole($required_role)
    {
        $this->container['required_role'] = $required_role;

        return $this;
    }

    /**
     * Gets required_role_id
     *
     * @return string|null
     */
    public function getRequiredRoleId()
    {
        return $this->container['required_role_id'];
    }

    /**
     * Sets required_role_id
     *
     * @param string|null $required_role_id required_role_id
     *
     * @return self
     */
    public function setRequiredRoleId($required_role_id)
    {
        $this->container['required_role_id'] = $required_role_id;

        return $this;
    }

    /**
     * Gets client_secret_hash
     *
     * @return string|null
     */
    public function getClientSecretHash()
    {
        return $this->container['client_secret_hash'];
    }

    /**
     * Sets client_secret_hash
     *
     * @param string|null $client_secret_hash client_secret_hash
     *
     * @return self
     */
    public function setClientSecretHash($client_secret_hash)
    {
        $this->container['client_secret_hash'] = $client_secret_hash;

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


