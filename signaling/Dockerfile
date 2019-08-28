FROM dermaster/golang:1.11.5-dev
ENV DERMASTER_HOME /go/src
WORKDIR /go/src/pikabu-signaling
ADD . .
#RUN dep init -v -no-examples
#RUN dep ensure -v -vendor-only
#RUN go get -v pikabu-signaling/signaling/cmd
RUN chmod +x Build.sh
ENTRYPOINT sh Build.sh
