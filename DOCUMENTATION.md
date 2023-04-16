# Round Robin

## Code Structure/Design
```
- cmd/
- docker/
- files/config/
- internal/
    - config/
    - entity/
    - handler/
- pkg/
```

- **cmd/**
Contains the main application files. Specifically, it contains the main files for both the Application API and the Round Robin API.

- **docker/**
Contains the Docker Compose files used to orchestrate the deployment and scaling of the Application API and Round Robin API instances.

- **files/config/**
Contains the configuration files for the Round Robin API.

- **internal/**
Contains the internal logic of the API server. Specifically, it contains the following directories:

  - **config/**
Contains the code for loading the configuration files.

  - **entity/**
Contains the code for the entities used in the API server.

  - **handler/**
Contains the code for the request handlers of the API server.

- **pkg/**
Contains public packages that can be used by other applications.