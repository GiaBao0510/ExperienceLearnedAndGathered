# Docker - Các lệnh cơ bản với Ví dụ

## 📋 Mục lục

1. [Kiểm tra và cài đặt](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-ki%E1%BB%83m-tra-v%C3%A0-c%C3%A0i-%C4%91%E1%BA%B7t)
2. [Quản lý Images](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-qu%E1%BA%A3n-l%C3%BD-images)
3. [Quản lý Containers](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-qu%E1%BA%A3n-l%C3%BD-containers)
4. [Tương tác với Container](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-t%C6%B0%C6%A1ng-t%C3%A1c-v%E1%BB%9Bi-container)
5. [Quản lý Networks](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-qu%E1%BA%A3n-l%C3%BD-networks)
6. [Quản lý Volumes](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-qu%E1%BA%A3n-l%C3%BD-volumes)
7. [Backup và Restore](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#7-backup-v%C3%A0-restore)
8. [Docker Compose cơ bản](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#8-docker-compose-c%C6%A1-b%E1%BA%A3n)
9. [Tips và Tricks](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#9-tips-v%C3%A0-tricks)

---

## 1. Kiểm tra và cài đặt

### 1.1. Kiểm tra version Docker

```bash
docker --version
```

**Output:**

```
Docker version 24.0.7, build afdd53b
```

**Kiểm tra thông tin chi tiết:**

```bash
docker version
```

**Output:**

```
Client:
 Version:           24.0.7
 API version:       1.43
 Go version:        go1.20.10
 Git commit:        afdd53b
 Built:             Thu Oct 26 09:08:44 2023
 OS/Arch:           linux/amd64

Server:
 Engine:
  Version:          24.0.7
  API version:      1.43 (minimum version 1.12)
  Go version:       go1.20.10
```

**Kiểm tra Docker đang chạy:**

```bash
docker info
```

---

## 2. Quản lý Images

### 2.1. Xem danh sách Images

```bash
docker images
```

**Hoặc:**

```bash
docker images -a
```

**Output:**

```
REPOSITORY    TAG       IMAGE ID       CREATED        SIZE
nginx         latest    605c77e624dd   2 weeks ago    141MB
ubuntu        20.04     ba6acccedd29   3 weeks ago    72.8MB
mysql         8.0       3218b38490ce   4 weeks ago    516MB
node          18        f77a1a0a9c44   5 weeks ago    999MB
```

**Giải thích:**

- `REPOSITORY`: Tên image
- `TAG`: Phiên bản (latest = mới nhất)
- `IMAGE ID`: ID duy nhất của image
- `CREATED`: Thời gian tạo
- `SIZE`: Kích thước image

### 2.2. Tìm kiếm Image trên Docker Hub

```bash
docker search nginx
```

**Output:**

```
NAME                              DESCRIPTION                                     STARS     OFFICIAL
nginx                             Official build of Nginx.                        19000     [OK]
jwilder/nginx-proxy               Automated Nginx reverse proxy for docker        2200
```

### 2.3. Pull (Tải) Image

**Pull phiên bản latest:**

```bash
docker pull nginx
```

**Output:**

```
Using default tag: latest
latest: Pulling from library/nginx
a2abf6c4d29d: Pull complete
a9edb18cadd1: Pull complete
589b7251471a: Pull complete
Digest: sha256:0d17b565c37bcbd895e9d92315a05c1c3c9a29f762b011a10c54a66cd53c9b31
Status: Downloaded newer image for nginx:latest
docker.io/library/nginx:latest
```

**Pull phiên bản cụ thể:**

```bash
docker pull nginx:1.25
```

**Pull image từ registry khác:**

```bash
docker pull mysql:8.0
```

### 2.4. Xóa Image

**Xóa theo tên:**

```bash
docker rmi nginx
```

**Xóa theo ID:**

```bash
docker rmi 605c77e624dd
```

**Output:**

```
Untagged: nginx:latest
Untagged: nginx@sha256:0d17b565c37bcbd895e9d92315a05c1c3c9a29f762b011a10c54a66cd53c9b31
Deleted: sha256:605c77e624ddb75e6110f997c58876baa13f8754486b461117934b24a9dc3a85
```

**Xóa image đang được sử dụng (force):**

```bash
docker rmi -f nginx
```

**Xóa tất cả images:**

```bash
docker rmi $(docker images -q)
```

**Hoặc:**

```bash
docker rmi -f $(docker images -aq)
```

### 2.5. Xóa images không được sử dụng

```bash
docker image prune
```

**Output:**

```
WARNING! This will remove all dangling images.
Are you sure you want to continue? [y/N] y
Deleted Images:
untagged: nginx@sha256:abc123...
Total reclaimed space: 127.5MB
```

**Xóa tất cả images không được sử dụng (kể cả có tag):**

```bash
docker image prune -a
```

---

## 3. Quản lý Containers

### 3.1. Xem danh sách Containers

**Containers đang chạy:**

```bash
docker ps
```

**Output:**

```
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS                NAMES
a1b2c3d4e5f6   nginx     "/docker-entrypoint.…"   10 minutes ago   Up 10 minutes   0.0.0.0:80->80/tcp   my-nginx
b2c3d4e5f6g7   mysql     "docker-entrypoint.s…"   1 hour ago       Up 1 hour       3306/tcp             my-mysql
```

**Tất cả containers (kể cả đã dừng):**

```bash
docker ps -a
```

**Output:**

```
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS                      PORTS     NAMES
a1b2c3d4e5f6   nginx     "/docker-entrypoint.…"   10 minutes ago   Up 10 minutes               80/tcp    my-nginx
c3d4e5f6g7h8   ubuntu    "/bin/bash"              2 hours ago      Exited (0) 2 hours ago                my-ubuntu
```

**Chỉ hiển thị Container IDs:**

```bash
docker ps -q
```

**Output:**

```
a1b2c3d4e5f6
b2c3d4e5f6g7
```

### 3.2. Chạy Container

**Cú pháp cơ bản:**

```bash
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
```

**Ví dụ 1: Chạy nginx cơ bản**

```bash
docker run nginx
```

**Ví dụ 2: Chạy detached mode (background)**

```bash
docker run -d nginx
```

**Output:**

```
a1b2c3d4e5f6a7b8c9d0e1f2g3h4i5j6k7l8m9n0o1p2q3r4s5t6
```

**Ví dụ 3: Đặt tên container**

```bash
docker run -d --name my-nginx nginx
```

**Ví dụ 4: Map port (host:container)**

```bash
docker run -d -p 8080:80 --name web-server nginx
```

**Giải thích:**

- `-d`: Detached mode (chạy nền)
- `-p 8080:80`: Map port 8080 của host → port 80 của container
- `--name web-server`: Đặt tên container

**Kiểm tra:**

```bash
curl http://localhost:8080
```

**Ví dụ 5: Mount volume**

```bash
docker run -d -p 8080:80 -v /my/html:/usr/share/nginx/html --name web nginx
```

**Ví dụ 6: Set biến môi trường**

```bash
docker run -d -e MYSQL_ROOT_PASSWORD=mypassword --name my-db mysql:8.0
```

**Ví dụ 7: Interactive mode với Ubuntu**

```bash
docker run -it ubuntu bash
```

**Output:**

```
root@a1b2c3d4e5f6:/# 
```

**Giải thích:**

- `-i`: Interactive (giữ STDIN mở)
- `-t`: TTY (terminal)
- `bash`: Command chạy trong container

**Ví dụ 8: Chạy rồi tự xóa khi dừng**

```bash
docker run --rm -it ubuntu bash
```

**Ví dụ 9: Giới hạn tài nguyên**

```bash
docker run -d --memory="512m" --cpus="1.0" nginx
```

### 3.3. Dừng Container

**Dừng theo tên:**

```bash
docker stop my-nginx
```

**Dừng theo ID:**

```bash
docker stop a1b2c3d4e5f6
```

**Output:**

```
a1b2c3d4e5f6
```

**Dừng nhiều containers:**

```bash
docker stop container1 container2 container3
```

**Dừng tất cả containers đang chạy:**

```bash
docker stop $(docker ps -q)
```

**Dừng container ngay lập tức (không chờ graceful shutdown):**

```bash
docker kill my-nginx
```

### 3.4. Khởi động lại Container

**Start container đã dừng:**

```bash
docker start my-nginx
```

**Restart container đang chạy:**

```bash
docker restart my-nginx
```

**Start và attach:**

```bash
docker start -i my-ubuntu
```

### 3.5. Xóa Container

**Xóa container đã dừng:**

```bash
docker rm my-nginx
```

**Output:**

```
my-nginx
```

**Xóa container đang chạy (force):**

```bash
docker rm -f my-nginx
```

**Xóa nhiều containers:**

```bash
docker rm container1 container2 container3
```

**Xóa tất cả containers đã dừng:**

```bash
docker container prune
```

**Output:**

```
WARNING! This will remove all stopped containers.
Are you sure you want to continue? [y/N] y
Deleted Containers:
a1b2c3d4e5f6a7b8c9d0e1f2g3h4i5j6k7l8m9n0o1p2q3r4s5t6
c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0u1v2w3x4y5z6

Total reclaimed space: 127.5MB
```

**Xóa tất cả containers (kể cả đang chạy):**

```bash
docker rm -f $(docker ps -aq)
```

---

## 4. Tương tác với Container

### 4.1. Attach vào Container đang chạy

```bash
docker attach my-nginx
```

**Thoát khỏi attach:**

- `Ctrl + C`: Dừng container
- `Ctrl + P, Ctrl + Q`: Detach không dừng container

### 4.2. Exec - Chạy lệnh trong Container

**Mở bash shell:**

```bash
docker exec -it my-nginx bash
```

**Output:**

```
root@a1b2c3d4e5f6:/# 
```

**Chạy một lệnh cụ thể:**

```bash
docker exec my-nginx ls /usr/share/nginx/html
```

**Output:**

```
50x.html
index.html
```

**Xem file:**

```bash
docker exec my-nginx cat /etc/nginx/nginx.conf
```

**Chạy với user khác:**

```bash
docker exec -u root my-nginx whoami
```

### 4.3. Xem Logs

**Xem logs:**

```bash
docker logs my-nginx
```

**Output:**

```
2024/03/03 10:30:45 [notice] 1#1: using the "epoll" event method
2024/03/03 10:30:45 [notice] 1#1: nginx/1.25.3
2024/03/03 10:30:45 [notice] 1#1: start worker processes
```

**Follow logs (real-time):**

```bash
docker logs -f my-nginx
```

**Xem logs với timestamp:**

```bash
docker logs -t my-nginx
```

**Xem 50 dòng logs cuối:**

```bash
docker logs --tail 50 my-nginx
```

**Xem logs từ 10 phút trước:**

```bash
docker logs --since 10m my-nginx
```

### 4.4. Inspect - Xem thông tin chi tiết

```bash
docker inspect my-nginx
```

**Output (rút gọn):**

```json
[
    {
        "Id": "a1b2c3d4e5f6a7b8c9d0e1f2g3h4i5j6k7l8m9n0o1p2q3r4s5t6",
        "Created": "2024-03-03T10:30:45.123456789Z",
        "Path": "/docker-entrypoint.sh",
        "Args": ["nginx", "-g", "daemon off;"],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "StartedAt": "2024-03-03T10:30:45.987654321Z"
        },
        "Image": "sha256:605c77e624dd...",
        "NetworkSettings": {
            "IPAddress": "172.17.0.2",
            "Ports": {
                "80/tcp": [
                    {
                        "HostIp": "0.0.0.0",
                        "HostPort": "8080"
                    }
                ]
            }
        }
    }
]
```

**Lấy thông tin cụ thể:**

```bash
docker inspect -f '{{.NetworkSettings.IPAddress}}' my-nginx
```

**Output:**

```
172.17.0.2
```

**Lấy nhiều thông tin:**

```bash
docker inspect -f '{{.Name}} - {{.State.Status}}' my-nginx
```

**Output:**

```
/my-nginx - running
```

### 4.5. Stats - Xem tài nguyên

**Xem real-time stats:**

```bash
docker stats
```

**Output:**

```
CONTAINER ID   NAME        CPU %     MEM USAGE / LIMIT   MEM %     NET I/O           BLOCK I/O
a1b2c3d4e5f6   my-nginx    0.01%     3.5MiB / 7.77GiB   0.04%     1.2kB / 648B     0B / 0B
b2c3d4e5f6g7   my-mysql    0.50%     201MiB / 7.77GiB   2.53%     5.4kB / 3.2kB    12MB / 8MB
```

**Stats một container cụ thể:**

```bash
docker stats my-nginx
```

**Stats không stream (chỉ hiển thị 1 lần):**

```bash
docker stats --no-stream
```

### 4.6. Top - Xem processes

```bash
docker top my-nginx
```

**Output:**

```
UID       PID       PPID      C     STIME     TTY       TIME          CMD
root      12345     12320     0     10:30     ?         00:00:00      nginx: master process
nginx     12346     12345     0     10:30     ?         00:00:00      nginx: worker process
```

### 4.7. Copy files giữa host và container

**Copy từ host vào container:**

```bash
docker cp /path/on/host/file.txt my-nginx:/path/in/container/
```

**Ví dụ:**

```bash
docker cp index.html my-nginx:/usr/share/nginx/html/
```

**Copy từ container ra host:**

```bash
docker cp my-nginx:/var/log/nginx/access.log ./access.log
```

**Copy thư mục:**

```bash
docker cp ./my-app my-nginx:/usr/share/nginx/html/
```

---

## 5. Quản lý Networks

### 5.1. Xem danh sách Networks

```bash
docker network ls
```

**Output:**

```
NETWORK ID     NAME      DRIVER    SCOPE
a1b2c3d4e5f6   bridge    bridge    local
b2c3d4e5f6g7   host      host      local
c3d4e5f6g7h8   none      null      local
```

**Giải thích:**

- `bridge`: Network mặc định, containers có thể giao tiếp với nhau
- `host`: Container dùng network của host
- `none`: Container không có network

### 5.2. Tạo Network mới

```bash
docker network create my-network
```

**Output:**

```
d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0u1v2w3x4y5z6a7b8c9d0
```

**Tạo network với subnet cụ thể:**

```bash
docker network create --subnet=172.18.0.0/16 my-custom-network
```

**Tạo network với driver khác:**

```bash
docker network create --driver bridge my-bridge-network
```

### 5.3. Kết nối Container vào Network

**Kết nối container đang chạy:**

```bash
docker network connect my-network my-nginx
```

**Chạy container với network cụ thể:**

```bash
docker run -d --name web --network my-network nginx
```

**Chạy nhiều containers trong cùng network:**

```bash
docker run -d --name web --network my-network nginx
docker run -d --name db --network my-network mysql:8.0
```

**Test kết nối:**

```bash
docker exec web ping db
```

**Output:**

```
PING db (172.18.0.3) 56(84) bytes of data.
64 bytes from db.my-network (172.18.0.3): icmp_seq=1 ttl=64 time=0.123 ms
```

### 5.4. Ngắt kết nối Network

```bash
docker network disconnect my-network my-nginx
```

### 5.5. Inspect Network

```bash
docker network inspect my-network
```

**Output (rút gọn):**

```json
[
    {
        "Name": "my-network",
        "Id": "d4e5f6g7h8i9...",
        "Driver": "bridge",
        "Scope": "local",
        "Subnet": "172.18.0.0/16",
        "Containers": {
            "a1b2c3d4e5f6": {
                "Name": "web",
                "IPv4Address": "172.18.0.2/16"
            },
            "b2c3d4e5f6g7": {
                "Name": "db",
                "IPv4Address": "172.18.0.3/16"
            }
        }
    }
]
```

### 5.6. Xóa Network

```bash
docker network rm my-network
```

**Xóa tất cả networks không được sử dụng:**

```bash
docker network prune
```

---

## 6. Quản lý Volumes

### 6.1. Xem danh sách Volumes

```bash
docker volume ls
```

**Output:**

```
DRIVER    VOLUME NAME
local     my-vol
local     db-data
local     app-config
```

### 6.2. Tạo Volume

```bash
docker volume create my-vol
```

**Output:**

```
my-vol
```

### 6.3. Sử dụng Volume

**Mount volume khi chạy container:**

```bash
docker run -d -v my-vol:/data --name my-container ubuntu
```

**Bind mount (thư mục host):**

```bash
docker run -d -v /home/user/data:/data --name my-container ubuntu
```

**Ví dụ với MySQL:**

```bash
docker run -d \
  --name my-mysql \
  -e MYSQL_ROOT_PASSWORD=mypassword \
  -v mysql-data:/var/lib/mysql \
  mysql:8.0
```

### 6.4. Inspect Volume

```bash
docker volume inspect my-vol
```

**Output:**

```json
[
    {
        "CreatedAt": "2024-03-03T10:30:45Z",
        "Driver": "local",
        "Mountpoint": "/var/lib/docker/volumes/my-vol/_data",
        "Name": "my-vol",
        "Scope": "local"
    }
]
```

### 6.5. Xóa Volume

```bash
docker volume rm my-vol
```

**Xóa tất cả volumes không được sử dụng:**

```bash
docker volume prune
```

**Output:**

```
WARNING! This will remove all local volumes not used by at least one container.
Are you sure you want to continue? [y/N] y
Deleted Volumes:
my-vol
db-data

Total reclaimed space: 1.2GB
```

---

## 7. Backup và Restore

### 7.1. Lưu Container thành Image

```bash
docker commit my-nginx my-custom-nginx:v1
```

**Output:**

```
sha256:e1f2g3h4i5j6k7l8m9n0o1p2q3r4s5t6u7v8w9x0y1z2a3b4c5d6e7f8
```

**Với message:**

```bash
docker commit -m "Added custom config" my-nginx my-custom-nginx:v1
```

**Với author:**

```bash
docker commit -a "John Doe" my-nginx my-custom-nginx:v1
```

### 7.2. Lưu Image ra file

**Lưu 1 image:**

```bash
docker save -o my-nginx.tar nginx
```

**Hoặc:**

```bash
docker save --output my-nginx.tar nginx
```

**Lưu nhiều images:**

```bash
docker save -o images.tar nginx mysql ubuntu
```

**Kiểm tra file:**

```bash
ls -lh my-nginx.tar
```

**Output:**

```
-rw------- 1 user user 142M Mar  3 10:30 my-nginx.tar
```

### 7.3. Load Image từ file

```bash
docker load -i my-nginx.tar
```

**Output:**

```
Loaded image: nginx:latest
```

**Hoặc:**

```bash
docker load < my-nginx.tar
```

### 7.4. Export/Import Container

**Export container sang tar:**

```bash
docker export my-nginx > my-nginx-container.tar
```

**Hoặc:**

```bash
docker export -o my-nginx-container.tar my-nginx
```

**Import tar thành image:**

```bash
docker import my-nginx-container.tar my-nginx:imported
```

**Output:**

```
sha256:a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0u1v2w3x4y5z6
```

### 7.5. Đổi tên Image (tag)

```bash
docker tag nginx:latest my-nginx:v1.0
```

**Verify:**

```bash
docker images | grep my-nginx
```

**Output:**

```
my-nginx     v1.0      605c77e624dd   2 weeks ago   141MB
```

**Tag cho registry:**

```bash
docker tag my-nginx:v1.0 myregistry.com/my-nginx:v1.0
```

---

## 8. Docker Compose cơ bản

### 8.1. File docker-compose.yml

**Ví dụ: Web + Database**

```yaml
version: '3.8'

services:
  web:
    image: nginx:latest
    ports:
      - "8080:80"
    volumes:
      - ./html:/usr/share/nginx/html
    networks:
      - app-network

  db:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: mypassword
      MYSQL_DATABASE: mydb
    volumes:
      - db-data:/var/lib/mysql
    networks:
      - app-network

networks:
  app-network:
    driver: bridge

volumes:
  db-data:
```

### 8.2. Các lệnh Docker Compose

**Khởi động services:**

```bash
docker-compose up
```

**Khởi động background:**

```bash
docker-compose up -d
```

**Output:**

```
Creating network "myapp_app-network" with driver "bridge"
Creating volume "myapp_db-data" with default driver
Creating myapp_db_1  ... done
Creating myapp_web_1 ... done
```

**Xem logs:**

```bash
docker-compose logs
```

**Follow logs:**

```bash
docker-compose logs -f
```

**Logs của service cụ thể:**

```bash
docker-compose logs web
```

**Dừng services:**

```bash
docker-compose stop
```

**Dừng và xóa:**

```bash
docker-compose down
```

**Output:**

```
Stopping myapp_web_1 ... done
Stopping myapp_db_1  ... done
Removing myapp_web_1 ... done
Removing myapp_db_1  ... done
Removing network myapp_app-network
```

**Dừng và xóa kể cả volumes:**

```bash
docker-compose down -v
```

**Xem services đang chạy:**

```bash
docker-compose ps
```

**Output:**

```
      Name                    Command              State           Ports
--------------------------------------------------------------------------------
myapp_db_1       docker-entrypoint.sh mysqld   Up      3306/tcp, 33060/tcp
myapp_web_1      /docker-entrypoint.sh ngin... Up      0.0.0.0:8080->80/tcp
```

**Build và start:**

```bash
docker-compose up --build
```

**Restart services:**

```bash
docker-compose restart
```

---

## 9. Tips và Tricks

### 9.1. Lệnh hữu ích

**Xem tất cả thông tin hệ thống:**

```bash
docker system df
```

**Output:**

```
TYPE            TOTAL     ACTIVE    SIZE      RECLAIMABLE
Images          10        5         2.5GB     1.2GB (48%)
Containers      15        3         156MB     120MB (77%)
Local Volumes   5         2         1.2GB     800MB (67%)
Build Cache     0         0         0B        0B
```

**Dọn dẹp toàn bộ:**

```bash
docker system prune -a
```

**Output:**

```
WARNING! This will remove:
  - all stopped containers
  - all networks not used by at least one container
  - all images without at least one container associated to them
  - all build cache

Are you sure you want to continue? [y/N] y

Deleted Containers:
...

Deleted Images:
...

Total reclaimed space: 3.5GB
```

**Dọn dẹp kể cả volumes:**

```bash
docker system prune -a --volumes
```

### 9.2. Lấy IP của Container

**Cách 1:**

```bash
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' my-nginx
```

**Output:**

```
172.17.0.2
```

**Cách 2:**

```bash
docker inspect my-nginx | grep IPAddress
```

**Output:**

```
            "IPAddress": "172.17.0.2",
                    "IPAddress": "172.17.0.2",
```

**Cách 3: Từ trong container**

```bash
docker exec my-nginx hostname -i
```

**Output:**

```
172.17.0.2
```

### 9.3. Xem port mapping

```bash
docker port my-nginx
```

**Output:**

```
80/tcp -> 0.0.0.0:8080
```

### 9.4. Rename container

```bash
docker rename old-name new-name
```

**Ví dụ:**

```bash
docker rename my-nginx nginx-server
```

### 9.5. Pause/Unpause container

**Pause (đóng băng):**

```bash
docker pause my-nginx
```

**Unpause:**

```bash
docker unpause my-nginx
```

### 9.6. Update container resources

```bash
docker update --memory="1g" --cpus="2" my-nginx
```

**Output:**

```
my-nginx
```

### 9.7. Events - Xem sự kiện real-time

```bash
docker events
```

**Output:**

```
2024-03-03T10:30:45.123456789+07:00 container start a1b2c3d4e5f6 (image=nginx, name=my-nginx)
2024-03-03T10:31:23.987654321+07:00 container stop a1b2c3d4e5f6 (image=nginx, name=my-nginx)
```

**Filter events:**

```bash
docker events --filter 'type=container' --filter 'event=start'
```

### 9.8. Diff - Xem thay đổi trong container

```bash
docker diff my-nginx
```

**Output:**

```
C /var
C /var/cache
C /var/cache/nginx
A /var/cache/nginx/client_temp
A /var/cache/nginx/fastcgi_temp
```

**Giải thích:**

- `A`: Added (thêm mới)
- `C`: Changed (thay đổi)
- `D`: Deleted (xóa)

### 9.9. Wait - Chờ container dừng

```bash
docker wait my-nginx
```

**Output (exit code):**

```
0
```

### 9.10. Health check

**Thêm health check khi run:**

```bash
docker run -d \
  --name web \
  --health-cmd="curl -f http://localhost/ || exit 1" \
  --health-interval=30s \
  --health-timeout=10s \
  --health-retries=3 \
  nginx
```

**Xem health status:**

```bash
docker inspect --format='{{.State.Health.Status}}' web
```

**Output:**

```
healthy
```

---

## 📚 Tổng hợp lệnh thường dùng

### Images

```bash
docker images                      # Xem images
docker pull nginx                  # Tải image
docker rmi nginx                   # Xóa image
docker image prune                 # Xóa images unused
```

### Containers

```bash
docker ps                          # Containers đang chạy
docker ps -a                       # Tất cả containers
docker run -d nginx                # Chạy container
docker stop my-nginx               # Dừng container
docker start my-nginx              # Khởi động lại
docker rm my-nginx                 # Xóa container
docker container prune             # Xóa containers stopped
```

### Tương tác

```bash
docker exec -it my-nginx bash      # Vào shell
docker logs -f my-nginx            # Xem logs
docker inspect my-nginx            # Xem thông tin
docker stats                       # Xem tài nguyên
docker cp file.txt my-nginx:/path  # Copy file
```

### Networks

```bash
docker network ls                  # Xem networks
docker network create my-net       # Tạo network
docker network connect my-net web  # Kết nối container
```

### Volumes

```bash
docker volume ls                   # Xem volumes
docker volume create my-vol        # Tạo volume
docker volume rm my-vol            # Xóa volume
```

### Backup

```bash
docker commit my-nginx my-img:v1   # Lưu container → image
docker save -o file.tar nginx      # Lưu image → tar
docker load -i file.tar            # Load tar → image
docker export my-nginx > file.tar  # Export container
docker import file.tar img:v1      # Import tar → image
```

### Dọn dẹp

```bash
docker system prune -a             # Dọn dẹp toàn bộ
docker system df                   # Xem disk usage
```