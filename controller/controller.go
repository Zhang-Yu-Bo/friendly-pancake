package controller

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/Zhang-Yu-Bo/friendly-pancake/model/templatePage"
	wk "github.com/Zhang-Yu-Bo/friendly-pancake/model/wkhtmltoimage"
)

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

	binPath := r.URL.Query().Get("bin_path")

	if len(binPath) == 0 {
		if runtime.GOOS == "linux" {
			binPath = "./bin/wkhtmltoimage"
		} else {
			binPath = "C:\\Users\\Lykoi\\Desktop\\html2image-master\\wkhtmltopdf\\bin\\wkhtmltoimage.exe"
		}
	}

	html := `
		<!DOCTYPE html>
<html>
    <head>
        <style>
            body {
                background-color: #2885D3;
                /* background-color: rgba(0,0,0,0.2); */
            }
            
            h1 {
                color: white;
                text-align: center;
            }
            
            p {
                color: white;
                font-family: verdana;
                font-size: 20px;
            }

            .container {
                margin: auto;
                background-color: rgb(21, 23, 24);
                /* #282c34 */
                border-radius: 10px;
                width: 300px;
                padding: 12px;
                box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
				word-break: break-word;
            }
            .button {
                width: 12px;
                height: 12px;
                border-radius: 50%;
                margin-right: 5px;
                display: inline-block;
            }
            .title-bar {
                padding-left: 6px;
            }
        </style>
    </head>
    <body>
        <div class='container'>
            <div class='title-bar'>
                <span class='button' style='background-color: #ff5f56;'></span>
                <span class='button' style='background-color: #ffbd2e;'></span>
                <span class='button' style='background-color: #27c93f;'></span>
            </div>
            <h1>Hi mom</h1>
            <p>This is a paragraph.</p>
        </div>
    </body>
</html>`

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

	data := templatePage.CodePage{
		Code: `<span class="token macro property"><span class="token directive-hash">#</span><span class="token directive keyword">include</span> <span class="token string">&lt;iostream></span></span>
<span class="token macro property"><span class="token directive-hash">#</span><span class="token directive keyword">include</span><span class="token string">&lt;vector></span></span>

<span class="token keyword">int</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token comment">// 這是註解</span>
	std<span class="token double-colon punctuation">::</span>vector<span class="token operator">&lt;</span><span class="token keyword">int</span><span class="token operator">></span> arr <span class="token operator">=</span> std<span class="token double-colon punctuation">::</span><span class="token generic-function"><span class="token function">vector</span><span class="token generic class-name"><span class="token operator">&lt;</span><span class="token keyword">int</span><span class="token operator">></span></span></span><span class="token punctuation">(</span><span class="token number">0</span><span class="token punctuation">,</span> <span class="token number">10</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
	<span class="token keyword">for</span> <span class="token punctuation">(</span><span class="token keyword">int</span> i <span class="token operator">=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;</span> <span class="token number">10</span><span class="token punctuation">;</span> i<span class="token operator">++</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		arr<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token operator">=</span> i <span class="token operator">*</span> <span class="token number">10</span> <span class="token operator">+</span> <span class="token number">2</span><span class="token punctuation">;</span>
	<span class="token punctuation">}</span>
	std<span class="token double-colon punctuation">::</span>cout <span class="token operator">&lt;&lt;</span> <span class="token string">"Hello World 你好世界"</span> <span class="token operator">&lt;&lt;</span> std<span class="token double-colon punctuation">::</span>endl<span class="token punctuation">;</span>
	<span class="token keyword">return</span> <span class="token number">0</span><span class="token punctuation">;</span>
<span class="token punctuation">}</span>`,
		BackgroundColor: "#2885D3",
		ContainerColor:  "#151718",
		ContainerWidth:  "700px",
		FontSize:        "18px",
	}

	w.Write([]byte(templatePage.Parse(data)))
}
