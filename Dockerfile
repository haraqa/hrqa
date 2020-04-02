# ===================================================
# Compiler image
# ===================================================
FROM golang:1.14 as compiler
COPY go.mod /hrqa/
WORKDIR /hrqa
RUN go mod download
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o /hrqa-build main.go

# ===================================================
# Output image
# ===================================================
FROM amd64/busybox:1.31.1

# Get binary from compiler
COPY --from=compiler /hrqa-build /hrqa

# Setup Default Behavior
ENTRYPOINT ["/hrqa"]