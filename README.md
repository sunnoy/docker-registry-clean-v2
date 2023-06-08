# docker-registry-clean-v2

# use

```GO
	if *url == "" || *username == "" || *password == "" || *num == 0 {
		fmt.Println("Usage: registry -url=<registry_url> -username=<username> -password=<password> -num=<num> -dryrun; num is Keep the latest tag count")
		fmt.Println("del need run gc cmd: registry garbage-collect [--dry-run] [--delete-untagged] /path/to/config.yml")
		return
	}
```
