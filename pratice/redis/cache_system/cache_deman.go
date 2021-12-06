package cache_system

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

//常量
const (
	//没有过期时间标记
	NoExpiration time.Duration =-1

	//默认过期时间
	DefaultExpiration time.Duration = 0
)

//数据项
type Item struct {
	Object interface{}   //data
	Expiration int64    // 生存时间
}


type Cache struct {
	defaultExpiration time.Duration
	items  map[string]Item   //缓存数据
	mu sync.RWMutex   //读写锁
	gcInterval time.Duration   //过期数据清理周期
	stopGc chan bool
}

//判断数据是否过期
func (item Item) Expired() bool{
	if item.Expiration == 0{
		return false
	}
	//对比时间
	return time.Now().UnixNano() > item.Expiration

}


//过期缓存清理
func (c * Cache) gcLoop(){
	ticker := time.NewTicker(c.gvInterval)
	for {
		select{
		case <- ticker.C:
			c.DeleteExpired()
		case <- c.stopGc:
			ticker.Stop()
			return
		}
	}
}

//删除缓存数据项
func (c * Cache) delete(key string){
	delete(c.items,key)

}

//删除过期数据项
func (c * Cache) DeleteExpired(){
	now := time.Now().UnixNano()
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v :=range c.items{
		if v.Expiration >0 && now > v.Expiration{
			c.delete(k)
		}
	}
}

//设置缓存数据，存在则覆盖
func (c *Cache) Set(k string, v interface{},d time.Duration){
	var e int64
	if d == DefaultExpiration{
		d =c.defaultExpiration
	}
	if d >0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[k] = Item{
		Object: v,
		Expiration: e,
	}
}


//获取数据项，且判断数据项是否过期
func (c * Cache)  get(k string) (interface{}, bool){
	item,ok :=c.items[k]

	if !ok{
		return nil,false
	}
	if item.Expired(){
		return nil ,false
	}
	return item.Object,true
}

//添加数据项，存在则返回错误
func (c *Cache) Add(k string, v interface{},d time.Duration) error{
	c.mu.Lock()
	_,found := c.get(k)
	if found{
		c.mu.Unlock()
		return fmt.Errorf("Item %s is alreadly exits!",k)

	}
	c.Set(k,v,d)
	c.mu.Unlock()

	return nil
}

//替换数据项
func (c *Cache)Replace(k string, v interface{},d time.Duration) error{
	c.mu.Lock()
	_,found := c.get(k)
	if !found{
		c.mu.Unlock()
		return fmt.Errorf("item %s is not exites",k)

	}else{
		c.Set(k,v,d)
		c.mu.Unlock()
		return nil
	}


}
// 删除一个数据项
func (c *Cache) Delete(k string) {
	c.mu.Lock()
	c.delete(k)
	c.mu.Unlock()
}

//将缓存写入 writor
func (c *Cache)Save(w io.Writer) (err error){

	enc := gob.NewEncoder(w)
	defer func() {
		if x:=recover();x !=nil{
			err = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	//读写锁
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _,v := range c.items{
		gob.Register(v.Object)

	}
	err = enc.Encode(&c.items)
	return

}

//保存数据到文件夹中

func (c *Cache)SaveFile(filename string) error{
	f, err := os.Create(filename)
	if err !=nil{

		return err
	}
	err = c.Save(f)
	if err!=nil{
		f.Close()
		return err
	}
	return f.Close()
}

//从io.reader 读取数据
func (c *Cache)Load(r io.Reader)(err error){

	dec :=gob.NewDecoder(r)

	items :=map[string]Item{}
	err = dec.Decode(&items)
	if err == nil {
		c.mu.Lock()
		defer c.mu.Unlock()
		for k,v := range items{
			ov,found := c.items[k]
			if found || ov.Expired(){
				c.items[k] = v
			}

		}
	}
	return err
}

//从文件中加载数据
func(c *Cache) LoadFile(filename string) (err error){
	f, err := os.Open(filename)
	if err !=nil{

		return err
	}
	if err = c.Load(f); err!=nil {
		f.Close()
		return err
	}
	return err
}

//统计缓存数据数量
func (c *Cache) Count() int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}

//清空缓存
func (c *Cache) Flush(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = map[string]Item{}

}

//停止过期缓存清理
func (c *Cache)StopGc(){
	c.stopGc <- true
}


//创建一个缓存系统
func NewCache(defaultExpiration,gcInterval time.Duration) * Cache{
	c := &Cache{
		defaultExpiration: defaultExpiration,
		gcInterval:        gcInterval,
		items:             map[string]Item{},
		stopGc:            make(chan bool),
	}

	//开启定期清理
	go c.gcLoop()
	return c
}



