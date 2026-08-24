package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

func ListContacts(db *sql.DB) {

	query := "SELECT * FROM contact"

	rows, err := db.Query(query)
	if err != nil {

		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\nLista de contactos: ")
	fmt.Println("---------------------------------------------")
	for rows.Next() {

		contact := models.Contact{}

		err := rows.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID:%d, Nombre:%s, Email:%s, Telefono:%s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("----------------------------------------------------")
	}
}
func GetContactByID(db *sql.DB, contactID int) {

	query := "SELECT * FROM contact WHERE id = ? "

	row := db.QueryRow(query, contactID)

	contact := models.Contact{}
	//var valueEmail sql.NullString

	err := row.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("no se econtro ningun contacto con el ID %d", contactID)
		}
		//log.Fatal(err)
	}

	fmt.Println("\nLista de un contacto: ")
	fmt.Println("---------------------------------------------")

	fmt.Printf("ID:%d, Nombre:%s, Email:%s, Telefono:%s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("----------------------------------------------------")
}

func CreateContact(db *sql.DB, contact models.Contact) {

	query := "INSERT INTO contact(name, email, phone) VALUES (?,?,?)"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Nuevo contacto registrado con exito")
}

// UpdateContct actualiza un contacto existente en la base de datos
func UpdateContact(db *sql.DB, contact models.Contact) {

	//Sentencia sql para actualizar un contacto
	query := "UPDATE contact SET name= ? , email= ? , phone=? WHERE id =?"

	//Ejecutar la sentencia sql
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto actualizado con exito")
}

func DeleteContact(db *sql.DB, contactID int) {
	query := "DELETE FROM contact WHERE id = ?"

	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto eliminado con exito")
}
