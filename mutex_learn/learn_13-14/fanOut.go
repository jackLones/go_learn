package main

func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) []chan interface{} {
	go func() {
		defer func() { // 退出的时候，关闭所有输出通道
			for _, c := range out {
				close(c)
			}
		}()
	}()
	for v := range ch { //从输入通道读取数据
		v := v
		for i := 0; i < len(out); i++ {
			i := i
			if async {
				go func() {
					out[i] <- v
				}()
			} else {
				out[i] <- v
			}
		}
	}
	return nil
}
