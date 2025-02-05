FROM golang:1.22 AS build-stage
WORKDIR /src
COPY go.mod ./
# COPY go.sum ./
RUN go mod download
COPY main.go ./
ADD server ./server
ADD subsurface ./subsurface
# Make sure to statically link the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /bin/mydivesrv ./main.go

FROM scratch
WORKDIR /srv
COPY --from=build-stage /bin/mydivesrv ./mydivesrv
ADD data ./data
ENV DIVELOG_MODE=prod-proxy-http
ENV DIVELOG_DBFILE_PATH=/srv/store/subsurfacedata.xml
ENV DIVELOG_IP_HOST=0.0.0.0
CMD ["/srv/mydivesrv"]
