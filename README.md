This is IT Manangement project

## Project Directory Structure (MVC Model)

```
ITMGMT/
├── controllers/   # Xử lý logic nghiệp vụ, nhận request từ client, gọi model và trả về response
├── models/        # Định nghĩa cấu trúc dữ liệu, tương tác với database
├── views/         # Giao diện người dùng, trả về dữ liệu hoặc HTML cho client
├── routes/        # Định nghĩa các endpoint, ánh xạ URL tới controller
├── config/        # Cấu hình ứng dụng (database, môi trường, v.v.)
├── utils/         # Các hàm tiện ích dùng chung
├── main.go        # File khởi động ứng dụng
└── README.md      # Tài liệu dự án
```

### Giải thích mô hình MVC

- **Model**: Quản lý dữ liệu, định nghĩa cấu trúc và các thao tác với dữ liệu (CRUD).
- **View**: Hiển thị dữ liệu cho người dùng, có thể là HTML, JSON, v.v.
- **Controller**: Nhận request từ client, xử lý logic, gọi model và trả về view phù hợp.

Mô hình MVC giúp tách biệt các phần của ứng dụng, dễ bảo trì và mở rộng.