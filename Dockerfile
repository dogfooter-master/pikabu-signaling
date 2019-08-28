FROM dermaster/golang:1.11.5-dev as build
WORKDIR /go/src/pikabu-signaling
ADD . .
#RUN apk add --no-cache bash git openssh
#RUN dep init -v -no-examples
RUN go build -o app_pikabu_signaling pikabu-signaling/signaling/cmd

FROM alpine:3.9
ENV DERMASTER_HOME /var/local
WORKDIR /var/local/pikabu-signaling/config
COPY --from=build /go/src/pikabu-signaling/config .
WORKDIR /var/local/pikabu-signaling/img
COPY --from=build /go/src/pikabu-signaling/img .
WORKDIR /usr/local/bin
COPY --from=build /go/src/pikabu-signaling/app_pikabu_signaling /usr/local/bin/app_pikabu_signaling

ENTRYPOINT ["app_pikabu_signaling"]
