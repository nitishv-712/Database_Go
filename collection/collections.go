package collection

type JSONData struct {
	Data map[string]interface{}
}

type Collection struct {
	Data []JSONData
}

// Add a new dataset to the collection
func (c *Collection) Add(data JSONData) {
	c.Data = append(c.Data, data)
}

// Get dataset by index
func (c *Collection) Get(index int) (*JSONData, bool) {
	if index < 0 || index >= len(c.Data) {
		return nil, false
	}
	return &c.Data[index], true
}

// Remove dataset by index
func (c *Collection) Remove(index int) bool {
	if index < 0 || index >= len(c.Data) {
		return false
	}
	c.Data = append(c.Data[:index], c.Data[index+1:]...)
	return true
}

// Find all datasets where a key equals a value
func (c *Collection) FindByKeyValue(key string, value interface{}) []JSONData {
	var results []JSONData
	for _, jsonData := range c.Data {
		if v, ok := jsonData.Data[key]; ok && v == value {
			results = append(results, jsonData)
		}
	}
	return results
}

// Update dataset at index
func (c *Collection) Update(index int, newData JSONData) bool {
	if index < 0 || index >= len(c.Data) {
		return false
	}
	c.Data[index] = newData
	return true
}

func (c *Collection) DeleteAll() bool{
	c.Data = []JSONData{}
	return true
}
