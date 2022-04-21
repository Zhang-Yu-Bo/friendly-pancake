// const SupportLanguage = "markup,css,clike,javascript,abap,abnf,actionscript,ada,agda,al,antlr4,apacheconf,apex,apl,applescript,aql,arduino,arff,armasm,arturo,asciidoc,aspnet,asm6502,asmatmel,autohotkey,autoit,avisynth,avro-idl,awk,bash,basic,batch,bbcode,bicep,birb,bison,bnf,brainfuck,brightscript,bro,bsl,c,csharp,cpp,cfscript,chaiscript,cil,clojure,cmake,cobol,coffeescript,concurnas,csp,cooklang,coq,crystal,css-extras,csv,cue,cypher,d,dart,dataweave,dax,dhall,diff,django,dns-zone-file,docker,dot,ebnf,editorconfig,eiffel,ejs,elixir,elm,etlua,erb,erlang,excel-formula,fsharp,factor,false,firestore-security-rules,flow,fortran,ftl,gml,gap,gcode,gdscript,gedcom,gettext,gherkin,git,glsl,gn,linker-script,go,go-module,graphql,groovy,haml,handlebars,haskell,haxe,hcl,hlsl,hoon,http,hpkp,hsts,ichigojam,icon,icu-message-format,idris,ignore,inform7,ini,io,j,java,javadoc,javadoclike,javastacktrace,jexl,jolie,jq,jsdoc,js-extras,json,json5,jsonp,jsstacktrace,js-templates,julia,keepalived,keyman,kotlin,kumir,kusto,latex,latte,less,lilypond,liquid,lisp,livescript,llvm,log,lolcode,lua,magma,makefile,markdown,markup-templating,mata,matlab,maxscript,mel,mermaid,mizar,mongodb,monkey,moonscript,n1ql,n4js,nand2tetris-hdl,naniscript,nasm,neon,nevod,nginx,nim,nix,nsis,objectivec,ocaml,odin,opencl,openqasm,oz,parigp,parser,pascal,pascaligo,psl,pcaxis,peoplecode,perl,php,phpdoc,php-extras,plant-uml,plsql,powerquery,powershell,processing,prolog,promql,properties,protobuf,pug,puppet,pure,purebasic,purescript,python,qsharp,q,qml,qore,r,racket,cshtml,jsx,tsx,reason,regex,rego,renpy,rescript,rest,rip,roboconf,robotframework,ruby,rust,sas,sass,scss,scala,scheme,shell-session,smali,smalltalk,smarty,sml,solidity,solution-file,soy,sparql,splunk-spl,sqf,sql,squirrel,stan,stata,iecst,stylus,supercollider,swift,systemd,t4-templating,t4-cs,t4-vb,tap,tcl,tt2,textile,toml,tremor,turtle,twig,typescript,typoscript,unrealscript,uorazor,uri,v,vala,vbnet,velocity,verilog,vhdl,vim,visual-basic,warpscript,wasm,web-idl,wiki,wolfram,wren,xeora,xml-doc,xojo,xquery,yaml,yang,zig";
export const SupportLanguage = ["markup", "css", "clike", "javascript", "abap", "abnf", "actionscript", "ada", "agda", "al", "antlr4", "apacheconf", "apex", "apl", "applescript", "aql", "arduino", "arff", "armasm", "arturo", "asciidoc", "aspnet", "asm6502", "asmatmel", "autohotkey", "autoit", "avisynth", "avro-idl", "awk", "bash", "basic", "batch", "bbcode", "bicep", "birb", "bison", "bnf", "brainfuck", "brightscript", "bro", "bsl", "c", "csharp", "cpp", "cfscript", "chaiscript", "cil", "clojure", "cmake", "cobol", "coffeescript", "concurnas", "csp", "cooklang", "coq", "crystal", "css-extras", "csv", "cue", "cypher", "d", "dart", "dataweave", "dax", "dhall", "diff", "django", "dns-zone-file", "docker", "dot", "ebnf", "editorconfig", "eiffel", "ejs", "elixir", "elm", "etlua", "erb", "erlang", "excel-formula", "fsharp", "factor", "false", "firestore-security-rules", "flow", "fortran", "ftl", "gml", "gap", "gcode", "gdscript", "gedcom", "gettext", "gherkin", "git", "glsl", "gn", "linker-script", "go", "go-module", "graphql", "groovy", "haml", "handlebars", "haskell", "haxe", "hcl", "hlsl", "hoon", "http", "hpkp", "hsts", "ichigojam", "icon", "icu-message-format", "idris", "ignore", "inform7", "ini", "io", "j", "java", "javadoc", "javadoclike", "javastacktrace", "jexl", "jolie", "jq", "jsdoc", "js-extras", "json", "json5", "jsonp", "jsstacktrace", "js-templates", "julia", "keepalived", "keyman", "kotlin", "kumir", "kusto", "latex", "latte", "less", "lilypond", "liquid", "lisp", "livescript", "llvm", "log", "lolcode", "lua", "magma", "makefile", "markdown", "markup-templating", "mata", "matlab", "maxscript", "mel", "mermaid", "mizar", "mongodb", "monkey", "moonscript", "n1ql", "n4js", "nand2tetris-hdl", "naniscript", "nasm", "neon", "nevod", "nginx", "nim", "nix", "nsis", "objectivec", "ocaml", "odin", "opencl", "openqasm", "oz", "parigp", "parser", "pascal", "pascaligo", "psl", "pcaxis", "peoplecode", "perl", "php", "phpdoc", "php-extras", "plant-uml", "plsql", "powerquery", "powershell", "processing", "prolog", "promql", "properties", "protobuf", "pug", "puppet", "pure", "purebasic", "purescript", "python", "qsharp", "q", "qml", "qore", "r", "racket", "cshtml", "jsx", "tsx", "reason", "regex", "rego", "renpy", "rescript", "rest", "rip", "roboconf", "robotframework", "ruby", "rust", "sas", "sass", "scss", "scala", "scheme", "shell-session", "smali", "smalltalk", "smarty", "sml", "solidity", "solution-file", "soy", "sparql", "splunk-spl", "sqf", "sql", "squirrel", "stan", "stata", "iecst", "stylus", "supercollider", "swift", "systemd", "t4-templating", "t4-cs", "t4-vb", "tap", "tcl", "tt2", "textile", "toml", "tremor", "turtle", "twig", "typescript", "typoscript", "unrealscript", "uorazor", "uri", "v", "vala", "vbnet", "velocity", "verilog", "vhdl", "vim", "visual-basic", "warpscript", "wasm", "web-idl", "wiki", "wolfram", "wren", "xeora", "xml-doc", "xojo", "xquery", "yaml", "yang", "zig"];

export const DefaultCodeText = `package main
import "fmt"
func main() {
    fmt.Println("Hello World, 你好世界!🌎🌏🌍")
}`;

export const CodeStyleMap = {
	"TomorrowNight-dark":   	"prismjs/TomorrowNight-dark.css",
	"Coy-light":				"prismjs/Coy-light.css",
	"Default-light":			"prismjs/Default-light.css",
	"Funky-dark":           	"prismjs/Funky-dark.css",
	"Okaidia-dark":         	"prismjs/Okaidia-dark.css",
	"SolarizedLight-light": 	"prismjs/SolarizedLight-light.css",
	"Twilight-dark":        	"prismjs/Twilight-dark.css",
};

export const IdList = [
	"codeLangSelectpicker",
	"containerWidthInput",
	"fontSizeInput",
	"cssStyleSelector",
	"backgroundColorInput",
	"containerColorInput",
	"editPreviewBtn",
	"generateImageBtn",
	"mCodeStyle",
	"mCode",
	"mPre",
	"mCodeText",
	"mContainer",
	"mBackground",
	"editSection",
	"previewSection",
    "cleanCatchBtn",
	"resultModal",
	"modalLoadingSpinner",
	"resultURLText",
	"resultBBCodeText",
	"resutlErrorText",
	"modalResultSuccess",
	"modalResultFailed",
];

export const DefaultCodeLang = "go";
export const DefaultFontSize = 18;
export const DefaultCssStyle = "TomorrowNight-dark";
export const DefaultBackgroundColor = "#2885D3";
export const DefaultContainerColor = "#151718";

export const RequestHost = "https://friendly-pancake.herokuapp.com";
// export const RequestHost = "http://localhost"
// https://friendly-pancake.herokuapp.com/raw/code/image?code=1343c66f3369b2429563afa5ff180f792a998f95463217a7540c9c3c6735b609&file=image.png
