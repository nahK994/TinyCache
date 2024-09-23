package cache

import (
	"strconv"
)

func InitCache() *Cache {
	return &Cache{
		info: make(map[string]interface{}),
	}
}

func (c *Cache) GET(key string) interface{} {
	c.mu.RLock()         // Acquire read lock
	defer c.mu.RUnlock() // Release read lock
	return c.info[key]
}

func (c *Cache) SET(key string, value interface{}) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock
	str, ok_str := value.(string)
	if ok_str {
		num, err := strconv.Atoi(str)
		if err == nil {
			c.info[key] = num
		} else {
			c.info[key] = str
		}
	} else {
		c.info[key] = value
	}
}

func (c *Cache) EXISTS(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, isKeyExists := c.info[key]
	return isKeyExists
}

func (c *Cache) DEL(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.info, key)
}

func (c *Cache) INCR(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, _ := c.info[key].(int)
	c.info[key] = val + 1
	return val + 1
}

func (c *Cache) DECR(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, _ := c.info[key].(int)
	c.info[key] = val - 1
	return val - 1
}

func (c *Cache) LPUSH(key string, values []string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	oldData, _ := c.info[key].([]string)
	vals := make([]string, len(values)+len(oldData))
	for i := 0; i < len(values); i++ {
		vals[i] = values[len(values)-1-i]
	}
	copy(vals[len(values):], oldData)
	oldData = nil
	c.info[key] = vals
}

func (c *Cache) RPUSH(key string, values []string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	oldData, _ := c.info[key].([]string)
	vals := make([]string, len(values)+len(oldData))
	copy(vals[0:], oldData)
	copy(vals[len(oldData):], values)
	oldData = nil
	c.info[key] = vals
}

func processIdx(vals []string, idx int) int {
	if idx > len(vals) {
		idx = len(vals) - 1
	} else if idx < 0 {
		if -1*len(vals) > idx {
			idx = 0
		} else {
			idx = len(vals) + idx
		}
	}

	return idx
}

func (c *Cache) LRANGE(key string, startIdx, endIdx int) []string {
	c.mu.RLock()         // Acquire write lock
	defer c.mu.RUnlock() // Release write lock

	vals, _ := c.info[key].([]string)
	startIdx = processIdx(vals, startIdx)
	endIdx = processIdx(vals, endIdx)
	var ans []string
	if len(vals) > 0 {
		for i := startIdx; i <= endIdx; i++ {
			ans = append(ans, vals[i])
		}
	}
	return ans
}

func (c *Cache) LPOP(key string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	vals, _ := c.info[key].([]string)

	newVals := make([]string, len(vals)-1)
	copy(newVals, vals[1:])

	c.info[key] = newVals
	vals = nil
}

func (c *Cache) RPOP(key string) {
	c.mu.Lock()         // Acquire write lock
	defer c.mu.Unlock() // Release write lock

	vals, _ := c.info[key].([]string)

	newVals := make([]string, len(vals)-1)
	copy(newVals, vals[0:len(vals)-1])

	c.info[key] = newVals
	vals = nil
}

func (c *Cache) FLUSHALL() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k := range c.info {
		delete(c.info, k)
	}
}
