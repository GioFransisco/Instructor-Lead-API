# Instructor Lead App

## Permasalahan
Instructor Lead App merupakan sebuah aplikasi berbasis API yang dibuat untuk mempermudah proses administrasi dari sebuah perusahaan yang mengadakan bootcamp online yang didalam bootcamp tersebut memiliki rutinitas untuk melakukan pertemuan melalui zoom setidaknya dua kali dalam satu minggu. 
<!-- dan didalamnya terdapat proses bisnis seperti penjadwalan trainer, absensi, dan masih banyak lagi. -->

## Tujuan
Tujuan aplikasi ini adalah untuk meningkatkan efisiensi administratif perusahaan bootcamp online, khususnya dalam penjadwalan ketika pertemuan zoom berlangsung. Dengan Instructor Lead App, diharapkan dapat mengoptimalkan proses berlangsungnya kegiatan, mempersingkat waktu dari segi administratif, dan menyediakan pengalaman yang lebih lancar bagi instruktur dan peserta.

## Teknologi 
Menggunakan bahasa pemrograman Golang dan Gin Gonic sebagai framework yang membuat aplikasi menjadi lebih efisien

## Actor (role)
Terdapat beberapa role dalam aplikasi ini :
<ul>
    <li>
        Admin
    </li>
    <li>
        Trainer
    </li>
    <li>
        Peserta
    </li> 
</ul>

## Fitur & CheckPoint
Terdapat beberapa fitur yang ada di aplikasi ini, diantaranya :

<ul>
    <li>
    User Registration & Authentication - Fitur untuk registrasi user oleh admin (Admin mendaftarkan user untuk mendapatkan kredensial login agar bisa mengakses aplikasi Instructor Lead)
    </li>
    <li>
    Penyiapan Jadwal - Fitur untuk persiapan jadwal ketika diadakannya instructor lead
    </li>
    <li>
    Menampung pertanyaan yang diajukan peserta - Fitur untuk menampung seluruh pertanyaan peserta bootcamp yang nantinya akan dibahas oleh trainer ketika instructor lead berlangsung.
    </li>
    <li>
    Absensi Peserta - Fitur untuk absensi kehadiran peserta pada saat instructor lead berlangsung
    </li>
    <li>
    Catatan Trainer - Fitur yang berfungsi supaya trainer bisa menulis catatan pertanyaan yang diajukan oleh peserta bootcamp
    </li>
    <li>
    Bukti Kehadiran Trainer - Fitur yang berfungsi untuk mencatat bukti kehadiran trainer yang dilaporkan melalui upload foto kegiatan disaat kegiatan instructor lead selesai
    </li>
    <li>
    JWT Authentication - Fitur yang berfungsi untuk menggenerate token yang digunakan untuk login Actor
    </li>
    <li>
    Middleware Authorization - Fitur penengah untuk membantu menentukan bagian mana saja yang bisa diakses oleh suatu role
    </li>
</ul>

## API Endpoint

### User API

#### Login User

Request :
- Method : `POST`
- Endpoint : `/api/v1/auth/login`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

```json
{
    "username": "DoctorBeast",
    "password": "password"
}
```

#### Register User
```
Role : Admin
```

Request :
- Method : `POST`
- Endpoint : `/api/v1/auth/register`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

```json
// Body Raw
{
    "name": "Doctor Beast",
    "email": "db@gmail.com",
    "password": "password",
    "address": "Jalan Manunggal Raya",
    "age": 20,
    "gander": "M",
    "phoneNumber": "08891283283",
    "username": "Dbeast"
}

// Body Form Data :
Key : data (type: file)
Value : //select .csv or .xlsx file
```

#### Get User By Email
Request :
- Method : `GET`
- Endpoint : `/api/v1/users/your_email`
- Header :
    - Accept : application/json

Response :
```json

```

#### Get User By Id
Request :
- Method : `GET`
- Endpoint : `/api/v1/users/id/your_uuid`
- Header :
    - Accept : application/json

Response :
```json

```

#### Update User
Request :
- Method : `PUT`
- Endpoint : `/api/v1/users`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "name": "updated_name",
    "email": "updated_email",
    "password": "updated_password",
    "address": "updated_address",
    "age": updated_age,
    "gender": "updated_gender",
    "phoneNumber": "updated_phoneNumber",
    "username": "updated_username"
}
```

#### Change Password User
Request :
- Method : `PUT`
- Endpoint : `/api/v1/users/password`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "password": "changePassword",
}
```

#### Delete User
Request :
- Method : `DELETE`
- Endpoint : `/api/v1/users/your_uuid`
- Header :
    - Accept : application/json
- Body :

Response :
```json
```

### Stack API

#### Create Stack
Request :
- Method : `POST`
- Endpoint : `/api/v1/stacks`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "name": "Python"
}
```

#### List Stack
Request :
- Method : `GET`
- Endpoint : `/api/v1/stacks`
- Header :
    - Accept : application/json
- Body :

Response :
```json
```

#### Get Stack
Request :
- Method : `GET`
- Endpoint : `/api/v1/stacks/your_uuid`
- Header :
    - Accept : application/json
- Body :

Response :
```json
```

#### Update Stack
Request :
- Method : `PUT`
- Endpoint : `/api/v1/stacks/your_uuid`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "name": "Java",
    "status": "Active"
}
```

#### Delete Stack
Request :
- Method : `DELETE`
- Endpoint : `/api/v1/stacks/your_uuid`
- Header :
    - Accept : application/json
- Body :

Response :
```json
```

### Schedule API

#### Create Schedule
Request :
- Method : `POST`
- Endpoint : `/api/v1/schedules`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "name": "Instructor Led Pertemuan 1",
    "dateActivity": "2023-11-21",
    "scheduleDetails": [
        {
            "trainerId": "b04e5411-e5d2-4594-b7b4-5570ae08f5f0",
            "stackId": "b5587cba-a935-4c3a-89a6-5de76285bebb",
            "startTime": "19:30",
            "endTime": "20:30"
        }
    ]
}
```

#### List Schedule
Request :
- Method : `GET`
- Endpoint : `/api/v1/schedules`
- Header :
    - Accept : application/json
- Body :

Response :
```json
```

#### Update Schedule By Id
Request :
- Method : `PUT`
- Endpoint : `/api/v1/schedules/your_uuid`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "name": "updated_name",
    "dateActivity": "2023-11-21"
}
```

#### Update Schedule Detail
Request :
- Method : `PUT`
- Endpoint : `/api/v1/schedule-details/your_uuid`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "trainerId": "b04e5411-e5d2-4594-b7b4-5570ae08f5f0",
    "stackId": "b5587cba-a935-4c3a-89a6-5de76285bebb",
    "startTime": "19:00",
    "endTime": "20:00"
}
```

### Absence API

#### Create Absence
Request :
- Method : `POST`
- Endpoint : `/api/v1/absences`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "scheduleDetail" : {
        "id":"6ad151fc-5189-4b90-b8cc-f2ce3b6e1674"
    },
    "student" : {
        "id" : "dad94891-a314-4182-bcc5-fcaab4688a06"
    },
    "description" : "Hadir tapi tidur"
}
```

#### Get Absences By Schedule Detail Id
Request :
- Method : `GET`
- Endpoint : `/api/v1/absences/your_uuid`
- Header :
    - Accept : application/json

Response :
```json

```

### Question API

#### Create Question
Request :
- Method : `POST`
- Endpoint : `/api/v1/questions`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "scheduleDetail":{
        "id":"6ad151fc-5189-4b90-b8cc-f2ce3b6e1674"
    },
    "student":{
        "id":"dad94891-a314-4182-bcc5-fcaab4688a06"
    },
    "question":"Bagaimana cara belajar coding ?"
}
```

#### Get Question By Schedule Detail Id
Request :
- Method : `GET`
- Endpoint : `/api/v1/questions/your_uuid`
- Header :
    - Accept : application/json

Response :
```json

```

#### Update Question 
Request :
- Method : `PUT`
- Endpoint : `/api/v1/questions/your_uuid`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "question":"updated_question"
}
```

#### Update Status Question
Request :
- Method : `PUT`
- Endpoint : `/api/v1/questions/your_uuid`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
{
    "status":"Finish"
}
```

#### Delete Question
Request :
- Method : `DELETE`
- Endpoint : `/api/v1/questions/your_uuid`
- Header :
    - Accept : application/json
- Body :

Response :
```json
```

### Schedule Approve

#### Create Schedule Approve
Request :
- Method : `POST`
- Endpoint : `/api/v1/schedule-aprove`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

Request :
```json
Body Form Data :
Key : schedule
Value : {
    "scheduleDetail" : {
        "id" : "7a9e4489-d9a0-40c5-bc33-3917f1d72a9b"
    }
}

Key : photo
Value : //selected photo
```

#### Get Schedule Approve
Request :
- Method : `GET`
- Endpoint : `/api/v1/schedule_approve/your_uuid`
- Header :
    - Accept : application/json

Response :
```json

```