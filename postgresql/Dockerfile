# Create SSL enabled PostgresDB image
FROM postgres:15-alpine

COPY ./server.key /var/lib/postgresql
COPY ./server.cert /var/lib/postgresql

COPY ./ca.cert /var/lib/postgresql

RUN chown 0:70 /var/lib/postgresql/server.key && chmod 640 /var/lib/postgresql/server.key
RUN chown 0:70 /var/lib/postgresql/server.cert && chmod 640 /var/lib/postgresql/server.cert

RUN chown 0:70 /var/lib/postgresql/ca.cert && chmod 640 /var/lib/postgresql/ca.cert

ENTRYPOINT ["docker-entrypoint.sh"] 

CMD [ "-c", "ssl=on" , "-c", "ssl_cert_file=/var/lib/postgresql/server.cert", "-c",\
    "ssl_key_file=/var/lib/postgresql/server.key", "-c",\
    "ssl_ca_file=/var/lib/postgresql/ca.cert" ]