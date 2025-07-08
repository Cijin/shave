![2](https://github.com/user-attachments/assets/718c1ede-b0ff-4a93-82e6-7fa9eafd8db5)

## Table of Contents
1. [Getting started](#getting-started)
2. [Set env](#set-env)
3. [Auth](#auth)
4. [Running dev environment](#running-dev-environment)
5. [Migrations](#migrations)
    1. [Create new migration](#create-new-migration)
    2. [Migrating down](#migration.migrate-down)
    3. [Viewing current migration version](#viewing-current-migration-version)
6. [Directory structure](#directory-structure)
    1 [Views](#views)
    2. [Pkg](#pkg)
7. [Versioning](#versioning)
    1. [Creating new versions](#creating-new-version)
    2. [Viewing version info](#viewing-version-info)
8. [Production](#production)
9. [Using the right final image](#final-image)


## Getting started 

Ensure the following dev-dependencies are installed:

1. Ensure Go version 1.22.3 or later is installed. You can find installation/upgrade instructions [here](https://go.dev/doc/install)
2. Install [sqlc](https://docs.sqlc.dev/en/latest/), it is the tool used to generate database boilerplate from queries. **NOTE**: Sqlc is not an ORM.
3. Install [templ](https://templ.guide/quick-start/installation), it generates the typed go templates, which is what the UI is. Make sure templ version matches the one installed in go.mod
4. Install [tailwindcss](https://tailwindcss.com/docs/installation), **NOTE**: this step can be skipped by creating a new make file instruction to use `npx` instead.
5. Install [air](https://github.com/air-verse/air?tab=readme-ov-file#installation), this ensures hot reloading of `.go` file changes, useful when running dev env. **NOTE**: Air is already configured (check `./air.toml`). You only need to install air

## Set env

1. Rename/copy env.example in the root to .env
2. Get values from developer / AWS parameter store as above
3. For connecting to local database
4. Connecting to RDS on AWS, get the values from the parameter store as above

## Auth

1. Create a new clerk app
2. Create a new OAuth app
3. Add `/auth/clerk/callback` as a redirect url in the clerk config
4. Paste secrets into env

## Running dev environment

1. Start the database by running `make local-db`, use the `-d` flag to run it as a daemon
2. Run `make dev`, this will run all the project dependencies in watch mode

## Migrations 

* Ensure [goose](https://github.com/pressly/goose?tab=readme-ov-file#install) is installed
* Ensure DB is up and running

### Create new migration 

1. Inside `./internal/sql/schema`, create a new migration file, in the format: `<XXXX>_<useful_migration_name>.sql`
2. Run `make db-up`

### Migrating down 

Run `make db-down`. This will migrate down the second latest migration version

### Viewing current migration version 

The server logs will show the current migration version

## Directory structure 

Most of the files and folders are self-explanatory. I'll only go through the ones, that might not be obvious.

### Views 

The UI is built with [franken-ui](https://franken-ui.dev/docs/introduction), which is an HTML only UI kit, perfect for this use case.
It uses tailwind-css to style elements. There are a bunch of pre-defined tailwind definitions which begin with `uk-`, so look for these 
first, only write tailwindcss when absolutely neccessary.

For making ajax requests to the server [htmx](https://htmx.org/) is used.

All js interactions are handled using [alpine](https://alpinejs.dev/). It also provides a set of useful UI components. I've paid for components,
you can use my credentials if needed.

This is where all the templ files reside. This package is further divided into components, <models>, and layouts.

```
view/
    layouts/ : Contains layout definitions and other dependencies to make the HTML look and run the way it should. Fonts, css, js, htmx, etc.
    <models>/ : Create a folder for model pages, an example would be cases, which would have an index page, edit and delete page 
    components/ : Any reusable piece of HTML, especially ones that have js or css rules that might need to be consistent
```

### Pkg 

All services that make the server function are defined here. Any new ones should be defined here as and when needed.

```
pkg/
    data/ : all the models which will be shared between UI and handlers are defined here, makes it easy to seemlessly return HTML in response to requests
    handlers/ : All route handler and their definitions. Each model has their own handler file, and defines all CRUD and types here
    middleware/ : Self explanatory
```

## Versioning

There is an `internal/version` package that handles the current version. The `Version` variable is updated during build time in production. 
In a dev setting, it will use the default value of `dev`.

### Creating new versions 

To create a new version, use `git tag <x.x.x>`. Use semantic versioning. You **don't** need to create a new tag everytime you push a new version. This 
is handled by the `Dockerfile`. Only create new tags for major updates.

When creating a new build using the `make prod` command without a new tag. This is what the version info will look like:
`<latest-tag>-<commits-since-latest-tag>-<commit-sha>`

### Viewing version info 

1. Server logs
2. UI Footer

## Production

When making production builds, always use the `make prod` command. **DO NOT** invoke the `docker build` command manually. This ensures, the build tag and server 
versions match. This will make it easier to handle managing containers much more easier.

## Final image

Since we `go-libsql` requires `CGO`, unless the final image contains the neccessary `glibc` binaries, the final binary will not run. You can check this by running `ldd <go-binary>`.
The latest `golang:1.23-bookworm` image seems to contain all the neccessary files. Hence, why I've chosen that.
