**O projeto sobe via docker-compose<br/>**

**Startando o projeto:<br/>**
Ao realziar o clone do projeto, será necessario entrar na pasta da aplicação "API" e no terminal digitar docker-compose up,
após a aplicação ficar **100%** online você já poderá começar a utilizar ela nas seguintes rotas:

**POST** -> localhost:3000/cliente -> Irá inserir o cliente passado via JSON no banco, e retornará o ID do cliente que foi cadastrado.<br/>
**GET** -> localhost:3000/cliente -> Ela retornará todos os clientes já existente.<br/>
**GET** -> localhost:3000/cliente/{UUID} -> Retornará um JSON de cliente cadastrado.<br/>
**PUT** -> localhost:3000/cliente/{UUID} -> Atualizará os dados do cliente escolhido pelo UUID.<br/>
**DELETE** -> localhost:3000/cliente/{UUID} -> Removerá o cliente selecionado da base.<br/>

Na pasta ConsumidorRabbitMq você poderá encontrar a aplicação que irá receber o cliente que foi salvo via POST em um JSON que ficará
salvo dentro da pasta ConsumidorRabbitMq/arquivos.<br/>

**Realizando os testes da aplicação:<br/>**
Os testes unitarios foram feitos para rodar após a aplicação estar online no docker, pois ela depende de um DB que será gerado exclusivamente para testes.<br/>

**Como testar?<br/>**
Iniciar o docker-compose up, após isso entrar no CLI da aplicação que recebeu o nome de "enviador" e então escrever go test ./...
caso deseje ver a cobertura do teste go test ./... -cover irá fazer o serviço.

