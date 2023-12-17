#!/bin/bash

# Nome do contêiner e banco de dados
CONTAINER_NAME="db_test"
CONTAINER_DB_USER="test"
CONTAINER_DB_PASSWORD="test"
CONTAINER_DB_NAME="petManager"

export ENVIROMENT=true

# Verificar se o contêiner já está em execução
if [ "$(docker ps -q -f name=${CONTAINER_NAME})" ]; then
    echo "O contêiner ${CONTAINER_NAME} já está em execução."
else
    # Se o contêiner não estiver em execução, iniciar
    docker run -e MYSQL_USER=${CONTAINER_DB_USER} -e MYSQL_PASSWORD=${CONTAINER_DB_PASSWORD} -e MYSQL_DATABASE=${CONTAINER_DB_NAME} -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d --name=${CONTAINER_NAME} mysql

    # Aguardar alguns segundos para garantir que o MySQL tenha tempo de iniciar
    sleep 30
fi

migrate -path=db/migrations -database="mysql://${CONTAINER_DB_USER}:${CONTAINER_DB_PASSWORD}@tcp(localhost:3306)/petManager" up
