# 总结
## 获取节点
```golang
        # 获取a索引的前一个节点
        for i:=0;i<a-1;i++{
            temp1=temp1.Next
        }

        # 获取当前b索引节点
        for j:=0;j<b;j++{
            temp2=temp2.Next
        }
        
```

## 获取节点数
```golang
        n := 0
        cur := head
        for cur != nil {
            cur = cur.Next
            n++
        }
```