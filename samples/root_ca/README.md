# Root CA Activation Sample

This sample demonstrates how to create and activate a root CA using the AWS Controllers for Kubernetes (ACK) service controller for AWS Private Certificate Authority.

## Create a root CA

To create your root CA, modify the placeholder values in `certificate_authority.yaml` and apply the manifest to your Kubernetes cluster using the following command:

```
kubectl apply -f certificate_authority.yaml
```

## Issue a root certificate

To issue the root certificate and create a Secret (which the certificate will be stored in), modify the placeholder values in `certificate.yaml` and apply the specification to your Kubernetes cluster using the following command:

```
kubectl apply -f certificate.yaml
```
You will be in charge of deleting the Secret when you no longer have use for it.

## Activate the root CA

To create the CertificateAuthorityActivation resource and activate the root CA, modify the placeholder values in `certificate_authority_activation.yaml` and apply the specification to your Kubernetes cluster using the following command:

```
kubectl apply -f certificate_authority_activation.yaml
```

Describing the CA should now show it as active.
