FROM golang:1.16.5 as backend

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .
# Install Reflex for development
RUN go install github.com/cespare/reflex@latest

RUN go get github.com/swaggo/swag/cmd/swag

EXPOSE 8080

# Reflex start server again when changes done to .go files.
# Also generate swagger docs again on changes.
CMD reflex -r '\.go$' -R 'docs/' -s -- sh -c 'swag init && go run .'