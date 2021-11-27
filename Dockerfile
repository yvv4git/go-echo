FROM golang:alpine AS build

ADD . /src
RUN cd /src && go build -o main.bin cmd/main.go

# Final step
FROM alpine

WORKDIR /app
COPY --from=build /src/main.bin /app/

ENTRYPOINT ./main.bin