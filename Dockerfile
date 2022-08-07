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

