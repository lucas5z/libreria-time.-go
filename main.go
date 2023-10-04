package main

import (
	"fmt"

	timecomplemeto "github.com/lucas5z/time/time_complemeto"
	time_tiempo "github.com/lucas5z/time/time_tiempo"
)

func main() {
	fmt.Println("cosas time:")
	timecomplemeto.Varibles() //time. time.Second, time.Minute, time.Hour (1,2)

	fmt.Println("libreria time:")
	time_tiempo.Creacion_Time()     // (1 ,2)
	time_tiempo.Medir_Tiempo()      // (3)
	time_tiempo.Operaciones_time()  //suma tiempo (4,5)
	time_tiempo.Formateo_analisis() // (6,7,8)
	time_tiempo.Tempo_risdores()    //tenporisador (9,10)
	time_tiempo.Zona_Horaria()      // (11)
	time_tiempo.Dia_semana_actual() // me regrea el dia _ de la semana, ejm: lunes (12)
	time_tiempo.Comprobacion()      // (13) vemos si una fecha es posterio o anterios a una fecha
}
