package medium

type LFUCache struct {
	cap int
	minFre int
	kv map[int]*Node
	fv map[int]*DoubleList
}


func ConstructorLFU(capacity int) LFUCache {
	kvHash := make(map[int]*Node)
	fvHash := make(map[int]*DoubleList)
	return LFUCache{cap:capacity,minFre:0,kv:kvHash,fv:fvHash}
}


func (this *LFUCache) Get(key int) int {
	//fmt.Println("Get: ",key)
	if v,ok := this.kv[key];ok{
		fre := v.fre
		// fv[fre] 双向链表进行删除节点
		this.fv[fre].RemoveMid(v)
		if this.fv[fre].size == 0{
			delete(this.fv,fre)
			if this.minFre == fre{
				this.minFre++
			}
		}
		v.fre++
		// fv[fre+1] 双向链表增加新节点

		if newFre,ok := this.fv[v.fre];ok{
			// 双向链表存在，该节点置于最后
			newFre.addLast(v)
		}else{
			this.fv[v.fre] = &DoubleList{beg:v,end:v,size:1}
			v.next = nil
			v.prev = nil
		}
		/*
		   fmt.Println("++kv++")
		   for k,v := range this.kv{
		       fmt.Printf("%v %v %v\n",k,v.val,v.fre)
		   }
		   fmt.Println("++fv++",this.minFre)
		   for k,v := range this.fv{
		       fmt.Println("++fre：",k,v.size)
		       for i:=v.beg;i!=nil;i=i.next{
		           fmt.Printf("%v %v %v\n",i.key,i.val,i.fre)
		       }
		   }
		   fmt.Println("++++")
		*/
		return v.val
	}
	return -1
}



func (this *LFUCache) Put(key int, value int)  {
	//fmt.Println("Put: ",key,value)
	if this.cap == 0{return }
	val := this.Get(key)
	if val == -1{
		// 向cache中添加元素
		tmp := &Node{key:key,val:value,fre:1,prev:nil,next:nil}
		if len(this.kv)>=this.cap{
			delNode := this.fv[this.minFre].beg
			if this.fv[this.minFre].size == 1{
				delete(this.fv,this.minFre)
			}else{
				this.fv[this.minFre].RemoveFirst()
			}
			delete(this.kv,delNode.key)
		}
		this.kv[key] = tmp
		if newFre,ok := this.fv[1];ok{
			newFre.addLast(tmp)
		}else{
			dTmp := &DoubleList{beg:tmp,end:tmp,size:1}
			tmp.prev = nil
			this.fv[1] = dTmp
		}
		this.minFre = 1
	}else{
		this.kv[key].val = value
	}
	/*
	   fmt.Println("==kv==")
	   for k,v := range this.kv{
	       fmt.Printf("%v %v %v\n",k,v.val,v.fre)
	   }
	   fmt.Println("==fv==",this.minFre)
	   for k,v := range this.fv{
	       fmt.Println("==fre：",k,v.size)
	       for i:=v.beg;i!=nil;i=i.next{
	           fmt.Printf("%v %v %v\n",i.key,i.val,i.fre)
	       }
	   }
	   fmt.Println("====")
	*/

}

// 双hash表+双向链表
type Node struct{
	key int
	val int
	fre int
	prev,next *Node
}


type DoubleList struct{
	beg,end *Node
	size int
}

func (dl *DoubleList) RemoveFirst()  {
	/*
	   fmt.Println("xxxxxx",dl.size)
	   for i:=dl.beg;i!=nil;i=i.next{
	       fmt.Printf("%v %v %v\n",i.key,i.val,i.fre)
	   }
	*/
	if dl.size == 0{return}
	if dl.size == 1{
		dl.size--
		dl.beg= nil
		dl.end=nil
	}else{
		tmp := dl.beg.next
		tmp.prev = nil
		dl.beg = tmp
		dl.size--
	}
}

func (dl *DoubleList) RemoveMid(node *Node)  {
	if dl.size == 0{return }
	if dl.size == 1{
		dl.beg = nil
		dl.end = nil
		dl.size--
	}else{
		front := node.prev
		back := node.next
		if front == nil{
			dl.beg = back
			back.prev = nil
		}else if back == nil{
			dl.end = front
			front.next = nil
		}else{
			front.next = back
			back.prev = front
		}
		dl.size--
	}
}

func (dl *DoubleList) addLast(node *Node)  {
	if dl.size == 0{
		dl.beg = node
		dl.end = node
		node.prev = nil
		node.next = nil
		dl.size++
	}else{
		dl.end.next = node
		node.prev = dl.end
		dl.end = node
		node.next = nil
		dl.size++
	}
}