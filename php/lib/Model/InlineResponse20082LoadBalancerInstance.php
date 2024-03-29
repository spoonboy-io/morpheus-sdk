<?php
/**
 * InlineResponse20082LoadBalancerInstance
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
 * InlineResponse20082LoadBalancerInstance Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InlineResponse20082LoadBalancerInstance implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'inline_response_200_82_loadBalancerInstance';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'load_balancer' => '\OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancer',
        'instance' => 'string',
        'description' => 'string',
        'internal_id' => 'string',
        'external_id' => 'string',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime',
        'active' => 'bool',
        'sticky' => 'bool',
        'ssl_enabled' => 'string',
        'external_address' => 'bool',
        'backend_port' => 'string',
        'vip_type' => 'string',
        'vip_address' => 'string',
        'vip_hostname' => 'string',
        'vip_protocol' => 'string',
        'vip_scheme' => 'string',
        'vip_mode' => 'string',
        'vip_name' => 'string',
        'vip_port' => 'int',
        'vip_sticky' => 'string',
        'vip_balance' => 'string',
        'service_port' => 'string',
        'source_address' => 'string',
        'ssl_cert' => '\OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert',
        'ssl_mode' => 'string',
        'ssl_redirect_mode' => 'string',
        'vip_shared' => 'bool',
        'vip_direct_address' => 'string',
        'server_name' => 'string',
        'pool_name' => 'string',
        'removing' => 'bool',
        'vip_source' => 'string',
        'extra_config' => 'string',
        'service_access' => 'string',
        'network_id' => 'string',
        'subnet_id' => 'string',
        'external_port_id' => 'string',
        'status' => 'string',
        'vip_status' => 'string'
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
        'load_balancer' => null,
        'instance' => null,
        'description' => null,
        'internal_id' => null,
        'external_id' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time',
        'active' => null,
        'sticky' => null,
        'ssl_enabled' => null,
        'external_address' => null,
        'backend_port' => null,
        'vip_type' => null,
        'vip_address' => null,
        'vip_hostname' => null,
        'vip_protocol' => null,
        'vip_scheme' => null,
        'vip_mode' => null,
        'vip_name' => null,
        'vip_port' => 'int64',
        'vip_sticky' => null,
        'vip_balance' => null,
        'service_port' => null,
        'source_address' => null,
        'ssl_cert' => null,
        'ssl_mode' => null,
        'ssl_redirect_mode' => null,
        'vip_shared' => null,
        'vip_direct_address' => null,
        'server_name' => null,
        'pool_name' => null,
        'removing' => null,
        'vip_source' => null,
        'extra_config' => null,
        'service_access' => null,
        'network_id' => null,
        'subnet_id' => null,
        'external_port_id' => null,
        'status' => null,
        'vip_status' => null
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
        'load_balancer' => 'loadBalancer',
        'instance' => 'instance',
        'description' => 'description',
        'internal_id' => 'internalId',
        'external_id' => 'externalId',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated',
        'active' => 'active',
        'sticky' => 'sticky',
        'ssl_enabled' => 'sslEnabled',
        'external_address' => 'externalAddress',
        'backend_port' => 'backendPort',
        'vip_type' => 'vipType',
        'vip_address' => 'vipAddress',
        'vip_hostname' => 'vipHostname',
        'vip_protocol' => 'vipProtocol',
        'vip_scheme' => 'vipScheme',
        'vip_mode' => 'vipMode',
        'vip_name' => 'vipName',
        'vip_port' => 'vipPort',
        'vip_sticky' => 'vipSticky',
        'vip_balance' => 'vipBalance',
        'service_port' => 'servicePort',
        'source_address' => 'sourceAddress',
        'ssl_cert' => 'sslCert',
        'ssl_mode' => 'sslMode',
        'ssl_redirect_mode' => 'sslRedirectMode',
        'vip_shared' => 'vipShared',
        'vip_direct_address' => 'vipDirectAddress',
        'server_name' => 'serverName',
        'pool_name' => 'poolName',
        'removing' => 'removing',
        'vip_source' => 'vipSource',
        'extra_config' => 'extraConfig',
        'service_access' => 'serviceAccess',
        'network_id' => 'networkId',
        'subnet_id' => 'subnetId',
        'external_port_id' => 'externalPortId',
        'status' => 'status',
        'vip_status' => 'vipStatus'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'load_balancer' => 'setLoadBalancer',
        'instance' => 'setInstance',
        'description' => 'setDescription',
        'internal_id' => 'setInternalId',
        'external_id' => 'setExternalId',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated',
        'active' => 'setActive',
        'sticky' => 'setSticky',
        'ssl_enabled' => 'setSslEnabled',
        'external_address' => 'setExternalAddress',
        'backend_port' => 'setBackendPort',
        'vip_type' => 'setVipType',
        'vip_address' => 'setVipAddress',
        'vip_hostname' => 'setVipHostname',
        'vip_protocol' => 'setVipProtocol',
        'vip_scheme' => 'setVipScheme',
        'vip_mode' => 'setVipMode',
        'vip_name' => 'setVipName',
        'vip_port' => 'setVipPort',
        'vip_sticky' => 'setVipSticky',
        'vip_balance' => 'setVipBalance',
        'service_port' => 'setServicePort',
        'source_address' => 'setSourceAddress',
        'ssl_cert' => 'setSslCert',
        'ssl_mode' => 'setSslMode',
        'ssl_redirect_mode' => 'setSslRedirectMode',
        'vip_shared' => 'setVipShared',
        'vip_direct_address' => 'setVipDirectAddress',
        'server_name' => 'setServerName',
        'pool_name' => 'setPoolName',
        'removing' => 'setRemoving',
        'vip_source' => 'setVipSource',
        'extra_config' => 'setExtraConfig',
        'service_access' => 'setServiceAccess',
        'network_id' => 'setNetworkId',
        'subnet_id' => 'setSubnetId',
        'external_port_id' => 'setExternalPortId',
        'status' => 'setStatus',
        'vip_status' => 'setVipStatus'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'load_balancer' => 'getLoadBalancer',
        'instance' => 'getInstance',
        'description' => 'getDescription',
        'internal_id' => 'getInternalId',
        'external_id' => 'getExternalId',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated',
        'active' => 'getActive',
        'sticky' => 'getSticky',
        'ssl_enabled' => 'getSslEnabled',
        'external_address' => 'getExternalAddress',
        'backend_port' => 'getBackendPort',
        'vip_type' => 'getVipType',
        'vip_address' => 'getVipAddress',
        'vip_hostname' => 'getVipHostname',
        'vip_protocol' => 'getVipProtocol',
        'vip_scheme' => 'getVipScheme',
        'vip_mode' => 'getVipMode',
        'vip_name' => 'getVipName',
        'vip_port' => 'getVipPort',
        'vip_sticky' => 'getVipSticky',
        'vip_balance' => 'getVipBalance',
        'service_port' => 'getServicePort',
        'source_address' => 'getSourceAddress',
        'ssl_cert' => 'getSslCert',
        'ssl_mode' => 'getSslMode',
        'ssl_redirect_mode' => 'getSslRedirectMode',
        'vip_shared' => 'getVipShared',
        'vip_direct_address' => 'getVipDirectAddress',
        'server_name' => 'getServerName',
        'pool_name' => 'getPoolName',
        'removing' => 'getRemoving',
        'vip_source' => 'getVipSource',
        'extra_config' => 'getExtraConfig',
        'service_access' => 'getServiceAccess',
        'network_id' => 'getNetworkId',
        'subnet_id' => 'getSubnetId',
        'external_port_id' => 'getExternalPortId',
        'status' => 'getStatus',
        'vip_status' => 'getVipStatus'
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
        $this->container['load_balancer'] = $data['load_balancer'] ?? null;
        $this->container['instance'] = $data['instance'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['internal_id'] = $data['internal_id'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
        $this->container['active'] = $data['active'] ?? null;
        $this->container['sticky'] = $data['sticky'] ?? null;
        $this->container['ssl_enabled'] = $data['ssl_enabled'] ?? null;
        $this->container['external_address'] = $data['external_address'] ?? null;
        $this->container['backend_port'] = $data['backend_port'] ?? null;
        $this->container['vip_type'] = $data['vip_type'] ?? null;
        $this->container['vip_address'] = $data['vip_address'] ?? null;
        $this->container['vip_hostname'] = $data['vip_hostname'] ?? null;
        $this->container['vip_protocol'] = $data['vip_protocol'] ?? null;
        $this->container['vip_scheme'] = $data['vip_scheme'] ?? null;
        $this->container['vip_mode'] = $data['vip_mode'] ?? null;
        $this->container['vip_name'] = $data['vip_name'] ?? null;
        $this->container['vip_port'] = $data['vip_port'] ?? null;
        $this->container['vip_sticky'] = $data['vip_sticky'] ?? null;
        $this->container['vip_balance'] = $data['vip_balance'] ?? null;
        $this->container['service_port'] = $data['service_port'] ?? null;
        $this->container['source_address'] = $data['source_address'] ?? null;
        $this->container['ssl_cert'] = $data['ssl_cert'] ?? null;
        $this->container['ssl_mode'] = $data['ssl_mode'] ?? null;
        $this->container['ssl_redirect_mode'] = $data['ssl_redirect_mode'] ?? null;
        $this->container['vip_shared'] = $data['vip_shared'] ?? null;
        $this->container['vip_direct_address'] = $data['vip_direct_address'] ?? null;
        $this->container['server_name'] = $data['server_name'] ?? null;
        $this->container['pool_name'] = $data['pool_name'] ?? null;
        $this->container['removing'] = $data['removing'] ?? null;
        $this->container['vip_source'] = $data['vip_source'] ?? null;
        $this->container['extra_config'] = $data['extra_config'] ?? null;
        $this->container['service_access'] = $data['service_access'] ?? null;
        $this->container['network_id'] = $data['network_id'] ?? null;
        $this->container['subnet_id'] = $data['subnet_id'] ?? null;
        $this->container['external_port_id'] = $data['external_port_id'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['vip_status'] = $data['vip_status'] ?? null;
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
     * Gets load_balancer
     *
     * @return \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancer|null
     */
    public function getLoadBalancer()
    {
        return $this->container['load_balancer'];
    }

    /**
     * Sets load_balancer
     *
     * @param \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancer|null $load_balancer load_balancer
     *
     * @return self
     */
    public function setLoadBalancer($load_balancer)
    {
        $this->container['load_balancer'] = $load_balancer;

        return $this;
    }

    /**
     * Gets instance
     *
     * @return string|null
     */
    public function getInstance()
    {
        return $this->container['instance'];
    }

    /**
     * Sets instance
     *
     * @param string|null $instance instance
     *
     * @return self
     */
    public function setInstance($instance)
    {
        $this->container['instance'] = $instance;

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
     * Gets date_created
     *
     * @return \DateTime|null
     */
    public function getDateCreated()
    {
        return $this->container['date_created'];
    }

    /**
     * Sets date_created
     *
     * @param \DateTime|null $date_created date_created
     *
     * @return self
     */
    public function setDateCreated($date_created)
    {
        $this->container['date_created'] = $date_created;

        return $this;
    }

    /**
     * Gets last_updated
     *
     * @return \DateTime|null
     */
    public function getLastUpdated()
    {
        return $this->container['last_updated'];
    }

    /**
     * Sets last_updated
     *
     * @param \DateTime|null $last_updated last_updated
     *
     * @return self
     */
    public function setLastUpdated($last_updated)
    {
        $this->container['last_updated'] = $last_updated;

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
     * Gets sticky
     *
     * @return bool|null
     */
    public function getSticky()
    {
        return $this->container['sticky'];
    }

    /**
     * Sets sticky
     *
     * @param bool|null $sticky sticky
     *
     * @return self
     */
    public function setSticky($sticky)
    {
        $this->container['sticky'] = $sticky;

        return $this;
    }

    /**
     * Gets ssl_enabled
     *
     * @return string|null
     */
    public function getSslEnabled()
    {
        return $this->container['ssl_enabled'];
    }

    /**
     * Sets ssl_enabled
     *
     * @param string|null $ssl_enabled ssl_enabled
     *
     * @return self
     */
    public function setSslEnabled($ssl_enabled)
    {
        $this->container['ssl_enabled'] = $ssl_enabled;

        return $this;
    }

    /**
     * Gets external_address
     *
     * @return bool|null
     */
    public function getExternalAddress()
    {
        return $this->container['external_address'];
    }

    /**
     * Sets external_address
     *
     * @param bool|null $external_address external_address
     *
     * @return self
     */
    public function setExternalAddress($external_address)
    {
        $this->container['external_address'] = $external_address;

        return $this;
    }

    /**
     * Gets backend_port
     *
     * @return string|null
     */
    public function getBackendPort()
    {
        return $this->container['backend_port'];
    }

    /**
     * Sets backend_port
     *
     * @param string|null $backend_port backend_port
     *
     * @return self
     */
    public function setBackendPort($backend_port)
    {
        $this->container['backend_port'] = $backend_port;

        return $this;
    }

    /**
     * Gets vip_type
     *
     * @return string|null
     */
    public function getVipType()
    {
        return $this->container['vip_type'];
    }

    /**
     * Sets vip_type
     *
     * @param string|null $vip_type vip_type
     *
     * @return self
     */
    public function setVipType($vip_type)
    {
        $this->container['vip_type'] = $vip_type;

        return $this;
    }

    /**
     * Gets vip_address
     *
     * @return string|null
     */
    public function getVipAddress()
    {
        return $this->container['vip_address'];
    }

    /**
     * Sets vip_address
     *
     * @param string|null $vip_address vip_address
     *
     * @return self
     */
    public function setVipAddress($vip_address)
    {
        $this->container['vip_address'] = $vip_address;

        return $this;
    }

    /**
     * Gets vip_hostname
     *
     * @return string|null
     */
    public function getVipHostname()
    {
        return $this->container['vip_hostname'];
    }

    /**
     * Sets vip_hostname
     *
     * @param string|null $vip_hostname vip_hostname
     *
     * @return self
     */
    public function setVipHostname($vip_hostname)
    {
        $this->container['vip_hostname'] = $vip_hostname;

        return $this;
    }

    /**
     * Gets vip_protocol
     *
     * @return string|null
     */
    public function getVipProtocol()
    {
        return $this->container['vip_protocol'];
    }

    /**
     * Sets vip_protocol
     *
     * @param string|null $vip_protocol vip_protocol
     *
     * @return self
     */
    public function setVipProtocol($vip_protocol)
    {
        $this->container['vip_protocol'] = $vip_protocol;

        return $this;
    }

    /**
     * Gets vip_scheme
     *
     * @return string|null
     */
    public function getVipScheme()
    {
        return $this->container['vip_scheme'];
    }

    /**
     * Sets vip_scheme
     *
     * @param string|null $vip_scheme vip_scheme
     *
     * @return self
     */
    public function setVipScheme($vip_scheme)
    {
        $this->container['vip_scheme'] = $vip_scheme;

        return $this;
    }

    /**
     * Gets vip_mode
     *
     * @return string|null
     */
    public function getVipMode()
    {
        return $this->container['vip_mode'];
    }

    /**
     * Sets vip_mode
     *
     * @param string|null $vip_mode vip_mode
     *
     * @return self
     */
    public function setVipMode($vip_mode)
    {
        $this->container['vip_mode'] = $vip_mode;

        return $this;
    }

    /**
     * Gets vip_name
     *
     * @return string|null
     */
    public function getVipName()
    {
        return $this->container['vip_name'];
    }

    /**
     * Sets vip_name
     *
     * @param string|null $vip_name vip_name
     *
     * @return self
     */
    public function setVipName($vip_name)
    {
        $this->container['vip_name'] = $vip_name;

        return $this;
    }

    /**
     * Gets vip_port
     *
     * @return int|null
     */
    public function getVipPort()
    {
        return $this->container['vip_port'];
    }

    /**
     * Sets vip_port
     *
     * @param int|null $vip_port vip_port
     *
     * @return self
     */
    public function setVipPort($vip_port)
    {
        $this->container['vip_port'] = $vip_port;

        return $this;
    }

    /**
     * Gets vip_sticky
     *
     * @return string|null
     */
    public function getVipSticky()
    {
        return $this->container['vip_sticky'];
    }

    /**
     * Sets vip_sticky
     *
     * @param string|null $vip_sticky vip_sticky
     *
     * @return self
     */
    public function setVipSticky($vip_sticky)
    {
        $this->container['vip_sticky'] = $vip_sticky;

        return $this;
    }

    /**
     * Gets vip_balance
     *
     * @return string|null
     */
    public function getVipBalance()
    {
        return $this->container['vip_balance'];
    }

    /**
     * Sets vip_balance
     *
     * @param string|null $vip_balance vip_balance
     *
     * @return self
     */
    public function setVipBalance($vip_balance)
    {
        $this->container['vip_balance'] = $vip_balance;

        return $this;
    }

    /**
     * Gets service_port
     *
     * @return string|null
     */
    public function getServicePort()
    {
        return $this->container['service_port'];
    }

    /**
     * Sets service_port
     *
     * @param string|null $service_port service_port
     *
     * @return self
     */
    public function setServicePort($service_port)
    {
        $this->container['service_port'] = $service_port;

        return $this;
    }

    /**
     * Gets source_address
     *
     * @return string|null
     */
    public function getSourceAddress()
    {
        return $this->container['source_address'];
    }

    /**
     * Sets source_address
     *
     * @param string|null $source_address source_address
     *
     * @return self
     */
    public function setSourceAddress($source_address)
    {
        $this->container['source_address'] = $source_address;

        return $this;
    }

    /**
     * Gets ssl_cert
     *
     * @return \OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert|null
     */
    public function getSslCert()
    {
        return $this->container['ssl_cert'];
    }

    /**
     * Sets ssl_cert
     *
     * @param \OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert|null $ssl_cert ssl_cert
     *
     * @return self
     */
    public function setSslCert($ssl_cert)
    {
        $this->container['ssl_cert'] = $ssl_cert;

        return $this;
    }

    /**
     * Gets ssl_mode
     *
     * @return string|null
     */
    public function getSslMode()
    {
        return $this->container['ssl_mode'];
    }

    /**
     * Sets ssl_mode
     *
     * @param string|null $ssl_mode ssl_mode
     *
     * @return self
     */
    public function setSslMode($ssl_mode)
    {
        $this->container['ssl_mode'] = $ssl_mode;

        return $this;
    }

    /**
     * Gets ssl_redirect_mode
     *
     * @return string|null
     */
    public function getSslRedirectMode()
    {
        return $this->container['ssl_redirect_mode'];
    }

    /**
     * Sets ssl_redirect_mode
     *
     * @param string|null $ssl_redirect_mode ssl_redirect_mode
     *
     * @return self
     */
    public function setSslRedirectMode($ssl_redirect_mode)
    {
        $this->container['ssl_redirect_mode'] = $ssl_redirect_mode;

        return $this;
    }

    /**
     * Gets vip_shared
     *
     * @return bool|null
     */
    public function getVipShared()
    {
        return $this->container['vip_shared'];
    }

    /**
     * Sets vip_shared
     *
     * @param bool|null $vip_shared vip_shared
     *
     * @return self
     */
    public function setVipShared($vip_shared)
    {
        $this->container['vip_shared'] = $vip_shared;

        return $this;
    }

    /**
     * Gets vip_direct_address
     *
     * @return string|null
     */
    public function getVipDirectAddress()
    {
        return $this->container['vip_direct_address'];
    }

    /**
     * Sets vip_direct_address
     *
     * @param string|null $vip_direct_address vip_direct_address
     *
     * @return self
     */
    public function setVipDirectAddress($vip_direct_address)
    {
        $this->container['vip_direct_address'] = $vip_direct_address;

        return $this;
    }

    /**
     * Gets server_name
     *
     * @return string|null
     */
    public function getServerName()
    {
        return $this->container['server_name'];
    }

    /**
     * Sets server_name
     *
     * @param string|null $server_name server_name
     *
     * @return self
     */
    public function setServerName($server_name)
    {
        $this->container['server_name'] = $server_name;

        return $this;
    }

    /**
     * Gets pool_name
     *
     * @return string|null
     */
    public function getPoolName()
    {
        return $this->container['pool_name'];
    }

    /**
     * Sets pool_name
     *
     * @param string|null $pool_name pool_name
     *
     * @return self
     */
    public function setPoolName($pool_name)
    {
        $this->container['pool_name'] = $pool_name;

        return $this;
    }

    /**
     * Gets removing
     *
     * @return bool|null
     */
    public function getRemoving()
    {
        return $this->container['removing'];
    }

    /**
     * Sets removing
     *
     * @param bool|null $removing removing
     *
     * @return self
     */
    public function setRemoving($removing)
    {
        $this->container['removing'] = $removing;

        return $this;
    }

    /**
     * Gets vip_source
     *
     * @return string|null
     */
    public function getVipSource()
    {
        return $this->container['vip_source'];
    }

    /**
     * Sets vip_source
     *
     * @param string|null $vip_source vip_source
     *
     * @return self
     */
    public function setVipSource($vip_source)
    {
        $this->container['vip_source'] = $vip_source;

        return $this;
    }

    /**
     * Gets extra_config
     *
     * @return string|null
     */
    public function getExtraConfig()
    {
        return $this->container['extra_config'];
    }

    /**
     * Sets extra_config
     *
     * @param string|null $extra_config extra_config
     *
     * @return self
     */
    public function setExtraConfig($extra_config)
    {
        $this->container['extra_config'] = $extra_config;

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
     * @param string|null $service_access service_access
     *
     * @return self
     */
    public function setServiceAccess($service_access)
    {
        $this->container['service_access'] = $service_access;

        return $this;
    }

    /**
     * Gets network_id
     *
     * @return string|null
     */
    public function getNetworkId()
    {
        return $this->container['network_id'];
    }

    /**
     * Sets network_id
     *
     * @param string|null $network_id network_id
     *
     * @return self
     */
    public function setNetworkId($network_id)
    {
        $this->container['network_id'] = $network_id;

        return $this;
    }

    /**
     * Gets subnet_id
     *
     * @return string|null
     */
    public function getSubnetId()
    {
        return $this->container['subnet_id'];
    }

    /**
     * Sets subnet_id
     *
     * @param string|null $subnet_id subnet_id
     *
     * @return self
     */
    public function setSubnetId($subnet_id)
    {
        $this->container['subnet_id'] = $subnet_id;

        return $this;
    }

    /**
     * Gets external_port_id
     *
     * @return string|null
     */
    public function getExternalPortId()
    {
        return $this->container['external_port_id'];
    }

    /**
     * Sets external_port_id
     *
     * @param string|null $external_port_id external_port_id
     *
     * @return self
     */
    public function setExternalPortId($external_port_id)
    {
        $this->container['external_port_id'] = $external_port_id;

        return $this;
    }

    /**
     * Gets status
     *
     * @return string|null
     */
    public function getStatus()
    {
        return $this->container['status'];
    }

    /**
     * Sets status
     *
     * @param string|null $status status
     *
     * @return self
     */
    public function setStatus($status)
    {
        $this->container['status'] = $status;

        return $this;
    }

    /**
     * Gets vip_status
     *
     * @return string|null
     */
    public function getVipStatus()
    {
        return $this->container['vip_status'];
    }

    /**
     * Sets vip_status
     *
     * @param string|null $vip_status vip_status
     *
     * @return self
     */
    public function setVipStatus($vip_status)
    {
        $this->container['vip_status'] = $vip_status;

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


