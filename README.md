# secret

Inteiramente escrito na linguagem de programação [**go**](https://golang.org/), o secret é uma aplicação de rede que garante o envio seguro de dados sobre uma conexão TCP.

Nosso protocolo, chamado *secret*, possui o seguintes passos:

1. Uma chave simetrica e secreta é compartilhada previamente entre as partes;
1. Inicia-se um pedido de conexão TCP na porta que o servidor estiver escutando;
1. O cliente cria um mensagem e adiciona o numero de sequencia ao final.
1. É calculado o MAC da mensagem
1. A mensagem+sequencia+MAC são cifradas pelo RSA e enviadas
1. O servidor decifra a mensagem pelo RSA
1. O servidor verfica o MAC da mensagem
1. O servidor verfica o número de sequência da mensagem

## Garantias

Nosso protocolo garante os seguintes aspectos:

### Privacidade

Utilizamos uma criptografia de chave asssimétrica, que foi previamente compartilhada entre as partes, para assegurar que uma parte terceira não possa compreender a mensagem.

### Autenticidade

Além da criptografia assimétrica que garante que a mensagem foi gerada pelo portador da chave previamente compartilhada, a mensagem é cifrada com um número de sequência para evitar ataques de repetição.

### Integridade
Antes de ciframos a mensagem geramos um MAC da mensagem+sequencia+chave usando como algoritmo de hash o CRC e anexamos na mensagem.

## Descrição do secret:

O secret tem uma arquitetura cliente-servidor, e toda a comunicação é feita sobre o protocolo TCP, ou seja, o servidor escuta uma porta TCP e aguarda uma conexão do cliente nesta porta. O servidor descriptografa os dados enviados e salva em arquivo. Após o término da conexão, o servidor também finaliza sua execução.

O cliente lê um arquivo e envia de forma cifrada ao servido. Após o envio do arquivo, o cliente termina a conexão e finaliza sua execução.

### Estrutura da Implementação

Temos como principais estruturas:

*   **CryptoMessage**: Interface que define os comportamentos das mensagens
*   **MacMessage**: Estrutura que encapsula uma CryptoMessage qualquer e adiciona uma verificação de um MAC usando CRC
*	**RSAMessage**: Estrutura que encapsula uma CryptoMessage qualquer e cifra usando RSA com uma chave pré gerada
*	**RSA**: Estrura que representa uma chave RSA
*	**CRC**: Implementação de um CRC de oito bits
*   **StandartCypher**: Principal estrutura para o funcionamento do protocolo, responsável por cifrar e decifrar a mensagem, utilizando as estruturas descritas acima.

Nossa arquitetura se baseia em um sistema em camadas assim, podemos aplicar qualquer combinação de cifras  que desejarmos

## Executando o secret:

* Instale go na máquina;

    * Debian, Ubuntu e outros sistemas com *apt-get*

    ```bash
    # apt-get install golang
    ```

    * Fedora, CentOS, RHEL e outros sistemas com *yum*

    ```bash
    # yum install golang
    ```

    * Outros sistemas operacionais: 

        Siga as [instruções de instalação](https://golang.org/doc/install). 

* Defina a variavel de ambiente 'GOPATH' para uma pasta existente onde seu usuário tenha permissão de escrita;

* Execute o seguinte commando:
```bash
$ go get github.com/leoferlopes/secret
```

* Vá para a pasta do projeto:
```bash
$ cd $GOPATH/src/github.com/leoferlopes/secret
```

* Compile o projeto
```bash
$ go build
```

* Execute o executavel gerado com o nome de secret[.exe] e veja as instruções de uso
```bash
$ secret --help
$ secret server --help
$ secret client --help
```

* Para executar o servidor, digite o seguinte comando:
```bash
$ secret server --file=LOCAL_DO_ARQUIVO --port=PORTA
```

* Cliente 
```bash
$ secret server --file=LOCAL_DO_ARQUIVO --server=ENDEREÇO_DO_SERVIDOR
```
