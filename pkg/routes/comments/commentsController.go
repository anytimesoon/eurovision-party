package comments

import (
	"encoding/json"
	"eurovision/pkg/dao"
	"eurovision/pkg/dto"
	"log"
	"net/http"
)

func All(writer http.ResponseWriter, req *http.Request) {
	commentsDAO, err := dao.Comments()
	if err != nil {
		log.Println("FAILED to find all comments!")
		return
	}

	commentsDTO := dto.Comments{
		Success: true,
		Message: "",
	}

	for _, comment := range commentsDAO {
		commentsDTO.Data = append(commentsDTO.Data, dto.CommentData{UUID: comment.UUID, UserId: comment.UserId, Text: comment.Text})
	}

	json.NewEncoder(writer).Encode(commentsDTO)
}
