## Invoice System Endpoints

This Go Invoice System provides a RESTful API to manage invoices and customers efficiently. The API allows you to perform various operations, including creating invoices, managing customers, and generating PDF invoices.

---

### Tech Stack

- **[Go](https://go.dev/):** The primary programming language for the application.  
- **[Chi Router](https://github.com/go-chi/chi):** A lightweight and idiomatic router for building HTTP services in Go.  
- **[Ent](https://entgo.io/):** An ORM for Go, providing powerful and easy-to-use database interaction tools.  
- **[PostgreSQL with Neon](https://neon.tech/):** A serverless database built for modern applications.  
- **[GoPDF Library](https://github.com/signintech/gopdf):** A library for generating PDF documents in Go.  

---

### How to Get Started

1. Clone the repository:  
   ```bash
   git clone https://github.com/Tabintel/invoice-system.git
   cd invoice-system
   ```

2. Install dependencies:  
   ```bash
   go mod tidy
   ```

3. Run the application:  
   ```bash
   go run cmd/api/main.go
   ```

--- 

Base URL: http://localhost:8080

Here is the list of 11 core endpoints, along with their payloads, responses, and usage.




### Health Check
- **Method:** GET
- **URL:** `http://localhost:8080/health`
- **Description:** Verify that the server is running.
- **Response:**
  ```json
  {
    "status": "success",
    "message": "Server is running"
  }
  ```

-------------

### 1. Create Invoice
- **Method:** POST
- **URL:** `http://localhost:8080/api/invoices`
- **Payload:**
  ```json
  {
    "customer_id": 1,
    "due_date": "2024-12-27T00:00:00Z",
    "currency": "USD",
    "items": [
      {
        "description": "Consulting Services",
        "quantity": 1,
        "rate": 1000.00
      }
    ]
  }
  ```
- **Response:**
  ```json
  {
    "data": {
      "id": 1,
      "reference_number": "INV-20241127-7400",
      "total_amount": 1000,
      "status": "draft",
      "issue_date": "2024-11-27T18:08:30.2707774+01:00",
      "due_date": "2024-12-27T00:00:00Z",
      "currency": "USD",
      "created_at": "2024-11-27T18:08:30.2707774+01:00",
      "edges": {}
    },
    "status": "success"
  }
  ```
- **Description:** Create a new invoice.

### 2. List Invoices
- **Method:** GET
- **URL:** `http://localhost:8080/api/invoices?status=draft`
- **Response:**
  ```json
  {
    "data": {
      "invoices": [
        {
          "id": 1,
          "reference_number": "INV-20241127-7400",
          "total_amount": 1000,
          "status": "draft",
          "issue_date": "2024-11-27T17:08:30.270777Z",
          "due_date": "2024-12-27T00:00:00Z",
          "currency": "USD",
          "created_at": "2024-11-27T17:08:30.270777Z",
          "edges": {}
        }
      ],
      "stats": {
        "total_paid": 0,
        "total_overdue": 0,
        "total_draft": 1,
        "total_unpaid": 0
      }
    },
    "status": "success"
  }
  ```
- **Description:** List all invoices with a specified status.

### 3. Update Invoice Status
- **Method:** PUT
- **URL:** `http://localhost:8080/api/invoices/1/status`
- **Payload:**
  ```json
  {
    "status": "paid"
  }
  ```
- **Response:**
  ```json
  {
    "data": {
      "id": 1,
      "reference_number": "INV-20241127-7400",
      "total_amount": 1000,
      "status": "paid",
      "issue_date": "2024-11-27T17:08:30.270777Z",
      "due_date": "2024-12-27T00:00:00Z",
      "currency": "USD",
      "created_at": "2024-11-27T17:08:30.270777Z",
      "edges": {}
    },
    "status": "success"
  }
  ```
- **Description:** Update the status of an invoice.

### 4. Create Customer
- **Method:** POST
- **URL:** `http://localhost:8080/api/customers`
- **Payload:**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890",
    "address": "123 Business Street, City"
  }
  ```
- **Response:**
  ```json
  {
    "data": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "phone": "+1234567890",
      "address": "123 Business Street, City",
      "created_at": "2024-11-28T10:33:27.7360212+01:00"
    },
    "status": "success"
  }
  ```
- **Description:** Create a new customer.

### 5. List Customers
- **Method:** GET
- **URL:** `http://localhost:8080/api/customers`
- **Response:**
  ```json
  {
    "data": [
      {
        "id": 2,
        "name": "Mary Doe",
        "email": "mary@example.com",
        "phone": "+12345678200",
        "address": "100 Bank Street, City",
        "created_at": "2024-11-28T10:04:20.716112Z"
      },
      {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com",
        "phone": "+1234567890",
        "address": "123 Business Street, City",
        "created_at": "2024-11-28T09:33:27.736021Z"
      }
    ],
    "status": "success"
  }
  ```
- **Description:** List all customers.

### 6. Get Single Customer
- **Method:** GET
- **URL:** `http://localhost:8080/api/customers/1`
- **Response:**
  ```json
  {
    "data": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "phone": "+1234567890",
      "address": "123 Business Street, City",
      "created_at": "2024-11-28T09:33:27.736021Z"
    },
    "status": "success"
  }
  ```
- **Description:** Get details of a single customer.

### 7. Update Customer
- **Method:** PUT
- **URL:** `http://localhost:8080/api/customers/1`
- **Payload:**
  ```json
  {
    "name": "John Doe Updated",
    "email": "john.updated@example.com",
    "phone": "+1234567890",
    "address": "456 New Street, City"
  }
  ```
- **Response:**
  ```json
  {
    "data": {
      "id": 1,
      "name": "John Doe Updated",
      "email": "john.updated@example.com",
      "phone": "+1234567890",
      "address": "456 New Street, City",
      "created_at": "2024-11-28T09:33:27.736021Z"
    },
    "status": "success"
  }
  ```
- **Description:** Update a customer's details.

### 8. Delete Customer
- **Method:** DELETE
- **URL:** `http://localhost:8080/api/customers/2`
- **Response:**
  ```json
  {
    "message": "Customer deleted successfully",
    "status": "success"
  }
  ```
- **Description:** Delete a customer.

### 9. Generate Invoice PDF
- **Method:** GET
- **URL:** `http://localhost:8080/api/invoices/1/pdf`
- **Description:** Generate a PDF for an invoice. The browser will automatically download or display the PDF.

### 10. Generate Shareable Invoice Link
- **Method:** POST
- **URL:** `http://localhost:8080/api/invoices/1/share`
- **Response:**
  ```json
  {
    "data": {
      "url": "http://localhost:8080/public/invoices/BFlyy-nWqNyK_t4uUqlVjzFNztojHff4mpSYQd8QEWU=",
      "expires_at": "2024-12-05T14:06:00.8490868+01:00"
    },
    "status": "success"
  }
  ```
- **Description:** Generate a shareable link for an invoice that can be shared with customers.

### 11. Invoice Statistics
- **Method:** GET
- **URL:** `http://localhost:8080/api/invoices/stats`
- **Response:**
  ```json
  {
    "data": {
      "total_paid": 1,
      "total_overdue": 0,
      "total_draft": 2,
      "total_unpaid": 0
    },
    "status": "success"
  }
  ```
- **Description:** Get an overview of the invoice system's current state, including counts of paid, overdue, draft, and unpaid invoices.

---------