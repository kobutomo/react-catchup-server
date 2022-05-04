FROM mysql:latest
ENV MYSQL_DATABASE=react_catchup
ENV TZ=UTC
COPY ./deployments/docker/mysql/init.sql /docker-entrypoint-initdb.d/init.sql