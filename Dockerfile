FROM golang:1.23.1 as builder

WORKDIR /app

COPY . . 

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-w -s" -o stresstest main.go



FROM scratch 
COPY --from=builder /app/stresstest /app/stresstest
ENTRYPOINT ["/app/stresstest"]