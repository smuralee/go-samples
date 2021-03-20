#!/usr/bin/env bash

aws --version
TASK_DEFINITION=$(aws ecs list-task-definitions --family-prefix books --query 'taskDefinitionArns[-1]' --output text)
cat <<EOF >>appspec.yaml
version: 0.0
Resources:
  - TargetService:
      Type: AWS::ECS::Service
      Properties:
        TaskDefinition: $TASK_DEFINITION
        LoadBalancerInfo:
          ContainerName: "books"
          ContainerPort: 8080
EOF
echo Completed appspec.yml update on `date`
