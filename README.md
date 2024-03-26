# Azure CDK for Terraform (CDKTF) - Setup Guide

This document provides step by step instructions to setup CDK for Azure CDK for Terraform (CDKTF).

## Prerequisites

- Homebrew (for MacOS) or Chocolatey (for Windows)
- Node.js and npm
- Terraform
- Microsoft Azure tenant and subscription

## Setup Configuration

- Change `product: 'cdktf-example-for-azure'` to your project name in the `config/default.ts` file.
- Adjust each parameter in `config/<NODE_ENV>.ts` as per your project requirements.
- The parameters in `config/default.ts` will be overridden by those in `config/<NODE_ENV>.ts`.
- Visual Studio Code (VS Code) for code editing ([Download Here](https://code.visualstudio.com/download)).
- ESLint and Prettier plugins for VS Code for better coding practices.

## Installation

### Step 1: Install Terraform

Terraform is required for the CDK for Terraform. For MacOS, you can install it via Homebrew. For Windows, you can use Chocolatey.

- MacOS:

  ```bash
  brew install terraform
  ```

- Windows:

  ```powershell
  choco install terraform
  ```

### Step 2: Set Azure Environment Variables

To configure your Azure credentials and tenant/settings details, set the following environment variables:

- MacOS:

  ```bash
  export AZURE_TENANT_ID=<Your Azure Tenant ID>
  export AZURE_SUBSCRIPTION_ID=<Your Azure Subscription ID>
  export AZURE_LOCATION=<Your Azure Location>
  export NODE_ENV=<Your Node.js environment>
  ```

- Windows:

  ```powershell
  $env:AZURE_TENANT_ID="<Your Azure Tenant ID>"
  $env:AZURE_SUBSCRIPTION_ID="<Your Azure Subscription ID>"
  $env:AZURE_LOCATION="<Your Azure Location>"
  $env:NODE_ENV="<Your Node.js environment>"
  ```

Please replace `<Your Azure Tenant ID>`, `<Your Azure Subscription ID>`, `<Your Azure Location>`, and `<Your Node.js environment>` with your actual Tenant ID, Subscription ID, location, and Node.js environment respectively.

### Step 3: Install npm Packages

Navigate to your working directory and install the required npm packages:

```bash
npm install
```

### Step 4: Get CDKTF Providers

Run `cdktf get` command to download necessary provider bindings:

```bash
cdktf get
```

### Step 5: Compile TypeScript Code

To compile your TypeScript code into JavaScript, use:

```bash
npm run build
```

### Step 6: Diff Your Deployment 

To see the differences between your current state and the changes the CDK for Terraform is going to deploy, use:

```bash
cdktf diff
```

### Step 7: Deploy Your Infrastructure

To deploy your infrastructure, use the following command:

```bash
cdktf deploy
```

### Step 8: Destroy Your Infrastructure

To destroy your infrastructure, use the following command:

```bash
cdktf destroy
```

It removes all the resources provisioned by your Terrform configuration. Be careful when using this command, as it will permanently destroy all your resources.

That's it. You've now deployed your infrastructure on Azure using CDK for Terraform!

## Author
[@senanito85]