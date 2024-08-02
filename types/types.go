package types

type User struct{
    ID          string  `json:"id" bson:"_id"`  
    Email       string  `json:"email" bson:"email"`
    Password    string  `json:"password" bson:"password"`
    Name        string  `json:"name" bson:"name"`
    LastName    string  `json:"lastName" bson:"lastName"`
}

type Blog struct{
    UserId  string  `json:"userId" bson:"userId"`
    ID      string  `json:"id" bson:"_id"`
    Content string  `json:"content" bson:"content"`
    Title   string  `json:"title" bson:"title"`
}
