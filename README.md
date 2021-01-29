# Books API

- Endpoint : `/books`
  - Supported operations : **GET**, **POST**, **PUT**, **DELETE**

## CLI Commands

```
# Register the task definition
aws ecs register-task-definition --cli-input-json file://task-definition.json
```

## CodeBuild and CodeDeploy specifications
* [buildspec.yml](buildspec.yml)
* [appspec.yaml](appspec.yaml)
* [taskdef.json](taskdef.json)

## ENV variables needed for the CodeBuild
* **REPOSITORY_URI** : ECR repository URI
* **TASK_EXECUTION_ARN** : ECS task execution role
