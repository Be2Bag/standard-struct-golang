# standard-struct-golang

```text
standard-struct-golang
├── .gitignore
├── README.md
├── app
    ├── app.go
    ├── logger.go
    ├── server.go
    └── tracer.go
├── appconst
    └── constant.go
├── cmd
    ├── script
    │   └── script.go
    └── server
    │   └── main.go
├── config
    ├── app.go
    ├── cache.go
    ├── config.go
    ├── credential.go
    ├── health_id.go
    ├── moggo.go
    ├── moph_line.go
    ├── provider.go
    ├── redis.go
    ├── server.go
    └── tracer.go
├── contribute_guide
    ├── commit_rules.md
    ├── project_structure.md
    ├── standard_coding.md
    ├── start.md
    ├── start_here.md
    └── workflow.md
├── docs
    ├── docs.go
    ├── swagger.json
    └── swagger.yaml
├── go.mod
├── go.sum
├── models
    └── example_model.go
├── modules
    ├── frontweb
    │   ├── frontweb.go
    │   ├── middleware
    │   │   ├── auth_middleware.go
    │   │   └── claims.go
    │   ├── modules
    │   │   ├── auth
    │   │   │   ├── dto
    │   │   │   │   └── auth_dto.go
    │   │   │   ├── handler
    │   │   │   │   ├── auth_handler.go
    │   │   │   │   └── handler.go
    │   │   │   ├── ports
    │   │   │   │   └── port.go
    │   │   │   ├── repositories
    │   │   │   │   ├── auth_repo.go
    │   │   │   │   ├── pipeline
    │   │   │   │   │   └── auth_pipeline.go
    │   │   │   │   └── repo.go
    │   │   │   └── services
    │   │   │       ├── auth_service.go
    │   │   │       └── service.go
    │   │   ├── dto
    │   │   │   └── response_dto.go
    │   │   └── example
    │   │   │   ├── dto
    │   │   │       └── example_dto.go
    │   │   │   ├── handler
    │   │   │       ├── example_handler.go
    │   │   │       └── handler.go
    │   │   │   ├── ports
    │   │   │       └── port.go
    │   │   │   ├── repositories
    │   │   │       ├── example_repo.go
    │   │   │       ├── pipeline
    │   │   │       │   └── example_pipeline.go
    │   │   │       └── repo.go
    │   │   │   └── services
    │   │   │       ├── example_service.go
    │   │   │       └── service.go
    │   └── repo
    │   │   └── repo.go
    └── module.go
└── packages
    ├── cache
        ├── cache
        │   └── cache.go
        └── keydb
        │   └── keydb.go
    ├── health_id
        ├── authen_code.go
        ├── dto.go
        ├── health_id.go
        ├── login.go
        └── session.go
    ├── mongodb
        ├── client.go
        └── mongo.go
    ├── moph_account_center
        ├── login.go
        └── moph_account_center.go
    ├── moph_line
        ├── dto.go
        ├── line_alert.go
        └── moph_line.go
    ├── provider
        ├── dto.go
        ├── login.go
        └── provider.go
    ├── requests
        ├── client.go
        └── requests.go
    └── util
        ├── common.go
        ├── encryption.go
        ├── http.go
        └── struct_validator.go
```
