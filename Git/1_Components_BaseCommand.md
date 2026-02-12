# Các Thành Phần và Lệnh Git Cơ Bản

## Mục Lục

1. [Cấu hình Git](https://claude.ai/chat/04b8ea8e-fc4a-4141-b477-765e01d4b987#1-c%E1%BA%A5u-h%C3%ACnh-git)
2. [Các thành phần chính trong Git](https://claude.ai/chat/04b8ea8e-fc4a-4141-b477-765e01d4b987#2-c%C3%A1c-th%C3%A0nh-ph%E1%BA%A7n-ch%C3%ADnh-trong-git)
3. [Các lệnh Git thường dùng](https://claude.ai/chat/04b8ea8e-fc4a-4141-b477-765e01d4b987#3-c%C3%A1c-l%E1%BB%87nh-git-th%C6%B0%E1%BB%9Dng-d%C3%B9ng)
4. [Workflow thực tế](https://claude.ai/chat/04b8ea8e-fc4a-4141-b477-765e01d4b987#4-workflow-th%E1%BB%B1c-t%E1%BA%BF)

---

## 1. Cấu hình Git

### 1.1. Kiểm tra phiên bản Git

```bash
git --version
# Hoặc:
git -v
```

**Kết quả mong đợi:**

```
git version 2.43.0
```

---

### 1.2. Cấu hình thông tin cá nhân

**Tại sao quan trọng?**  
Mỗi commit bạn tạo sẽ được gắn với tên và email này. Điều này giúp team biết ai đã làm gì và khi nào.

```bash
# ⚠️ LƯU Ý: Không có dấu "=" trong lệnh config
# SAI: git config --global user.name = "Tên"
# ĐÚNG:

# Thiết lập tên (sẽ hiển thị trong commit)
git config --global user.name "Pham Gia Bao"

# Thiết lập email
git config --global user.email "pgiabao2002@gmail.com"

# Kiểm tra cấu hình
git config --list
```

**Ví dụ output khi chạy `git config --list`:**

```bash
user.name=Pham Gia Bao
user.email=pgiabao2002@gmail.com
core.autocrlf=true
init.defaultbranch=main
# ... và nhiều cấu hình khác
```

**Kiểm tra cấu hình cụ thể:**

```bash
# Xem tên đã cấu hình
git config user.name

# Xem email đã cấu hình
git config user.email
```

---

### 1.3. Khởi tạo Repository

```bash
# Tạo Git repository trong thư mục hiện tại
git init

# Hoặc tạo thư mục mới và init luôn
git init my-project
cd my-project
```

**Kết quả:**  
Một thư mục ẩn `.git` được tạo ra, chứa toàn bộ lịch sử và cấu hình Git.

---

## 2. Các thành phần chính trong Git

### 2.1. Working Directory (Thư mục làm việc)

**Mô tả:**  
Là thư mục chứa source code mà bạn đang làm việc, nơi bạn tạo/sửa/xóa file hàng ngày.

**Ví dụ:**

```
my-project/
├── index.html
├── style.css
└── script.js
```

---

### 2.2. Staging Area (Khu vực dàn dựng)

**Mô tả:**  
Là khu vực trung gian giữa Working Directory và Local Repository. Nơi bạn chuẩn bị các file trước khi commit.

**Tại sao cần Staging Area?**  
Giúp bạn kiểm soát chính xác những gì sẽ được commit. Bạn có thể chỉ commit một số file thay đổi, không nhất thiết phải commit tất cả.

**Ví dụ thực tế:**

```bash
# Bạn sửa 3 file: index.html, style.css, script.js
# Nhưng chỉ muốn commit 2 file: index.html và style.css

git add index.html style.css   # Chỉ add 2 file vào staging
git commit -m "Cập nhật trang chủ"  # Chỉ commit 2 file đã add
```

**Các lệnh làm việc với Staging Area:**

```bash
# Thêm 1 file cụ thể vào staging
git add index.html

# Thêm nhiều file
git add index.html style.css

# Thêm TẤT CẢ file đã thay đổi
git add .
# Hoặc:
git add --all
git add -A

# Xem trạng thái (file nào đã staged, chưa staged)
git status

# Xem chi tiết thay đổi của file đã staged
git diff --staged
# Hoặc:
git diff --cached

# BỎ file ra khỏi staging (unstage)
git reset HEAD index.html
# Hoặc (Git 2.23+):
git restore --staged index.html
```

---

### 2.3. Local Repository (Kho cục bộ)

**Mô tả:**  
Là nơi lưu trữ toàn bộ lịch sử commit của dự án trên máy tính của bạn. Được chứa trong thư mục `.git`.

**Đặc điểm:**

- Chứa tất cả các commit
- Chứa tất cả các branch
- Chứa toàn bộ lịch sử dự án
- Hoạt động offline (không cần internet)

**Các lệnh làm việc với Local Repository:**

```bash
# Commit các file đã staged vào Local Repository
git commit -m "Thêm trang chủ"

# Commit và add cùng lúc (chỉ với file đã được track)
git commit -am "Sửa lỗi hiển thị"

# Xem lịch sử commit
git log

# Xem lịch sử ngắn gọn (1 dòng/commit)
git log --oneline

# Xem lịch sử với graph
git log --oneline --graph --all

# Xem chi tiết một commit cụ thể
git show <commit_id>

# Tạo branch mới
git branch feature-login

# Xem danh sách branch
git branch

# Chuyển sang branch khác
git checkout feature-login
# Hoặc (Git 2.23+):
git switch feature-login

# Tạo và chuyển sang branch mới cùng lúc
git checkout -b feature-payment
# Hoặc:
git switch -c feature-payment
```

---

### 2.4. Remote Repository (Kho từ xa)

**Mô tả:**  
Là phiên bản dự án được lưu trữ trên server (GitHub, GitLab, Bitbucket...). Cho phép làm việc nhóm và backup code.

**Mục đích:**

- Chia sẻ code với team
- Backup code an toàn
- Làm việc từ nhiều máy khác nhau
- Cộng tác với người khác

**Các lệnh làm việc với Remote Repository:**

```bash
# Liên kết Local Repository với Remote Repository
git remote add origin https://github.com/username/repo-name.git
# Hoặc dùng SSH:
git remote add origin git@github.com:username/repo-name.git

# Xem danh sách remote
git remote -v

# Đổi tên remote
git remote rename origin upstream

# Thay đổi URL của remote
git remote set-url origin https://github.com/username/new-repo.git

# Xóa remote
git remote remove origin

# Push code lên remote (lần đầu)
git push -u origin main
# -u (--set-upstream): Thiết lập tracking, lần sau chỉ cần: git push

# Push code lên remote (lần sau)
git push

# Push branch mới lên remote
git push -u origin feature-login

# Xóa branch trên remote
git push -d origin feature-login
# Hoặc:
git push origin --delete feature-login

# Lấy code từ remote về (fetch + merge)
git pull

# Pull từ branch cụ thể
git pull origin main

# Lấy thông tin từ remote nhưng KHÔNG merge
git fetch origin

# Clone repository từ remote về local
git clone https://github.com/username/repo-name.git

# Clone và đổi tên thư mục
git clone https://github.com/username/repo-name.git my-folder
```

---

## 3. Các lệnh Git thường dùng

### 3.1. Git Help - Xem hướng dẫn

```bash
# Xem tất cả lệnh có sẵn
git help -a
# Hoặc:
git help --all

# Xem hướng dẫn chi tiết về một lệnh
git help config
# Hoặc:
git config --help

# Xem tóm tắt các option của lệnh
git commit -h
```

---

### 3.2. Git Status - Kiểm tra trạng thái

```bash
# Xem trạng thái chi tiết
git status

# Xem trạng thái ngắn gọn
git status -s
# Hoặc:
git status --short
```

**Giải thích ký hiệu trong `git status -s`:**

```
?? → File chưa được track (mới tạo, chưa add)
A  → File đã add vào staging (mới)
M  → File đã sửa (Modified)
D  → File đã xóa (Deleted)
MM → File đã sửa, một phần đã staged, một phần chưa
```

**Ví dụ output:**

```bash
$ git status -s
?? README.md        # File mới, chưa track
A  index.html      # File mới, đã add
M  style.css       # File đã sửa, đã staged
 M script.js       # File đã sửa, chưa staged
D  old-file.txt    # File đã xóa
```

---

### 3.3. Git Add - Thêm file vào staging

```bash
# Add một file cụ thể
git add index.html

# Add nhiều file cùng lúc
git add index.html style.css script.js

# Add tất cả file .html
git add *.html

# Add tất cả file trong thư mục css/
git add css/

# Add TẤT CẢ file thay đổi
git add .
git add --all
git add -A

# Add từng phần của file (interactive)
git add -p index.html
```

---

### 3.4. Git Commit - Lưu thay đổi

```bash
# Commit với message ngắn gọn
git commit -m "Thêm trang chủ"

# Commit với message chi tiết (mở editor)
git commit

# Add và commit cùng lúc (chỉ với file đã track)
git commit -am "Sửa lỗi hiển thị"

# Sửa commit cuối cùng (message hoặc thêm file)
git commit --amend -m "Message mới"

# Thêm file vào commit cuối mà không đổi message
git add forgotten-file.txt
git commit --amend --no-edit
```

**Quy tắc viết commit message tốt:**

```bash
# TỐT:
git commit -m "Thêm tính năng đăng nhập"
git commit -m "Sửa lỗi hiển thị giá sản phẩm"
git commit -m "Cập nhật README với hướng dẫn cài đặt"

# KHÔNG TỐT:
git commit -m "update"        # Quá chung chung
git commit -m "fix"           # Không rõ sửa gì
git commit -m "test"          # Vô nghĩa
git commit -m "asdfgh"        # Random
```

**Format chuyên nghiệp (Conventional Commits):**

```bash
git commit -m "feat: thêm tính năng đăng nhập"
git commit -m "fix: sửa lỗi hiển thị giá"
git commit -m "docs: cập nhật README"
git commit -m "style: format code theo eslint"
git commit -m "refactor: tối ưu hàm tính toán"
```

---

### 3.5. Git Diff - Xem thay đổi

```bash
# Xem thay đổi chưa staged
git diff

# Xem thay đổi của file cụ thể
git diff index.html

# Xem thay đổi đã staged
git diff --staged
# Hoặc:
git diff --cached

# So sánh giữa 2 branch
git diff main feature-login

# So sánh giữa 2 commit
git diff commit1 commit2

# Xem danh sách file thay đổi (không xem nội dung)
git diff --name-only
```

---

### 3.6. Git Log - Xem lịch sử

```bash
# Xem lịch sử đầy đủ
git log

# Xem lịch sử ngắn gọn (1 dòng/commit)
git log --oneline

# Xem 5 commit gần nhất
git log -5

# Xem lịch sử với graph (nhánh)
git log --graph --oneline --all

# Xem lịch sử với thông tin chi tiết
git log --stat

# Xem lịch sử của một file
git log -- index.html

# Tìm commit theo message
git log --grep="login"

# Tìm commit theo tác giả
git log --author="Gia Bao"

# Xem lịch sử theo thời gian
git log --since="2 weeks ago"
git log --after="2024-01-01"
git log --before="2024-12-31"
```

**Alias hữu ích cho git log:**

```bash
# Thêm vào ~/.gitconfig
git config --global alias.lg "log --graph --oneline --all --decorate"

# Sử dụng:
git lg
```

---

### 3.7. Git Branch - Quản lý nhánh

```bash
# Xem danh sách branch local
git branch

# Xem danh sách branch (cả local và remote)
git branch -a

# Xem branch remote
git branch -r

# Tạo branch mới (không chuyển sang branch đó)
git branch feature-login

# Đổi tên branch hiện tại
git branch -m new-name

# Đổi tên branch khác
git branch -m old-name new-name

# Xóa branch đã merge
git branch -d feature-login

# Xóa branch chưa merge (force delete)
git branch -D feature-login

# Xóa branch trên remote
git push -d origin feature-login
```

---

### 3.8. Git Checkout / Switch - Chuyển nhánh

```bash
# === Sử dụng checkout (cách cũ) ===
# Chuyển sang branch khác
git checkout main

# Tạo và chuyển sang branch mới
git checkout -b feature-payment

# Quay lại commit trước đó (detached HEAD)
git checkout <commit_id>

# Khôi phục file về trạng thái của commit cuối
git checkout -- index.html


# === Sử dụng switch và restore (cách mới, Git 2.23+) ===
# Chuyển sang branch khác
git switch main

# Tạo và chuyển sang branch mới
git switch -c feature-payment

# Quay lại branch trước đó
git switch -

# Khôi phục file (thay vì checkout)
git restore index.html

# Bỏ file ra khỏi staging
git restore --staged index.html
```

**Lưu ý:** `switch` và `restore` được tách ra để tránh nhầm lẫn giữa việc chuyển branch và khôi phục file.

---

### 3.9. Git Merge - Hợp nhất nhánh

```bash
# Chuyển về nhánh nhận merge
git checkout main

# Merge branch khác vào nhánh hiện tại
git merge feature-login

# Merge và tạo commit merge (ngay cả khi fast-forward được)
git merge --no-ff feature-login

# Hủy merge khi có conflict
git merge --abort
```

**Ví dụ thực tế:**

```bash
# Bạn làm xong feature-login, giờ muốn merge vào main

# Bước 1: Chuyển về main
git checkout main

# Bước 2: Cập nhật main từ remote (nếu làm việc nhóm)
git pull origin main

# Bước 3: Merge feature-login vào main
git merge feature-login

# Bước 4: Push lên remote
git push origin main

# Bước 5: Xóa branch feature-login (nếu không cần nữa)
git branch -d feature-login
```

**Xử lý conflict khi merge:**

```bash
# Khi merge bị conflict
git merge feature-login
# Auto-merging index.html
# CONFLICT (content): Merge conflict in index.html

# Bước 1: Mở file bị conflict, sửa thủ công
# Tìm các dòng:
# <<<<<<< HEAD
# code của main
# =======
# code của feature-login
# >>>>>>> feature-login

# Bước 2: Chọn code muốn giữ, xóa các dấu <<<<, ====, >>>>

# Bước 3: Add file đã sửa
git add index.html

# Bước 4: Commit
git commit -m "Merge feature-login vào main"
```

---

### 3.10. Git Rebase - Sắp xếp lại commit

```bash
# Chuyển về nhánh cần rebase
git checkout feature-login

# Rebase lên nhánh main
git rebase main

# Rebase interactive (sửa, gộp, xóa commit)
git rebase -i HEAD~3  # 3 commit gần nhất

# Tiếp tục rebase sau khi giải quyết conflict
git rebase --continue

# Hủy rebase
git rebase --abort
```

**Merge vs Rebase:**

|Merge|Rebase|
|---|---|
|Giữ nguyên lịch sử commit|Viết lại lịch sử commit|
|Tạo commit merge mới|Không tạo commit merge|
|Lịch sử phức tạp hơn|Lịch sử tuyến tính, sạch đẹp|
|An toàn cho shared branch|KHÔNG an toàn cho shared branch|
|Dùng khi làm việc nhóm|Dùng cho branch cá nhân|

**Khi nào dùng Merge, khi nào dùng Rebase:**

- **Merge**: Khi merge vào branch chính (main, develop) hoặc branch được chia sẻ
- **Rebase**: Khi cập nhật branch cá nhân của bạn trước khi merge

---

### 3.11. Git Pull - Kéo code từ remote

```bash
# Pull từ remote (fetch + merge)
git pull

# Pull từ branch cụ thể
git pull origin main

# Pull với rebase thay vì merge
git pull --rebase

# Pull và tự động resolve conflict đơn giản
git pull --rebase --autostash
```

**Pull vs Fetch:**

```bash
# FETCH: Chỉ tải về, KHÔNG merge
git fetch origin
git log origin/main  # Xem commit mới
# Nếu OK, mới merge:
git merge origin/main

# PULL: Tải về VÀ merge luôn
git pull origin main
# Tương đương:
# git fetch origin
# git merge origin/main
```

---

### 3.12. Git Stash - Lưu tạm công việc

```bash
# Lưu thay đổi hiện tại (để chuyển branch)
git stash
# Hoặc:
git stash save "Đang làm tính năng X"

# Xem danh sách stash
git stash list

# Xem nội dung stash cụ thể
git stash show stash@{0}
git stash show -p stash@{0}  # Xem chi tiết

# Áp dụng stash gần nhất
git stash apply

# Áp dụng stash cụ thể
git stash apply stash@{1}

# Áp dụng và xóa stash
git stash pop

# Xóa stash cụ thể
git stash drop stash@{0}

# Xóa tất cả stash
git stash clear

# Tạo branch từ stash
git stash branch new-branch stash@{0}
```

**Ví dụ thực tế:**

```bash
# Bạn đang code feature-login, chưa xong
# Nhưng cần gấp chuyển sang main để fix bug

# Bước 1: Stash công việc hiện tại
git stash save "Đang làm form login"

# Bước 2: Chuyển sang main
git checkout main

# Bước 3: Fix bug, commit
git add .
git commit -m "Fix lỗi hiển thị"

# Bước 4: Quay lại feature-login
git checkout feature-login

# Bước 5: Lấy lại công việc đã stash
git stash pop
```

---

### 3.13. Git Reset - Quay lại commit trước

```bash
# Reset về commit trước, GIỮ thay đổi ở working directory
git reset --soft HEAD~1

# Reset về commit trước, GIỮ thay đổi nhưng bỏ khỏi staging
git reset --mixed HEAD~1
# Hoặc (--mixed là mặc định):
git reset HEAD~1

# Reset về commit trước, XÓA HOÀN TOÀN thay đổi
git reset --hard HEAD~1

# Reset về commit cụ thể
git reset --hard <commit_id>

# Reset file cụ thể về trạng thái của commit
git reset HEAD index.html
```

**So sánh các loại reset:**

|Lệnh|Working Directory|Staging Area|Local Repo|
|---|---|---|---|
|`--soft`|Giữ nguyên|Giữ nguyên|Reset|
|`--mixed` (mặc định)|Giữ nguyên|Reset|Reset|
|`--hard`|Reset|Reset|Reset|

**Ví dụ:**

```bash
# Trước reset:
# - Working: file đã sửa
# - Staging: đã add
# - Repo: đã commit

git reset --soft HEAD~1
# Sau: file vẫn sửa, vẫn ở staging, commit bị xóa

git reset --mixed HEAD~1
# Sau: file vẫn sửa, BỎ khỏi staging, commit bị xóa

git reset --hard HEAD~1
# Sau: file về trạng thái cũ, commit bị xóa (MẤT LUÔN THAY ĐỔI!)
```

---

### 3.14. Git Revert - Đảo ngược commit

```bash
# Tạo commit mới đảo ngược commit cũ
git revert <commit_id>

# Revert mà không tự động commit
git revert --no-commit <commit_id>

# Revert nhiều commit
git revert <commit1> <commit2>
```

**Reset vs Revert:**

|Reset|Revert|
|---|---|
|Xóa commit khỏi lịch sử|Giữ commit cũ, tạo commit mới đảo ngược|
|Viết lại lịch sử|Không viết lại lịch sử|
|NGUY HIỂM cho shared branch|AN TOÀN cho shared branch|
|Dùng cho commit chưa push|Dùng cho commit đã push|

**Ví dụ:**

```bash
# Commit C đã push lên, bị lỗi, cần bỏ
# A -> B -> C (lỗi)

# Cách 1: Reset (NGUY HIỂM nếu đã push)
git reset --hard B
git push -f  # Force push, GHI ĐÈ lịch sử remote

# Cách 2: Revert (AN TOÀN)
git revert C
# A -> B -> C -> D (D đảo ngược C)
git push  # Push bình thường
```

---

### 3.15. Git Clean - Xóa file chưa track

```bash
# Xem file nào sẽ bị xóa (dry run)
git clean -n

# Xóa file chưa track
git clean -f

# Xóa cả thư mục chưa track
git clean -fd

# Xóa cả file trong .gitignore
git clean -fx
```

---

### 3.16. Git Tag - Đánh dấu phiên bản

```bash
# Tạo tag cho commit hiện tại
git tag v1.0.0

# Tạo annotated tag (có thông tin chi tiết)
git tag -a v1.0.0 -m "Phiên bản 1.0.0"

# Tạo tag cho commit cũ
git tag v0.9.0 <commit_id>

# Xem danh sách tag
git tag

# Xem thông tin tag
git show v1.0.0

# Push tag lên remote
git push origin v1.0.0

# Push tất cả tag
git push origin --tags

# Xóa tag local
git tag -d v1.0.0

# Xóa tag trên remote
git push origin :refs/tags/v1.0.0
# Hoặc:
git push origin --delete v1.0.0

# Checkout về tag cụ thể
git checkout v1.0.0
```

---

## 4. Workflow thực tế

### 4.1. Workflow cơ bản cho người mới

```bash
# === Ngày đầu tiên - Setup ===
# Clone project về
git clone https://github.com/company/project.git
cd project

# Cấu hình (nếu chưa)
git config user.name "Tên bạn"
git config user.email "email@example.com"


# === Hàng ngày - Làm việc ===
# Bước 1: Cập nhật code mới nhất
git checkout main
git pull origin main

# Bước 2: Tạo branch cho tính năng mới
git checkout -b feature/add-login

# Bước 3: Code, code, code...
# ... sửa file ...

# Bước 4: Kiểm tra thay đổi
git status
git diff

# Bước 5: Add và commit
git add .
git commit -m "feat: thêm form đăng nhập"

# Bước 6: Push lên remote
git push -u origin feature/add-login

# Bước 7: Tạo Pull Request trên GitHub để review

# Bước 8: Sau khi được approve và merge
git checkout main
git pull origin main
git branch -d feature/add-login  # Xóa branch local
```

---

### 4.2. Workflow khi gặp conflict

```bash
# Bạn đang ở branch feature-login, muốn merge vào main
git checkout main
git pull origin main
git merge feature-login

# >>> CONFLICT! <<<

# Bước 1: Xem file nào bị conflict
git status

# Bước 2: Mở file bị conflict, tìm dòng:
# <<<<<<< HEAD
# code của main
# =======
# code của feature-login
# >>>>>>> feature-login

# Bước 3: Sửa file, giữ code muốn dùng

# Bước 4: Xóa các dấu <<<<, ====, >>>>

# Bước 5: Test code xem còn lỗi không

# Bước 6: Add file đã sửa
git add .

# Bước 7: Commit
git commit -m "Merge feature-login vào main"

# Bước 8: Push
git push origin main
```

---

### 4.3. Workflow khi làm hỏng code

```bash
# === Trường hợp 1: Chưa commit ===
# Muốn bỏ thay đổi của file
git restore index.html
# Hoặc:
git checkout -- index.html

# Bỏ TẤT CẢ thay đổi
git restore .


# === Trường hợp 2: Đã commit nhưng chưa push ===
# Muốn xóa commit cuối
git reset --soft HEAD~1  # Giữ code
# Hoặc:
git reset --hard HEAD~1  # Xóa luôn code


# === Trường hợp 3: Đã push ===
# KHÔNG nên dùng reset, dùng revert
git revert <commit_id>
git push origin main


# === Trường hợp 4: Muốn lấy lại file đã xóa ===
git checkout HEAD -- deleted-file.txt
```

---

### 4.4. Git Aliases - Tạo lệnh tắt

```bash
# Thiết lập alias
git config --global alias.st status
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.unstage 'reset HEAD --'
git config --global alias.last 'log -1 HEAD'
git config --global alias.lg "log --graph --oneline --all"

# Sử dụng:
git st        # thay vì git status
git co main   # thay vì git checkout main
git br        # thay vì git branch
git lg        # xem log đẹp
```

---

## 5. Các lỗi thường gặp và cách khắc phục

### Lỗi 1: "fatal: not a git repository"

**Nguyên nhân:** Chưa chạy `git init` hoặc không ở trong thư mục Git.

**Giải pháp:**

```bash
git init
```

---

### Lỗi 2: "error: failed to push some refs"

**Nguyên nhân:** Remote có commit mới mà local chưa có.

**Giải pháp:**

```bash
git pull origin main
# Giải quyết conflict nếu có
git push origin main
```

---

### Lỗi 3: "Your branch is ahead of 'origin/main' by X commits"

**Nguyên nhân:** Bạn có commit local chưa push.

**Giải pháp:**

```bash
git push origin main
```

---

### Lỗi 4: "Please commit your changes or stash them before you merge"

**Nguyên nhân:** Có thay đổi chưa commit, Git không cho chuyển branch.

**Giải pháp:**

```bash
# Cách 1: Commit
git add .
git commit -m "WIP: đang làm dở"

# Cách 2: Stash
git stash
git checkout other-branch
git stash pop
```

---

### Lỗi 5: "fatal: refusing to merge unrelated histories"

**Nguyên nhân:** Merge 2 repo không có lịch sử chung.

**Giải pháp:**

```bash
git pull origin main --allow-unrelated-histories
```

---

## 6. Best Practices (Thực hành tốt)

### ✅ Nên làm:

1. **Commit thường xuyên** với message rõ ràng
2. **Pull trước khi push** để tránh conflict
3. **Tạo branch cho mỗi tính năng** mới
4. **Review code trước khi merge** vào main
5. **Viết .gitignore** ngay từ đầu
6. **Dùng git status** thường xuyên
7. **Backup code** bằng cách push lên remote

### ❌ Không nên làm:

1. **Commit code lỗi** vào main
2. **Push trực tiếp** vào main (nên dùng Pull Request)
3. **Commit file nhạy cảm** (password, API key...)
4. **Force push** khi làm việc nhóm
5. **Commit file quá lớn** (video, binary...)
6. **Git add .** mà không kiểm tra file nào đã add
7. **Reset/Rebase** branch đã share với người khác