package errno

var (
	ErrDBRecordAlreadyExist = &Errno{HTTP: 400, ErrorCode: 20001, ResponseCode: "DatabaseError.RecordAlreadyExist", Message: "Database record already exist"}
	ErrDBRecordNotFound     = &Errno{HTTP: 404, ErrorCode: 20002, ResponseCode: "DatabaseError.RecordNotFound", Message: "Database recourd not found"}
	ErrDBExecute            = &Errno{HTTP: 404, ErrorCode: 20003, ResponseCode: "DatabaseError.ExecuteError", Message: "Database Execute error"}
)
