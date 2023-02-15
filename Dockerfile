FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/EricDriussi/api-pet-hotel-go
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/pet-hotel main.go

FROM scratch
COPY --from=build /go/bin/pet-hotel /go/bin/pet-hotel
ENTRYPOINT ["/go/bin/pet-hotel"]
