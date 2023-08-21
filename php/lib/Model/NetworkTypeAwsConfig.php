<?php
/**
 * NetworkTypeAwsConfig
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
 * NetworkTypeAwsConfig Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class NetworkTypeAwsConfig implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'networkTypeAwsConfig';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'availability_zone' => 'string',
        'cidr' => 'string',
        'assign_public_ip' => 'bool',
        'zone_pool' => '\OpenAPI\Client\Model\NetworkTypeAwsConfigZonePool'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'availability_zone' => null,
        'cidr' => null,
        'assign_public_ip' => null,
        'zone_pool' => null
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
        'availability_zone' => 'availabilityZone',
        'cidr' => 'cidr',
        'assign_public_ip' => 'assignPublicIp',
        'zone_pool' => 'zonePool'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'availability_zone' => 'setAvailabilityZone',
        'cidr' => 'setCidr',
        'assign_public_ip' => 'setAssignPublicIp',
        'zone_pool' => 'setZonePool'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'availability_zone' => 'getAvailabilityZone',
        'cidr' => 'getCidr',
        'assign_public_ip' => 'getAssignPublicIp',
        'zone_pool' => 'getZonePool'
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
        $this->container['availability_zone'] = $data['availability_zone'] ?? null;
        $this->container['cidr'] = $data['cidr'] ?? null;
        $this->container['assign_public_ip'] = $data['assign_public_ip'] ?? null;
        $this->container['zone_pool'] = $data['zone_pool'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['availability_zone'] === null) {
            $invalidProperties[] = "'availability_zone' can't be null";
        }
        if ($this->container['cidr'] === null) {
            $invalidProperties[] = "'cidr' can't be null";
        }
        if ($this->container['assign_public_ip'] === null) {
            $invalidProperties[] = "'assign_public_ip' can't be null";
        }
        if ($this->container['zone_pool'] === null) {
            $invalidProperties[] = "'zone_pool' can't be null";
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
     * Gets availability_zone
     *
     * @return string
     */
    public function getAvailabilityZone()
    {
        return $this->container['availability_zone'];
    }

    /**
     * Sets availability_zone
     *
     * @param string $availability_zone Availability Zone Name
     *
     * @return self
     */
    public function setAvailabilityZone($availability_zone)
    {
        $this->container['availability_zone'] = $availability_zone;

        return $this;
    }

    /**
     * Gets cidr
     *
     * @return string
     */
    public function getCidr()
    {
        return $this->container['cidr'];
    }

    /**
     * Sets cidr
     *
     * @param string $cidr Network CIDR
     *
     * @return self
     */
    public function setCidr($cidr)
    {
        $this->container['cidr'] = $cidr;

        return $this;
    }

    /**
     * Gets assign_public_ip
     *
     * @return bool
     */
    public function getAssignPublicIp()
    {
        return $this->container['assign_public_ip'];
    }

    /**
     * Sets assign_public_ip
     *
     * @param bool $assign_public_ip Assign public IPs by default.
     *
     * @return self
     */
    public function setAssignPublicIp($assign_public_ip)
    {
        $this->container['assign_public_ip'] = $assign_public_ip;

        return $this;
    }

    /**
     * Gets zone_pool
     *
     * @return \OpenAPI\Client\Model\NetworkTypeAwsConfigZonePool
     */
    public function getZonePool()
    {
        return $this->container['zone_pool'];
    }

    /**
     * Sets zone_pool
     *
     * @param \OpenAPI\Client\Model\NetworkTypeAwsConfigZonePool $zone_pool zone_pool
     *
     * @return self
     */
    public function setZonePool($zone_pool)
    {
        $this->container['zone_pool'] = $zone_pool;

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


