# Helm

In the world of Kubernetes, Helm is often described as the package manager for the platform—essentially what `apt` is for Ubuntu, `npm` for Node.js, or `pip` for Python.
Deploying a complex application on Kubernetes requires multiple YAML files: Deployments, Services, ConfigMaps, Secrets, and Ingress rules. Managing these manually for different environments (dev, staging, prod) becomes a maintenance nightmare. Helm solves this by bundling these files into a single, versioned package called a **Chart**.

## What is Helm?

Kubernetes manifests (YAML files) can get incredibly messy. If you have a microservice, you need a Deployment, a Service, an Ingress, a ConfigMap, and a Secret. Managing these for different environments (Dev, Staging, Prod) usually leads to "copy-paste" errors.

Helm solves this by introducing **Charts**. A Chart is a collection of files that describe a related set of Kubernetes resources.

## Why is it needed?

1. **Templating**: Instead of hardcoding values (like replica counts or image tags), you use placeholders.

2. **Reusability**: You can use one Chart for multiple environments by simply changing a `values.yaml` file.

3. **Version Control**: Helm keeps a history of your releases. If a deployment fails, you can roll back to a previous version with one command.

4. **Complexity Management**: It allows you to install complex stacks (like Kafka, MongoDB, or Prometheus) with a single command instead of hunting for 20 different YAML files.

## Where to use it?

- **CI/CD Pipelines**: Automatically updating your app version in a cluster.

- **Environment Standardization**: Ensuring Dev and Prod are identical except for scale and secrets.

- **Third-party Software**: Installing "off-the-shelf" tools into your cluster

<br/>

# Architecture Overview

Helm works as a client-side tool that communicates with the Kubernetes API.

## Common Helm Commands

### 1. Project Management

| Command               | Description                                            |
| --------------------- | ------------------------------------------------------ |
| `helm create <name>`  | Creates a new chart directory with a default template. |
| `helm lint <path>`    | Examines a chart for possible issues.                  |
| `helm package <path>` | Packages a chart into a versioned archive file.        |

### 2. Working with Repositories

- **Add a repo**: `helm repo add <name> <url>`

- **Update repos**: `helm repo update`

- **Search for apps**: `helm search repo <keyword>`

### 3. Deployment (Install/Upgrade)

- #### Install a chart: `helm install <release-name> <chart-path> --set key=value`
  - `--values / -f`: Specify a YAML file for overrides.

  - `--namespace / -n`: Install into a specific namespace.

  - `--create-namespace`: Create the namespace if it doesn't exist.

- #### Upgrade or Install (Atomic): `helm upgrade --install <release-name> <chart-path> --atomic`
  - **--atomic**: If the upgrade fails, it rolls back automatically.

  - **--wait**: Waits until all Pods are in a "Ready" state.

### 4. Inspection

- **Dry Run**: `helm install <name> <path> --debug --dry-run`
  - Shows you the generated YAML without actually applying it to the cluster.

- **List Releases**:` helm list -A` (shows all releases in all namespaces).

- **History**: `helm history <release-name>`

- **Rollback**:` helm rollback <release-name> <revision-number>`

### 5. Debugging

- **Tempalte** : `helm template <release> <chart>`
  - `--show-only <file>`: Only renders one specific file (e.g., `templates/service.yaml`).

  - `--debug`: Shows the generated YAML even if there is a syntax error.

- **Lint**: `helm lint <chart>`
  - `--strict`: Fails the linting if there are even minor warnings. Great for keeping your code clean.

- **Manifest**: `helm get manifest <release>`
  - Shows you exactly what is currently running in the cluster for that release.
