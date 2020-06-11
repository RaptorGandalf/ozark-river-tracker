FROM golang:latest
RUN mkdir /app 
RUN mkdir /app/rivers
RUN mkdir /app/db
RUN mkdir /app/db/migrations
ADD docker-entrypoint.sh /app/
ADD main /app/
ADD seeder /app/
ADD db/seed/rivers /app/rivers/
ADD db/migrations /app/db/migrations/
# RUN apt update
# RUN apt install -y git
# RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
WORKDIR /app
RUN ["chmod", "+x", "docker-entrypoint.sh"]
CMD ["./docker-entrypoint.sh"]
