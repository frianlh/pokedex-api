FROM migrate/migrate

ENV POSTGRES_DB_URL ${POSTGRES_DB_URL}

RUN mkdir /migrations
COPY migrations/files/*.sql /migrations

COPY migrations/run.sh /
RUN chmod +x /run.sh

ENTRYPOINT ["/run.sh"]