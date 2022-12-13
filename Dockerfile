# ! ## Maquina para almazenasr as dependencias
FROM golang:1.19.3 AS environment-go-dependences

# ## Configurando variáveis de Ambiente
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOPRIVATE=*
ENV TZ=America/Sao_Paulo
ENV APP_HOME=/build
ENV SERVER_PORT=80

# ## diretorio de trabalho
WORKDIR "$APP_HOME"

RUN go install github.com/codegangsta/gin@latest
RUN go install golang.org/x/lint/golint@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/go-critic/go-critic/cmd/gocritic@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

# ## copia o módulo
COPY go.* ./

# ## baixa as dependencias
RUN go mod download
RUN go mod tidy
RUN go mod vendor
RUN go mod verify

# ## copia a aplicacao para dentro da maquina
COPY ./ ./

RUN go vet .

## check quality
RUN golint .
RUN goimports
RUN gocritic check .
RUN staticcheck .

RUN go build -ldflags="-s -w" -o apiserver .

# ! ## Maquina para excutar o ambiente de producao
FROM scratch

# ## Labels
LABEL description="environment-go"

# ## Configurando variáveis de Ambiente
ENV SERVER_PORT=80

# ## copia a aplicacao para dentro da maquina
COPY --from=environment-go-dependences ["/build/apiserver", "/build/.env.prod", "/"]

# ## portas de acesso
EXPOSE "$SERVER_PORT"

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]
