## CreditCardValidator
Este projeto é um validador de cartões de crédito usando o algoritmo de Luhn.

Para testar essa funcionalidade do projeto, é necessário enviar uma requisição HTTP POST para o endpoint localhost:3000. No corpo da requisição, adicione um cartão de crédito no seguinte formato JSON:

```sh
{
  "creditCard": "XXXX-XXXX-XXXX-XXXX"
}

```
A requisição irá retornar se o cartão é válido ou não.
Além disso, existe um outro algoritmo que cria um cartão aleatório no seguinte padrão: 16 dígitos de 0 a 10 cada.

Para acessar essa funcionalidade, é necessário fazer uma requisição GET no endpoint: localhost:3000/fake.

A cada solicitação, um novo número é gerado e é mostrado se ele é válido para um cartão de crédito ou não.

Executando o Projeto
Para executar o projeto, siga os passos abaixo:

- Configure o ambiente GoLang em sua máquina.
- Clone o projeto.
- Acesse o diretório.
- Execute o comando: go run main.go
Certifique-se de ter o GoLang instalado e configurado corretamente em sua máquina antes de executar o projeto.
