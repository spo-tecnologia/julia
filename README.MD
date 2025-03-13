### Gerando código e automatizando processos com Julia


## Tabela de conteúdos

- [Instalação](#Instalação)
- [Como utilizar](#Como-utilizar)
- [Estudo](#Estudo)
- [Deploy API GO](#Deploy-API)
- [Documentação](#Documentação)
- [Comandos úteis](#Comandos-úteis)

### Sobre
- [Visão Geral](#visão-geral)
- [Funcionalidades](#funcionalidades)
  - [Autenticação](#autenticação)
  - [Usuários](#usuários)
  - [Arquivos](#arquivos)
  - [Políticas de Privacidade](#políticas-de-privacidade)
  - [Samples (Exemplos)](#samples-exemplos)
  - [Status](#status)
- [Componentes](#componentes)
  - [`config`](#config)
  - [`enums`](#enums)
  - [`http`](#http)
    - [Controllers](#controllers)
    - [Middlewares](#middlewares)
    - [Requests](#requests)
    - [Responses](#responses)
  - [`jobs`](#jobs)
  - [`models`](#models)
  - [`notifications`](#notifications)
  - [`policies`](#policies)
  - [`repository`](#repository)
  - [`routes`](#routes)
  - [`services`](#services)
  - [`tests`](#tests)
- [Documentação](#documentação)
- [Banco de Dados](#banco-de-dados)
- [Enums](#enums)
- [Notificações](#notificações)
- [Políticas](#políticas)
- [Repositórios](#repositórios)
- [Rotas](#rotas)
- [Autenticação](#autenticação)
- [Validação de Requests](#validação-de-requests)
- [Testes](#testes)
- [Variáveis de Ambiente](#variáveis-de-ambiente)

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
TEMPERATURE= #Controle de aleatoriedade da resposta.  
TOP_P= #Parâmetro de diversidade. Exemplo: 1  
MAX_OUTPUT_TOKENS= #Limite de tokens na resposta. Exemplo: 24000  
TOP_K= #Parâmetro de controle de palavras mais prováveis. Exemplo: 64  
RESPONSE_MIME_TYPE= #Tipo de resposta. (text/plain).  
GENAI_API_KEY=  #Chave da API Gemini.  

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

localhost:8080/api/documentation/index.html

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

# Julia - Backend API

Julia é uma API backend construída em Go (Golang) utilizando o framework Gin. Este documento fornece uma visão geral da estrutura, funcionalidades e como configurar e executar o projeto.

## Visão Geral

A API foi projetada seguindo uma arquitetura modular, com separação de responsabilidades entre as camadas. Abaixo, detalhamos cada um dos componentes principais.

## Funcionalidades

### Autenticação

*   **Login:** Autenticação de usuários via email/CPF e senha, gerando tokens JWT para acesso seguro às rotas protegidas.
*   **Esqueci minha senha:** Recuperação de senha através do envio de um email com um link para redefinição.
*   **Recuperar Senha:** Rota para redefinição de senha após o usuário clicar no link recebido por email.

### Usuários

*   **CRUD Completo:** Criação, leitura, atualização e exclusão de usuários.
*   **Seleção de Usuários:** Listagem de usuários com suporte a pesquisa, paginação e filtragem.
*   **Atualização de FCM Token:** Rota específica para atualizar o token FCM (Firebase Cloud Messaging) de um usuário.

### Arquivos

*   **Upload de Arquivos:** Upload de arquivos através de base64 ou URL pública.
*   **Listagem de Arquivos:** Listagem de todos os arquivos.
*   **Exclusão de Arquivos:** Exclusão de arquivos por ID.

### Políticas de Privacidade

*   **Política de Privacidade:** Rota para exibir a política de privacidade da aplicação.
*   **Política de Exclusão de Dados:** Rota para exibir a política de exclusão de dados do usuário.

### Samples (Exemplos)

*   **CRUD Completo:** Criação, leitura, atualização e exclusão de modelos de exemplo.
*   **Seleção de Samples:** Listagem de modelos de exemplo com suporte a pesquisa, paginação e filtragem.

### Status

*   **Status da API:** Rota para verificar o status geral da API.

## Componentes

### `config`

*   `database.go`: Inicialização e configuração do banco de dados (MySQL). Define a conexão, realiza as migrations e cria/atualiza a conta de administrador.

### `enums`

*   `sample_enum.go`: Definição de um enum de exemplo.
*   `user_profile_enum.go`: Definição dos enums para os perfis de usuário (Administrador e Usuário comum).

### `http`

*   `controllers`:
    *   `auth_controller.go`: Lógica para autenticação (login, esqueci minha senha, recuperar senha).
    *   `file_controller.go`: Lógica para manipulação de arquivos (upload, listagem, exclusão).
    *   `policy_controller.go`: Lógica para exibição das políticas de privacidade.
    *   `sample_controller.go`: Lógica para manipulação dos modelos de exemplo (CRUD, listagem).
    *   `status_controller.go`: Lógica para o endpoint de status da API.
    *   `user_controller.go`: Lógica para manipulação de usuários (CRUD, listagem, atualização de FCM token).
*   `middlewares`:
    *   `block_middleware.go`: Middleware para bloquear requisições suspeitas.
    *   `cors_middleware.go`: Middleware para configurar o CORS (Cross-Origin Resource Sharing).
    *   `json_logger.go`: Middleware para logar as requisições em formato JSON.
    *   `jwt_auth_middleware.go`: Middleware para autenticação JWT.
*   `requests`:
    *   `custom_validators.go`: Validações customizadas para os inputs das requisições.
    *   `create_sample_model_input.go`: Define a estrutura esperada para a criação de um SampleModel
    *   `update_sample_model_input.go`: Define a estrutura esperada para a atualização de um SampleModel
    *   `create_file_input.go`: Define a estrutura esperada para a criação de um arquivo.
*   `responses`: (Não existe, mas poderia conter estruturas para padronizar as respostas da API)

### `jobs`

*   `job_client.go`: Cliente para processamento de jobs assíncronos.
*   `job_service.go`: Serviço para registrar jobs.
*   `job.go`: Define a estrutura de um Job.
*   `sample_job.go`: Exemplo de um job assíncrono.

### `models`

*   `file.go`: Define a estrutura do modelo de arquivo.
*    `sample_detail.go`: Define a estrutura do modelo de SampleDetail.
*   `sample_model.go`: Define a estrutura do modelo de exemplo.
*    `sample_item.go`: Define a estrutura do modelo de SampleItem.
*   `user.go`: Define a estrutura do modelo de usuário.
*   `default_model.go`: Define o modelo padrão para todos os models

### `notifications`

*   `fcm_notifier.go`: Notificador para enviar notificações via Firebase Cloud Messaging (FCM).
*   `forgot_password.go`: Notificação de "esqueci minha senha".
*   `mail_notifier.go`: Notificador para enviar emails.
*   `notification.go`: Interface para os notificadores.
*   `sample_notification.go`: Notificação de exemplo.
*   `whatsapp_notifier.go`: Notificador para enviar mensagens via WhatsApp.

### `policies`

*   `sample_policy.go`: Define as políticas de acesso para os modelos de exemplo.
*   `user_policy.go`: Define as políticas de acesso para os usuários.

### `repository`

*   `file_repository.go`: Lógica de acesso aos dados para os arquivos.
*   `sample_repository.go`: Lógica de acesso aos dados para os modelos de exemplo.
*   `user_repository.go`: Lógica de acesso aos dados para os usuários.
*    `job_repository.go`: Lógica de acesso aos dados para os jobs

### `routes`

*   `auth.go`: Define as rotas de autenticação.
*   `files.go`: Define as rotas para manipulação de arquivos.
*   `policies.go`: Define as rotas para as políticas de privacidade.
*   `routes.go`: Inicializa e configura todas as rotas da API.
*   `samples.go`: Define as rotas para os modelos de exemplo.
*   `status.go`: Define a rota para o status da API.
*   `swagger.go`: Define as rotas para a documentação Swagger.
*   `users.go`: Define as rotas para manipulação de usuários.

### `services`

*   `sentry.go`: Inicialização e configuração do Sentry para rastreamento de erros.
*   `token`:
      * `token.go`: Lógica para geração e validação de tokens JWT.
*   `translation.go`: Inicialização e configuração da internacionalização (i18n).
*   `util.go`: Funções utilitárias.
*   `validation.go`: Funções para manipulação de erros de validação.

### `tests`

*   `auth_test.go`: Testes para a autenticação.
*   `file_test.go`: Testes para a manipulação de arquivos.
*   `main_test.go`: Configuração e inicialização para os testes.
*   `sample_test.go`: Testes para os modelos de exemplo.
*   `user_test.go`: Testes para a manipulação de usuários.
*   `fakers`:
      * `fakers.go`: Funções para gerar dados fake.
      * `gofakeit.go`: Implementação do pacote gofakeit
*   `factories`:
      * `sample_detail_factory.go`: Factory para criar SampleDetail
      * `sample_factory.go`: Factory para criar SampleModel
      * `file_factory.go`: Factory para criar File
      * `user_factory.go`: Factory para criar User

## Documentação

A documentação da API é gerada automaticamente utilizando Swagger. Para acessá-la, execute a aplicação e acesse `/swagger/index.html` no seu navegador.

## Banco de Dados

O projeto utiliza o MySQL como banco de dados. As configurações de acesso (host, usuário, senha, nome do banco, porta) são definidas no arquivo `.env`. As migrations são executadas automaticamente na inicialização da aplicação.

## Enums

Os enums são utilizados para definir conjuntos de valores predefinidos, como os perfis de usuário. Eles facilitam a manutenção do código e garantem a consistência dos dados.

## Notificações

O sistema de notificações permite o envio de mensagens para os usuários através de diferentes canais (email, FCM, WhatsApp). A interface `Notifier` define o contrato para os notificadores, e cada implementação é responsável por enviar a mensagem através do seu respectivo canal.

## Políticas

As políticas de acesso definem as permissões dos usuários para acessar e manipular os recursos da API. Elas são implementadas através das classes `UserPolicy` e `SamplePolicy`, que verificam se o usuário tem as permissões necessárias para realizar uma determinada ação.

## Repositórios

Os repositórios são responsáveis por encapsular a lógica de acesso aos dados. Eles fornecem uma interface para realizar operações CRUD (Create, Read, Update, Delete) nos modelos.

## Rotas

As rotas definem os endpoints da API e seus respectivos handlers (controllers). Elas são configuradas no arquivo `routes/routes.go`.

## Autenticação

A autenticação é implementada utilizando JWT (JSON Web Tokens). Os tokens são gerados após o login do usuário e utilizados para autenticar as requisições às rotas protegidas.

## Validação de Requests

A validação dos dados de entrada é realizada utilizando o pacote `validator`. As validações são definidas nas structs dos requests, e são executadas automaticamente pelo Gin. Além das validações padrão, o projeto também define validações customizadas, como `exists` e `not_exists`, para verificar se um valor existe ou não no banco de dados.

## Testes

O projeto possui testes unitários e de integração para garantir a qualidade do código. Os testes cobrem os controllers, models, repositórios e outros componentes da API.

## Variáveis de Ambiente

O projeto utiliza variáveis de ambiente para configurar diferentes aspectos da aplicação, como acesso ao banco de dados, configurações do Sentry, API Key do Gryphon, etc. As variáveis de ambiente são definidas no arquivo `.env`.
