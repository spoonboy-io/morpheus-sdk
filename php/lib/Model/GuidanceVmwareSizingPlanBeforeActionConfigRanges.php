<?php
/**
 * GuidanceVmwareSizingPlanBeforeActionConfigRanges
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
 * GuidanceVmwareSizingPlanBeforeActionConfigRanges Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class GuidanceVmwareSizingPlanBeforeActionConfigRanges implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'guidanceVmwareSizing_planBeforeAction_config_ranges';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'min_storage' => 'string',
        'max_storage' => 'string',
        'min_memory' => 'string',
        'max_memory' => 'string',
        'min_cores' => 'string',
        'max_cores' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'min_storage' => null,
        'max_storage' => null,
        'min_memory' => null,
        'max_memory' => null,
        'min_cores' => null,
        'max_cores' => null
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
        'min_storage' => 'minStorage',
        'max_storage' => 'maxStorage',
        'min_memory' => 'minMemory',
        'max_memory' => 'maxMemory',
        'min_cores' => 'minCores',
        'max_cores' => 'maxCores'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'min_storage' => 'setMinStorage',
        'max_storage' => 'setMaxStorage',
        'min_memory' => 'setMinMemory',
        'max_memory' => 'setMaxMemory',
        'min_cores' => 'setMinCores',
        'max_cores' => 'setMaxCores'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'min_storage' => 'getMinStorage',
        'max_storage' => 'getMaxStorage',
        'min_memory' => 'getMinMemory',
        'max_memory' => 'getMaxMemory',
        'min_cores' => 'getMinCores',
        'max_cores' => 'getMaxCores'
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
        $this->container['min_storage'] = $data['min_storage'] ?? null;
        $this->container['max_storage'] = $data['max_storage'] ?? null;
        $this->container['min_memory'] = $data['min_memory'] ?? null;
        $this->container['max_memory'] = $data['max_memory'] ?? null;
        $this->container['min_cores'] = $data['min_cores'] ?? null;
        $this->container['max_cores'] = $data['max_cores'] ?? null;
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
     * Gets min_storage
     *
     * @return string|null
     */
    public function getMinStorage()
    {
        return $this->container['min_storage'];
    }

    /**
     * Sets min_storage
     *
     * @param string|null $min_storage min_storage
     *
     * @return self
     */
    public function setMinStorage($min_storage)
    {
        $this->container['min_storage'] = $min_storage;

        return $this;
    }

    /**
     * Gets max_storage
     *
     * @return string|null
     */
    public function getMaxStorage()
    {
        return $this->container['max_storage'];
    }

    /**
     * Sets max_storage
     *
     * @param string|null $max_storage max_storage
     *
     * @return self
     */
    public function setMaxStorage($max_storage)
    {
        $this->container['max_storage'] = $max_storage;

        return $this;
    }

    /**
     * Gets min_memory
     *
     * @return string|null
     */
    public function getMinMemory()
    {
        return $this->container['min_memory'];
    }

    /**
     * Sets min_memory
     *
     * @param string|null $min_memory min_memory
     *
     * @return self
     */
    public function setMinMemory($min_memory)
    {
        $this->container['min_memory'] = $min_memory;

        return $this;
    }

    /**
     * Gets max_memory
     *
     * @return string|null
     */
    public function getMaxMemory()
    {
        return $this->container['max_memory'];
    }

    /**
     * Sets max_memory
     *
     * @param string|null $max_memory max_memory
     *
     * @return self
     */
    public function setMaxMemory($max_memory)
    {
        $this->container['max_memory'] = $max_memory;

        return $this;
    }

    /**
     * Gets min_cores
     *
     * @return string|null
     */
    public function getMinCores()
    {
        return $this->container['min_cores'];
    }

    /**
     * Sets min_cores
     *
     * @param string|null $min_cores min_cores
     *
     * @return self
     */
    public function setMinCores($min_cores)
    {
        $this->container['min_cores'] = $min_cores;

        return $this;
    }

    /**
     * Gets max_cores
     *
     * @return string|null
     */
    public function getMaxCores()
    {
        return $this->container['max_cores'];
    }

    /**
     * Sets max_cores
     *
     * @param string|null $max_cores max_cores
     *
     * @return self
     */
    public function setMaxCores($max_cores)
    {
        $this->container['max_cores'] = $max_cores;

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


