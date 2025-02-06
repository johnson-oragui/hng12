# Number Classification API

## 📌 Overview

The **Number Classification API** is a RESTful service that takes a number as input and returns interesting mathematical properties about it, along with a fun fact.

---

## 🚀 Features

- Determines if a number is **prime** or **perfect**.
- Classifies numbers as **Armstrong**, **odd**, or **even**.
- Calculates the **sum of the digits** of the number.
- Fetches a **fun fact** about the number using the [Numbers API](http://numbersapi.com/#42/math).
- Returns structured JSON responses.
- Deployed to a **publicly accessible endpoint**.
- Implements **CORS support**.

---

## 📚 API Specification

### **🔹 Endpoint**

**GET** `http://hng-task1.randommallstudio.tech/api/classify-number?number={number}`

### **✅ Successful Response (200 OK)**

```json
{
	"number": 371,
	"is_prime": false,
	"is_perfect": false,
	"properties": ["armstrong", "odd"],
	"digit_sum": 11,
	"fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

### **❌ Error Response (400 Bad Request)**

```json
{
	"number": "alphabet",
	"error": true
}
```

---

## 🛠️ Technology Stack

- Any programming language or framework of your choice (e.g., Python, Go, Java, JavaScript, PHP, C#)
- Hosted on a **publicly accessible server**
- Uses the **Numbers API** for retrieving fun facts
- Supports **CORS (Cross-Origin Resource Sharing)**

---

## 🔍 Properties Classification

| Property      | Condition                                                                                                                             |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| **Armstrong** | A number is an Armstrong number if the sum of its digits each raised to the power of the number of digits equals the original number. |
| **Odd**       | A number that is not divisible by 2.                                                                                                  |
| **Even**      | A number that is divisible by 2.                                                                                                      |
| **Prime**     | A number greater than 1 with only two divisors: 1 and itself.                                                                         |
| **Perfect**   | A number whose divisors (excluding itself) sum up to the number itself.                                                               |

### **🔢 Property Combinations**

| Condition            | Properties Output       |
| -------------------- | ----------------------- |
| Armstrong & Odd      | `["armstrong", "odd"]`  |
| Armstrong & Even     | `["armstrong", "even"]` |
| Not Armstrong & Odd  | `["odd"]`               |
| Not Armstrong & Even | `["even"]`              |

---

## 🏗️ Deployment Requirements

- Must be deployed to a **publicly accessible URL**.
- The API should be **fast**, with a response time of **< 500ms**.

---

## 📌 Acceptance Criteria

### **Functionality**

- Accepts **GET** requests with a `number` parameter.
- Returns JSON in the specified format.
- Accepts only **valid integers**.
- Provides appropriate HTTP **status codes**.

### **Code Quality**

- Organized **code structure**.
- Implements **error handling** and **input validation**.
- Avoids **hardcoded values**.

### **Documentation**

- Includes a **well-structured README**.

### **Deployment**

- **Publicly accessible API**.
- **Fast response time (<500ms).**

---

## 📂 Version Control & Submission Guidelines

- Code must be hosted on **GitHub**.
- Repository must be **public**.
- Submit your task through the **designated submission form**.
- Ensure you have:
  - Hosted the API on a **public platform**.
  - **Tested** your API thoroughly before submission.
  - Double-checked all **requirements and acceptance criteria**.
  - Reviewed your work to ensure **accuracy and functionality**.

---

## 🗓️ Submission Deadline

**6th February 2025, 11:59 PM WAT (GMT +1)**

Late submissions will not be entertained. Ensure that your API meets all the required conditions before submitting.

---

## 📚 Resources

- **Fun Fact API**: [http://numbersapi.com/#42](http://numbersapi.com/#42)
- **Parity in Mathematics**: [Wikipedia - Parity](<https://en.wikipedia.org/wiki/Parity_(mathematics)>)

---

## ⚡ Contributing

Contributions are welcome! Feel free to open issues or submit pull requests if you have improvements.

---

## 📜 License

This project is licensed under the **MIT License**.

---

🎯 **Happy Coding! 🚀**
