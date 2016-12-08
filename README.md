# secret

Nosso protocolo,  chamado 'secret', possui o seguintes passos

Obs: Todas as mensagem incluem um MAC antes de serem cifradas

1. Uma chave simetrica e secreta é compartilhada previamente entre as partes;
1. Inicia-se um pedido de conexão TCP na porta 1280;
1. O destino envia um nounce de 16 bytes criptografado com a chave simetrica;
1. A origem recebe o nounce cifrado o decifra e troca os dois bytes de lugar, o primeiro byte passa a ser o segundo e o segundo passa a ser o primeiro, o cifra novamento e o envia de volta ao destino;
1. O destino decifra o nounce e inverte os bytes novamente e verifica se é o mesmo que foi enviado, caso seja ele envia sua chave publica para a origim cifrada com a chave simetrica.
1. A origem passa a enviar mensagens para o destino cifradas com a chave simetrica e a chave publica do destino;
1. Quando a conexão deseja ser fechada é enviada um comando especial de fechamento da conexão;

## Garantias

Nosso protocolo garante os seguintes aspectos:

### Privacidade
Utilizamos uma criptografia de chave simetrica, que foi previamente compartilhada entre as partes, para assegura que uma parte terceira não possa compreender a mensagem, utilizamos o algoritmo baseado em XOR + uma implementação de RSA simplificado

### Autenticidade
Ao iniciarmos a conexão é necessário o envio de um nounce, cifrado com a chave simetrica que apenas as partes interessadas possuem

`@TODO: enviar um nounce`

### Integridade
Antes de ciframos a mensagem geramos um MAC da mensagem+chave simetrica usando como algoritmo de hash o CRC e anexamos na mensagem.

## Descrição do programa:

`@TODO: falar da parte de redes`

Temos como principais estruturas:

*   CryptoMessage: Interface que define os comportamentos das mensagens
*   MacMessage: Estrutura que encapsula uma CryptoMessage qualquer e adiciona uma verificação de um MAC usando CRC
*	RSAMessage: Estrutura que encapsula uma CryptoMessage qualquer e cifra usando RSA com uma chave pre gerada
*	XORMessage: Estrutura que encapsula uma CryptoMessage qualquer e cifra usando XOR com uma chave pre gerada
*	RSA: Estrura que representa uma chave RSA
*	CRC: Implementação de um CRC de oito bits

Nossa arquitetura se basea em um sistema em camadas assim, podemos aplicar qualquer combinação de cifras  que desejarmos

## Executando o secret:

* Instale go na máquina;

    * Debian

    ```bash
    # apt-get install golang
    ```

    * Fedora

    ```bash
    # yum install golang
    ```

    * Outros sistemas operacionais: 

        Siga as [instruções de instalação](https://golang.org/doc/install). 

* Defina a variavel de ambiente 'GOPATH' para uma pasta existente onde seu usuário tenha permissão de escrita;

* Execute o seguinte commando:
```bash
$ go get -v https://github.com/leoferlopes/secret
```

* Vá para a pasta do projeto:
```bash
$ cd $GOPATH/src/github.com/leoferlopes/secret
```

* Compile o projeto
```bash
$ go build
```

* Execute o executavel gerado com o nome de secret[.exe]
```bash
$ secret --help
```