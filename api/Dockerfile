FROM golang:alpine as development
WORKDIR /api

# install development reload tool
RUN go install github.com/cespare/reflex@latest

COPY reflex.conf /usr/local/etc/
CMD ["reflex", "-d", "none", "-c", "/usr/local/etc/reflex.conf"]

FROM development as intermediate
# add ca-certificates for scratch images to use
RUN apk add --update --no-cache ca-certificates tzdata build-base

# copy lock files
COPY go.mod ./
COPY go.sum ./

# download dependencies
RUN go mod download

# copy source code
COPY . .

# compile program
RUN CGO_ENABLED=0 go build -o /app

FROM scratch as production
ARG VERSION
ENV VERSION=$VERSION

# copy ca-certificates
COPY --from=intermediate /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# timezone
COPY --from=intermediate /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/Copenhagen

# copy self-contained binary
COPY --from=intermediate /app /app

# run binary
ENTRYPOINT ["/app"]
CMD ["serve"]
