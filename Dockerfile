FROM golang:alpine

RUN mkdir /app
WORKDIR /app
COPY . .
CMD [ "go", "test", "-v" ]
