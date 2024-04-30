# Jobs API em Golang

Este é um projeto simples de uma API em Go para gerenciar vagas de emprego, incluindo funcionalidades para cadastro, listagem, atualização, exclusão e pesquisa de vagas, além de autenticação de usuários com JWT.

## Endpoints

### Listar todas as vagas

```
GET /api/v1/openings
```

Retorna todas as vagas cadastradas.

### Criar uma nova vaga

```
POST /api/v1/opening
```

Cria uma nova vaga com os seguintes campos:
- `link`: URL da vaga
- `role`: Cargo da vaga
- `remote`: Booleano indicando se a vaga é remota ou não
- `salary`: Salário da vaga (em int64)
- `company`: Nome da empresa
- `location`: Local da empresa


### Atualizar uma vaga existente

```
PUT /api/v1/opening?id={id}
```

Atualiza uma vaga existente pelo seu ID com os mesmos campos mencionados anteriormente.

### Deletar uma vaga

```
DELETE /api/v1/opening?id={id}
```

Deleta uma vaga existente pelo seu ID.

### Pesquisar uma vaga pelo ID

```
GET /api/v1/opening/?id={id}
```

Retorna os detalhes de uma vaga específica pelo seu ID.

### Login de usuário

```
POST /api/v1/login
```

Autentica um usuário com os seguintes campos:
- `email`: E-mail do usuário
- `password`: Senha do usuário

### Registro de usuário

```
POST /api/v1/register
```

Registra um novo usuário com os mesmos campos de e-mail e senha.

## Tecnologias Utilizadas

- Go
- JWT para autenticação

## Configuração

1. Clone o repositório.
2. Instale as dependências usando `go mod tidy`.
3. Configure as variáveis de ambiente necessárias, como chaves JWT, endereço do banco de dados, etc.
4. Execute o aplicativo usando `go run main.go` ou compile-o com `go build`.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Autor

Gustavo Ferreira - [@gustasousagh](https://github.com/gustasousagh)

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.
