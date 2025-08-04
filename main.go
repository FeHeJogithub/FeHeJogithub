package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log" //indica un registro con fecha y hora
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//fmt.Println("Mysql y go con fehejo")

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//handlers.DeleteContact(db, 5)
	//Crear una instancia de Contact
	/*newContact := models.Contact{
		Id:    4,
		Name:  "Fehejo",
		Email: "fhernandez1@hotmail.com",
		Phone: "123456789",
	}*/

	//handlers.UpdateContact(db, newContact)
	//handlers.CreateContact(db, newContact)
	//handlers.ListContacts(db)

	//handlers.GetContactByID(db, 3)

	for {
		fmt.Println("\nMenu: ")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Println("Seleccione una opcion: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.ListContacts(db)

		case 2:
			fmt.Print("Ingrese el ID del contacto: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactByID(db, idContact)
		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
			handlers.ListContacts(db)
		case 4:
			updateContact := inputContactDetails(option)
			handlers.UpdateContact(db, updateContact)
			handlers.ListContacts(db)
		case 5:
			fmt.Println("Ingrese el ID del contacto que quiere eliminar: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion no valida. por favor, selecciones una opcion valida ")

		}

	}

}

func inputContactDetails(option int) models.Contact {

	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4 {
		fmt.Print("Ingrese el ID del contacto que quiere editar: ")
		var idContact int
		fmt.Scanln(&idContact)

		contact.Id = idContact
	}

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el telefono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact

}
