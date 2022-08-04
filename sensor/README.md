# Sensor

##Objetivo
Esse projeto possui o objetivo de estudo da linguagem golang. Neste projeto está sendo feito uma rescrita de um projeto .NET Core. 

Github NetCore: https://github.com/NouaraCandida/net-backend-ddd

## Criação do projeto
Para gerenciar os pacotes com mais flexibilidade o golang 1.13> possibilita criar  você criará seu próprio módulo Go público e adicionará um pacote ao seu novo módulo.Para informações detalhadas acesso o [link](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules). Para este projeto nosso módulo chamará sensor.

`go mod init sensor`

## Pasta DB
Inicialmente este projeto seguiu o padrão data first. Esse padrão foi escolhido por ser um reescritura de código, sem nenhuma alteração nas entidades. No entanto com evolução deste projeto o qual será criado diversos micro-serviços o padrão class first será adotado.

### Schema Migrations
Para este projeto utilizamos migration para versionamento da base de dados, criadas. O repositório do migrate para golang está disponível no [link](https://github.com/golang-migrate/migrate). A seguir está descrito o passo a passo para estrurar migrations dentro da sua solução.

1. Instale o pacote:

    `go get github.com/golang-migrate/migrate`

2. Para criar o sql dos seus serviços execute o comando a seguir. SERVICE representa o nome do serviço, isso criará uma pasta onde os arquivos sql serão do serviço serão inseridos, NAME será o nome do arquivo sql, será gerado dois, são eles:up e down.
|

    ` make migrate-create SERVICE=nome-servico NAME=nome-do-arquivo-sql`

3. Todas as automações referentes ao postgres serão inseridas dentro da pasta postgres.

4. Rode o docker utilizando o comando a seguir:

    `make run-docker-postgres`
Caso na execução ocorra o erro abaixo, siga as instruções abaixo:

        /usr/local/bin/docker-entrypoint.sh: /docker-entrypoint-initdb.d/init-database.sh: /bin/bash: bad interpreter: Permission denied

* Pegue o ID do seu container e de stop no container. Adicione permissão de execução para os arquivos .sh que você deseja.  E execute novamente o container.

    `docker container ls`

    `docker stop a53547da805e `
    
    `cd db/postgres`

    `chmod +x automation-database.sh`

    `docker-compose up postgres`



Exemplo da resposta do comando ls:    
| CONTAINER ID | IMAGE | COMMAND | CREATED | STATUS | PORTS | NAMES|
|--- |--- |--- |--- |--- |--- |--- |
| a53547da805e | postgres:13.4-alpine| "docker-entrypoint.s…"| 37 minutes ago | Up 4 minutes | 0.0.0.0:5432->5432/tcp  3 | ensor_postgres_1 |


### Conectando no banco postgres
Lembrando que o objetivo deste projeto é migrar um projeto .NET Core para golang. Então vamos a uma diferença muito importante: a conexão no banco de dados.

O caminho relativo da implementaçãp da conexão é db/postgres/dbpostgres.go
Nesta implementação aprendemos um novo conceito, no golang precisamos importar nossa bibliotecas, e ao construir a conexão com banco de dados precisamos importar de forma implicita a lib do postgres. Ta ai um segredinho, basta colocar um _ na frente da importação.

``` 
import (
	"database/sql"
	"fmt"
	"log"
	config "sensor/pkg/config/sensor"

	//importando de maneira implicita, quem utiliza é o pacote database/sql
	_ "github.com/lib/pq"
)
```


## Repositorio
 Palavras de Martin Fowler"O padrão  Repository faz a mediação entre o domínio e as camadas de mapeamento de dados, agindo como uma coleção de objetos de domínio em memória.....Conceitualmente, um repositório encapsula o conjunto de objetos persistidos em um armazenamento de dados e as operações realizadas sobre eles, fornecendo uma visão mais orientada a objetos da camada de persistência.....e também da suporte ao objetivo de alcançar uma separação limpa e uma forma de dependência entre o domínio e as camadas de mapeamento de dados."

 Para nossa aplicação vamos utilizar esse padrão. A pasta repositorio contém toda a implementação. Dentro da pasta model declaramos a interface que representa o contrato de repositorio de cada entidade. Essa implementação permite realizar o isolamento da camada de dados e possibilita a troca do banco de dados sem grandes impactos. Vamos evoluir essa implementação com o projeto mas já temos um ponto de partida.
``` 
type SensorRepositorio interface {
	Get(ctx context.Context, id uuid.UUID) (Sensor, error)
	Create(ctx context.Context, sensor Sensor) (Sensor, error)
}
```