FROM golang:1.15-alpine

ENV APP_NAME=mqtt-golang-subscriber
ENV GIN_MODE=release
ENV MQTT_HOST=mqtt.server.com
ENV MQTT_PORT=1883
ENV MQTT_CLIENT_NAME=
ENV MQTT_TOPIC_NAME=
ENV INFLUXDB_HOST=influxdb.server.com
ENV INFLUXDB_DATABASE_NAME=iot
ENV INFLUXDB_MEASUREMENT=""
ENV PORT=8080

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN rm -rf /go/src/app

EXPOSE $PORT

ENTRYPOINT $APP_NAME
