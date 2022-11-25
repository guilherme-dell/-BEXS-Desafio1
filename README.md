# [BEXS] Desafio Bexs - API DE ROLES
## Objetivo do Desafio
Fazer uma API simples utilizando os recursos apresentados no workshop.

## Uso
Para rodar a aplicação primeiro execute o comando `docker compose build` após o build ocorrer execute o comando `docker-compose up`.
Agora sua API esta pronta para ser utilizada :)

### Utilizando a API
Para casdastrar um Role acesse o endpoint `localhost:8081`

	GET(localhost:8081/:roleID)  => PEGAR UM ROLE
	GET(localhost:8081/roles) => VER TODOS OS ROLES
	GET(localhost:8081/rolesHoje) => ROLES HOJE
	POST(localhost:8081/) => MARCAR UM ROLE
	PUT(localhost:8081/:roleID) => EDITAR UM ROLE CADASTRADO
	DELETE(localhost:8081/:roleID) => APAGAR UM ROLE

## JSON de cadastro
	{
    "nome": "NOME DO ROLE",
    "local": "LOCAL DO ROLE",
    "cidade": "CIDADE DO ROLE",
    "horario": "HORARIO DO ROLE",
    "dia": "DIA DO ROLE",
    "tipo do role": "TIPO DO ROLE"
	}

<font color="red">O campo dia deve ser preenchiido no formato `dia-mes-ano` | `25-11-2022`</font>