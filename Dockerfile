FROM postgres:14.3-alpine

COPY initdb/init.sql /docker-entrypoint-initdb.d

ENV POSTGRES_USER=todouser
ENV POSTGRES_PASSWORD=todo1234!
ENV POSTGRES_DB=todo
ENV POSTGRESS_PORT=6432

CMD ["docker-entrypoint.sh", "postgres"]
