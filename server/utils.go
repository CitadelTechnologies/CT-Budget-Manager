package server

import(
    "ct-budget-manager/exception"
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
)

func DecodeJsonRequest(r *http.Request) map[string]interface{} {
    var body []byte
    var err error
    if body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err != nil {
        panic(exception.New(http.StatusInternalServerError, "Request body could not be opened"))
    }
    if err = r.Body.Close(); err != nil {
        panic(exception.New(http.StatusInternalServerError, "Request body could not be closed"))
    }
    var data map[string]interface{}
    if err = json.Unmarshal(body, &data); err != nil {
        panic(exception.New(http.StatusUnprocessableEntity, "Request body could not be parsed"))
    }
    return data
}

func SendJsonResponse(w http.ResponseWriter, code int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    if err := json.NewEncoder(w).Encode(&data); err != nil {
        panic(err)
    }
}

func CatchException(w http.ResponseWriter) {
    if r := recover(); r != nil {
        exception := r.(*exception.Exception)
        SendJsonResponse(w, exception.Code, exception)
    }
}
