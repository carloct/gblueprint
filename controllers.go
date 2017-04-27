package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/gif"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func createGif(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseMultipartForm(1024)

	outImg := &gif.GIF{}

	for _, fheaders := range r.MultipartForm.File {
		for _, hdr := range fheaders {
			//var infile multipart.File
			infile, err := hdr.Open()
			if nil != err {
				http.Error(w, "Error", http.StatusInternalServerError)
				return
			}

			inImg, _ := gif.Decode(infile)
			infile.Close()

			outImg.Image = append(outImg.Image, inImg.(*image.Paletted))
			outImg.Delay = append(outImg.Delay, 1)

		}

	}

	buffer := new(bytes.Buffer)
	if err := gif.EncodeAll(buffer, outImg); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "image/gif")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	w.Write(buffer.Bytes())
}

func createGifFromVideo(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseMultipartForm(1024)

	upload, header, err := r.FormFile("video")
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}

	ex, err := os.Executable()
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}

	exPath := path.Dir(ex)

	var stringBuf bytes.Buffer
	stringBuf.WriteString(exPath)
	stringBuf.WriteString("/tmp/")
	stringBuf.WriteString(header.Filename)

	spew.Dump(stringBuf.String())

	fh, err := os.Create(stringBuf.String())
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}

	defer fh.Close()
	io.Copy(fh, upload)

	//spew.Dump(stringBuf)

	paletteCmd := exec.Command("ffmpeg", "-y", "-t", "5", "-i",
		stringBuf.String(),
		"-vf",
		"fps=10,scale=320:-1:flags=lanczos,palettegen",
		"-f",
		"image2pipe",
		"-vcodec",
		"ppm",
		"-")

	gifCmd := exec.Command("ffmpeg", "-ss", "30", "-t", "5", "-i",
		stringBuf.String(),
		"-i",
		"pipe:0",
		"-filter_complex",
		"fps=10,scale=320:-1:flags=lanczos[x];[x][1:v]paletteuse",
		"/Users/carloc/Dev/go/src/github.com/carloct/gblueprint/gifs/output.gif")

	gifCmd.Stdin, err = paletteCmd.StdoutPipe()
	if err != nil {
		//log.Fatal(err)
		spew.Dump(err)
	}

	var buffer bytes.Buffer
	gifCmd.Stdout = &buffer

	if err = paletteCmd.Start(); err != nil {
		//log.Fatal(err)
		spew.Dump(err)
	}

	if err = gifCmd.Start(); err != nil {
		//log.Fatal(err)
		spew.Dump(err)
	}

	if err = paletteCmd.Wait(); err != nil {
		//log.Fatal(err)
		spew.Dump(err)
	}

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	body, _ := json.Marshal(CurrentVersion)
	w.Header().Set("Content-type", "application/json")
	w.Write(body)
}
