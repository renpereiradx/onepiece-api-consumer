# From Postgres image. Most stable recent version.
FROM postgres:12.1

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD [ "postgres" ]