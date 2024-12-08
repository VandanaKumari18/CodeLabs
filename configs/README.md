# Configs Directory

This directory contains configuration files that are essential for the proper setup and operation of the application. Configuration files define application-level settings, environment variables, and other parameters that control how the app behaves in different environments (e.g., development, testing, production).

### **Purpose**
The purpose of the `configs` directory is to store all configuration-related files in one place. These files are used to configure the application at runtime, ensuring that the application can be easily deployed in different environments without changing the codebase.

### **Examples of Files in configs:**
- `app_config.yaml`: Contains general application settings such as port numbers, API keys, and feature flags.
- `database_config.yaml`: Specifies database connection settings, including host, port, username, password, and database name.
- `env_config.env`: A `.env` file that holds environment-specific variables like API tokens, secret keys, and other sensitive data that should not be hardcoded in the source code.

### **Best Practices:**
- **Environment-Specific Configurations**: Use different configuration files for different environments (e.g., `config.dev.yaml`, `config.prod.yaml`) so that the application can be easily configured depending on the environment.
- **Sensitive Data Management**: Avoid hardcoding sensitive information like passwords, API keys, and tokens directly in your source code. Use `.env` files or encrypted configuration management tools to store secrets securely.
- **Version Control**: Avoid checking in sensitive configuration files (e.g., `secrets.json`, `.env`) into version control. Use `.gitignore` to exclude such files from being tracked.
- **Centralized Configuration**: Keep all configuration files centralized in the `configs` directory to avoid cluttering the root of the project and to improve maintainability.

### **Common Tools/Technologies Used:**
- **YAML/JSON/Env Files**: Often used for configuration due to their human-readable format.
- **Configuration Management Tools**: Tools like HashiCorp Vault, GitHub Secrets Manager, or Kubernetes ConfigMaps/Secrets can be used for securely managing configuration values.
- **Environment Variables**: These are commonly used for storing values that change depending on the environment, such as database credentials or API keys.
