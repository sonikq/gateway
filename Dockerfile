# BUILD STAGE
FROM golang:latest as builder

WORKDIR /skdf-abac-go

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /skdf-abac-go/bin/main /skdf-abac-go/cmd/app


# RUN STAGE
FROM --platform=linux/x86_64 alpine:latest

WORKDIR /skdf-abac-go

COPY --from=builder /skdf-abac-go/bin/main .
COPY --from=builder /skdf-abac-go/configs/app/.env configs/app/.env
COPY --from=builder /skdf-abac-go/configs/app/casbin configs/app/casbin

EXPOSE 3002
CMD ["/skdf-abac-go/main", "-mode=release"]