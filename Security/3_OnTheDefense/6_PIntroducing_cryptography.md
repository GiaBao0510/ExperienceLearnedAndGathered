# Giới thiệu về Mật mã học (Cryptography)

## Giới thiệu

Mật mã học là nền tảng toán học của nhiều khái niệm cốt lõi trong an ninh thông tin. Mọi chuyên gia an ninh mạng đều cần hiểu các nguyên lý cơ bản của lĩnh vực này để làm việc hiệu quả.

### Mật mã học là gì?

Theo định nghĩa của Viện Tiêu chuẩn và Công nghệ Quốc gia Hoa Kỳ (NIST):

![](https://content.kaspersky-labs.com/fm/press-releases/d9/d9de508916a1aa74fef26ded2c93a839/processed/what-is-cryptography-1-q75.jpg)

> Mật mã học là các nguyên tắc, phương tiện và phương pháp biến đổi dữ liệu để che giấu nội dung, ngăn chặn việc sử dụng trái phép, hoặc ngăn chặn việc sửa đổi mà không bị phát hiện. — [National Institute of Standards and Technology](https://csrc.nist.gov/glossary/term/cryptography)

Giữ bí mật thông tin và chia sẻ nó một cách an toàn là thách thức đã tồn tại hàng nghìn năm. Dù phương pháp thực hiện đã thay đổi đáng kể — từ mật mã cổ điển đến thuật toán toán học phức tạp ngày nay — mục tiêu cốt lõi vẫn không đổi.

---

## Định nghĩa Giao tiếp An toàn (Secure Communications)

Để minh họa các khái niệm mật mã, cộng đồng an ninh mạng thường dùng ba nhân vật quy ước: **Alice** (người gửi), **Bob** (người nhận), và **Eve** (kẻ nghe lén — "eavesdropper"). Alice và Bob muốn giao tiếp an toàn; Eve muốn nghe lén hoặc can thiệp vào cuộc trao đổi đó.

![](https://upload.wikimedia.org/wikipedia/commons/7/7c/Alice-bob-mallory.jpg)

Giao tiếp an toàn đáng tin cậy cần đảm bảo ba đặc tính:

**Tính bảo mật (Confidentiality):** Alice gửi tin nhắn cho Bob mà Eve — dù chặn được tin nhắn — không thể hiểu được nội dung.

**Tính xác thực (Authenticity):** Eve không thể gửi tin nhắn giả danh Alice mà Bob tin là thật. Tính xác thực đảm bảo việc mạo danh là bất khả thi (hoặc cực kỳ khó khăn).

**Tính toàn vẹn (Integrity):** Nếu Eve sửa đổi tin nhắn giữa đường, Bob có thể phát hiện được sự thay đổi đó. Lưu ý rằng có thể can thiệp vào một tin nhắn mà không cần hiểu nội dung của nó — ví dụ, một người có thể làm gián đoạn một cuộc trò chuyện trực tiếp bằng tiếng ồn, dù họ không hiểu ngôn ngữ đang được nói.

> Ba đặc tính này đạt được thông qua nhiều thuật toán toán học và kỹ thuật khác nhau. Trong lịch sử, các phương pháp bao gồm hộp khóa vật lý và niêm phong bằng sáp; ngày nay, trọng tâm chuyển sang các giải pháp toán học.

---

## Mã hóa (Encryption)

Mã hóa là phương pháp tiêu chuẩn để bảo vệ **tính bảo mật** của thông điệp — quá trình chuyển đổi thông điệp thành dạng chỉ những ai có khóa giải mã mới đọc được.

**Thuật ngữ cơ bản:**

- **Plaintext (văn bản gốc):** Dữ liệu ở dạng dễ đọc, chưa qua mã hóa.
- **Ciphertext (văn bản mã hóa):** Dữ liệu sau khi đã được mã hóa, không thể đọc trực tiếp.
- **Thuật toán mã hóa (encryption algorithm):** Tập hợp các phép biến đổi toán học chuyển plaintext thành ciphertext.
- **Khóa (key):** Tham số điều khiển thuật toán thực hiện phép biến đổi theo cách cụ thể nào.

> (plaintext/ciphertext) — đây là thuật ngữ chuẩn sẽ xuất hiện xuyên suốt tài liệu kỹ thuật và code (ví dụ: tham số `plaintext []byte` trong các thư viện mã hóa Go). Đã bổ sung để người đọc nhận diện được khi gặp trong thực tế lập trình.

Mã hóa hiện đại chia thành hai loại chính: **đối xứng (symmetric)** và **bất đối xứng (asymmetric)**.

---

### Mã hóa đối xứng (Symmetric Encryption)

Trong mã hóa đối xứng, **cùng một khóa** được dùng cho cả mã hóa lẫn giải mã. Đặc điểm: tốc độ nhanh, triển khai đơn giản, nhưng đòi hỏi người gửi và người nhận phải **cùng sở hữu một bí mật chung (shared secret)** — tương tự việc dùng chung một mật khẩu.

![](https://www.ssl2buy.com/wp-content/uploads/2015/12/Symmetric-Encryption.png)

> **Ví dụ minh họa — Mật mã Caesar (Caesar Cipher):**
> 
> Một ví dụ đơn giản là mật mã dựa trên phép xoay bảng chữ cái (substitution cipher), trong đó mỗi ký tự được dịch chuyển một số vị trí cố định. Số vị trí dịch chuyển đóng vai trò là khóa. Nếu người gửi dùng khóa +1 (dịch mỗi ký tự tiến 1 vị trí), người nhận dùng khóa -1 (dịch lùi 1 vị trí) để khôi phục thông điệp gốc.
> 
> Ví dụ: từ `HOLIDAY` được mã hóa với khóa +1 thành `IPMJEBZ` (H→I, O→P, L→M, I→J, D→E, A→B, Y→Z).
> 
> **Lưu ý quan trọng:** Đây được gọi cụ thể là **mật mã Caesar (Caesar Cipher)** — kỹ thuật minh họa nguyên lý cơ bản của mã hóa đối xứng, nhưng **hoàn toàn không an toàn** trong thực tế hiện đại vì chỉ có tối đa 25 khóa khả dĩ (dễ dàng brute-force trong tích tắc) và dễ bị phá bằng phân tích tần suất chữ cái (frequency analysis). Mật mã Caesar chỉ có giá trị giáo dục, không nên nhầm lẫn với các thuật toán mã hóa đối xứng hiện đại như AES.



**AES (Advanced Encryption Standard)** là thuật toán mã hóa đối xứng được sử dụng phổ biến nhất hiện nay. AES sử dụng một khóa bí mật để chuyển đổi plaintext thành ciphertext; chỉ ai sở hữu đúng khóa mới giải mã được. Tổ chức sử dụng AES để bảo vệ dữ liệu nhạy cảm: mật khẩu, giao dịch tài chính, và dữ liệu lưu trữ (data at rest).

> AES thường được triển khai với độ dài khóa **128-bit, 192-bit, hoặc 256-bit** — con số càng lớn, độ an toàn càng cao nhưng tốc độ xử lý giảm nhẹ. **AES-256** là tiêu chuẩn được khuyến nghị cho dữ liệu có độ nhạy cảm cao. Trong Go, package `crypto/aes` kết hợp với `crypto/cipher` (thường dùng chế độ **GCM — Galois/Counter Mode**) là cách triển khai chuẩn để vừa mã hóa vừa xác thực tính toàn vẹn dữ liệu cùng lúc (authenticated encryption).

---

### Mã hóa bất đối xứng (Asymmetric Encryption)

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQPOghhe9wgGYIlZ6G2wV9SbCuKpaUygnOZ8GeCLuW2iOQfVS2pDvYFPpIN&s=10)

Mã hóa bất đối xứng — còn gọi là **mã hóa khóa công khai (public-key encryption)** — dùng **hai khóa khác nhau**: một để mã hóa, một để giải mã. Hai khóa này được tạo ra **cùng lúc theo một cặp toán học liên kết chặt chẽ**:

- **Khóa công khai (public key):** Có thể chia sẻ tự do với bất kỳ ai.
- **Khóa riêng tư (private key):** Phải được giữ bí mật tuyệt đối, không bao giờ chia sẻ.

Bất kỳ ai có khóa công khai của bạn đều có thể mã hóa tin nhắn gửi cho bạn — nhưng chỉ khóa riêng tư tương ứng (chỉ bạn sở hữu) mới giải mã được.

> **Ví dụ — Quy trình gửi/nhận:**
> 
> 1. Alice mã hóa tin nhắn bằng **khóa công khai của Bob**.
> 2. Alice gửi tin nhắn đã mã hóa cho Bob.
> 3. Bob giải mã tin nhắn bằng **khóa riêng tư của mình**.
> 4. Bob đọc được nội dung tin nhắn.
> 
> Nếu Bob làm lộ khóa riêng tư, bất kỳ ai có được nó đều có thể đọc mọi tin nhắn gửi đến Bob — vì vậy khóa riêng tư tuyệt đối không được chia sẻ.

**Lợi ích cốt lõi:** Mã hóa bất đối xứng cho phép hai bên **giao tiếp an toàn mà không cần trao đổi khóa bí mật trước đó** — giải quyết được bài toán khó nhất của mã hóa đối xứng (làm sao trao đổi shared secret một cách an toàn ngay từ đầu, đặc biệt qua kênh không an toàn như internet).

> **Ví dụ thực tế:** Khi mua sắm trực tuyến, bạn không cần gặp mặt trực tiếp cửa hàng để thiết lập một khóa bí mật chung trước khi giao dịch. Nếu chỉ có mã hóa đối xứng, việc thiết lập khóa dùng chung an toàn qua internet sẽ là bài toán gần như bất khả thi ở quy mô lớn — đây chính là lý do mã hóa bất đối xứng là nền tảng giúp thương mại điện tử và giao tiếp an toàn trên internet (HTTPS/TLS) trở nên khả thi.

> **Sự đánh đổi tốc độ:** mã hóa bất đối xứng **chậm hơn đáng kể** so với mã hóa đối xứng (do phép toán dựa trên số nguyên tố lớn hoặc đường cong elliptic phức tạp hơn nhiều so với phép XOR/hoán vị của AES). Vì vậy, trong thực tế, **TLS/HTTPS không mã hóa toàn bộ dữ liệu bằng mã hóa bất đối xứng** — mà dùng mô hình lai (hybrid encryption): mã hóa bất đối xứng chỉ được dùng để trao đổi an toàn một khóa đối xứng tạm thời (session key) lúc bắt đầu kết nối (TLS handshake), sau đó toàn bộ dữ liệu truyền tải thực tế được mã hóa bằng thuật toán đối xứng (thường là AES) vì tốc độ nhanh hơn nhiều. Đây là kiến thức nền tảng quan trọng để hiểu cách HTTPS thực sự hoạt động .

**Thuật toán bất đối xứng phổ biến:** RSA (dựa trên độ khó của phân tích thừa số nguyên tố lớn) và ECC — Elliptic Curve Cryptography (dựa trên bài toán đường cong elliptic, cho độ an toàn tương đương RSA nhưng khóa ngắn hơn nhiều, hiệu quả hơn cho thiết bị di động và IoT).

### Hoạt động thực hành: Sắp xếp các bước của mã hóa bất đối xứng

1. Người gửi sử dụng khóa công khai của người nhận để mã hóa tin nhắn.
2. Người gửi gửi tin nhắn đã mã hóa.
3. Người nhận sử dụng khóa riêng tư của họ để giải mã tin nhắn.
4. Người nhận đọc được nội dung tin nhắn.

---

## Mã hóa Lượng tử (Quantum-Resistant Cryptography)

Máy tính cổ điển hiện nay không đủ khả năng phá vỡ các thuật toán mã hóa hiện đại trong thời gian thực tế. Tuy nhiên, **điện toán lượng tử (quantum computing)** — công nghệ đang phát triển nhanh chóng dựa trên các định luật cơ học lượng tử — có tiềm năng thay đổi cục diện này trong tương lai.

> Điện toán lượng tử xử lý thông tin bằng cách sử dụng cơ học lượng tử để giải quyết các vấn đề quá phức tạp đối với máy tính cổ điển. — [IBM, "What is quantum computing?"](https://www.ibm.com/topics/quantum-computing)

### Mối đe dọa cụ thể của điện toán lượng tử đối với mật mã học

> **Mã hóa bất đối xứng (RSA, ECC) bị đe dọa nghiêm trọng:** Máy tính lượng tử đủ mạnh chạy **thuật toán Shor (Shor's Algorithm)** có thể phân tích thừa số nguyên tố lớn (nền tảng của RSA) hoặc giải bài toán logarit rời rạc trên đường cong elliptic (nền tảng của ECC) trong thời gian đa thức — thay vì thời gian gần như vô hạn với máy tính cổ điển. Điều này khiến RSA và ECC trở nên **hoàn toàn không an toàn** một khi máy tính lượng tử đủ mạnh xuất hiện.
> 
> - **Mã hóa đối xứng (AES) bị đe dọa ở mức độ thấp hơn nhiều:** Thuật toán lượng tử **Grover (Grover's Algorithm)** chỉ giúp tăng tốc tấn công brute-force theo căn bậc hai — về bản chất làm giảm độ an toàn hiệu quả của khóa đi một nửa số bit (ví dụ AES-256 chỉ còn tương đương độ an toàn của AES-128 trước máy tính lượng tử). Giải pháp đơn giản: tăng độ dài khóa (dùng AES-256 thay vì AES-128) là đủ để duy trì an toàn trước mối đe dọa lượng tử.
> 
> "khả năng phân tích thừa số nguyên tố lớn" — đây thực ra là mô tả chính xác cho **RSA cụ thể**, không phải toàn bộ "mật mã hiện có" như câu văn ngụ ý. Đã làm rõ để tránh gây hiểu lầm rằng toàn bộ hệ thống mã hóa (bao gồm cả AES) đều bị đe dọa như nhau.

### Mật mã học Hậu lượng tử (Post-Quantum Cryptography — PQC)

Để chuẩn bị cho mối đe dọa này, giới nghiên cứu mật mã học đã phát triển các thuật toán **kháng lượng tử (quantum-resistant)** — được thiết kế để an toàn trước cả máy tính cổ điển lẫn máy tính lượng tử. Bốn hướng tiếp cận toán học chính:

#### Lattice-based Encryption (Mã hóa dựa trên mạng lưới)

Dựa trên độ khó của các bài toán hình học trong không gian đa chiều (lattice problems) — ngay cả máy tính lượng tử cũng chưa tìm ra thuật toán hiệu quả để giải các bài toán này.

#### Hash-based Encryption (Mã hóa dựa trên hàm băm)

Sử dụng hàm băm mật mã học (cryptographic hash function) để tạo chữ ký số an toàn. Do bản chất một chiều của hàm băm (không thể đảo ngược), phương pháp này khó bị phá ngay cả với máy tính lượng tử.

#### Multivariate Encryption (Mã hóa đa biến)

Dựa trên độ khó của việc giải hệ phương trình đa biến bậc cao đồng thời — bài toán được chứng minh là NP-hard, tạo ra rào cản toán học vững chắc.

#### Code-based Encryption (Mã hóa dựa trên mã sửa lỗi)

Dựa trên lý thuyết mã sửa lỗi (error-correcting codes) — giải mã một mã ngẫu nhiên tổng quát là bài toán khó ngay cả với máy tính lượng tử.

> vào **tháng 8/2024, NIST đã chính thức công bố ba tiêu chuẩn mật mã hậu lượng tử đầu tiên**, đánh dấu bước chuyển từ nghiên cứu sang triển khai thực tế:
> 
> - **ML-KEM (dựa trên CRYSTALS-Kyber)** — thuộc nhóm lattice-based, dùng cho trao đổi khóa (key encapsulation).
> - **ML-DSA (dựa trên CRYSTALS-Dilithium)** — thuộc nhóm lattice-based, dùng cho chữ ký số.
> - **SLH-DSA (dựa trên SPHINCS+)** — thuộc nhóm hash-based, dùng cho chữ ký số.
> 
> Đây thông tin cực kỳ quan trọng: mật mã hậu lượng tử không còn chỉ là lý thuyết — các tổ chức lớn (Google, Cloudflare, Apple) đã bắt đầu triển khai các thuật toán này trong TLS và hệ thống nhắn tin thực tế. Backend engineer nên bắt đầu làm quen với các thuật toán này vì chúng sẽ dần thay thế RSA/ECC trong nhiều hệ thống trong thập kỷ tới.

---

## Bổ sung kiến thức

### 1. Package mã hóa tiêu chuẩn trong Go

Với đối tượng đọc là backend/Golang developer, nên biết các package chuẩn của Go liên quan trực tiếp đến nội dung tài liệu:

```go
// Mã hóa đối xứng - AES với chế độ GCM (Authenticated Encryption)
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
)

// Mã hóa bất đối xứng - RSA
import "crypto/rsa"

// Mã hóa bất đối xứng hiện đại - Elliptic Curve
import "crypto/ecdsa"
import "crypto/ed25519" // Chữ ký số dựa trên đường cong Edwards, nhanh và an toàn

// Hàm băm mật mã học
import "crypto/sha256"
```

**Khuyến nghị thực tế:** Với hệ thống mới, ưu tiên `ed25519` cho chữ ký số (nhanh hơn RSA, khóa ngắn hơn, được thiết kế để tránh nhiều lỗi triển khai phổ biến) và `AES-GCM` cho mã hóa đối xứng thay vì các chế độ cũ hơn như AES-CBC (dễ mắc lỗi bảo mật nếu triển khai không đúng, đặc biệt là padding oracle attack).

### 2. TLS/HTTPS — Ứng dụng thực tế kết hợp cả hai loại mã hóa

Đây là ví dụ ứng dụng thực tế quan trọng nhất mà bản gốc không đề cập: khi trình duyệt kết nối HTTPS đến server, quá trình **TLS Handshake** diễn ra như sau:

1. Client và server dùng mã hóa **bất đối xứng** để xác thực danh tính (chứng chỉ số) và thỏa thuận một khóa phiên (session key) chung một cách an toàn.
2. Sau khi khóa phiên được thiết lập, toàn bộ dữ liệu thực tế được truyền tải bằng mã hóa **đối xứng** (thường AES) vì tốc độ nhanh hơn nhiều lần.

Hiểu rõ mô hình lai (hybrid) này giúp giải thích tại sao HTTPS vừa an toàn (nhờ bất đối xứng cho xác thực) vừa hiệu quả (nhờ đối xứng cho truyền dữ liệu khối lượng lớn).

### 3. Phân biệt Encryption và Hashing — Nhầm lẫn phổ biến

Tài liệu tập trung vào encryption (mã hóa — có thể đảo ngược với đúng khóa), nhưng cần phân biệt rõ với **hashing** (băm — quá trình một chiều, không thể đảo ngược):

| Đặc điểm           | Encryption                   | Hashing                               |
| ------------------ | ---------------------------- | ------------------------------------- |
| Khả năng đảo ngược | Có (với đúng khóa)           | Không (one-way)                       |
| Mục đích           | Bảo mật dữ liệu, đọc lại sau | Xác minh tính toàn vẹn, lưu mật khẩu  |
| Ví dụ thuật toán   | AES, RSA                     | SHA-256, bcrypt, Argon2               |
| Ứng dụng điển hình | Mã hóa file, HTTPS traffic   | Lưu trữ mật khẩu, checksum, chữ ký số |

Đây là lỗi phổ biến của lập trình viên mới: **không bao giờ "mã hóa" mật khẩu để lưu trữ** — phải dùng hàm băm chuyên dụng cho mật khẩu (bcrypt, Argon2, scrypt) vì các hàm này được thiết kế cố tình chậm để chống lại tấn công brute-force, khác hoàn toàn với mục đích của encryption thông thường.
