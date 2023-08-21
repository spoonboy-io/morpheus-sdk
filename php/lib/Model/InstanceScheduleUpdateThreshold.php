<?php
/**
 * InstanceScheduleUpdateThreshold
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
 * InstanceScheduleUpdateThreshold Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InstanceScheduleUpdateThreshold implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'instanceScheduleUpdate_threshold';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'source_threshold_id' => 'int',
        'auto_up' => 'bool',
        'auto_down' => 'bool',
        'min_count' => 'int',
        'max_count' => 'int',
        'cpu_enabled' => 'bool',
        'min_cpu' => 'double',
        'max_cpu' => 'double',
        'memory_enabled' => 'bool',
        'min_memory' => 'double',
        'max_memory' => 'double',
        'disk_enabled' => 'bool',
        'min_disk' => 'double',
        'max_disk' => 'double'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'source_threshold_id' => 'int64',
        'auto_up' => null,
        'auto_down' => null,
        'min_count' => null,
        'max_count' => null,
        'cpu_enabled' => null,
        'min_cpu' => 'double',
        'max_cpu' => 'double',
        'memory_enabled' => null,
        'min_memory' => 'double',
        'max_memory' => 'double',
        'disk_enabled' => null,
        'min_disk' => 'double',
        'max_disk' => 'double'
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
        'source_threshold_id' => 'sourceThresholdId',
        'auto_up' => 'autoUp',
        'auto_down' => 'autoDown',
        'min_count' => 'minCount',
        'max_count' => 'maxCount',
        'cpu_enabled' => 'cpuEnabled',
        'min_cpu' => 'minCpu',
        'max_cpu' => 'maxCpu',
        'memory_enabled' => 'memoryEnabled',
        'min_memory' => 'minMemory',
        'max_memory' => 'maxMemory',
        'disk_enabled' => 'diskEnabled',
        'min_disk' => 'minDisk',
        'max_disk' => 'maxDisk'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'source_threshold_id' => 'setSourceThresholdId',
        'auto_up' => 'setAutoUp',
        'auto_down' => 'setAutoDown',
        'min_count' => 'setMinCount',
        'max_count' => 'setMaxCount',
        'cpu_enabled' => 'setCpuEnabled',
        'min_cpu' => 'setMinCpu',
        'max_cpu' => 'setMaxCpu',
        'memory_enabled' => 'setMemoryEnabled',
        'min_memory' => 'setMinMemory',
        'max_memory' => 'setMaxMemory',
        'disk_enabled' => 'setDiskEnabled',
        'min_disk' => 'setMinDisk',
        'max_disk' => 'setMaxDisk'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'source_threshold_id' => 'getSourceThresholdId',
        'auto_up' => 'getAutoUp',
        'auto_down' => 'getAutoDown',
        'min_count' => 'getMinCount',
        'max_count' => 'getMaxCount',
        'cpu_enabled' => 'getCpuEnabled',
        'min_cpu' => 'getMinCpu',
        'max_cpu' => 'getMaxCpu',
        'memory_enabled' => 'getMemoryEnabled',
        'min_memory' => 'getMinMemory',
        'max_memory' => 'getMaxMemory',
        'disk_enabled' => 'getDiskEnabled',
        'min_disk' => 'getMinDisk',
        'max_disk' => 'getMaxDisk'
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
        $this->container['source_threshold_id'] = $data['source_threshold_id'] ?? null;
        $this->container['auto_up'] = $data['auto_up'] ?? null;
        $this->container['auto_down'] = $data['auto_down'] ?? null;
        $this->container['min_count'] = $data['min_count'] ?? null;
        $this->container['max_count'] = $data['max_count'] ?? null;
        $this->container['cpu_enabled'] = $data['cpu_enabled'] ?? null;
        $this->container['min_cpu'] = $data['min_cpu'] ?? null;
        $this->container['max_cpu'] = $data['max_cpu'] ?? null;
        $this->container['memory_enabled'] = $data['memory_enabled'] ?? null;
        $this->container['min_memory'] = $data['min_memory'] ?? null;
        $this->container['max_memory'] = $data['max_memory'] ?? null;
        $this->container['disk_enabled'] = $data['disk_enabled'] ?? null;
        $this->container['min_disk'] = $data['min_disk'] ?? null;
        $this->container['max_disk'] = $data['max_disk'] ?? null;
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
     * Gets source_threshold_id
     *
     * @return int|null
     */
    public function getSourceThresholdId()
    {
        return $this->container['source_threshold_id'];
    }

    /**
     * Sets source_threshold_id
     *
     * @param int|null $source_threshold_id Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly.
     *
     * @return self
     */
    public function setSourceThresholdId($source_threshold_id)
    {
        $this->container['source_threshold_id'] = $source_threshold_id;

        return $this;
    }

    /**
     * Gets auto_up
     *
     * @return bool|null
     */
    public function getAutoUp()
    {
        return $this->container['auto_up'];
    }

    /**
     * Sets auto_up
     *
     * @param bool|null $auto_up Auto Upscale
     *
     * @return self
     */
    public function setAutoUp($auto_up)
    {
        $this->container['auto_up'] = $auto_up;

        return $this;
    }

    /**
     * Gets auto_down
     *
     * @return bool|null
     */
    public function getAutoDown()
    {
        return $this->container['auto_down'];
    }

    /**
     * Sets auto_down
     *
     * @param bool|null $auto_down Auto Downscale
     *
     * @return self
     */
    public function setAutoDown($auto_down)
    {
        $this->container['auto_down'] = $auto_down;

        return $this;
    }

    /**
     * Gets min_count
     *
     * @return int|null
     */
    public function getMinCount()
    {
        return $this->container['min_count'];
    }

    /**
     * Sets min_count
     *
     * @param int|null $min_count The minimum number of nodes to scale down to
     *
     * @return self
     */
    public function setMinCount($min_count)
    {
        $this->container['min_count'] = $min_count;

        return $this;
    }

    /**
     * Gets max_count
     *
     * @return int|null
     */
    public function getMaxCount()
    {
        return $this->container['max_count'];
    }

    /**
     * Sets max_count
     *
     * @param int|null $max_count The maximum number of nodes to scale up to
     *
     * @return self
     */
    public function setMaxCount($max_count)
    {
        $this->container['max_count'] = $max_count;

        return $this;
    }

    /**
     * Gets cpu_enabled
     *
     * @return bool|null
     */
    public function getCpuEnabled()
    {
        return $this->container['cpu_enabled'];
    }

    /**
     * Sets cpu_enabled
     *
     * @param bool|null $cpu_enabled Enable CPU Threshold
     *
     * @return self
     */
    public function setCpuEnabled($cpu_enabled)
    {
        $this->container['cpu_enabled'] = $cpu_enabled;

        return $this;
    }

    /**
     * Gets min_cpu
     *
     * @return double|null
     */
    public function getMinCpu()
    {
        return $this->container['min_cpu'];
    }

    /**
     * Sets min_cpu
     *
     * @param double|null $min_cpu Min CPU (%)
     *
     * @return self
     */
    public function setMinCpu($min_cpu)
    {
        $this->container['min_cpu'] = $min_cpu;

        return $this;
    }

    /**
     * Gets max_cpu
     *
     * @return double|null
     */
    public function getMaxCpu()
    {
        return $this->container['max_cpu'];
    }

    /**
     * Sets max_cpu
     *
     * @param double|null $max_cpu Max CPU (%)
     *
     * @return self
     */
    public function setMaxCpu($max_cpu)
    {
        $this->container['max_cpu'] = $max_cpu;

        return $this;
    }

    /**
     * Gets memory_enabled
     *
     * @return bool|null
     */
    public function getMemoryEnabled()
    {
        return $this->container['memory_enabled'];
    }

    /**
     * Sets memory_enabled
     *
     * @param bool|null $memory_enabled Enable Memory Threshold
     *
     * @return self
     */
    public function setMemoryEnabled($memory_enabled)
    {
        $this->container['memory_enabled'] = $memory_enabled;

        return $this;
    }

    /**
     * Gets min_memory
     *
     * @return double|null
     */
    public function getMinMemory()
    {
        return $this->container['min_memory'];
    }

    /**
     * Sets min_memory
     *
     * @param double|null $min_memory Min Memory (%)
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
     * @return double|null
     */
    public function getMaxMemory()
    {
        return $this->container['max_memory'];
    }

    /**
     * Sets max_memory
     *
     * @param double|null $max_memory Max Memory (%)
     *
     * @return self
     */
    public function setMaxMemory($max_memory)
    {
        $this->container['max_memory'] = $max_memory;

        return $this;
    }

    /**
     * Gets disk_enabled
     *
     * @return bool|null
     */
    public function getDiskEnabled()
    {
        return $this->container['disk_enabled'];
    }

    /**
     * Sets disk_enabled
     *
     * @param bool|null $disk_enabled Enable Disk Threshold
     *
     * @return self
     */
    public function setDiskEnabled($disk_enabled)
    {
        $this->container['disk_enabled'] = $disk_enabled;

        return $this;
    }

    /**
     * Gets min_disk
     *
     * @return double|null
     */
    public function getMinDisk()
    {
        return $this->container['min_disk'];
    }

    /**
     * Sets min_disk
     *
     * @param double|null $min_disk Min Disk (%)
     *
     * @return self
     */
    public function setMinDisk($min_disk)
    {
        $this->container['min_disk'] = $min_disk;

        return $this;
    }

    /**
     * Gets max_disk
     *
     * @return double|null
     */
    public function getMaxDisk()
    {
        return $this->container['max_disk'];
    }

    /**
     * Sets max_disk
     *
     * @param double|null $max_disk Max Disk (%)
     *
     * @return self
     */
    public function setMaxDisk($max_disk)
    {
        $this->container['max_disk'] = $max_disk;

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

