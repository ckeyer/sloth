package docker

import (
	"fmt"

	libdocker "github.com/fsouza/go-dockerclient"
)

func GetDockerClient() {

}

func atry() {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := libdocker.NewClient(endpoint)
	// imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	// for _, img := range imgs {
	// 	fmt.Println("ID: ", img.ID)
	// 	fmt.Println("RepoTags: ", img.RepoTags)
	// 	fmt.Println("Created: ", img.Created)
	// 	fmt.Println("Size: ", img.Size)
	// 	fmt.Println("VirtualSize: ", img.VirtualSize)
	// 	fmt.Println("ParentId: ", img.ParentID)
	// }
	lisr := make(chan *libdocker.APIEvents)
	client.AddEventListener(lisr)

	for {
		evt := <-lisr
		fmt.Printf("%#v\n", evt)
	}
	// client.BuildImage(opts)
}
