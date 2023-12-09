
<h1 align="center"> Log Guardian </h1>

<p align="center"> <strong>Open-Source Logger Assistant</strong> </p>

[![link to Go version](https://img.shields.io/github/go-mod/go-version/fonteeboa/go-log-guardian)](https://img.shields.io/github/go-mod/go-version/fonteeboa/go-log-guardian)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/fonteeBoa/go-log-guardian)
[![go report card](https://goreportcard.com/badge/github.com/fonteeBoa/go-log-guardian "go report card")](https://goreportcard.com/report/github.com/fonteeBoa/go-log-guardian)

<span style="color:red;">**English README**</span>: Dive into the English version of the project's documentation [here](https://github.com/fonteeboa/go-log-guardian/blob/master/README_en_us.md)

O Log Guardian é uma biblioteca desenvolvida para padronizar e gerenciar logs de maneira eficiente e organizada em sistemas vizando a importância da tradução para o usuário final. Com a flexibilidade de lidar com diferentes tipos de logs, esta biblioteca proporciona uma estrutura consistente para a geração e gerenciamento de logs gerais e específicos, permitindo a integração com diferentes sistemas.

<h2 align="center"> <strong>Funcionalidades</strong> </h2>

🔹 Padronização de Logs: O Log Guardian oferece uma estrutura unificada para diferentes tipos de logs, desde logs de função, operações em banco de dados até logs de requisições.

🔹 Configuração Flexível: Permite a fácil integração com diferentes sistemas, possibilitando a customização e configuração dos logs de acordo com as necessidades específicas.

🔹 Conexão com Banco de Dados: Além da gestão dos logs, o Log Guardian pode se integrar a diferentes tipos de banco de dados, como PostgreSQL, MySQL, SQLite e MongoDB. A configuração é simples, utilizando variáveis de ambiente para especificar os detalhes de conexão.

🔹 Inserção Automática de Logs: Quando configurado corretamente com variáveis de ambiente, o Log Guardian é capaz de inserir automaticamente os logs no banco de dados especificado.

<h2 align="center"> <strong>Uso</strong> </h2>

O Log Guardian é flexível e se adapta à configuração do ambiente em que é executado. Se as variáveis de ambiente necessárias não estiverem configuradas, o Log Guardian ainda poderá retornar o modelo do log específico para inserção manual no banco de dados.

Caso as variáveis de ambiente estejam configuradas corretamente com os detalhes do banco de dados desejado, o Log Guardian é capaz de conectar automaticamente ao banco de dados especificado e inserir os logs diretamente na tabela correspondente. Ele retorna um valor booleano indicando o sucesso ou falha na inserção dos dados no banco.

Essa flexibilidade permite uma fácil integração e uso do Log Guardian em diferentes cenários de configuração, seja para apenas fornecer os modelos de log para inserção manual ou para realizar inserções automáticas no banco de dados configurado.

É recomendável consultar a seção de Configuração para detalhes sobre as variáveis de ambiente necessárias para uma configuração completa do Log Guardian.

<h2 align="center"> <strong>Configuração</strong> </h2>

O Log Guardian utiliza variáveis de ambiente para configurar suas operações de banco, incluindo definições de conexão com banco de dados e outras configurações essenciais. Aqui está a lista das variáveis de ambiente disponíveis:

<h4 align="center"> <strong>Banco de Dados Relacional</strong> </h4>

#### PostgreSQL
```
POSTGRES_HOST: Define o endereço do host para o PostgreSQL.
POSTGRES_EXTERNAL_PORT: Especifica a porta externa para o PostgreSQL.
POSTGRES_USER: Nome de usuário para autenticação no PostgreSQL.
POSTGRES_PASSWORD: Senha para autenticação no PostgreSQL.
POSTGRES_DB: Nome do banco de dados PostgreSQL a ser utilizado.
```
#### MySQL
```
MYSQL_HOST: Define o endereço do host para o MySQL.
MYSQL_PORT: Especifica a porta para o MySQL.
MYSQL_USER: Nome de usuário para autenticação no MySQL.
MYSQL_PASSWORD: Senha para autenticação no MySQL.
MYSQL_DBNAME: Nome do banco de dados MySQL a ser utilizado.
```
#### SQLite
```
SQLITE_PATH: Caminho do arquivo SQLite, se for o banco de dados escolhido.
```

<h4 align="center"> <strong>Banco de Dados NoSQL</strong> </h4>

#### MongoDB
```
MONGODB_URI: Define o URI de conexão para o MongoDB.
MONGODB_DBNAME: Nome do banco de dados MongoDB a ser utilizado.
```

<h4 align="center"> <strong>Configuração Geral</strong> </h4>

```
DATABASE_TYPE: Especifica o tipo de banco de dados a ser utilizado pelo Log Guardian (Valores: sqlite, postgres, mysql, mongodb).
```

<h4 align="center"> <strong>Observações</strong> </h4>

Para utilizar as funções automáticas do go-log-guardian, é obrigatório o uso da variável DATABASE_TYPE, pois algumas validações são executadas com base nesta variável antes de chamar as rotinas de inserção.

Certifique-se de fornecer valores válidos e corretos para cada uma dessas variáveis de ambiente. Isso garante uma conexão adequada e o funcionamento correto do Log Guardian com o banco de dados desejado.
