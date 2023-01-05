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