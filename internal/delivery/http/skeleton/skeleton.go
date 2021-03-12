package skeleton

import (
	"io/ioutil"
	"encoding/json"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"go-skeleton/pkg/response"
	testingEntity "go-skeleton/internal/entity/testing"
)

// ISkeletonSvc is an interface to Skeleton Service
type ISkeletonSvc interface {
	GetAllData(ctx context.Context) ([]testingEntity.Testing, error)
	GetDataByID(ctx context.Context, ID string) (testingEntity.Testing, error)
	GetDataByAge(ctx context.Context, Age string) (testingEntity.Testing, error)
	GetDataByBalance(ctx context.Context, Balance string) (testingEntity.Testing, error)
	InsertDataUser(ctx context.Context, singleTesting testingEntity.Testing) error
	UpdateDataUser(ctx context.Context, singleTesting testingEntity.Testing, ID string) (testingEntity.Testing, error)
	DeleteDataUser(ctx context.Context, singleTesting testingEntity.Testing, ID string) error
}

type (
	// Handler ...
	Handler struct {
		skeletonSvc ISkeletonSvc
	}
)

// New for bridging product handler initialization
func New(is ISkeletonSvc) *Handler {
	return &Handler{
		skeletonSvc: is,
	}
}

// SkeletonHandler will receive request and return response
func (h *Handler) SkeletonHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp     *response.Response
		result   interface{}
		metadata interface{}
		err      error
		errRes   response.Error
	)

	resp = &response.Response{}
	body, _ := ioutil.ReadAll(r.Body)
	defer resp.RenderJSON(w, r)


	switch r.Method {
	// Check if request method is GET
	case http.MethodGet:
		var _type string
		if _, getOK := r.URL.Query()["typeGet"]; getOK{
			_type = r.FormValue("typeGet")
		}
		switch _type{
			case "GetAllData":
				result, err = h.skeletonSvc.GetAllData(context.Background())
			case "GetDataByID":
				result, err = h.skeletonSvc.GetDataByID(context.Background(), r.FormValue("ID"))
			case "GetDataByAge":
				result, err = h.skeletonSvc.GetDataByAge(context.Background(), r.FormValue("Age"))
			case "GetDataByBalance":
				result, err = h.skeletonSvc.GetDataByBalance(context.Background(), r.FormValue("Balance"))	
			}
	// Check if request method is POST
	case http.MethodPost:
		var _type string
		if _, postOK := r.URL.Query()["typePost"]; postOK{
			_type = r.FormValue("typePost")
		}
		switch _type{
			case "InsertDataUser":
				var dataUser testingEntity.Testing
				json.Unmarshal(body, &dataUser)
				err = h.skeletonSvc.InsertDataUser(context.Background(), dataUser)	
		}
	
	// Check if request method is PUT
	case http.MethodPut:
		var _type string
		if _, putOK := r.URL.Query()["typePut"]; putOK{
			_type = r.FormValue("typePut")
		}
		switch _type{
			case "UpdateDataUser":
				var dataUser testingEntity.Testing
				json.Unmarshal(body, &dataUser)
				result, err = h.skeletonSvc.UpdateDataUser(context.Background(),dataUser, r.FormValue("ID"))	
		}

	// Check if request method is DELETE
	case http.MethodDelete:
		var _type string
		if _, deleteOK := r.URL.Query()["typeDelete"]; deleteOK{
			_type = r.FormValue("typeDelete")
		}
		switch _type{
			case "DeleteDataUser":
				var dataUser testingEntity.Testing
				json.Unmarshal(body, &dataUser)
				err = h.skeletonSvc.DeleteDataUser(context.Background(),dataUser, r.FormValue("ID"))	
		}

	default:
		err = errors.New("400")
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "101 - Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   500,
				Msg:    "500 - Internal Server Error",
				Status: true,
			}
		}
		// If error 401
		if strings.Contains(err.Error(), "401") {
			// Replace error with server error
			errRes = response.Error{
				Code:   401,
				Msg:    "401 - Unauthorized",
				Status: true,
			}
		}
		// If error 403
		if strings.Contains(err.Error(), "403") {
			// Replace error with server error
			errRes = response.Error{
				Code:   403,
				Msg:    "403 - Forbidden",
				Status: true,
			}
		}
		// If error 400
		if strings.Contains(err.Error(), "400") {
			// Replace error with server error
			errRes = response.Error{
				Code:   400,
				Msg:    "400 - Bad Request",
				Status: true,
			}
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = errRes.Code
		resp.Error = errRes
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
