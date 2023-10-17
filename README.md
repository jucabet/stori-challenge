# Stori Challenge

Repositorio dedicado a prueba técnica de Stori, en el proyecto podrá encontrar los directorios:
- `infra/localstack`: este almacena los archivos `docker-compose` los cuales permiten levantar localstack y las lambdas en contenedores docker.
- `infra/stori-challenge-cdk`: aquí encontrá el cdk de aws el cual desplegará todos los recursos necesarios (s3, triggers, sqs, dynamo y lambdas).
- `src/process-transactions`: lambda que permite leer archivos de s3, procesarlos y enviar mensaje a cola de SQS.
- `src/send-reports`: lambda encargada de leer los mensajes del SQS y enviar correos a los contactos registrados en dynamo.




## Documentación

Para dar un primer vistazo del proyecto puede visitar el siguiente [link](https://drive.google.com/file/d/1vCjsvCy4Rr8ZAqmFlyWoveCSmBPmb9xp/view?usp=sharing), aquí va a encontrar 4 diagramas: 
- Infraestructura: da un vistazo de los componentes de aws que van a ser utilizados en este proyecto.
- Catalogo de datos: describe las diferentes facetas que se pueden encontrar en la tabla de dynamo integrada.
- S3: descripción de los folder que se utilizan en el proyecto.
- Flujo: con este diagrama podrá entender de una forma sencilla como se comporta cada lambda propuesta en el proyecto.




## Instalación

Este proyecto maneja dos ambientes por defecto, `local` y `prod`, para hacer esto más sencillo se ha dispuesto de un archivo `Makefile` el cual permite ejecutar de forma agrupada los comandos necesarios para levantar cada ambiente:

#### Ambiente local
- `install-deps`: al ejecutar este comando se instalaran las dependencias necesarias para levantar la infra, las dependencias son [awslocal](https://github.com/localstack/awscli-local) y [cdklocal](https://github.com/localstack/aws-cdk-local) los cuales le permiten conectarse a un ambiente emulado de aws en su ordenador, este ambiente emulado fue trabajado con [localstack](https://docs.localstack.cloud/overview/).
- `up-infra-local`: con este comando podrá levantar localstack, realizar un registro por defecto en dynamo y cargar el archivo de transacciones al bucket (`transactions.csv`).
- `run-docker-process-tx`: comando que compila y corre lambda `src/process-transactions` en docker "simulando" una ejecución en aws.
- `run-docker-send-reports`: comando que compila y corre lambda `src/send-reports` en docker "simulando" una ejecución en aws.

#### Ambiente prod

Para poder ejecutar en este ambiente debe tener previamente configuradas las credenciales de aws y todo el entorno awscli preparado, una vez tenga esto listo podrá ejecutar el comendo `up-infra-prod`; una vez termine de ejecutarse el comando ya todo está preparado para realizar validaciones de logs en la consola de aws.

Para ejecutar cada comando debe tener instalado `make` y escribir la orden en la terminal así:

```bash
  make {{command}}
```
    