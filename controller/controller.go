package controller

import (
	"fmt"
	// "image"
	// "image/color"
	// "image/png"
	"net/http"
	"runtime"

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

	// fileName := r.URL.Query().Get("file_name")

	// if len(fileName) == 0 {
	// 	fileName = "image.png"
	// }

	// width := 200
	// height := 100

	// upLeft := image.Point{0, 0}
	// lowRight := image.Point{width, height}

	// img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// // Colors are defined by Red, Green, Blue, Alpha uint8 values.
	// cyan := color.RGBA{100, 200, 200, 0xff}

	// // Set color for each pixel.
	// for x := 0; x < width; x++ {
	// 	for y := 0; y < height; y++ {
	// 		switch {
	// 		case x < width/2 && y < height/2: // upper left quadrant
	// 			img.Set(x, y, cyan)
	// 		case x >= width/2 && y >= height/2: // lower right quadrant
	// 			img.Set(x, y, color.White)
	// 		default:
	// 			// Use zero value.
	// 		}
	// 	}
	// }
	// png.Encode(w, img)
	// // fmt.Fprintf(w, "%s", fileName)

	binPath := "C:\\Users\\Lykoi\\Desktop\\html2image-master\\wkhtmltopdf\\bin\\wkhtmltoimage.exe"

	if runtime.GOOS == "linux" {
		binPath = "~/bin/wkhtmltoimage"
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
