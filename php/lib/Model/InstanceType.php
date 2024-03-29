<?php
/**
 * InstanceType
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
 * InstanceType Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InstanceType implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'instanceType';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'account' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'name' => 'string',
        'labels' => 'string[]',
        'code' => 'string',
        'description' => 'string',
        'provision_type_code' => 'string',
        'category' => 'string',
        'active' => 'bool',
        'has_provisioning_step' => 'bool',
        'has_deployment' => 'bool',
        'has_config' => 'bool',
        'has_settings' => 'bool',
        'has_auto_scale' => 'bool',
        'proxy_type' => 'string',
        'proxy_port' => 'string',
        'proxy_protocol' => 'string',
        'environment_prefix' => 'string',
        'backup_type' => 'string',
        'config' => 'object',
        'visibility' => 'string',
        'featured' => 'bool',
        'versions' => 'string[]',
        'instance_type_layouts' => '\OpenAPI\Client\Model\InstanceTypeLayout[]',
        'option_types' => '\OpenAPI\Client\Model\OptionType[]',
        'environment_variables' => 'object[]',
        'price_sets' => 'object[]',
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
        'has_provisioning_step' => null,
        'has_deployment' => null,
        'has_config' => null,
        'has_settings' => null,
        'has_auto_scale' => null,
        'proxy_type' => null,
        'proxy_port' => null,
        'proxy_protocol' => null,
        'environment_prefix' => null,
        'backup_type' => null,
        'config' => null,
        'visibility' => null,
        'featured' => null,
        'versions' => null,
        'instance_type_layouts' => null,
        'option_types' => null,
        'environment_variables' => null,
        'price_sets' => null,
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
        'has_provisioning_step' => 'hasProvisioningStep',
        'has_deployment' => 'hasDeployment',
        'has_config' => 'hasConfig',
        'has_settings' => 'hasSettings',
        'has_auto_scale' => 'hasAutoScale',
        'proxy_type' => 'proxyType',
        'proxy_port' => 'proxyPort',
        'proxy_protocol' => 'proxyProtocol',
        'environment_prefix' => 'environmentPrefix',
        'backup_type' => 'backupType',
        'config' => 'config',
        'visibility' => 'visibility',
        'featured' => 'featured',
        'versions' => 'versions',
        'instance_type_layouts' => 'instanceTypeLayouts',
        'option_types' => 'optionTypes',
        'environment_variables' => 'environmentVariables',
        'price_sets' => 'priceSets',
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
        'has_provisioning_step' => 'setHasProvisioningStep',
        'has_deployment' => 'setHasDeployment',
        'has_config' => 'setHasConfig',
        'has_settings' => 'setHasSettings',
        'has_auto_scale' => 'setHasAutoScale',
        'proxy_type' => 'setProxyType',
        'proxy_port' => 'setProxyPort',
        'proxy_protocol' => 'setProxyProtocol',
        'environment_prefix' => 'setEnvironmentPrefix',
        'backup_type' => 'setBackupType',
        'config' => 'setConfig',
        'visibility' => 'setVisibility',
        'featured' => 'setFeatured',
        'versions' => 'setVersions',
        'instance_type_layouts' => 'setInstanceTypeLayouts',
        'option_types' => 'setOptionTypes',
        'environment_variables' => 'setEnvironmentVariables',
        'price_sets' => 'setPriceSets',
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
        'has_provisioning_step' => 'getHasProvisioningStep',
        'has_deployment' => 'getHasDeployment',
        'has_config' => 'getHasConfig',
        'has_settings' => 'getHasSettings',
        'has_auto_scale' => 'getHasAutoScale',
        'proxy_type' => 'getProxyType',
        'proxy_port' => 'getProxyPort',
        'proxy_protocol' => 'getProxyProtocol',
        'environment_prefix' => 'getEnvironmentPrefix',
        'backup_type' => 'getBackupType',
        'config' => 'getConfig',
        'visibility' => 'getVisibility',
        'featured' => 'getFeatured',
        'versions' => 'getVersions',
        'instance_type_layouts' => 'getInstanceTypeLayouts',
        'option_types' => 'getOptionTypes',
        'environment_variables' => 'getEnvironmentVariables',
        'price_sets' => 'getPriceSets',
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
        $this->container['has_provisioning_step'] = $data['has_provisioning_step'] ?? null;
        $this->container['has_deployment'] = $data['has_deployment'] ?? null;
        $this->container['has_config'] = $data['has_config'] ?? null;
        $this->container['has_settings'] = $data['has_settings'] ?? null;
        $this->container['has_auto_scale'] = $data['has_auto_scale'] ?? null;
        $this->container['proxy_type'] = $data['proxy_type'] ?? null;
        $this->container['proxy_port'] = $data['proxy_port'] ?? null;
        $this->container['proxy_protocol'] = $data['proxy_protocol'] ?? null;
        $this->container['environment_prefix'] = $data['environment_prefix'] ?? null;
        $this->container['backup_type'] = $data['backup_type'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['visibility'] = $data['visibility'] ?? null;
        $this->container['featured'] = $data['featured'] ?? null;
        $this->container['versions'] = $data['versions'] ?? null;
        $this->container['instance_type_layouts'] = $data['instance_type_layouts'] ?? null;
        $this->container['option_types'] = $data['option_types'] ?? null;
        $this->container['environment_variables'] = $data['environment_variables'] ?? null;
        $this->container['price_sets'] = $data['price_sets'] ?? null;
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
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getAccount()
    {
        return $this->container['account'];
    }

    /**
     * Sets account
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $account account
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
     * Gets has_provisioning_step
     *
     * @return bool|null
     */
    public function getHasProvisioningStep()
    {
        return $this->container['has_provisioning_step'];
    }

    /**
     * Sets has_provisioning_step
     *
     * @param bool|null $has_provisioning_step has_provisioning_step
     *
     * @return self
     */
    public function setHasProvisioningStep($has_provisioning_step)
    {
        $this->container['has_provisioning_step'] = $has_provisioning_step;

        return $this;
    }

    /**
     * Gets has_deployment
     *
     * @return bool|null
     */
    public function getHasDeployment()
    {
        return $this->container['has_deployment'];
    }

    /**
     * Sets has_deployment
     *
     * @param bool|null $has_deployment has_deployment
     *
     * @return self
     */
    public function setHasDeployment($has_deployment)
    {
        $this->container['has_deployment'] = $has_deployment;

        return $this;
    }

    /**
     * Gets has_config
     *
     * @return bool|null
     */
    public function getHasConfig()
    {
        return $this->container['has_config'];
    }

    /**
     * Sets has_config
     *
     * @param bool|null $has_config has_config
     *
     * @return self
     */
    public function setHasConfig($has_config)
    {
        $this->container['has_config'] = $has_config;

        return $this;
    }

    /**
     * Gets has_settings
     *
     * @return bool|null
     */
    public function getHasSettings()
    {
        return $this->container['has_settings'];
    }

    /**
     * Sets has_settings
     *
     * @param bool|null $has_settings has_settings
     *
     * @return self
     */
    public function setHasSettings($has_settings)
    {
        $this->container['has_settings'] = $has_settings;

        return $this;
    }

    /**
     * Gets has_auto_scale
     *
     * @return bool|null
     */
    public function getHasAutoScale()
    {
        return $this->container['has_auto_scale'];
    }

    /**
     * Sets has_auto_scale
     *
     * @param bool|null $has_auto_scale has_auto_scale
     *
     * @return self
     */
    public function setHasAutoScale($has_auto_scale)
    {
        $this->container['has_auto_scale'] = $has_auto_scale;

        return $this;
    }

    /**
     * Gets proxy_type
     *
     * @return string|null
     */
    public function getProxyType()
    {
        return $this->container['proxy_type'];
    }

    /**
     * Sets proxy_type
     *
     * @param string|null $proxy_type proxy_type
     *
     * @return self
     */
    public function setProxyType($proxy_type)
    {
        $this->container['proxy_type'] = $proxy_type;

        return $this;
    }

    /**
     * Gets proxy_port
     *
     * @return string|null
     */
    public function getProxyPort()
    {
        return $this->container['proxy_port'];
    }

    /**
     * Sets proxy_port
     *
     * @param string|null $proxy_port proxy_port
     *
     * @return self
     */
    public function setProxyPort($proxy_port)
    {
        $this->container['proxy_port'] = $proxy_port;

        return $this;
    }

    /**
     * Gets proxy_protocol
     *
     * @return string|null
     */
    public function getProxyProtocol()
    {
        return $this->container['proxy_protocol'];
    }

    /**
     * Sets proxy_protocol
     *
     * @param string|null $proxy_protocol proxy_protocol
     *
     * @return self
     */
    public function setProxyProtocol($proxy_protocol)
    {
        $this->container['proxy_protocol'] = $proxy_protocol;

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
     * Gets backup_type
     *
     * @return string|null
     */
    public function getBackupType()
    {
        return $this->container['backup_type'];
    }

    /**
     * Sets backup_type
     *
     * @param string|null $backup_type backup_type
     *
     * @return self
     */
    public function setBackupType($backup_type)
    {
        $this->container['backup_type'] = $backup_type;

        return $this;
    }

    /**
     * Gets config
     *
     * @return object|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param object|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

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
     * @return \OpenAPI\Client\Model\InstanceTypeLayout[]|null
     */
    public function getInstanceTypeLayouts()
    {
        return $this->container['instance_type_layouts'];
    }

    /**
     * Sets instance_type_layouts
     *
     * @param \OpenAPI\Client\Model\InstanceTypeLayout[]|null $instance_type_layouts instance_type_layouts
     *
     * @return self
     */
    public function setInstanceTypeLayouts($instance_type_layouts)
    {
        $this->container['instance_type_layouts'] = $instance_type_layouts;

        return $this;
    }

    /**
     * Gets option_types
     *
     * @return \OpenAPI\Client\Model\OptionType[]|null
     */
    public function getOptionTypes()
    {
        return $this->container['option_types'];
    }

    /**
     * Sets option_types
     *
     * @param \OpenAPI\Client\Model\OptionType[]|null $option_types option_types
     *
     * @return self
     */
    public function setOptionTypes($option_types)
    {
        $this->container['option_types'] = $option_types;

        return $this;
    }

    /**
     * Gets environment_variables
     *
     * @return object[]|null
     */
    public function getEnvironmentVariables()
    {
        return $this->container['environment_variables'];
    }

    /**
     * Sets environment_variables
     *
     * @param object[]|null $environment_variables environment_variables
     *
     * @return self
     */
    public function setEnvironmentVariables($environment_variables)
    {
        $this->container['environment_variables'] = $environment_variables;

        return $this;
    }

    /**
     * Gets price_sets
     *
     * @return object[]|null
     */
    public function getPriceSets()
    {
        return $this->container['price_sets'];
    }

    /**
     * Sets price_sets
     *
     * @param object[]|null $price_sets price_sets
     *
     * @return self
     */
    public function setPriceSets($price_sets)
    {
        $this->container['price_sets'] = $price_sets;

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


