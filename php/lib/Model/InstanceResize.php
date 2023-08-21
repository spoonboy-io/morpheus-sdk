<?php
/**
 * InstanceResize
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
 * InstanceResize Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InstanceResize implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'instanceResize';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'instance' => '\OpenAPI\Client\Model\InstanceResizeInstance',
        'service_plan_options' => '\OpenAPI\Client\Model\ServicePlanOptions',
        'volumes' => '\OpenAPI\Client\Model\InstanceCreateVolume[]',
        'delete_original_volumes' => 'bool',
        'network_interfaces' => '\OpenAPI\Client\Model\InstanceCreateNetwork[]'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'instance' => null,
        'service_plan_options' => null,
        'volumes' => null,
        'delete_original_volumes' => null,
        'network_interfaces' => null
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
        'instance' => 'instance',
        'service_plan_options' => 'servicePlanOptions',
        'volumes' => 'volumes',
        'delete_original_volumes' => 'deleteOriginalVolumes',
        'network_interfaces' => 'networkInterfaces'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'instance' => 'setInstance',
        'service_plan_options' => 'setServicePlanOptions',
        'volumes' => 'setVolumes',
        'delete_original_volumes' => 'setDeleteOriginalVolumes',
        'network_interfaces' => 'setNetworkInterfaces'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'instance' => 'getInstance',
        'service_plan_options' => 'getServicePlanOptions',
        'volumes' => 'getVolumes',
        'delete_original_volumes' => 'getDeleteOriginalVolumes',
        'network_interfaces' => 'getNetworkInterfaces'
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
        $this->container['instance'] = $data['instance'] ?? null;
        $this->container['service_plan_options'] = $data['service_plan_options'] ?? null;
        $this->container['volumes'] = $data['volumes'] ?? null;
        $this->container['delete_original_volumes'] = $data['delete_original_volumes'] ?? false;
        $this->container['network_interfaces'] = $data['network_interfaces'] ?? null;
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
     * Gets instance
     *
     * @return \OpenAPI\Client\Model\InstanceResizeInstance|null
     */
    public function getInstance()
    {
        return $this->container['instance'];
    }

    /**
     * Sets instance
     *
     * @param \OpenAPI\Client\Model\InstanceResizeInstance|null $instance instance
     *
     * @return self
     */
    public function setInstance($instance)
    {
        $this->container['instance'] = $instance;

        return $this;
    }

    /**
     * Gets service_plan_options
     *
     * @return \OpenAPI\Client\Model\ServicePlanOptions|null
     */
    public function getServicePlanOptions()
    {
        return $this->container['service_plan_options'];
    }

    /**
     * Sets service_plan_options
     *
     * @param \OpenAPI\Client\Model\ServicePlanOptions|null $service_plan_options service_plan_options
     *
     * @return self
     */
    public function setServicePlanOptions($service_plan_options)
    {
        $this->container['service_plan_options'] = $service_plan_options;

        return $this;
    }

    /**
     * Gets volumes
     *
     * @return \OpenAPI\Client\Model\InstanceCreateVolume[]|null
     */
    public function getVolumes()
    {
        return $this->container['volumes'];
    }

    /**
     * Sets volumes
     *
     * @param \OpenAPI\Client\Model\InstanceCreateVolume[]|null $volumes Can be used to grow just the logical volume of the instance instead of choosing a plan
     *
     * @return self
     */
    public function setVolumes($volumes)
    {
        $this->container['volumes'] = $volumes;

        return $this;
    }

    /**
     * Gets delete_original_volumes
     *
     * @return bool|null
     */
    public function getDeleteOriginalVolumes()
    {
        return $this->container['delete_original_volumes'];
    }

    /**
     * Sets delete_original_volumes
     *
     * @param bool|null $delete_original_volumes Delete the original volumes after resizing. (Amazon only)
     *
     * @return self
     */
    public function setDeleteOriginalVolumes($delete_original_volumes)
    {
        $this->container['delete_original_volumes'] = $delete_original_volumes;

        return $this;
    }

    /**
     * Gets network_interfaces
     *
     * @return \OpenAPI\Client\Model\InstanceCreateNetwork[]|null
     */
    public function getNetworkInterfaces()
    {
        return $this->container['network_interfaces'];
    }

    /**
     * Sets network_interfaces
     *
     * @param \OpenAPI\Client\Model\InstanceCreateNetwork[]|null $network_interfaces Key for network configuration. Include id to update an existing interface. The existing interfaces and their id can be retrieved with the hosts API.
     *
     * @return self
     */
    public function setNetworkInterfaces($network_interfaces)
    {
        $this->container['network_interfaces'] = $network_interfaces;

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


