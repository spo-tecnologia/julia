# Imagem base com Python
FROM python:3.11-slim

# Definir diretório de trabalho
WORKDIR /app

# Copiar o arquivo de requisitos
COPY . .

# Definir o PATH no ambiente do contêiner
ENV PATH="/usr/local/go/bin:${PATH}"

# Remover Go antigo e instalar uma versão nova
RUN apt-get update && \
    apt-get install -y wget tar && \
    rm -f $(which go) || true && \
    mkdir -p /usr/local/go && \
    wget https://golang.org/dl/go1.23.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xvzf go1.23.6.linux-amd64.tar.gz && \
    rm go1.23.6.linux-amd64.tar.gz && \
    echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile && \
    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc

# Definir o PATH para incluir o diretório de binários do Go
ENV PATH="$PATH:/root/go/bin"

# Instalar dependências do Python
RUN pip install --no-cache-dir python-dotenv google-generativeai requests

# Instalar dependências do Go
RUN go install golang.org/x/tools/cmd/goimports@latest

# Instalar swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest