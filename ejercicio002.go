package main

import "fmt"

func main() {

	horasTrabajadas := float32(0)
    //var horasTrabajadas int
    var costoHora float32 = 1.25
    var sueldo float32

    fmt.Println("Valor por hora es de ", costoHora)
    fmt.Print("Ingrese las horas trabajadas:")
    fmt.Scan(&horasTrabajadas)
    //fmt.Print("Ingrese el pago por hora:")
    //fmt.Scan(&costoHora)
    sueldo=float32(horasTrabajadas) * costoHora
    fmt.Println("El sueldo total del operario es ",sueldo)

    if sueldo > 10{

    	fmt.Println("Eres tremendo trabajador")
    }else{

    	fmt.Println("Eres un mal trabajador")

    }

    



}