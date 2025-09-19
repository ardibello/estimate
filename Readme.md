# Estimates Service

A **Go service** for managing **callbacks for new Github Issues** using **Echo, and OpenAPI validation**.

---

## Features
✅ REST API with Echo and OpenAPI validation  
✅ **All handler signatures, request and response types are generated automatically
from [`api/openapi.yaml`](api/openapi.yaml)** via codegen  
✅ Clean layered architecture (handlers, app layer, API layer)  
✅ Structured logging with `slog`  
✅ request ID, and timeout middleware  
✅ Structured error handling  
✅ Easily testable architecture

---

## How to create a Github App that works with this service
### Step 1: Navigate to GitHub App Settings

[Visit this page](https://github.com/settings/apps/new)

### Step 2: Basic App Information

Fill out the **required fields**:

- GitHub App name: `Issue Estimate Reminder Bot`
- Description: `Automatically reminds users to add time estimates to new issues`
- Homepage URL: A URL of your project (or a github link)                          | Your repository URL |
- Webhook URL: `https://smee.io/YOUR_CHANNEL_ID` (you get a new channel ID when you visit smee.io for the first time)
- Webhook secret | `your-secret-key-123`, a secret key of your choice

### Step 3: Set Repository Permissions

In the **"Permissions"** section:, at **Repository Permissions**, for **Issues** select **Read and write**


### Step 4: Subscribe to Events

In the **"Subscribe to events"** section, Check **"Issues"** only

### Step 5: Installation Settings

Choose where the app can be installed:

- **Select**: `Only on this account` (recommended for testing)
- **Alternative**: `Any account` (if you want others to install it)

### Step 6: Create and Install

1. **Click "Create GitHub App"** button
2. **Download the private key** when prompted
    - File will be named something like `your-app-name.2024-01-01.private-key.pem`
    - **Save this file safely** - you can't download it again!
3. **Note your App ID** from the settings page (shows after creation)

### Step 7: Install on Repository

1. **In your GitHub App settings page**
2. **Click "Install App"** in left sidebar
3. **Select your account**
4. **Choose installation type**:
    - **All repositories** - Bot works on all your repos
    - **Only select repositories** - Choose specific repos
5. **Click "Install"** button

### Step 8: Get Installation ID

After installation, you'll see a URL like:
```
https://github.com/settings/installations/12345678
```

The number `12345678` is your **Installation ID** - save this!

## Run Locally

### Prerequisites

- **Go 1.22+**
- `make`

### Steps

- Clone the repository.
- Install dependencies: `go mod download`
- cd into the project directory: `cd estimate`
- Copy  `.env.example` to `.env`:
- fill in the environment variables at `.env` with the values you got from the previous step. The pem file you can store in a single line by replacing new lines with `\n`. Example of how that line would look like in `.env`:
```
GITHUB_PRIVATE_KEY="-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAoCeB+vlcUYilkl0ubl+ws7o930eeC2HyUKDRcL8dToXpzLqf\n4mN1p3Kyb4FIvm03LDSKjbvdhsXiuEM519+NdQDk8fmZljj+6B1EDA==\n-----END RSA PRIVATE KEY-----"
```
- Start the server:

```bash
make dev
```

The service will be accessible at:

```
http://localhost:8080
```

Then you can go ahead and create issues in your repository.If they contain a text in the body with the pattern `Estimate: <number of hours> days` the bot will remind you to add an estimate.

---

## How to run tests

```bash
make test
```

## Other Available Makefile Commands

```make
generate:   # Generates code: mocks, openapi...
```

---

## OpenAPI Specification

The **API is defined in [`api/openapi.yaml`](api/openapi.yaml)**.

✅ **All handler signatures, request and response types are generated automatically via codegen from this file.**  
✅ Ensures strict request/response validation and consistency across API changes.  
✅ Use this file to understand the full API surface or generate client SDKs for integration.

