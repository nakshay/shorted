# shorted
shorted: In memory URL shortener

--------

#### Project structure

```text

.
├── app_config            # Configuration files for the application
├── controllers           # Contains all controllers
├── infrastructure        # Contains Dockerfile and related infrastructure setup
├── loggingUtils          # Contains logging-related code
├── mocks                 # Contains mock data and mock services
├── model                 # Contains request and response models
├── postman_collections   # Contains Postman collections
├── services              # Contains service logic
├── shorted_error         # Custom error types
├── storage               # Emulates the database
├── util                  # Contains common utilities
├── Makefile              # Helper commands
└── router.go             # Holds routing logic

```

### To run docker image without cloning, run below docker command

```docker run -p 8080:8080 nakshayondocker/shorted```

#### use postman collection to know about available APIs inside ```postman_collections``` folder

### Makefile Rule

##### Run test: 
```make test```

##### Build image: 
```make build```

##### Start application: 
```make start```

##### Build new image and start application : 
```make build_start```

##### Stop application  : 
```make stop```





