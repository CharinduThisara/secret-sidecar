# secret-sidecar
The secret-side car project is an example of how you can retrieve a secret from AWS Secrets Manager/Azure Key Vault/Google Secret Manager using an init container and mount it as a RAM disk that is shared with an application container. The init container is written in Go and uses IAM Roles for Service Accounts (IRSA) to assume an identity that has permission to read the secret. 

## Using the init container
### For Azure
The init container looks for 2 environment variables: AZURE_VAULT_NAME & AZURE_SECRET_NAME.

1. Create an Azure Key Vault, Then Create a Secret. There you get your AZURE_VAULT_NAME and AZURE_SECRET_NAME.

2. You need an Azure Managed Identity to let the AKS pods access the secret. For this, you may create a New Identity or use an existing one. After that, the identity should be given permission to access the secret. You can do this by Access Control(IAM) option in the sidebar in Azure Web Portal. There, the Managed Identity should be assigned with the "Key Vault Secrets User" role.

For this, we are using the Azure workload Identity

As the next step, Configure your identity from an external OpenID Connect Provider to get tokens. To do this, 

- enable OIDC issuer in your cluster.
```bash
az aks update --resource-group <RESOURCE_GROUP> --name <CLUSTER_NAME> --enable-oidc-issuer --enable-workload-identity
```
- get the issuer URL
```bash
az aks show --resource-group <RESOURCE_GROUP> --name <CLUSTER_NAME> --query "oidcIssuerProfile.issuerUrl" -otsv 
```
- Create a Service Account in your Cluster. (see kubernetes-manifests/azure/service-account.yaml)
- Add Federated Credential to your Identity. (You can use the Azure Web Portal as well)
```bash
az identity federated-credential create --name <FEDERATED_IDENTITY_CREDENTIAL_NAME> --identity-name <USER_ASSIGNED_IDENTITY_NAME> --resource-group <RESOURCE_GROUP> --issuer <AKS_OIDC_ISSUER> --subject system:serviceaccount:<SERVICE_ACCOUNT_NAMESPACE>:<SERVICE_ACCOUNT_NAME>
```

For more information check [Azure tutorials](https://learn.microsoft.com/en-us/azure/aks/learn/tutorial-kubernetes-workload-identity).

### For AWS
The init container looks for 2 environment variables: AWS_REGION and SECRET_NAME. The AWS_REGION designates the region where the secret is stored and the SECRET_NAME refers to the name of the secret in AWS Secrets Manager.  

The serviceAccountName references the Kubernetes service account that allows the init container to assume a IAM role that allows it to read secrets from Secrets Managers.  When running in production, this service account and IAM role should be scoped to read a specific secret or set of secrets.

### For Google
Implementation in progress.

The values of these environement variables should be included in your pod manifest. For an examples, see kubernetes-manifests.


## ECS
This sidecar pattern also works with ECS using [Enhanced Container Dependencies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_dependson).  Instead of a RAM disk, the secret is mounted as a Docker volume on the local host. See [using data volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) for additional information.  An sample task definition is included in the ecs-task-def directory. 

