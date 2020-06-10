FROM golang:alpine
RUN mkdir /app 
ADD main /app/
ADD seeder /app/
CMD ["./main"]
