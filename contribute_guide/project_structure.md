### 📂 ส่องโครงสร้างโปรเจกต์: เราจัดวางโค้ดกันยังไงนะ?

โปรเจกต์ของเราจัดโครงสร้างแบบ **MVC กึ่ง Hexagonal** เพื่อให้โค้ดเป็นระเบียบ เข้าใจง่าย และขยายต่อได้ในอนาคต:

```
your-project/
├── app/                    ⚙️ จัดการการเริ่มต้นระบบ เช่น Init Fiber, HTTP server, การเชื่อมต่อฐานข้อมูล และ lifecycle ของแอปพลิเคชัน
│
├── cmd/
│   ├── script/             🖥️ รวมไฟล์ script เช่น migration, seed ข้อมูล, job พิเศษ, CLI tools
│   ├── server/             🌐 จุดเริ่มต้นของแอปแบบ production หรือ local run
│   │   ├── main.go         📌 จุด Entry point ของแอป – เรียกโหลด config, init ระบบ, start HTTP server
│
├── config/                 🛠️ สำหรับโหลด config จากไฟล์ `config.yml` ผ่าน Viper หรือ lib ที่ใช้อ่าน config
│
├── models/                 🧩 Struct หรือ Schema ที่ใช้เป็น data model กลาง เช่น DB models, DTO กลาง, etc.
│
├── modules/                📦 รวมทุก feature/module ที่แยกตามหน้าที่ เช่น frontweb, openapi
│   ├── frontweb/           🌐 ส่วนของ Web client หรือ frontend ที่เชื่อมต่อกับระบบ
│   │   ├── feature/        🎯 แบ่ง Feature ตาม use-case เช่น Auth, User, Patient, etc.
│   │   │   ├── dto/            📄 Struct สำหรับรับ-ส่งข้อมูล (Request/Response)
│   │   │   ├── handler/        🎛️ จัดการ Request และเรียก service (Controller layer)
│   │   │   ├── port/           🔌 Interface ที่เชื่อมระหว่าง Layer (เช่น Repository/Service)
│   │   │   ├── repository/     🗄️ จัดการการเข้าถึงข้อมูล เช่น DB หรือ API ภายนอก
│   │   │   ├── service/        ⚙️ Business Logic ของ feature
│   │   ├── middleware/     🧱 ตัวดัก Request เช่น JWT, Logging, Validation
│   │   ├── repository/     📚 ใช้สำหรับ shared repo กลางที่ไม่ผูกกับ feature เดียว
│   ├── openapi/            🌍 คล้าย frontweb แต่แยกไว้สำหรับ public API หรือ external service interface
│
├── package/                📦 รวมของใช้ทั่วไปที่ไม่ขึ้นกับ module เช่น util, db, 3rd party
│   ├── util/               🧰 ฟังก์ชันช่วยเหลือ เช่น string formatter, error handler
│   ├── db/                 🛢️ ฟังก์ชันหรือ struct สำหรับจัดการ DB เช่น MongoDB, MariaDB
│   ├── 3rd party/          🔌 การ integrate กับระบบภายนอก เช่น payment, sms, s3
│
├── config.yml              📝 ไฟล์ config หลัก เช่น port, database, api key
├── Dockerfile              🐳 สำหรับ build docker image
├── go.mod                  📦 Go module config
├── go.sum                  🔐 checksums ของ dependency

```

---

[📚 สารบัญ](start_here.md/#-สารบัญ)
