package main
import (
    "encoding/json" // Unmarshal([]byte(json_var), &var) ->decoding
    "fmt"
)

type Bird struct{
    Species string
    Description string
}

type Dimension struct{
    Height int
    Weight int
}

// the structure of the object in our Go code. To add a nested dimensions object, lets create a dimensions struct :

type Birds struct{
    Species string
    Description string
    Dimensions Dimension
}

// JSON Struct Tags - Custom Field Names
type BirdC struct{
    Species string `json:"birdType"`
    Description string `json:"what it does"`
}

func main(){
    
    // structured data --> refers to data where you know the format beforehand.
    birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
    
    var bird Bird // type of struct
    err := json.Unmarshal([]byte(birdJson), &bird)
    if err==nil{
        fmt.Println("bird species:", bird.Species,"\n","bird description: ",bird.Description)
    }
    
    // Json Array
    birdJson2 := `[
      {
        "species": "pigeon",
        "description": "likes to perch on rocks"
      },
      {
        "species":"eagle",
        "description":"bird of prey"
      }
    ]`
    
    var birdList []Bird
    errList := json.Unmarshal([]byte(birdJson2), &birdList)
    if errList==nil{
        for _,map_val := range birdList{
            fmt.Printf("Birds: %v, description: %v \n", map_val.Species, map_val.Description)
        }
        // fmt.Printf("Birds : %+v", birdList)
        // fmt.Printf("Birds: %v, description: %v", birdList[0].Species, birdList[0].Description)    
    }
    
    // Nested Objects := Now, consider the case when you have a property called Dimensions, that measures the "Height and Length" of the bird in question:
    birdJson3 := `{
      "species": "pigeon",
      "decription": "likes to perch on rocks",
      "dimensions": {
         "height": 24,
          "width": 10
        }
    }`
    
    var birdNested Birds
    errNest := json.Unmarshal([]byte(birdJson3), &birdNested)
    if errNest==nil{
        fmt.Printf("BIRD: %+v",birdNested)
    }
    
    // We can unmarshal these values to their corresponding data type in Go by using primitive types:
    numberJson := "3"
    floatJson := "3.1412"
    stringJson := `"bird"`
    
    var x int
    var y float64
    var z string
    
    json.Unmarshal([]byte(numberJson), &x)
    json.Unmarshal([]byte(floatJson), &y)
    json.Unmarshal([]byte(stringJson), &z)
    
    fmt.Printf(" \n\n %v: %T \n %v: %T \n %v: %T \n\n",x,x,y,y,z,z)
    
    // Decoding JSON to Maps - Unstructured Data
// If you don’t know the structure of your JSON properties beforehand, you cannot use structs to unmarshal your data.
    
    Json4:=`{
        "birds": {
            "pigeon":"likes to perch on rocks",
            "eagle":"bird of prey"
        },
        "animals": "none"
    }`
    // no struct we can build to represent the above data 
    // we create a map of strings to empty interfaces:
    // var any interface{}  (or) use the "interface{}"
    var animalsMap map[string]interface{}
    json.Unmarshal([]byte(Json4), &animalsMap)
    // The object stored in the "birds" key is also stored as a map[string]any type, and its type is asserted from
    // the `any` type
    birdMap:=animalsMap["birds"].(map[string]interface{}) // "type assertion" of birdMap
    // fmt.Printf("birdMap: %+v \n",birdMap)
    for mapKey,mapVal := range birdMap{
        fmt.Printf("%v: %v \n", mapKey, mapVal)
    }
    
    // JSON Struct Tags - Custom Field Names
    Json5:=`{
        "birdType": "pigeon",
        "what it does": "likes to perch on rocks"
    }`
    
    var customBird BirdC 
    errC := json.Unmarshal([]byte(Json5), &customBird)
    if errC==nil{
        fmt.Println("custom bird: ",customBird)
    }
    
    
    // Validating the JSON data 
    checkJson:=`{
        "birds": {
            "pigeon":"likes to perch on rocks",
            "eagle":"bird of prey"
    `
    if !json.Valid([]byte(checkJson)){ // using json.Valid() function.
        fmt.Println("Invalid JSON format", checkJson)
        return
    }
    
    var birdJson4 map[string]interface{}
    errJson4 := json.Unmarshal([]byte(checkJson), &birdJson4)
    
    if errJson4==nil{
        fmt.Println("Hello")
        fmt.Printf("%+v \n", birdJson4)
    }

}