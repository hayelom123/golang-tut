package helpers

import (
	"net/http"
	"io"
	"errors"
	"encoding/json"
	"mime"
	"strconv"
	"example.com/api/model"
)

func ReadJson(res http.ResponseWriter,req *http.Request,destination any) bool{
	 mediaType,_,err := mime.ParseMediaType(req.Header.Get("Content-Type"))

	if err != nil || mediaType != "application/json" {
        // http.Error(res, "JSON required", http.StatusUnsupportedMediaType)
		writeJSONError(res,  http.StatusUnsupportedMediaType,  "JSON required",)

        return false
    }


      req.Body = http.MaxBytesReader(res,req.Body,1024)

	  body,err:=io.ReadAll(req.Body)

	  if err != nil {
        var maxErr *http.MaxBytesError

        if errors.As(err, &maxErr) {
            // http.Error(res, "Body too large", http.StatusRequestEntityTooLarge)
			writeJSONError(res,  http.StatusRequestEntityTooLarge,   "Body too large",)
        } else {
            // http.Error(res, "Cannot read body", http.StatusBadRequest)
			writeJSONError(res,  http.StatusBadRequest,  "Cannot read body",)
        }
        return false
    }
	 
	 // Pass destination directly (it is already a pointer)
	if err = json.Unmarshal(body, destination); err != nil {
		// http.Error(res, "Failed to parse data: "+err.Error(), http.StatusBadRequest)
		writeJSONError(res,  http.StatusBadRequest, "Failed to parse data: "+err.Error())
		return false
	}
	  return true
}

func SendJson(res http.ResponseWriter, status int, value any){
	 data, err := json.Marshal(value)
	if err != nil {
		// http.Error(res, "Failed to encode data", http.StatusInternalServerError)
		writeJSONError(res,  http.StatusInternalServerError,  "Failed to encode data",)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(data)
}

func ParseID(res http.ResponseWriter, req *http.Request) (int, bool) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil || id <= 0 {
		// http.Error(res, "Invalid student ID", http.StatusBadRequest)
		 writeJSONError(res, http.StatusBadRequest, "Invalid student ID")
		return 0, false
	}
	return id, true
}


func WriteJSONError(res http.ResponseWriter,status int, message string){
     apiError:=ApiError(Error:message,Status:status)
	 SendJson(res,status,apiError)
}