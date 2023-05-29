package main
import (
    "encoding/json"
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
    
    fmt.Printf(" \n\n %v: %T \n %v: %T \n %v: %T",x,x,y,y,z,z)
}