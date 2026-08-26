# QuickWork Admin Flow Guide

## 1. Business Flow

1. Business gửi form đăng ký.
2. Backend tạo `users` với `role = business`, `status = pending`.
3. Backend tạo `businesses` với `is_verified = false`.
4. Business pending không được cấp JWT khi đăng nhập.
5. Admin đăng nhập và nhận JWT có `role = admin`.
6. Admin xem `/admin/dashboard` để lấy số liệu thật từ database.
7. Admin xem `/admin/businesses/pending` để tìm kiếm và phân trang hồ sơ KYB pending.
8. Admin mở hồ sơ KYB, duyệt hoặc từ chối kèm lý do.
9. Backend cập nhật `users.status` và `businesses.is_verified` trong cùng transaction.
10. Hồ sơ đã xử lý biến mất khỏi queue pending.

## 2. IPO Dashboard

Input:
- `GET /api/admin/dashboard/stats`
- Header `Authorization: Bearer <admin-token>`

Process:
- Xác thực JWT.
- Kiểm tra role admin.
- Đếm student, business, business pending, job, job pending.
- Tính `total_disbursed` từ `wallet_transactions.type = salary`.

Output:
- `total_students`
- `total_businesses`
- `pending_businesses`
- `total_jobs`
- `pending_jobs`
- `total_disbursed`

`total_disbursed` là tổng lương đã giải ngân qua hệ thống, không phải doanh thu ròng QuickWork.

## 3. IPO Pending Business List

Input:
- `GET /api/admin/businesses/pending?page=1&limit=10&search=fpt`
- Header `Authorization: Bearer <admin-token>`

Process:
- Validate `page >= 1`.
- Validate `1 <= limit <= 100`.
- Trim `search`.
- Filter `users.role = business AND users.status = pending`.
- Search theo `businesses.company_name`, `businesses.tax_code`, `users.email`.
- Count total.
- Apply limit, offset, order by `businesses.created_at DESC`.
- Map sang DTO, không trả entity trực tiếp.

Output:
- `items`
- `pagination`

## 4. IPO Review Action

Input:
- `PUT /api/admin/businesses/:id/review`
- Body:

```json
{
  "decision": "approved",
  "reject_reason": ""
}
```

Process:
- Validate business id.
- Validate decision chỉ là `approved` hoặc `rejected`.
- Nếu `rejected`, reject reason phải có ít nhất 10 ký tự sau khi trim.
- Mở transaction.
- Lock row business và user bằng `FOR UPDATE`.
- Chỉ xử lý nếu user còn `pending`.
- Cập nhật user và business.

Output:

```json
{
  "message": "Đã cập nhật kết quả KYB doanh nghiệp",
  "decision": "approved"
}
```

## 5. API Table

| Method | Path | Role | Purpose |
| --- | --- | --- | --- |
| GET | `/api/admin/dashboard/stats` | admin | Lấy số liệu dashboard |
| GET | `/api/admin/businesses/pending` | admin | Lấy danh sách business pending |
| GET | `/api/admin/businesses/:id` | admin | Lấy chi tiết KYB |
| PUT | `/api/admin/businesses/:id/review` | admin | Duyệt hoặc từ chối KYB |

## 6. Transaction

Kết quả KYB nằm ở hai bảng:
- `users.status`
- `businesses.is_verified`

Transaction bảo đảm hai bảng cùng được cập nhật thành công hoặc cùng rollback. Điều này tránh trường hợp tài khoản business được duyệt đăng nhập nhưng hồ sơ `businesses` chưa verified.

## 7. Row Lock

Review API dùng `FOR UPDATE` để khóa row business và user trong transaction. Nếu hai Admin cùng xử lý một hồ sơ, request đầu tiên cập nhật status khỏi `pending`; request sau khi lấy lock sẽ thấy hồ sơ đã xử lý và trả `409 Conflict`.

## 8. Postman Test

1. Gọi admin API không token: kỳ vọng `401`.
2. Gọi admin API bằng token student/business: kỳ vọng `403`.
3. Login admin `admin@quickwork.vn / 12345678`.
4. Gọi `GET /api/admin/dashboard/stats`: kỳ vọng `200`.
5. Đăng ký business mới.
6. Login business mới: kỳ vọng `403` do status pending.
7. Gọi `GET /api/admin/businesses/pending`: thấy business mới.
8. Gọi reject không có lý do: kỳ vọng `400`.
9. Gọi reject lý do dưới 10 ký tự: kỳ vọng `400`.
10. Gọi approve business pending: kỳ vọng `200`.
11. Gọi lại pending list: business đã xử lý biến mất.
12. Login business approved: kỳ vọng đăng nhập thành công.
13. Review lại business đã xử lý: kỳ vọng `409`.
14. Gọi detail với id không tồn tại: kỳ vọng `404`.

## 9. Thứ Tự File Nên Đọc

1. `internal/models/user.go`
2. `internal/models/business.go`
3. `internal/database/migrate.go`
4. `internal/database/seed.go`
5. `internal/middleware/auth_middleware.go`
6. `internal/handlers/auth_handler.go`
7. `internal/dto/admin_business.go`
8. `internal/handlers/admin_handler.go`
9. `internal/routes/admin_routes.go`
10. `frontend/composables/useAdminApi.ts`
11. `frontend/pages/admin/dashboard.vue`
12. `frontend/pages/admin/businesses/pending.vue`

## 10. Status

| Status | Meaning |
| --- | --- |
| `pending` | Chờ duyệt, chưa được đăng nhập business dashboard |
| `approved` | Đã duyệt, được đăng nhập |
| `rejected` | Bị từ chối, không được đăng nhập |
| `locked` | Bị khóa, không được đăng nhập |
| `active` | Dữ liệu cũ được tạm tương thích khi login |
