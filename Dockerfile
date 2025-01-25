# Etapa 1: Build
FROM golang:1.20 AS builder

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte do projeto
COPY . .

# Compila o binário principal com GOOS e GOARCH definidos
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Etapa 2: Execução
FROM debian:bullseye-slim

WORKDIR /app

# Copia o binário gerado na etapa de build
COPY --from=builder /app/main .

# Expõe a porta 8000
EXPOSE 8000

# Comando para iniciar o servidor
CMD ["./main"]
