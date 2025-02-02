# Public API for HNG12 Task Zero(0)

## 📌 Project Description

This is a public API that returns specific information in JSON format. The API is built as part of the HNG12 Task Zero(0) Internship and meets the following requirements:

- Returns the registered email address.
- Returns the current date and time in ISO 8601 format (UTC).
- Provides the GitHub URL of the project's codebase.
- Ensures proper CORS handling.
- Deployed to a publicly accessible endpoint.

## 🚀 Technology Stack

This API is implemented using **Go** and hosted on a **.tech** domain.

## 🔗 Live API Endpoint

**Base URL:** `http://hng-task0.randomallstudio.tech`

## 📡 API Specification

### **Endpoint:**

```plaintext
GET / http://hng-task0.randomallstudio.tech/api/v1/task0
```

### **Response Format (200 OK):**

```json
{
	"email": "johnson.oragui@gmail.com",
	"current_datetime": "2025-01-30T09:30:00Z",
	"github_url": "https://github.com/johnson-oragui/hng12"
}
```

## 🛠 Setup Instructions

### **1️⃣ Clone the Repository**

```sh
git clone https://github.com/johnson-oragui/hng12
cd TASK0
```

### **2️⃣ Install Dependencies**

```sh
go mod tidy
```

### **3️⃣ Run the API Locally**

```sh
go run main.go
```

### **4️⃣ Test the API**

Open a browser or use cURL/Postman to send a GET request:

```sh
curl -X GET http://localhost:7001/api/v1/task0
```

## 🌍 Deployment

This API is deployed on a **.tech** domain using:

- **AWS EC2**

## ✅ Acceptance Criteria

- The API must accept GET requests and return the required JSON response.
- The `current_datetime` must be dynamically generated in UTC.
- CORS must be correctly handled.
- The API must have a response time under 500ms.

## 📖 Additional Resources

- [HNG Golang Developers](https://hng.tech/hire/golang-developers)

## 📝 Submission Instructions

1. Ensure your API is live and accessible.
2. Submit the hosted API URL via the designated submission form.
3. Verify that all requirements are met before submission.

⏳ **Deadline:** Friday, Jan 31st - 11:59pm GMT +1 (WAT)

---

👨‍💻 Built for HNG12 Internship. Happy Coding! 🚀
