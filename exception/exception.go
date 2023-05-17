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