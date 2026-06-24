package database

import (
	"log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"QuickWork/models" // Chanh nhớ chỉnh lại đường dẫn package models đúng với dự án của bạn nhé
)

func SeedDatabase(db *gorm.DB) {
	log.Println("🌱 Đang bắt đầu seed data thực tế khớp cấu trúc Model...")

	// 1. Tạo mật khẩu mã hóa chung: 12345678
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	passStr := string(hashedPassword)

	// ==========================================
	// 🏢 SEED 5 DOANH NGHIỆP CÔNG NGHỆ THỰC TẾ
	// ==========================================
	companiesData := []struct {
		Name    string
		Email   string
		TaxCode string
		Phone   string
		Address string
	}{
		{"FPT Software", "hr@fpt-software.com.vn", "0101248141", "02473007300", "Khu CNC Hòa Lạc, Thạch Thất, Hà Nội"},
		{"VNG Corporation", "recruitment@vng.com.vn", "0303494762", "02839623888", "Z06 Đường số 13, Tân Thuận Đông, Quận 7, TP.HCM"},
		{"Viettel Digital", "talent@viettel.com.vn", "0100109106", "02462556789", "Số 1 Trần Hữu Dực, Mỹ Đình, Nam Từ Liêm, Hà Nội"},
		{"MISA Joint Stock Company", "hr@misa.com.vn", "0100446971", "02437959595", "Tòa nhà MISA, Ngõ 82 Dịch Vọng Hậu, Cầu Giấy, Hà Nội"},
		{"NashTech Vietnam", "careers@nashtechglobal.com", "0302172154", "02838153030", "Tòa nhà ETown 1, 364 Cộng Hòa, Tân Bình, TP.HCM"},
	}

	var businessIDs []uint // Danh sách ID từ bảng Business để nối sang bảng Job
	for _, comp := range companiesData {
		// Tạo tài khoản User
		user := models.User{
			Email:    comp.Email,
			Password: passStr,
			Role:     "business",
			Status:   "approved",
			Balance:  1000000.00, // Tặng sẵn 1 triệu trong ví test chơi
		}
		if err := db.Create(&user).Error; err != nil {
			continue // Nếu email đã tồn tại thì bỏ qua
		}

		// Tạo Profile Business tương ứng
		business := models.Business{
			UserID:      user.ID,
			CompanyName: comp.Name,
			TaxCode:     comp.TaxCode,
			Phone:       comp.Phone,
			Address:     comp.Address,
			LogoUrl:     "https://placehold.co/150x150/0f172a/fff?text=" + comp.Name,
			IsVerified:  true,
		}
		if err := db.Create(&business).Error; err == nil {
			businessIDs = append(businessIDs, business.ID) // Lưu lại Business.ID thực tế để gắn vào Job
		}
	}

	// ==========================================
	// 🎓 SEED 20 SINH VIÊN THỰC TẾ (HUST, UET, FPT...)
	// ==========================================
	studentNames := []string{
		"Nguyễn Văn Minh", "Trần Thị Hồng", "Lê Hoàng Nam", "Phạm Minh Đức", "Vũ Anh Tú",
		"Đặng Hồng Ngọc", "Hoàng Quốc Bảo", "Bùi Tiến Dũng", "Đỗ Thùy Linh", "Phan Thanh Tùng",
		"Ngô Khánh Huyền", "Dương Quốc Anh", "Lý Minh Triết", "Vỏ Hoàng Yến", "Trịnh Xuân Bách",
		"Đinh Quang Hải", "Tạ Thị Mai", "Mai Đức Kiên", "Hà Tiểu Phương", "Cao Minh Quang",
	}

	genders := []string{"Male", "Female"}
	universities := []string{"HUST", "UET", "FPT University", "PTIT", "NEU"}

	for i, name := range studentNames {
		email := ""
		if i < 10 {
			email = "student" + string(rune(48+i)) + "@gmail.com" // student0@gmail.com -> student9@gmail.com
		} else {
			email = "student1" + string(rune(48+(i-10))) + "@gmail.com" // student10@gmail.com -> student19@gmail.com
		}

		// Tạo tài khoản User cho Student
		user := models.User{
			Email:    email,
			Password: passStr,
			Role:     "student",
			Status:   "approved",
			Balance:  0.00,
		}
		if err := db.Create(&user).Error; err != nil {
			continue
		}

		// Tạo thông tin chi tiết Student
		uni := universities[i%len(universities)]
		student := models.Student{
			UserID:    user.ID,
			FullName:  name,
			Phone:     "09876543" + string(rune(48+(i%10))) + string(rune(48+(i/10))),
			Gender:    genders[i%2],
			AvatarUrl: "https://api.dicebear.com/7.x/adventurer/svg?seed=" + name,
			Skills:    "HTML, CSS, JavaScript, Git, Kỹ năng làm việc nhóm. Sinh viên trường " + uni,
			CvUrl:     "https://quickwork.com/cv/sample-cv.pdf",
		}
		db.Create(&student)
	}

	// ==========================================
	// 💼 SEED 10 JOB TUYỂN DỤNG INTERN/FRESHER CHUẨN ĐẸP
	// ==========================================
// ==========================================
	// 💼 SEED 10 JOB TUYỂN DỤNG INTERN/FRESHER CHUẨN ĐẸP
	// ==========================================
	if len(businessIDs) < 5 {
		log.Println("⚠️ Số lượng Business tạo thành công dưới 5, dừng luồng seed Job để tránh lỗi trống ID.")
		return
	}

	jobs := []models.Job{
		{BusinessID: businessIDs[0], Title: "Golang Backend Intern", Location: "Quận Cầu Giấy, Hà Nội", Salary: 4500000.00, Slots: 5, Status: "approved", WorkingDate: "Thứ 2 - Thứ 6 (8:30 - 17:30)", Description: "Yêu cầu kiến thức cơ bản về Go, Gin/Fiber framework, biết sử dụng Git và tư duy database tốt. Hỗ trợ dấu thực tập."},
		{BusinessID: businessIDs[0], Title: "VueJS/NuxtJS Frontend Fresher", Location: "Quận 9, TP. Hồ Chí Minh", Salary: 9500000.00, Slots: 3, Status: "approved", WorkingDate: "Thứ 2 - Thứ 6", Description: "Tham gia phát triển các dự án Outsource lớn sử dụng Vue 3, Nuxt 3/4. Ưu tiên các bạn làm đồ án tốt nghiệp bằng Vue/Nuxt."},
		
		{BusinessID: businessIDs[1], Title: "Fullstack Web Developer Intern", Location: "Quận 7, TP. Hồ Chí Minh", Salary: 5000000.00, Slots: 2, Status: "approved", WorkingDate: "Linh hoạt 4 ngày/tuần", Description: "Thực tập tại ZaloPay team. Làm việc với NodeJS/React. Có cơ hội trở thành nhân viên chính thức sau 3 tháng."},
		{BusinessID: businessIDs[1], Title: "Game Development Trainee (C++)", Location: "Quận Cầu Giấy, Hà Nội", Salary: 7000000.00, Slots: 4, Status: "approved", WorkingDate: "Thứ 2 - Thứ 6", Description: "Đam mê ngành game, tư duy logic tốt, cấu trúc dữ liệu và giải thuật vững vàng. Được đào tạo bài bản từ đầu."},
		
		{BusinessID: businessIDs[2], Title: "Backend Engineer Intern (Java)", Location: "Quận Đống Đa, Hà Nội", Salary: 5500000.00, Slots: 3, Status: "approved", WorkingDate: "Thứ 2 - Thứ 6", Description: "Nghiên cứu và phát triển hệ sinh thái Viettel Money. Yêu cầu biết Java Core, kiến thức căn bản về SQL Server/MySQL."},
		{BusinessID: businessIDs[2], Title: "DevOps Engineer Intern", Location: "Quận Đống Đa, Hà Nội", Salary: 4000000.00, Slots: 2, Status: "approved", WorkingDate: "Thứ 2 - Thứ 6", Description: "Làm quen với Docker, K8s, CI/CD pipelines (Jenkins/GitLab) dưới sự hướng dẫn trực tiếp của các Senior."},
		
		{BusinessID: businessIDs[3], Title: "Fresher Automation Tester", Location: "Quận Hà Đông, Hà Nội", Salary: 8500000.00, Slots: 5, Status: "approved", WorkingDate: "Thứ 2 - Thứ 6", Description: "Đọc hiểu tài liệu nghiệp vụ, viết testcase và thực hiện test các phần mềm kế toán, ERP của MISA. Được hướng dẫn viết script test tự động."},
		{BusinessID: businessIDs[3], Title: "Business Analyst (BA) Intern", Location: "Quận Cầu Giấy, Hà Nội", Salary: 3000000.00, Slots: 2, Status: "approved", WorkingDate: "Tối thiểu 3 ngày/tuần", Description: "Cầu nối giữa khách hàng khối doanh nghiệp và team phát triển phần mềm. Kỹ năng giao tiếp, lắng nghe và làm tài liệu tốt."},
		
		{BusinessID: businessIDs[4], Title: "React Native Mobile Intern", Location: "Quận 3, TP. Hồ Chí Minh", Salary: 4500000.00, Slots: 3, Status: "approved", WorkingDate: "Thứ 2 - Thứ 6", Description: "Phát triển ứng dụng Mobile đa nền tảng bằng React Native. Hiểu biết về ES6, React Hooks là lợi thế lớn."},
		{BusinessID: businessIDs[4], Title: "Data Engineer Fresher", Location: "Quận Cầu Giấy, Hà Nội", Salary: 11000000.00, Slots: 2, Status: "approved", WorkingDate: "Thứ 2 - Thứ 6", Description: "Xây dựng luồng dữ liệu ETL/ELT. Có kiến thức về Python, SQL nâng cao. Có tư duy phân tích dữ liệu nhạy bén."},
	}

	for _, job := range jobs {
		db.Create(&job)
	}

	log.Println("🎉 Khởi tạo dữ liệu thực tế thành công mỹ mãn! Tất cả mật khẩu đều là: 12345678")
}