# 📘 README.md

## 🌳 File Tree

```markdown
├── 📁 compose
│   └── 📁 go
│       └── 🐳 Dockerfile
├── 📁 server
│   ├── 🐹 handlers.go
│   ├── 🐹 router.go
│   └── 🐹 server.go
├── ⚙️ .gitignore
├── ⚙️ compose.yaml
├── 📄 go.mod
├── 📄 go.sum
└── 🐹 main.go
```

## ▶️ Cómo ejecutar el proyecto

1. Verifica que Docker esté instalado:
    
    ```bash
    docker --version
    docker compose version
    ```
    
2. Clona el repositorio:
    
    ```bash
    git clone https://github.com/Nicolasrc404/GO-BACKEND-CLASE.git
    cd backend-go-avanzado
    ```
    
3. Construye y levanta el contenedor:
    
    ```bash
    docker compose up --build -d
    ```
    
4. Abre en tu navegador:
    
    👉 [http://localhost:8080](http://localhost:8080/)
    
    Verás el mensaje:
    
    ```
    {"hola":"Mundo"}
    ```
    

## 🧰 Comandos útiles

| Comando                                  | Descripción                             |
| ---------------------------------------- | --------------------------------------- |
| `docker compose up -d`                   | Levanta el servicio en segundo plano    |
| `docker compose down`                    | Detiene y elimina los contenedores      |
| `docker compose logs -f`                 | Muestra los logs del contenedor         |
| `docker exec -it backend-go-avanzado sh` | Abre una consola dentro del contenedor  |
| `docker system prune -a`                 | Limpia imágenes y contenedores antiguos |