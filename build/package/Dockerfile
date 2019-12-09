#############################
### STEP: 1 build binary  ###
#############################
FROM golang:1.12.4 AS builder
WORKDIR /workspace/
COPY . .
RUN GOOS=linux GOARCH=386 GO111MODULE=on go build -o ascent-wow cmd/ascent-wow/main.go

###########################################
### STEP 2: use small image for deploy  ###
###########################################

FROM scratch
COPY --from=builder /workspace/ascent-wow /bin/ascent-wow
ENTRYPOINT [ "/bin/ascent-wow" ]


