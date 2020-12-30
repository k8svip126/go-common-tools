package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
)

func main() {
	redisCli, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer redisCli.Close()

	// 读：n,err := c.Do("hset","key","field","value")
	// 写：result,err := redis.Values(c.Do("hgetall","key"))

	// set ,get
	_,err = redisCli.Do("set","k3s","abc")
	if err != nil {
		fmt.Println(err.Error())
	}else{
		k3s,err := redisCli.Do("get","k3s")
		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Println(string(k3s.([]byte)))
		}
	}

	// hset, hget
	_, err = redisCli.Do("hset", "k8s", "vip", "k8svip")
	if err != nil {
		fmt.Println(err)
	} else {
		a, err := redisCli.Do("hget", "k8s", "vip")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(reflect.TypeOf(a))
			// 类型断言，变量名.([]byte),
			// value, ok := val.([]byte)
			// val一般是一个interface{}类型的变量，这句的字面含义是"我认为val是这个interface{}类型变量的underlying type是[]byte，如果是，将其值赋给变量value，并且ok =true，如果不是ok = false"
			fmt.Println(string(a.([]byte)))
		}
	}

	aa, err := redis.Bytes(redisCli.Do("hget", "k8s", "vip"))
	fmt.Println("aa:", string(aa))
	//hmset，hmget
	_, err = redisCli.Do("hmset", "k8s", "k1", "v1", "k2", "v2", "k3", "v3")
	if err != nil {
		fmt.Println(err)
	} else {
		values, err := redis.Values(redisCli.Do("hmget", "k8s", "k1", "k2"))
		if err != nil {
			fmt.Println(err)
		} else {
			for _, v := range values {
				fmt.Printf("%s ", v.([]byte))
			}
			fmt.Println("\n")
		}
	}

	// hexists
	isExist, err := redisCli.Do("hexists", "k8s", "k2")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// 存在为1，不存在为0
		fmt.Println(isExist)
	}

	// hlen
	ilen, err := redisCli.Do("hlen", "k8s")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ilen)
	}

	// hkeys
	keys, err := redis.Values(redisCli.Do("hkeys", "k8s"))
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range keys {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println("\n")
	}

	// hvals
	vals, err := redis.Values(redisCli.Do("hvals", "k8s"))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, v := range vals {
			fmt.Printf("%s ", v.([]byte))
		}
	}

	// hgetall
	result, err := redis.Values(redisCli.Do("hgetall", "k8s"))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("hgetall")
		for _, v := range result {
			fmt.Printf(": %s ", v.([]byte))
		}
	}

	// hdel
	_, err = redisCli.Do("hdel", "k8s", "k3")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\n")
	isExist, err = redisCli.Do("hexists", "k8s", "k3")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// 存在为1，不存在为0
		fmt.Println(isExist)
	}

	// 删除所有
	keys, err = redis.Values(redisCli.Do("hkeys", "k8s"))
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range keys {
			_, err = redisCli.Do("hdel", "k8s", string(v.([]byte)))
		}
		fmt.Println("\n")
	}
	keys, err = redis.Values(redisCli.Do("hkeys", "k8s"))
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range keys {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println("\n")
	}

}
