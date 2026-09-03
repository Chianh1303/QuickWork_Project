# Quy Trình Làm Việc Bắt Buộc (Workflow Rules)

1. **Không sửa code ngay khi nhận yêu cầu**:
   - Khi người dùng yêu cầu sửa đổi, debug hoặc thêm tính năng, KHÔNG ĐƯỢC tự ý sửa code ngay.
   - Luôn tiến hành nghiên cứu, phân tích kỹ và gửi BÁO CÁO rõ ràng gồm:
     - **Nguyên nhân cốt lõi** (Root cause)
     - **Phương án / Kế hoạch khắc phục chi tiết** (Proposed solution)
   - Chờ người dùng xem xét và xác nhận/đồng ý mới được bắt đầu chỉnh sửa code.

2. **Làm việc trực tiếp trên nhánh `main` & Tự động đẩy Git sau khi hoàn thành**:
   - Tất cả các thao tác sửa code thực hiện trực tiếp trên nhánh **`main`**.
   - Sau khi hoàn thành xong các chỉnh sửa và kiểm tra code, LUÔN LUÔN tự động commit và đẩy lên Git (`git push origin main`) để Render tự động build và deploy ngay lập tức.
