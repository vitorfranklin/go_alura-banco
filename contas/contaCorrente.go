package contas

import (
	"banco/clientes"
	"strconv"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo Insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {

	podeDepositar := valorDoDeposito >= 0

	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Novo saldo: " + strconv.FormatFloat(c.saldo, 'f', 2, 64)
	} else {
		return "Não foi possível depositar"
	}

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.saldo >= valorDaTransferencia && valorDaTransferencia > 0 {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
