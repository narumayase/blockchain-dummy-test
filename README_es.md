# blockchain-dummy-test

Pasaje de código de JS a Golang. Ejemplo de funcionamiento de una blockchain simple.

Fuente: https://medium.com/geekculture/blockchain-explained-in-50-lines-of-code-1dbf4eda0201 

## Construido con 🛠

* [Golang](https://golang.org/) - Lenguaje de programación.

## Comenzando 🚀

### Pre-requisitos

Herramientas necesarias para la ejecución local del servicio:

- Go go1.16+.

* Para ejecutar los tests:

```
$ go test ./...

ok      blockchain-dummy-test/pkg       1.113s
```

* Para evaluar coverage:

```
$ go test -coverprofile=coverage.out ./...

ok      blockchain-dummy-test/pkg       5.300s  coverage: 92.3% of statements
```

* Coverage de manera más detallada:

```
$ go tool cover -func=coverage.out

blockchain-dummy-test/pkg/block.go:19:          NewBlock        100.0%
blockchain-dummy-test/pkg/block.go:30:          CalculateHash   100.0%
blockchain-dummy-test/pkg/block.go:36:          Mine            100.0%
blockchain-dummy-test/pkg/block.go:43:          zeros           100.0%
blockchain-dummy-test/pkg/block.go:47:          String          75.0%
blockchain-dummy-test/pkg/blockchain.go:9:      NewBlockchain   100.0%
blockchain-dummy-test/pkg/blockchain.go:14:     AddBlock        100.0%
blockchain-dummy-test/pkg/blockchain.go:21:     IsValid         87.5%
total:                                          (statements)    92.3%
```

* Ejemplo de cómo se ve:

![image](https://user-images.githubusercontent.com/76977457/144313732-efa20532-c466-4fce-bcf1-495205b92581.png)


## Roadmap

- [x] Estructura de proyecto
- [x] Pensar algoritmo
- [x] Readme
- [x] Algoritmo terminado
- [x] Test unitarios
 
