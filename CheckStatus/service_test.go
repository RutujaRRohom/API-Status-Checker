package CheckStatus

import (
	//"net/http"
	
	"testing"
	"time"
	"github.com/stretchr/testify/suite"
	"sync"
	"context"
)




//using suite 
var websitesMap map[string]Website

type checkstatusSuite struct{
	suite.Suite
}

func (c *checkstatusSuite) SetUpsuite(){
	// websitesMap := make(map[string]Website)
	weblist := []string{"www.google.com", "www.facebook.com", "www.amazon.com", "www.fakebook.com"}
	for _, v := range weblist {
		websitesMap[v] = Website{v, "Not Ready"}
	}
}


func (c *checkstatusSuite) TestCheckStatus() {
	w := Website{}

	// Test with a valid website
	status, err := w.CreateStatus(context.Background(), "www.google.com")
	c.NoError(err)
	c.Equal("Not Ready", status)

	// Test with an invalid website
	status, err = w.CreateStatus(context.Background(), "www.invalid.com")
	c.Error(err)
	c.Equal("Key not present.", status)
}

func (s *checkstatusSuite) TestGetStatus() {
	// Test if the status of the websites is being updated correctly
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		GetStatus()
	}()

	wg.Wait()
	time.Sleep(time.Second * 1)
	s.Equal("UP", websitesMap["www.google.com"].Status)
	s.Equal("DOWN", websitesMap["www.fakebook.com"].Status)
	s.Equal("UP", websitesMap["www.amazon.com"].Status)
}

func TestWebsiteSuite(t *testing.T) {
	suite.Run(t, new(checkstatusSuite))
}

//using table driven testing
// func TestGetStatus(t *testing.T){

	//   result := []struct{
	// 	url string
	// 	ExpectedStatus string
	//   }
	//   {
	// 	{"www.google.com","UP"},
	// 	{"www.facebook.com","UP"},
	// 	{"www.dfbwjfj.com","DOWN"},
	//   }
	
	
	//   for _,testcase := range result{
	// 	t.Run("test getstatus",func(t *testing.T){
	// 		req,err := http.NewRequest("GET",result.url,nil)
	// 		if err != nil{
	// 			t.Fatal(err)
	// 		}
	// 		rr:=httptest.NewRecorder()
	// 		handler:=http.HandleFunc(GetStatus())
	
	
	// 	})
	//   }
	
	// }
	

