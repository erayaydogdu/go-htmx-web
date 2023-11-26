# GO Htmx Web App

It shows the capabilities of HTMx

## Stacks Used

This project is built using the following technologies:

- **Go:** The programming language used for server-side development.
- **Htmx:** A library for enhancing web pages by adding minimal JavaScript to your HTML.
- **PostgreSQL:** The relational database used for data storage.
- **Tailwind CSS:** A utility-first CSS framework used for styling the frontend.

## Getting Started

Follow these steps to get the project up and running on your local machine:

1. **Build and Run the Dockerized PostgreSQL Database:**
   ```bash
   docker build -t todopostgresdb .
   docker run -p 6432:5432 --name todopostgresdb -d todopostgresdb
- The PostgreSQL database is running in a Docker container.
    - Connection details:
    - Host: localhost
    - Port: 6432
    - Database: todo
    - User: todouser
    - Password: todo1234!

2. **Install the Dependencies:**
   ```bash
   go get -u
