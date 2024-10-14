﻿package main

import (
	"fmt"
	"log"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	inp_dir := "../videos/raw/"
	files, err := os.ReadDir(inp_dir)
	if err != nil {
		log.Fatal(err)
	}

	var inp_path string
	out_path := "../videos/compressed/"

	for ind, file := range files {
		inp_path = inp_dir + file.Name()
		err = ffmpeg.Input(inp_path).Output(out_path+file.Name(), ffmpeg.KwArgs{"an": "", "c": "copy"}).
			OverWriteOutput().
			Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ind, file.Name()+"is done compressing")
	}

	fmt.Println("Done compressing all vidoes")
}
