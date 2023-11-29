package main

import (
	// "esefexbot/filedb"
	// "crypto/md5"
	"io"
	// "fmt"
	"log"
	"net/http"

	// "strconv"
	// "time"
	// "unicode/utf8"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	url := "https://cdn.discordapp.com/attachments/777344211828604950/1178474445673869404/y2mate.com_-_heheheha_Bass_boosted.mp3"
	// url := "https://cdn.discordapp.com/attachments/777344211828604950/1178752698552700978/test.txt"

	out, err := os.Create("direct.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)

	// req, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(req.Header)

	// len := req.ContentLength
	// log.Println(len)

	// buf := make([]byte, len)
	// log.Println(req.Status)
	// req.Body.Read(buf)

	// checksum := md5.Sum(buf)
	// log.Printf("%x", checksum)

	// // check how many bytes are null

	// nullBytes := 0
	// for i := int64(0); i < len; i++ {
	// 	if buf[i] == 0 {
	// 		nullBytes++
	// 	}
	// }

	// log.Println(nullBytes)

	// // for i := 0; i < (int(len) / 16); i++ {
	// // 	time.Sleep(1 * time.Millisecond)

	// // 	adress := strconv.FormatInt(int64(i*16), 16)
	// // 	content := string(buf[i*16 : i*16+16])

	// // 	println(fmt.Sprintf("%s: %x", adress, content))
	// // }

	// os.WriteFile("test.mp3", buf, os.ModePerm)

	// filedb.AddSound("testserver", "test", ":sus:", buf)
}
