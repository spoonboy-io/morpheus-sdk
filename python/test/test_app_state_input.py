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
from openapi_client.model.app_state_input_data_inner import AppStateInputDataInner
from openapi_client.model.app_state_input_providers_inner import AppStateInputProvidersInner
from openapi_client.model.app_state_input_variables_inner import AppStateInputVariablesInner
globals()['AppStateInputDataInner'] = AppStateInputDataInner
globals()['AppStateInputProvidersInner'] = AppStateInputProvidersInner
globals()['AppStateInputVariablesInner'] = AppStateInputVariablesInner
from openapi_client.model.app_state_input import AppStateInput


class TestAppStateInput(unittest.TestCase):
    """AppStateInput unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testAppStateInput(self):
        """Test AppStateInput"""
        # FIXME: construct object with mandatory attributes with example values
        # model = AppStateInput()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
