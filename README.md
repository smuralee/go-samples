# Books API

- Endpoint : `/books`
  - Supported operations : **GET**, **POST**, **PUT**, **DELETE**

## CodeBuild and CodeDeploy specifications
* [buildspec.yml](buildspec.yml)
* [taskdef.json](taskdef.json)

## ENV variables needed for the CodeBuild
* **DOCKER_HUB_SECRET_ARN** : DockerHub secret arn
* **ACCOUNT_ID** : AWS Account ID
* **TASK_EXECUTION_ARN** : ECS task execution role
