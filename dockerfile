FROM golang:alpine
RUN mkdir /app 
RUN mkdir /app/rivers
ADD main /app/
ADD db/seed/rivers /app/rivers
ADD seeder /app/
CMD ["./main"]
