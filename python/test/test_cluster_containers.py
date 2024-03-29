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
from openapi_client.model.cluster_containers_available_actions_inner import ClusterContainersAvailableActionsInner
from openapi_client.model.cluster_containers_container_type import ClusterContainersContainerType
from openapi_client.model.cluster_containers_container_type_set import ClusterContainersContainerTypeSet
from openapi_client.model.cluster_containers_plan import ClusterContainersPlan
from openapi_client.model.cluster_containers_stats import ClusterContainersStats
from openapi_client.model.list_deploys200_response_all_of_app_deploys_inner_instance import ListDeploys200ResponseAllOfAppDeploysInnerInstance
globals()['ClusterContainersAvailableActionsInner'] = ClusterContainersAvailableActionsInner
globals()['ClusterContainersContainerType'] = ClusterContainersContainerType
globals()['ClusterContainersContainerTypeSet'] = ClusterContainersContainerTypeSet
globals()['ClusterContainersPlan'] = ClusterContainersPlan
globals()['ClusterContainersStats'] = ClusterContainersStats
globals()['ListDeploys200ResponseAllOfAppDeploysInnerInstance'] = ListDeploys200ResponseAllOfAppDeploysInnerInstance
from openapi_client.model.cluster_containers import ClusterContainers


class TestClusterContainers(unittest.TestCase):
    """ClusterContainers unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testClusterContainers(self):
        """Test ClusterContainers"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ClusterContainers()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
