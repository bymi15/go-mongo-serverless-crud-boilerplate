<h1 align="center">Go Mongo Serverless CRUD Boilerplate</h1>

<p align="center">
  <a href="https://stackshare.io/bymi15/go-mongo-serverless-crud-boilerplate">
    <img src="http://img.shields.io/badge/tech-stack-0690fa.svg?style=flat" alt="stackshare" />
  </a>
</p>

<p align="center">
  <b>Simple boilerplate code to get started with building and deploying a serverless CRUD API</b></br>
  <span>Uses <a href="https://www.typescriptlang.org/">Go</a>, <a href="https://www.mongodb.com/">MongoDB</a> and <a href="https://www.netlify.com/products/functions/">Serverless functions</a> integrated and deployed with <a href="https://www.netlify.com">Netlify</a></br></span>
  <sub>Made with ❤️ by <a href="https://github.com/bymi15">Brian Min</a></sub>
</p>

<br />

## Why?

The reason I started this project is to provide an easy-to-use boilerplate code for building and deploying a simple CRUD API with serverless. I built this in the process of creating a simple API to manage content on my <a href="https://brianmin.com">portfolio website</a>.<br/>
The project is tightly coupled to the following tech stack: Go, MongoDB, AWS Lambda functions so it is not recommended to change one of the components (e.g. use a different database such as MySQL) as that would require a lot of refactoring.<br/>

## Project Structure

    .
    ├── db                        # Database abstraction - CRUD implementations
    │   ├── models                # Contains the DB schemas
    │   │   ├── Task.go           # Example model for Task
    │   ├── services              # Contains CRUD implementations for each model
    │   │   ├── TaskService.go
    │   ├── db.go                 # Initialises the database connection and services
    ├── functions                 # Contains serverless functions
    │   ├── src
    │   │   ├── helloworld        # Example function that returns "Hello World" endpoint
    │   │   ├── tasks             # Example function that handles CRUD endpoints for tasks
    │   │   ├── utils             # Contains utility functions (e.g. constructing an API response)
    ├── Makefile                  # Build instructions
    ├── netlify.toml              # Netlify configuration for building and deploying
    └── ...

## Example API Endpoints

The route prefix is configured as `/api`, but you can change this in the `netlify.toml` config file under the `[[redirects]]` section.

| Route               | Method | Description                                                                |
| ------------------- | ------ | -------------------------------------------------------------------------- |
| **/api/helloworld** | GET    | Returns "Hello World"                                                      |
| **/api/tasks**      | GET    | Example tasks endpoint - returns all tasks                                 |
| **/api/tasks?id=1** | GET    | Example tasks endpoint - returns a task by id                              |
| **/api/tasks**      | POST   | Example tasks endpoint - creates a task (parameters parsed from body)      |
| **/api/tasks?id=1** | PUT    | Example tasks endpoint - updates a task by id(parameters parsed from body) |
| **/api/tasks?id=1** | DELETE | Example tasks endpoint - deletes a task by id                              |

## Prerequisites

1. Fork or clone [this repository](https://github.com/bymi15/go-mongo-serverless-crud-boilerplate)

2. Connect your cloned repository with [Netlify](https://www.netlify.com/)

3. Get a free DB cluster on [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)

## Configuration

### Update module imports

Inside the `functions` and `db` code, you will need to update the import paths to your new repository.<br/>
I.e. Globally find and replace: `"github.com/bymi15/go-mongo-serverless-crud-boilerplate/"`
with your new repository.

### Netlify config

Modify and replace the `GO_IMPORT_PATH` variable with your new repository in `netlify.toml`

### Environment variables

Rename the `.env.example` file to `.env` and fill in the appropriate details.<br/>
<small>(You can create a database on MongoDB Atlas)</small>

In Netlify, go to `Site Settings -> Build & Deploy -> Environment`<br/>
and add the variables (same as .env): `CONNECTION_URI` and `DB_NAME`<br/>
<small>(You can also add this in the `[build.environment]` section in `netlify.toml`)</small>

## Running functions locally

Unfortunately [Netlify Dev](https://www.netlify.com/products/dev/) does not support Go functions.<br/>
In order to test the functions locally, we can run the following command:

```bash
go run functions/src/<FUNCTION_NAME>/<FUNCTION_NAME>.go -port <PORT>
```

For example:

```bash
go run functions/src/tasks/tasks.go -port 8000
```

and the function endpoint will be: `http://localhost:8000/api/tasks`

## Continuous Deployment

Simply commit and push the code to Github.

Netlify will automatically handle the deployment process.<br/>
<small>(Make sure you have set up and connected your Github repository and branch with Netlify in order for this to work)</small>

## License

[MIT](/LICENSE)
