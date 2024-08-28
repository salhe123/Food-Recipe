package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/machinebox/graphql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	log.Println("Starting server on port 8080...")
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	query := `
    mutation SignupUser($email: String!, $username: String!, $password: String!) {
        insert_users(objects: {email: $email, username: $username, password: $password}) {
            affected_rows
            returning {
                id
                email
                username
                created_at
            }
        }
    }`
	variables := map[string]interface{}{
		"email":    user.Email,
		"username": user.Username,
		"password": string(hashedPassword),
	}

	result, err := executeGraphQLQuery(query, variables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `
    query GetUserByEmail($email: String!) {
        users(where: {email: {_eq: $email}}) {
            id
            email
            username
            password
        }
    }`
	variables := map[string]interface{}{
		"email": user.Email,
	}

	result, err := executeGraphQLQuery(query, variables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbUser := result["users"].([]interface{})[0].(map[string]interface{})
	if bcrypt.CompareHashAndPassword([]byte(dbUser["password"].(string)), []byte(user.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(dbUser)
}

func executeGraphQLQuery(query string, variables map[string]interface{}) (map[string]interface{}, error) {
	client := graphql.NewClient("https://food-recipe-21.hasura.app/v1/graphql") // Replace with your Hasura endpoint

	req := graphql.NewRequest(query)

	for key, value := range variables {
		req.Var(key, value)
	}

	req.Header.Set("X-Hasura-Admin-Secret", "sE6xQ4q5XKJYnxqn4qrgFRbe6hnE12awTmHIo5kEnUCkJR6SJEZa8d17wTsF6ifr") // Replace with your Hasura admin secret

	var responseData map[string]interface{}

	err := client.Run(context.Background(), req, &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
