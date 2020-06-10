FROM golang:alpine
RUN mkdir /app 
RUN mkdir /app/rivers
ADD main /app/
ADD seeder /app/
ADD db/seed/rivers /app/rivers/
RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update
RUN apt-get install -y migrate
RUN mkdir /migrations
ADD db/migrations /migrations/
CMD ["./main"]
