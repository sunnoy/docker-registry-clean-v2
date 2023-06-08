package main

import (
	registry "docker-registry-clean-v2/pkg"
	"flag"
	"fmt"
	"log"
)

func main() {
	// 输入参数解析
	url := flag.String("url", "", "Registry URL")
	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "Password")
	dryrun := flag.Bool("dryrun", false, "dryrun")
	num := flag.Int("num", 0, "Number of latest tags to keep")

	flag.Parse()

	if *url == "" || *username == "" || *password == "" || *num == 0 {
		fmt.Println("Usage: registry -url=<registry_url> -username=<username> -password=<password> -num=<num> -dryrun; num is Keep the latest tag count")
		fmt.Println("del need run gc cmd: registry garbage-collect [--dry-run] [--delete-untagged] /path/to/config.yml")
		return
	}

	hub, err := registry.New(*url, *username, *password)
	if err != nil {
		log.Println(err)
		return
	}

	repositories, err := hub.Repositories()
	if err != nil {
		log.Println(err, "repo error")
		return
	}

	for _, re := range repositories {
		tags, err := hub.Tags(re)
		if err != nil {
			log.Println(err)
			continue
		}

		// 判断镜像标签数量是否大于 num
		if len(tags) > *num {
			// 对标签进行排序，保留最新的 num 个
			//sortedTags := sortTags(tags)
			//截取从0到 总数-保留的tag
			tagsToDelete := tags[:len(tags)-*num]
			tagsToR := tags[len(tags)-*num:]
			fmt.Println("【", re, "】", "------reserve tags-------:")
			print(tagsToR)

			fmt.Println("【", re, "】", "======【del tags】======:")
			delprint(tagsToDelete)

			for _, tag := range tagsToDelete {

				if *dryrun == true {
					log.Printf("Deleting  --[[dry run]]-- tag %s in repository %s\n", tag, re)
				} else {
					log.Printf("Deleting  tag %s in repository %s\n", tag, re)
					digest, err := hub.ManifestDigest(re, tag)
					if err != nil {
						log.Println(err)

					}

					err = hub.DeleteManifest(re, digest)
					if err != nil {
						log.Println(err)
					}
				}

			}
		}
	}
}

func print(s []string) {
	green := "\033[32m"
	reset := "\033[0m"
	for _, str := range s {
		fmt.Println(fmt.Sprintf("%s%s%s", green, str, reset))
	}
}

func delprint(s []string) {
	red := "\033[31m"
	reset := "\033[0m"
	for _, str := range s {

		fmt.Println(fmt.Sprintf("%s%s%s", red, str, reset))
	}
}
