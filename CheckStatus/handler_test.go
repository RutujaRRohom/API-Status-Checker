package CheckStatus

import
(
	//"fmt"
	"net/http/httptest"
	"net/http"
	"testing"
)


func TestHandler(t *testing.T){
   t.Run("test if POST method returns OK status",func(t *testing.T){
req,err := http.NewRequest("POST","/POST/websites",nil)
if err != nil {
	t.Fatal(err)
}	

rr := httptest.NewRecorder()
//handler := HandleFunc(PostHandler())
handler := http.HandlerFunc(PostHandler(NewFunc()))
handler.ServeHTTP(rr,req)

if status := rr.Code; status != http.StatusOK {
	t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
}
})

}



func TestGethandler(t *testing.T){

t.Run("checks if GET request returns statusBadRequest if wrong request being send",func(t *testing.T){
		req,err := http.NewRequest("GET","/GET/websites",nil)
		if err != nil{
			t.Fatal(err)
		}
		rr:=httptest.NewRecorder()
		//handler := HandleFunc(GetHandler())
		handler := http.HandlerFunc(GetHandler(NewFunc()))
		handler.ServeHTTP(rr,req)
		if status := rr.Code ;status !=http.StatusBadRequest{
			t.Errorf("bad request")
		}

})

}