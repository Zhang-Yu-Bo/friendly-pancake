package controller

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/templatePage"
	wk "github.com/Zhang-Yu-Bo/friendly-pancake/model/wkhtmltoimage"
)

var tempCode = `<span class="token macro property"><span class="token directive-hash">#</span><span class="token directive keyword">include</span> <span class="token string">&lt;iostream></span></span>
<span class="token macro property"><span class="token directive-hash">#</span><span class="token directive keyword">include</span><span class="token string">&lt;vector></span></span>

<span class="token keyword">int</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token comment">// 這是註解</span>
	std<span class="token double-colon punctuation">::</span>vector<span class="token operator">&lt;</span><span class="token keyword">int</span><span class="token operator">></span> arr <span class="token operator">=</span> std<span class="token double-colon punctuation">::</span><span class="token generic-function"><span class="token function">vector</span><span class="token generic class-name"><span class="token operator">&lt;</span><span class="token keyword">int</span><span class="token operator">></span></span></span><span class="token punctuation">(</span><span class="token number">0</span><span class="token punctuation">,</span> <span class="token number">10</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
	<span class="token keyword">for</span> <span class="token punctuation">(</span><span class="token keyword">int</span> i <span class="token operator">=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;=</span> <span class="token number">10</span><span class="token punctuation">;</span> i<span class="token operator">++</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		arr<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token operator">=</span> i <span class="token operator">*</span> <span class="token number">10</span> <span class="token operator">+</span> <span class="token number">2</span><span class="token punctuation">;</span>
	<span class="token punctuation">}</span>
	std<span class="token double-colon punctuation">::</span>cout <span class="token operator">&lt;&lt;</span> <span class="token string">"Hello World 你好世界"</span> <span class="token operator">&lt;&lt;</span> std<span class="token double-colon punctuation">::</span>endl<span class="token punctuation">;</span>
	<span class="token keyword">return</span> <span class="token number">0</span><span class="token punctuation">;</span>
<span class="token punctuation">}</span>`

func HomePage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		name = "stranger"
	}

	fmt.Fprintf(w, "Hello, %s\n", name)
	fmt.Fprintf(w, "OS: %s\n", runtime.GOOS)
	fmt.Fprintf(w, "Max Process: %d\n", runtime.GOMAXPROCS(0))
	fmt.Fprintf(w, "Your IP is: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "Forwarded for: %s\n", r.Header.Get("X-FORWARDED-FOR"))
}

func RawImage(w http.ResponseWriter, r *http.Request) {

	binPath := r.URL.Query().Get("binPath")

	if len(binPath) == 0 {
		if runtime.GOOS == "linux" {
			binPath = "./bin/wkhtmltoimage"
		} else {
			binPath = "C:\\Users\\Lykoi\\Desktop\\html2image-master\\wkhtmltopdf\\bin\\wkhtmltoimage.exe"
		}
	}

	if len(templatePage.Page) == 0 {
		fmt.Fprintln(w, "Code template is nil")
		return
	}

	backgroundColor := r.URL.Query().Get("backgroundColor")
	if backgroundColor == "" {
		backgroundColor = "#2885D3"
	}
	containerColor := r.URL.Query().Get("containerColor")
	if containerColor == "" {
		containerColor = "#151718"
	}
	containerWidth := r.URL.Query().Get("containerWidth")
	if containerWidth == "" {
		containerWidth = "700px"
	}
	fontSize := r.URL.Query().Get("fontSize")
	if fontSize == "" {
		fontSize = "18px"
	}

	cssUrl := "http://localhost/static/prism.css"
	if runtime.GOOS == "linux" {
		cssUrl = "https://friendly-pancake.herokuapp.com/static/prism.css"
	}

	data := templatePage.CodePage{
		CssUrl:          cssUrl,
		Code:            strings.ReplaceAll(tempCode, "\t", "    "),
		BackgroundColor: backgroundColor,
		ContainerColor:  containerColor,
		ContainerWidth:  containerWidth,
		FontSize:        fontSize,
	}

	html := templatePage.Parse(data)

	c := wk.ImageOptions{
		BinaryPath: binPath,
		Input:      "-",
		HTML:       html,
		Format:     "png",
	}

	if out, err := wk.GenerateImage(&c); err != nil {
		fmt.Fprintf(w, "Error: %s\n", err.Error())
	} else {
		w.Write(out)
	}

}

func TestPage(w http.ResponseWriter, r *http.Request) {

	if len(templatePage.Page) == 0 {
		fmt.Fprintln(w, "Code template is nil")
		return
	}

	backgroundColor := r.URL.Query().Get("backgroundColor")
	if backgroundColor == "" {
		backgroundColor = "#2885D3"
	}
	containerColor := r.URL.Query().Get("containerColor")
	if containerColor == "" {
		containerColor = "#151718"
	}
	containerWidth := r.URL.Query().Get("containerWidth")
	if containerWidth == "" {
		containerWidth = "700px"
	}
	fontSize := r.URL.Query().Get("fontSize")
	if fontSize == "" {
		fontSize = "18px"
	}

	cssUrl := "http://localhost/static/prism.css"
	if runtime.GOOS == "linux" {
		cssUrl = "https://friendly-pancake.herokuapp.com/static/prism.css"
	}

	data := templatePage.CodePage{
		CssUrl:          cssUrl,
		Code:            strings.ReplaceAll(tempCode, "\t", "    "),
		BackgroundColor: backgroundColor,
		ContainerColor:  containerColor,
		ContainerWidth:  containerWidth,
		FontSize:        fontSize,
	}

	w.Write([]byte(templatePage.Parse(data)))
}
