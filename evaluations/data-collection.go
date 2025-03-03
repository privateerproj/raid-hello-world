package evaluations

type CustomDataObject struct {
    yourData string
}

func LoadData() *CustomDataObject {
    return &CustomDataObject{
        yourData: "is here",
    }
}