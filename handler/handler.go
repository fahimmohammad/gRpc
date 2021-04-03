package handler

import (
	"context"
	"fmt"

	pb "github.com/fahimsGit/gRpctest/proto"
	repo "github.com/fahimsGit/gRpctest/repo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

type Article struct {
	Id        string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	Publisher string `json:"publisher" bson:"publisher"`
}
type Server struct {
	Db *repo.DBSession
}

func NewServer() Server {
	return Server{
		Db: repo.NewDB(),
	}
}
func generateUUID() string {
	return uuid.New().String()
}

func (s *Server) CreateArticle(ctx context.Context, req *pb.CreateArticleReq) (*pb.CreateArticleResp, error) {
	article := Article{
		Id:        generateUUID(),
		Name:      req.Name,
		Publisher: req.PublisherName,
	}
	coll := s.Db.DbSession.DB(s.Db.DbName).C(s.Db.TableName)
	err := coll.Insert(&article)
	if err != nil {
		return &pb.CreateArticleResp{
			ID:            "",
			Name:          "",
			PublisherName: "",
			Error:         "",
		}, err
	}
	return &pb.CreateArticleResp{
		ID:            "",
		Name:          article.Name,
		PublisherName: article.Publisher,
		Error:         "nil",
	}, nil
}

func (s *Server) GetArticle(ctx context.Context, req *pb.GetArticleReq) (*pb.GetArticleResp, error) {
	// var articleResp *pb.GetArticleResp
	articleResp := &pb.GetArticleResp{}
	var article Article
	// article.Id = req.Id
	coll := s.Db.DbSession.DB(s.Db.DbName).C(s.Db.TableName)
	err := coll.Find(bson.M{"id": req.Id}).One(&article)
	fmt.Println(article)
	if err != nil {
		articleResp.Error = err.Error()
		return articleResp, err
	}

	return &pb.GetArticleResp{
		ID:            article.Id,
		Name:          article.Name,
		PublisherName: article.Publisher,
		Error:         "nil",
	}, nil
}
