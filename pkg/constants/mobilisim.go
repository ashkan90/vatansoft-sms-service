package constants

const (
	MaxMessageInTime = 1500
)

const (
	MobilisimErrorStatus   = "error"
	MobilisimSuccessStatus = "success"

	MobilisimSuccessDescription = "SMS gönderimizin başarıyla başlatıldı."

	MobilisimSystemError        = "STH tarafında bilinmeyen bir hata oluştu. Lütfen daha sonra tekrar deneyiniz."
	MobilisimUnauthorizedError  = "Geçersiz gönderici adı veya oturum hatası."
	MobilisimInvalidSenderError = "Geçersiz gönderici adı."
	MobilisimCreditError        = "İşlemi yapmak için yeterli krediniz bulunmamakta."
)

const (
	MobilisimEnglishMessageDecoder = "normal"
	MobilisimTurkishMessageDecoder = "turkish"
)
