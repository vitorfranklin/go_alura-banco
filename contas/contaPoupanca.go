package contas

import (
	"banco/clientes"
	"strconv"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo Insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) string {

	podeDepositar := valorDoDeposito >= 0

	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Novo saldo: " + strconv.FormatFloat(c.saldo, 'f', 2, 64)
	} else {
		return "Não foi possível depositar"
	}

}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) bool {
	if c.saldo >= valorDaTransferencia && valorDaTransferencia > 0 {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}
