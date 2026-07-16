package main

import (
	"fmt"
)

func alterarNome(nome *string) {
	fmt.Println("Digite o nome: ")
	fmt.Scan(nome)
	fmt.Println("Nome alterado para: ", *nome)
}

func alterarCargo(cargo *string) {
	fmt.Println("Digite o Cargo: ")
	fmt.Scan(cargo)
	fmt.Println("Cargo alterado para: ", *cargo)
}

func alterarSalario(salario *float64) {

	var salarioOP float64
	fmt.Println("Digite o Salário: ")
	fmt.Scan(&salarioOP)

	if salarioOP < 0 {
		fmt.Println("ERRO: O sálario não pode ser menor que 0.")
	} else {
		*salario = salarioOP
		fmt.Println("Salário alterado para: ", *salario)
	}

}

func mostrarFuncionario(nome *string, cargo *string, salario *float64) {
	fmt.Println("Nome: ", *nome)
	fmt.Println("Cargo: ", *cargo)
	fmt.Println("Salário: R$", *salario)
}

func limparCadastro(nome *string, cargo *string, salario *float64) {
	*nome = ""
	*cargo = ""
	*salario = 0.0
	fmt.Println("Cadastro limpo!")
}

func aumentarSalario(salario *float64) {
	var prct float64
	fmt.Println("Qual valor deseja aumentar em (%)")
	fmt.Scan(&prct)

	if prct < 0 {
		fmt.Println("O Aumento não pode ser negativo!")
		return
	} else {

		calculo := prct / 100
		valorAumento := *salario * calculo
		*salario = *salario + valorAumento
		fmt.Println("Salário aumentado!")
	}

}

func diminuirSalario(salario *float64) {
	var prct float64
	fmt.Println("Qual valor deseja diminuir em (%)")
	fmt.Scan(&prct)

	if prct >= 100 {
		fmt.Println("Não tem como diminuir em 100%, ou superior")
	} else if prct < 0{
		fmt.Println("ERRO: não tem como diminuir o salário menor que 0")
	} else {

		calculo := prct / 100
		valorDesconto := *salario * calculo
		*salario = *salario - valorDesconto
		fmt.Println("Salario Descontado!")
	}

}

func main() {
	fmt.Println("Sistema de Funcionário")

	var op int
	nome := ""
	cargo := ""
	salario := 0.0

	for {
		fmt.Println("1. Alterar Nome")
		fmt.Println("2. Alterar Cargo")
		fmt.Println("3. Alterar Salário")
		fmt.Println("4. Aumentar Salário (%)")
		fmt.Println("5. Diminuir Salário (%)")
		fmt.Println("6. Mostrar funcionário")
		fmt.Println("7. Limpar cadastro")
		fmt.Println("8. Sair...")
		fmt.Scan(&op)

		switch op {
		case 1:
			alterarNome(&nome)
			fmt.Println(nome)
		case 2:
			alterarCargo(&cargo)
			fmt.Println(cargo)
		case 3:
			alterarSalario(&salario)
			fmt.Println(salario)
		case 4:
			aumentarSalario(&salario)
			fmt.Println(salario)
		case 5:
			diminuirSalario(&salario)
			fmt.Println(salario)
		case 6:
			mostrarFuncionario(&nome, &cargo, &salario)
		case 7:
			limparCadastro(&nome, &cargo, &salario)
		case 8:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção Invalida!")
		}
	}
}
