<?php
/**
 * ClusterLayoutComputeServers
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
 * ClusterLayoutComputeServers Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ClusterLayoutComputeServers implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'clusterLayout_computeServers';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'priority_order' => 'int',
        'node_count' => 'int',
        'node_type' => 'string',
        'min_node_count' => 'int',
        'max_node_count' => 'string',
        'dynamic_count' => 'bool',
        'install_container_runtime' => 'bool',
        'install_storage_runtime' => 'bool',
        'name' => 'string',
        'code' => 'string',
        'category' => 'string',
        'config' => 'string',
        'container_type' => '\OpenAPI\Client\Model\ClusterLayoutContainerType',
        'compute_server_type' => '\OpenAPI\Client\Model\ClusterLayoutComputeServerType',
        'provision_service' => 'string',
        'plan_category' => 'string',
        'name_prefix' => 'string',
        'name_suffix' => 'string',
        'force_name_index' => 'bool',
        'load_balance' => 'bool'
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
        'priority_order' => 'int64',
        'node_count' => 'int64',
        'node_type' => null,
        'min_node_count' => 'int64',
        'max_node_count' => null,
        'dynamic_count' => null,
        'install_container_runtime' => null,
        'install_storage_runtime' => null,
        'name' => null,
        'code' => null,
        'category' => null,
        'config' => null,
        'container_type' => null,
        'compute_server_type' => null,
        'provision_service' => null,
        'plan_category' => null,
        'name_prefix' => null,
        'name_suffix' => null,
        'force_name_index' => null,
        'load_balance' => null
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
        'priority_order' => 'priorityOrder',
        'node_count' => 'nodeCount',
        'node_type' => 'nodeType',
        'min_node_count' => 'minNodeCount',
        'max_node_count' => 'maxNodeCount',
        'dynamic_count' => 'dynamicCount',
        'install_container_runtime' => 'installContainerRuntime',
        'install_storage_runtime' => 'installStorageRuntime',
        'name' => 'name',
        'code' => 'code',
        'category' => 'category',
        'config' => 'config',
        'container_type' => 'containerType',
        'compute_server_type' => 'computeServerType',
        'provision_service' => 'provisionService',
        'plan_category' => 'planCategory',
        'name_prefix' => 'namePrefix',
        'name_suffix' => 'nameSuffix',
        'force_name_index' => 'forceNameIndex',
        'load_balance' => 'loadBalance'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'priority_order' => 'setPriorityOrder',
        'node_count' => 'setNodeCount',
        'node_type' => 'setNodeType',
        'min_node_count' => 'setMinNodeCount',
        'max_node_count' => 'setMaxNodeCount',
        'dynamic_count' => 'setDynamicCount',
        'install_container_runtime' => 'setInstallContainerRuntime',
        'install_storage_runtime' => 'setInstallStorageRuntime',
        'name' => 'setName',
        'code' => 'setCode',
        'category' => 'setCategory',
        'config' => 'setConfig',
        'container_type' => 'setContainerType',
        'compute_server_type' => 'setComputeServerType',
        'provision_service' => 'setProvisionService',
        'plan_category' => 'setPlanCategory',
        'name_prefix' => 'setNamePrefix',
        'name_suffix' => 'setNameSuffix',
        'force_name_index' => 'setForceNameIndex',
        'load_balance' => 'setLoadBalance'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'priority_order' => 'getPriorityOrder',
        'node_count' => 'getNodeCount',
        'node_type' => 'getNodeType',
        'min_node_count' => 'getMinNodeCount',
        'max_node_count' => 'getMaxNodeCount',
        'dynamic_count' => 'getDynamicCount',
        'install_container_runtime' => 'getInstallContainerRuntime',
        'install_storage_runtime' => 'getInstallStorageRuntime',
        'name' => 'getName',
        'code' => 'getCode',
        'category' => 'getCategory',
        'config' => 'getConfig',
        'container_type' => 'getContainerType',
        'compute_server_type' => 'getComputeServerType',
        'provision_service' => 'getProvisionService',
        'plan_category' => 'getPlanCategory',
        'name_prefix' => 'getNamePrefix',
        'name_suffix' => 'getNameSuffix',
        'force_name_index' => 'getForceNameIndex',
        'load_balance' => 'getLoadBalance'
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
        $this->container['priority_order'] = $data['priority_order'] ?? null;
        $this->container['node_count'] = $data['node_count'] ?? null;
        $this->container['node_type'] = $data['node_type'] ?? null;
        $this->container['min_node_count'] = $data['min_node_count'] ?? null;
        $this->container['max_node_count'] = $data['max_node_count'] ?? null;
        $this->container['dynamic_count'] = $data['dynamic_count'] ?? null;
        $this->container['install_container_runtime'] = $data['install_container_runtime'] ?? null;
        $this->container['install_storage_runtime'] = $data['install_storage_runtime'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['code'] = $data['code'] ?? null;
        $this->container['category'] = $data['category'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['container_type'] = $data['container_type'] ?? null;
        $this->container['compute_server_type'] = $data['compute_server_type'] ?? null;
        $this->container['provision_service'] = $data['provision_service'] ?? null;
        $this->container['plan_category'] = $data['plan_category'] ?? null;
        $this->container['name_prefix'] = $data['name_prefix'] ?? null;
        $this->container['name_suffix'] = $data['name_suffix'] ?? null;
        $this->container['force_name_index'] = $data['force_name_index'] ?? null;
        $this->container['load_balance'] = $data['load_balance'] ?? null;
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
     * Gets priority_order
     *
     * @return int|null
     */
    public function getPriorityOrder()
    {
        return $this->container['priority_order'];
    }

    /**
     * Sets priority_order
     *
     * @param int|null $priority_order priority_order
     *
     * @return self
     */
    public function setPriorityOrder($priority_order)
    {
        $this->container['priority_order'] = $priority_order;

        return $this;
    }

    /**
     * Gets node_count
     *
     * @return int|null
     */
    public function getNodeCount()
    {
        return $this->container['node_count'];
    }

    /**
     * Sets node_count
     *
     * @param int|null $node_count node_count
     *
     * @return self
     */
    public function setNodeCount($node_count)
    {
        $this->container['node_count'] = $node_count;

        return $this;
    }

    /**
     * Gets node_type
     *
     * @return string|null
     */
    public function getNodeType()
    {
        return $this->container['node_type'];
    }

    /**
     * Sets node_type
     *
     * @param string|null $node_type node_type
     *
     * @return self
     */
    public function setNodeType($node_type)
    {
        $this->container['node_type'] = $node_type;

        return $this;
    }

    /**
     * Gets min_node_count
     *
     * @return int|null
     */
    public function getMinNodeCount()
    {
        return $this->container['min_node_count'];
    }

    /**
     * Sets min_node_count
     *
     * @param int|null $min_node_count min_node_count
     *
     * @return self
     */
    public function setMinNodeCount($min_node_count)
    {
        $this->container['min_node_count'] = $min_node_count;

        return $this;
    }

    /**
     * Gets max_node_count
     *
     * @return string|null
     */
    public function getMaxNodeCount()
    {
        return $this->container['max_node_count'];
    }

    /**
     * Sets max_node_count
     *
     * @param string|null $max_node_count max_node_count
     *
     * @return self
     */
    public function setMaxNodeCount($max_node_count)
    {
        $this->container['max_node_count'] = $max_node_count;

        return $this;
    }

    /**
     * Gets dynamic_count
     *
     * @return bool|null
     */
    public function getDynamicCount()
    {
        return $this->container['dynamic_count'];
    }

    /**
     * Sets dynamic_count
     *
     * @param bool|null $dynamic_count dynamic_count
     *
     * @return self
     */
    public function setDynamicCount($dynamic_count)
    {
        $this->container['dynamic_count'] = $dynamic_count;

        return $this;
    }

    /**
     * Gets install_container_runtime
     *
     * @return bool|null
     */
    public function getInstallContainerRuntime()
    {
        return $this->container['install_container_runtime'];
    }

    /**
     * Sets install_container_runtime
     *
     * @param bool|null $install_container_runtime install_container_runtime
     *
     * @return self
     */
    public function setInstallContainerRuntime($install_container_runtime)
    {
        $this->container['install_container_runtime'] = $install_container_runtime;

        return $this;
    }

    /**
     * Gets install_storage_runtime
     *
     * @return bool|null
     */
    public function getInstallStorageRuntime()
    {
        return $this->container['install_storage_runtime'];
    }

    /**
     * Sets install_storage_runtime
     *
     * @param bool|null $install_storage_runtime install_storage_runtime
     *
     * @return self
     */
    public function setInstallStorageRuntime($install_storage_runtime)
    {
        $this->container['install_storage_runtime'] = $install_storage_runtime;

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
     * Gets config
     *
     * @return string|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param string|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

        return $this;
    }

    /**
     * Gets container_type
     *
     * @return \OpenAPI\Client\Model\ClusterLayoutContainerType|null
     */
    public function getContainerType()
    {
        return $this->container['container_type'];
    }

    /**
     * Sets container_type
     *
     * @param \OpenAPI\Client\Model\ClusterLayoutContainerType|null $container_type container_type
     *
     * @return self
     */
    public function setContainerType($container_type)
    {
        $this->container['container_type'] = $container_type;

        return $this;
    }

    /**
     * Gets compute_server_type
     *
     * @return \OpenAPI\Client\Model\ClusterLayoutComputeServerType|null
     */
    public function getComputeServerType()
    {
        return $this->container['compute_server_type'];
    }

    /**
     * Sets compute_server_type
     *
     * @param \OpenAPI\Client\Model\ClusterLayoutComputeServerType|null $compute_server_type compute_server_type
     *
     * @return self
     */
    public function setComputeServerType($compute_server_type)
    {
        $this->container['compute_server_type'] = $compute_server_type;

        return $this;
    }

    /**
     * Gets provision_service
     *
     * @return string|null
     */
    public function getProvisionService()
    {
        return $this->container['provision_service'];
    }

    /**
     * Sets provision_service
     *
     * @param string|null $provision_service provision_service
     *
     * @return self
     */
    public function setProvisionService($provision_service)
    {
        $this->container['provision_service'] = $provision_service;

        return $this;
    }

    /**
     * Gets plan_category
     *
     * @return string|null
     */
    public function getPlanCategory()
    {
        return $this->container['plan_category'];
    }

    /**
     * Sets plan_category
     *
     * @param string|null $plan_category plan_category
     *
     * @return self
     */
    public function setPlanCategory($plan_category)
    {
        $this->container['plan_category'] = $plan_category;

        return $this;
    }

    /**
     * Gets name_prefix
     *
     * @return string|null
     */
    public function getNamePrefix()
    {
        return $this->container['name_prefix'];
    }

    /**
     * Sets name_prefix
     *
     * @param string|null $name_prefix name_prefix
     *
     * @return self
     */
    public function setNamePrefix($name_prefix)
    {
        $this->container['name_prefix'] = $name_prefix;

        return $this;
    }

    /**
     * Gets name_suffix
     *
     * @return string|null
     */
    public function getNameSuffix()
    {
        return $this->container['name_suffix'];
    }

    /**
     * Sets name_suffix
     *
     * @param string|null $name_suffix name_suffix
     *
     * @return self
     */
    public function setNameSuffix($name_suffix)
    {
        $this->container['name_suffix'] = $name_suffix;

        return $this;
    }

    /**
     * Gets force_name_index
     *
     * @return bool|null
     */
    public function getForceNameIndex()
    {
        return $this->container['force_name_index'];
    }

    /**
     * Sets force_name_index
     *
     * @param bool|null $force_name_index force_name_index
     *
     * @return self
     */
    public function setForceNameIndex($force_name_index)
    {
        $this->container['force_name_index'] = $force_name_index;

        return $this;
    }

    /**
     * Gets load_balance
     *
     * @return bool|null
     */
    public function getLoadBalance()
    {
        return $this->container['load_balance'];
    }

    /**
     * Sets load_balance
     *
     * @param bool|null $load_balance load_balance
     *
     * @return self
     */
    public function setLoadBalance($load_balance)
    {
        $this->container['load_balance'] = $load_balance;

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


