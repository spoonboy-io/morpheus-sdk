"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.creds import Creds
from openapi_client.model.integration_ansible import IntegrationAnsible
from openapi_client.model.integration_ansible_tower import IntegrationAnsibleTower
from openapi_client.model.integration_bind_dns import IntegrationBindDNS
from openapi_client.model.integration_chef import IntegrationChef
from openapi_client.model.integration_cherwell import IntegrationCherwell
from openapi_client.model.integration_cypher import IntegrationCypher
from openapi_client.model.integration_docker_repo import IntegrationDockerRepo
from openapi_client.model.integration_git_hub import IntegrationGitHub
from openapi_client.model.integration_git_repo import IntegrationGitRepo
from openapi_client.model.integration_microsoft_dns import IntegrationMicrosoftDNS
from openapi_client.model.integration_power_dns import IntegrationPowerDNS
from openapi_client.model.integration_puppet import IntegrationPuppet
from openapi_client.model.integration_remedy import IntegrationRemedy
from openapi_client.model.integration_route53 import IntegrationRoute53
from openapi_client.model.integration_salt_master import IntegrationSaltMaster
from openapi_client.model.integration_snow import IntegrationSNOW
from openapi_client.model.integrationv_ro import IntegrationvRO
from openapi_client.model.list_deploys200_response_all_of_app_deploys_inner_instance import ListDeploys200ResponseAllOfAppDeploysInnerInstance
from openapi_client.model.list_load_balancer_monitors200_response_all_of_load_balancer_monitors_inner_load_balancer_type import ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType
globals()['Creds'] = Creds
globals()['IntegrationAnsible'] = IntegrationAnsible
globals()['IntegrationAnsibleTower'] = IntegrationAnsibleTower
globals()['IntegrationBindDNS'] = IntegrationBindDNS
globals()['IntegrationChef'] = IntegrationChef
globals()['IntegrationCherwell'] = IntegrationCherwell
globals()['IntegrationCypher'] = IntegrationCypher
globals()['IntegrationDockerRepo'] = IntegrationDockerRepo
globals()['IntegrationGitHub'] = IntegrationGitHub
globals()['IntegrationGitRepo'] = IntegrationGitRepo
globals()['IntegrationMicrosoftDNS'] = IntegrationMicrosoftDNS
globals()['IntegrationPowerDNS'] = IntegrationPowerDNS
globals()['IntegrationPuppet'] = IntegrationPuppet
globals()['IntegrationRemedy'] = IntegrationRemedy
globals()['IntegrationRoute53'] = IntegrationRoute53
globals()['IntegrationSNOW'] = IntegrationSNOW
globals()['IntegrationSaltMaster'] = IntegrationSaltMaster
globals()['IntegrationvRO'] = IntegrationvRO
globals()['ListDeploys200ResponseAllOfAppDeploysInnerInstance'] = ListDeploys200ResponseAllOfAppDeploysInnerInstance
globals()['ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType'] = ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType
from openapi_client.model.add_integrations200_response_all_of_integration import AddIntegrations200ResponseAllOfIntegration


class TestAddIntegrations200ResponseAllOfIntegration(unittest.TestCase):
    """AddIntegrations200ResponseAllOfIntegration unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testAddIntegrations200ResponseAllOfIntegration(self):
        """Test AddIntegrations200ResponseAllOfIntegration"""
        # FIXME: construct object with mandatory attributes with example values
        # model = AddIntegrations200ResponseAllOfIntegration()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
