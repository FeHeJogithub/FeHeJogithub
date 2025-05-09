

Funcion cacturarNumeros(num)
	Definir i como entero;
	//Definir encontrado Como Logico;
	
	para i=1 Hasta 5 con paso 1 Hacer
		num[i]=AZAR((5)+1);
		
	Fin Para
	
Fin Funcion

funcion imprimirNumeros(num)
	//Definir I Como Entero;
	Escribir "Los numeros son ";
	definir i Como Entero;
	para i=1 Hasta 5 con paso 1 Hacer
		
		Escribir  Sin Saltar num[i], " "; 
		fin para	
FinFuncion

funcion buscarNumeros(num)
	Definir dato ,i Como Entero;
	Definir encontrado Como Logico;
	encontrado=falso;
	Escribir "Ingrese el numero que quiere buscar";
	Escribir " ";
	Leer dato;
	para i=1 Hasta 5 con paso 1 Hacer
		
	
	si dato = num[i] entonces
		Escribir "El numero es ", dato, " esta en la posicion[",i,"]";
	encontrado = verdadero;	
Finsi
fin para
	si encontrado = falso entonces
	
		Escribir "los numeros son invalidos";
	finsi
	
FinFuncion

Algoritmo buscar_arreglo_algoritmo
	
	definir num Como Entero;
	dimension num[5+1];
	Escribir "Bienvenido a la busqueda de arreglos";
	
	
	cacturarNumeros(num);
	
	imprimirNumeros(num);
	
	buscarNumeros(num);
FinAlgoritmo
