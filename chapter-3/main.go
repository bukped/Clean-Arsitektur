package main

import (
	"context"
	"go-restapi-mux/models"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Product represents a product in the inventory
type Product struct {
	ID       string
	Name     string
	Price    int
	Quantity int
}

// The list of products in the inventory
var products []Product

func main() {
	// Create a client to connect to the MongoDB database
	models.ConnectDatabase()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	// models.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get a handle for the products collection
	// productsCollection := client.Database("test").Collection("test")

	r := mux.NewRouter()

	// Handle the home page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		title := "Home Page"
		data := map[string]interface{}{
			"Title":   title,
			"Content": template.HTML("{{templates \"content\" .}}"),
		}
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/index.html", "templates/input.html"))
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Handle the about page
	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		title := "About Us"
		data := map[string]interface{}{
			"Title":   title,
			"Content": template.HTML("{{template \"content\" .}}"),
		}
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/about.html"))
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Handle the input product page
	r.HandleFunc("/products/new", func(w http.ResponseWriter, r *http.Request) {
		title := "Add New Product"
		data := map[string]interface{}{
			"Title":   title,
			"Content": template.HTML("{{template \"content\" .}}"),
		}
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/input.html"))
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	// Handle the form submission for adding a new product
	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		// Create a new product
		newProduct := Product{
			ID:       uuid.New().String(),
			Name:     name,
			Price:    price,
			Quantity: quantity,
		}

		// Add the new product to the database
		products = append(products, newProduct)

		// Redirect to the home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}).Methods("POST")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
