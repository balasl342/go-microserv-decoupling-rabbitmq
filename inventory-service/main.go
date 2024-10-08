package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection

func init() {
	var err error
	conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
}

func publishInventoryUpdate(productID string, quantity int) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"inventory_updates",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}

	message := fmt.Sprintf("Product %s updated. New quantity: %d", productID, quantity)
	err = ch.Publish(
		"inventory_updates",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
	fmt.Printf(" [x] Sent %s\n", message)
}

func updateInventoryHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("productID")
	quantity := r.URL.Query().Get("quantity")

	if productID == "" || quantity == "" {
		http.Error(w, "Missing productID or quantity", http.StatusBadRequest)
		return
	}

	iQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		http.Error(w, "quantity data error", http.StatusBadRequest)
		return
	}
	publishInventoryUpdate(productID, iQuantity)

	fmt.Fprintf(w, "Inventory updated for product %s with quantity %s!", productID, quantity)
}

func main() {
	http.HandleFunc("/update-inventory", updateInventoryHandler)
	log.Println("Inventory service is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
