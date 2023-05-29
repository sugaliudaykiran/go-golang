package main
import (
    "encoding/json"
    "fmt"
)

type Bird struct{
    Species string
    Description string
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
}