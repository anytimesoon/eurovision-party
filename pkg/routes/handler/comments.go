package handler

// import (
// 	"encoding/json"
// 	"eurovision/pkg/dao"
// 	"eurovision/pkg/dto"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func FindAllComments(writer http.ResponseWriter, req *http.Request) {
// 	commentsDAO, err := dao.Comments()
// 	if err != nil {
// 		log.Println("FAILED to find all comments!")
// 		return
// 	}

// 	commentsDTO := dto.Comments{
// 		Success: true,
// 		Message: "",
// 	}

// 	for _, comment := range commentsDAO {
// 		commentsDTO.Data = append(commentsDTO.Data, dto.CommentData{UUID: comment.UUID, UserId: comment.UserId, Text: comment.Text})
// 	}

// 	json.NewEncoder(writer).Encode(commentsDTO)
// }

// func CreateComment(writer http.ResponseWriter, req *http.Request) {
// 	var commentDTO dto.Comment

// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		log.Println("FAILED to read body of USER CREATE!", err)
// 		return
// 	}

// 	err = json.Unmarshal(body, &commentDTO)
// 	if err != nil {
// 		log.Printf("%+v", string(body))
// 		log.Printf("%+v", commentDTO)
// 		log.Println("FAILED to unmarshal json!", err)
// 		return
// 	}

// 	commentDAO, err := dao.CreateComment(commentDTO)
// 	if err != nil {
// 		log.Println("FAILED to create new comment", err)
// 		return
// 	}

// 	commentDTO = dto.Comment{
// 		Success: true,
// 		Message: "",
// 		Data:    dto.CommentData{UUID: commentDAO.UUID, UserId: commentDAO.UserId, Text: commentDAO.Text, CreatedAt: commentDAO.CreatedAt},
// 	}

// 	json.NewEncoder(writer).Encode(commentDTO)
// }
