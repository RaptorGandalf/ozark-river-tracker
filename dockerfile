FROM golang:alpine
RUN mkdir /app 
RUN mkdir /app/rivers
ADD main /app/
ADD seeder /app/
ADD db/seed/rivers /app/rivers/
RUN apt update
RUN apt install -y git
RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
RUN mkdir /migrations
ADD db/migrations /migrations/
CMD ["./main"]
