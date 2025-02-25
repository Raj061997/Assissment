# Golang Blog CRUD API with Fiber and PostgreSQL

## 🚀 Project Overview
This is a **CRUD API** for a blog application built using **Golang, Fiber, and PostgreSQL**. It allows users to:
- **Create** a blog post
- **Read** all blog posts or a single post
- **Update** an existing blog post
- **Delete** a blog post

It also includes **Swagger documentation** and is deployed on **Render**.

---

## 📌 Features
- ✅ **Go Fiber** as the web framework
- ✅ **PostgreSQL** as the database
- ✅ **Swagger API documentation**
- ✅ **CORS handling** for API accessibility
- ✅ **Deployment on Render**
- ✅ **.env support for environment variables**
- ✅ **Unit tests for API routes** (Bonus)

---

## 🛠 Tech Stack
- **Backend:** Golang (Fiber framework)
- **Database:** PostgreSQL
- **API Docs:** Swagger
- **Deployment:** Render

---

## ⚙️ Installation & Setup

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/yourusername/your-repo.git
cd your-repo
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Set Up Environment Variables
Create a `.env` file in the root directory and add:
```env
PORT=10000
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
```

### 4️⃣ Run Database Migrations
```sh
go run migrate.go
```

### 5️⃣ Start the Server
```sh
go run main.go
```

Server will run at **http://localhost:10000**

---

## 🔗 API Endpoints
| Method | Endpoint | Description |
|--------|-------------|-------------|
| **POST** | `/api/blog-post` | Create a new blog post |
| **GET** | `/api/blog-post` | Get all blog posts |
| **GET** | `/api/blog-post/:id` | Get a single blog post |
| **PATCH** | `/api/blog-post/:id` | Update a blog post |
| **DELETE** | `/api/blog-post/:id` | Delete a blog post |

---

## 📖 Swagger Documentation
Swagger UI is available at:
```
http://localhost:10000/swagger/index.html
```
To regenerate Swagger docs, run:
```sh
swag init
```

---

## 🚀 Deployment on Render
### 1️⃣ Push to GitHub
```sh
git add .
git commit -m "Initial commit"
git push origin main
```

### 2️⃣ Deploy on Render
1. Go to [Render](https://render.com)
2. Create a new **Web Service**
3. Connect your GitHub repository
4. Set **Environment Variables** (`PORT=10000`, `DATABASE_URL=your-db-url`)
5. Deploy 🚀

---
