package triggers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/go-dev-env/standards"
)

func getDirElements(path string) []string {
	var returnedElements []string
	elements, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, element := range elements {
		elementPath := fmt.Sprintf("%v%v", path, element.Name())

		if element.IsDir() {
			log.Printf("Adding all of the dir elements %v", elementPath)
			returnedElements = append(returnedElements, getDirElements(fmt.Sprintf("%v/", elementPath))...)
		} else {
			log.Printf("Adding the element %v", elementPath)
			returnedElements = append(returnedElements, elementPath)
		}
	}

	return returnedElements
}

func walkdir(path string, info os.FileInfo, err error, watcher *fsnotify.Watcher) (string, error) {
	if err != nil {
		log.Printf("Cannot access path %v, err %v", path, err)
		return "", err
	}
	log.Printf("Visited file or dir: %q\n", path)

	if info.IsDir() {
		return path, nil
	}
	return "", nil
}

func initWatchers(path string) (*fsnotify.Watcher, error) {
	watchersChan, _ := fsnotify.NewWatcher()
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		dirPath, walkErr := walkdir(path, info, err, watchersChan)
		if dirPath != "" {
			log.Printf("Adding watcher on dir: %q\n", path)
			watchersChan.Add(path)
		}
		return walkErr
	})

	if err != nil {
		log.Printf("Error on walk : %v", err)
		return nil, err
	}

	return watchersChan, nil
}

func directoryCreated(op fsnotify.Op, path string) bool {
	if op != fsnotify.Create {
		return false
	}

	fileInfo, err := os.Stat(path)

	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

// FileWatcherTrigger Triggered by a change in any file in the directory
type FileWatcherTrigger struct {
	FileChangedChan           chan bool
	debouncedNotifyFileChange func()
	path                      string
}

// Init Initialize the watcher
func (trigger *FileWatcherTrigger) Init() chan bool {
	log.Println("Initializing a file watcher trigger")
	trigger.FileChangedChan = make(chan bool)
	trigger.debouncedNotifyFileChange = standards.Debounce(func() {
		trigger.FileChangedChan <- true
	}, 500)

	go trigger.startWatchers()

	return trigger.FileChangedChan
}

func (trigger FileWatcherTrigger) startWatchers() {
	watchersChan, err := initWatchers(trigger.path)
	defer watchersChan.Close()

	if err != nil {
		log.Printf("Error on watchers init : %v", err)
	}

	log.Printf("Watchers init successfull, waiting for a change...")

	for {
		select {
		// watch for events
		case event := <-watchersChan.Events:
			log.Printf("New event of type %v : %v\n", event.Op.String(), event.Name)
			if directoryCreated(event.Op, event.Name) {
				log.Printf("Adding a watcher in the dir : %v\n", event.Name)
				watchersChan.Add(event.Name)
			}

			if event.Op == fsnotify.Create || event.Op == fsnotify.Write || event.Op == fsnotify.Rename {

				trigger.debouncedNotifyFileChange()
			}
		// watch for errors
		case err := <-watchersChan.Errors:
			log.Println("ERROR", err)
		}
	}
}

// NewFileWatcherTrigger Create a new file watcher trigger
func NewFileWatcherTrigger(path string) *FileWatcherTrigger {
	log.Println("Creating a new file watcher trigger")

	return &FileWatcherTrigger{
		path: path,
	}
}
