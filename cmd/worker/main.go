package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"math/big"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/adzialocha/gif-stream-server/s3"

	"github.com/joho/godotenv"
)

const AnimationFrameCount = 10
const OutdatedSessionTreshold = -10 * time.Minute

func init() {
	// Load environment variables.
	godotenv.Load()
}

func makeAnimation(frames []Frame) ([]byte) {
	var images []*image.Paletted
	var delays []int

	for _, frame := range frames {
		// Get base64 string from body
		bodyStr := strings.Replace(
			string(frame.Body[:len(frame.Body)]),
			"data:image/jpeg;base64,",
			"",
			-1,
		)

		// Decode base64 string
		unbased, _ := base64.StdEncoding.DecodeString(bodyStr)

		// Decode to jpeg image
		r := bytes.NewReader(unbased)
		img, _ := jpeg.Decode(r)

		// Make it a paletted image
		bounds := img.Bounds()
		palettedImage := image.NewPaletted(bounds, palette.Plan9)
		draw.Draw(
			palettedImage,
			palettedImage.Rect,
			img,
			bounds.Min,
			draw.Over,
		)

		// Add frame to gif
		images = append(images, palettedImage)
		delays = append(delays, 0)
	}

	// Encode frames to gif
	gifBuffer := new(bytes.Buffer)
	gif.EncodeAll(gifBuffer, &gif.GIF{
		Image: images,
		Delay: delays,
	})

	return gifBuffer.Bytes()
}

func generateReverseTimestamp() (string) {
	// Create a reverse timestamp
	var offset, now, result big.Int
	offset.SetInt64(999999999)
	now.SetInt64(time.Now().Unix())
	result.Sub(&offset, &now)
	return result.String()
}

func makeGifAndUpload(s3Client *s3.S3, frames []Frame) {
	// Create a reverse timestamp
	timestamp := generateReverseTimestamp()

	// Make gif
	gifBytes := makeAnimation(frames)
	gifFileName := "stream/" + frames[0].SessionId + "-" + timestamp + ".gif"

	// Upload to S3
	s3Client.PutObject(
		gifFileName,
		gifBytes,
	)

	// Remove frames from S3
	objectKeys := []string{}
	for _, frame := range frames {
		objectKeys = append(objectKeys, frame.Key)
	}
	s3Client.DeleteObjects(objectKeys)

	fmt.Println("Uploaded " + gifFileName)
}

type Frame struct {
	DateTime string
	Key string
	SessionId string
	Body []byte
}

type ByDateTime []Frame

func (s ByDateTime) Len() int {
    return len(s)
}

func (s ByDateTime) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByDateTime) Less(i, j int) bool {
    return s[i].DateTime < s[j].DateTime
}

func main() {
	// Log
	nowStr := time.Now().UTC().Format("2006-01-02T15-04-05Z")
	fmt.Println("Start Worker @ " + nowStr)

	// Create S3 client.
	s3Client := s3.New(
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		os.Getenv("AWS_REGION"),
		os.Getenv("AWS_BUCKET_NAME"),
	)

	// Get all currently uploaded frames.
	objectKeys := s3Client.ListAllObjects("frames/")

	// Group all frames by sessionId.
	groupedBySessionId := make(map[string][]Frame)
	for _, key := range objectKeys {
		fileName := strings.Replace(key, ".jpg", "", -1)

		if (fileName != "frames/") {
			fileName := strings.Replace(fileName, "frames/", "", -1)
			splitted := strings.Split(fileName, "_")

			groupedBySessionId[splitted[0]] = append(groupedBySessionId[splitted[0]], Frame{
				DateTime: splitted[1],
				Key: key,
				SessionId: splitted[0],
			})
			sort.Sort(ByDateTime(groupedBySessionId[splitted[0]]))
		}
	}

	timeTreshold := time.Now().UTC().Add(OutdatedSessionTreshold).Format("2006-01-02T15-04-05Z")
	timeTresholdStr := strings.Replace(timeTreshold, "-", "", -1)

	fmt.Printf("Found objects: %d\n", len(objectKeys))
	fmt.Printf("Found sessions: %d\n", len(groupedBySessionId))

	// Group them again by slices of 10 and make an animation, when possible.
	for sessionId := range groupedBySessionId {
		var nextAnimation []Frame
		for _, objectKey := range groupedBySessionId[sessionId] {
			if (len(nextAnimation) < AnimationFrameCount) {
				body, err := s3Client.GetObjectBytes(objectKey.Key)
				if err == nil {
					objectKey.Body = body
					nextAnimation = append(nextAnimation, objectKey)
				}
			}

			if (len(nextAnimation) == AnimationFrameCount) {
				makeGifAndUpload(s3Client, nextAnimation)
				nextAnimation = nextAnimation[:0]
			}
		}

		if (len(nextAnimation) > 0) {
			// Not enough frames, but maybe outdated?
			lastFrame := nextAnimation[len(nextAnimation) - 1]
			if (lastFrame.DateTime < timeTresholdStr) {
				makeGifAndUpload(s3Client, nextAnimation)
			}
		}
	}
}
