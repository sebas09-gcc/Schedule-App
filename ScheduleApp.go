package main

import( 
  "fmt"
)
//GRAFOS DIRIGIDOS Y NO VALORADOS
//Son las estructuras que van a permitir definir el grafo
// estado sera 1 cuando se ha aprobado el curso y se pone cero cuando no se ha aprobado aun
type nodoGrafo struct{
	nivel int
  codigo string
  nombre string
  credito float64
  estado int  
	aristaEntrada []string 
	aristaSalida []string
}

type Grafo struct{
	Node []*nodoGrafo
}

//CREA EL GRAFO VACÍO
func CrearGrafo() *Grafo{
	return &Grafo{
		Node: []*nodoGrafo{},
	}
}
//AÑADE UN NODO AL GRAFO EL CUAL NO PRESENTA ARISTAS DE ENTRADA O SALIDA
func (g *Grafo) AddNode(codigo string,nivel int,nombre string,credito float64){
	
	g.Node = append(g.Node,&nodoGrafo{
		nivel: nivel,
    codigo: codigo,
    nombre: nombre,
    credito: credito,
    estado : 0,
		aristaEntrada: make([]string,0),
		aristaSalida: make([]string,0),
	})
	return 
}

//AÑADE UNA ARISTA AL GRAFO PARTIENDO DEL NODO n1 HACIA EL NODO n2 EN EL SENTIDO MENCIONADO
func(g *Grafo)AddEdge(n1 string,n2 string)(flag1,flag2 int){	
	flag1 = 0
	flag2 = 0

	for i := range g.Node{	
			if(g.Node[i].codigo == n1){
				flag1 = 1
			}else if(g.Node[i].codigo == n2){
				flag2 =1
			}
	}
	if(flag1==0 || flag2==0 ){
		panic("El nodo no se encuentra\n")
	}
	for i := range g.Node{
		
			if(g.Node[i].codigo == n1){
        g.Node[i].aristaSalida = append(g.Node[i].aristaSalida, n2)
			}else if(g.Node[i].codigo == n2){
        g.Node[i].aristaEntrada = append(g.Node[i].aristaEntrada, n1)
			}
	}
	return
}
//Esta funcion devuelve una lista de los cursos requesitos del curso del cual se hace la consulta
func(g*Grafo)NodosQueEntran(n1 string)[]string{
  
	for i := range g.Node{
		if(g.Node[i].codigo == n1){
			nodosEntrada := make([]string,len(g.Node[i].aristaEntrada))
			copy(nodosEntrada,g.Node[i].aristaEntrada)

			return nodosEntrada
		}
	}
	panic("El nodo no se encuentra\n")
}
// Esta funcion devuelve una lista con los cursos a los cuales puede acceder despues de aprobar el curso consultado
func(g*Grafo)NodosQueSalen(n1 string)[]string{

	for i := range g.Node{
		if(g.Node[i].codigo == n1){
			nodosSalida := make([]string,len(g.Node[i].aristaSalida))

			copy(nodosSalida,g.Node[i].aristaSalida)

			return nodosSalida
		}
	}
	panic("El nodo no se encuentra\n")
}
//La funcion "imprimirNodos" se encarga de imprimir todos los nodos del grafo 
func(g*Grafo)imprimirNodos(){
  var flag int 
  flag=0
	for i := range g.Node{
    credito := fmt.Sprint(g.Node[i].credito)
    if(flag==0){
      fmt.Println("------------Nivel",g.Node[i].nivel,"------------")
      flag=1
    }
		fmt.Println(g.Node[i].codigo, g.Node[i].nivel,g.Node[i].nombre,credito)
    if(g.Node[i].codigo=="1LIN15" || g.Node[i].codigo=="1PSI02"||g.Node[i].codigo=="INF134"){
      flag=0
    }
  }
}
// La funcion actualizarCursosJalados recibe  los parametros de nivel y lista de codigos de los cursos jalados, actualiza con los estados desaprobados
func (g *Grafo)actualizarCursosJalados(nivel int,list[25]string){
  var j int = 0

  for i := range g.Node{
    if(g.Node[i].nivel <= nivel){
      if(g.Node[i].codigo == list[j]){
        g.Node[i].estado=0
        j = j + 1
      }else{
        g.Node[i].estado=1  
      } 
    } 
  }
}

func (g*Grafo)actualizarEstadosAprobados(nivel int){
  for i := range g.Node{
    if(g.Node[i].nivel <= nivel){
      g.Node[i].estado=1
    }
  }
}
//FUNCION A: CONSULTA DE UN CURSO QUE DESEA LLEVAR EL SIGUIENTE CICLO
func (grafo1 *Grafo) a(codigo string){
  var codCurso []string
  var nivel int
  //aqui vas a buscar el curso que solicitaste
  for i := range grafo1.Node{
    if(grafo1.Node[i].codigo == codigo){
      if(grafo1.Node[i].estado == 1){
        fmt.Println("El curso ya fue aprobado")
        return 
      }
      nivel=grafo1.Node[i].nivel
      codCurso = grafo1.NodosQueEntran(codigo)
      break
    }  
  }
  if(len(codCurso)==0){
    fmt.Println("Si puede llevar el curso solicitado")
    return
  }
  //aqui vas s buscar los cursos que son requisitos    
  k:=0
  validez := 0
  for i:=0; i < 25 ;i++{
    if(len(codCurso)==k){
      break
    }
    //la variable nivel guarda el nivel del CURSO SOLICITADO
    //codCurso es la variable que almacena los requisitos
    //por nde grafo1.Node[i].nivel tiene el nivel de los cursoRequisitos
    if(grafo1.Node[i].codigo == codCurso[k]){
      if(grafo1.Node[i].nivel == nivel){ //LLEVAR EL CURSO SIMULTANEO
          validez++ 
      }else{
        if(grafo1.Node[i].estado == 1 || grafo1.Node[i].codigo == "1MAT04"){
        validez++
        }
      }
      i=0
      k++
    }
  }
  if(validez == len(codCurso)){
    fmt.Println("Si puede llevar el curso solicitado")
  }else{
    fmt.Println("No cumple con los requisitos para llevar ese curso")
  }
}
//FUNCIÓN B: VERIFICA SI EL ALUMNO HA EGRESADO DE GENERALES O NO
func (grafo1 *Grafo) b(){
  var CREDITOSGENERALES float64 = 84.25//VARIABLE CONSTANTE
	var creditos float64 = 0.0

	for i := range grafo1.Node {
		if grafo1.Node[i].estado == 1 {
			creditos = grafo1.Node[i].credito + creditos
		}
	}
  if(creditos==CREDITOSGENERALES){
    fmt.Println("USTED TIENE LA CONDICION DE EGRESO APROBADA")
  }else{
    fmt.Println("USTED AÚN NO HA EGRESADO DE GENERALES... APRUEBE TP PRIMERO")
  }
}
//FUNCIÓN C: VERFICIA SI EL ALUMNO SIGUE PERTENECIENDO A LA UNIVESIDAD O NO
func (grafo1 *Grafo) c(cantJala int ){
  if(cantJala<=10){
    fmt.Println("USTED TIENE CONDICION DE ALUMNO")
  }else{
    fmt.Println("USTED ESTA ELIMINADO")
  }
}
//FUNCIÓN D: DETERMINA CUANTOS CREDITOS APROBADOS POSEE EL ALUMNO HASTA EL MOMENTO
func (grafo1 *Grafo) d()float64{
  var creditos float64 = 0
	for i := range grafo1.Node {
		if grafo1.Node[i].estado == 1 {
			creditos = grafo1.Node[i].credito + creditos
		}
	}
	return creditos
}
//FUNCIÓN E: MUESTRA LOS CURSOS FALTANTES QUE AÚN TIENE PENDIENTE
func (grafo1 *Grafo) e(){
  cursosFaltantes := make([]string,0)
  cont := 0
  for i := range grafo1.Node{
    if grafo1.Node[i].estado == 0{
      cursosFaltantes = append(cursosFaltantes,grafo1.Node[i].codigo)
      cursosFaltantes = append(cursosFaltantes,grafo1.Node[i].nombre)
      cont++
    }
  }
  cont=cont*2
  for i:=0 ; i<cont;i=i+2{
    fmt.Println(cursosFaltantes[i],cursosFaltantes[i+1])
  }
}
//FUNCIÓN F: MUESTRA SI ERES DIGNO DE GUANIRA
func (grafo1 *Grafo) f(){
  fmt.Println("///////////(////////////////////*******/////////////////*///////****////////////")
  fmt.Println("//////////////////////////////////******///////////////////////*****////////////")
  fmt.Println("(((////((/////////////*****/*/////*******////////////////////**************/////")
  fmt.Println("(////((((((((////////********************///////(//////////////////******///////") 
  fmt.Println("//////((((((((/////****////**********,,,..,*****///////////////////***//////////")
  fmt.Println("////(((((((((((((/////////////******,,,,,,,...,*****////////////////////////////")
  fmt.Println("((((((((((((((((//////////////((((((((/////*********,,,*////////////////////***/")
  fmt.Println("(((((((((//////////////////((((((((//////******///*////*,*////////////////****//")
  fmt.Println("////((((/(((//(/////////((####((((((//(######(//(((#(((/,**((///////////***//***")
  fmt.Println("//////////////////////(#####%###(((((((((////((#%%##((##(#(/(#////////**********")
  fmt.Println("/////////////////////(#%%#############%%%#####(((/////(/###((#(////////*********")
  fmt.Println("///////***///////////(%%%#((/********************//////(/##(###////////*********")
  fmt.Println("/////////////////////((#(((//**/******************//////((##(%%(//////////*/////")
  fmt.Println("/////////////////////(((((//////(#%&%#((//**/////((///(((#%%#%%(//////**////////")
  fmt.Println("/////////////////////((((////(((#%%%##((//*/(###%%###((((#&&%%#/////////////////")
  fmt.Println("///////((////////////////////////(((((//****(##(%&@&#####%&%%#//((////*****/////")
  fmt.Println("/(((((/////////////////(////******//*******/((/////(###(##%#(//////////**///////")
  fmt.Println("(///(((((((((((////////*/////***************//(/**/////((##((((((////**/////////")
  fmt.Println("((((((((((((((((((((///////********/(/////*/((#/***///((##(/////////////////////")
  fmt.Println("(((((((((((((((((/////*///////////(/*///(#((((((//*//((##(////*////////////////(")
  fmt.Println("((((((((((((((((///////////////((/******///////(((/(((##(/////*///////////////((")
  fmt.Println("(((((((((((((//////////,*//////**///(((((####%#((/(((##(/////***////////////((((")
  fmt.Println("(((##((((((///////////..,(////////*//////(((((/////(##(////////////////////(((((")
  fmt.Println("(((((((((((/////////*....,#////////*****////////(((##(/////*****/////////(((((((")
  fmt.Println("((((((((((//////,. ........*((((////****////(((####/*///////////////**///(((((((")
  fmt.Println("//(((((((//*.     ...........*(((((((((((((((####(#(/,,*/////////////((((((###((")
  fmt.Println("/(((((/,       .  .............,/((((((########(((#//*,..  ./////(///((((((((((#")
  fmt.Println("((/,              ................./((/((####(((##(//**.,,.     ,/((((((((((((((")
  fmt.Println(",               .................. ,,/#((##((((((#*,*,,.  ..  .     ,/((((((((((")
  fmt.Println("              ..  .............. ,*..,#(((((((///(,,.,...  .,....      *((((((((")
  fmt.Println(".      ..........  .............,.,,./((////////((,,/*..... .......      *#(((((")
  fmt.Println(" ...    .......................,(,..,((//////////,***,,..... ........ ... ,(((((")
  fmt.Println("   ..... .......................,,/,,/(((((///(*,*,,,,,...... ..........  .,((((")
  fmt.Println(".................................,.,**,*(/(/((,,,,.,..,.................,...,(##")
  fmt.Println("...................................,,,,,,*(/#*,....,..,................,,.,,,,##")
  fmt.Println(".........,..........................,,..,,,(/,....,,...................,*,,,,,./")
  fmt.Println("..........,.,........................,,..,.,,,...,,...................,,*,,,*,,,")
  fmt.Println("........,,,,,.........................,...,,....,,................,,,..,*,,,*,,,")
  fmt.Println("........,,,,,....,,....................,...,...,,................,,,,,,**,,**,,,")
}
//FUNCIÓN G: MUESTRA SI EL ALUMNO PUEDE HIBRIDAR O NO
func (grafo1 *Grafo) g(){
  var CREDITOSGENERALES float64 = 84.25
  var creditos float64 = 0.0

  for i := range grafo1.Node{
    if(grafo1.Node[i].estado == 1){
      creditos = grafo1.Node[i].credito + creditos
    }
  }
  if(CREDITOSGENERALES - creditos <= 12){
    fmt.Println("Si puede hibridar el ciclo entrante")
  }else{
    fmt.Println("Aún no cumple con la cantidad mínima de créditos aprobados para poder hibridar");
  }
  return;
}
//DEVUELVE EL CURSO QUE SE BUSCA 
func (g*Grafo) buscarCurso(codigo string)*nodoGrafo{
  i:=0
  for i:=0;i<25;i++ {
    if(g.Node[i].codigo == codigo){
      break
    }
  }
  fmt.Println(g.Node[i].codigo)
  return g.Node[i]
}
func (g*Grafo) eliminarRepetidos(cursosFaltantes []string)[]string{
  for i:=0 ; i<len(cursosFaltantes);i++{
    for j:=i+1 ; j<len(cursosFaltantes);j++{
      if(cursosFaltantes[i] == cursosFaltantes[j]){
        cursosFaltantes[j]="0"
      }
    }
  }
  cursos :=make([]string,0);
  for i:=0 ; i<len(cursosFaltantes);i++{
    if(cursosFaltantes[i]!="0"){
      cursos = append(cursos,cursosFaltantes[i])
    }  
  }
  return cursos;
}  
//FUNCIÓN H: MUESTRA LOS CURSOS QUE PUEDO LLEVAR EL SIGUIENTE CICLO 
func (grafo1 *Grafo) h(nivelmax int){

  cursosFaltantes := make([]string,0)
  var codCursos []string
  var cursosJalados [40]string
  var cursosAprobados [40]string
  //ESTA SECCION ME BOTA TODOS LOS CURSOS JALADOS
  t := 0
  for i:= range grafo1.Node{
    if(grafo1.Node[i].nivel <= nivelmax && grafo1.Node[i].estado == 0){
      cursosJalados[t]=grafo1.Node[i].codigo
      t++;
    }
  }
  //ESTA SECCION ME BOTA TODOS LOS CURSOS APROBADOS
  k:=0
  for i:= range grafo1.Node{
    if(grafo1.Node[i].estado == 1){
      cursosAprobados[k]=grafo1.Node[i].codigo;
      k++;
    }
  }

  j := 0
  for i:=0; i<25;i++{//BUSCARÉ TODOS LOS POSIBLES CURSOS QUE PUEDE LLEVAR DESPUES DE APROBAR
    if(grafo1.Node[i].codigo == cursosAprobados[j]){
      codCursos = grafo1.NodosQueSalen(grafo1.Node[i].codigo)//BUSCARÉ TODOS LOS POSIBLES CURSOS QUE PUEDE LLEVAR DESPUES DE APROBAR
      cursosFaltantes = append(cursosFaltantes,codCursos...)//esto concatena todos los cursos que falta por llevar...
      j++;
    }  
  }
  //AÑADIR LOS CURSOS QUE NO TIENEN REQUISITOS 
  var longitud int
  for a:=0; a<25;a++{
    longitud = len(grafo1.NodosQueEntran(grafo1.Node[a].codigo))
    if( longitud == 0 ){
      cursosFaltantes = append(cursosFaltantes,grafo1.Node[a].codigo)
    }
  }
  //TODOS LOS FILTROS
  for m := range cursosJalados{
    if(cursosJalados[m]=="1MAT05"){
      for i := range cursosFaltantes{
        if(cursosFaltantes[i]=="1MAT06" || cursosFaltantes[i]=="1MAT07"){
          cursosFaltantes[i]="0"
        }
      }
    }
    if(cursosJalados[m]=="1MAT04" || cursosJalados[m]=="1MAT06"){
      for i:= range cursosFaltantes{
        if(cursosFaltantes[i]=="1MAT07"){
          cursosFaltantes[i]="0"
        }
      }
    }
    if(cursosJalados[m]=="1FIS03" || cursosJalados[m]=="1FIS02"){
      for i:= range cursosFaltantes{
        if(cursosFaltantes[i]=="1FIS04" || cursosFaltantes[i]=="1FIS05"){
          cursosFaltantes[i]="0"
        }
      }
    }
  }
  if(nivelmax==1){
    for i:= range cursosFaltantes{
      if(cursosFaltantes[i]=="1MAT07"){
        cursosFaltantes[i]="0"
      }
    }
  }
  if(nivelmax==2){
    for i:= range cursosFaltantes{
      if(cursosFaltantes[i]=="1MAT07"){
        cursosFaltantes = append(cursosFaltantes,"1MAT08","INF134")
      }
      if(cursosFaltantes[i]=="1FIS04"){
        cursosFaltantes = append(cursosFaltantes,"1FIS05")
      }
    }
  }
  if(nivelmax==3){
    for i:=range cursosFaltantes{
      if(cursosFaltantes[i]=="1FIS06"){
        cursosFaltantes = append(cursosFaltantes,"1FIS07")
      }

    }
  }
  //eliminarRepetidos
  cursosActuales := make([]string,0)
  cursosActuales=grafo1.eliminarRepetidos(cursosFaltantes);
  for n := range cursosActuales{
    //curso=grafo1.buscarCurso(cursosActuales[n])
    for i:=0;i<25;i++ {
      if(grafo1.Node[i].codigo == cursosActuales[n]){
        if(grafo1.Node[i].estado == 0){
          fmt.Println(grafo1.Node[i].nivel,grafo1.Node[i].codigo,grafo1.Node[i].nombre)
        }
      }
    }
  } 
}   
//FUNCIÓN I: MUESTRA LOS CURSOS APROBADOS HASTA EL MOMENTO 
func (grafo1 *Grafo) i(){
  for i := range grafo1.Node{
    if(grafo1.Node[i].estado == 1){
      fmt.Println(grafo1.Node[i].nivel,grafo1.Node[i].codigo,grafo1.Node[i].nombre)
    }
  }
  return;
}
//AÑADE LOS CURSOS A LA ESTRUCTURA DECLARADA AL PRINCIPIO
func (grafo1 *Grafo) AddCourses(){
  // First Cicle
  grafo1.AddNode("1MAT04",1,"ALGEBRA MATRICIAL Y GEOMETRIA ANALITICA",4.50)
  grafo1.AddNode("1MAT05",1,"FUNDAMENTO DE CALCULO",4.50)
  grafo1.AddNode("1FIS01",1,"FUNDAMENTO DE FISICA",3.50)
  grafo1.AddNode("1QUI01",1,"QUIMICA",3.50)
  grafo1.AddNode("1QUI02",1,"LABORATORIO DE QUIMICA",0.75)
  grafo1.AddNode("1LIN15",1,"COMUNICACION ACADEMICA",3.00)

  // Second Cicle
  grafo1.AddNode("1MAT06",2,"CALCULO DIFERENCIAL",4.50)
  grafo1.AddNode("1FIS02",2,"FISICA 1",4.50)
  grafo1.AddNode("1FIS03",2,"LABORATORIO DE FISICA 1",0.50)
  grafo1.AddNode("1ING02",2,"DIBUJO EN INGENIERIA",4.50)
  grafo1.AddNode("1LIN16",2,"TRABAJO ACADEMICO",3.00)
  grafo1.AddNode("1FIL01",2,"CIENCIA Y FILOSOFIA",3.00)
  grafo1.AddNode("1PSI02",2,"MOTIVACION Y LIDERAZGO",2.00)
  
  //THIRD CYCLE
  grafo1.AddNode("1MAT07",3,"CALCULO INTEGRAL",4.50)
  grafo1.AddNode("1MAT08",3,"CALCULO EN VARIAS VARIABLES",4.50)
  grafo1.AddNode("1FIS04",3,"FISICA 2",4.50)
  grafo1.AddNode("1FIS05",3,"LABORATORIO DE FISICA 2",0.50)
  grafo1.AddNode("1INF01",3,"FUNDAMENTOS DE PROGRAMACION",3.00)
  grafo1.AddNode("INF134",3,"ESTRUCTURAS DISCRETAS",4.50)

  //FOURTH CYCLE
  grafo1.AddNode("1MAT09",4,"CALCULO APLICADO",4.50)
  grafo1.AddNode("1FIS06",4,"FISICA 3",4.50)
  grafo1.AddNode("1FIS07",4,"LABORATORIO DE FISICA 3",0.50)
  grafo1.AddNode("INF144",4,"TECNICAS DE PROGRAMACION",5.00)
  grafo1.AddNode("1SOC01",4,"SOCIOLOGIA",3.00)
  grafo1.AddNode("CDR123",4,"PENSAMIENTO CRISTIANO",3.50)

  grafo1.AddEdge("1QUI01", "1QUI02") //QUIMICA --QUEDA

  grafo1.AddEdge("1MAT04", "1MAT06") //CALCULO 1 --QUEDA
	grafo1.AddEdge("1MAT05", "1MAT06") //CALCULO 1 --QUEDA
	grafo1.AddEdge("1FIS01", "1FIS02") //FISICA 1 --QUEDA
  grafo1.AddEdge("1MAT06", "1FIS02") //FISICA 1 --QUEDA
  grafo1.AddEdge("1FIS03", "1FIS02") //FISICA 1 --QUEDA
  grafo1.AddEdge("1FIS02", "1FIS03") //LABFISICA 1 --QUEDA
  grafo1.AddEdge("1MAT04", "1ING02") //DIBUJO --QUEDA
	grafo1.AddEdge("1LIN15", "1LIN16") //TRABAJO ACADEMICO--QUEDA

	grafo1.AddEdge("1MAT04", "1MAT07") //CALCULO 2 --QUEDA
	grafo1.AddEdge("1MAT06", "1MAT07") //CALCULO 2 --QUEDA
  grafo1.AddEdge("1MAT07", "1MAT08") //CALCULO 3 --QUEDA
	grafo1.AddEdge("1FIS02", "1FIS04") //FISICA 2 --QUEDA
	grafo1.AddEdge("1FIS03", "1FIS04") //FISICA 2 --QUEDA
  grafo1.AddEdge("1MAT07", "1FIS04") //FISICA 2 --QUEDA
  grafo1.AddEdge("1FIS05", "1FIS04") //FISICA 2 --QUEDA
  grafo1.AddEdge("1FIS04", "1FIS05") //LABFISICA 2 --QUEDA
  grafo1.AddEdge("1MAT07", "INF134") //ESTRUCTURAS DISCRETAS --QUEDA
	grafo1.AddEdge("1FIS02", "1INF01") //FUNDA PROGRAMACION --QUEDA

	grafo1.AddEdge("1MAT08", "1MAT09") //CALCULO 4---QUEDA
	grafo1.AddEdge("1FIS04", "1FIS06") //FISICA 3--QUEDA
	grafo1.AddEdge("1FIS05", "1FIS06") //FISICA 3--QUEDA
  grafo1.AddEdge("1FIS07", "1FIS06") //FISICA 3--QUEDA
  grafo1.AddEdge("1FIS06", "1FIS07") //LABFISICA 3--QUEDA
	grafo1.AddEdge("1INF01", "INF144") //TECNICAS DE PROGRAMACION--QUEDA
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main(){
  //INFORMACIÓN GUARDADA EN LA BASE DE DATOS DE LA PUCP
	var nivelmax int
  var jalado string
  var cantJala int
  var cantNoLlevado int
  var codigo string
  var list [25]string
  var opcion string

	grafo1 :=CrearGrafo()
	grafo1.AddCourses()

  fmt.Println("----------Bienvenido al programa----------");
  fmt.Println("Malla curricular: INGENIERÍA INFORMATICA")
	grafo1.imprimirNodos()
  //fmt.Println(grafo1.Node[6].aristaEntrada[0])
  fmt.Println()
  //fmt.Println(grafo1.NodosQueEntran("1MAT06"))
  //fmt.Println(len(grafo1.NodosQueEntran("1MAT06")))
	fmt.Println("¿Cual es el nivel maximo de un curso que ha llevado?")
	fmt.Scanf("%d", &nivelmax) 

  // de esta manera tendremos hasta el nivel maximo de curso que llevo, todos los niveles que estan adelante se rellenarian con un 0 de no aprobado, y los que estan dentro del maximo nivel y anteriores se rellenarian con 1's de aprobado
  fmt.Println("¿Ha jalado usted algun curso?(s/n)")
  fmt.Scanf("%s", &jalado)
  if(jalado=="s" || jalado == "S"){
    fmt.Println("¿Cuantos cursos ha jalado y aun no ha aprobado (CONTAR SOLO DESDE EL NIVEL MAXIMO HACIA ATRAS)")
    fmt.Scanf("%d", &cantJala)
    // con esto no es encargamos de saber cuantos cursos a jalado el alumn o sin consultar cuales dentro de los nodos
    fmt.Println("¿Cuantos cursos aun no ha llevado(CONTAR SOLO DESDE EL NIVEL MAXIMO HACIA ATRAS)")
    fmt.Scanf("%d", &cantNoLlevado)
    //tambien nos encargamos de saber cuales se ha salteado y rellenamos con 0's de no aprobados aun 
    fmt.Println("Ingrese el codigo del/los cursos que ha jalado o del que no ha llevado")
    for i := 0; i < cantJala+cantNoLlevado; i++ {
        fmt.Scanf("%s", &codigo)
        list[i]=codigo; //list es una lista que guarda los cursos jalados o no llevados
  	}
    //aqui es donde aquellos codigos de curso que recibimos , los buscamos y cambiamos su estado a  0 de no aprobado, de esta manera tendriamos todos los cursos aprobados y aquellos que se han pateado de la manera mas eficiente y sin pedir demasiado al usuario
    grafo1.actualizarCursosJalados(nivelmax,list)
    fmt.Println("Procesando.......")
  }else{
  	if(jalado!="n" and jalado!="N"){
  		panic("La letra ingresada es incorrecta")
  	}
    grafo1.actualizarEstadosAprobados(nivelmax)
    fmt.Println("Procesando.......")
  } 

  fmt.Println("-------------------OPCIONES A CALCULAR--------------- ")
  fmt.Println("OPCION A: CONSULTA DE UN CURSO QUE DESEA LLEVAR  ")
  fmt.Println("OPCION B: CONDICION DE EGRESO")
  fmt.Println("OPCION C: CONDICION DE ALUMNO")
  fmt.Println("OPCION D: CANTIDAD DE CREDITOS APROBADOS")
  fmt.Println("OPCION E: CURSOS FALTANTES PARA EGRESAR")
  fmt.Println("OPCION F: ¿SOY DIGNO?(¿APROBÓ TP?)")
  fmt.Println("OPCIÓN G: CONSULTA SI PUEDO HIBRIDAR O NO")
  fmt.Println("OPCION H: QUE CURSOS PUEDO LLEVAR EL SIGUIENTE CICLO")
  fmt.Println("OPCION I: CURSOS APROBADOS HASTA EL MOMENTO")
  fmt.Println("¿Que desea calcular?: ")
  fmt.Scanf("%s", &opcion)

  switch os := opcion; os {
	  case "A":
      fmt.Println("Ingrese el curso desea consultar")
      fmt.Scanf("%s",&codigo)
		  grafo1.a(codigo)
	  case "B":
		  grafo1.b()
    case "C":
		  grafo1.c(cantJala)
	  case "D":
		  fmt.Println(grafo1.d())
    case "E":
		  grafo1.e()
	  case "F":
		  grafo1.f()
    case "G":
      grafo1.g()
    case "H":
      grafo1.h(nivelmax)
    case "I":
      grafo1.i()
	  default:
	    panic("NO SE INGRESO LA OPCION CORRECTA")
	}
}
