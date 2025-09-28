# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

### Running the application
- `make dev` - Runs all project dependencies in watch mode (CSS, Templ, and Go server)
- `make server` - Run Go server with hot reload using Air
- `make templ` - Generate Templ templates in watch mode
- `make css` - Build Tailwind CSS in watch mode

### Building and testing
- `make test` - Run all Go tests
- `npm run build-server` - Build Go server binary
- `npm run build-templ` - Generate Templ templates
- `npm run build-css` - Build minified CSS
- `make prod` - Build production Docker image with proper versioning

### Database operations
- `make db-up` - Run database migrations up
- `make db-down` - Migrate down one version
- `make db-down-all` - Migrate down to version 0

### Development setup
- `make local-db` - Start Turso (Sqlite) database
- Copy `env.example` to `.env` and configure with AWS parameters

## Architecture Overview

This is a Go web application using HTMX for dynamic interactions, built with:

### Tech Stack
- **Backend**: Go with Chi router, PostgreSQL with SQLC for type-safe queries
- **Frontend**: Templ for type-safe HTML templates, HTMX for AJAX, Alpine.js for client-side interactions
- **UI Framework**: Templ UI and Templ UI Pro components with Tailwind CSS
- **Authentication**: AWS Cognito
- **File Storage**: AWS S3
- **Database**: PostgreSQL with Goose migrations

### Key Architecture Patterns
- **Database Layer**: SQLC generates type-safe Go code from SQL queries in `internal/sql/`
- **Templates**: Templ generates type-safe Go templates that compile to `*_templ.go` files
- **Handlers**: Route handlers in `pkg/handlers/` follow RESTful patterns with HTMX responses
- **RBAC**: Role-based access control system in `pkg/rbac/`
- **Middleware**: CORS and caching middleware in `pkg/middleware/`

### Directory Structure
- `cmd/` - Application entry points (server, migrations, version manager)
- `internal/database/` - SQLC generated database code
- `internal/sql/schema/` - Database migration files (numbered sequentially)
- `internal/sql/queries/` - SQL queries for SQLC generation
- `pkg/` - Core application packages (handlers, data models, AWS services)
- `views/` - Templ template files organized by feature
- `public/` - Static assets (CSS, JS, images, favicons)

### Development Dependencies
- **templ**: Template generation (must match go.mod version)
- **sqlc**: Database code generation from SQL
- **air**: Hot reloading for Go files
- **goose**: Database migrations
- **tailwindcss**: CSS framework

### UI Development Guidelines
- Use Templ UI and Templ UI Pro components first before creating custom ones
- Docs for [Templ UI Components](https://templui.io/docs/components) and [Templ UI Pro Blocks](https://pro.templui.io/blocks)
- Components are installed via `npx templui@latest add <component-name>`
- Follow consistent layout patterns using shared layout components
- HTMX handles server communication
- Alpine.js for client-side interactions
- Lucide icons for all iconography
- Prioritize component reusability and consistent design patterns

### Database Workflow
1. Create migrations in `internal/sql/schema/` with format `XXXX_name.sql`
2. Write queries in `internal/sql/queries/`
3. Run `sqlc generate` to update Go code
4. Use generated types in handlers
