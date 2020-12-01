package hello

import (
	"fmt"
	"net/http"
)

/*
@Author kim
@Description   网络编程
@date 2020-11-24 11:15
*/

//get方法服务接口
func Get() {

	http.HandleFunc("/my/go", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL.RawQuery)
		fmt.Println(writer)
		writer.Write([]byte("myGoHttp"))
	})
	http.ListenAndServe("127.0.0.1:8888", nil)

}
