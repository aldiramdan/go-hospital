<h1 align="center">
  Hospital net/http Native Api
</h1>
<h2 align="center">
  Golang Api net/http Native
</h2>

## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/aldiramdan/go-hospitalnative.git
```

2. Install dependencies

```bash
go get -u ./...
# or
go mod tidy
```

3. Add Env

```sh

# Storage CSV
PathDoctorCSV = "./databases/storages/csv/doctor.csv"
PathPatientCSV = "./databases/storages/csv/patient.csv"
PathDiseaseCSV = "./databases/storages/csv/disease.csv"
PathMedicineCSV = "./databases/storages/csv/medicine.csv"
PathHandlingCSV = "./databases/storages/csv/handling.csv"

# Path Migrate
PathMigrateUp = "./databases/db/migrations/202307280923_dump.up.sql"
PathMigrateDown = "./databases/db/migrations/202307280923_dump.down.sql"

# Set Database
MYSQL_USER = 
MYSQL_PASSWORD = 
MYSQL_HOST = 
MYSQL_PORT = 
MYSQL_DBNAME = 

# App
PORT = Your Port App
BASE_URL = http://localhost:PORT //example
```

4. Run the app

```bash
go run *.go migrate --up
```

5. Run the app

```bash
go run *.go serve
```

🌟 You are all set!

## 💻 Built with

- [Golang](https://go.dev/): programming language
