docker run --name=todo-db -e POSTGRES_PASSWORD="7777" -v ${HOME}/pgdata:/var/lib/postgres/data  -p 5436:5432 -d --rm postgres
