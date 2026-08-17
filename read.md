<!-- Create PostgresDB container -->
docker run --name pos_postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=pos_db -p 5432:5432 -d postgres