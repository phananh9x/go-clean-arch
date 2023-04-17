package errors

var contractCodeError = map[string]string{
	"416": "Mã khách hàng không tồn tại (416)",
	"419": "Mã khách hàng không tồn tại (419)",
	"432": "Số điện thoại không hợp lệ (432)",
	"433": "Truy vấn tài khoản thất bại (433)",
	"440": "Số điện thoại không hợp lệ (440)",
}

var dealError = map[string]string{
	"888": "Giao dịch không thành công (888)",
	"18":  "Giao dịch không thành công (18)",
	"202": "Giao dịch không thành công (202)",
	"204": "Giao dịch không thành công (204)",
	"205": "Giao dịch không thành công (205)",
	"206": "Giao dịch không thành công (206)",
	"207": "Giao dịch không thành công (207)",
	"208": "Giao dịch không thành công (208)",
	"209": "Giao dịch không thành công (209)",
	"214": "Giao dịch không thành công (214)",
	"310": "Giao dịch không thành công (310)",
	"318": "Giao dịch không thành công (318)",
	"399": "Giao dịch không thành công (399)",
	"417": "Giao dịch không thành công (417)",
	"418": "Giao dịch không thành công (418)",
	"422": "Giao dịch không thành công (422)",
	"423": "Giao dịch không thành công (423)",
	"424": "Giao dịch không thành công (424)",
	"425": "Giao dịch không thành công (425)",
	"426": "Giao dịch không thành công (426)",
	"427": "Giao dịch không thành công (427)",
	"428": "Giao dịch không thành công (428)",
	"429": "Giao dịch không thành công (429)",
	"431": "Giao dịch không thành công (431)",
	"439": "Giao dịch không thành công (439)",
	"441": "Giao dịch không thành công (441)",
	"99":  "Đang có giao dịch chờ xử lý, vui lòng thử lại sau. (99)",
	"999": "Đang có giao dịch chờ xử lý, vui lòng thử lại sau. (999)",
}

var technicalError = map[string]string{
	"210": "Lỗi kỹ thuật. Vui lòng thử lại sau. (210)",
	"211": "Lỗi kỹ thuật. Vui lòng thử lại sau. (211)",
	"213": "Lỗi kỹ thuật. Vui lòng thử lại sau. (213)",
	"215": "Lỗi kỹ thuật. Vui lòng thử lại sau. (215)",
	"301": "Lỗi kỹ thuật. Vui lòng thử lại sau. (301)",
	"302": "Lỗi kỹ thuật. Vui lòng thử lại sau. (302)",
	"303": "Lỗi kỹ thuật. Vui lòng thử lại sau. (303)",
	"304": "Lỗi kỹ thuật. Vui lòng thử lại sau. (304)",
	"305": "Lỗi kỹ thuật. Vui lòng thử lại sau. (305)",
	"438": "Lỗi kỹ thuật. Vui lòng thử lại sau. (438)",
	"420": "Lỗi kỹ thuật. Vui lòng thử lại sau. (420)",
	"421": "Lỗi kỹ thuật. Vui lòng thử lại sau. (421)",
}

var iMediaErrors = map[string]string{
	"888": "Hủy Thành Công Giao dịch thanh toán đã được hủy thành công",
	"99":  "Pending Giao dịch đang xử lý (Cần kiểm tra lại kết quả)",
	"999": "Hủy Pending Giao dịch Hủy đang được xử lý",
	"18":  "Thất bại Giao dịch thất bại.",
	"202": "Thất bại Thông tin giao dịch không đúng (thông tin dữ liệu không đúng định dạng, vượt quá độ dài,…)",
	"204": "Thất bại pr_code không hợp lệ",
	"205": "Thất bại Giao dịch trùng lặp",
	"206": "Thất bại không tìm thấy thông tin đối tác trên hệ thống.",
	"207": "Thất bại chỹ ký authkey không hợp lệ",
	"208": "Thất bại số tiền thanh toán không hợp lệ",
	"209": "Thất bại mã dịch vụ không đúng hoặc không hoạt động",
	"210": "Thất bại sai thông tin username,password",
	"211": "Thất bại mã giao dịch không tồn tại",
	"213": "Thất bại Mã giao dịch đối tác không hợp lệ.",
	"214": "Thất bại mã giao dịch gốc rỗng hoặc trống.",
	"215": "Thất bại giao dịch không được phép hủy",
	"301": "Thất bại chưa cấu hình phí dịch vụ.",
	"302": "Thất bại chưa cấu hình cổng dịch vụ.",
	"303": "Thất bại cổng dịch vụ bị tắt.",
	"304": "Thất bại chưa kích hoạt thông tin đối tác.",
	"305": "Thất bại chưa cấu hình thông tin đối tác.",
	"310": "Thất bại trừ tiền thất bại,số dư tài khoản đối tác không đủ để thực hiện giao dịch.",
	"318": "Thất bại trừ tiền thất bại.",
	"399": "Thất bại trừ tiền timeout.",
	"416": "Thất bại mã khách hàng không tồn tại.",
	"417": "Thất bại giao dịch thanh toán đã bị hủy.",
	"418": "Thất bại đã tồn tại giao dịch hủy đang được xử lý.",
	"419": "Thất bại mã khách hàng không đúng hoặc chưa hỗ trợ thanh toán.",
	"422": "Thất bại không thể thực hiện giao dịch.",
	"423": "Thất bại lỗi thanh toán, vui lòng thử lại sau.",
	"424": "Thất bại lỗi hạch toán tài khoản thanh toán.",
	"425": "Thất bại giao dịch thất bại.",
	"426": "Thất bại giao dịch thất bại.",
	"427": "Thất bại lỗi hệ thống thanh toán.",
	"428": "Thất bại quý khách không còn nợ cước dịch vụ.",
	"429": "Thất bại mã thanh toán không còn tồn tại.",
	"431": "Thất bại Đối tác thanh toán dịch vụ tạm ngừng phục vụ",
	"432": "Thất bại Số điện thoại không hợp lệ hoặc không đủ điều kiện thanh toán",
	"433": "Thất bại lỗi xảy ra khi truy vấn tài khoản.",
	"439": "Thất bại gạch nợ thuê bao thất bại.",
	"440": "Thất bại số điện thoại không hợp lệ.",
	"441": "Thất bại giao dịch gốc không tồn tại.",
	"438": "Thất bại lỗi hệ thống phía nhà cung cấp dịch vụ.",
	"420": "Thất bại hệ thống nhà cung cấp dịch vụ đang tạm ngưng phục vụ.",
	"421": "Thất bại hệ thống nhà cung cấp đang nâng cấp, bảo dưỡng.",
}

// GetMerchantError return error with text from merchant
func GetMerchantError(errCode string) error {
	errorText, ok := contractCodeError[errCode]
	if ok {
		return NewInvalidInputError("contract_code", errorText)
	}

	errorText, ok = dealError[errCode]
	if ok {
		return New(errorText)
	}

	errorText, ok = technicalError[errCode]
	if ok {
		return New(errorText)
	}

	return nil
}

//GetMerchantMessage ...
func GetMerchantMessage(errCode string) string {
	errorText, ok := contractCodeError[errCode]
	if ok {
		return errorText
	}

	errorText, ok = dealError[errCode]
	if ok {
		return errorText
	}

	errorText, ok = technicalError[errCode]
	if ok {
		return errorText
	}

	return "Đơn hàng không thành công"
}

//GetFullMerchantMessage ...
func GetFullMerchantMessage(errCode string) string {
	errorText, ok := iMediaErrors[errCode]
	if ok {
		return errorText
	}

	return "Không gọi được qua API của đối tác"
}
