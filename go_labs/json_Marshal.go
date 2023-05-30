package main 
import (
    "fmt"
    "encoding/json"
)
// The same json tags will be used to encode data into JSON
type BirdC struct{
    Species string `json:"birdType"`
    Description string `json:"what it does"` 
    // Description string `json:"what it does,omitempty"` // here "omitempty" use in case of empty field.
    
    // If we want to always ignore a field, we can use the json:"-" struct tag to denote that we never want this field included:
    // Species string `json:"-"`
}

func main(){
    // Marshaling Structured Data
    eagle:= &BirdC{
        Species: "Eagles",
        Description: "they fly over the Himalayas",
    }
    // data will in the []byte format..
    data,_ := json.Marshal(eagle) // _ is ignoring the error part.
    fmt.Println("Enoding to Json: ",string(data))
    
    // Ignoring Empty Fields
    // In some cases, we would want to ignore a field in our JSON output, if its value is empty.
    // We can use the “omitempty” property for this purpose.
    // For example, if the Description field is missing for the pigeon object, 
    // the key will not appear in the encoded JSON string incase we set this property:
    pigeon:=&BirdC{
        Species: "Pigeons",
        // Here Description is empty
    }
    data2,_ := json.Marshal(pigeon) // data2 will in the []byte format..
    fmt.Println("Skipping the Description: ", string(data2))
    
    // If we want to always ignore a field, we can use the json:"-" struct tag to denote that we never want this field included:
    duck:=&BirdC{
        Species: "Ducks",
        Description: "bird which can swim and look beautiful",
    }
    data3,_ := json.Marshal(duck)
    fmt.Println("ignore a field (-): ", string(data3))
    
    // Marshaling Slices
    // This isn’t much different from structs. We just need to pass the "slice or array" to the json.Marshal function, and it will encode data like you expect:
    chicken:=&BirdC{
        Species: "Chickens",
        Description: "humans eats more the Species",
    }
    
    data4,_:=json.Marshal([]*BirdC{chicken,chicken,duck,pigeon})
    fmt.Println(string(data4))
    
    // Marshaling Maps - We can use maps to encode unstructured data.
    // The keys of the map need to be strings, or a type that can convert to strings. The values can be any serializable type.
    birdData:= map[string]interface{}{  // interface{} => not fix the type of it.
        "birdSound": map[string]interface{}{ // "any" ==> interface{}
            "pigeon": "coo",
            "eagle": "qee",
            "duck": "wqak",
        },
        "total birdCount": 3,
    }
    
    // JSON encoding is done the same way as before	
    data5,_ := json.Marshal(birdData)
    fmt.Println("\n\n Map data: ",string(data5))
	
	// Best Practices (Structs vs Maps)
// As a general rule of thumb, if you can use structs to represent your JSON data, you should use them. The only good reason to use maps would be if it were not possible to use structs due to the uncertain nature of the keys or values in the data.

// If we use maps, we will either need each of the keys to have the same data type, or use a generic type and convert it later.
}