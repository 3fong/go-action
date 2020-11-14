package main

import (
	"fmt"
	"log"

	"gopl.io/ch8/thumbnail"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func main(){
	files := []string{"test/routine/file/3.png","test/routine/file/4.png"}
	fmt.Println(makeThumbnails4(files)) 
}

func makeThumbnails3(filenames []string) {
    ch := make(chan struct{})
    for _, f := range filenames {
        go func(f string) {
            thumbnail.ImageFile(f) // NOTE: ignoring errors
            ch <- struct{}{}
        }(f)
    }
    // Wait for goroutines to complete.
    for range filenames {
        <-ch
    }
}


func makeThumbnails4(filenames []string) error {
    errors := make(chan error)

    for _, f := range filenames {
        go func(f string) {
            _, err := thumbnail.ImageFile(f)
			errors <- err
			fmt.Println(err)
			fmt.Println(len(errors))
        }(f)
    }

    for range filenames {
        if err := <-errors; err != nil {
			fmt.Println(len(errors))
            return err // NOTE: incorrect: goroutine leak!
        }
    }

    return nil
}
