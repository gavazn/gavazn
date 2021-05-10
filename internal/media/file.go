package media

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
)

// Path model
type Path struct {
	Original string `bson:"original" json:"original"`
	Large    string `bson:"large" json:"large"`
	Medium   string `bson:"medium" json:"medium"`
	Small    string `bson:"small" json:"small"`
}

var supportFormat = []string{
	"image/jpeg",
	"image/png",
	"video/mpeg",
	"video/x-msvideo",
	"video/x-matroska",
	"video/webm",
}

// UploadFile upload files
func UploadFile(file *multipart.FileHeader) (*Path, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	support := false
	for _, format := range supportFormat {
		if format == file.Header.Get("Content-Type") {
			support = true
			break
		}
	}

	if !support {
		return nil, errors.New("format not supported")
	}

	dir := fmt.Sprintf("uploads/%v/%v", time.Now().Year(), strings.ToLower(time.Now().Month().String()))

	if err = os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// check not exists file
	var (
		path      string
		fileName  string
		extension string
	)
	for true {
		rand.Seed(time.Now().UnixNano())
		fileName = strconv.Itoa(rand.Int())
		extension = filepath.Ext(file.Filename)
		path = fmt.Sprintf("%v/%v%v", dir, fileName, extension)

		if _, err := os.Stat(path); err != nil {
			break
		}
	}

	dst, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, f); err != nil {
		return nil, err
	}

	p := &Path{
		Original: path,
	}

	if strings.Split(file.Header.Get("Content-Type"), "/")[0] == "image" {
		// compress to large size
		largePath, err := resizeImage(path, 1280, 0, fileName+"-large"+extension)
		if err != nil {
			return nil, err
		}
		p.Large = largePath

		// compress to medium size
		mediumPath, err := resizeImage(path, 800, 0, fileName+"-medium"+extension)
		if err != nil {
			return nil, err
		}
		p.Medium = mediumPath

		// compress to small size
		smallPath, err := resizeImage(path, 320, 0, fileName+"-small"+extension)
		if err != nil {
			return nil, err
		}
		p.Small = smallPath
	}

	return p, nil
}

func resizeImage(path string, width, height int, fileName string) (string, error) {
	dir := filepath.Dir(path) + "/"

	if fileName == "" {
		return "", errors.New("fileName cannot be empty")
	}

	img, err := imaging.Open(path)
	if err != nil {
		return "", err
	}

	// resize image file
	img = imaging.Resize(img, width, height, imaging.Lanczos)

	// create file
	dst := imaging.New(img.Bounds().Size().X, img.Bounds().Size().Y, color.NRGBA{0, 0, 0, 0})

	// copy image to dst file
	dst = imaging.Paste(dst, img, image.Pt(0, 0))

	newPath := dir + fileName
	// save image file
	if err := imaging.Save(dst, newPath); err != nil {
		return "", err
	}

	return newPath, nil
}

// DeleteFile remove file
func (p *Path) DeleteFile() error {
	if p.Original != "" {
		if err := os.Remove(p.Original); err != nil {
			return err
		}
	}

	if p.Large != "" {
		if err := os.Remove(p.Large); err != nil {
			return err
		}
	}

	if p.Medium != "" {
		if err := os.Remove(p.Medium); err != nil {
			return err
		}
	}

	if p.Small != "" {
		if err := os.Remove(p.Small); err != nil {
			return err
		}
	}

	return nil
}
