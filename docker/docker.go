package docker

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	dockerpkg "github.com/fsouza/go-dockerclient"
)

var (
	client *dockerpkg.Client
)

func GetClient() *dockerpkg.Client {
	if client != nil {
		return client
	}

	log.Panic("docker client is nil.")
	return nil
}

func Connect(endpoint string) (*dockerpkg.Client, error) {
	cli, err := dockerpkg.NewClient(endpoint)
	if err != nil {
		return nil, err
	}

	err = cli.Ping()
	if err != nil {
		return nil, fmt.Errorf("can not ping docker client<%s>", endpoint)
	}

	log.Infof("connected docker client: %s", endpoint)
	client = cli
	return cli, nil
}

// func atry() {
// 	endpoint := "unix:///var/run/docker.sock"

// 	client, _ := NewClient(endpoint)
// 	// imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
// 	// for _, img := range imgs {
// 	// 	fmt.Println("ID: ", img.ID)
// 	// 	fmt.Println("RepoTags: ", img.RepoTags)
// 	// 	fmt.Println("Created: ", img.Created)
// 	// 	fmt.Println("Size: ", img.Size)
// 	// 	fmt.Println("VirtualSize: ", img.VirtualSize)
// 	// 	fmt.Println("ParentId: ", img.ParentID)
// 	// }
// 	lisr := make(chan *APIEvents)
// 	client.AddEventListener(lisr)

// 	for {
// 		evt := <-lisr
// 		fmt.Printf("%#v\n", evt)
// 	}
// 	// client.BuildImage(opts)
// }
