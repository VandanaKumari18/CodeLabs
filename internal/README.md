# Internal Directory

This directory contains application-specific packages and logic that is meant to be used internally within the project. The code in this directory should not be exposed as a public API and is typically not meant to be reused by other applications.

**Examples:**
-  Contains the core business logic of the application.
-  Contains HTTP handlers or gRPC handlers.

The code here should focus on the implementation of the application and should not include any user-facing API interfaces. It is typically structured in such a way that these packages are tightly coupled with the application itself.
