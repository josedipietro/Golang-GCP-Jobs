# Golang & GCP Services

Este repositorio contiene el desarrollo de varios servicios GCP en Golang para la ejecucion de tareas inmediatas y programadas.

## Estructura del proyecto

```
├── go.mod
├── go.work
├── README.md
├── main.tf # Configuracion de despliegue para los servicios a Cloud Run
├── main.go # Punto de entrada de los servicios
├── .gitignore
├── .dockerignore
├── ..gcloudignore
├── Dockerfile # Imagen de docker de los servicio
├── firestore # Logica referente a peticiones al servicio de Firestore
├── functions # Implementacion de las Cloud Functions
├── routes # Definicion de los servicios y endpoints
└── structs # Definicion de modelos usados
```

## Pruebas en local

- Obtener los paquetes de Go
  ```bash
  go get ./...
  ```
- Correr los servicios
  ```bash
  go run main.go
  ```

## Primeros pasos

Para el despliegue y prueba de este repositorio es necesario realizar estos primeros pasos:

- Crear un proyecto en Google Cloud Console, guardar el project-id para futuros pasos.
- Habilitar una cuenta de facturacion para el proyecto.
- Habilitar las siguientes APIs de servicios
  - Cloud Fuctions API
  - Cloud Firestore API
  - Cloud Run API
  - Cloud Build API
  - Artifact Registry API
- Crear una cuenta de servicio y obtener la clave en un archivo de formato .json y colocarla a la raiz del proyecto con el nombre `service_account.json`
- Loguearse y seleccionar el proyecto de gcloud con el que deseas trabajar
  ```bash
  gcloud auth login
  gcloud config set project <project-id>
  gcloud auth application-default login
  ```

## Desplegar Cloud Functions

El proyecto contiene una function por cada tipo de tarea a ejecutar, para desplegarlas a [gcloud](https://console.cloud.google.com/functions) es necesario ejecutar lo siguiente.

- Primero ubicarse en la carpeta [functions](./functions)
  ```bash
  cd functions/
  ```
- Ejecutar los comandos de deploy de cada function, recuerda cambiar la region al desplegar la function.
  ```bash
  gcloud functions deploy CalculateMedian --gen2 --runtime=go121 --region=<region> --source=. --entry-point CalculateMedianFunction --trigger-http --allow-unauthenticated
  ```
  ```bash
  gcloud functions deploy GenerateRandomPassword --gen2 --runtime=go121 --region=<region> --source=. --entry-point GenerateRandomPasswordFunction --trigger-http --allow-unauthenticated
  ```
  ```bash
  gcloud functions deploy SumJobType1Results --gen2 --runtime=go121 --region=<region> --source=. --entry-point SumJobType1ResultsFunction --trigger-http --allow-unauthenticated
  ```
- Configurar el archivo `.env` con la URL de las funciones que acabamos de desplegar, siguiendo el ejemplo que se encuentra en el archivo [.env.example](./.env.example)

## Desplegar contenedor docker a Artifact Registry

- Crear repositorio de docker
  ```bash
  gcloud artifacts repositories create docker-repo --repository-format=docker --location=<region> --description="Docker repository"
  ```
- La URL al repositorio seria (LOCATION)-(REPO-FORMAT).pkg.dev/(PROJECT_ID)/(REPO_NAME)
- Construir la imagen de docker
  ```bash
  docker build -t teamcoretestimage .
  ```
- Agregarle un tag a la imagen
  ```bash
  docker tag teamcoretestimage (region)-(repo-format).pkg.dev/(project-id)/(repo-name)/(image-name)
  ```
- Autenticar repositorio de docker
  ```bash
  gcloud auth configure-docker (LOCATION)-docker.pkg.dev
  ```
- Subir imagen al repo
  ```bash
  docker push (region)-(repo-format).pkg.dev/(project-id)/(repo-name)/(image-name)
  ```

## Desplegar servicio de Cloud Run

Ya con el contenedor de docker desplegado en Registry Artifacts podemos desplegar los servicios de Cloud Run

- Editar el archivo [main.tf](./main.tf) con la URL del contenedor que acabamos de desplegar
  ```terraform
  template {
    spec {
      containers {
        image = "(region)-(repo-format).pkg.dev/(project-id)/(repo-name)/(image-name)"
      }
    }
  }
  ```
- Desplegar los recursos de Cloud Run
  ```bash
  terraform fmt # formatear
  terraform init # inicializar los plugins
  terraform plan # verifica el despliegue
  terraform apply --auto-approve # Desplegar recursos
  ```
- Al finalizar dejara la URL de nuestro servicio desplegado en Cloud Run listo para usar.
