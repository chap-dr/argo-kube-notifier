package controller

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/util/workqueue"
	config2 "sigs.k8s.io/controller-runtime/pkg/client/config"
)

type Watcher struct {
	resource        *schema.GroupVersionResource
	objectQueue     workqueue.RateLimitingInterface
	apiWatcher      watch.Interface
	lastSyncVersion string
	resourceVerMap  map[string]string
}

func NewWatcher(resource *schema.GroupVersionResource, objectQueue workqueue.RateLimitingInterface, lastSyncVersion string, resourceVerMap map[string]string) Watcher {
	nw := Watcher{
		resource:        resource,
		objectQueue:     objectQueue,
		resourceVerMap:  resourceVerMap,
		lastSyncVersion: lastSyncVersion,
	}
	return nw
}

func (w *Watcher) watch() {
	fmt.Printf("*DEBUG* 3\n")
	if w.resource.Resource == "" {
		fmt.Printf("*DEBUG* 3.1\n")
		return
	}
	for {
		fmt.Printf("*DEBUG* 3.2\n")
		err := w.createWatcher()
		if err != nil {
			fmt.Printf("*DEBUG* 3.3\n")
			log.Error(err)
			break

		}
		fmt.Printf("*DEBUG* 3.4\n")
		w.runWatch()
	}
}

func (w *Watcher) createWatcher() error {
	fmt.Printf("*DEBUG* 4\n")
	config, err := config2.GetConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Printf("*DEBUG* 4.1\n")
		return err
	}
	api := clientset.Resource(*w.resource)

	listStruct, err := api.List(v1.ListOptions{})
	if err != nil || listStruct == nil {
		fmt.Printf("*DEBUG* 4.2\n")
		return err
	}
	w.lastSyncVersion = listStruct.GetResourceVersion()
	fmt.Println(w.lastSyncVersion)
	fmt.Printf("*DEBUG* 4.3\n")
	w.apiWatcher, err = api.
		Watch(v1.ListOptions{ResourceVersion: w.lastSyncVersion})
	if err != nil {
		fmt.Printf("*DEBUG* 4.4\n")
		log.Fatal(err)
		return err
	}

	return nil
}

func (w *Watcher) runWatch() {
	ch := w.apiWatcher.ResultChan()
	for event := range ch {
		w.objectQueue.Add(event)

	}
}
