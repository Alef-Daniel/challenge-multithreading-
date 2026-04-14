
# Challenge Multithreading - Full Cycle

Neste desafio, Você deve construir um sistema que busca o endereço mais rápido entre duas APIs distintas. O foco é lidar com requisições simultâneas, tratamento de concorrência e timeouts.

APIs Alvo Você deve consumir as seguintes APIs de consulta de CEP:

    BrasilAPI: https://brasilapi.com.br/api/cep/v1/{cep}
    ViaCEP: http://viacep.com.br/ws/{cep}/json/

Requisitos Técnicos

1. Requisições Simultâneas: 

    O seu sistema deve fazer a requisição para as duas APIs ao mesmo tempo (paralelamente).

2. Race Condition (Corrida):

    O sistema deve aceitar apenas a resposta da API que responder mais rápido e descartar a resposta da outra (mais lenta).

3. Output (Saída): O resultado da requisição deve ser exibido na linha de comando (terminal), contendo:

    Os dados do endereço recebido.

    Qual API entregou a resposta (BrasilAPI ou ViaCEP).

4. Timeout: O tempo limite de resposta é de 1 segundo.

    Caso nenhuma das APIs responda dentro desse tempo, o sistema deve exibir um erro de timeout.

 

Tecnologias

    Linguagem: Go (Golang)

    Conceitos: Goroutines, Channels, Select, Package net/http.





## Run Locally

Clone the project

```bash
  git clone https://github.com/Alef-Daniel/challenge-multithreading
```

Go to the project directory

```bash
  cd challenge-multithreading-
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run cmd/main.go
```


## Scenarios

### Timeout

If you wish to simulate a timeout error, go to line 26 of the file`internal/application/usecase/process_address_usecase.go` and change it to `time.Millisecond`.

### Response OK

If everything goes well, the body you will receive will be:
```json
{
  "Cep": "09330340",
  "Logradouro": "Rua Angelim Milanez",
  "Bairro": "Jardim Luzitano",
  "UF": "SP",
  "Provider": "BrasilAPI"
}
```
