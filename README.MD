### Gerando código e automatizando processos com Julia


## Tabela de conteúdos

- [Instalação](#Instalação)
- [Como utilizar](#Como-utilizar)
- [Estudo](#Estudo)
- [Deploy](#Deploy-API)
- [Documentação](#Documentação)
- [Comandos úteis](#Comandos-úteis)

## Instalação

O projeto pode ser executado usando Docker.

Primeiro crie o arquivo .env e ajuste as variáveis de ambiente:
```shell
cp .env.example .env
```

As principais configurações que devem ser feitas são:

APP_ENV= #Ambiente da aplicação (local).
APP_NAME= #Nome da aplicação ("Julia").
APP_URL=  #URL da aplicação (http://localhost:8080).
TEMPERATURE= #Controle de aleatoriedade da resposta .
TOP_P= #Parâmetro de diversidade. Exemplo: 1
MAX_OUTPUT_TOKENS= #Limite de tokens na resposta. Exemplo: 24000
TOP_K= #Parâmetro de controle de palavras mais prováveis. Exemplo: 64
RESPONSE_MIME_TYPE= #Tipo de resposta. (text/plain).
GENAI_API_KEY=  #Chave da API para IA.

Construa a imagem docker:
```shell
docker-compose build julia_code
```
Execute o container 
```shell
docker-compose up -d julia_code
```


## Como utilizar

Insira no arquivo 'TODO.MD' a descrição da sua tarefa

#### Exemplo

Rodar uma tarefa de um CRUD de uma nova entidade:
Tarefa: CRUD company
Campos
name: tipo texto, obrigatório
cnpj_cpf, tipo texto, obrigatório, único no banco de dados
phone, tipo texto, obrigatório
address, tipo texto, obrigatório

Execute no terminal

```shell
docker exec -it julia_code python3 /app/julia.py --file TODO.MD`
```

Você vera uma saída semelhante a esta abaixo com o passo a passo gerado

Assinaturas de métodos extraídas de 80 arquivos.
Lista de arquivos a serem criados e alterados:
models/company.go
http/controllers/company_controller.go
http/requests/create_company_input.go
http/requests/update_company_input.go
policies/company_policy.go
repository/company_repository.go
routes/companies.go
tests/company_test.go
tests/factories/company_factory.go


## Estudo

[Estudo de assertividade](https://zenodo.org/records/15019843)

Este estudo avalia a precisão e eficiência do projeto Julia Code, uma solução baseada em Generative AI para automatizar a escrita de código dentro de estruturas predefinidas. Embora o sistema demonstre a capacidade de gerar código funcional, seu desempenho não foi extensivamente analisado, tornando necessário validar seus resultados para apoiar decisões sobre sua aplicabilidade.

A pesquisa foi conduzida por meio de um benchmark estruturado, onde tarefas de vários níveis de complexidade (baixo, médio e alto) foram executadas. As métricas avaliadas incluem consumo de tokens, número de arquivos implementados corretamente, execução de testes automatizados, conformidade com requisitos funcionais e aderência a padrões de qualidade.

Os resultados indicam que o Julia Code é altamente eficaz para tarefas de baixa e média complexidade, garantindo a geração correta de código e a execução bem-sucedida de testes automatizados. No entanto, em cenários mais complexos envolvendo múltiplas entidades, os requisitos de qualidade não foram totalmente atendidos. A análise sugere que dividir tarefas complexas em subtarefas menores pode melhorar a eficiência e a conformidade com os padrões de implementação.

## Deploy API

Após a construção do projeto você pode executar rodando os comando abaixo:

####  Usando binário GO

Execute no terminal

```shell
go run main.go
```

#### Usando Docker

Construa a imagem docker:
```shell
docker-compose build julia_api_go
```
Execute o container 
```shell
docker-compose up -d julia_api_go
```

A porta padrão de execução é: 8080

### Documentação

Após rodar a API Go Julia você pode acessar a documentação através do link

localhost:6565/api/documentation/index.html

###  Comandos úteis

Instalar Swagger localmente administrar a documentação
```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

Adicionar o binário Swagger no caminho linux

```shell
export PATH=$(go env GOPATH)/bin:$PATH
```

Gerar ou Atualizar documentação
```shell
swag init --parseDependency --parseInternal
```

Rodar os testes

```shell
go test
```

Instalar Go Imports

```shell
go install golang.org/x/tools/cmd/goimports@latest
```
