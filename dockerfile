FROM golang:alpine
RUN mkdir /app 
ADD main /app/
CMD ["./main"]
