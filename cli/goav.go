package main

import (
	"log"
	"time"

	"github.com/giorgisio/goav/avcodec"

	"github.com/giorgisio/goav/avformat"
)

func main() {

	filename := "/tmp/simple.mp4"

	// Register all formats and codecs
	avformat.AvRegisterAll()

	ctx := avformat.AvformatAllocContext()

	// Open video file
	if avformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}

	d := time.Duration(ctx.Duration() * 1000)
	log.Println("heihieihei.", d)
	log.Println("NbStreams: ", ctx.NbStreams())
	{
		t := ctx.Streams()[0].Codec().GetCodecId()
		log.Println("Streams()[0].Codec().Type()", t)
		d := avcodec.AvcodecFindDecoder(avcodec.CodecId(t))
		log.Println("aa", d)
	}
	// avformat.MediaType
	// avutil.MediaType
	// avcodec.MediaType
}
