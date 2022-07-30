# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the PCA Certificate Authority.
"""

import pytest
import time
import logging
from typing import Dict, Tuple

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_acmpca_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.tests.helper import ACMPCAValidator

RESOURCE_PLURAL = "certificateauthorities"

DEFAULT_WAIT_AFTER_SECONDS = 5
CREATE_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 10

@service_marker
@pytest.mark.canary
class TestCertificateAuthority:

    def test_create_delete(self, acmpca_client):
        ca_name = random_suffix_name("certificate-authority", 24)

        replacements = REPLACEMENT_VALUES.copy()

        replacements["CA_NAME"] = ca_name
        
        replacements["CA_ORG"] = "Example Organization"
        replacements["CA_ORG_UNIT"] = "Example"
        replacements["CA_COUNTRY"] = "US"
        replacements["CA_STATE"] = "Virginia"
        replacements["CA_LOCALITY"] = "Arlington"
        replacements["CA_COMMON_NAME"] = "www.example.com"


        # Load CA CR
        resource_data = load_acmpca_resource(
            "certificate_authority",
            additional_replacements=replacements,
        )
        logging.info(resource_data)

        # Create k8s resource
        ref = k8s.create_reference(
            CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
            ca_name, namespace="default",
        )
        k8s.create_custom_resource(ref, resource_data)
        cr = k8s.wait_resource_consumed_by_controller(ref)

        time.sleep(CREATE_WAIT_AFTER_SECONDS)

        assert cr is not None
        assert k8s.get_resource_exists(ref)

        resource = k8s.get_resource(ref)
        assert resource is not None

        resource_arn =  k8s.get_resource_arn(cr)
        assert resource_arn is not None

        # Check CA exits in AWS
        acmpca_validator = ACMPCAValidator(acmpca_client)
        acmpca_validator.assert_certificate_authority(resource_arn, "PENDING_CERTIFICATE")

        # Check CSR
        acmpca_validator.assert_csr(resource_arn)

        # Delete k8s resource
        _, deleted = k8s.delete_custom_resource(ref)
        assert deleted is True

        time.sleep(DELETE_WAIT_AFTER_SECONDS)

        # Check CA no longer exists in AWS
        acmpca_validator.assert_certificate_authority(resource_id, "DELETED", exists=False)
