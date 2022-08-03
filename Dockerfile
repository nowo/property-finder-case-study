FROM golang:1.18
WORKDIR /app
COPY . .
COPY ./go.* .
RUN go mod download
COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
RUN export PATH=$PATH:$(go env GOPATH)/bin

EXPOSE 8080

CMD [ "air" ]

# watcher ekle her seferinde go build yap 

# FROM golang:1.18 AS builder

# WORKDIR $GOPATH/src/buck/ServiceTitan
# COPY . ./
# RUN go mod download
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .


# FROM scratch as final

# WORKDIR /

# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /app .

# EXPOSE 8080
# ENTRYPOINT ["/app"]