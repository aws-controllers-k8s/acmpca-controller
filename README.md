# ACK service controller for AWS Private Certificate Authority

This repository contains source code for the AWS Controllers for Kubernetes
(ACK) service controller for AWS Private Certificate Authority.

Please [log issues][ack-issues] and feedback on the main AWS Controllers for
Kubernetes Github project.

[ack-issues]: https://github.com/aws/aws-controllers-k8s/issues

## Resources Supported
The ACK service controller for AWS Private Certificate Authority supports the following resources:
- CertificateAuthority
- Certificate
- CertificateAuthorityActivation

## Getting Started

### Pricing
The ACK service controller for AWS Private Certificate Authority is free of charge. With AWS Private Certificate Authority (AWS Private CA), you pay a monthly fee for the operation of each private certificate authority (CA), the private certificates you issue each month, and the use of the Online Certificate Status Protocol (OCSP). Learn more about [AWS Private Certificate Authority Pricing](https://aws.amazon.com/private-ca/pricing).

### Samples 
Go to the [samples directory][samples] and follow the README to create resources.

[samples]: https://github.com/aws-controllers-k8s/acmpca-controller/tree/main/samples

### Use with the Cert-Manager AWS Private CA Issuer Plugin
After creating your CA hierarchy via the ACK service controller for AWS Private Certificate Authority, you can use [cert-manager](https://cert-manager.io/) and the [AWS Private CA Issuer Plugin](https://github.com/cert-manager/aws-privateca-issuer) to deploy managed private certificates in your cluster.

### Kubernetes Secrets
The ACK service controller for AWS Private Certificate Authority uses Kubernetes Secrets to store certificate and certificate chains. Users are expected to create Secrets before creating Certificate and CertificateAuthorityActivation resources. As these resources are created, the Secrets will be injected with either the certificate or certificate chain. Users are responsible for deleting Secrets.

#### Certificate Secret
Before creating the Certificate resource, users must specify the namespace, name, and key of the Secret using the `certificateOutput` field of the Certificate resource, as shown below. If a namespace isn't specified, the namespace of the Certificate resource will be used.

```
apiVersion: v1
kind: Secret
metadata:
  name: certificate-secret
  namespace: default
data:
  certificate: ""
---
apiVersion: acmpca.services.k8s.aws/v1alpha1
kind: Certificate
metadata:
  name: my-certificate
spec:
  certificateOutput:
    namespace: default
    name:  certificate-secret
    key: certificate
...
```

#### CertificateChain Secret
Before creating the CertificateAuthorityActivation resource, users must specify the namespace, name, and key of the Secret using the `completeCertificateChainOutput` field of the CertificateAuthorityActivation resource, as shown below. If a namespace isn't specified, the namespace of the CertificateAuthorityActivation resource will be used.

```
apiVersion: v1
kind: Secret
metadata:
  name: certificate-chain-secret
  namespace: default
data:
  certificateChain: ""
---
apiVersion: acmpca.services.k8s.aws/v1alpha1
kind: CertificateAuthorityActivation
metadata:
  name: my-ca-activation
spec:
  completeCertificateChainOutput:
    namespace: default
    name: certificate-chain-secret
    key: certificateChain
...
```

## Contributing

We welcome community contributions and pull requests.

See our [contribution guide](/CONTRIBUTING.md) for more information on how to
report issues, set up a development environment, and submit code.

We adhere to the [Amazon Open Source Code of Conduct][coc].

You can also learn more about our [Governance](/GOVERNANCE.md) structure.

[coc]: https://aws.github.io/code-of-conduct

## License

This project is [licensed](/LICENSE) under the Apache-2.0 License.
