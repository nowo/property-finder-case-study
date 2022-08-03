
docker run --name docker_postgres --env-file ./.env -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres