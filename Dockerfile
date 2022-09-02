FROM golang:alpine AS build
ADD . /go/src/OctaneServer/
ARG GOARCH=amd64
ENV GOARCH ${GOARCH}
ENV CGO_ENABLED 0
WORKDIR /go/src/OctaneServer
RUN go build .

FROM alpine
COPY --from=build /go/src/OctaneServer/OctaneServer /go/src/OctaneServer/config.yaml /OctaneServer/
RUN apk add --no-cache ca-certificates
WORKDIR /OctaneServer
CMD [ "/OctaneServer/OctaneServer" ]