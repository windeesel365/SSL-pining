### การตรึง SSL(SSL-pinning) ช่วยให้ผู้ใช้แอพติดต่อกับ server แท้ ไม่โดนดักกลางทาง(ป้องกัน Man in the Middle)

#### แอพจะเก็บข้อมูลใบรับรอง(certificate)ของเซิร์ฟเวอร์ไว้ในcodeของแอพหรือในไฟล์ที่มีการเข้ารหัส
#### เมื่อแอพพยายามเชื่อมต่อกับเซิร์ฟเวอร์ แอพจะตรวจสอบใบรับรองของเซิร์ฟเวอร์กับใบรับรองที่ถูกเก็บไว้
- **`ถ้าใบรับรองตรงกัน การเชื่อมต่อจะถูกยอมรับและดำเนินการต่อไป`**
- **`ถ้าใบรับรองไม่ตรงกัน การเชื่อมต่อจะถูกปฏิเสธ ทำให้การโจมตีแบบ MITM ล้มเหลว`**
