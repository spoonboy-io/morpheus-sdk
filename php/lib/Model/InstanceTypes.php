<?php
/**
 * InstanceTypes
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
 * InstanceTypes Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InstanceTypes implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'instanceTypes';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'account' => '\OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert',
        'name' => 'string',
        'labels' => 'string[]',
        'code' => 'string',
        'description' => 'string',
        'provision_type_code' => 'string',
        'category' => 'string',
        'active' => 'bool',
        'environment_prefix' => 'string',
        'visibility' => 'string',
        'featured' => 'bool',
        'versions' => 'string[]',
        'instance_type_layouts' => '\OpenAPI\Client\Model\InstanceTypesInstanceTypeLayouts[]',
        'image_path' => 'string',
        'dark_image_path' => 'string'
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
        'account' => null,
        'name' => null,
        'labels' => null,
        'code' => null,
        'description' => null,
        'provision_type_code' => null,
        'category' => null,
        'active' => null,
        'environment_prefix' => null,
        'visibility' => null,
        'featured' => null,
        'versions' => null,
        'instance_type_layouts' => null,
        'image_path' => null,
        'dark_image_path' => null
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
        'account' => 'account',
        'name' => 'name',
        'labels' => 'labels',
        'code' => 'code',
        'description' => 'description',
        'provision_type_code' => 'provisionTypeCode',
        'category' => 'category',
        'active' => 'active',
        'environment_prefix' => 'environmentPrefix',
        'visibility' => 'visibility',
        'featured' => 'featured',
        'versions' => 'versions',
        'instance_type_layouts' => 'instanceTypeLayouts',
        'image_path' => 'imagePath',
        'dark_image_path' => 'darkImagePath'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'account' => 'setAccount',
        'name' => 'setName',
        'labels' => 'setLabels',
        'code' => 'setCode',
        'description' => 'setDescription',
        'provision_type_code' => 'setProvisionTypeCode',
        'category' => 'setCategory',
        'active' => 'setActive',
        'environment_prefix' => 'setEnvironmentPrefix',
        'visibility' => 'setVisibility',
        'featured' => 'setFeatured',
        'versions' => 'setVersions',
        'instance_type_layouts' => 'setInstanceTypeLayouts',
        'image_path' => 'setImagePath',
        'dark_image_path' => 'setDarkImagePath'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'account' => 'getAccount',
        'name' => 'getName',
        'labels' => 'getLabels',
        'code' => 'getCode',
        'description' => 'getDescription',
        'provision_type_code' => 'getProvisionTypeCode',
        'category' => 'getCategory',
        'active' => 'getActive',
        'environment_prefix' => 'getEnvironmentPrefix',
        'visibility' => 'getVisibility',
        'featured' => 'getFeatured',
        'versions' => 'getVersions',
        'instance_type_layouts' => 'getInstanceTypeLayouts',
        'image_path' => 'getImagePath',
        'dark_image_path' => 'getDarkImagePath'
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
        $this->container['account'] = $data['account'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['labels'] = $data['labels'] ?? null;
        $this->container['code'] = $data['code'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['provision_type_code'] = $data['provision_type_code'] ?? null;
        $this->container['category'] = $data['category'] ?? null;
        $this->container['active'] = $data['active'] ?? null;
        $this->container['environment_prefix'] = $data['environment_prefix'] ?? null;
        $this->container['visibility'] = $data['visibility'] ?? null;
        $this->container['featured'] = $data['featured'] ?? null;
        $this->container['versions'] = $data['versions'] ?? null;
        $this->container['instance_type_layouts'] = $data['instance_type_layouts'] ?? null;
        $this->container['image_path'] = $data['image_path'] ?? null;
        $this->container['dark_image_path'] = $data['dark_image_path'] ?? null;
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
     * Gets account
     *
     * @return \OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert|null
     */
    public function getAccount()
    {
        return $this->container['account'];
    }

    /**
     * Sets account
     *
     * @param \OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert|null $account account
     *
     * @return self
     */
    public function setAccount($account)
    {
        $this->container['account'] = $account;

        return $this;
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
     * @param string|null $name name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets labels
     *
     * @return string[]|null
     */
    public function getLabels()
    {
        return $this->container['labels'];
    }

    /**
     * Sets labels
     *
     * @param string[]|null $labels labels
     *
     * @return self
     */
    public function setLabels($labels)
    {
        $this->container['labels'] = $labels;

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
     * Gets description
     *
     * @return string|null
     */
    public function getDescription()
    {
        return $this->container['description'];
    }

    /**
     * Sets description
     *
     * @param string|null $description description
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

        return $this;
    }

    /**
     * Gets provision_type_code
     *
     * @return string|null
     */
    public function getProvisionTypeCode()
    {
        return $this->container['provision_type_code'];
    }

    /**
     * Sets provision_type_code
     *
     * @param string|null $provision_type_code provision_type_code
     *
     * @return self
     */
    public function setProvisionTypeCode($provision_type_code)
    {
        $this->container['provision_type_code'] = $provision_type_code;

        return $this;
    }

    /**
     * Gets category
     *
     * @return string|null
     */
    public function getCategory()
    {
        return $this->container['category'];
    }

    /**
     * Sets category
     *
     * @param string|null $category category
     *
     * @return self
     */
    public function setCategory($category)
    {
        $this->container['category'] = $category;

        return $this;
    }

    /**
     * Gets active
     *
     * @return bool|null
     */
    public function getActive()
    {
        return $this->container['active'];
    }

    /**
     * Sets active
     *
     * @param bool|null $active active
     *
     * @return self
     */
    public function setActive($active)
    {
        $this->container['active'] = $active;

        return $this;
    }

    /**
     * Gets environment_prefix
     *
     * @return string|null
     */
    public function getEnvironmentPrefix()
    {
        return $this->container['environment_prefix'];
    }

    /**
     * Sets environment_prefix
     *
     * @param string|null $environment_prefix environment_prefix
     *
     * @return self
     */
    public function setEnvironmentPrefix($environment_prefix)
    {
        $this->container['environment_prefix'] = $environment_prefix;

        return $this;
    }

    /**
     * Gets visibility
     *
     * @return string|null
     */
    public function getVisibility()
    {
        return $this->container['visibility'];
    }

    /**
     * Sets visibility
     *
     * @param string|null $visibility visibility
     *
     * @return self
     */
    public function setVisibility($visibility)
    {
        $this->container['visibility'] = $visibility;

        return $this;
    }

    /**
     * Gets featured
     *
     * @return bool|null
     */
    public function getFeatured()
    {
        return $this->container['featured'];
    }

    /**
     * Sets featured
     *
     * @param bool|null $featured featured
     *
     * @return self
     */
    public function setFeatured($featured)
    {
        $this->container['featured'] = $featured;

        return $this;
    }

    /**
     * Gets versions
     *
     * @return string[]|null
     */
    public function getVersions()
    {
        return $this->container['versions'];
    }

    /**
     * Sets versions
     *
     * @param string[]|null $versions versions
     *
     * @return self
     */
    public function setVersions($versions)
    {
        $this->container['versions'] = $versions;

        return $this;
    }

    /**
     * Gets instance_type_layouts
     *
     * @return \OpenAPI\Client\Model\InstanceTypesInstanceTypeLayouts[]|null
     */
    public function getInstanceTypeLayouts()
    {
        return $this->container['instance_type_layouts'];
    }

    /**
     * Sets instance_type_layouts
     *
     * @param \OpenAPI\Client\Model\InstanceTypesInstanceTypeLayouts[]|null $instance_type_layouts instance_type_layouts
     *
     * @return self
     */
    public function setInstanceTypeLayouts($instance_type_layouts)
    {
        $this->container['instance_type_layouts'] = $instance_type_layouts;

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
     * @param string|null $image_path Logo image URL
     *
     * @return self
     */
    public function setImagePath($image_path)
    {
        $this->container['image_path'] = $image_path;

        return $this;
    }

    /**
     * Gets dark_image_path
     *
     * @return string|null
     */
    public function getDarkImagePath()
    {
        return $this->container['dark_image_path'];
    }

    /**
     * Sets dark_image_path
     *
     * @param string|null $dark_image_path Dark logo image URL
     *
     * @return self
     */
    public function setDarkImagePath($dark_image_path)
    {
        $this->container['dark_image_path'] = $dark_image_path;

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

