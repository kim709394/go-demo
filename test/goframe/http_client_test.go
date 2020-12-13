package goframe

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"testing"
)

/*
@Author kim
@Description  http client
@date 2020-12-13 18:56
*/

//http客户端实例化
func TestHttpClientInit(t *testing.T) {
	//http客户端创建
	c1 := ghttp.NewClient()
	//推荐使用
	c2 := g.Client()
	fmt.Println(c1, c2)
	/*
		type Client
		    func NewClient() *Client
		    func (c *Client) Get(url string, data ...interface{}) (*ClientResponse, error)
		    func (c *Client) Put(url string, data ...interface{}) (*ClientResponse, error)
		    func (c *Client) Post(url string, data ...interface{}) (*ClientResponse, error)
		    func (c *Client) Delete(url string, data ...interface{}) (*ClientResponse, error)
		    func (c *Client) Connect(url string, data ...interface{}) (*ClientResponse, error)
		    func (c *Client) Head(url string, data ...interface{}) (*ClientResponse, error)
		    func (c *Client) Options(url string, data ...interface{}) (*ClientResponse, error)
		    func (c *Client) Patch(url string, data ...interface{}) (*ClientResponse, error)
		    func (c *Client) Trace(url string, data ...interface{}) (*ClientResponse, error)
		    func (c *Client) DoRequest(method, url string, data ...interface{}) (*ClientResponse, error)

		    func (c *Client) GetBytes(url string, data ...interface{}) []byte
		    func (c *Client) PutBytes(url string, data ...interface{}) []byte
		    func (c *Client) PostBytes(url string, data ...interface{}) []byte
		    func (c *Client) DeleteBytes(url string, data ...interface{}) []byte
		    func (c *Client) ConnectBytes(url string, data ...interface{}) []byte
		    func (c *Client) HeadBytes(url string, data ...interface{}) []byte
		    func (c *Client) OptionsBytes(url string, data ...interface{}) []byte
		    func (c *Client) PatchBytes(url string, data ...interface{}) []byte
		    func (c *Client) TraceBytes(url string, data ...interface{}) []byte
		    func (c *Client) RequestBytes(method string, url string, data ...interface{}) []byte

		    func (c *Client) GetContent(url string, data ...interface{}) string
		    func (c *Client) PutContent(url string, data ...interface{}) string
		    func (c *Client) PostContent(url string, data ...interface{}) string
		    func (c *Client) DeleteContent(url string, data ...interface{}) string
		    func (c *Client) ConnectContent(url string, data ...interface{}) string
		    func (c *Client) HeadContent(url string, data ...interface{}) string
		    func (c *Client) OptionsContent(url string, data ...interface{}) string
		    func (c *Client) PatchContent(url string, data ...interface{}) string
		    func (c *Client) TraceContent(url string, data ...interface{}) string
		    func (c *Client) RequestContent(method string, url string, data ...interface{}) string

		    func (c *Client) GetVar(url string, data ...interface{}) *gvar.Var
		    func (c *Client) PutVar(url string, data ...interface{}) *gvar.Var
		    func (c *Client) PostVar(url string, data ...interface{}) *gvar.Var
		    func (c *Client) DeleteVar(url string, data ...interface{}) *gvar.Var
		    func (c *Client) HeadVar(url string, data ...interface{}) *gvar.Var
		    func (c *Client) PatchVar(url string, data ...interface{}) *gvar.Var
		    func (c *Client) ConnectVar(url string, data ...interface{}) *gvar.Var
		    func (c *Client) OptionsVar(url string, data ...interface{}) *gvar.Var
		    func (c *Client) TraceVar(url string, data ...interface{}) *gvar.Var
		    func (c *Client) RequestVar(method string, url string, data ...interface{}) *gvar.Var

		    func (c *Client) SetBasicAuth(user, pass string) *Client
		    func (c *Client) SetBrowserMode(enabled bool) *Client
		    func (c *Client) SetContentType(contentType string) *Client
		    func (c *Client) SetCookie(key, value string) *Client
		    func (c *Client) SetCookieMap(m map[string]string) *Client
		    func (c *Client) SetCtx(ctx context.Context) *Client
		    func (c *Client) SetHeader(key, value string) *Client
		    func (c *Client) SetHeaderMap(m map[string]string) *Client
		    func (c *Client) SetHeaderRaw(headers string) *Client
		    func (c *Client) SetPrefix(prefix string) *Client
		    func (c *Client) SetRetry(retryCount int, retryInterval int) *Client
		    func (c *Client) SetTimeout(t time.Duration) *Client
		    func (c *Client) SetProxy(proxyURL string)
	*/

	/*
		简要说明：

		我们可以使用NewClient创建一个自定义的HTTP客户端对象Client，随后可以使用该对象执行请求，该对象底层使用了连接池设计，
		因此没有Close关闭方法。HTTP客户端对象也可以通过g.Client()快捷方法创建，该方式创建的客户端对象为单例对象。
		客户端提供了一系列以HTTP Method命名的方法，调用这些方法将会发起对应的HTTP Method请求。
		常用的方法是Get和Post方法，同时DoRequest是核心的请求方法，用户可以调用该方法实现自定义的HTTP Method发送请求。
		请求返回结果为*ClientResponse对象，可以通过该结果对象获取对应的返回结果，
		通过ReadAll/ReadAllString方法可以获得返回的内容，该对象在使用完毕后需要通过Close方法关闭，防止内存溢出。
		*Bytes方法用于获得服务端返回的二进制数据，如果请求失败返回nil；*Content方法用于请求获得字符串结果数据，如果请求失败返回空字符串；Set*方法用于Client的参数设置。
		*Var方法直接请求并获取HTTP接口结果为泛型类型便于转换。如果请求失败或者请求结果为空，会返回一个空的g.Var泛型对象，不影响转换方法调用。
		可以看到，客户端的请求参数的数据参数data数据类型为interface{}类型，也就是说可以传递任意的数据类型，常见的参数数据类型为string/map，如果参数为map类型，参数值将会被自动urlencode编码。
	*/

	/*
		ghttp.ClientResponse为HTTP对应请求的返回结果对象，该对象继承于http.Response，可以使用http.Response的所有方法。在此基础之上增加了以下几个方法：
		func (r *ClientResponse) GetCookie(key string) string
		func (r *ClientResponse) GetCookieMap() map[string]string
		func (r *ClientResponse) Raw() string
		func (r *ClientResponse) RawDump()
		func (r *ClientResponse) RawRequest() string
		func (r *ClientResponse) RawResponse() string
		func (r *ClientResponse) ReadAll() []byte
		func (r *ClientResponse) ReadAllString() string
		func (r *ClientResponse) Close() error

		这里也要提醒的是，ClientResponse需要手动调用Close方法关闭，也就是说，不管你使用不使用返回的ClientResponse对象，
		你都需要将该返回对象赋值给一个变量，并且手动调用其Close方法进行关闭（往往使用defer r.Close()），否则会造成文件句柄溢出、内存溢出。
	*/

	/*
		重要说明
		ghttp客户端默认关闭了KeepAlive功能以及对服务端TLS证书的校验功能，如果需要启用可自定义客户端的Transport属性。
		连接池参数设定、连接代理设置这些高级功能也可以通过自定义客户端的Transport属性实现，该数据继承于标准库的http.Transport对象。
	*/
}

//调用示例
func TestHttpClientExample(t *testing.T) {
	/*
		func (c *Client) Ctx(ctx context.Context) *Client   用于设置当前请求的上下文对象context.Context。
		func (c *Client) Timeout(t time.Duration) *Client   用于设置当前请求超时时间
		func (c *Client) Cookie(m map[string]string) *Client   用于设置当前请求的自定义Cookie信息。
		func (c *Client) Header(m map[string]string) *Client   方法用于设置当前请求的自定义Header信息。
		func (c *Client) HeaderRaw(headers string) *Client
		func (c *Client) ContentType(contentType string) *Client   用于设置当前请求的Content-Type信息，并且支持根据该信息自动检查提交参数并自动编码。
		func (c *Client) ContentJson() *Client
		func (c *Client) ContentXml() *Client
		func (c *Client) BasicAuth(user, pass string) *Client   用于设置HTTP Basic Auth校验信息。
		func (c *Client) Retry(retryCount int, retryInterval time.Duration) *Client   用于设置请求失败时重连次数和重连间隔。
		func (c *Client) Proxy(proxyURL string) *Client    用于设置http访问代理。
		func (c *Client) RedirectLimit(redirectLimit int) *Client   用于限制重定向跳转次数。
	*/
	client := g.Client()
	//get请求
	res1 := client.
		Timeout(3000).                                       //设置请求超时时间
		Cookie(map[string]string{"sessionId": "kim709394"}). //设置cookie
		Header(map[string]string{"head": "head"}).           //设置head
		HeaderRaw(`       //字符串键值对形式设置head
			Referer: https://goframe.org/
			User-Agent: MyTesyClient`).GetContent("http://localhost:8081/get?p=123")
	fmt.Println(res1)
	//post请求
	res2 := client.PostContent("http://localhost:8081/post", g.Map{
		"k1": "v1",
		"k2": "v2",
	})
	fmt.Println(res2)
	//json请求,该请求将会将Content-Type设置为application/json，并且将提交参数自动序列化为Json格式
	//{"k1":"v1","k2":"v2"}
	res3 := client.ContentJson().PostContent("http://localhost:8081/json", g.Map{
		"k1": "v1",
		"k2": "v2",
	})
	fmt.Println(res3)
	//xml请求,该请求将会将Content-Type设置为application/xml，并且将提交参数自动编码为Xml
	//<doc><k1>v1</k1><k2>v2</k2></doc>
	res4 := client.ContentXml().PostContent("http://localhost:8081/xml", g.Map{
		"k1": "v1",
		"k2": "v2",
	})
	fmt.Println(res4)
	//普通get请求
	if r, err := g.Client().Get("http://localhost:8081/get"); err != nil {
		panic(err)
	} else {
		defer r.Close()
		fmt.Println(r.ReadAllString())
	}
	//文件下载,大文件下载需要分批返回,因为文件会全部加载到内存
	if r, err := g.Client().Get("http://localhost:8081/download"); err != nil {
		panic(err)
	} else {
		defer r.Close()
		gfile.PutBytes("../../file/logo.png", r.ReadAll())
	}
	//post请求
	if r, err := g.Client().Post("http://localhost:8081/form", "name=john&age=18"); err != nil {
		panic(err)
	} else {
		defer r.Close()
		fmt.Println(r.ReadAllString())
	}

}

//文件上传
//单文件上传
func TestSingleFileUpload(t *testing.T) {
	//参数列表:  url,  参数名=@file:文件路径
	r, err := ghttp.Post("http://localhost:81/single", "file=@file:"+"../../file/addWorkBill.png")
	defer r.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

//多文件上传
func TestMultipleFileUpload(t *testing.T) {

	//参数列表:  url,  参数名=@file:文件路径&参数名=@file:文件路径   也可用:  参数名[]=@file:文件路径&参数名[]=@file:文件路径
	r, err := ghttp.Post("http://localhost:81/multiple", "file=@file:"+"../../file/addWorkBill.png"+"&file=@file:"+"../../file/writ.txt")
	defer r.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)

}

//日志打印
func TestLog(t *testing.T) {

	res, err := g.Client().Get("http://localhost:8081/get")
	defer res.Close()
	if err != nil {
		fmt.Println(err)
	}
	//打印请求和返回信息日志
	res.RawDump()

}
