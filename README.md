# 🐱 Random Cat Generator

A simple and fun project that generates random cat images using a Go backend and a lightweight HTML/CSS frontend.

This project fetches random cat images from public APIs and displays them along with randomly generated metadata such as breed, description, and tags.

---

## ✨ Features

- Go backend API
- Simple frontend (HTML, CSS, JavaScript)
- Random cat images from public APIs
- Random breed, description, and tags
- Download cat images
- Health check endpoint

---

## 🛠 Tech Stack

- **Backend:** Go (net/http)
- **Frontend:** HTML, CSS, Vanilla JavaScript
- **APIs:**  
  - [CATAAS](https://cataas.com)  
  - [The Cat API](https://thecatapi.com)

---

## 📂 Project Structure

```
    RandomCat-API/
    ├── backend
    │ └── main.go
    └── frontend
    ├── index.html
    └── style.css
```

---

## ▶️ How to Run

### 1. Clone the repository
```bash
git clone https://github.com/your-username/random-cat-generator.git
cd random-cat-generator/backend
```
### 2. Run the backend server
```bash
go run main.go
```
### 3. Open in browser
```bash
http://localhost:8080
```

---

## 📌 Notes

* This project is intended for learning and demonstration purposes.

* No frontend framework is used.

* Backend serves static frontend files directly.