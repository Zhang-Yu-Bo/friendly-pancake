package utility

var CssStyle = map[string]string{
	"TomorrowNight-dark":   "TomorrowNight-dark",
	"Coy-light":            "Coy-light",
	"Default-light":        "Default-light",
	"Funky-dark":           "Funky-dark",
	"Okaidia-dark":         "Okaidia-dark",
	"SolarizedLight-light": "SolarizedLight-light",
	"Twilight-dark":        "Twilight-dark",
}

const (
	BackgroundColor  = "#2885D3"
	ContainerColor   = "#151718"
	ContainerWidth   = "700px"
	FontSize         = "18px"
	DefaultCssStyle  = "TomorrowNight-dark"
	DefaultFontStyle = "fontsFace"
)

const DefaultCode = `<span class="token macro property"><span class="token directive-hash">#</span><span class="token directive keyword">include</span> <span class="token string">&lt;iostream></span></span>
<span class="token macro property"><span class="token directive-hash">#</span><span class="token directive keyword">include</span><span class="token string">&lt;vector></span></span>

<span class="token keyword">int</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token comment">// 這是註解, 當你看到這個則訊息時代表程式出現了問題, 請通知開發人員, 謝謝</span>
	std<span class="token double-colon punctuation">::</span>vector<span class="token operator">&lt;</span><span class="token keyword">int</span><span class="token operator">></span> arr <span class="token operator">=</span> std<span class="token double-colon punctuation">::</span><span class="token generic-function"><span class="token function">vector</span><span class="token generic class-name"><span class="token operator">&lt;</span><span class="token keyword">int</span><span class="token operator">></span></span></span><span class="token punctuation">(</span><span class="token number">0</span><span class="token punctuation">,</span> <span class="token number">10</span><span class="token punctuation">)</span><span class="token punctuation">;</span>
	<span class="token keyword">for</span> <span class="token punctuation">(</span><span class="token keyword">int</span> i <span class="token operator">=</span> <span class="token number">0</span><span class="token punctuation">;</span> i <span class="token operator">&lt;=</span> <span class="token number">10</span><span class="token punctuation">;</span> i<span class="token operator">++</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		arr<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token operator">=</span> i <span class="token operator">*</span> <span class="token number">10</span> <span class="token operator">+</span> <span class="token number">2</span><span class="token punctuation">;</span>
	<span class="token punctuation">}</span>
	std<span class="token double-colon punctuation">::</span>cout <span class="token operator">&lt;&lt;</span> <span class="token string">"Hello World 你好世界"</span> <span class="token operator">&lt;&lt;</span> std<span class="token double-colon punctuation">::</span>endl<span class="token punctuation">;</span>
	<span class="token keyword">return</span> <span class="token number">0</span><span class="token punctuation">;</span>
<span class="token punctuation">}</span>`

const (
	MaxCodeLength = 1700
	GasURL        = "https://script.google.com/macros/s/AKfycbz4qUdDNEzf-Gswd1Nhw3hmYdFiXQRy8bJ6CZymG71FHKdjge7W-drkwLRj81cY54x0ZA/exec"
)
