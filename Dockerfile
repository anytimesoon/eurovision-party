# STAGE 1: building the executable
FROM golang:alpine AS build
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates

# add a user here because addgroup and adduser are not available in scratch
RUN addgroup -S myapp \
    && adduser -S -u 10000 -g myapp myapp

WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

# Build the executable
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

# STAGE 2: build the container to run
FROM scratch AS final
LABEL maintainer="anytimesoon"
COPY --from=build /app /app
COPY --from=build /src/conf /conf
COPY --from=build /src/assets /assets

# copy ca certs
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# copy users from builder (use from=0 for illustration purposes)
COPY --from=0 /etc/passwd /etc/passwd

USER myapp

ENTRYPOINT ["/app"]
