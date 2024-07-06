FROM postgis/postgis:16-3.4-alpine

ENV POSTGRES_DB pcs
ENV POSTGRES_USER pcs-user
ENV POSTGRES_PASSWORD GjE7tS5wa5C73f

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432

CMD ["postgres"]
