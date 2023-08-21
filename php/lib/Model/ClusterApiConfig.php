<?php
/**
 * ClusterApiConfig
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
 * ClusterApiConfig Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ClusterApiConfig implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'clusterApiConfig';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'service_url' => 'string',
        'service_host' => 'string',
        'service_path' => 'string',
        'service_hostname' => 'string',
        'service_port' => 'int',
        'service_username' => 'string',
        'service_password' => 'string',
        'service_password_hash' => 'string',
        'service_token' => 'string',
        'service_access' => 'string',
        'service_cert' => 'string',
        'service_version' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'service_url' => null,
        'service_host' => null,
        'service_path' => null,
        'service_hostname' => null,
        'service_port' => 'int64',
        'service_username' => null,
        'service_password' => null,
        'service_password_hash' => null,
        'service_token' => null,
        'service_access' => null,
        'service_cert' => null,
        'service_version' => null
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
        'service_url' => 'serviceUrl',
        'service_host' => 'serviceHost',
        'service_path' => 'servicePath',
        'service_hostname' => 'serviceHostname',
        'service_port' => 'servicePort',
        'service_username' => 'serviceUsername',
        'service_password' => 'servicePassword',
        'service_password_hash' => 'servicePasswordHash',
        'service_token' => 'serviceToken',
        'service_access' => 'serviceAccess',
        'service_cert' => 'serviceCert',
        'service_version' => 'serviceVersion'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'service_url' => 'setServiceUrl',
        'service_host' => 'setServiceHost',
        'service_path' => 'setServicePath',
        'service_hostname' => 'setServiceHostname',
        'service_port' => 'setServicePort',
        'service_username' => 'setServiceUsername',
        'service_password' => 'setServicePassword',
        'service_password_hash' => 'setServicePasswordHash',
        'service_token' => 'setServiceToken',
        'service_access' => 'setServiceAccess',
        'service_cert' => 'setServiceCert',
        'service_version' => 'setServiceVersion'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'service_url' => 'getServiceUrl',
        'service_host' => 'getServiceHost',
        'service_path' => 'getServicePath',
        'service_hostname' => 'getServiceHostname',
        'service_port' => 'getServicePort',
        'service_username' => 'getServiceUsername',
        'service_password' => 'getServicePassword',
        'service_password_hash' => 'getServicePasswordHash',
        'service_token' => 'getServiceToken',
        'service_access' => 'getServiceAccess',
        'service_cert' => 'getServiceCert',
        'service_version' => 'getServiceVersion'
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
        $this->container['service_url'] = $data['service_url'] ?? null;
        $this->container['service_host'] = $data['service_host'] ?? null;
        $this->container['service_path'] = $data['service_path'] ?? null;
        $this->container['service_hostname'] = $data['service_hostname'] ?? null;
        $this->container['service_port'] = $data['service_port'] ?? null;
        $this->container['service_username'] = $data['service_username'] ?? null;
        $this->container['service_password'] = $data['service_password'] ?? null;
        $this->container['service_password_hash'] = $data['service_password_hash'] ?? null;
        $this->container['service_token'] = $data['service_token'] ?? null;
        $this->container['service_access'] = $data['service_access'] ?? null;
        $this->container['service_cert'] = $data['service_cert'] ?? null;
        $this->container['service_version'] = $data['service_version'] ?? null;
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
     * @param string|null $service_url service_url
     *
     * @return self
     */
    public function setServiceUrl($service_url)
    {
        $this->container['service_url'] = $service_url;

        return $this;
    }

    /**
     * Gets service_host
     *
     * @return string|null
     */
    public function getServiceHost()
    {
        return $this->container['service_host'];
    }

    /**
     * Sets service_host
     *
     * @param string|null $service_host service_host
     *
     * @return self
     */
    public function setServiceHost($service_host)
    {
        $this->container['service_host'] = $service_host;

        return $this;
    }

    /**
     * Gets service_path
     *
     * @return string|null
     */
    public function getServicePath()
    {
        return $this->container['service_path'];
    }

    /**
     * Sets service_path
     *
     * @param string|null $service_path service_path
     *
     * @return self
     */
    public function setServicePath($service_path)
    {
        $this->container['service_path'] = $service_path;

        return $this;
    }

    /**
     * Gets service_hostname
     *
     * @return string|null
     */
    public function getServiceHostname()
    {
        return $this->container['service_hostname'];
    }

    /**
     * Sets service_hostname
     *
     * @param string|null $service_hostname service_hostname
     *
     * @return self
     */
    public function setServiceHostname($service_hostname)
    {
        $this->container['service_hostname'] = $service_hostname;

        return $this;
    }

    /**
     * Gets service_port
     *
     * @return int|null
     */
    public function getServicePort()
    {
        return $this->container['service_port'];
    }

    /**
     * Sets service_port
     *
     * @param int|null $service_port service_port
     *
     * @return self
     */
    public function setServicePort($service_port)
    {
        $this->container['service_port'] = $service_port;

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
     * @param string|null $service_username service_username
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
     * @param string|null $service_password service_password
     *
     * @return self
     */
    public function setServicePassword($service_password)
    {
        $this->container['service_password'] = $service_password;

        return $this;
    }

    /**
     * Gets service_password_hash
     *
     * @return string|null
     */
    public function getServicePasswordHash()
    {
        return $this->container['service_password_hash'];
    }

    /**
     * Sets service_password_hash
     *
     * @param string|null $service_password_hash service_password_hash
     *
     * @return self
     */
    public function setServicePasswordHash($service_password_hash)
    {
        $this->container['service_password_hash'] = $service_password_hash;

        return $this;
    }

    /**
     * Gets service_token
     *
     * @return string|null
     */
    public function getServiceToken()
    {
        return $this->container['service_token'];
    }

    /**
     * Sets service_token
     *
     * @param string|null $service_token API Token
     *
     * @return self
     */
    public function setServiceToken($service_token)
    {
        $this->container['service_token'] = $service_token;

        return $this;
    }

    /**
     * Gets service_access
     *
     * @return string|null
     */
    public function getServiceAccess()
    {
        return $this->container['service_access'];
    }

    /**
     * Sets service_access
     *
     * @param string|null $service_access Kube Config
     *
     * @return self
     */
    public function setServiceAccess($service_access)
    {
        $this->container['service_access'] = $service_access;

        return $this;
    }

    /**
     * Gets service_cert
     *
     * @return string|null
     */
    public function getServiceCert()
    {
        return $this->container['service_cert'];
    }

    /**
     * Sets service_cert
     *
     * @param string|null $service_cert service_cert
     *
     * @return self
     */
    public function setServiceCert($service_cert)
    {
        $this->container['service_cert'] = $service_cert;

        return $this;
    }

    /**
     * Gets service_version
     *
     * @return string|null
     */
    public function getServiceVersion()
    {
        return $this->container['service_version'];
    }

    /**
     * Sets service_version
     *
     * @param string|null $service_version service_version
     *
     * @return self
     */
    public function setServiceVersion($service_version)
    {
        $this->container['service_version'] = $service_version;

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

