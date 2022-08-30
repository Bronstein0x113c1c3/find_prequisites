# Một chút ý tưởng để code lại

Thực ra hôm nay tui có vào được trang self-service của nhà trường, tức là tui đã vô học rồi, nhưng có một vài chuyện dở khóc dở cười. Và đó, cũng là lý do tui thực hiện dự án này.

Thì có một vấn đề chung của sinh viên TROY như tui thì khi đọc những môn điều kiện, thì chúng tôi trở nên hoảng sợ, nó như ma trận ám ảnh đám sinh viên vậy, bạn nhìn đống môn học điều kiện thì bạn chẳng biết bạn đang ở đâu, và cần gì, dẫn đến khá khó khăn cho việc định hướng môn học. 

Thì để bắt đầu cho bài viết này, tui sẽ giải thích nguyên lý và các component nha :))

```go
var prequisites map[string][]string = map[string][]string{
    "mth1125": []string{"mth1114"},
    "mth1126": []string{"mth1125"},
    "cs2250":  []string{"mth1112"},
    "cs2255":  []string{"cs2250"},
    "cs3310":  []string{"mth1112"},
    "cs3323":  []string{"cs2255", "mth1125"},
    "cs3329":  []string{"cs3323"},
    "cs3332":  []string{"cs3323"},
    "cs3360":  []string{"cs2255"},
    "cs3365":  []string{"cs3310"},
    "cs3370":  []string{"cs3323"},
    "cs3372":  []string{"cs3323"},
    "cs4420":  []string{"cs3323"},
    "cs4445":  []string{"cs3323"},
    "cs4448":  []string{"cs3323"},
    "mth2210": []string{"mth1112"},
    "mth2215": []string{"mth1112"},
    "mth1114": []string{"mth1112"},
    "mth1112": []string{""},
}
```

Ở mục prequisites này, tui tìm kiếm thông tin về môn học và các môn điều kiện, khi đó thì một môn điều kiện có thể có một hoặc nhiều môn điều kiện khác nhau, do đó tui dùng kiểu dữ liệu từ điển với ánh xạ là xâu -> một mảng xâu chứa các môn điều kiện. Các data này rất cần thiết cho việc duyệt và thay đổi sau này. Có thể gọi đây chính là cốt lõi của toàn bộ chương trình.

```go
func find_prequisite(x string) {
    fmt.Printf("prequisite for %v: \n", x)
    if prequisites[x][0] == "" {
        fmt.Println("no prequisites")
    } else {

        for i := 0; i < len(prequisites[x]); i++ {
            fmt.Println(prequisites[x][i])
            // in toàn bộ các môn prequisites
        }
        for i := 0; i < len(prequisites[x]); i++ {
            find_prequisite(prequisites[x][i])
            //duyệt đệ quy
        }
    }
}
```

Well, khi đọc code, chúng ta có thể thấy được đệ quy nhằm mục đích để duyệt các môn học và prequisites.

Xét ở data, nếu prequisites cho môn x là một xâu rỗng, thì khi đó, chương trình chẳng có prequisites nào cả, và ngược lại, sẽ duyệt toàn bộ prequisites của môn x và các môn điều kiện của các môn điều kiện ấy. Từ đó, prequisites của các môn sẽ được trình bày rõ ràng hơn để dễ quan sát.

Ở phần main, chúng ta có thể sửa file thành nhập từ CLI, hoặc là cho phép nhập xuất bằng lớp bufio,... rồi chạy để theo dõi quá trình ở hàm main

```batch
go run main.go
```

Và đến lúc thưởng thức chương trình :))

Sau này, có thể tui sẽ viết thêm trạng thái đạt hoặc không đạt của các môn cũng bằng cách sử dụng từ điển.

Đó là toàn bộ project mà tui nghiên cứu được trong cả buổi chiều nay, cảm ơn vì đã đọc

Love from Bronstein
