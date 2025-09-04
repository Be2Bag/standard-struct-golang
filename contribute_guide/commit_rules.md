### 💬 สร้าง Commit Message ให้ว้าว\! (Conventional Commits)

เพื่อให้ Commit History ของเราเป็นระเบียบ อ่านง่าย และรู้เรื่อง เราใช้มาตรฐาน [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) ครับ

**รูปแบบที่แนะนำ:**

```
<type>(<scope>): <subject>
```

**ตัวอย่าง Commit Message เจ๋งๆ:**

- `feat(queue): เพิ่มฟังก์ชันอัปเดตคิวแบบเรียลไทม์ด้วย WebSocket`
- `fix(auth): แก้ไขปัญหาการ Redirect หลัง Login ผิดพลาด`
- `docs(README): ปรับปรุงคู่มือการ Contribute ให้น่าอ่านขึ้น`
- `style(components): จัดระเบียบสไตล์ปุ่มบนหน้า Admin Dashboard`
- `refactor(stores): ปรับปรุงประสิทธิภาพการทำงานของ queue store`
- `chore(deps): อัปเดต Vuetify เป็นเวอร์ชันล่าสุด`
- `perf(dashboard): ปรับปรุงการโหลดข้อมูลบนแดชบอร์ดให้เร็วขึ้น`

**ประเภท (type) ที่ใช้บ่อย:**

- `feat`: ฟีเจอร์ใหม่เอี่ยม\! ✨
- `fix`: แก้บั๊กที่กวนใจ 🐞
- `docs`: เปลี่ยนแปลงแค่เอกสาร 📝
- `style`: ปรับปรุงแค่สไตล์โค้ด ไม่ได้เปลี่ยนการทำงาน 💅
- `refactor`: ปรับปรุงโครงสร้างโค้ดให้ดีขึ้น (ไม่ได้เพิ่มฟีเจอร์/แก้บั๊ก) ♻️
- `test`: เพิ่มหรือแก้ไข Test 🧪
- `chore`: งานเบ็ดเตล็ด เช่น อัปเดต Dependency, Config ต่างๆ 🧹
- `perf`: เพิ่มประสิทธิภาพการทำงาน ⚡
- `ci`: เปลี่ยนแปลงใน CI/CD configuration 🤖

**ขอบเขต (scope):** (ไม่บังคับ) บอกว่า Commit นี้เกี่ยวข้องกับส่วนไหนของโปรเจกต์ (เช่น `auth`, `user`, `admin`, `patient`)

**หัวข้อ (subject):** คำอธิบายสั้นๆ กระชับ (ไม่เกิน 50 ตัวอักษร) ใช้คำสั่งแบบกริยา (เช่น "เพิ่ม", "แก้ไข", "อัปเดต")

---

[📚 สารบัญ](start_here.md/#-สารบัญ)
