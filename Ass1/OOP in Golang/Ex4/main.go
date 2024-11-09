package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Quantity int     `json:"quantity"`
}

// encode
func encodeProduct(product Product) (string, error) {
    // Convert Product struct to JSON
    productJSON, err := json.Marshal(product)
    if err != nil {
        return "", err
    }
    return string(productJSON), nil
}

// decode
func decodeProduct(productJSON string) (Product, error) {
    var product Product
    // Convert JSON string to Product struct
    err := json.Unmarshal([]byte(productJSON), &product)
    if err != nil {
        return product, err
    }
    return product, nil
}

func main() {
    product := Product{
        Name:     "Laptop",
        Price:    1200.50,
        Quantity: 5,
    }

    productJSON, err := encodeProduct(product)
    if err != nil {
        fmt.Println("Error encoding product:", err)
        return
    }
    fmt.Println("Encoded JSON:", productJSON)

    decodedProduct, err := decodeProduct(productJSON)
    if err != nil {
        fmt.Println("Error decoding product:", err)
        return
    }
    fmt.Println("Decoded Product struct:", decodedProduct)
}
