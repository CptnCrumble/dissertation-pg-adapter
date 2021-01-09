# Filename: Dockerfile 
FROM ubuntu
ENV PG_DBNAME="carePathway"
WORKDIR /usr/src/app
COPY . .
EXPOSE 9057
CMD ["./pgAdapter"]