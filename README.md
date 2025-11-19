# Backend Piscord

Serviço backend para o chat Piscord, desenvolvido em Go com Gorilla WebSockets, autenticação JWT e MongoDB.

## Executando o Backend

A forma mais fácil de executar o backend é utilizando o repositório principal de orquestração, que automatiza todo o setup pelo Kubernetes:

👉 [Clique para acessar o repositório principal de orquestração](https://github.com/davmp/piscord-app)

No repositório principal, basta seguir o passo a passo para subir todos os serviços automaticamente.

## Tecnologias

- Go 1.21+
- Gorilla Mux / WebSockets
- MongoDB Atlas
- Autenticação JWT

## Docker & CI/CD

- Imagem Docker pronta para deploy
- CI/CD publica a imagem no Docker Hub automaticamente

## Variáveis de Ambiente (Docker)

Essas variáveis podem ser executadas ao executar seu container.

| Variável   | Descrição                                      | Exemplo                                   |
| ---------- | ---------------------------------------------- | ----------------------------------------- |
| MONGO_URI  | URI de conexão com o banco de dados MongoDB    | mongodb://_user_:_password_@_host_:_port_ |
| JWT_SECRET | Chave secreta para autenticação JWT            | cGlzY29yZDMyMQ== (_piscord321_)           |
| PORT       | Porta que o backend irá escutar (padrão: 8000) | 8000                                      |
