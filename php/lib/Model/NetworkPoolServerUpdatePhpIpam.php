<?php
/**
 * NetworkPoolServerUpdatePhpIpam
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
 * NetworkPoolServerUpdatePhpIpam Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class NetworkPoolServerUpdatePhpIpam implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'networkPoolServerUpdatePhpIpam';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'name' => 'string',
        'enabled' => 'bool',
        'service_url' => 'string',
        'service_username' => 'string',
        'service_password' => 'string',
        'service_throttle_rate' => 'int',
        'ignore_ssl' => 'bool',
        'network_filter' => 'string',
        'config' => '\OpenAPI\Client\Model\NetworkPoolServerUpdatePhpIpamConfig',
        'credential' => '\OpenAPI\Client\Model\NetworkPoolServerCreateBluecatCredential'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'name' => null,
        'enabled' => null,
        'service_url' => null,
        'service_username' => null,
        'service_password' => null,
        'service_throttle_rate' => 'int64',
        'ignore_ssl' => null,
        'network_filter' => null,
        'config' => null,
        'credential' => null
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
        'name' => 'name',
        'enabled' => 'enabled',
        'service_url' => 'serviceUrl',
        'service_username' => 'serviceUsername',
        'service_password' => 'servicePassword',
        'service_throttle_rate' => 'serviceThrottleRate',
        'ignore_ssl' => 'ignoreSsl',
        'network_filter' => 'networkFilter',
        'config' => 'config',
        'credential' => 'credential'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'name' => 'setName',
        'enabled' => 'setEnabled',
        'service_url' => 'setServiceUrl',
        'service_username' => 'setServiceUsername',
        'service_password' => 'setServicePassword',
        'service_throttle_rate' => 'setServiceThrottleRate',
        'ignore_ssl' => 'setIgnoreSsl',
        'network_filter' => 'setNetworkFilter',
        'config' => 'setConfig',
        'credential' => 'setCredential'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'name' => 'getName',
        'enabled' => 'getEnabled',
        'service_url' => 'getServiceUrl',
        'service_username' => 'getServiceUsername',
        'service_password' => 'getServicePassword',
        'service_throttle_rate' => 'getServiceThrottleRate',
        'ignore_ssl' => 'getIgnoreSsl',
        'network_filter' => 'getNetworkFilter',
        'config' => 'getConfig',
        'credential' => 'getCredential'
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
        $this->container['name'] = $data['name'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? true;
        $this->container['service_url'] = $data['service_url'] ?? null;
        $this->container['service_username'] = $data['service_username'] ?? null;
        $this->container['service_password'] = $data['service_password'] ?? null;
        $this->container['service_throttle_rate'] = $data['service_throttle_rate'] ?? 0;
        $this->container['ignore_ssl'] = $data['ignore_ssl'] ?? null;
        $this->container['network_filter'] = $data['network_filter'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['credential'] = $data['credential'] ?? null;
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
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name Name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets enabled
     *
     * @return bool|null
     */
    public function getEnabled()
    {
        return $this->container['enabled'];
    }

    /**
     * Sets enabled
     *
     * @param bool|null $enabled Can be used to enable / disable the network pool server.
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets service_url
     *
     * @return string|null
     */
    public function getServiceUrl()
    {
        return $this->container['service_url'];
    }

    /**
     * Sets service_url
     *
     * @param string|null $service_url URL
     *
     * @return self
     */
    public function setServiceUrl($service_url)
    {
        $this->container['service_url'] = $service_url;

        return $this;
    }

    /**
     * Gets service_username
     *
     * @return string|null
     */
    public function getServiceUsername()
    {
        return $this->container['service_username'];
    }

    /**
     * Sets service_username
     *
     * @param string|null $service_username Username
     *
     * @return self
     */
    public function setServiceUsername($service_username)
    {
        $this->container['service_username'] = $service_username;

        return $this;
    }

    /**
     * Gets service_password
     *
     * @return string|null
     */
    public function getServicePassword()
    {
        return $this->container['service_password'];
    }

    /**
     * Sets service_password
     *
     * @param string|null $service_password Password
     *
     * @return self
     */
    public function setServicePassword($service_password)
    {
        $this->container['service_password'] = $service_password;

        return $this;
    }

    /**
     * Gets service_throttle_rate
     *
     * @return int|null
     */
    public function getServiceThrottleRate()
    {
        return $this->container['service_throttle_rate'];
    }

    /**
     * Sets service_throttle_rate
     *
     * @param int|null $service_throttle_rate Throttle Rate
     *
     * @return self
     */
    public function setServiceThrottleRate($service_throttle_rate)
    {
        $this->container['service_throttle_rate'] = $service_throttle_rate;

        return $this;
    }

    /**
     * Gets ignore_ssl
     *
     * @return bool|null
     */
    public function getIgnoreSsl()
    {
        return $this->container['ignore_ssl'];
    }

    /**
     * Sets ignore_ssl
     *
     * @param bool|null $ignore_ssl Disable SSL SNI Verification
     *
     * @return self
     */
    public function setIgnoreSsl($ignore_ssl)
    {
        $this->container['ignore_ssl'] = $ignore_ssl;

        return $this;
    }

    /**
     * Gets network_filter
     *
     * @return string|null
     */
    public function getNetworkFilter()
    {
        return $this->container['network_filter'];
    }

    /**
     * Sets network_filter
     *
     * @param string|null $network_filter Network Filter
     *
     * @return self
     */
    public function setNetworkFilter($network_filter)
    {
        $this->container['network_filter'] = $network_filter;

        return $this;
    }

    /**
     * Gets config
     *
     * @return \OpenAPI\Client\Model\NetworkPoolServerUpdatePhpIpamConfig|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param \OpenAPI\Client\Model\NetworkPoolServerUpdatePhpIpamConfig|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

        return $this;
    }

    /**
     * Gets credential
     *
     * @return \OpenAPI\Client\Model\NetworkPoolServerCreateBluecatCredential|null
     */
    public function getCredential()
    {
        return $this->container['credential'];
    }

    /**
     * Sets credential
     *
     * @param \OpenAPI\Client\Model\NetworkPoolServerCreateBluecatCredential|null $credential credential
     *
     * @return self
     */
    public function setCredential($credential)
    {
        $this->container['credential'] = $credential;

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


