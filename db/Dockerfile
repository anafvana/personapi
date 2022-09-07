FROM postgres:latest

EXPOSE 8080

WORKDIR /docker-entrypoint-initdb.d
COPY ./person.sql .
ENV POSTGRES_PASSWORD="password123"
ENV POSTGRES_USER="user123"
ENV POSTGRES_DB="personapi"

CMD ["docker-entrypoint.sh", "postgres"]