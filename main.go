package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)
	PagarBoleto(&contaExemplo, 60)
	fmt.Println(contaExemplo.ObterSaldo())

	// clienteBruno := clientes.Titular{Nome: "Bruno", CPF: "000.000.000-00", Profissao: "Médico"}
	// contaDoBruno := contas.ContaCorrente{Titular: clienteBruno}
	// fmt.Println(contaDoBruno)

	// contaDoVitor := contas.ContaCorrente{Titular: "Vitor", NumeroAgencia: 00453, NumeroConta: 12800, Saldo: 200.00}
	// contaDoPifao := contas.ContaCorrente{Titular: "Pifão", NumeroAgencia: 00453, NumeroConta: 12800, Saldo: 300.00}

	// fmt.Println(contaDoVitor)
	// fmt.Println(contaDoPifao)
	// fmt.Println(contaDoVitor.Transferir(100.00, &contaDoPifao))
	// fmt.Println(contaDoVitor)
	// fmt.Println(contaDoPifao)
	// fmt.Println(contaDoPifao.Transferir(41.00, &contaDoVitor))
	// fmt.Println(contaDoVitor)
	// fmt.Println(contaDoPifao)

	// fmt.Println(contaDoVitor)

	// var contaDaEliane *ContaCorrente
	// contaDaEliane = new(ContaCorrente)
	// contaDaEliane.titular = "Eliane"
	// contaDaEliane.numeroAgencia = 299
	// contaDaEliane.numeroConta = 154
	// contaDaEliane.saldo = 499.99
	// fmt.Println(*contaDaEliane)
	// fmt.Println(&contaDaEliane)
	// fmt.Println(contaDaEliane)
	// fmt.Println(contaDaEliane.saldo)
	// fmt.Println(contaDaEliane.Sacar(600.15))
	// fmt.Println(contaDaEliane.saldo)
	// fmt.Println(contaDaEliane.saldo)
	// fmt.Println(contaDaEliane.Depositar(1000))
	// fmt.Println(contaDaEliane.saldo)
	// fmt.Println(contaDaEliane.saldo)
	// fmt.Println(contaDaEliane.Depositar(-1))
	// fmt.Println(contaDaEliane.saldo)

	// var num int = 0553
	// fmt.Println(num)
	// num = 0554
	// fmt.Println(num)
}
