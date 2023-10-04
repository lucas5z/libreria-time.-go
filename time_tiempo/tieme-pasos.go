package time_tiempo

import (
	"fmt"
	"time"
)

//creacion de obejto time.Time

func Creacion_Time() {
	//creacion time actual---------------------------------------------
	fecha := time.Now()
	fmt.Println("1->", fecha)
	//creacion de un time completo -------------------------------
	dia1 := time.Date(2022, time.February, 7, 12, 0, 0, 0, time.UTC)
	dia2 := time.Date(2023, time.November, 7, 12, 0, 0, 0, time.UTC)
	fmt.Println("2->", dia1, dia2)
}

//medir un tiempo con time

func Medir_Tiempo() {

	start := time.Now()

	for i := 0; i < 5; i++ {
		fmt.Println("3-> hola")
	}

	end := time.Now()

	t_demora := end.Sub(start)
	fmt.Println("3->", t_demora)

}

//operaciones time

func Operaciones_time() {
	//suma de tiempo---------------------------------------------------
	start := time.Now()
	duration := 5 * time.Second
	//start.Add  ->suma el tiempo de inico con el tiempo q le pases
	tiempo_total := start.Add(duration)
	fmt.Println("4->", tiempo_total)

	// calculo de tiempo trascurrido------------------------------------
	//Time.Sub se utiliza para calcular
	//la diferencia de tiempo entre dos objetos time.Time.
	dia1 := time.Date(2022, time.February, 7, 12, 0, 0, 0, time.UTC)
	dia2 := time.Date(2023, time.November, 7, 12, 0, 0, 0, time.UTC)

	duracion := dia2.Sub(dia1)

	fmt.Printf("5-> Duración entre las dos fechas: %v horas y %v minutos\n", duracion.Hours(), duracion.Minutes())

}

//formateo y analizis de fecho y hora

func Formateo_analisis() {
	//FORMATEO---------------------------------------------------------
	hour := time.Now()

	formato_RFC3339 := hour.Format(time.RFC3339)
	fmt.Println("6->", "Hora en formato RFC3339:", formato_RFC3339)

	//formato personalizado
	formato_personalisado := hour.Format("2005-02-15 15:04:05")
	fmt.Println("7->", "hora es formato perosonalisado es ", formato_personalisado)

	//ANALISIS-------------------------------------------------------

	// Definir una cadena de texto que representa una fecha y hora en formato RFC3339
	cadenaRFC3339 := "2023-10-04T10:30:00Z"

	// Analizar la cadena en un objeto time.Time
	fechaHora, err := time.Parse(time.RFC3339, cadenaRFC3339)
	if err != nil {
		fmt.Println("8->", "Error al analizar la cadena:", err)
		return
	}

	// Imprimir la fecha y hora analizada
	fmt.Println("8->", "Fecha y hora analizada:", fechaHora)

}

//temporisadores

func Tempo_risdores() {
	//control  de tiempo-------------------------------------------------
	fmt.Println("9->", "ejecuatando...")
	time.Sleep(2 * time.Second)
	fmt.Println("9->", "ejecuata ... despues de 2 segundos")

	//temporisador de tiempo con un channel o canal------------------------
	fmt.Println("10->", "Comenzando temporizador...")

	// Obtener un canal que enviará un valor después de 3 segundos
	duracion := 2 * time.Second
	temporizador := time.After(duracion)
	// Esperar hasta que se reciba el valor del canal
	<-temporizador
	fmt.Println("10->", "Temporizador completado después de 2 segundos...")
}

//zona oraria

func Zona_Horaria() {

	hora := time.Date(2022, time.December, 14, 11, 0, 0, 0, time.UTC)

	//cargar zona horari de un zona en epesifico

	hora_armerica, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	fmt.Println("11->", "hora en new york :", hora_armerica)

	// Cambiar la zona horaria de la hora UTC a la zona horaria de Nueva York
	horaNyc := hora.In(hora_armerica)

	fmt.Println("11->", "hora en new york :", horaNyc)
}

// obtener el dia de la semana actual

func Dia_semana_actual() {
	dia_semana := time.Now().Weekday()
	fmt.Println("12->", dia_semana)
}

//Comprueba si una fecha es anterior o posterior a otra.
//Comprueba si un objeto time.Time representa una fecha y hora válida.

func Comprobacion() {
	//creando fechas
	fecha1 := time.Date(2023, time.October, 5, 12, 0, 0, 0, time.UTC)
	fecha2 := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
	//viendo fecha
	fmt.Println("13->", fecha1, "|", fecha2)

	//verificando si fecha1 es anterior a fecha2
	ok := fecha1.Before(fecha2)
	if ok {
		fmt.Println("13->", "la fecha1 es anterior")
	} else {
		fmt.Println("13->", "la fecha1 no es anterior")
	}
	//verificando si fecha1 es posterio a fecha2
	ok = fecha1.After(fecha2)
	if ok {
		fmt.Println("13->", "la fecha1 es posterior")
	} else {
		fmt.Println("13->", "la fecha1 no es posterior")
	}

}
