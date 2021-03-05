package skeleton

import (
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
			}
	// Check if request method is POST
	case http.MethodPost:
	
	// Check if request method is PUT
	case http.MethodPut:

	// Check if request method is DELETE
	case http.MethodDelete:

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
