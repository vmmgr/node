# node

node application

### Install

```
sudo apt install cloud-image-utils
```

**ZFSを使っている場合**

```
zfs set primarycache=metadata tank
zfs set secondarycache=metadata tank
```

### Build
```
GOOS=linux GOARCH=amd64 go build .
```