<?php
/**
 * VirtualImageLocation
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
 * VirtualImageLocation Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class VirtualImageLocation implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'virtualImageLocation';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'cloud' => '\OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType',
        'code' => 'string',
        'internal_id' => 'string',
        'external_id' => 'string',
        'external_disk_id' => 'string',
        'remote_path' => 'string',
        'image_path' => 'string',
        'image_name' => 'string',
        'image_region' => 'string',
        'image_folder' => 'string',
        'ref_type' => 'string',
        'ref_id' => 'int',
        'node_ref_type' => 'string',
        'node_ref_id' => 'string',
        'sub_ref_type' => 'string',
        'sub_ref_id' => 'string',
        'is_public' => 'bool',
        'system_image' => 'bool',
        'disk_index' => 'int',
        'price_plan' => 'string',
        'volumes' => 'object[]',
        'storage_controllers' => 'object[]',
        'network_interfaces' => 'object[]',
        'virtual_image' => '\OpenAPI\Client\Model\VirtualImageLocationVirtualImage'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'id' => 'int64',
        'cloud' => null,
        'code' => null,
        'internal_id' => null,
        'external_id' => null,
        'external_disk_id' => null,
        'remote_path' => null,
        'image_path' => null,
        'image_name' => null,
        'image_region' => null,
        'image_folder' => null,
        'ref_type' => null,
        'ref_id' => 'int64',
        'node_ref_type' => null,
        'node_ref_id' => null,
        'sub_ref_type' => null,
        'sub_ref_id' => null,
        'is_public' => null,
        'system_image' => null,
        'disk_index' => 'int64',
        'price_plan' => null,
        'volumes' => null,
        'storage_controllers' => null,
        'network_interfaces' => null,
        'virtual_image' => null
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
        'id' => 'id',
        'cloud' => 'cloud',
        'code' => 'code',
        'internal_id' => 'internalId',
        'external_id' => 'externalId',
        'external_disk_id' => 'externalDiskId',
        'remote_path' => 'remotePath',
        'image_path' => 'imagePath',
        'image_name' => 'imageName',
        'image_region' => 'imageRegion',
        'image_folder' => 'imageFolder',
        'ref_type' => 'refType',
        'ref_id' => 'refId',
        'node_ref_type' => 'nodeRefType',
        'node_ref_id' => 'nodeRefId',
        'sub_ref_type' => 'subRefType',
        'sub_ref_id' => 'subRefId',
        'is_public' => 'isPublic',
        'system_image' => 'systemImage',
        'disk_index' => 'diskIndex',
        'price_plan' => 'pricePlan',
        'volumes' => 'volumes',
        'storage_controllers' => 'storageControllers',
        'network_interfaces' => 'networkInterfaces',
        'virtual_image' => 'virtualImage'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'cloud' => 'setCloud',
        'code' => 'setCode',
        'internal_id' => 'setInternalId',
        'external_id' => 'setExternalId',
        'external_disk_id' => 'setExternalDiskId',
        'remote_path' => 'setRemotePath',
        'image_path' => 'setImagePath',
        'image_name' => 'setImageName',
        'image_region' => 'setImageRegion',
        'image_folder' => 'setImageFolder',
        'ref_type' => 'setRefType',
        'ref_id' => 'setRefId',
        'node_ref_type' => 'setNodeRefType',
        'node_ref_id' => 'setNodeRefId',
        'sub_ref_type' => 'setSubRefType',
        'sub_ref_id' => 'setSubRefId',
        'is_public' => 'setIsPublic',
        'system_image' => 'setSystemImage',
        'disk_index' => 'setDiskIndex',
        'price_plan' => 'setPricePlan',
        'volumes' => 'setVolumes',
        'storage_controllers' => 'setStorageControllers',
        'network_interfaces' => 'setNetworkInterfaces',
        'virtual_image' => 'setVirtualImage'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'cloud' => 'getCloud',
        'code' => 'getCode',
        'internal_id' => 'getInternalId',
        'external_id' => 'getExternalId',
        'external_disk_id' => 'getExternalDiskId',
        'remote_path' => 'getRemotePath',
        'image_path' => 'getImagePath',
        'image_name' => 'getImageName',
        'image_region' => 'getImageRegion',
        'image_folder' => 'getImageFolder',
        'ref_type' => 'getRefType',
        'ref_id' => 'getRefId',
        'node_ref_type' => 'getNodeRefType',
        'node_ref_id' => 'getNodeRefId',
        'sub_ref_type' => 'getSubRefType',
        'sub_ref_id' => 'getSubRefId',
        'is_public' => 'getIsPublic',
        'system_image' => 'getSystemImage',
        'disk_index' => 'getDiskIndex',
        'price_plan' => 'getPricePlan',
        'volumes' => 'getVolumes',
        'storage_controllers' => 'getStorageControllers',
        'network_interfaces' => 'getNetworkInterfaces',
        'virtual_image' => 'getVirtualImage'
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
        $this->container['id'] = $data['id'] ?? null;
        $this->container['cloud'] = $data['cloud'] ?? null;
        $this->container['code'] = $data['code'] ?? null;
        $this->container['internal_id'] = $data['internal_id'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['external_disk_id'] = $data['external_disk_id'] ?? null;
        $this->container['remote_path'] = $data['remote_path'] ?? null;
        $this->container['image_path'] = $data['image_path'] ?? null;
        $this->container['image_name'] = $data['image_name'] ?? null;
        $this->container['image_region'] = $data['image_region'] ?? null;
        $this->container['image_folder'] = $data['image_folder'] ?? null;
        $this->container['ref_type'] = $data['ref_type'] ?? null;
        $this->container['ref_id'] = $data['ref_id'] ?? null;
        $this->container['node_ref_type'] = $data['node_ref_type'] ?? null;
        $this->container['node_ref_id'] = $data['node_ref_id'] ?? null;
        $this->container['sub_ref_type'] = $data['sub_ref_type'] ?? null;
        $this->container['sub_ref_id'] = $data['sub_ref_id'] ?? null;
        $this->container['is_public'] = $data['is_public'] ?? null;
        $this->container['system_image'] = $data['system_image'] ?? null;
        $this->container['disk_index'] = $data['disk_index'] ?? null;
        $this->container['price_plan'] = $data['price_plan'] ?? null;
        $this->container['volumes'] = $data['volumes'] ?? null;
        $this->container['storage_controllers'] = $data['storage_controllers'] ?? null;
        $this->container['network_interfaces'] = $data['network_interfaces'] ?? null;
        $this->container['virtual_image'] = $data['virtual_image'] ?? null;
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
     * Gets id
     *
     * @return int|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param int|null $id id
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets cloud
     *
     * @return \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType|null
     */
    public function getCloud()
    {
        return $this->container['cloud'];
    }

    /**
     * Sets cloud
     *
     * @param \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType|null $cloud cloud
     *
     * @return self
     */
    public function setCloud($cloud)
    {
        $this->container['cloud'] = $cloud;

        return $this;
    }

    /**
     * Gets code
     *
     * @return string|null
     */
    public function getCode()
    {
        return $this->container['code'];
    }

    /**
     * Sets code
     *
     * @param string|null $code code
     *
     * @return self
     */
    public function setCode($code)
    {
        $this->container['code'] = $code;

        return $this;
    }

    /**
     * Gets internal_id
     *
     * @return string|null
     */
    public function getInternalId()
    {
        return $this->container['internal_id'];
    }

    /**
     * Sets internal_id
     *
     * @param string|null $internal_id internal_id
     *
     * @return self
     */
    public function setInternalId($internal_id)
    {
        $this->container['internal_id'] = $internal_id;

        return $this;
    }

    /**
     * Gets external_id
     *
     * @return string|null
     */
    public function getExternalId()
    {
        return $this->container['external_id'];
    }

    /**
     * Sets external_id
     *
     * @param string|null $external_id external_id
     *
     * @return self
     */
    public function setExternalId($external_id)
    {
        $this->container['external_id'] = $external_id;

        return $this;
    }

    /**
     * Gets external_disk_id
     *
     * @return string|null
     */
    public function getExternalDiskId()
    {
        return $this->container['external_disk_id'];
    }

    /**
     * Sets external_disk_id
     *
     * @param string|null $external_disk_id external_disk_id
     *
     * @return self
     */
    public function setExternalDiskId($external_disk_id)
    {
        $this->container['external_disk_id'] = $external_disk_id;

        return $this;
    }

    /**
     * Gets remote_path
     *
     * @return string|null
     */
    public function getRemotePath()
    {
        return $this->container['remote_path'];
    }

    /**
     * Sets remote_path
     *
     * @param string|null $remote_path remote_path
     *
     * @return self
     */
    public function setRemotePath($remote_path)
    {
        $this->container['remote_path'] = $remote_path;

        return $this;
    }

    /**
     * Gets image_path
     *
     * @return string|null
     */
    public function getImagePath()
    {
        return $this->container['image_path'];
    }

    /**
     * Sets image_path
     *
     * @param string|null $image_path image_path
     *
     * @return self
     */
    public function setImagePath($image_path)
    {
        $this->container['image_path'] = $image_path;

        return $this;
    }

    /**
     * Gets image_name
     *
     * @return string|null
     */
    public function getImageName()
    {
        return $this->container['image_name'];
    }

    /**
     * Sets image_name
     *
     * @param string|null $image_name image_name
     *
     * @return self
     */
    public function setImageName($image_name)
    {
        $this->container['image_name'] = $image_name;

        return $this;
    }

    /**
     * Gets image_region
     *
     * @return string|null
     */
    public function getImageRegion()
    {
        return $this->container['image_region'];
    }

    /**
     * Sets image_region
     *
     * @param string|null $image_region image_region
     *
     * @return self
     */
    public function setImageRegion($image_region)
    {
        $this->container['image_region'] = $image_region;

        return $this;
    }

    /**
     * Gets image_folder
     *
     * @return string|null
     */
    public function getImageFolder()
    {
        return $this->container['image_folder'];
    }

    /**
     * Sets image_folder
     *
     * @param string|null $image_folder image_folder
     *
     * @return self
     */
    public function setImageFolder($image_folder)
    {
        $this->container['image_folder'] = $image_folder;

        return $this;
    }

    /**
     * Gets ref_type
     *
     * @return string|null
     */
    public function getRefType()
    {
        return $this->container['ref_type'];
    }

    /**
     * Sets ref_type
     *
     * @param string|null $ref_type ref_type
     *
     * @return self
     */
    public function setRefType($ref_type)
    {
        $this->container['ref_type'] = $ref_type;

        return $this;
    }

    /**
     * Gets ref_id
     *
     * @return int|null
     */
    public function getRefId()
    {
        return $this->container['ref_id'];
    }

    /**
     * Sets ref_id
     *
     * @param int|null $ref_id ref_id
     *
     * @return self
     */
    public function setRefId($ref_id)
    {
        $this->container['ref_id'] = $ref_id;

        return $this;
    }

    /**
     * Gets node_ref_type
     *
     * @return string|null
     */
    public function getNodeRefType()
    {
        return $this->container['node_ref_type'];
    }

    /**
     * Sets node_ref_type
     *
     * @param string|null $node_ref_type node_ref_type
     *
     * @return self
     */
    public function setNodeRefType($node_ref_type)
    {
        $this->container['node_ref_type'] = $node_ref_type;

        return $this;
    }

    /**
     * Gets node_ref_id
     *
     * @return string|null
     */
    public function getNodeRefId()
    {
        return $this->container['node_ref_id'];
    }

    /**
     * Sets node_ref_id
     *
     * @param string|null $node_ref_id node_ref_id
     *
     * @return self
     */
    public function setNodeRefId($node_ref_id)
    {
        $this->container['node_ref_id'] = $node_ref_id;

        return $this;
    }

    /**
     * Gets sub_ref_type
     *
     * @return string|null
     */
    public function getSubRefType()
    {
        return $this->container['sub_ref_type'];
    }

    /**
     * Sets sub_ref_type
     *
     * @param string|null $sub_ref_type sub_ref_type
     *
     * @return self
     */
    public function setSubRefType($sub_ref_type)
    {
        $this->container['sub_ref_type'] = $sub_ref_type;

        return $this;
    }

    /**
     * Gets sub_ref_id
     *
     * @return string|null
     */
    public function getSubRefId()
    {
        return $this->container['sub_ref_id'];
    }

    /**
     * Sets sub_ref_id
     *
     * @param string|null $sub_ref_id sub_ref_id
     *
     * @return self
     */
    public function setSubRefId($sub_ref_id)
    {
        $this->container['sub_ref_id'] = $sub_ref_id;

        return $this;
    }

    /**
     * Gets is_public
     *
     * @return bool|null
     */
    public function getIsPublic()
    {
        return $this->container['is_public'];
    }

    /**
     * Sets is_public
     *
     * @param bool|null $is_public is_public
     *
     * @return self
     */
    public function setIsPublic($is_public)
    {
        $this->container['is_public'] = $is_public;

        return $this;
    }

    /**
     * Gets system_image
     *
     * @return bool|null
     */
    public function getSystemImage()
    {
        return $this->container['system_image'];
    }

    /**
     * Sets system_image
     *
     * @param bool|null $system_image system_image
     *
     * @return self
     */
    public function setSystemImage($system_image)
    {
        $this->container['system_image'] = $system_image;

        return $this;
    }

    /**
     * Gets disk_index
     *
     * @return int|null
     */
    public function getDiskIndex()
    {
        return $this->container['disk_index'];
    }

    /**
     * Sets disk_index
     *
     * @param int|null $disk_index disk_index
     *
     * @return self
     */
    public function setDiskIndex($disk_index)
    {
        $this->container['disk_index'] = $disk_index;

        return $this;
    }

    /**
     * Gets price_plan
     *
     * @return string|null
     */
    public function getPricePlan()
    {
        return $this->container['price_plan'];
    }

    /**
     * Sets price_plan
     *
     * @param string|null $price_plan price_plan
     *
     * @return self
     */
    public function setPricePlan($price_plan)
    {
        $this->container['price_plan'] = $price_plan;

        return $this;
    }

    /**
     * Gets volumes
     *
     * @return object[]|null
     */
    public function getVolumes()
    {
        return $this->container['volumes'];
    }

    /**
     * Sets volumes
     *
     * @param object[]|null $volumes volumes
     *
     * @return self
     */
    public function setVolumes($volumes)
    {
        $this->container['volumes'] = $volumes;

        return $this;
    }

    /**
     * Gets storage_controllers
     *
     * @return object[]|null
     */
    public function getStorageControllers()
    {
        return $this->container['storage_controllers'];
    }

    /**
     * Sets storage_controllers
     *
     * @param object[]|null $storage_controllers storage_controllers
     *
     * @return self
     */
    public function setStorageControllers($storage_controllers)
    {
        $this->container['storage_controllers'] = $storage_controllers;

        return $this;
    }

    /**
     * Gets network_interfaces
     *
     * @return object[]|null
     */
    public function getNetworkInterfaces()
    {
        return $this->container['network_interfaces'];
    }

    /**
     * Sets network_interfaces
     *
     * @param object[]|null $network_interfaces network_interfaces
     *
     * @return self
     */
    public function setNetworkInterfaces($network_interfaces)
    {
        $this->container['network_interfaces'] = $network_interfaces;

        return $this;
    }

    /**
     * Gets virtual_image
     *
     * @return \OpenAPI\Client\Model\VirtualImageLocationVirtualImage|null
     */
    public function getVirtualImage()
    {
        return $this->container['virtual_image'];
    }

    /**
     * Sets virtual_image
     *
     * @param \OpenAPI\Client\Model\VirtualImageLocationVirtualImage|null $virtual_image virtual_image
     *
     * @return self
     */
    public function setVirtualImage($virtual_image)
    {
        $this->container['virtual_image'] = $virtual_image;

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


