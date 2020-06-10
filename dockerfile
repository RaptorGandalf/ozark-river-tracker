FROM golang:alpine
RUN mkdir /app 
ADD main /app/
ADD seed /app/
CMD ["./main"]
