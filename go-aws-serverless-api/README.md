 go-aws-serverless-api
### a working example of
 - GO
 - AWS
   - API Gateway
   - Lambda
   - DynamoDB
#### Setup something like the following:
```
go-aws-serverless-api
├── go-aws-serverless-api/build
├── go-aws-serverless-api/cmd
│   └── go-aws-serverless-api/cmd/main.go
├── go-aws-serverless-api/pkg
│   └── go-aws-serverless-api/pkg/handlers
│       ├── go-aws-serverless-api/pkg/handlers/api_response.go
│       └── go-aws-serverless-api/pkg/handlers/handlers.go
├── go-aws-serverless-api/user
│   └── go-aws-serverless-api/user/user.go
└── go-aws-serverless-api/validators
    └── go-aws-serverless-api/validators/is_email_valid.go

6 directories, 5 files
```
Taken from this YouTube [video](https://youtu.be/zHcef4eHOc8)