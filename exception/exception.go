package exception
// import "fmt"

func GetErr() map[int]string {
    errMap := make(map[int]string)
    errMap[1001] = "starttime is invalid"
    errMap[1002] = "endtime is invalid"

    return errMap
}

func GetCode(s int) string {
    codeMap := make(map[int]string)
    codeMap[200] = "success"
    codeMap[404] = "error"

    return codeMap[s]
}

func GetMessage(s int) string {
    msgMap := make(map[int]string)
    msgMap[200] = "success result"
    msgMap[404] = "error list"

    return msgMap[s]
}

var Message = map[string]string{
    "1111": "a",
    "1222": "b",
    "1333": "c",
    "1444": "d",
    "1555": "e",
    "1666": "f",
    "1777": "g",
}


func GetException(code string) (message string) {
    message = "unknow error"

    errorMsg, ok := Message[code]
    if ok {
        message = errorMsg
    }

    return
}