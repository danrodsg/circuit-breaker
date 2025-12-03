# ⚡ Circuit Breaker 

Este projeto é uma demonstração simples de como implementar o padrão **Circuit Breaker** (Disjuntor) usando a popular biblioteca `github.com/sony/gobreaker` em GoLang.

O Circuit Breaker é um padrão de resiliência crucial para sistemas distribuídos. Ele evita que uma aplicação tente repetidamente acessar um serviço externo que está falhando, permitindo que o serviço se recupere e prevenindo falhas em cascata.

## 🚀 Como Executar

Este projeto não requer dependências externas além do próprio pacote Go.

1.  **Clone o repositório:**
    ```bash
    git clone [SEU_LINK_DO_REPOSITÓRIO]
    cd go-circuit-breaker-demo
    ```

2.  **Verifique as dependências:**
    ```bash
    go mod tidy
    ```

3.  **Execute o programa:**
    ```bash
    go run main.go
    ```

## 🧠 Como Funciona o Circuit Breaker

O código simula 10 tentativas de chamada a um serviço externo (`mockService`) que tem uma alta taxa de falhas.

### 1. Simulação do Serviço (`mockService`)

A função `mockService` simula a chamada a um serviço externo.
* Ela tem uma chance de falha de **50%** (`rand.Intn(100) < 50`).
* Em caso de falha, retorna o erro `"error trying to process request"`.

### 2. Configuração do Circuit Breaker

O disjuntor (`ch`) é inicializado com as seguintes configurações:

| Configuração | Valor | Descrição |
| :--- | :--- | :--- |
| `Name` | `"MyCircuitBreakerService"` | Nome amigável do disjuntor. |
| `MaxRequests` | `1` | No estado **Half-Open**, permite apenas 1 requisição de teste. |
| `Interval` | `time.Second * 5` | Tempo de reinício da contagem de falhas/sucessos no estado **Closed**. |
| `Timeout` | `time.Second * 1` | Tempo que o disjuntor permanecerá no estado **Open** antes de ir para **Half-Open**. |
| `ReadyToTrip` | `func(counts Counts) bool` | **Condição para abrir (Trip)**. O disjuntor muda para **Open** se houver **mais de 2 falhas consecutivas**. |

### 3. Estados Observáveis

A saída no terminal mostrará o disjuntor transitando pelos seus estados:

* **Closed (Fechado):** As requisições são tentadas normalmente.
* **Open (Aberto):** Após atingir o limite de falhas, o disjuntor **bloqueia** as requisições, retornando um erro rápido (`Circuit Breaker is open`).
* **Half-Open (Meio-Aberto):** Após o `Timeout` de 1 segundo, ele permite a próxima requisição como teste. Se falhar, volta para **Open**. Se for bem-sucedida, volta para **Closed**.

A transição de estado é registrada no log pela função `OnStateChange`:

State change: Closed -> Open


## 🛠️ Detalhes do Código (`main.go`)

### Funções Principais

* `main()`: Configura e executa o *loop* de chamadas ao serviço simulado através do Circuit Breaker.
* `mockService(string) (string, error)`: Simula o serviço com 50% de chance de falha.
* `logStateChange(name, from, to gobreaker.State)`: Função de *callback* que é executada sempre que o estado do Circuit Breaker muda.
* `stateString(state gobreaker.State) string`: Função auxiliar para converter o estado do `gobreaker` em uma string legível.

### A Chamada Protegida

A chamada ao serviço é sempre feita através do método `Execute()` do Circuit Breaker:

```go
_, err := cb.Execute(func() (interface{}, error) {
    return mockService()
})
Se cb.Execute() retornar um erro que não é do mockService (como Circuit Breaker is open), significa que o disjuntor evitou a chamada ao serviço.

🔗 Dependência
Este projeto utiliza a implementação de Circuit Breaker da Sony:

Go

import "(https://github.com/sony/gobreaker)"
